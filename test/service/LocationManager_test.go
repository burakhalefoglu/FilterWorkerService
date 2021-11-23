package test

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"FilterWorkerService/test/Mock/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var locationModel = model.LocationModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	Continent:  "Asia",
	Country:    "China",
	City:       "Pekin",
	Region:     "West",
	Org:        "ChinaTel",
}

var LocationResponseModel = model.LocationResponseModel{
	ProjectId:  locationModel.ProjectId,
	ClientId:   locationModel.ClientId,
	CustomerId: locationModel.CustomerId,
	Continent:  1,
	Country:    30,
	City:       187,
	Region:     93,
	Org:        8,
}

func Test_AddLocation_Success(t *testing.T) {
	var testLocationDal = new(repository.MockLocationDal)
	var testcache = new(service.MockCacheService)
	testcache.On("ManageCache", "Continent", locationModel.Continent).Return(int64(1), true, "")
	testcache.On("ManageCache", "Country", locationModel.Country).Return(int64(30), true, "")
	testcache.On("ManageCache", "City", locationModel.City).Return(int64(187), true, "")
	testcache.On("ManageCache", "Region", locationModel.Region).Return(int64(93), true, "")
	testcache.On("ManageCache", "Org", locationModel.Org).Return(int64(8), true, "")
	testLocationDal.On("Add", &LocationResponseModel).Return(nil)
	var manager = concrete.LocationManager{
		ICacheService: testcache,
		IJsonParser:   &gojson.GoJson{},
		ILocationDal:  testLocationDal,
	}
	bytModel, _ := manager.IJsonParser.EncodeJson(locationModel)
	var v, s, m = manager.AddLocation(bytModel)
	assert.Equal(t, &LocationResponseModel, v)
	assert.Equal(t, true, s)
	assert.Equal(t, "", m)
}
