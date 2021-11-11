package abstract

import "FilterWorkerService/internal/model"

type IGameSessionEveryLoginService interface {
	AddGameSession(data *model.GameSessionEveryLoginRespondModel) (s bool, m string)
	UpdateGameSession(modelResponse *model.GameSessionEveryLoginRespondModel) (s bool, m string)
}
