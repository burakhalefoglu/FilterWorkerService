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

func (sc *ScreenClickManager) ConvertRawModelToResponseModel(data *[]byte) (respondModel *model.ScreenClickRespondModel, s bool, m string) {
	firstModel := model.ScreenClickModel{}
	err := sc.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return &model.ScreenClickRespondModel{}, false, err.Error()
	}
	hour := int64(firstModel.CreationAt.Hour())
	yearOfDay := int64(firstModel.CreationAt.YearDay())
	year := int64(firstModel.CreationAt.Year())
	modelResponse := model.ScreenClickRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.FirstClickSessionYearOfDay = yearOfDay
	modelResponse.FirstClickSessionYear = year
	modelResponse.FirstClickSessionHour = hour
	modelResponse.LastClickSessionYearOfDay = yearOfDay
	modelResponse.LastClickSessionYear = year
	modelResponse.LastClickSessionHour = hour
	modelResponse.FirstStartXCor = firstModel.StartXCor
	modelResponse.FirstStartYCor = firstModel.StartYCor
	modelResponse.FirstFinishXCor = firstModel.FinishXCor
	modelResponse.FirstFinishYCor = firstModel.FinishYCor
	modelResponse.LastStartXCor = firstModel.StartXCor
	modelResponse.LastStartYCor = firstModel.StartYCor
	modelResponse.LastFinishXCor = firstModel.FinishXCor
	modelResponse.LastFinishYCor = firstModel.FinishYCor
	modelResponse.FirstTouchCount = int64(firstModel.TouchCount)
	modelResponse.LastTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstMinusLastTouchCount = modelResponse.FirstTouchCount - modelResponse.LastTouchCount
	modelResponse.FirstFingerId = int64(firstModel.FingerId)
	modelResponse.LastFingerId = int64(firstModel.FingerId)
	modelResponse.FirstDayClickCount = int64(firstModel.TouchCount)
	modelResponse.TotalClickDay = modelResponse.LastClickSessionYearOfDay - modelResponse.FirstClickSessionYearOfDay
	modelResponse.TotalClickCount = int64(firstModel.TouchCount)
	modelResponse.TotalClickSessionCount = 1
	modelResponse.TotalStartXCor = int64(firstModel.StartXCor)
	modelResponse.TotalStartYCor = int64(firstModel.StartYCor)
	modelResponse.TotalFinishXCor = int64(firstModel.FinishXCor)
	modelResponse.TotalFinishYCor = int64(firstModel.FinishYCor)
	modelResponse.SessionBasedAvegareStartXCor = float64(modelResponse.TotalStartXCor) / float64(modelResponse.TotalClickSessionCount)
	modelResponse.SessionBasedAvegareStartYCor = float64(modelResponse.TotalStartYCor) / float64(modelResponse.TotalClickSessionCount)
	modelResponse.SessionBasedAvegareFinishXCor = float64(modelResponse.TotalFinishXCor) / float64(modelResponse.TotalClickSessionCount)
	modelResponse.SessionBasedAvegareFinishYCor = float64(modelResponse.TotalFinishYCor) / float64(modelResponse.TotalClickSessionCount)
	modelResponse.SessionBasedAvegareClickCount = float64(modelResponse.TotalClickCount) / float64(modelResponse.TotalClickSessionCount)
	modelResponse.DailyAvegareClickCount = float64(modelResponse.TotalClickCount) / float64(modelResponse.TotalClickDay)
	modelResponse.LastTouchCountMinusSessionBasedAvegareClickCount = float64(modelResponse.LastTouchCount) - float64(modelResponse.SessionBasedAvegareClickCount)
	return &modelResponse, true, ""
}

func (sc *ScreenClickManager) AddScreenClick(data *model.ScreenClickRespondModel) (s bool, m string) {
	logErr := sc.IScreenClickDal.Add(data)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func (sc *ScreenClickManager) UpdateScreenClick(modelResponse *model.ScreenClickRespondModel) (s bool, m string) {
	oldModel, err := sc.IScreenClickDal.GetScreenClickByCustomerId(modelResponse.CustomerId)
	if err != nil {
		return false, err.Error()
	}
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	//oldModel.FirstClickSessionYearOfDay
	//oldModel.FirstClickSessionYear
	//oldModel.FirstClickSessionHour
	oldModel.LastClickSessionYearOfDay = modelResponse.LastClickSessionYearOfDay
	oldModel.LastClickSessionYear = modelResponse.LastClickSessionYear
	oldModel.LastClickSessionHour = modelResponse.LastClickSessionHour
	//oldModel.FirstStartXCor
	//oldModel.FirstStartYCor
	//oldModel.FirstFinishXCor
	//oldModel.FirstFinishYCor
	oldModel.LastStartXCor = modelResponse.LastStartXCor
	oldModel.LastStartYCor = modelResponse.LastStartYCor
	oldModel.LastFinishXCor = modelResponse.LastFinishXCor
	oldModel.LastFinishYCor = modelResponse.LastFinishYCor
	//oldModel.FirstTouchCount
	oldModel.LastTouchCount = modelResponse.LastTouchCount
	oldModel.FirstMinusLastTouchCount = oldModel.FirstTouchCount - oldModel.LastTouchCount
	//oldModel.FirstFingerId
	oldModel.LastFingerId = modelResponse.LastFingerId
	oldModel.FirstDayClickCount = calculateFirstDayClickCount(modelResponse, oldModel)
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
	logErr := sc.IScreenClickDal.UpdateScreenClickByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func calculateFirstDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) int64 {
	if (oldModel.FirstClickSessionYearOfDay == modelResponse.FirstClickSessionYearOfDay) && (oldModel.FirstClickSessionYear == modelResponse.FirstClickSessionYear) {
		oldModel.FirstDayClickCount = oldModel.FirstDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.FirstDayClickCount
	}
	return oldModel.FirstDayClickCount
}
