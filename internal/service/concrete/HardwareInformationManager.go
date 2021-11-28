package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IHardwareInformationDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/logger"
)

type hardwareInformationManager struct {
	ICacheService           *ICacheService.ICacheService
	IHardwareInformationDal *IHardwareInformationDal.IHardwareInformationDal
	IJsonParser             *IJsonParser.IJsonParser
	ILog                    *logger.ILog
}

func HardwareInformationManagerConstructor() *hardwareInformationManager {
	return &hardwareInformationManager{
		ICacheService:           &IoC.CacheService,
		IHardwareInformationDal: &IoC.HardwareInformationDal,
		IJsonParser:             &IoC.JsonParser,
		ILog:                    &IoC.Logger,
	}
}

func (h *hardwareInformationManager) AddHardwareInformation(data *[]byte) (v interface{}, s bool, m string) {
	// Todo : 1 Model karşılanacak
	firstmodel := model.HardwareInformationModel{}
	Err := (*h.IJsonParser).DecodeJson(data, &firstmodel)
	if Err != nil {
		(*h.ILog).SendErrorLog("HardwareInformationManager", "AddHardwareInformation",
			"byte array to HardwareInformationModel", "Json Parser Decode Err: ", Err.Error())
		return &model.HardwareInformationResponseModel{}, false, Err.Error()
	}
	// Todo: 2 Filtreler Buraya Yazılacak
	modelResponse := model.HardwareInformationResponseModel{}
	modelResponse.ClientId = firstmodel.ClientId
	modelResponse.ProjectId = firstmodel.ProjectId
	modelResponse.CustomerId = firstmodel.CustomerId
	modelResponse.DeviceType = int64(firstmodel.DeviceType)
	modelResponse.GraphicsDeviceType = int64(firstmodel.GraphicsDeviceType)
	modelResponse.GraphicsMemorySize = int64(firstmodel.GraphicsMemorySize)
	modelResponse.OperatingSystem, _, _ = (*h.ICacheService).ManageCache("OperatingSystem", firstmodel.OperatingSystem)
	modelResponse.ProcessorCount = int64(firstmodel.ProcessorCount)
	modelResponse.ProcessorType, _, _ = (*h.ICacheService).ManageCache("ProcessorType", firstmodel.ProcessorType)
	modelResponse.SystemMemorySize = int64(firstmodel.SystemMemorySize)
	// Todo : 3 Model burada kayıt edilecek
	defer (*h.ILog).SendInfoLog("HardwareInformationManager", "AddHardwareInformation",
		modelResponse.ClientId, modelResponse.ProjectId)
	logErr := (*h.IHardwareInformationDal).Add(&modelResponse)
	if logErr != nil {
		(*h.ILog).SendErrorLog("HardwareInformationManager", "AddHardwareInformation",
			"HardwareInformationDal_Add", logErr.Error())
		return &modelResponse, false, logErr.Error()
	}
	return &modelResponse, true, ""
}
