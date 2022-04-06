package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IHardwareDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type hardwareManager struct {
	ICacheService *ICacheService.ICacheService
	IHardwareDal  *IHardwareDal.IHardwareDal
	IJsonParser   *IJsonParser.IJsonParser
}

func HardwareManagerConstructor() *hardwareManager {
	return &hardwareManager{
		ICacheService: &IoC.CacheService,
		IHardwareDal:  &IoC.HardwareDal,
		IJsonParser:   &IoC.JsonParser,
	}
}

func (h *hardwareManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	// Todo : 1 Model karşılanacak
	firstmodel := model.HardwareModel{}
	convertErr := (*h.IJsonParser).DecodeJson(data, &firstmodel)
	if convertErr != nil {
		clogger.Error(&map[string]interface{}{"Byte array to HardwareModel  HardwareManager Json Parser Decode ERROR: ": convertErr.Error()})

		// log.Fatal("HardwareInformationManager", "AddHardwareInformation",
		// 	"byte array to HardwareInformationModel", "Json Parser Decode Err: ", convertErr.Error())
		return false, convertErr.Error()
	}
	// Todo: 2 Filtreler Buraya Yazılacak
	modelResponse := model.HardwareResponseModel{}
	modelResponse.Id = firstmodel.Id
	modelResponse.ClientId = firstmodel.ClientId
	modelResponse.ProjectId = firstmodel.ProjectId
	modelResponse.CustomerId = firstmodel.CustomerId
	modelResponse.DeviceType = int16(firstmodel.DeviceType)
	modelResponse.GraphicsDeviceType = int16(firstmodel.GraphicsDeviceType)
	modelResponse.GraphicsMemorySize = int16(firstmodel.GraphicsMemorySize)
	modelResponse.OperatingSystem, _, _ = (*h.ICacheService).ManageCache("OperatingSystem", firstmodel.OperatingSystem)
	modelResponse.ProcessorCount = int16(firstmodel.ProcessorCount)
	modelResponse.ProcessorType, _, _ = (*h.ICacheService).ManageCache("ProcessorType", firstmodel.ProcessorType)
	modelResponse.SystemMemorySize = int16(firstmodel.SystemMemorySize)
	// Todo : 3 Model burada kayıt edilecek
	// defer log.Print("HardwareInformationManager", "AddHardwareInformation",
	// 	modelResponse.ClientId, modelResponse.ProjectId)
	oldModel, err := (*h.IHardwareDal).GetById(modelResponse.ClientId, modelResponse.ProjectId)

	switch {

	case err != nil && err.Error() != "not found":
		clogger.Error(&map[string]interface{}{
			fmt.Sprintf("Get clientId: %d, projectId: %d hardware_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): err.Error(),
		})

	case err != nil && err.Error() == "not found":

		logErr := (*h.IHardwareDal).Add(&modelResponse)
		if logErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Add clientId: %d, projectId: %d hardware_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): logErr.Error(),
			})
			return false, logErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Add clientId: %d, projectId: %d hardware_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return true, "Added"

	case err == nil:
		_, updateResult, updateErr := h.UpdateHardware(&modelResponse, oldModel)
		if updateErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Update clientId: %d, projectId: %d game_session_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): updateErr.Error(),
			})
			return updateResult, updateErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Update clientId: %d, projectId: %d game_session_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return updateResult, "Updated"

	default:

		return false, ""

	}
	return true, ""
}

func (g *hardwareManager) UpdateHardware(modelResponse *model.HardwareResponseModel, oldModel *model.HardwareResponseModel) (updatedModel *model.HardwareResponseModel, s bool, m error) {
	oldModel.Id = modelResponse.Id
	oldModel.ClientId = modelResponse.ClientId
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.DeviceType = modelResponse.DeviceType
	oldModel.GraphicsDeviceType = modelResponse.GraphicsDeviceType
	oldModel.GraphicsMemorySize = modelResponse.GraphicsMemorySize
	oldModel.OperatingSystem = modelResponse.OperatingSystem
	oldModel.ProcessorCount = modelResponse.ProcessorCount
	oldModel.ProcessorType = modelResponse.ProcessorType
	oldModel.SystemMemorySize = modelResponse.SystemMemorySize

	logErr := (*g.IHardwareDal).UpdateById(oldModel.ClientId, oldModel.ProjectId, oldModel)
	if logErr != nil {
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}
