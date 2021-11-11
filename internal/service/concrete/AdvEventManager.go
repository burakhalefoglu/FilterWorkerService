package concrete

import (
	model "FilterWorkerService/internal/model"
	IAdvEventDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"

	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type AdvEventManager struct {
	IAdvEventDal  IAdvEventDal.IAdvEventDal
	IJsonParser   IJsonParser.IJsonParser
	ICacheService ICacheService.ICacheService
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
	year := int64(firstModel.TrigerdTime.Year())
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
	modelResponse.AverageAdvDailyVideoClickCount = float64(modelResponse.TotalVideoAdvCount) / float64(modelResponse.TotalAdvDay)
	modelResponse.FirstAdvYearOfDay = yearOfDay
	modelResponse.FirstAdvYear = year
	modelResponse.FirstAdvClickHour = hour
	//modelResponse.FirstVideoClickYearOfDay
	//modelResponse.FirstVideoClickHour
	modelResponse.FirstAdvType, _, _ = a.ICacheService.ManageCache("AdvType", firstModel.AdvType)
	modelResponse.LastAdvYearOfDay = yearOfDay
	modelResponse.LastAdvYear = year
	//modelResponse.LastVideoClickYearOfDay
	modelResponse.LastAdvClickHour = hour
	//modelResponse.FirstDayVideoClickCount
	//modelResponse.LastDayVideoClickCount
	modelResponse.LastAdvType, _, _ = a.ICacheService.ManageCache("AdvType", firstModel.AdvType)
	modelResponse.LastMinusFirstDayVideoClickCount = modelResponse.LastDayVideoClickCount - modelResponse.FirstDayVideoClickCount
	modelResponse.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount = modelResponse.LastDayVideoClickCount - int64(modelResponse.AverageAdvDailyVideoClickCount)
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
	oldModel, err := a.IAdvEventDal.GetAdvEventByCustomerId(modelResponse.CustomerId)
	if err != nil {
		return false, err.Error()
	}
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalAdvDay = (modelResponse.LastAdvYearOfDay - oldModel.FirstAdvYearOfDay) + 365*(modelResponse.LastAdvYear-oldModel.FirstAdvYear)
	oldModel.TotalAdvCount = oldModel.TotalAdvCount + modelResponse.TotalAdvCount
	oldModel.TotalVideoAdvCount = oldModel.TotalVideoAdvCount + modelResponse.TotalVideoAdvCount
	oldModel.TotalInterstitialAdvCount = oldModel.TotalInterstitialAdvCount + modelResponse.TotalInterstitialAdvCount
	oldModel.LevelBasedAverageInterstitialAdvCount = calculateAdvLevelBasedAvgInterstitialCount(oldModel)
	oldModel.LevelBasedAverageVideoAdvCount = calculateAdvLevelBasedAvgVideoCount(oldModel)
	oldModel.AverageAdvDailyVideoClickCount = float64(oldModel.TotalVideoAdvCount) / float64(oldModel.TotalAdvDay)
	//oldModel.FirstAdvYearOfDay = oldModel.FirstAdvYearOfDay
	//oldModel.FirstAdvYear
	//oldModel.FirstAdvClickHour = oldModel.FirstAdvClickHour
	//oldModel.FirstVideoClickYearOfDay
	//oldModel.FirstVideoClickHour
	//oldModel.FirstAdvType
	oldModel.LastAdvYearOfDay = modelResponse.LastAdvYearOfDay
	oldModel.LastAdvYear = modelResponse.LastAdvYear
	//modelResponse.LastVideoClickYearOfDay
	//modelResponse.LastAdvClickHour = hour
	oldModel.LastAdvType = modelResponse.LastAdvType
	oldModel.FirstDayVideoClickCount = calculateFirstDayVideoClickCount(modelResponse, oldModel)
	oldModel.LastDayVideoClickCount = calculateLastDayVideoClickCount(modelResponse, oldModel)
	oldModel.LastMinusFirstDayVideoClickCount = oldModel.LastDayVideoClickCount - oldModel.FirstDayVideoClickCount
	oldModel.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount = oldModel.LastDayVideoClickCount - int64(oldModel.AverageAdvDailyVideoClickCount)
	oldModel.SundayVideoAdvClickCount = oldModel.SundayVideoAdvClickCount + modelResponse.SundayVideoAdvClickCount
	oldModel.MondayVideoAdvClickCount = oldModel.MondayVideoAdvClickCount + modelResponse.MondayVideoAdvClickCount
	oldModel.TuesdayVideoAdvClickCount = oldModel.TuesdayVideoAdvClickCount + modelResponse.TuesdayVideoAdvClickCount
	oldModel.WednesdayVideoAdvClickCount = oldModel.WednesdayVideoAdvClickCount + modelResponse.WednesdayVideoAdvClickCount
	oldModel.ThursdayVideoAdvClickCount = oldModel.ThursdayVideoAdvClickCount + modelResponse.ThursdayVideoAdvClickCount
	oldModel.FridayVideoAdvClickCount = oldModel.FridayVideoAdvClickCount + modelResponse.FridayVideoAdvClickCount
	oldModel.SaturdayVideoAdvClickCount = oldModel.SaturdayVideoAdvClickCount + modelResponse.SaturdayVideoAdvClickCount
	oldModel.VideoAdvClick0To5HourCount = oldModel.VideoAdvClick0To5HourCount + modelResponse.VideoAdvClick0To5HourCount
	oldModel.VideoAdvClick6To11HourCount = oldModel.VideoAdvClick6To11HourCount + modelResponse.VideoAdvClick6To11HourCount
	oldModel.VideoAdvClick12To17HourCount = oldModel.VideoAdvClick12To17HourCount + modelResponse.VideoAdvClick12To17HourCount
	oldModel.VideoAdvClick18To23HourCount = oldModel.VideoAdvClick18To23HourCount + modelResponse.VideoAdvClick18To23HourCount
	oldModel.AmVideoAdvClickCount = oldModel.AmVideoAdvClickCount + modelResponse.AmVideoAdvClickCount
	oldModel.PmVideoAdvClickCount = oldModel.PmVideoAdvClickCount + modelResponse.PmVideoAdvClickCount
	oldModel.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount = oldModel.LastDayVideoClickCount - int64(oldModel.AverageAdvDailyVideoClickCount)
	logErr := a.IAdvEventDal.UpdateAdvEventByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func calculateFirstDayVideoClickCount(modelResponse *model.AdvEventRespondModel, oldmodel *model.AdvEventRespondModel) int64 {
	if (oldmodel.FirstVideoClickYearOfDay == modelResponse.FirstVideoClickYearOfDay) && (oldmodel.FirstAdvYear == modelResponse.FirstAdvYear) {
		oldmodel.FirstDayVideoClickCount = oldmodel.FirstDayVideoClickCount + modelResponse.FirstDayVideoClickCount
		return oldmodel.FirstDayVideoClickCount
	}
	return oldmodel.FirstDayVideoClickCount
}
func calculateLastDayVideoClickCount(modelResponse *model.AdvEventRespondModel, oldmodel *model.AdvEventRespondModel) int64 {
	if (oldmodel.LastVideoClickYearOfDay == modelResponse.LastVideoClickYearOfDay) && (oldmodel.FirstVideoClickYearOfDay != modelResponse.FirstVideoClickYearOfDay) {
		oldmodel.LastDayVideoClickCount = oldmodel.LastDayVideoClickCount + modelResponse.LastDayVideoClickCount
		return oldmodel.LastDayVideoClickCount
	}
	return modelResponse.LastDayVideoClickCount
}

func determineAdvDay(modelResponse *model.AdvEventRespondModel, day int64) {
	switch day {
	case 0:
		modelResponse.SundayVideoAdvClickCount = 1
	case 1:
		modelResponse.MondayVideoAdvClickCount = 1
	case 2:
		modelResponse.TuesdayVideoAdvClickCount = 1
	case 3:
		modelResponse.WednesdayVideoAdvClickCount = 1
	case 4:
		modelResponse.ThursdayVideoAdvClickCount = 1
	case 5:
		modelResponse.FridayVideoAdvClickCount = 1
	case 6:
		modelResponse.SaturdayVideoAdvClickCount = 1
	}
}

func determineAdvHour(modelResponse *model.AdvEventRespondModel, hour int64) {
	switch {
	case hour <= 5:
		modelResponse.VideoAdvClick0To5HourCount = 1
	case (hour > 5) && (hour <= 11):
		modelResponse.VideoAdvClick6To11HourCount = 1
	case (hour > 11) && (hour <= 17):
		modelResponse.VideoAdvClick12To17HourCount = 1
	case (hour > 17) && (hour <= 23):
		modelResponse.VideoAdvClick18To23HourCount = 1
	}
}

func determineAdvAmPm(modelResponse *model.AdvEventRespondModel, hour int64) {
	switch {
	case hour <= 12:
		modelResponse.AmVideoAdvClickCount = 1
	default:
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
