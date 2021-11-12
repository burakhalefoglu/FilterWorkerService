package abstract

import "FilterWorkerService/internal/model"

type ILevelBaseSessionService interface {
	AddLevelBaseSession(data *model.LevelBaseSessionRespondModel) (s bool, m string)
	UpdateLevelBaseSession(modelResponse *model.LevelBaseSessionRespondModel) (s bool, m string)
}
