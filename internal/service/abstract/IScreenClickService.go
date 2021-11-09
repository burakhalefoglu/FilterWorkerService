package abstract

import "FilterWorkerService/internal/model"

type IScreenClickService interface {
	AddScreenClick(data *model.ScreenClickRespondModel) (s bool, m string)
	UpdateScreenClick(modelResponse *model.ScreenClickRespondModel) (s bool, m string)
}
