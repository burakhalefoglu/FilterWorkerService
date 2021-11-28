package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IBuyingEventDal "FilterWorkerService/internal/repository/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/logger"
)

type buyingEventManager struct {
	IBuyingEventDal *IBuyingEventDal.IBuyingEventDal
	IJsonParser     *IJsonParser.IJsonParser
	ILog            *logger.ILog
}

func BuyingEventManagerConstructor() *buyingEventManager {
	return &buyingEventManager{
		IBuyingEventDal: &IoC.BuyingEventDal,
		IJsonParser:     &IoC.JsonParser,
		ILog:            &IoC.Logger,
	}
}

func (b *buyingEventManager) ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string) {
	firstModel := model.BuyingEventModel{}
	Err := (*b.IJsonParser).DecodeJson(data, &firstModel)
	if Err != nil {
		(*b.ILog).SendErrorLog("BuyingEventManager", "ConvertRawModelToResponseModel",
			"byte array to BuyingEventModel", "Json Parser Decode Err: ", Err.Error())
		return &model.BuyingEventRespondModel{}, false, Err.Error()
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
	modelResponse.FirstBuyingMinute = int64(firstModel.InWhatMinutes)
	modelResponse.SecondBuyingYearOfDay = 0
	modelResponse.SecondBuyingHour = 0
	modelResponse.ThirdBuyingYearOfDay = 0
	modelResponse.ThirdBuyingHour = 0
	modelResponse.PenultimateBuyingYearOfDay = 0
	modelResponse.PenultimateBuyingHour = 0
	modelResponse.LastBuyingYearOfDay = 0
	modelResponse.LastBuyingYear = 0
	modelResponse.LastBuyingHour = 0
	modelResponse.FirstDayBuyingCount = 1
	modelResponse.PenultimateDayBuyingCount = 0
	modelResponse.LastDayBuyingCount = 0
	modelResponse.LastMinusFirstDayBuyingCount = -1
	DetermineBuyingDay(&modelResponse, day)
	DetermineBuyingHour(&modelResponse, hour)
	DetermineBuyingAmPm(&modelResponse, hour)
	modelResponse.BuyingDayAverageBuyingCount = 1
	CalculateBuyingLevelBasedAvgBuyingCount(&modelResponse)

	defer (*b.ILog).SendInfoLog("BuyingEventManager", "ConvertRawModelToResponseModel",
		modelResponse.ClientId, modelResponse.ProjectId)
	oldModel, err := (*b.IBuyingEventDal).GetBuyingEventById(modelResponse.ClientId)
	if err != nil {
		(*b.ILog).SendErrorLog("BuyingEventManager", "ConvertRawModelToResponseModel",
			"BuyingEventDal_GetBuyingEventById", err.Error())
	}
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := (*b.IBuyingEventDal).Add(&modelResponse)
		if logErr != nil {
			(*b.ILog).SendErrorLog("BuyingEventManager", "ConvertRawModelToResponseModel",
				"BuyingEventDal_Add", logErr.Error())
			return nil, false, logErr.Error()
		}
		return &modelResponse, true, "Added"

	case err == nil:
		updModel, updateResult, updateErr := b.UpdateBuyingEvent(&modelResponse, oldModel)
		if updateErr != nil {
			return nil, updateResult, updateErr.Error()
		}
		return updModel, updateResult, "Updated"

	default:

		return nil, false, err.Error()

	}

}

func (b *buyingEventManager) UpdateBuyingEvent(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (updatedModel *model.BuyingEventRespondModel, s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalBuyingCount = oldModel.TotalBuyingCount + modelResponse.TotalBuyingCount
	oldModel.TotalBuyingDay = (modelResponse.FirstBuyingYearOfDay - oldModel.FirstBuyingYearOfDay) + 365*(modelResponse.FirstBuyingYear-oldModel.FirstBuyingYear)
	//oldModel.FirstBuyingYearOfDay
	//oldModel.FirstBuyingYear
	//oldModel.FirstBuyingHour
	//oldModel.FirstBuyingMinute
	oldModel.SecondBuyingYearOfDay, oldModel.SecondBuyingHour = CalculateSecondBuying(modelResponse, oldModel)
	oldModel.ThirdBuyingYearOfDay, modelResponse.ThirdBuyingHour = CalculateThirdBuying(modelResponse, oldModel)

	oldModel.FirstDayBuyingCount = CalculateFirstDayBuyingCount(modelResponse, oldModel)
	oldModel.PenultimateDayBuyingCount = CalculatePenultimateDayBuyingCount(modelResponse, oldModel)
	oldModel.LastDayBuyingCount = CalculateLastDayBuyingCount(modelResponse, oldModel)

	oldModel.PenultimateBuyingYearOfDay = oldModel.LastBuyingYearOfDay
	oldModel.PenultimateBuyingHour = oldModel.LastBuyingHour
	oldModel.LastBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
	oldModel.LastBuyingYear = modelResponse.FirstBuyingYear
	oldModel.LastBuyingHour = modelResponse.FirstBuyingHour
	oldModel.LastBuyingMinute = modelResponse.FirstBuyingMinute

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
	CalculateBuyingLevelBasedAvgBuyingCount(oldModel)

	defer (*b.ILog).SendInfoLog("BuyingEventManager", "UpdateBuyingEvent",
		oldModel.ClientId, oldModel.ProjectId)
	logErr := (*b.IBuyingEventDal).UpdateBuyingEventById(oldModel.ClientId, oldModel)
	if logErr != nil {
		(*b.ILog).SendErrorLog("BuyingEventManager", "UpdateBuyingEvent",
			"BuyingEventDal_UpdateBuyingEventById", logErr.Error())
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateSecondBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (day int64, hour int64) {
	if oldModel.TotalBuyingCount == 2 {
		oldModel.SecondBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.SecondBuyingHour = modelResponse.FirstBuyingHour
		return oldModel.SecondBuyingYearOfDay, oldModel.SecondBuyingHour
	}
	return oldModel.SecondBuyingYearOfDay, oldModel.SecondBuyingHour
}

func CalculateThirdBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (day int64, hour int64) {
	if oldModel.TotalBuyingCount == 3 {
		oldModel.ThirdBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.ThirdBuyingHour = modelResponse.FirstBuyingHour
		return oldModel.ThirdBuyingYearOfDay, oldModel.ThirdBuyingHour
	}
	return oldModel.ThirdBuyingYearOfDay, oldModel.ThirdBuyingHour
}

func CalculatePenultimateDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (count int64) {
	if ((modelResponse.FirstBuyingYearOfDay + (365 * modelResponse.FirstBuyingYear)) > (oldModel.LastBuyingYearOfDay + (365 * oldModel.LastBuyingYear))) && ((modelResponse.FirstBuyingYearOfDay + 365*modelResponse.FirstBuyingYear) != (oldModel.FirstBuyingYearOfDay + 365*oldModel.FirstBuyingYear)) {
		return oldModel.LastDayBuyingCount
	}
	return oldModel.PenultimateDayBuyingCount
}

func CalculateFirstDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) int64 {
	if ((oldModel.FirstBuyingYearOfDay) == (modelResponse.FirstBuyingYearOfDay)) && ((oldModel.FirstBuyingYear) == (modelResponse.FirstBuyingYear)) {
		oldModel.FirstDayBuyingCount = oldModel.FirstDayBuyingCount + modelResponse.FirstDayBuyingCount
		return oldModel.FirstDayBuyingCount
	}
	return oldModel.FirstDayBuyingCount
}

func CalculateLastDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) int64 {
	switch {
	case ((oldModel.LastBuyingYearOfDay + (365 * oldModel.LastBuyingYear)) == (modelResponse.FirstBuyingYearOfDay + (365 * modelResponse.FirstBuyingYear))) && ((modelResponse.FirstBuyingYearOfDay + 365*modelResponse.FirstBuyingYear) != (oldModel.FirstBuyingYearOfDay + 365*oldModel.FirstBuyingYear)):
		oldModel.LastDayBuyingCount = oldModel.LastDayBuyingCount + modelResponse.FirstDayBuyingCount
		return oldModel.LastDayBuyingCount
	case ((modelResponse.FirstBuyingYearOfDay + (365 * modelResponse.FirstBuyingYear)) > (oldModel.LastBuyingYearOfDay + (365 * oldModel.LastBuyingYear))) && ((modelResponse.FirstBuyingYearOfDay + 365*modelResponse.FirstBuyingYear) != (oldModel.FirstBuyingYearOfDay + 365*oldModel.FirstBuyingYear)):
		return modelResponse.FirstDayBuyingCount
	default:
		return oldModel.LastDayBuyingCount

	}
}

func DetermineBuyingDay(modelResponse *model.BuyingEventRespondModel, day int64) {
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

func DetermineBuyingHour(modelResponse *model.BuyingEventRespondModel, hour int64) {
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

func DetermineBuyingAmPm(modelResponse *model.BuyingEventRespondModel, hour int64) {
	switch {
	case hour <= 12:
		modelResponse.AmBuyingCount = 1
	default:
		modelResponse.PmBuyingCount = 1
	}
}

func CalculateBuyingLevelBasedAvgBuyingCount(modelResponse *model.BuyingEventRespondModel) {
	switch modelResponse.LevelIndex {
	case 0:
		modelResponse.LevelBasedAverageBuyingCount = float64(modelResponse.TotalBuyingCount)
	default:
		modelResponse.LevelBasedAverageBuyingCount = float64(modelResponse.TotalBuyingCount) / float64(modelResponse.LevelIndex)

	}
}
