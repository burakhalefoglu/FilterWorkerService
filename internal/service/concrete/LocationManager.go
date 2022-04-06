package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	ILocationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type locationManager struct {
	ICacheService *ICacheService.ICacheService
	IJsonParser   *Ijsonparser.IJsonParser
	ILocationDal  *ILocationDal.ILocationDal
}

func LocationManagerConstructor() *locationManager {
	return &locationManager{
		ICacheService: &IoC.CacheService,
		IJsonParser:   &IoC.JsonParser,
		ILocationDal:  &IoC.LocationDal,
	}
}

func (l *locationManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {

	firstmodel := model.LocationModel{}
	convertErr := (*l.IJsonParser).DecodeJson(data, &firstmodel)
	if convertErr != nil {
		clogger.Error(&map[string]interface{}{"Byte array to LocationModel  LocationManager Json Parser Decode ERROR: ": convertErr.Error()})
		return false, convertErr.Error()
	}
	modelResponse := model.LocationResponseModel{}
	modelResponse.Id = firstmodel.Id
	modelResponse.ProjectId = firstmodel.ProjectId
	modelResponse.ClientId = firstmodel.ClientId
	modelResponse.CustomerId = firstmodel.CustomerId
	modelResponse.Region, _, _ = (*l.ICacheService).ManageCache("Region", firstmodel.Region)
	modelResponse.Country, _, _ = (*l.ICacheService).ManageCache("Country", firstmodel.Country)
	modelResponse.Org, _, _ = (*l.ICacheService).ManageCache("Org", firstmodel.Org)
	modelResponse.City, _, _ = (*l.ICacheService).ManageCache("City", firstmodel.City)
	modelResponse.Continent, _, _ = (*l.ICacheService).ManageCache("Continent", firstmodel.Continent)
	// defer log.Print("LocationManager", "AddLocation",
	// 	modelResponse.ClientId, modelResponse.ProjectId)

	oldModel, err := (*l.ILocationDal).GetById(modelResponse.ClientId, modelResponse.ProjectId)

	switch {

	case err != nil && err.Error() != "not found":
		clogger.Error(&map[string]interface{}{
			fmt.Sprintf("Get clientId: %d, projectId: %d location_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): err.Error(),
		})

	case err != nil && err.Error() == "not found":

		logErr := (*l.ILocationDal).Add(&modelResponse)
		if logErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Add clientId: %d, projectId: %d location_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): logErr.Error(),
			})
			return false, logErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Add clientId: %d, projectId: %d location_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return true, "Added"

	case err == nil:
		_, updateResult, updateErr := l.UpdateLocation(&modelResponse, oldModel)
		if updateErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Update clientId: %d, projectId: %d location_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): updateErr.Error(),
			})
			return updateResult, updateErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Update clientId: %d, projectId: %d location_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return updateResult, "Updated"

	default:

		return false, ""

	}
	return false, ""
}

func (l *locationManager) UpdateLocation(modelResponse *model.LocationResponseModel, oldModel *model.LocationResponseModel) (updatedModel *model.LocationResponseModel, s bool, m error) {
	oldModel.Id = modelResponse.Id
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.Region = modelResponse.Region
	oldModel.Country = modelResponse.Country
	oldModel.Org = modelResponse.Org
	oldModel.City = modelResponse.City
	oldModel.Continent = modelResponse.Continent

	logErr := (*l.ILocationDal).UpdateById(oldModel.ClientId, oldModel.ProjectId, oldModel)
	if logErr != nil {
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}
