package apisrv

import (
	"crypto/tls"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
)

var testServer *httptest.Server
var server *Server

func TestMain(m *testing.M) {
	common.InitConfig()
	common.SetLogLevel()
	dbConfig := viper.GetStringMap("test_database")
	for _, iConfig := range dbConfig {
		config := common.InterfaceToInterfaceMap(iConfig)
		viper.Set("database.type", config["type"])
		viper.Set("database.connection", config["connection"])
		viper.Set("database.dialect", config["dialect"])
		RunTestForDB(m)
	}
}

func RunTestForDB(m *testing.M) {
	var err error
	server, err = NewServer()
	if err != nil {
		log.Fatal(err)
	}
	testServer = httptest.NewUnstartedServer(server.Echo)
	testServer.TLS = new(tls.Config)
	testServer.TLS.NextProtos = append(testServer.TLS.NextProtos, "h2")
	testServer.StartTLS()
	defer testServer.Close()

	viper.Set("keystone.authurl", testServer.URL+"/v3")
	err = server.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()
	log.Info("starting test")
	code := m.Run()
	log.Info("finished test")
	if code != 0 {
		os.Exit(code)
	}
}

func RunTest(t *testing.T, file string) {
	testData, err := LoadTest(file)
	assert.NoError(t, err, "failed to load test data")
	clients := map[string]*Client{}

	for key, client := range testData.Clients {
		//Rewrite endpoint for test server
		client.Endpoint = testServer.URL
		client.AuthURL = testServer.URL + "/v3"
		client.InSecure = true
		client.Init()

		clients[key] = client

		err := clients[key].Login()
		assert.NoError(t, err, "client failed to login")
	}
	for _, cleanTask := range testData.Cleanup {
		fmt.Println(cleanTask)
		clientID := cleanTask["client"]
		if clientID == "" {
			clientID = "default"
		}
		client := clients[clientID]
		// delete existing resources.
		log.Debug(cleanTask["path"])
		_, err := client.Delete(cleanTask["path"], nil) // nolint
		if err != nil {
			log.Debug(err)
		}
	}
	for _, task := range testData.Workflow {
		log.Debug("[Step] ", task.Name)
		task.Request.Data = common.YAMLtoJSONCompat(task.Request.Data)
		clientID := "default"
		if task.Client != "" {
			clientID = task.Client
		}
		client := clients[clientID]
		_, err := client.DoRequest(task.Request)
		assert.NoError(t, err, fmt.Sprintf("task %v failed", task))
		task.Expect = common.YAMLtoJSONCompat(task.Expect)
		ok := common.AssertEqual(t, task.Expect, task.Request.Output, fmt.Sprintf("task %v failed", task))
		if !ok {
			break
		}
	}
}

func LoadTest(file string) (*TestScenario, error) {
	var testScenario TestScenario
	err := common.LoadFile(file, &testScenario)
	return &testScenario, err
}

type Task struct {
	Name    string      `yaml:"name"`
	Client  string      `yaml:"client"`
	Request *Request    `yaml:"request"`
	Expect  interface{} `yaml:"expect"`
}

type TestScenario struct {
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Tables      []string            `yaml:"tables"`
	Clients     map[string]*Client  `yaml:"clients"`
	Cleanup     []map[string]string `yaml:"cleanup"`
	Workflow    []*Task             `yaml:"workflow"`
}
