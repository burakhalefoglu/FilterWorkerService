package abstract

import "FilterWorkerService/internal/model"

type IGameSessionEveryLoginService interface {
	ConvertRawModelToResponseModel(data *[]byte) (gameSession *model.GameSessionEveryLoginRespondModel, s bool, m string)
}
