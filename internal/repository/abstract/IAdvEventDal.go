package abstract

import "FilterWorkerService/internal/model"

type IAdvEventDal interface {
	Add(data *model.AdvEventRespondModel) error
}