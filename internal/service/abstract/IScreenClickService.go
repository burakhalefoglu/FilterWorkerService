package abstract

import "FilterWorkerService/internal/model"

type IScreenClickService interface {
	ConvertRawModelToResponseModel(data *[]byte) (respondModel *model.ScreenClickRespondModel, s bool, m string)
}
