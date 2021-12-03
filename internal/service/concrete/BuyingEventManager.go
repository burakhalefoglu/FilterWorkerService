package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IBuyingEventDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/logger"
)

type buyingEventManager struct {
	IBuyingEventDal *IBuyingEventDal.IBuyingEventDal
	IJsonParser     *IJsonParser.IJsonParser
	ICacheService   *ICacheService.ICacheService
	ILog            *logger.ILog
}

func BuyingEventManagerConstructor() *buyingEventManager {
	return &buyingEventManager{
		IBuyingEventDal: &IoC.BuyingEventDal,
		IJsonParser:     &IoC.JsonParser,
		ICacheService:   &IoC.CacheService,
		ILog:            &IoC.Logger,
	}
}

func (b *buyingEventManager) ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string) {
	firstModel := model.BuyingEventModel{}
	convertErr := (*b.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		(*b.ILog).SendErrorLog("BuyingEventManager", "ConvertRawModelToResponseModel",
			"byte array to BuyingEventModel", "Json Parser Decode Err: ", convertErr.Error())
		return &model.BuyingEventRespondModel{}, false, convertErr.Error()
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
	modelResponse.TotalBuyingHour = 0
	modelResponse.FirstBuyingYearOfDay = yearOfDay
	modelResponse.FirstBuyingYear = year
	modelResponse.FirstBuyingHour = hour
	modelResponse.FirstBuyingMinute = int64(firstModel.InWhatMinutes)
	modelResponse.FirstBuyingProductType, _, _ = (*b.ICacheService).ManageCache("ProductType", firstModel.ProductType)
	modelResponse.SecondBuyingYearOfDay = 0
	modelResponse.SecondBuyingHour = 0
	modelResponse.SecondBuyingMinute = 0
	modelResponse.SecondBuyingProductType = 0
	modelResponse.ThirdBuyingYearOfDay = 0
	modelResponse.ThirdBuyingHour = 0
	modelResponse.ThirdBuyingMinute = 0
	modelResponse.ThirdBuyingProductType = 0
	modelResponse.FourthBuyingYearOfDay = 0
	modelResponse.FourthBuyingHour = 0
	modelResponse.FourthBuyingMinute = 0
	modelResponse.FourthBuyingProductType = 0
	modelResponse.FifthBuyingYearOfDay = 0
	modelResponse.FifthBuyingHour = 0
	modelResponse.FifthBuyingMinute = 0
	modelResponse.FifthBuyingProductType = 0
	modelResponse.PenultimateBuyingYearOfDay = 0
	modelResponse.PenultimateBuyingHour = 0
	modelResponse.PenultimateBuyingMinute = 0
	modelResponse.PenultimateBuyingProductType = 0
	modelResponse.LastBuyingYearOfDay = 0
	modelResponse.LastBuyingYear = 0
	modelResponse.LastBuyingHour = 0
	modelResponse.LastBuyingMinute = 0
	modelResponse.LastBuyingProductType = 0
	modelResponse.FirstDayBuyingCount = 1
	modelResponse.SecondDayBuyingCount = 0
	modelResponse.ThirdDayBuyingCount = 0
	modelResponse.FourthDayBuyingCount = 0
	modelResponse.FifthDayBuyingCount = 0
	modelResponse.SixthDayBuyingCount = 0
	modelResponse.SeventhDayBuyingCount = 0
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
	if err != nil && err.Error() != "null data error" {
		(*b.ILog).SendErrorLog("BuyingEventManager", "ConvertRawModelToResponseModel",
			"BuyingEventDal_GetBuyingEventById", err.Error())
	}
	switch {
	case  err != nil && err.Error() == "null data error":

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
	oldModel.TotalBuyingHour = ((modelResponse.FirstBuyingYearOfDay+365*modelResponse.FirstBuyingYear)*24 + modelResponse.FirstBuyingHour) - ((oldModel.FirstBuyingYearOfDay+365*oldModel.FirstBuyingYear)*24 + oldModel.FirstBuyingHour)
	CalculateSecondBuying(modelResponse, oldModel)
	CalculateThirdBuying(modelResponse, oldModel)
	CalculateFourthBuying(modelResponse, oldModel)
	CalculateFifthBuying(modelResponse, oldModel)
	CalculateFirstDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateSecondDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateThirdDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateFourthDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateFifthDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateSixthDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateSeventhDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	oldModel.PenultimateDayBuyingCount = CalculatePenultimateDayBuyingCount(modelResponse, oldModel)
	oldModel.LastDayBuyingCount = CalculateLastDayBuyingCount(modelResponse, oldModel)

	oldModel.PenultimateBuyingYearOfDay = oldModel.LastBuyingYearOfDay
	oldModel.PenultimateBuyingHour = oldModel.LastBuyingHour
	oldModel.PenultimateBuyingMinute = oldModel.LastBuyingMinute
	oldModel.PenultimateBuyingProductType = oldModel.LastBuyingProductType
	oldModel.LastBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
	oldModel.LastBuyingYear = modelResponse.FirstBuyingYear
	oldModel.LastBuyingHour = modelResponse.FirstBuyingHour
	oldModel.LastBuyingMinute = modelResponse.FirstBuyingMinute
	oldModel.LastBuyingProductType = modelResponse.FirstBuyingProductType

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

func CalculateSecondBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) {
	switch oldModel.TotalBuyingCount {
	case 2:
		oldModel.SecondBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.SecondBuyingHour = modelResponse.FirstBuyingHour
		oldModel.SecondBuyingMinute = modelResponse.FirstBuyingMinute
		oldModel.SecondBuyingProductType = modelResponse.FirstBuyingProductType
	}
}

func CalculateThirdBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) {
	switch oldModel.TotalBuyingCount {
	case 3:
		oldModel.ThirdBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.ThirdBuyingHour = modelResponse.FirstBuyingHour
		oldModel.ThirdBuyingMinute = modelResponse.FirstBuyingMinute
		oldModel.ThirdBuyingProductType = modelResponse.FirstBuyingProductType
	}
}

func CalculateFourthBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) {
	switch oldModel.TotalBuyingCount {
	case 4:
		oldModel.FourthBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.FourthBuyingHour = modelResponse.FirstBuyingHour
		oldModel.FourthBuyingMinute = modelResponse.FourthBuyingMinute
		oldModel.FourthBuyingProductType = modelResponse.FourthBuyingProductType
	}
}

func CalculateFifthBuying(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) {
	switch oldModel.TotalBuyingCount {
	case 5:
		oldModel.FifthBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.FifthBuyingHour = modelResponse.FirstBuyingHour
		oldModel.FifthBuyingMinute = modelResponse.FourthBuyingMinute
		oldModel.FifthBuyingProductType = modelResponse.FourthBuyingProductType
	}
}

func CalculatePenultimateDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel) (count int64) {
	if ((modelResponse.FirstBuyingYearOfDay + (365 * modelResponse.FirstBuyingYear)) > (oldModel.LastBuyingYearOfDay + (365 * oldModel.LastBuyingYear))) && ((modelResponse.FirstBuyingYearOfDay + 365*modelResponse.FirstBuyingYear) != (oldModel.FirstBuyingYearOfDay + 365*oldModel.FirstBuyingYear)) {
		return oldModel.LastDayBuyingCount
	}
	return oldModel.PenultimateDayBuyingCount
}

func CalculateFirstDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel, total_adv_hour int64) {
	switch {
	case total_adv_hour <= 24:
		oldModel.FirstDayBuyingCount = oldModel.FirstDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateSecondDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel, total_adv_hour int64) {
	switch {
	case total_adv_hour <= 48 && total_adv_hour > 24:
		oldModel.SecondDayBuyingCount = oldModel.SecondDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateThirdDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel, total_adv_hour int64) {
	switch {
	case total_adv_hour <= 72 && total_adv_hour > 48:
		oldModel.ThirdDayBuyingCount = oldModel.ThirdDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateFourthDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel, total_adv_hour int64) {
	switch {
	case total_adv_hour <= 96 && total_adv_hour > 72:
		oldModel.FourthDayBuyingCount = oldModel.FourthDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateFifthDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel, total_adv_hour int64) {
	switch {
	case total_adv_hour <= 120 && total_adv_hour > 96:
		oldModel.FifthDayBuyingCount = oldModel.FifthDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateSixthDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel, total_adv_hour int64) {
	switch {
	case total_adv_hour <= 144 && total_adv_hour > 120:
		oldModel.SixthDayBuyingCount = oldModel.SixthDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateSeventhDayBuyingCount(modelResponse *model.BuyingEventRespondModel, oldModel *model.BuyingEventRespondModel, total_adv_hour int64) {
	switch {
	case total_adv_hour <= 168 && total_adv_hour > 144:
		oldModel.SeventhDayBuyingCount = oldModel.SeventhDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
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
