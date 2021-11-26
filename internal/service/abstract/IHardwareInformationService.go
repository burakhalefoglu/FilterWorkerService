package abstract

import "FilterWorkerService/internal/model"

type IHardwareInformationService interface {
	AddHardwareInformation(data *[]byte) (respondHardwareModel *model.HardwareInformationResponseModel, s bool, m string)
}