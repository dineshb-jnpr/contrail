package services

import (
	"context"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//RESTCreatePort handle a Create REST service.
// nolint
func (service *ContrailService) RESTCreatePort(c echo.Context) error {
	requestData := &models.CreatePortRequest{}
	if err := c.Bind(requestData); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "port",
		}).Debug("bind failed on create")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.CreatePort(ctx, requestData)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusCreated, response)
}

//CreatePort handle a Create API
// nolint
func (service *ContrailService) CreatePort(
	ctx context.Context,
	request *models.CreatePortRequest) (*models.CreatePortResponse, error) {
	model := request.Port
	if model.UUID == "" {
		model.UUID = uuid.NewV4().String()
	}
	auth := common.GetAuthCTX(ctx)
	if auth == nil {
		return nil, common.ErrorUnauthenticated
	}

	if model.FQName == nil {
		if model.DisplayName != "" {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.DisplayName}
		} else {
			model.FQName = []string{auth.DomainID(), auth.ProjectID(), model.UUID}
		}
	}
	model.Perms2 = &models.PermType2{}
	model.Perms2.Owner = auth.ProjectID()

	return service.Next().CreatePort(ctx, request)
}

//RESTUpdatePort handles a REST Update request.
// nolint
func (service *ContrailService) RESTUpdatePort(c echo.Context) error {
	//id := c.Param("id")
	request := &models.UpdatePortRequest{}
	if err := c.Bind(request); err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"resource": "port",
		}).Debug("bind failed on update")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid JSON format")
	}
	ctx := c.Request().Context()
	response, err := service.UpdatePort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//UpdatePort handles a Update request.
// nolint
func (service *ContrailService) UpdatePort(
	ctx context.Context,
	request *models.UpdatePortRequest) (*models.UpdatePortResponse, error) {
	model := request.Port
	if model == nil {
		return nil, common.ErrorBadRequest("Update body is empty")
	}
	return service.Next().UpdatePort(ctx, request)
}

//RESTDeletePort delete a resource using REST service.
// nolint
func (service *ContrailService) RESTDeletePort(c echo.Context) error {
	id := c.Param("id")
	request := &models.DeletePortRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	_, err := service.DeletePort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusNoContent, nil)
}

//RESTGetPort a REST Get request.
// nolint
func (service *ContrailService) RESTGetPort(c echo.Context) error {
	id := c.Param("id")
	request := &models.GetPortRequest{
		ID: id,
	}
	ctx := c.Request().Context()
	response, err := service.GetPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}

//RESTListPort handles a List REST service Request.
// nolint
func (service *ContrailService) RESTListPort(c echo.Context) error {
	var err error
	spec := models.GetListSpec(c)
	request := &models.ListPortRequest{
		Spec: spec,
	}
	ctx := c.Request().Context()
	response, err := service.ListPort(ctx, request)
	if err != nil {
		return common.ToHTTPError(err)
	}
	return c.JSON(http.StatusOK, response)
}
