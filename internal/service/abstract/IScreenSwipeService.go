package abstract

import "FilterWorkerService/internal/model"

type IScreenSwipeService interface {
	AddScreenSwipe(data *model.ScreenSwipeRespondModel) (s bool, m string)
	UpdateScreenSwipe(modelResponse *model.ScreenSwipeRespondModel) (s bool, m string)
}
