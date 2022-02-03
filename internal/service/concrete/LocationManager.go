package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	ILocationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
	"log"
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

func (l *locationManager) AddLocation(data *[]byte) (v interface{}, s bool, m string) {

	firstmodel := model.LocationModel{}
	convertErr := (*l.IJsonParser).DecodeJson(data, &firstmodel)
	if convertErr != nil {
		log.Fatal("LocationManager", "AddLocation",
			"byte array to LocationModel", "Json Parser Decode Err: ", convertErr.Error())
		return &model.LocationResponseModel{}, false, convertErr.Error()
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
	defer log.Print("LocationManager", "AddLocation",
		modelResponse.ClientId, modelResponse.ProjectId)
	logErr := (*l.ILocationDal).Add(&modelResponse)
	if logErr != nil {
		log.Fatal("LocationManager", "AddLocation",
			"LocationDal_Add", logErr.Error())
		return &modelResponse, false, logErr.Error()
	}
	return &modelResponse, true, ""
}
