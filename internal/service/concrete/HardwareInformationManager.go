package concrete

import (
	model "FilterWorkerService/internal/model"
	IHardwareInformationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type HardwareInformationManager struct {
	ICacheService           ICacheService.ICacheService
	IHardwareInformationDal IHardwareInformationDal.IHardwareInformationDal
	IJsonParser             IJsonParser.IJsonParser
}

func (h *HardwareInformationManager) AddHardwareInformation(data *[]byte) (respondHardwareModel *model.HardwareInformationResponseModel, s bool, m string) {
	// Todo : 1 Model karşılanacak
	firstmodel := model.HardwareInformationModel{}
	err := h.IJsonParser.DecodeJson(data, &firstmodel)
	if err != nil {
		return &model.HardwareInformationResponseModel{}, false, err.Error()
	}
	// Todo: 2 Filtreler Buraya Yazılacak
	modelResponse := model.HardwareInformationResponseModel{}
	modelResponse.ClientId = firstmodel.ClientId
	modelResponse.ProjectId = firstmodel.ProjectId
	modelResponse.CustomerId = firstmodel.CustomerId
	modelResponse.DeviceType = int64(firstmodel.DeviceType)
	modelResponse.GraphicsDeviceType = int64(firstmodel.GraphicsDeviceType)
	modelResponse.GraphicsMemorySize = int64(firstmodel.GraphicsMemorySize)
	modelResponse.OperatingSystem, _, _ = h.ICacheService.ManageCache("OperatingSystem", firstmodel.OperatingSystem)
	modelResponse.ProcessorCount = int64(firstmodel.ProcessorCount)
	modelResponse.ProcessorType, _, _ = h.ICacheService.ManageCache("ProcessorType", firstmodel.ProcessorType)
	modelResponse.SystemMemorySize = int64(firstmodel.SystemMemorySize)
	// Todo : 3 Model burada kayıt edilecek
	logErr := h.IHardwareInformationDal.Add(&modelResponse)
	if logErr != nil {
		return &modelResponse, false, logErr.Error()
	}
	return &modelResponse, true, ""
}
