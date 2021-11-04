package abstract

import "FilterWorkerService/internal/model"

type ILocationDal interface {
	Add(data *model.LocationResponseModel) error
}
