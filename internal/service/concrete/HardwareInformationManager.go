package concrete

import (
	model2 "FilterWorkerService/internal/model"
	IHardwareInformationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type HardwareInformationManager struct {
	ICacheService           ICacheService.ICacheService
	IHardwareInformationDal IHardwareInformationDal.IHardwareInformationDal
	IJsonParser             IJsonParser.IJsonParser
}

func (h *HardwareInformationManager) AddHardwareInformation(data *[]byte) (s bool, m string) {
	// Todo : 1 Model karşılanacak
	model := model2.HardwareInformationModel{}
	err := h.IJsonParser.DecodeJson(data, &model)
	if err != nil {
		return false, err.Error()
	}
	
	// Todo: 2 Filtreler Buraya Yazılacak
	modelResponse := model2.HardwareInformationResponseModel{}
	modelResponse.ClientId = model.ClientId
	modelResponse.ProjectId = model.ProjectId
	modelResponse.CustomerId = model.CustomerId
	modelResponse.DeviceType = int64(model.DeviceType)
	modelResponse.GraphicsDeviceType = int64(model.GraphicsDeviceType)
	modelResponse.GraphicsMemorySize = int64(model.GraphicsMemorySize)
	modelResponse.OperatingSystem, s, m = h.ICacheService.ManageCache("OperatingSystem", model.OperatingSystem)
	modelResponse.ProcessorCount = int64(model.ProcessorCount)
	modelResponse.SystemMemorySize = int64(model.SystemMemorySize)
	// Todo : 3 Model burada kayıt edilecek
	logErr := h.IHardwareInformationDal.Add(&modelResponse)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}
