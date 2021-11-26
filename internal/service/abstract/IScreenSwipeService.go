package abstract

import "FilterWorkerService/internal/model"

type IScreenSwipeService interface {
	ConvertRawModelToResponseModel(data *[]byte) (convertedModel *model.ScreenSwipeRespondModel, s bool, m string)
}
