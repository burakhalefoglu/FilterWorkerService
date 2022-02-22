package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IBuyingEventDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"fmt"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
)

type buyingEventManager struct {
	IBuyingEventDal *IBuyingEventDal.IBuyingEventDal
	IJsonParser     *IJsonParser.IJsonParser
	ICacheService   *ICacheService.ICacheService
}

func BuyingEventManagerConstructor() *buyingEventManager {
	return &buyingEventManager{
		IBuyingEventDal: &IoC.BuyingEventDal,
		IJsonParser:     &IoC.JsonParser,
		ICacheService:   &IoC.CacheService,
	}
}

func (b *buyingEventManager) ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string) {
	firstModel := model.BuyingEventModel{}
	convertErr := (*b.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		clogger.Error(&logger.Messages{"Byte array to BuyingEventModel  BuyingEventManager Json Parser Decode ERROR: ": convertErr.Error()})
		// log.Fatal("BuyingEventManager", "ConvertRawModelToResponseModel",
		// 	"byte array to BuyingEventModel", "Json Parser Decode Err: ", convertErr.Error())
		return &model.BuyingEventResponseModel{}, false, convertErr.Error()
	}
	hour := int16(firstModel.TriggeredTime.Hour())
	day := int16(firstModel.TriggeredTime.Weekday())
	yearOfDay := int16(firstModel.TriggeredTime.YearDay())
	year := int16(firstModel.TriggeredTime.Year())
	var value, _, _ = (*b.ICacheService).ManageCache("ProductType", firstModel.ProductType)
	var ProductType = byte(value)
	modelResponse := model.BuyingEventResponseModel{}
	modelResponse.Id = firstModel.Id
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int16(firstModel.LevelIndex)
	modelResponse.TotalBuyingCount = 1
	modelResponse.TotalBuyingDay = 1
	modelResponse.TotalBuyingHour = 0
	modelResponse.FirstBuyingYearOfDay = yearOfDay
	modelResponse.FirstBuyingYear = year
	modelResponse.FirstBuyingHour = hour
	modelResponse.FirstBuyingMinute = int16(firstModel.InMinutes)
	modelResponse.FirstBuyingProductType = ProductType
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
	modelResponse.SixthBuyingYearOfDay = 0
	modelResponse.SixthBuyingHour = 0
	modelResponse.SixthBuyingMinute = 0
	modelResponse.SixthBuyingProductType = 0
	modelResponse.SeventhBuyingYearOfDay = 0
	modelResponse.SeventhBuyingHour = 0
	modelResponse.SeventhBuyingMinute = 0
	modelResponse.SeventhBuyingProductType = 0
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
	DetermineBuyingDay(&modelResponse, day)
	DetermineBuyingHour(&modelResponse, hour)
	DetermineBuyingAmPm(&modelResponse, hour)
	modelResponse.BuyingDayAverageBuyingCount = 1
	CalculateBuyingLevelBasedAvgBuyingCount(&modelResponse)

	// defer log.Print("BuyingEventManager", "ConvertRawModelToResponseModel",
	// 	modelResponse.ClientId, modelResponse.ProjectId)
	oldModel, err := (*b.IBuyingEventDal).GetById(modelResponse.ClientId, modelResponse.ProjectId)
	// if err != nil && err.Error() != "null data error" {
	// 	log.Fatal("BuyingEventManager", "ConvertRawModelToResponseModel",
	// 		"BuyingEventDal_GetBuyingEventById", err.Error())
	// }
	switch {

	case err != nil && err.Error() != "not found":
		clogger.Error(&logger.Messages{
			fmt.Sprintf("Get clientId: %d, projectId: %d buying_event_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): err.Error(),
		})

	case err != nil && err.Error() == "not found":

		logErr := (*b.IBuyingEventDal).Add(&modelResponse)
		if logErr != nil {
			clogger.Error(&logger.Messages{
				fmt.Sprintf("Add clientId: %d, projectId: %d buying_event_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): logErr.Error(),
			})

			return nil, false, logErr.Error()
		}
		clogger.Info(&logger.Messages{
			fmt.Sprintf("Add clientId: %d, projectId: %d buying_event_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return &modelResponse, true, "Added"

	case err == nil:
		updModel, updateResult, updateErr := b.UpdateBuyingEvent(&modelResponse, oldModel)
		if updateErr != nil {

			clogger.Error(&logger.Messages{
				fmt.Sprintf("Update clientId: %d, projectId: %d buying_event_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): updateErr.Error(),
			})
			return nil, updateResult, updateErr.Error()
		}
		clogger.Info(&logger.Messages{
			fmt.Sprintf("Update clientId: %d, projectId: %d buying_event_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return updModel, updateResult, "Updated"

	default:

		return nil, false, ""

	}
	return nil, false, ""

}

func (b *buyingEventManager) UpdateBuyingEvent(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel) (updatedModel *model.BuyingEventResponseModel, s bool, m error) {
	oldModel.Id = modelResponse.Id
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalBuyingCount = oldModel.TotalBuyingCount + modelResponse.TotalBuyingCount
	oldModel.TotalBuyingDay = (int32(modelResponse.FirstBuyingYearOfDay) - int32(oldModel.FirstBuyingYearOfDay)) + 365*(int32(modelResponse.FirstBuyingYear)-int32(oldModel.FirstBuyingYear))
	oldModel.TotalBuyingHour = ((int32(modelResponse.FirstBuyingYearOfDay)+365*int32(modelResponse.FirstBuyingYear))*24 + int32(modelResponse.FirstBuyingHour)) - ((int32(oldModel.FirstBuyingYearOfDay)+365*int32(oldModel.FirstBuyingYear))*24 + int32(oldModel.FirstBuyingHour))
	CalculateSecondBuying(modelResponse, oldModel)
	CalculateThirdBuying(modelResponse, oldModel)
	CalculateFourthBuying(modelResponse, oldModel)
	CalculateFifthBuying(modelResponse, oldModel)
	CalculateSixthBuying(modelResponse, oldModel)
	CalculateSeventhBuying(modelResponse, oldModel)
	CalculateFirstDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateSecondDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateThirdDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateFourthDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateFifthDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateSixthDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	CalculateSeventhDayBuyingCount(modelResponse, oldModel, oldModel.TotalBuyingHour)
	// oldModel.PenultimateDayBuyingCount = CalculatePenultimateDayBuyingCount(modelResponse, oldModel)
	// oldModel.LastDayBuyingCount = CalculateLastDayBuyingCount(modelResponse, oldModel)

	oldModel.PenultimateBuyingYearOfDay = oldModel.LastBuyingYearOfDay
	oldModel.PenultimateBuyingHour = oldModel.LastBuyingHour
	oldModel.PenultimateBuyingMinute = oldModel.LastBuyingMinute
	oldModel.PenultimateBuyingProductType = oldModel.LastBuyingProductType
	oldModel.LastBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
	oldModel.LastBuyingYear = modelResponse.FirstBuyingYear
	oldModel.LastBuyingHour = modelResponse.FirstBuyingHour
	oldModel.LastBuyingMinute = modelResponse.FirstBuyingMinute
	oldModel.LastBuyingProductType = modelResponse.FirstBuyingProductType

	//oldModel.LastMinusFirstDayBuyingCount = oldModel.LastDayBuyingCount - oldModel.FirstDayBuyingCount
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
	oldModel.BuyingDayAverageBuyingCount = float32(oldModel.TotalBuyingCount) / float32(oldModel.TotalBuyingDay)
	CalculateBuyingLevelBasedAvgBuyingCount(oldModel)

	// defer log.Print("BuyingEventManager", "UpdateBuyingEvent",
	// 	oldModel.ClientId, oldModel.ProjectId)
	logErr := (*b.IBuyingEventDal).UpdateById(oldModel.ClientId, oldModel.ProjectId, oldModel)
	if logErr != nil {
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateSecondBuying(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel) {
	switch oldModel.TotalBuyingCount {
	case 2:
		oldModel.SecondBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.SecondBuyingHour = modelResponse.FirstBuyingHour
		oldModel.SecondBuyingMinute = modelResponse.FirstBuyingMinute
		oldModel.SecondBuyingProductType = modelResponse.FirstBuyingProductType
	}
}

func CalculateThirdBuying(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel) {
	switch oldModel.TotalBuyingCount {
	case 3:
		oldModel.ThirdBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.ThirdBuyingHour = modelResponse.FirstBuyingHour
		oldModel.ThirdBuyingMinute = modelResponse.FirstBuyingMinute
		oldModel.ThirdBuyingProductType = modelResponse.FirstBuyingProductType
	}
}

func CalculateFourthBuying(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel) {
	switch oldModel.TotalBuyingCount {
	case 4:
		oldModel.FourthBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.FourthBuyingHour = modelResponse.FirstBuyingHour
		oldModel.FourthBuyingMinute = modelResponse.FourthBuyingMinute
		oldModel.FourthBuyingProductType = modelResponse.FourthBuyingProductType
	}
}

func CalculateFifthBuying(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel) {
	switch oldModel.TotalBuyingCount {
	case 5:
		oldModel.FifthBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.FifthBuyingHour = modelResponse.FirstBuyingHour
		oldModel.FifthBuyingMinute = modelResponse.FourthBuyingMinute
		oldModel.FifthBuyingProductType = modelResponse.FourthBuyingProductType
	}
}

func CalculateSixthBuying(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel) {
	switch oldModel.TotalBuyingCount {
	case 6:
		oldModel.SixthBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.SixthBuyingHour = modelResponse.FirstBuyingHour
		oldModel.SixthBuyingMinute = modelResponse.FourthBuyingMinute
		oldModel.SixthBuyingProductType = modelResponse.FourthBuyingProductType
	}
}

func CalculateSeventhBuying(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel) {
	switch oldModel.TotalBuyingCount {
	case 7:
		oldModel.SeventhBuyingYearOfDay = modelResponse.FirstBuyingYearOfDay
		oldModel.SeventhBuyingHour = modelResponse.FirstBuyingHour
		oldModel.SeventhBuyingMinute = modelResponse.FourthBuyingMinute
		oldModel.SeventhBuyingProductType = modelResponse.FourthBuyingProductType
	}
}

func CalculateFirstDayBuyingCount(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel, total_adv_hour int32) {
	switch {
	case total_adv_hour <= 24:
		oldModel.FirstDayBuyingCount = oldModel.FirstDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateSecondDayBuyingCount(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel, total_adv_hour int32) {
	switch {
	case total_adv_hour <= 48 && total_adv_hour > 24:
		oldModel.SecondDayBuyingCount = oldModel.SecondDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateThirdDayBuyingCount(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel, total_adv_hour int32) {
	switch {
	case total_adv_hour <= 72 && total_adv_hour > 48:
		oldModel.ThirdDayBuyingCount = oldModel.ThirdDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateFourthDayBuyingCount(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel, total_adv_hour int32) {
	switch {
	case total_adv_hour <= 96 && total_adv_hour > 72:
		oldModel.FourthDayBuyingCount = oldModel.FourthDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateFifthDayBuyingCount(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel, total_adv_hour int32) {
	switch {
	case total_adv_hour <= 120 && total_adv_hour > 96:
		oldModel.FifthDayBuyingCount = oldModel.FifthDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateSixthDayBuyingCount(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel, total_adv_hour int32) {
	switch {
	case total_adv_hour <= 144 && total_adv_hour > 120:
		oldModel.SixthDayBuyingCount = oldModel.SixthDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func CalculateSeventhDayBuyingCount(modelResponse *model.BuyingEventResponseModel, oldModel *model.BuyingEventResponseModel, total_adv_hour int32) {
	switch {
	case total_adv_hour <= 168 && total_adv_hour > 144:
		oldModel.SeventhDayBuyingCount = oldModel.SeventhDayBuyingCount + modelResponse.FirstDayBuyingCount
	}
}

func DetermineBuyingDay(modelResponse *model.BuyingEventResponseModel, day int16) {
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

func DetermineBuyingHour(modelResponse *model.BuyingEventResponseModel, hour int16) {
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

func DetermineBuyingAmPm(modelResponse *model.BuyingEventResponseModel, hour int16) {
	switch {
	case hour <= 12:
		modelResponse.AmBuyingCount = 1
	default:
		modelResponse.PmBuyingCount = 1
	}
}

func CalculateBuyingLevelBasedAvgBuyingCount(modelResponse *model.BuyingEventResponseModel) {
	switch modelResponse.LevelIndex {
	case 0:
		modelResponse.LevelBasedAverageBuyingCount = float32(modelResponse.TotalBuyingCount)
	default:
		modelResponse.LevelBasedAverageBuyingCount = float32(modelResponse.TotalBuyingCount) / float32(modelResponse.LevelIndex)
	}
}
