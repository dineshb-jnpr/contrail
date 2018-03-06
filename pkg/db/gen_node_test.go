// nolint
package db

import (
	"context"
	"testing"
	"time"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/pkg/errors"
)

//For skip import error.
var _ = errors.New("")

func TestNode(t *testing.T) {
	// t.Parallel()
	db := &DB{
		DB:      testDB,
		Dialect: NewDialect("mysql"),
	}
	db.initQueryBuilders()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mutexMetadata := common.UseTable(db.DB, "metadata")
	mutexTable := common.UseTable(db.DB, "node")
	// mutexProject := UseTable(db.DB, "node")
	defer func() {
		mutexTable.Unlock()
		mutexMetadata.Unlock()
		if p := recover(); p != nil {
			panic(p)
		}
	}()
	model := models.MakeNode()
	model.UUID = "node_dummy_uuid"
	model.FQName = []string{"default", "default-domain", "node_dummy"}
	model.Perms2.Owner = "admin"
	var err error

	// Create referred objects

	var Keypaircreateref []*models.NodeKeypairRef
	var KeypairrefModel *models.Keypair
	KeypairrefModel = models.MakeKeypair()
	KeypairrefModel.UUID = "node_keypair_ref_uuid"
	KeypairrefModel.FQName = []string{"test", "node_keypair_ref_uuid"}
	_, err = db.CreateKeypair(ctx, &models.CreateKeypairRequest{
		Keypair: KeypairrefModel,
	})
	KeypairrefModel.UUID = "node_keypair_ref_uuid1"
	KeypairrefModel.FQName = []string{"test", "node_keypair_ref_uuid1"}
	_, err = db.CreateKeypair(ctx, &models.CreateKeypairRequest{
		Keypair: KeypairrefModel,
	})
	KeypairrefModel.UUID = "node_keypair_ref_uuid2"
	KeypairrefModel.FQName = []string{"test", "node_keypair_ref_uuid2"}
	_, err = db.CreateKeypair(ctx, &models.CreateKeypairRequest{
		Keypair: KeypairrefModel,
	})
	if err != nil {
		t.Fatal("ref create failed", err)
	}
	Keypaircreateref = append(Keypaircreateref, &models.NodeKeypairRef{UUID: "node_keypair_ref_uuid", To: []string{"test", "node_keypair_ref_uuid"}})
	Keypaircreateref = append(Keypaircreateref, &models.NodeKeypairRef{UUID: "node_keypair_ref_uuid2", To: []string{"test", "node_keypair_ref_uuid2"}})
	model.KeypairRefs = Keypaircreateref

	//create project to which resource is shared
	projectModel := models.MakeProject()
	projectModel.UUID = "node_admin_project_uuid"
	projectModel.FQName = []string{"default-domain-test", "admin-test"}
	projectModel.Perms2.Owner = "admin"
	var createShare []*models.ShareType
	createShare = append(createShare, &models.ShareType{Tenant: "default-domain-test:admin-test", TenantAccess: 7})
	model.Perms2.Share = createShare

	_, err = db.CreateProject(ctx, &models.CreateProjectRequest{
		Project: projectModel,
	})
	if err != nil {
		t.Fatal("project create failed", err)
	}

	//    //populate update map
	//    updateMap := map[string]interface{}{}
	//
	//
	//    common.SetValueByPath(updateMap, ".UUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Username", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Type", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".SSHKey", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PrivateMachineState", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".PrivateMachineProperties", ".", "test")
	//
	//
	//
	//    if ".Perms2.Share" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Perms2.Share", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Perms2.GlobalAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Password", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentUUID", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".ParentType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".MacAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IPAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.UserVisible", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OwnerAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Owner", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.OtherAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.GroupAccess", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Permissions.Group", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.LastModified", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Enable", ".", true)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Description", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Creator", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".IDPerms.Created", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".Hostname", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GCPMachineType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".GCPImage", ".", "test")
	//
	//
	//
	//    if ".FQName" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".FQName", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".FQName", ".", `{"test": "test"}`)
	//    }
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.IpmiUsername", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.IpmiPassword", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.IpmiAddress", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.DeployRamdisk", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DriverInfo.DeployKernel", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".DisplayName", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.MemoryMB", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.DiskGB", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.CPUCount", ".", 1.0)
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".BMProperties.CPUArch", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".AwsInstanceType", ".", "test")
	//
	//
	//
	//    common.SetValueByPath(updateMap, ".AwsAmi", ".", "test")
	//
	//
	//
	//    if ".Annotations.KeyValuePair" == ".Perms2.Share" {
	//        var share []interface{}
	//        share = append(share, map[string]interface{}{"tenant":"default-domain-test:admin-test", "tenant_access":7})
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", share)
	//    } else {
	//        common.SetValueByPath(updateMap, ".Annotations.KeyValuePair", ".", `{"test": "test"}`)
	//    }
	//
	//
	//    common.SetValueByPath(updateMap, "uuid", ".", "node_dummy_uuid")
	//    common.SetValueByPath(updateMap, "fq_name", ".", []string{"default", "default-domain", "access_control_list_dummy"})
	//    common.SetValueByPath(updateMap, "perms2.owner", ".", "admin")
	//
	//    // Create Attr values for testing ref update(ADD,UPDATE,DELETE)
	//
	//    var Keypairref []interface{}
	//    Keypairref = append(Keypairref, map[string]interface{}{"operation":"delete", "uuid":"node_keypair_ref_uuid", "to": []string{"test", "node_keypair_ref_uuid"}})
	//    Keypairref = append(Keypairref, map[string]interface{}{"operation":"add", "uuid":"node_keypair_ref_uuid1", "to": []string{"test", "node_keypair_ref_uuid1"}})
	//
	//
	//
	//    common.SetValueByPath(updateMap, "KeypairRefs", ".", Keypairref)
	//
	//
	_, err = db.CreateNode(ctx,
		&models.CreateNodeRequest{
			Node: model,
		})

	if err != nil {
		t.Fatal("create failed", err)
	}

	//    err = common.DoInTransaction(db, func (tx *sql.Tx) error {
	//        return UpdateNode(tx, model.UUID, updateMap)
	//    })
	//    if err != nil {
	//        t.Fatal("update failed", err)
	//    }

	//Delete ref entries, referred objects

	err = DoInTransaction(ctx, db.DB, func(ctx context.Context) error {
		tx := GetTransaction(ctx)
		stmt, err := tx.Prepare("delete from `ref_node_keypair` where `from` = ? AND `to` = ?;")
		if err != nil {
			return errors.Wrap(err, "preparing KeypairRefs delete statement failed")
		}
		_, err = stmt.Exec("node_dummy_uuid", "node_keypair_ref_uuid")
		_, err = stmt.Exec("node_dummy_uuid", "node_keypair_ref_uuid1")
		_, err = stmt.Exec("node_dummy_uuid", "node_keypair_ref_uuid2")
		if err != nil {
			return errors.Wrap(err, "KeypairRefs delete failed")
		}
		return nil
	})
	_, err = db.DeleteKeypair(ctx,
		&models.DeleteKeypairRequest{
			ID: "node_keypair_ref_uuid"})
	if err != nil {
		t.Fatal("delete ref node_keypair_ref_uuid  failed", err)
	}
	_, err = db.DeleteKeypair(ctx,
		&models.DeleteKeypairRequest{
			ID: "node_keypair_ref_uuid1"})
	if err != nil {
		t.Fatal("delete ref node_keypair_ref_uuid1  failed", err)
	}
	_, err = db.DeleteKeypair(
		ctx,
		&models.DeleteKeypairRequest{
			ID: "node_keypair_ref_uuid2",
		})
	if err != nil {
		t.Fatal("delete ref node_keypair_ref_uuid2 failed", err)
	}

	//Delete the project created for sharing
	_, err = db.DeleteProject(ctx, &models.DeleteProjectRequest{
		ID: projectModel.UUID})
	if err != nil {
		t.Fatal("delete project failed", err)
	}

	response, err := db.ListNode(ctx, &models.ListNodeRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Nodes) != 1 {
		t.Fatal("expected one element", err)
	}

	ctxDemo := context.WithValue(ctx, "auth", common.NewAuthContext("default", "demo", "demo", []string{}))
	_, err = db.DeleteNode(ctxDemo,
		&models.DeleteNodeRequest{
			ID: model.UUID},
	)
	if err == nil {
		t.Fatal("auth failed")
	}

	_, err = db.CreateNode(ctx,
		&models.CreateNodeRequest{
			Node: model})
	if err == nil {
		t.Fatal("Raise Error On Duplicate Create failed", err)
	}

	_, err = db.DeleteNode(ctx,
		&models.DeleteNodeRequest{
			ID: model.UUID})
	if err != nil {
		t.Fatal("delete failed", err)
	}

	response, err = db.ListNode(ctx, &models.ListNodeRequest{
		Spec: &models.ListSpec{Limit: 1}})
	if err != nil {
		t.Fatal("list failed", err)
	}
	if len(response.Nodes) != 0 {
		t.Fatal("expected no element", err)
	}
	return
}
