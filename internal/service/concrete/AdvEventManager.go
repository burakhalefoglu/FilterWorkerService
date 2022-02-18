package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IAdvEventDal "FilterWorkerService/internal/repository/abstract"
	ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"fmt"
	"log"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
)

type advEventManager struct {
	IAdvEventDal  *IAdvEventDal.IAdvEventDal
	IJsonParser   *IJsonParser.IJsonParser
	ICacheService *ICacheService.ICacheService
}

func AdvEventManagerConstructor() *advEventManager {
	return &advEventManager{IAdvEventDal: &IoC.AdvEventDal,
		IJsonParser:   &IoC.JsonParser,
		ICacheService: &IoC.CacheService,
	}
}

func (a *advEventManager) ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string) {
	firstModel := model.AdvEventDataModel{}
	convertErr := (*a.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		clogger.Error(&logger.Messages{
			"AdvEventManager ConvertRawModelToResponseModel Json Parser Decode Err: ": convertErr.Error(),
		})
		// log.Fatal("AdvEventManager", "ConvertRawModelToResponseModel",
		// 	"byte array to AdvEventModel", "Json Parser Decode Err: ", convertErr.Error())
		return nil, false, convertErr.Error()
	}
	hour := int16(firstModel.TriggeredTime.Hour())
	day := int16(firstModel.TriggeredTime.Weekday())
	yearOfDay := int16(firstModel.TriggeredTime.YearDay())
	year := int16(firstModel.TriggeredTime.Year())
	minute := int16(firstModel.InMinutes)
	value, _, _ := (*a.ICacheService).ManageCache("AdvType", firstModel.AdvType)
	adv_type := byte(value)
	modelResponse := model.AdvEventResponseModel{}
	modelResponse.Id = firstModel.Id
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int16(firstModel.LevelIndex)
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
	modelResponse.FirstAdvClickMinute = minute
	modelResponse.FirstAdvType = adv_type
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
	modelResponse.SixthAdvYearOfDay   = 0
	modelResponse.SixthAdvHour        = 0
	modelResponse.SixthAdvMinute      = 0
	modelResponse.SixthAdvType        = 0
	modelResponse.SeventhAdvYearOfDay = 0
	modelResponse.SeventhAdvHour      = 0
	modelResponse.SeventhAdvMinute    = 0
	modelResponse.SeventhAdvType      = 0
	modelResponse.PenultimateAdvYearOfDay = 0
	modelResponse.PenultimateAdvHour = 0
	modelResponse.PenultimateAdvMinute = 0
	modelResponse.PenultimateAdvType = 0
	modelResponse.LastAdvYearOfDay = 0
	modelResponse.LastAdvYear = 0
	modelResponse.LastAdvClickHour = 0
	modelResponse.LastAdvClickMinute = 0
	modelResponse.LastAdvType = 0

	modelResponse.FirstFiveMinutesAdvClickCount  = 1
	modelResponse.FirstTenMinutesAdvClickCount  = 1
	modelResponse.FirstQuarterHourAdvClickCount = 1
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

	// defer log.Print("AdvEventManager", "ConvertRawModelToResponseModel",
	// 	modelResponse.ClientId, modelResponse.ProjectId)

	oldModel, err := (*a.IAdvEventDal).GetById(modelResponse.ClientId, modelResponse.ProjectId)
	if err != nil && err.Error() != "null data error" {
		log.Fatal("AdvEventManager", "ConvertRawModelToResponseModel",
			"AdvEventDal_GetAdvEventById", err.Error())
	}
	switch {
	case err != nil && err.Error() == "null data error":

		logErr := (*a.IAdvEventDal).Add(&modelResponse)
		if logErr != nil {
			clogger.Error(&logger.Messages{
			"Add adv_event_data err: ": logErr.Error(),
		})
			// log.Fatal("AdvEventManager", "ConvertRawModelToResponseModel",
			// 	"AdvEventDal_Add", logErr.Error())
			return nil, false, logErr.Error()
		}
		clogger.Info(&logger.Messages{
			"Add adv_event_data  : ": "SUCCESS",
		})
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

func (a *advEventManager) UpdateAdvEvent(modelResponse *model.AdvEventResponseModel,
	oldModel *model.AdvEventResponseModel) (updatedModel *model.AdvEventResponseModel, s bool, m error) {
	oldModel.Id = modelResponse.Id
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalAdvDay = (int32(modelResponse.FirstAdvYearOfDay) - int32(oldModel.FirstAdvYearOfDay)) + 365*(int32(modelResponse.FirstAdvYear)-int32(oldModel.FirstAdvYear))
	oldModel.TotalAdvCount = oldModel.TotalAdvCount + modelResponse.TotalAdvCount
	oldModel.TotalAdvHour = ((int32(modelResponse.FirstAdvYearOfDay)+365*int32(modelResponse.FirstAdvYear))*24 + int32(modelResponse.FirstAdvClickHour)) - ((int32(oldModel.FirstAdvYearOfDay)+365*int32(oldModel.FirstAdvYear))*24 + int32(oldModel.FirstAdvClickHour))
	oldModel.TotalAdvMinute = (((int32(modelResponse.FirstAdvYearOfDay)+365*int32(modelResponse.FirstAdvYear))*24+int32(modelResponse.FirstAdvClickHour))*60 + int32(modelResponse.FirstAdvClickMinute)) - (((int32(oldModel.FirstAdvYearOfDay)+365*int32(oldModel.FirstAdvYear))*24+int32(oldModel.FirstAdvClickHour))*60 + int32(oldModel.FirstAdvClickMinute))
	CalculateAdvLevelBasedAvgClickCount(oldModel)
	oldModel.AverageAdvDailyClickCount = CalculateAverageAdvDailyClickCount(oldModel)

	CalculateSecondAdv(modelResponse, oldModel)
	CalculateThirdAdv(modelResponse, oldModel)
	CalculateFourthAdv(modelResponse, oldModel)
	CalculateFifthAdv(modelResponse, oldModel)
	CalculateSixthAdv(modelResponse, oldModel)
	CalculateSeventhAdv(modelResponse, oldModel)

	oldModel.PenultimateDayAdvClickCount = CalculatePenultimateDayAdvDay(modelResponse, oldModel)
	oldModel.LastDayAdvClickCount = CalculateLastDayAdvClickCount(modelResponse, oldModel)

	oldModel.PenultimateAdvYearOfDay = oldModel.LastAdvYearOfDay
	oldModel.PenultimateAdvHour = oldModel.LastAdvClickHour
	oldModel.PenultimateAdvMinute = oldModel.LastAdvClickMinute
	oldModel.LastAdvYearOfDay = modelResponse.FirstAdvYearOfDay
	oldModel.LastAdvYear = modelResponse.FirstAdvYear
	oldModel.LastAdvClickHour = modelResponse.FirstAdvClickHour
	oldModel.LastAdvClickMinute = modelResponse.FirstAdvClickMinute
	oldModel.LastAdvType = modelResponse.FirstAdvType
	
	CalculateFirstFiveMinutesTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateFirstTenMinutesTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateFirstQuarterHourTotalAdvCount(modelResponse, oldModel,oldModel.TotalAdvMinute)
	CalculateFirstHalfHourTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateFirstHourTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateFirstTwoHourTotalAdvCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateFirstThreeHourAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateFirstSixHourAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateFirstTwelveHourAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)

	CalculateFirstDayAdvClickCount(modelResponse, oldModel, oldModel.TotalAdvMinute)
	CalculateSecondDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	CalculateThirdDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	CalculateFourthDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	CalculateFifthDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	CalculateSixthDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)
	CalculateSeventhDayTotalSessionCount(modelResponse, oldModel, oldModel.TotalAdvHour)

	oldModel.LastMinusFirstDayAdvClickCount = oldModel.LastDayAdvClickCount - oldModel.FirstDayAdvClickCount
	oldModel.LastMinusPenultimateDayAdvClickCount = oldModel.LastDayAdvClickCount - oldModel.PenultimateDayAdvClickCount
	oldModel.LastDayAdvClickCountMinusAverageDailyAdvClickCount = float32(oldModel.LastDayAdvClickCount) - oldModel.AverageAdvDailyClickCount

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

	// defer log.Print("AdvEventManager", "UpdateAdvEvent",
	// 	oldModel.ClientId, oldModel.ProjectId)
		
	logErr := (*a.IAdvEventDal).UpdateById(oldModel.ClientId, oldModel.ProjectId, oldModel)
	if logErr != nil {

		clogger.Error(&logger.Messages{
			"Update adv_event_data err: ": logErr.Error(),
		})
		// log.Fatal("AdvEventManager", "UpdateAdvEvent",
		// 	"AdvEventDal_UpdateAdvEventById", logErr.Error())
		return oldModel, false, logErr
	}
	clogger.Info(&logger.Messages{
		fmt.Sprintf("Update %d %d adv_event_data  : ",oldModel.ClientId, oldModel.ProjectId): "SUCCESS",
	})
	return oldModel, true, nil
}

func CalculateAverageAdvDailyClickCount(oldModel *model.AdvEventResponseModel) (count float32) {
	if oldModel.TotalAdvDay == 0 {
		return float32(oldModel.TotalAdvCount)
	}
	return float32(oldModel.TotalAdvCount) / float32(oldModel.TotalAdvDay)
}

func CalculateFirstFiveMinutesTotalAdvCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 5 {
		oldModel.FirstFiveMinutesAdvClickCount = oldModel.FirstFiveMinutesAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstTenMinutesTotalAdvCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 10 {
		oldModel.FirstTenMinutesAdvClickCount = oldModel.FirstTenMinutesAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstQuarterHourTotalAdvCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 15 {
		oldModel.FirstQuarterHourAdvClickCount = oldModel.FirstQuarterHourAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstHalfHourTotalAdvCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 30 {
		oldModel.FirstHalfHourAdvClickCount = oldModel.FirstHalfHourAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstHourTotalAdvCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 60 {
		oldModel.FirstHourAdvClickCount = oldModel.FirstHourAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstTwoHourTotalAdvCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 120 {
		oldModel.FirstTwoHourAdvClickCount = oldModel.FirstTwoHourAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstThreeHourAdvClickCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 180 {
		oldModel.FirstThreeHourAdvClickCount = oldModel.FirstThreeHourAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstSixHourAdvClickCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 360 {
		oldModel.FirstSixHourAdvClickCount = oldModel.FirstSixHourAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFirstTwelveHourAdvClickCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 720 {
		oldModel.FirstTwelveHourAdvClickCount = oldModel.FirstTwelveHourAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateSecondDayTotalSessionCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_hour int32) {
	if total_adv_hour <= 48 && total_adv_hour > 24 {
		oldModel.SecondDayAdvClickCount = oldModel.SecondDayAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateThirdDayTotalSessionCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_hour int32) {
	if total_adv_hour <= 72 && total_adv_hour > 48 {
		oldModel.ThirdDayAdvClickCount = oldModel.ThirdDayAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFourthDayTotalSessionCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_hour int32) {
	if total_adv_hour <= 96 && total_adv_hour > 72 {
		oldModel.FourthDayAdvClickCount = oldModel.FourthDayAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateFifthDayTotalSessionCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_hour int32) {
	if total_adv_hour <= 120 && total_adv_hour > 96 {
		oldModel.FifthDayAdvClickCount = oldModel.FifthDayAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateSixthDayTotalSessionCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_hour int32) {
	if total_adv_hour <= 144 && total_adv_hour > 120 {
		oldModel.SixthDayAdvClickCount = oldModel.SixthDayAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateSeventhDayTotalSessionCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_hour int32) {
	if total_adv_hour <= 168 && total_adv_hour > 144 {
		oldModel.SeventhDayAdvClickCount = oldModel.SeventhDayAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculateSecondAdv(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) {
	if oldModel.TotalAdvCount == 2 {
		oldModel.SecondAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.SecondAdvHour = modelResponse.FirstAdvClickHour
		oldModel.SecondAdvMinute = modelResponse.FirstAdvClickMinute
		oldModel.SecondAdvType = modelResponse.FirstAdvType
	}
}

func CalculateThirdAdv(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) {
	if oldModel.TotalAdvCount == 3 {
		oldModel.ThirdAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.ThirdAdvHour = modelResponse.FirstAdvClickHour
		oldModel.ThirdAdvMinute = modelResponse.FirstAdvClickMinute
		oldModel.ThirdAdvType = modelResponse.FirstAdvType
	}
}

func CalculateFourthAdv(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) {
	if oldModel.TotalAdvCount == 4 {
		oldModel.FourthAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.FourthAdvHour = modelResponse.FirstAdvClickHour
		oldModel.FourthAdvMinute = modelResponse.FirstAdvClickMinute
		oldModel.FourthAdvType = modelResponse.FirstAdvType
	}
}

func CalculateFifthAdv(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) {
	if oldModel.TotalAdvCount == 5 {
		oldModel.FifthAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.FifthAdvHour = modelResponse.FirstAdvClickHour
		oldModel.FifthAdvMinute = modelResponse.FirstAdvClickMinute
		oldModel.FifthAdvType = modelResponse.FirstAdvType
	}
}

func CalculateSixthAdv(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) {
	if oldModel.TotalAdvCount == 6 {
		oldModel.SixthAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.SixthAdvHour = modelResponse.FirstAdvClickHour
		oldModel.SixthAdvMinute = modelResponse.FirstAdvClickMinute
		oldModel.SixthAdvType = modelResponse.FirstAdvType
	}
}

func CalculateSeventhAdv(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) {
	if oldModel.TotalAdvCount == 7 {
		oldModel.SeventhAdvYearOfDay = modelResponse.FirstAdvYearOfDay
		oldModel.SeventhAdvHour = modelResponse.FirstAdvClickHour
		oldModel.SeventhAdvMinute = modelResponse.FirstAdvClickMinute
		oldModel.SeventhAdvType = modelResponse.FirstAdvType
	}
}

func CalculateFirstDayAdvClickCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel, total_adv_minute int32) {
	if total_adv_minute <= 1440 {
		oldModel.FirstDayAdvClickCount = oldModel.FirstDayAdvClickCount + modelResponse.FirstDayAdvClickCount
	}
}

func CalculatePenultimateDayAdvDay(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) (count int16) {
	switch {
	case ((modelResponse.FirstAdvYearOfDay + (365 * modelResponse.FirstAdvYear)) > (oldModel.LastAdvYearOfDay + (365 * oldModel.LastAdvYear))) && ((modelResponse.FirstAdvYearOfDay + 365*modelResponse.FirstAdvYear) != (oldModel.FirstAdvYearOfDay + 365*oldModel.FirstAdvYear)):
		return oldModel.LastDayAdvClickCount

	default:
		return oldModel.PenultimateDayAdvClickCount
	}
}

func CalculateLastDayAdvClickCount(modelResponse *model.AdvEventResponseModel, oldModel *model.AdvEventResponseModel) int16 {
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

func DetermineAdvDay(modelResponse *model.AdvEventResponseModel, day int16) {
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

func DetermineAdvHour(modelResponse *model.AdvEventResponseModel, hour int16) {
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

func DetermineAdvAmPm(modelResponse *model.AdvEventResponseModel, hour int16) {
	switch {
	case hour <= 12:
		modelResponse.AmAdvClickCount = 1
	default:
		modelResponse.PmAdvClickCount = 1
	}
}

func CalculateAdvLevelBasedAvgClickCount(modelResponse *model.AdvEventResponseModel) {
	switch modelResponse.LevelIndex {
	case 0:
		modelResponse.LevelBasedAverageAdvCount = float32(modelResponse.TotalAdvCount)
	default:
		modelResponse.LevelBasedAverageAdvCount = float32(modelResponse.TotalAdvCount) / float32(modelResponse.LevelIndex)

	}

}
