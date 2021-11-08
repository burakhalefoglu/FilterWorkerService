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



func (b *BuyingEventManager) ConvertRawModelToResponseModel(data *[]byte) (respondModel *model.BuyingEventRespondModel, s bool, m string){
	firstModel := model.BuyingEventModel{}
	err := b.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return &model.BuyingEventRespondModel{}, false, err.Error()
	}
	hour := int64(firstModel.TrigerdTime.Hour())
	day := int64(firstModel.TrigerdTime.Weekday())
	yearOfDay := int64(firstModel.TrigerdTime.YearDay())
	
	modelResponse := model.BuyingEventRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = 	int64(firstModel.LevelIndex)
	modelResponse.TotalBuyingCount = 1
	modelResponse.TotalBuyingDay = 1
	modelResponse.FirstBuyingYearOfDay = yearOfDay
	modelResponse.FirstBuyingHour = hour
	modelResponse.LastBuyingYearOfDay = yearOfDay
	modelResponse.LastBuyingHour = hour
	modelResponse.FirstDayBuyingCount = 1
	modelResponse.LastDayBuyingCount = 1
	modelResponse.LastMinusFirstDayBuyingCount = modelResponse.LastDayBuyingCount - modelResponse.FirstDayBuyingCount
	determineBuyingDay(&modelResponse, day)
	determineBuyingHour(&modelResponse, hour)
	determineBuyingAmPm(&modelResponse, hour)
	modelResponse.BuyingDayAverageBuyingCount = float64(modelResponse.TotalBuyingCount) / float64(modelResponse.TotalBuyingDay)
	modelResponse.LevelBasedAverageBuyingCount = calculateBuyingLevelBasedAvgBuyingCount(&modelResponse)
	return &modelResponse, true, ""
}

func (b *BuyingEventManager) AddBuyingEvent(data *model.BuyingEventRespondModel) (s bool, m string) {
	logErr := b.IBuyingEventDal.Add(data)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func (b *BuyingEventManager) UpdateBuyingEventByCustomerId(modelResponse *model.BuyingEventRespondModel) (s bool, m string){
	oldModel, err := b.IBuyingEventDal.GetBuyingEventByCustomerId(modelResponse.CustomerId)
	if err != nil {
		return false, err.Error()
	}
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalBuyingCount = oldModel.TotalBuyingCount + modelResponse.TotalBuyingCount
	oldModel.TotalBuyingDay = calculateTotalBuyingDay(modelResponse, oldModel)
	oldModel.LastBuyingYearOfDay = modelResponse.LastBuyingYearOfDay
	oldModel.LastBuyingHour = modelResponse.LastBuyingHour
	oldModel.FirstDayBuyingCount = calculateFirstDayBuyingCount(modelResponse, oldModel)
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
	oldModel.BuyingDayAverageBuyingCount = float64(oldModel.TotalBuyingCount)/ float64(oldModel.TotalBuyingDay)
	oldModel.LevelBasedAverageBuyingCount = calculateBuyingLevelBasedAvgBuyingCount(oldModel)
	logErr := b.IBuyingEventDal.UpdateBuyingEventByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func calculateFirstDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) int64 {
	if (oldModel.FirstBuyingYearOfDay == modelResponse.FirstBuyingYearOfDay){
		oldModel.FirstDayBuyingCount = oldModel.FirstDayBuyingCount + modelResponse.FirstDayBuyingCount
		return oldModel.FirstDayBuyingCount
	}
	return oldModel.FirstDayBuyingCount
}

func calculateLastDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) int64 {
	if (oldModel.LastBuyingYearOfDay == modelResponse.LastBuyingYearOfDay) && (oldModel.FirstBuyingYearOfDay != modelResponse.FirstBuyingYearOfDay){
		oldModel.LastDayBuyingCount = oldModel.LastDayBuyingCount + modelResponse.LastDayBuyingCount
		return oldModel.LastDayBuyingCount
	}	
	return oldModel.LastDayBuyingCount
}

func calculateTotalBuyingDay(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) int64 {
	if modelResponse.LastBuyingYearOfDay != oldModel.LastBuyingYearOfDay {
		oldModel.TotalBuyingDay = oldModel.TotalBuyingDay + modelResponse.TotalBuyingDay
		return oldModel.TotalBuyingDay
	}
	return oldModel.TotalBuyingDay
}

func determineBuyingDay(modelResponse *model.BuyingEventRespondModel, day int64) {
	if day == 0 {
		modelResponse.SundayBuyingCount = 1
	} else if day == 1 {
		modelResponse.MondayBuyingCount = 1
	} else if day == 2 {
		modelResponse.TuesdayBuyingCount = 1
	} else if day == 3 {
		modelResponse.WednesdayBuyingCount = 1
	} else if day == 4 {
		modelResponse.ThursdayBuyingCount = 1
	} else if day == 5 {
		modelResponse.FridayBuyingCount = 1
	} else if day == 6 {
		modelResponse.SaturdayBuyingCount = 1
	}
}

func determineBuyingHour(modelResponse *model.BuyingEventRespondModel, hour int64) {
	if hour <= 5 {
		modelResponse.Buying0To5HourCount = 1
	} else if (hour > 5) && (hour <= 11) {
		modelResponse.Buying6To11HourCount = 1
	} else if (hour > 11) && (hour <= 17) {
		modelResponse.Buying12To17HourCount = 1
	} else if (hour > 17) && (hour <= 23) {
		modelResponse.Buying18To23HourCount = 1
	}
}

func determineBuyingAmPm(modelResponse *model.BuyingEventRespondModel, hour int64) {
	if hour <= 12 {
		modelResponse.AmBuyingCount = 1
	} 
	modelResponse.PmBuyingCount = 1
}

func calculateBuyingLevelBasedAvgBuyingCount(modelResponse *model.BuyingEventRespondModel) float64 {
	if modelResponse.LevelIndex == 0 {
		modelResponse.LevelBasedAverageBuyingCount = float64(modelResponse.TotalBuyingCount)
		return modelResponse.LevelBasedAverageBuyingCount
	}
	modelResponse.LevelBasedAverageBuyingCount = float64(modelResponse.TotalBuyingCount) / float64(modelResponse.LevelIndex)
	return modelResponse.LevelBasedAverageBuyingCount
}
