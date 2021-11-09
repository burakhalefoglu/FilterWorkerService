package concrete

import (
	model2 "FilterWorkerService/internal/model"
	ILocationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	Ijsonparser "FilterWorkerService/pkg/jsonParser"
)

type LocationManager struct {
	ICacheService ICacheService.ICacheService
	IJsonParser   Ijsonparser.IJsonParser
	ILocationDal  ILocationDal.ILocationDal
}

func (l *LocationManager) AddLocation(data *[]byte) (s bool, m string) {

	model := model2.LocationModel{}
	err := l.IJsonParser.DecodeJson(data, &model)
	if err != nil {
		return false, err.Error()
	}

	modelResponse := model2.LocationResponseModel{}
	modelResponse.ProjectId = model.ProjectId
	modelResponse.ClientId = model.ClientId
	modelResponse.CustomerId = model.CustomerId
	modelResponse.Region, _, _ = l.ICacheService.ManageCache("Region", model.Region)
	modelResponse.Country, _, _ = l.ICacheService.ManageCache("Country", model.Country)
	modelResponse.Org, _, _ = l.ICacheService.ManageCache("Org", model.Org)
	modelResponse.City, _, _ = l.ICacheService.ManageCache("City", model.City)
	modelResponse.Continent, _, _ = l.ICacheService.ManageCache("Continent", model.Continent)

	locerr := l.ILocationDal.Add(&modelResponse)
	if locerr != nil {
		return false, locerr.Error()
	}
	return true, ""
}
