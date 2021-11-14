package concrete

import (
	model "FilterWorkerService/internal/model"
	IAdvEventDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"

	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type AdvEventManager struct {
	IAdvEventDal  IAdvEventDal.IAdvEventDal
	IJsonParser   IJsonParser.IJsonParser
	ICacheService ICacheService.ICacheService
}

func (a *AdvEventManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.AdvEventModel{}
	err := a.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	hour := int64(firstModel.TrigerdTime.Hour())
	day := int64(firstModel.TrigerdTime.Weekday())
	yearOfDay := int64(firstModel.TrigerdTime.YearDay())
	year := int64(firstModel.TrigerdTime.Year())
	minute := int64(firstModel.TrigerdTime.Minute())
	modelResponse := model.AdvEventRespondModel{}
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.TotalAdvDay = 1
	modelResponse.TotalAdvCount = 1
	modelResponse.LevelBasedAverageAdvCount = CalculateAdvLevelBasedAvgClickCount(&modelResponse)
	modelResponse.AverageAdvDailyClickCount = 1
	modelResponse.FirstAdvYearOfDay = yearOfDay
	modelResponse.FirstAdvYear = year
	modelResponse.FirstAdvClickHour = hour
	modelResponse.FirstADvClickMinute = minute
	modelResponse.FirstAdvType, _, _ = a.ICacheService.ManageCache("AdvType", firstModel.AdvType)
	modelResponse.SecondAdvYearOfDay = 0
	modelResponse.SecondAdvHour = 0
	modelResponse.SecondAdvMinute = 0
	modelResponse.ThirdAdvYearOfDay = 0
	modelResponse.ThirdAdvHour = 0
	modelResponse.ThirdAdvMinute = 0
	modelResponse.PenultimateAdvYearOfDay = 0
	modelResponse.PenultimateAdvHour = 0
	modelResponse.PenultimateAdvMinute = 0
	modelResponse.LastAdvYearOfDay = yearOfDay
	modelResponse.LastAdvYear = year
	modelResponse.LastAdvClickHour = hour
	modelResponse.LastAdvClickMinute = minute
	modelResponse.LastAdvType, _, _ = a.ICacheService.ManageCache("AdvType", firstModel.AdvType)
	modelResponse.FirstDayAdvClickCount = 1
	modelResponse.PenultimateDayAdvClickCount = 0
	modelResponse.LastDayAdvClickCount = 1
	modelResponse.LastMinusFirstDayAdvClickCount = 0
	modelResponse.LastMinusPenultimateDayAdvClickCount = 0
	modelResponse.LastDayAdvClickCountMinusAverageDailyAdvClickCount = 0
	DetermineAdvDay(&modelResponse, day)
	DetermineAdvHour(&modelResponse, hour)
	DetermineAdvAmPm(&modelResponse, hour)

	oldModel, err := a.IAdvEventDal.GetAdvEventByCustomerId(modelResponse.CustomerId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := a.IAdvEventDal.Add(&modelResponse)
		if logErr != nil {
			return false, logErr.Error()
		}
		return true, ""

	case err == nil:
		updateResult, updateErr := a.updateAdvEvent(&modelResponse, oldModel)
		if updateErr != nil {
			return updateResult, updateErr.Error()
		}
		return updateResult, ""

	default:

		return false, err.Error()

	}

}


func (a *AdvEventManager) updateAdvEvent(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalAdvDay = (modelResponse.LastAdvYearOfDay - oldModel.FirstAdvYearOfDay) + 365*(modelResponse.LastAdvYear-oldModel.FirstAdvYear)
	oldModel.TotalAdvCount = oldModel.TotalAdvCount + modelResponse.TotalAdvCount
	oldModel.LevelBasedAverageAdvCount = CalculateAdvLevelBasedAvgClickCount(oldModel)
	oldModel.AverageAdvDailyClickCount = float64(oldModel.TotalAdvCount) / float64(oldModel.TotalAdvDay)
	//oldModel.FirstAdvYearOfDay = oldModel.FirstAdvYearOfDay
	//oldModel.FirstAdvYear
	//oldModel.FirstAdvClickHour = oldModel.FirstAdvClickHour
	//oldModel.FirstADvClickMinute
	//oldModel.FirstAdvType
	oldModel.SecondAdvYearOfDay, oldModel.SecondAdvHour, oldModel.SecondAdvMinute = CalculateSecondAdv(modelResponse, oldModel)
	oldModel.ThirdAdvYearOfDay, oldModel.ThirdAdvHour, oldModel.ThirdAdvMinute = CalculateThirdAdv(modelResponse, oldModel)
	oldModel.PenultimateAdvYearOfDay = oldModel.LastAdvYearOfDay
	oldModel.PenultimateAdvHour = oldModel.LastAdvClickHour
	oldModel.PenultimateAdvMinute = oldModel.LastAdvClickMinute
	oldModel.LastAdvYearOfDay = modelResponse.LastAdvYearOfDay
	oldModel.LastAdvYear = modelResponse.LastAdvYear
	oldModel.LastAdvClickHour = modelResponse.LastAdvClickHour
	oldModel.LastAdvClickMinute = modelResponse.LastAdvClickMinute
	oldModel.LastAdvType = modelResponse.LastAdvType
	oldModel.FirstDayAdvClickCount = CalculateFirstDayAdvClickCount(modelResponse, oldModel)
	oldModel.PenultimateDayAdvClickCount = CalculatePenultimateDayAdvDay(modelResponse, oldModel)
	oldModel.LastDayAdvClickCount = CalculateLastDayAdvClickCount(modelResponse, oldModel)
	oldModel.LastMinusFirstDayAdvClickCount = oldModel.LastDayAdvClickCount - oldModel.FirstDayAdvClickCount
	oldModel.LastMinusPenultimateDayAdvClickCount = oldModel.LastDayAdvClickCount - oldModel.PenultimateDayAdvClickCount
	oldModel.LastDayAdvClickCountMinusAverageDailyAdvClickCount = oldModel.LastDayAdvClickCount - int64(oldModel.AverageAdvDailyClickCount)
	oldModel.SundayAdvClickCount = oldModel.SundayAdvClickCount + modelResponse.SundayAdvClickCount
	oldModel.MondayAdvClickCount = oldModel.MondayAdvClickCount + modelResponse.MondayAdvClickCount
	oldModel.TuesdayAdvClickCount = oldModel.TuesdayAdvClickCount + modelResponse.TuesdayAdvClickCount
	oldModel.WednesdayAdvClickCount = oldModel.WednesdayAdvClickCount + modelResponse.WednesdayAdvClickCount
	oldModel.ThursdayAdvClickCount = oldModel.ThursdayAdvClickCount + modelResponse.ThursdayAdvClickCount
	oldModel.FridayAdvClickCount = oldModel.FridayAdvClickCount + modelResponse.FridayAdvClickCount
	oldModel.SaturdayAdvClickCount = oldModel.SaturdayAdvClickCount + modelResponse.SaturdayAdvClickCount
	oldModel.AdvClick0To5HourCount = oldModel.AdvClick0To5HourCount + modelResponse.AdvClick0To5HourCount
	oldModel.AdvClick6To11HourCount = oldModel.AdvClick6To11HourCount + modelResponse.AdvClick6To11HourCount
	oldModel.AdvClick12To17HourCount = oldModel.AdvClick12To17HourCount + modelResponse.AdvClick12To17HourCount
	oldModel.AdvClick18To23HourCount = oldModel.AdvClick18To23HourCount + modelResponse.AdvClick18To23HourCount
	oldModel.AmAdvClickCount = oldModel.AmAdvClickCount + modelResponse.AmAdvClickCount
	oldModel.PmAdvClickCount = oldModel.PmAdvClickCount + modelResponse.PmAdvClickCount

	logErr := a.IAdvEventDal.UpdateAdvEventByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr
	}
	return true, nil
}

func CalculateSecondAdv(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (day int64, hour int64, minute int64) {
	if oldModel.TotalAdvCount == 2 {
		oldModel.SecondAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.SecondAdvHour = modelResponse.FirstAdvClickHour
		oldModel.SecondAdvMinute = modelResponse.FirstADvClickMinute
		return oldModel.SecondAdvYearOfDay, oldModel.SecondAdvHour, oldModel.SecondAdvMinute
	}
	return oldModel.SecondAdvYearOfDay, oldModel.SecondAdvHour, oldModel.SecondAdvMinute
}

func CalculateThirdAdv(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (day int64, hour int64, minute int64) {
	if oldModel.TotalAdvCount == 3 {
		oldModel.ThirdAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.ThirdAdvHour = modelResponse.FirstAdvClickHour
		oldModel.ThirdAdvMinute = modelResponse.FirstADvClickMinute
		return oldModel.ThirdAdvYearOfDay, oldModel.ThirdAdvHour, oldModel.ThirdAdvMinute
	}
	return oldModel.ThirdAdvYearOfDay, oldModel.ThirdAdvHour, oldModel.ThirdAdvMinute
}

func CalculateFirstDayAdvClickCount(modelResponse *model.AdvEventRespondModel, oldmodel *model.AdvEventRespondModel) int64 {
	if (oldmodel.FirstAdvYearOfDay == modelResponse.FirstAdvYearOfDay) && (oldmodel.FirstAdvYear == modelResponse.FirstAdvYear) {
		oldmodel.FirstDayAdvClickCount = oldmodel.FirstDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldmodel.FirstDayAdvClickCount
	}
	return oldmodel.FirstDayAdvClickCount
}

func CalculatePenultimateDayAdvDay(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (count int64) {
	if (oldModel.LastAdvYearOfDay != modelResponse.LastAdvYearOfDay) && (oldModel.FirstAdvYearOfDay != modelResponse.FirstAdvYearOfDay) {
		oldModel.PenultimateDayAdvClickCount = oldModel.LastDayAdvClickCount
		return oldModel.PenultimateDayAdvClickCount
	}
	return oldModel.PenultimateDayAdvClickCount
}

func CalculateLastDayAdvClickCount(modelResponse *model.AdvEventRespondModel, oldmodel *model.AdvEventRespondModel) int64 {
	if (oldmodel.LastAdvYearOfDay == modelResponse.LastAdvYearOfDay) && (oldmodel.FirstAdvYearOfDay != modelResponse.FirstAdvYearOfDay) {
		oldmodel.LastDayAdvClickCount = oldmodel.LastDayAdvClickCount + modelResponse.LastDayAdvClickCount
		return oldmodel.LastDayAdvClickCount

	} else if (oldmodel.LastDayAdvClickCount != modelResponse.LastDayAdvClickCount) && (oldmodel.FirstAdvYearOfDay != modelResponse.FirstAdvYearOfDay) {
		return modelResponse.LastDayAdvClickCount
	}

	return oldmodel.LastDayAdvClickCount
}

func DetermineAdvDay(modelResponse *model.AdvEventRespondModel, day int64) {
	switch day {
	case 0:
		modelResponse.SundayAdvClickCount = 1
	case 1:
		modelResponse.MondayAdvClickCount = 1
	case 2:
		modelResponse.TuesdayAdvClickCount = 1
	case 3:
		modelResponse.WednesdayAdvClickCount = 1
	case 4:
		modelResponse.ThursdayAdvClickCount = 1
	case 5:
		modelResponse.FridayAdvClickCount = 1
	case 6:
		modelResponse.SaturdayAdvClickCount = 1
	}
}

func DetermineAdvHour(modelResponse *model.AdvEventRespondModel, hour int64) {
	switch {
	case hour <= 5:
		modelResponse.AdvClick0To5HourCount = 1
	case (hour > 5) && (hour <= 11):
		modelResponse.AdvClick6To11HourCount = 1
	case (hour > 11) && (hour <= 17):
		modelResponse.AdvClick12To17HourCount = 1
	case (hour > 17) && (hour <= 23):
		modelResponse.AdvClick18To23HourCount = 1
	}
}

func DetermineAdvAmPm(modelResponse *model.AdvEventRespondModel, hour int64) {
	switch {
	case hour <= 12:
		modelResponse.AmAdvClickCount = 1
	default:
		modelResponse.PmAdvClickCount = 1
	}
}

func CalculateAdvLevelBasedAvgClickCount(modelResponse *model.AdvEventRespondModel) float64 {
	if modelResponse.LevelIndex == 0 {
		modelResponse.LevelBasedAverageAdvCount = float64(modelResponse.TotalAdvCount)
		return modelResponse.LevelBasedAverageAdvCount
	}
	modelResponse.LevelBasedAverageAdvCount = float64(modelResponse.TotalAdvCount) / float64(modelResponse.LevelIndex)
	return modelResponse.LevelBasedAverageAdvCount
}
