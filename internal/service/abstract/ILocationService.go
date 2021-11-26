package abstract

import "FilterWorkerService/internal/model"

type ILocationService interface {
	AddLocation(data *[]byte) (respondLocationModel *model.LocationResponseModel, s bool, m string)
}
