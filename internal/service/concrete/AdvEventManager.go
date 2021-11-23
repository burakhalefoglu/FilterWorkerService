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

func (a *AdvEventManager) ConvertRawModelToResponseModel(data *[]byte) (adv *model.AdvEventRespondModel, s bool, m string) {
	firstModel := model.AdvEventModel{}
	err := a.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return nil, false, err.Error()
	}
	hour := int64(firstModel.TrigerdTime.Hour())
	day := int64(firstModel.TrigerdTime.Weekday())
	yearOfDay := int64(firstModel.TrigerdTime.YearDay())
	year := int64(firstModel.TrigerdTime.Year())
	minute := int64(firstModel.InMinutes)
	modelResponse := model.AdvEventRespondModel{}
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.TotalAdvDay = 1
	modelResponse.TotalAdvCount = 1
	modelResponse.TotalAdvHour = 0
	modelResponse.TotalAdvMinute = 1
	CalculateAdvLevelBasedAvgClickCount(&modelResponse)
	modelResponse.AverageAdvDailyClickCount = 1
	modelResponse.FirstAdvYearOfDay = yearOfDay
	modelResponse.FirstAdvYear = year
	modelResponse.FirstWeekDay = day
	modelResponse.FirstAdvClickHour = hour
	modelResponse.FirstADvClickMinute = minute
	modelResponse.FirstAdvType, _, _ = a.ICacheService.ManageCache("AdvType", firstModel.AdvType)
	modelResponse.SecondAdvYearOfDay = 0
	modelResponse.SecondAdvHour = 0
	modelResponse.SecondAdvMinute = 0
	modelResponse.SecondAdvType = 0
	modelResponse.ThirdAdvYearOfDay = 0
	modelResponse.ThirdAdvHour = 0
	modelResponse.ThirdAdvMinute = 0
	modelResponse.ThirdAdvType = 0
	modelResponse.FourthAdvYearOfDay = 0
	modelResponse.FourthAdvHour = 0
	modelResponse.FourthAdvMinute = 0
	modelResponse.FourthAdvType = 0
	modelResponse.FifthAdvYearOfDay = 0
	modelResponse.FifthAdvHour = 0
	modelResponse.FifthAdvMinute = 0
	modelResponse.FifthAdvType = 0
	modelResponse.PenultimateAdvYearOfDay = 0
	modelResponse.PenultimateAdvHour = 0
	modelResponse.PenultimateAdvMinute = 0
	modelResponse.PenultimateAdvType = 0
	modelResponse.LastAdvYearOfDay = 0
	modelResponse.LastAdvYear = 0
	modelResponse.LastAdvClickHour = 0
	modelResponse.LastAdvClickMinute = 0
	modelResponse.LastAdvType = 0

	modelResponse.FirstHalfHourAdvClickCount = 1
	modelResponse.FirstHourAdvClickCount = 1
	modelResponse.FirstTwoHourAdvClickCount = 1
	modelResponse.FirstThreeHourAdvClickCount = 1
	modelResponse.FirstSixHourAdvClickCount = 1
	modelResponse.FirstTwelveHourAdvClickCount = 1

	modelResponse.FirstDayAdvClickCount = 1
	modelResponse.SecondDayAdvClickCount = 0
	modelResponse.ThirdDayAdvClickCount = 0
	modelResponse.FourthDayAdvClickCount = 0
	modelResponse.FifthDayAdvClickCount = 0
	modelResponse.SixthDayAdvClickCount = 0
	modelResponse.SeventhDayAdvClickCount = 0

	modelResponse.PenultimateDayAdvClickCount = 0
	modelResponse.LastDayAdvClickCount = 0
	modelResponse.LastMinusFirstDayAdvClickCount = -1
	modelResponse.LastMinusPenultimateDayAdvClickCount = 0
	modelResponse.LastDayAdvClickCountMinusAverageDailyAdvClickCount = -1
	DetermineAdvDay(&modelResponse, day)
	DetermineAdvHour(&modelResponse, hour)
	DetermineAdvAmPm(&modelResponse, hour)

	oldModel, err := a.IAdvEventDal.GetAdvEventById(modelResponse.ClientId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := a.IAdvEventDal.Add(&modelResponse)
		if logErr != nil {
			return nil, false, logErr.Error()
		}
		return &modelResponse, true, "Added"

	case err == nil:
		updatedModel, updateResult, updateErr := a.UpdateAdvEvent(&modelResponse, oldModel)
		if updateErr != nil {
			return nil, updateResult, "Update went wrong!"
		}
		return updatedModel, updateResult, "Updated"

	default:
		return nil, false, err.Error()
	}

}

func (a *AdvEventManager) UpdateAdvEvent(modelResponse *model.AdvEventRespondModel,
	oldModel *model.AdvEventRespondModel) (updatedModel *model.AdvEventRespondModel, s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalAdvDay = (modelResponse.FirstAdvYearOfDay - oldModel.FirstAdvYearOfDay) + 365*(modelResponse.FirstAdvYear-oldModel.FirstAdvYear)
	oldModel.TotalAdvCount = oldModel.TotalAdvCount + modelResponse.TotalAdvCount
	oldModel.TotalAdvHour = ((modelResponse.FirstAdvYearOfDay+365*modelResponse.FirstAdvYear)*24 + modelResponse.FirstAdvClickHour) - ((oldModel.FirstAdvYearOfDay+365*oldModel.FirstAdvYear)*24 + oldModel.FirstAdvClickHour)
	oldModel.TotalAdvMinute = (((modelResponse.FirstAdvYearOfDay+365*modelResponse.FirstAdvYear)*24+modelResponse.FirstAdvClickHour)*60 + modelResponse.FirstADvClickMinute) - (((oldModel.FirstAdvYearOfDay+365*oldModel.FirstAdvYear)*24+oldModel.FirstAdvClickHour)*60 + oldModel.FirstADvClickMinute)
	CalculateAdvLevelBasedAvgClickCount(oldModel)
	oldModel.AverageAdvDailyClickCount = CalculateAverageAdvDailyClickCount(oldModel)
	//oldModel.FirstAdvYearOfDay = oldModel.FirstAdvYearOfDay
	//oldModel.FirstAdvYear
	//oldModel.FirstWeekDay
	//oldModel.FirstAdvClickHour = oldModel.FirstAdvClickHour
	//oldModel.FirstADvClickMinute
	//oldModel.FirstAdvType
	oldModel.SecondAdvYearOfDay, oldModel.SecondAdvHour, oldModel.SecondAdvMinute, oldModel.SecondAdvType = CalculateSecondAdv(modelResponse, oldModel)
	oldModel.ThirdAdvYearOfDay, oldModel.ThirdAdvHour, oldModel.ThirdAdvMinute, oldModel.ThirdAdvType = CalculateThirdAdv(modelResponse, oldModel)
	modelResponse.FourthAdvYearOfDay, modelResponse.FourthAdvHour, modelResponse.FourthAdvMinute, modelResponse.FourthAdvType = CalculateFourthAdv(modelResponse, oldModel)
	modelResponse.FifthAdvYearOfDay, modelResponse.FifthAdvHour, modelResponse.FifthAdvMinute, modelResponse.FifthAdvType = CalculateFifthAdv(modelResponse, oldModel)

	oldModel.PenultimateDayAdvClickCount = CalculatePenultimateDayAdvDay(modelResponse, oldModel)
	oldModel.LastDayAdvClickCount = CalculateLastDayAdvClickCount(modelResponse, oldModel)

	oldModel.PenultimateAdvYearOfDay = oldModel.LastAdvYearOfDay
	oldModel.PenultimateAdvHour = oldModel.LastAdvClickHour
	oldModel.PenultimateAdvMinute = oldModel.LastAdvClickMinute
	oldModel.LastAdvYearOfDay = modelResponse.FirstAdvYearOfDay
	oldModel.LastAdvYear = modelResponse.FirstAdvYear
	oldModel.LastAdvClickHour = modelResponse.FirstAdvClickHour
	oldModel.LastAdvClickMinute = modelResponse.FirstADvClickMinute
	oldModel.LastAdvType = modelResponse.FirstAdvType

	oldModel.FirstHalfHourAdvClickCount = CalculateFirstHalfHourTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	oldModel.FirstHourAdvClickCount = CalculateFirstHourTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	oldModel.FirstTwoHourAdvClickCount = CalculateFirstTwoHourTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	oldModel.FirstThreeHourAdvClickCount = CalculateFirstThreeHourAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	oldModel.FirstSixHourAdvClickCount = CalculateFirstSixHourAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	oldModel.FirstTwelveHourAdvClickCount = CalculateFirstTwelveHourAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)

	oldModel.FirstDayAdvClickCount = CalculateFirstDayAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	oldModel.SecondDayAdvClickCount = CalculateSecondDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	oldModel.ThirdDayAdvClickCount = CalculateThirdDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	oldModel.FourthDayAdvClickCount = CalculateFourthDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	oldModel.FifthDayAdvClickCount = CalculateFifthDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	oldModel.SixthDayAdvClickCount = CalculateSixthDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	oldModel.SeventhDayAdvClickCount = CalculateSeventhDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)

	oldModel.LastMinusFirstDayAdvClickCount = oldModel.LastDayAdvClickCount - oldModel.FirstDayAdvClickCount
	oldModel.LastMinusPenultimateDayAdvClickCount = oldModel.LastDayAdvClickCount - oldModel.PenultimateDayAdvClickCount
	oldModel.LastDayAdvClickCountMinusAverageDailyAdvClickCount = float64(oldModel.LastDayAdvClickCount) - oldModel.AverageAdvDailyClickCount

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

	logErr := a.IAdvEventDal.UpdateAdvEventById(oldModel.ClientId, oldModel)
	if logErr != nil {
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateAverageAdvDailyClickCount(oldModel *model.AdvEventRespondModel) (count float64){
	if oldModel.TotalAdvDay == 0{
		return float64(oldModel.TotalAdvCount)
	}
	return float64(oldModel.TotalAdvCount) / float64(oldModel.TotalAdvDay)
}

func CalculateFirstHalfHourTotalAdvCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_minute int64) (count int64) {
	if total_adv_minute <= 30 {
		oldModel.FirstHalfHourAdvClickCount = oldModel.FirstHalfHourAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FirstHalfHourAdvClickCount
	}
	return oldModel.FirstHalfHourAdvClickCount
}

func CalculateFirstHourTotalAdvCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_minute int64) (count int64) {
	if total_adv_minute <= 60 {
		oldModel.FirstHourAdvClickCount = oldModel.FirstHourAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FirstHourAdvClickCount
	}
	return oldModel.FirstHourAdvClickCount
}

func CalculateFirstTwoHourTotalAdvCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_minute int64) (count int64) {
	if total_adv_minute <= 120 {
		oldModel.FirstTwoHourAdvClickCount = oldModel.FirstTwoHourAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FirstTwoHourAdvClickCount
	}
	return oldModel.FirstTwoHourAdvClickCount
}

func CalculateFirstThreeHourAdvClickCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_minute int64) (count int64) {
	if total_adv_minute <= 180 {
		oldModel.FirstThreeHourAdvClickCount = oldModel.FirstThreeHourAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FirstThreeHourAdvClickCount
	}
	return oldModel.FirstThreeHourAdvClickCount
}

func CalculateFirstSixHourAdvClickCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_minute int64) (count int64) {
	if total_adv_minute <= 360 {
		oldModel.FirstSixHourAdvClickCount = oldModel.FirstSixHourAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FirstSixHourAdvClickCount
	}
	return oldModel.FirstSixHourAdvClickCount
}

func CalculateFirstTwelveHourAdvClickCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_minute int64) (count int64) {
	if total_adv_minute <= 720 {
		oldModel.FirstTwelveHourAdvClickCount = oldModel.FirstTwelveHourAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FirstSixHourAdvClickCount
	}
	return oldModel.FirstSixHourAdvClickCount
}

func CalculateSecondDayTotalSessionCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_hour int64) (count int64) {
	if total_adv_hour <= 48 && total_adv_hour > 24 {
		oldModel.SecondDayAdvClickCount = oldModel.SecondDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.SecondDayAdvClickCount
	}
	return oldModel.SecondDayAdvClickCount
}

func CalculateThirdDayTotalSessionCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_hour int64) (count int64) {
	if total_adv_hour <= 72 && total_adv_hour > 48 {
		oldModel.ThirdDayAdvClickCount = oldModel.ThirdDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.ThirdDayAdvClickCount
	}
	return oldModel.ThirdDayAdvClickCount
}

func CalculateFourthDayTotalSessionCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_hour int64) (count int64) {
	if total_adv_hour <= 96 && total_adv_hour > 72 {
		oldModel.FourthDayAdvClickCount = oldModel.FourthDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FourthDayAdvClickCount
	}
	return oldModel.FourthDayAdvClickCount
}

func CalculateFifthDayTotalSessionCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_hour int64) (count int64) {
	if total_adv_hour <= 120 && total_adv_hour > 96 {
		oldModel.FifthDayAdvClickCount = oldModel.FifthDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FifthDayAdvClickCount
	}
	return oldModel.FifthDayAdvClickCount
}

func CalculateSixthDayTotalSessionCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_hour int64) (count int64) {
	if total_adv_hour <= 144 && total_adv_hour > 120 {
		oldModel.SixthDayAdvClickCount = oldModel.SixthDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.SixthDayAdvClickCount
	}
	return oldModel.SixthDayAdvClickCount
}

func CalculateSeventhDayTotalSessionCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_hour int64) (count int64) {
	if total_adv_hour <= 168 && total_adv_hour > 144 {
		oldModel.SeventhDayAdvClickCount = oldModel.SeventhDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.SeventhDayAdvClickCount
	}
	return oldModel.SeventhDayAdvClickCount
}

func CalculateSecondAdv(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (day int64, hour int64, minute int64, Advtype int64) {
	if oldModel.TotalAdvCount == 2 {
		oldModel.SecondAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.SecondAdvHour = modelResponse.FirstAdvClickHour
		oldModel.SecondAdvMinute = modelResponse.FirstADvClickMinute
		oldModel.SecondAdvType = modelResponse.FirstAdvType
		return oldModel.SecondAdvYearOfDay, oldModel.SecondAdvHour, oldModel.SecondAdvMinute, oldModel.SecondAdvType
	}
	return oldModel.SecondAdvYearOfDay, oldModel.SecondAdvHour, oldModel.SecondAdvMinute, oldModel.SecondAdvType
}

func CalculateThirdAdv(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (day int64, hour int64, minute int64, Atype int64) {
	if oldModel.TotalAdvCount == 3 {
		oldModel.ThirdAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.ThirdAdvHour = modelResponse.FirstAdvClickHour
		oldModel.ThirdAdvMinute = modelResponse.FirstADvClickMinute
		oldModel.ThirdAdvType = modelResponse.FirstAdvType
		return oldModel.ThirdAdvYearOfDay, oldModel.ThirdAdvHour, oldModel.ThirdAdvMinute, oldModel.ThirdAdvType
	}
	return oldModel.ThirdAdvYearOfDay, oldModel.ThirdAdvHour, oldModel.ThirdAdvMinute, oldModel.ThirdAdvType
}

func CalculateFourthAdv(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (day int64, hour int64, minute int64, Atype int64) {
	if oldModel.TotalAdvCount == 4 {
		oldModel.FourthAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.FourthAdvHour = modelResponse.FirstAdvClickHour
		oldModel.FourthAdvMinute = modelResponse.FirstADvClickMinute
		oldModel.FourthAdvType = modelResponse.FirstAdvType
		return oldModel.FourthAdvYearOfDay, oldModel.FourthAdvHour, oldModel.FourthAdvMinute, oldModel.FourthAdvType
	}
	return oldModel.FourthAdvYearOfDay, oldModel.FourthAdvHour, oldModel.FourthAdvMinute, oldModel.FourthAdvType
}

func CalculateFifthAdv(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (day int64, hour int64, minute int64, Atype int64) {
	if oldModel.TotalAdvCount == 5 {
		oldModel.FifthAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.FifthAdvHour = modelResponse.FirstAdvClickHour
		oldModel.FifthAdvMinute = modelResponse.FirstADvClickMinute
		oldModel.FifthAdvType = modelResponse.FirstAdvType
		return oldModel.FifthAdvYearOfDay, oldModel.FifthAdvHour, oldModel.FifthAdvMinute, oldModel.FifthAdvType
	}
	return oldModel.FifthAdvYearOfDay, oldModel.FifthAdvHour, oldModel.FifthAdvMinute, oldModel.FifthAdvType
}

func CalculateFirstDayAdvClickCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel, total_adv_minute int64) int64 {
	if total_adv_minute <= 1440 {
		oldModel.FirstDayAdvClickCount = oldModel.FirstDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.FirstDayAdvClickCount
	}
	return oldModel.FirstDayAdvClickCount
}

func CalculatePenultimateDayAdvDay(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) (count int64) {
	switch {
	case ((modelResponse.FirstAdvYearOfDay + (365 * modelResponse.FirstAdvYear)) > (oldModel.LastAdvYearOfDay + (365 * oldModel.LastAdvYear))) && ((modelResponse.FirstAdvYearOfDay + 365*modelResponse.FirstAdvYear) != (oldModel.FirstAdvYearOfDay + 365*oldModel.FirstAdvYear)):
		return oldModel.LastDayAdvClickCount

	default:
		return oldModel.PenultimateDayAdvClickCount
	}
}

func CalculateLastDayAdvClickCount(modelResponse *model.AdvEventRespondModel, oldModel *model.AdvEventRespondModel) int64 {
	switch {
	case ((oldModel.LastAdvYearOfDay + 365*(oldModel.LastAdvYear)) == (modelResponse.FirstAdvYearOfDay + 365*(modelResponse.FirstAdvYear))):
		oldModel.LastDayAdvClickCount = oldModel.LastDayAdvClickCount + modelResponse.FirstDayAdvClickCount
		return oldModel.LastDayAdvClickCount

	case ((modelResponse.FirstAdvYearOfDay + (365 * modelResponse.FirstAdvYear)) > (oldModel.LastAdvYearOfDay + (365 * oldModel.LastAdvYear))) && ((modelResponse.FirstAdvYearOfDay + 365*modelResponse.FirstAdvYear) != (oldModel.FirstAdvYearOfDay + 365*oldModel.FirstAdvYear)):
		return modelResponse.FirstDayAdvClickCount

	default:
		return oldModel.LastDayAdvClickCount
	}
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

func CalculateAdvLevelBasedAvgClickCount(modelResponse *model.AdvEventRespondModel) {
	switch modelResponse.LevelIndex {
	case 0:
		modelResponse.LevelBasedAverageAdvCount = float64(modelResponse.TotalAdvCount)
	default:
		modelResponse.LevelBasedAverageAdvCount = float64(modelResponse.TotalAdvCount) / float64(modelResponse.LevelIndex)

	}

}
