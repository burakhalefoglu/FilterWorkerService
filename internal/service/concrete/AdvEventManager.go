package concrete

import (
	model "FilterWorkerService/internal/model"
	IAdvEventDal "FilterWorkerService/internal/repository/abstract"

	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type AdvEventManager struct {
	IAdvEventDal IAdvEventDal.IAdvEventDal
	IJsonParser  IJsonParser.IJsonParser
}

func (a *AdvEventManager) ConvertRawModelToResponseModel(data *[]byte) (respondModel *model.AdvEventRespondModel, s bool, m string) {
	firstModel := model.AdvEventModel{}
	err := a.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return &model.AdvEventRespondModel{}, false, err.Error()
	}
	hour := int64(firstModel.TrigerdTime.Hour())
	day := int64(firstModel.TrigerdTime.Weekday())
	yearOfDay := int64(firstModel.TrigerdTime.YearDay())

	modelResponse := model.AdvEventRespondModel{}
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.TotalAdvDay = 1
	modelResponse.TotalAdvCount = 1
	//modelResponse.TotalVideoAdvCount   ------> 1 video veya geçiş reklamı çıkmaasına göre belirlenecek
	//modelResponse.TotalInterstitialAdvCount
	modelResponse.LevelBasedAverageInterstitialAdvCount = calculateAdvLevelBasedAvgInterstitialCount(&modelResponse)
	modelResponse.LevelBasedAverageVideoAdvCount = calculateAdvLevelBasedAvgVideoCount(&modelResponse)
	modelResponse.AverageDailyVideoAdvClickCount = float64(modelResponse.TotalVideoAdvCount) / float64(modelResponse.TotalAdvDay)
	modelResponse.FirstAdvYearOfDay = yearOfDay
	modelResponse.FirstAdvClickHour = hour
	//modelResponse.FirstVideoClickYearOfDay
	modelResponse.LastAdvYearOfDay = yearOfDay
	//modelResponse.LastVideoClickYearOfDay
	modelResponse.LastAdvClickHour = hour
	//modelResponse.FirstDayVideoClickCount
	modelResponse.LastDayVideoClickCount = 0
	modelResponse.LastMinusFirstDayVideoClickCount = modelResponse.LastDayVideoClickCount - modelResponse.FirstDayVideoClickCount
	modelResponse.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount = modelResponse.LastDayVideoClickCount - int64(modelResponse.AverageDailyVideoAdvClickCount)
	determineAdvDay(&modelResponse, day)
	determineAdvHour(&modelResponse, hour)
	determineAdvAmPm(&modelResponse, hour)
	return &modelResponse, true, ""
}

func (a *AdvEventManager) AddAdvEvent(data *model.AdvEventRespondModel) (s bool, m string) {

	logErr := a.IAdvEventDal.Add(data)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func (a *AdvEventManager) UpdateAdvEvent(modelResponse *model.AdvEventRespondModel) (s bool, m string) {
	oldmodel, err := a.IAdvEventDal.GetAdvEventByCustomerId(modelResponse.CustomerId)
	if err != nil {
		return false, err.Error()
	}
	// modelResponse.ClientId = oldmodel.ClientId
	// modelResponse.ProjectId = oldmodel.ProjectId
	//modelResponse.CustomerId = oldmodel.CustomerId
	//modelResponse.LevelIndex = oldmodel.LevelIndex
	modelResponse.TotalAdvDay = calculateTotalAdvDay(modelResponse, oldmodel)
	modelResponse.TotalAdvCount = oldmodel.TotalAdvCount + modelResponse.TotalAdvCount
	modelResponse.TotalVideoAdvCount = modelResponse.TotalVideoAdvCount + oldmodel.TotalAdvCount
	modelResponse.TotalInterstitialAdvCount = modelResponse.TotalInterstitialAdvCount + oldmodel.TotalInterstitialAdvCount
	modelResponse.LevelBasedAverageInterstitialAdvCount = calculateAdvLevelBasedAvgInterstitialCount(modelResponse)
	modelResponse.LevelBasedAverageVideoAdvCount = calculateAdvLevelBasedAvgVideoCount(modelResponse)
	modelResponse.AverageDailyVideoAdvClickCount = float64(modelResponse.TotalVideoAdvCount) / float64(modelResponse.TotalAdvDay)
	modelResponse.FirstDayVideoClickCount = calculateFirstDayVideoClickCount(modelResponse, oldmodel)
	modelResponse.LastDayVideoClickCount = calculateLastDayVideoClickCount(modelResponse, oldmodel)
	modelResponse.LastMinusFirstDayVideoClickCount = modelResponse.LastDayVideoClickCount - modelResponse.FirstDayVideoClickCount
	modelResponse.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount = modelResponse.LastDayVideoClickCount - int64(modelResponse.AverageDailyVideoAdvClickCount)

	logErr := a.IAdvEventDal.UpdateAdvEventByCustomerId(modelResponse.CustomerId, modelResponse)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func calculateFirstDayVideoClickCount(modelResponse *model.AdvEventRespondModel, oldmodel *model.AdvEventRespondModel) int64 {
	if (oldmodel.FirstVideoClickYearOfDay == modelResponse.FirstAdvYearOfDay){
		modelResponse.FirstDayVideoClickCount = oldmodel.FirstDayVideoClickCount + modelResponse.FirstDayVideoClickCount
		return modelResponse.FirstDayVideoClickCount
	}
	modelResponse.FirstDayVideoClickCount = oldmodel.FirstDayVideoClickCount
	return modelResponse.FirstDayVideoClickCount
}
func calculateLastDayVideoClickCount(modelResponse *model.AdvEventRespondModel, oldmodel *model.AdvEventRespondModel) int64 {
	if (oldmodel.LastVideoClickYearOfDay == modelResponse.LastVideoClickYearOfDay) && (oldmodel.FirstVideoClickYearOfDay != modelResponse.FirstVideoClickYearOfDay){
		modelResponse.LastDayVideoClickCount = oldmodel.LastDayVideoClickCount + modelResponse.LastDayVideoClickCount
		return modelResponse.LastDayVideoClickCount
	}
	modelResponse.LastDayVideoClickCount = oldmodel.LastDayVideoClickCount
	return modelResponse.LastDayVideoClickCount
}

func calculateTotalAdvDay(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) int64 {
	if (modelResponse.LastAdvYearOfDay != oldModel.LastAdvYearOfDay) {
		modelResponse.TotalAdvDay = oldModel.TotalAdvDay + modelResponse.TotalAdvDay
		return modelResponse.TotalAdvDay
	}
	modelResponse.TotalAdvDay = oldModel.TotalAdvDay
	return modelResponse.TotalAdvDay
}

func determineAdvDay(modelResponse *model.AdvEventRespondModel, day int64) {
	if day == 0 {
		modelResponse.SundayVideoAdvClickCount = 1
	} else if day == 1 {
		modelResponse.MondayVideoAdvClickCount = 1
	} else if day == 2 {
		modelResponse.TuesdayVideoAdvClickCount = 1
	} else if day == 3 {
		modelResponse.WednesdayVideoAdvClickCount = 1
	} else if day == 4 {
		modelResponse.ThursdayVideoAdvClickCount = 1
	} else if day == 5 {
		modelResponse.FridayVideoAdvClickCount = 1
	} else if day == 6 {
		modelResponse.SaturdayVideoAdvClickCount = 1
	}
}

func determineAdvHour(modelResponse *model.AdvEventRespondModel, hour int64) {
	if hour <= 5 {
		modelResponse.VideoAdvClick0To5HourCount = 1
	} else if (hour > 5) && (hour <= 11) {
		modelResponse.VideoAdvClick6To11HourCount = 1
	} else if (hour > 11) && (hour <= 17) {
		modelResponse.VideoAdvClick12To17HourCount = 1
	} else if (hour > 17) && (hour <= 23) {
		modelResponse.VideoAdvClick18To23HourCount = 1
	}
}

func determineAdvAmPm(modelResponse *model.AdvEventRespondModel, hour int64) {
	if hour <= 12 {
		modelResponse.AmVideoAdvClickCount = 1
	} else if hour > 12 {
		modelResponse.PmVideoAdvClickCount = 1
	}
}

func calculateAdvLevelBasedAvgInterstitialCount(modelResponse *model.AdvEventRespondModel) float64 {
	if modelResponse.LevelIndex == 0 {
		modelResponse.LevelBasedAverageInterstitialAdvCount = float64(modelResponse.TotalInterstitialAdvCount)
		return modelResponse.LevelBasedAverageInterstitialAdvCount
	}
	modelResponse.LevelBasedAverageInterstitialAdvCount = float64(modelResponse.TotalInterstitialAdvCount) / float64(modelResponse.LevelIndex)
	return modelResponse.LevelBasedAverageInterstitialAdvCount
}

func calculateAdvLevelBasedAvgVideoCount(modelResponse *model.AdvEventRespondModel) float64 {
	if modelResponse.LevelIndex == 0 {
		modelResponse.LevelBasedAverageVideoAdvCount = float64(modelResponse.TotalVideoAdvCount)
		return modelResponse.LevelBasedAverageVideoAdvCount
	}
	modelResponse.LevelBasedAverageVideoAdvCount = float64(modelResponse.TotalVideoAdvCount) / float64(modelResponse.LevelIndex)
	return modelResponse.LevelBasedAverageVideoAdvCount
}
