package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	ILocationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/logger"
)

type locationManager struct {
	ICacheService *ICacheService.ICacheService
	IJsonParser   *Ijsonparser.IJsonParser
	ILocationDal  *ILocationDal.ILocationDal
	ILog          *logger.ILog
}

func LocationManagerConstructor() *locationManager {
	return &locationManager{
		ICacheService: &IoC.CacheService,
		IJsonParser:   &IoC.JsonParser,
		ILocationDal:  &IoC.LocationDal,
		ILog:          &IoC.Logger,
	}
}

func (l *locationManager) AddLocation(data *[]byte) (v interface{}, s bool, m string) {

	firstmodel := model.LocationModel{}
	Err := (*l.IJsonParser).DecodeJson(data, &firstmodel)
	if Err != nil {
		(*l.ILog).SendErrorLog("LocationManager", "AddLocation",
			"byte array to LocationModel", "Json Parser Decode Err: ", Err.Error())
		return &model.LocationResponseModel{}, false, Err.Error()
	}
	modelResponse := model.LocationResponseModel{}
	modelResponse.ProjectId = firstmodel.ProjectId
	modelResponse.ClientId = firstmodel.ClientId
	modelResponse.CustomerId = firstmodel.CustomerId
	modelResponse.Region, _, _ = (*l.ICacheService).ManageCache("Region", firstmodel.Region)
	modelResponse.Country, _, _ = (*l.ICacheService).ManageCache("Country", firstmodel.Country)
	modelResponse.Org, _, _ = (*l.ICacheService).ManageCache("Org", firstmodel.Org)
	modelResponse.City, _, _ = (*l.ICacheService).ManageCache("City", firstmodel.City)
	modelResponse.Continent, _, _ = (*l.ICacheService).ManageCache("Continent", firstmodel.Continent)
	defer (*l.ILog).SendInfoLog("LocationManager", "AddLocation",
		modelResponse.ClientId, modelResponse.ProjectId)
	logErr := (*l.ILocationDal).Add(&modelResponse)
	if logErr != nil {
		(*l.ILog).SendErrorLog("LocationManager", "AddLocation",
			"LocationDal_Add", logErr.Error())
		return &modelResponse, false, logErr.Error()
	}
	return &modelResponse, true, ""
}
