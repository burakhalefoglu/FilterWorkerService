package abstract

import "FilterWorkerService/internal/model"

type ILevelBaseSessionService interface {
	ConvertRawModelToResponseModel(data *[]byte) (convertedModel *model.LevelBaseSessionRespondModel, s bool, m string)
}
