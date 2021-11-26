package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	ILocationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
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

func (l *locationManager) AddLocation(data *[]byte) (respondLocationModel *model.LocationResponseModel, s bool, m string) {

	firstmodel := model.LocationModel{}
	err := (*l.IJsonParser).DecodeJson(data, &firstmodel)
	if err != nil {
		return &model.LocationResponseModel{}, false, err.Error()
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

	locerr := (*l.ILocationDal).Add(&modelResponse)
	if locerr != nil {
		return &modelResponse, false, locerr.Error()
	}
	return &modelResponse, true, ""
}
