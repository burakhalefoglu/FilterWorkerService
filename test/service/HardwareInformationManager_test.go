package test

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"FilterWorkerService/test/Mock/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

var harwareModel = model.HardwareInformationModel{
	ProjectId:          "Test",
	ClientId:           "Test",
	CustomerId:         "Test",
	DeviceType:         87893,
	GraphicsDeviceType: 9056,
	GraphicsMemorySize: 4000,
	OperatingSystem:    "IOS",
	ProcessorCount:     16,
	ProcessorType:      "Intel",
	SystemMemorySize:   4000,
}

var hardwareResponseModel = model.HardwareInformationResponseModel{
	ProjectId:          harwareModel.ProjectId,
	ClientId:           harwareModel.ClientId,
	CustomerId:         harwareModel.CustomerId,
	DeviceType:         int64(harwareModel.DeviceType),
	GraphicsDeviceType: int64(harwareModel.GraphicsDeviceType),
	GraphicsMemorySize: int64(harwareModel.GraphicsMemorySize),
	OperatingSystem:    1,
	ProcessorCount:     int64(harwareModel.ProcessorCount),
	ProcessorType:      6,
	SystemMemorySize:   int64(harwareModel.SystemMemorySize),
}

func Test_AddHardwareInformation_AddSuccess(t *testing.T){
	var testHardwareInfoDal = new(repository.MockHardwareInformationDal)
	var testCache = new(service.MockCacheService)	
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.HardwareInformationDal = testHardwareInfoDal
	IoC.CacheService = testCache
	var manager = concrete.HardwareInformationManagerConstructor()
	var hardwareModel_test = harwareModel
	var hardwareResponseModel_test = hardwareResponseModel
	var hardwareModel_test_byte, _ = json.EncodeJson(hardwareModel_test)

	testCache.On("ManageCache", "OperatingSystem", harwareModel.OperatingSystem).Return(int64(1), true, "")
	testCache.On("ManageCache", "ProcessorType", harwareModel.ProcessorType).Return(int64(6), true, "")
	testHardwareInfoDal.On("Add", &hardwareResponseModel).Return(nil)

	var v,s,m = manager.AddHardwareInformation(hardwareModel_test_byte)
	var value, success = v.(model.HardwareInformationResponseModel)
	if success == true {
		assert.Equal(t, &hardwareResponseModel_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, true, s)
	assert.Equal(t, "", m)
}