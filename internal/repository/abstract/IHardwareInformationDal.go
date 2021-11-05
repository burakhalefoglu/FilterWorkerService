package abstract

import "FilterWorkerService/internal/model"

type IHardwareInformationDal interface {
	Add(data *model.HardwareInformationResponseModel) error
}