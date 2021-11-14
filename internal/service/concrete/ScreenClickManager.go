package concrete

import (
	model "FilterWorkerService/internal/model"
	IScreenClickDal "FilterWorkerService/internal/repository/abstract"

	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type ScreenClickManager struct {
	IScreenClickDal IScreenClickDal.IScreenClickDal
	IJsonParser     IJsonParser.IJsonParser
}

func (sc *ScreenClickManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.ScreenClickModel{}
	err := sc.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	hour := int64(firstModel.CreationAt.Hour())
	yearOfDay := int64(firstModel.CreationAt.YearDay())
	year := int64(firstModel.CreationAt.Year())
	minute := int64(firstModel.CreationAt.Minute())
	modelResponse := model.ScreenClickRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.FirstClickSessionYearOfDay = yearOfDay
	modelResponse.FirstClickSessionYear = year
	modelResponse.FirstClickSessionHour = hour
	modelResponse.FirstClickSessionMinute = minute
	modelResponse.FirstTouchCount = int64(firstModel.TouchCount)
	modelResponse.SecondClickSessionHour = 0
	modelResponse.SecondClickSessionMinute = 0
	modelResponse.SecondTouchCount = 0
	modelResponse.ThirdClickSessionHour = 0
	modelResponse.ThirdClickSessionMinute = 0
	modelResponse.ThirdTouchCount = 0
	modelResponse.PenultimateClickSessionHour = 0
	modelResponse.PenultimateClickSessionMinute = 0
	modelResponse.PenultimateTouchCount = 0
	modelResponse.LastClickSessionYearOfDay = yearOfDay
	modelResponse.LastClickSessionYear = year
	modelResponse.LastClickSessionHour = hour
	modelResponse.LastClickSessionMinute = minute
	modelResponse.LastTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstStartXCor = firstModel.StartXCor
	modelResponse.FirstStartYCor = firstModel.StartYCor
	modelResponse.FirstFinishXCor = firstModel.FinishXCor
	modelResponse.FirstFinishYCor = firstModel.FinishYCor
	modelResponse.SecondStartXCor = 0
	modelResponse.SecondStartYCor = 0
	modelResponse.SecondFinishXCor = 0
	modelResponse.SecondFinishYCor = 0
	modelResponse.ThirdStartXCor = 0
	modelResponse.ThirdStartYCor = 0
	modelResponse.ThirdFinishXCor = 0
	modelResponse.ThirdFinishYCor = 0
	modelResponse.PenultimateStartXCor = 0
	modelResponse.PenultimateStartYCor = 0
	modelResponse.PenultimateFinishXCor = 0
	modelResponse.PenultimateFinishYCor = 0
	modelResponse.LastStartXCor = firstModel.StartXCor
	modelResponse.LastStartYCor = firstModel.StartYCor
	modelResponse.LastFinishXCor = firstModel.FinishXCor
	modelResponse.LastFinishYCor = firstModel.FinishYCor
	modelResponse.FirstMinusLastTouchCount = modelResponse.FirstTouchCount - modelResponse.LastTouchCount
	modelResponse.FirstFingerId = int64(firstModel.FingerId)
	modelResponse.PenultimateFingerId = 0
	modelResponse.LastFingerId = int64(firstModel.FingerId)
	modelResponse.FirstDayClickCount = int64(firstModel.TouchCount)
	modelResponse.PenultimateDayClickCount = 0
	modelResponse.LastDayClickCount = int64(firstModel.TouchCount)
	modelResponse.TotalClickDay = 1
	modelResponse.TotalClickCount = 1
	modelResponse.TotalClickSessionCount = 1
	modelResponse.TotalStartXCor = int64(firstModel.StartXCor)
	modelResponse.TotalStartYCor = int64(firstModel.StartYCor)
	modelResponse.TotalFinishXCor = int64(firstModel.FinishXCor)
	modelResponse.TotalFinishYCor = int64(firstModel.FinishYCor)
	modelResponse.SessionBasedAvegareStartXCor = float64(firstModel.StartXCor)
	modelResponse.SessionBasedAvegareStartYCor = float64(firstModel.StartYCor)
	modelResponse.SessionBasedAvegareFinishXCor = float64(firstModel.FinishXCor)
	modelResponse.SessionBasedAvegareFinishYCor = float64(firstModel.FinishYCor)
	modelResponse.SessionBasedAvegareClickCount = float64(firstModel.TouchCount)
	modelResponse.DailyAvegareClickCount = float64(firstModel.TouchCount)
	modelResponse.LastTouchCountMinusSessionBasedAvegareClickCount = 0

	oldModel, err := sc.IScreenClickDal.GetScreenClickById(modelResponse.ClientId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := sc.IScreenClickDal.Add(&modelResponse)
		if logErr != nil {
			return false, logErr.Error()
		}
		return true, ""

	case err == nil:
		updateResult, updateErr := sc.updateScreenClick(&modelResponse, oldModel)
		if updateErr != nil {
			return updateResult, updateErr.Error()
		}
		return updateResult, ""

	default:

		return false, err.Error()

	}
}

func (sc *ScreenClickManager) updateScreenClick(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	//oldModel.FirstClickSessionYearOfDay
	//oldModel.FirstClickSessionYear
	//oldModel.FirstClickSessionHour
	//oldModel.FirstClickSessionMinute
	//oldModel.FirstTouchCount
	oldModel.SecondClickSessionHour, oldModel.SecondClickSessionMinute, oldModel.SecondTouchCount = calculateSecondClick(modelResponse, oldModel)
	oldModel.ThirdClickSessionHour, oldModel.ThirdClickSessionMinute, oldModel.ThirdTouchCount = calculateThirdClick(modelResponse, oldModel)
	oldModel.PenultimateClickSessionHour = oldModel.LastClickSessionHour
	oldModel.PenultimateClickSessionMinute = oldModel.LastClickSessionMinute
	oldModel.PenultimateTouchCount = oldModel.LastTouchCount
	oldModel.LastClickSessionYearOfDay = modelResponse.LastClickSessionYearOfDay
	oldModel.LastClickSessionYear = modelResponse.LastClickSessionYear
	oldModel.LastClickSessionHour = modelResponse.LastClickSessionHour
	oldModel.LastClickSessionMinute = modelResponse.LastClickSessionMinute
	oldModel.LastTouchCount = modelResponse.LastTouchCount
	//oldModel.FirstStartXCor
	//oldModel.FirstStartYCor
	//oldModel.FirstFinishXCor
	//oldModel.FirstFinishYCor
	oldModel.SecondStartXCor, oldModel.SecondStartYCor, oldModel.SecondFinishXCor, oldModel.SecondFinishYCor = calculateSecondRotate(modelResponse, oldModel)
	oldModel.ThirdStartXCor, modelResponse.ThirdStartYCor, modelResponse.ThirdFinishXCor, modelResponse.ThirdFinishYCor = calculateThirdRotate(modelResponse, oldModel)
	oldModel.PenultimateStartXCor = oldModel.LastStartXCor
	oldModel.PenultimateStartYCor = oldModel.LastStartYCor
	oldModel.PenultimateFinishXCor = oldModel.LastFinishXCor
	oldModel.PenultimateFinishYCor = oldModel.LastFinishYCor
	oldModel.LastStartXCor = modelResponse.LastStartXCor
	oldModel.LastStartYCor = modelResponse.LastStartYCor
	oldModel.LastFinishXCor = modelResponse.LastFinishXCor
	oldModel.LastFinishYCor = modelResponse.LastFinishYCor
	oldModel.FirstMinusLastTouchCount = oldModel.FirstTouchCount - oldModel.LastTouchCount
	//oldModel.FirstFingerId
	oldModel.PenultimateFingerId = oldModel.LastFingerId
	oldModel.LastFingerId = modelResponse.LastFingerId
	oldModel.FirstDayClickCount = calculateFirstDayClickCount(modelResponse, oldModel)
	oldModel.PenultimateDayClickCount = calculatePenultimateDayClickCount(modelResponse, oldModel)
	oldModel.LastDayClickCount = calculateLastDayClickCount(modelResponse, oldModel)
	oldModel.TotalClickDay = (oldModel.LastClickSessionYearOfDay - oldModel.FirstClickSessionYearOfDay) + 365*(oldModel.LastClickSessionYear-oldModel.FirstClickSessionYear)
	oldModel.TotalClickCount = oldModel.TotalClickCount + modelResponse.TotalClickCount
	oldModel.TotalClickSessionCount = oldModel.TotalClickSessionCount + modelResponse.TotalClickSessionCount
	oldModel.TotalStartXCor = oldModel.TotalStartXCor + modelResponse.TotalStartXCor
	oldModel.TotalStartYCor = oldModel.TotalStartYCor + modelResponse.TotalStartYCor
	oldModel.TotalFinishXCor = oldModel.TotalFinishXCor + modelResponse.TotalFinishXCor
	oldModel.TotalFinishYCor = oldModel.TotalFinishYCor + modelResponse.TotalFinishYCor
	oldModel.SessionBasedAvegareStartXCor = float64(oldModel.TotalStartXCor) / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareStartYCor = float64(oldModel.TotalStartYCor) / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareFinishXCor = float64(oldModel.TotalFinishXCor) / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareFinishYCor = float64(oldModel.TotalFinishYCor) / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareClickCount = float64(oldModel.TotalClickCount) / float64(oldModel.TotalClickSessionCount)
	oldModel.DailyAvegareClickCount = float64(oldModel.TotalClickCount) / float64(oldModel.TotalClickDay)
	oldModel.LastTouchCountMinusSessionBasedAvegareClickCount = float64(oldModel.LastTouchCount) - float64(oldModel.SessionBasedAvegareClickCount)
	logErr := sc.IScreenClickDal.UpdateScreenClickById(oldModel.ClientId, oldModel)
	if logErr != nil {
		return false, logErr
	}
	return true, nil
}

func calculateSecondRotate(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (startx float64, starty float64, finishx float64, finishy float64) {
	if oldModel.TotalClickCount == 1 {
		oldModel.SecondStartXCor = modelResponse.FirstStartXCor
		oldModel.SecondStartYCor = modelResponse.FirstStartYCor
		oldModel.SecondFinishXCor = modelResponse.FirstFinishXCor
		oldModel.SecondFinishYCor = modelResponse.FirstFinishYCor
		return oldModel.SecondStartXCor, oldModel.SecondStartYCor, oldModel.SecondFinishXCor, oldModel.SecondFinishYCor
	}
	return oldModel.SecondStartXCor, oldModel.SecondStartYCor, oldModel.SecondFinishXCor, oldModel.SecondFinishYCor
}

func calculateThirdRotate(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (startx float64, starty float64, finishx float64, finishy float64) {
	if oldModel.TotalClickCount == 2 {
		oldModel.ThirdStartXCor = modelResponse.FirstStartXCor
		oldModel.ThirdStartYCor = modelResponse.FirstStartYCor
		oldModel.ThirdFinishXCor = modelResponse.FirstFinishXCor
		oldModel.ThirdFinishYCor = modelResponse.FirstFinishYCor
		return oldModel.ThirdStartXCor, oldModel.ThirdStartYCor, oldModel.ThirdFinishXCor, oldModel.ThirdFinishYCor
	}
	return oldModel.ThirdStartXCor, oldModel.ThirdStartYCor, oldModel.ThirdFinishXCor, oldModel.ThirdFinishYCor
}

func calculateSecondClick(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (hour int64, minute int64, count int64) {
	if oldModel.TotalClickCount == 1 {
		oldModel.SecondClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.SecondClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.SecondTouchCount = modelResponse.FirstTouchCount
		return oldModel.SecondClickSessionHour, oldModel.SecondClickSessionMinute, oldModel.SecondTouchCount
	}
	return oldModel.SecondClickSessionHour, oldModel.SecondClickSessionMinute, oldModel.SecondTouchCount
}

func calculateThirdClick(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (hour int64, minute int64, count int64) {
	if oldModel.TotalClickCount == 2 {
		oldModel.ThirdClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.ThirdClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.ThirdTouchCount = modelResponse.FirstTouchCount
		return oldModel.ThirdClickSessionHour, oldModel.ThirdClickSessionMinute, oldModel.ThirdTouchCount
	}
	return oldModel.ThirdClickSessionHour, oldModel.ThirdClickSessionMinute, oldModel.ThirdTouchCount
}

func calculateFirstDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) int64 {
	if (oldModel.FirstClickSessionYearOfDay == modelResponse.FirstClickSessionYearOfDay) && (oldModel.FirstClickSessionYear == modelResponse.FirstClickSessionYear) {
		oldModel.FirstDayClickCount = oldModel.FirstDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.FirstDayClickCount
	}
	return oldModel.FirstDayClickCount
}

func calculatePenultimateDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (count int64) {
	if oldModel.LastClickSessionYearOfDay != modelResponse.LastClickSessionYearOfDay {
		oldModel.PenultimateDayClickCount = oldModel.LastDayClickCount
		return oldModel.PenultimateDayClickCount
	}
	return oldModel.PenultimateDayClickCount
}

func calculateLastDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (count int64) {
	if (oldModel.LastClickSessionYearOfDay == modelResponse.LastClickSessionYearOfDay) && (oldModel.FirstClickSessionYearOfDay != modelResponse.FirstClickSessionYearOfDay) {
		oldModel.LastDayClickCount = oldModel.LastDayClickCount + modelResponse.LastDayClickCount
		return oldModel.LastDayClickCount
	} else if (oldModel.LastClickSessionYearOfDay != modelResponse.LastClickSessionYearOfDay) && (oldModel.FirstClickSessionYearOfDay != modelResponse.FirstClickSessionYearOfDay) {
		return modelResponse.LastDayClickCount
	}
	return oldModel.LastDayClickCount
}
