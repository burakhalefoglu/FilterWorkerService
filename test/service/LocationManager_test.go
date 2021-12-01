package test

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/Log"
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
	var testCache = new(service.MockCacheService)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.LocationDal = testLocationDal
	IoC.Logger = testLog
	IoC.CacheService = testCache
	var manager = concrete.LocationManagerConstructor()
	var locationModel_test = locationModel
	var LocationResponseModel_test = LocationResponseModel
	var locationModel_test_byte, _ = json.EncodeJson(locationModel_test)

	testCache.On("ManageCache", "Continent", locationModel_test.Continent).Return(int64(1), true, "")
	testCache.On("ManageCache", "Country", locationModel_test.Country).Return(int64(30), true, "")
	testCache.On("ManageCache", "City", locationModel_test.City).Return(int64(187), true, "")
	testCache.On("ManageCache", "Region", locationModel_test.Region).Return(int64(93), true, "")
	testCache.On("ManageCache", "Org", locationModel_test.Org).Return(int64(8), true, "")
	testLocationDal.On("Add", &LocationResponseModel_test).Return(nil)

	var v, s, m = manager.AddLocation(locationModel_test_byte)
	var value, success = v.(model.LocationResponseModel)
	if success == true {
		assert.Equal(t, &LocationResponseModel_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, &LocationResponseModel, v)
	assert.Equal(t, true, s)
	assert.Equal(t, "", m)
}
