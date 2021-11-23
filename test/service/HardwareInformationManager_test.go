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

func Test_AddHardwareInformation_Success(t *testing.T){
	var testHardwareInfoDal = new(repository.MockHardwareInformationDal)
	var testCache = new(service.MockCacheService)
	var manager = concrete.HardwareInformationManager{
		ICacheService:           testCache,
		IHardwareInformationDal: testHardwareInfoDal,
		IJsonParser:             &gojson.GoJson{},
	}
	
	testCache.On("ManageCache", "OperatingSystem", harwareModel.OperatingSystem).Return(int64(1), true, "")
	testCache.On("ManageCache", "ProcessorType", harwareModel.ProcessorType).Return(int64(6), true, "")
	testHardwareInfoDal.On("Add", &hardwareResponseModel).Return(nil)
	var byteData,_ = manager.IJsonParser.EncodeJson(harwareModel)
	var v,s,m = manager.AddHardwareInformation(byteData)
	assert.Equal(t, &hardwareResponseModel, v)
	assert.Equal(t, true, s)
	assert.Equal(t, "", m)
}