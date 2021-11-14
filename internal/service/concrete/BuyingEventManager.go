package concrete

import (
	model "FilterWorkerService/internal/model"
	IBuyingEventDal "FilterWorkerService/internal/repository/abstract"

	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type BuyingEventManager struct {
	IBuyingEventDal IBuyingEventDal.IBuyingEventDal
	IJsonParser     IJsonParser.IJsonParser
}

func (b *BuyingEventManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.BuyingEventModel{}
	err := b.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	hour := int64(firstModel.TrigerdTime.Hour())
	day := int64(firstModel.TrigerdTime.Weekday())
	yearOfDay := int64(firstModel.TrigerdTime.YearDay())
	year := int64(firstModel.TrigerdTime.Year())
	modelResponse := model.BuyingEventRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.TotalBuyingCount = 1
	modelResponse.TotalBuyingDay = 1
	modelResponse.FirstBuyingYearOfDay = yearOfDay
	modelResponse.FirstBuyingYear = year
	modelResponse.FirstBuyingHour = hour
	modelResponse.SecondBuyingYearOfDay = 0
	modelResponse.SecondBuyingHour = 0
	modelResponse.ThirdBuyingYearOfDay = 0
	modelResponse.ThirdBuyingHour = 0
	modelResponse.PenultimateBuyingYearOfDay = 0
	modelResponse.PenultimateBuyingHour = 0
	modelResponse.LastBuyingYearOfDay = yearOfDay
	modelResponse.LastBuyingYear = year
	modelResponse.LastBuyingHour = hour
	modelResponse.FirstDayBuyingCount = 1
	modelResponse.PenultimateDayBuyingCount = 0
	modelResponse.LastDayBuyingCount = 1
	modelResponse.LastMinusFirstDayBuyingCount = modelResponse.LastDayBuyingCount - modelResponse.FirstDayBuyingCount
	determineBuyingDay(&modelResponse, day)
	determineBuyingHour(&modelResponse, hour)
	determineBuyingAmPm(&modelResponse, hour)
	modelResponse.BuyingDayAverageBuyingCount = float64(modelResponse.TotalBuyingCount) / float64(modelResponse.TotalBuyingDay)
	modelResponse.LevelBasedAverageBuyingCount = calculateBuyingLevelBasedAvgBuyingCount(&modelResponse)


	oldModel, err := b.IBuyingEventDal.GetBuyingEventByCustomerId(modelResponse.CustomerId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := b.IBuyingEventDal.Add(&modelResponse)
		if logErr != nil {
			return false, logErr.Error()
		}
		return true, ""

	case err == nil:
		updateResult, updateErr := b.updateBuyingEventByCustomerId(&modelResponse, oldModel)
		if updateErr != nil {
		return updateResult, updateErr.Error()
	}
		return updateResult, ""

	default:

		return false, err.Error()

	}
	
}


func (b *BuyingEventManager) updateBuyingEventByCustomerId(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (s bool, m error) {
	
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalBuyingCount = oldModel.TotalBuyingCount + modelResponse.TotalBuyingCount
	oldModel.TotalBuyingDay = (modelResponse.LastBuyingYearOfDay - oldModel.FirstBuyingYearOfDay) + 365*(modelResponse.LastBuyingYear-oldModel.FirstBuyingYear)
	//oldModel.FirstBuyingYearOfDay
	//oldModel.FirstBuyingYear
	//oldModel.FirstBuyingHour
	oldModel.SecondBuyingYearOfDay,oldModel.SecondBuyingHour = calculateSecondBuying(modelResponse, oldModel)
	oldModel.ThirdBuyingYearOfDay,modelResponse.ThirdBuyingHour = calculateThirdBuying(modelResponse, oldModel)
	oldModel.PenultimateBuyingYearOfDay = oldModel.LastBuyingYearOfDay
	oldModel.PenultimateBuyingHour = oldModel.LastBuyingHour
	oldModel.LastBuyingYearOfDay = modelResponse.LastBuyingYearOfDay
	oldModel.LastBuyingYear = modelResponse.LastBuyingYear
	oldModel.LastBuyingHour = modelResponse.LastBuyingHour
	oldModel.FirstDayBuyingCount = calculateFirstDayBuyingCount(modelResponse, oldModel)
	oldModel.PenultimateDayBuyingCount = calculatePenultimateDayBuyingCount(modelResponse, oldModel)
	oldModel.LastDayBuyingCount = calculateLastDayBuyingCount(modelResponse, oldModel)	
	oldModel.LastMinusFirstDayBuyingCount = oldModel.LastDayBuyingCount - oldModel.FirstDayBuyingCount
	oldModel.SundayBuyingCount = oldModel.SundayBuyingCount + modelResponse.SundayBuyingCount
	oldModel.MondayBuyingCount = oldModel.MondayBuyingCount + modelResponse.MondayBuyingCount
	oldModel.TuesdayBuyingCount = oldModel.TuesdayBuyingCount + modelResponse.TuesdayBuyingCount
	oldModel.WednesdayBuyingCount = oldModel.WednesdayBuyingCount + modelResponse.WednesdayBuyingCount
	oldModel.ThursdayBuyingCount = oldModel.ThursdayBuyingCount + modelResponse.ThursdayBuyingCount
	oldModel.FridayBuyingCount = oldModel.FridayBuyingCount + modelResponse.FridayBuyingCount
	oldModel.SaturdayBuyingCount = oldModel.SaturdayBuyingCount + modelResponse.SaturdayBuyingCount
	oldModel.AmBuyingCount = oldModel.AmBuyingCount + modelResponse.AmBuyingCount
	oldModel.PmBuyingCount = oldModel.PmBuyingCount + modelResponse.PmBuyingCount
	oldModel.Buying0To5HourCount = oldModel.Buying0To5HourCount + modelResponse.Buying0To5HourCount
	oldModel.Buying6To11HourCount = oldModel.Buying6To11HourCount + modelResponse.Buying6To11HourCount
	oldModel.Buying12To17HourCount = oldModel.Buying12To17HourCount + modelResponse.Buying12To17HourCount
	oldModel.Buying18To23HourCount = oldModel.Buying18To23HourCount + modelResponse.Buying18To23HourCount
	oldModel.BuyingDayAverageBuyingCount = float64(oldModel.TotalBuyingCount) / float64(oldModel.TotalBuyingDay)
	oldModel.LevelBasedAverageBuyingCount = calculateBuyingLevelBasedAvgBuyingCount(oldModel)
	logErr := b.IBuyingEventDal.UpdateBuyingEventByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr
	}
	return true, nil
}

func calculateSecondBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (day int64, hour int64) {
	if oldModel.TotalBuyingCount == 2 {
		oldModel.SecondBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.SecondBuyingHour = modelResponse.FirstBuyingHour
		return oldModel.SecondBuyingYearOfDay, oldModel.SecondBuyingHour
	}
	return oldModel.SecondBuyingYearOfDay, oldModel.SecondBuyingHour
}

func calculateThirdBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (day int64, hour int64) {
	if oldModel.TotalBuyingCount == 3 {
		oldModel.ThirdBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.ThirdBuyingHour = modelResponse.FirstBuyingHour
		return oldModel.ThirdBuyingYearOfDay, oldModel.ThirdBuyingHour
	}
	return oldModel.ThirdBuyingYearOfDay, oldModel.ThirdBuyingHour
}


func calculatePenultimateDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (count int64) {
	if (oldModel.LastBuyingYearOfDay != modelResponse.LastBuyingYearOfDay) && (oldModel.FirstBuyingYearOfDay != modelResponse.FirstBuyingYearOfDay) {
		oldModel.PenultimateDayBuyingCount = oldModel.LastDayBuyingCount
		return oldModel.PenultimateDayBuyingCount
	}
	return oldModel.PenultimateDayBuyingCount
}



func calculateFirstDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) int64 {
	if (oldModel.FirstBuyingYearOfDay == modelResponse.FirstBuyingYearOfDay) && (oldModel.FirstBuyingYear == modelResponse.FirstBuyingYear) {
		oldModel.FirstDayBuyingCount = oldModel.FirstDayBuyingCount + modelResponse.FirstDayBuyingCount
		return oldModel.FirstDayBuyingCount
	}
	return oldModel.FirstDayBuyingCount
}

func calculateLastDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) int64 {
	if (oldModel.LastBuyingYearOfDay == modelResponse.LastBuyingYearOfDay) && (oldModel.FirstBuyingYearOfDay != modelResponse.FirstBuyingYearOfDay) {
		oldModel.LastDayBuyingCount = oldModel.LastDayBuyingCount + modelResponse.LastDayBuyingCount
		return oldModel.LastDayBuyingCount

	} else if (oldModel.LastBuyingYearOfDay != modelResponse.LastBuyingYearOfDay) && (oldModel.FirstBuyingYearOfDay != modelResponse.FirstBuyingYearOfDay) {
		return modelResponse.LastDayBuyingCount
	}

	return oldModel.LastDayBuyingCount
}

func determineBuyingDay(modelResponse *model.BuyingEventRespondModel, day int64) {
	switch day {
	case 0:
		modelResponse.SundayBuyingCount = 1
	case 1:
		modelResponse.MondayBuyingCount = 1
	case 2:
		modelResponse.TuesdayBuyingCount = 1
	case 3:
		modelResponse.WednesdayBuyingCount = 1
	case 4:
		modelResponse.ThursdayBuyingCount = 1
	case 5:
		modelResponse.FridayBuyingCount = 1
	case 6:
		modelResponse.SaturdayBuyingCount = 1
	}
}

func determineBuyingHour(modelResponse *model.BuyingEventRespondModel, hour int64) {
	switch {
	case hour <= 5:
		modelResponse.Buying0To5HourCount = 1
	case (hour > 5) && (hour <= 11):
		modelResponse.Buying6To11HourCount = 1
	case (hour > 11) && (hour <= 17):
		modelResponse.Buying12To17HourCount = 1
	case (hour > 17) && (hour <= 23):
		modelResponse.Buying18To23HourCount = 1
	}
}

func determineBuyingAmPm(modelResponse *model.BuyingEventRespondModel, hour int64) {
	switch {
	case hour <= 12:
		modelResponse.AmBuyingCount = 1
	default:
		modelResponse.PmBuyingCount = 1
	}
}

func calculateBuyingLevelBasedAvgBuyingCount(modelResponse *model.BuyingEventRespondModel) float64 {
	if modelResponse.LevelIndex == 0 {
		modelResponse.LevelBasedAverageBuyingCount = float64(modelResponse.TotalBuyingCount)
		return modelResponse.LevelBasedAverageBuyingCount
	}
	modelResponse.LevelBasedAverageBuyingCount = float64(modelResponse.TotalBuyingCount) / float64(modelResponse.LevelIndex)
	return modelResponse.LevelBasedAverageBuyingCount
}
