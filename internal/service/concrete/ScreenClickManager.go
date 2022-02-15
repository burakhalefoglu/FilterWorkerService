package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IScreenClickDal "FilterWorkerService/internal/repository/abstract"
	"log"

	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type screenClickManager struct {
	IScreenClickDal *IScreenClickDal.IScreenClickDal
	IJsonParser     *IJsonParser.IJsonParser
}

func ScreenClickManagerConstructor() *screenClickManager {
	return &screenClickManager{
		IScreenClickDal: &IoC.ScreenClickDal,
		IJsonParser:     &IoC.JsonParser,
	}
}

func (sc *screenClickManager) ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string) {
	firstModel := model.ScreenClickModel{}
	convertErr := (*sc.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		log.Fatal("ScreenClickManager", "ConvertRawModelToResponseModel",
			"byte array to ScreenClickModel", "Json Parser Decode Err: ", convertErr.Error())
		return &model.ScreenClickRespondModel{}, false, convertErr.Error()
	}
	hour := int16(firstModel.CreationAt.Hour())
	yearOfDay := int16(firstModel.CreationAt.YearDay())
	year := int16(firstModel.CreationAt.Year())
	minute := int16(firstModel.CreationAt.Minute())
	touchCount := int32(firstModel.TouchCount)
	modelResponse := model.ScreenClickRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int16(firstModel.LevelIndex)
	modelResponse.FirstClickSessionYearOfDay = yearOfDay
	modelResponse.FirstClickSessionYear = year
	modelResponse.FirstClickSessionHour = hour
	modelResponse.FirstClickSessionMinute = minute
	modelResponse.FirstTouchCount = int16(firstModel.TouchCount)
	modelResponse.SecondClickSessionHour = 0
	modelResponse.SecondClickSessionMinute = 0
	modelResponse.SecondTouchCount = 0
	modelResponse.ThirdClickSessionHour = 0
	modelResponse.ThirdClickSessionMinute = 0
	modelResponse.ThirdTouchCount = 0

	modelResponse.FourthClickSessionHour = 0
	modelResponse.FourthClickSessionMinute = 0
	modelResponse.FourthTouchCount = 0
	modelResponse.FifthClickSessionHour = 0
	modelResponse.FifthClickSessionMinute = 0
	modelResponse.FifthTouchCount = 0
	modelResponse.SixthClickSessionHour = 0
	modelResponse.SixthClickSessionMinute = 0
	modelResponse.SixthTouchCount = 0
	modelResponse.SeventhClickSessionHour = 0
	modelResponse.SeventhClickSessionMinute = 0
	modelResponse.SeventhTouchCount = 0

	modelResponse.PenultimateClickSessionHour = 0
	modelResponse.PenultimateClickSessionMinute = 0
	modelResponse.PenultimateTouchCount = 0
	modelResponse.LastClickSessionYearOfDay = 0
	modelResponse.LastClickSessionYear = 0
	modelResponse.LastClickSessionHour = 0
	modelResponse.LastClickSessionMinute = 0
	modelResponse.LastTouchCount = 0

	modelResponse.FirstStartXCor = float32(firstModel.StartXCor)
	modelResponse.FirstStartYCor = float32(firstModel.StartYCor)
	modelResponse.FirstFinishXCor = float32(firstModel.FinishXCor)
	modelResponse.FirstFinishYCor = float32(firstModel.FinishYCor)
	modelResponse.SecondStartXCor = 0
	modelResponse.SecondStartYCor = 0
	modelResponse.SecondFinishXCor = 0
	modelResponse.SecondFinishYCor = 0
	modelResponse.ThirdStartXCor = 0
	modelResponse.ThirdStartYCor = 0
	modelResponse.ThirdFinishXCor = 0
	modelResponse.ThirdFinishYCor = 0

	modelResponse.FourthStartXCor = 0
	modelResponse.FourthStartYCor = 0
	modelResponse.FourthFinishXCor = 0
	modelResponse.FourthFinishYCor = 0
	modelResponse.FifthStartXCor = 0
	modelResponse.FifthStartYCor = 0
	modelResponse.FifthFinishXCor = 0
	modelResponse.FifthFinishYCor = 0

	modelResponse.SixthStartXCor = 0
	modelResponse.SixthStartYCor = 0
	modelResponse.SixthFinishXCor = 0
	modelResponse.SixthFinishYCor = 0
	modelResponse.SeventhStartXCor = 0
	modelResponse.SeventhStartYCor = 0
	modelResponse.SeventhFinishXCor = 0
	modelResponse.SeventhFinishYCor = 0

	modelResponse.PenultimateStartXCor = 0
	modelResponse.PenultimateStartYCor = 0
	modelResponse.PenultimateFinishXCor = 0
	modelResponse.PenultimateFinishYCor = 0
	modelResponse.LastStartXCor = 0
	modelResponse.LastStartYCor = 0
	modelResponse.LastFinishXCor = 0
	modelResponse.LastFinishYCor = 0

	modelResponse.FirstFiveMinutesTouchCount = touchCount
	modelResponse.FirstTenMinutesTouchCount = touchCount
	modelResponse.FirstQuarterHourTouchCount = touchCount
	modelResponse.FirstHalfHourTouchCount = touchCount
	modelResponse.FirstHourTouchCount = touchCount
	modelResponse.FirstTwoHourTouchCount = touchCount
	modelResponse.FirstThreeHourTouchCount = touchCount
	modelResponse.FirstSixHourTouchCount = touchCount
	modelResponse.FirstTwelveHourTouchCount = touchCount

	modelResponse.FirstMinusLastTouchCount = int16(firstModel.TouchCount)
	modelResponse.FirstFingerId = byte(firstModel.FingerId)
	modelResponse.PenultimateFingerId = 0
	modelResponse.LastFingerId = 0

	modelResponse.FirstDayClickCount = touchCount
	modelResponse.SecondDayClickCount = 0
	modelResponse.ThirdDayClickCount = 0
	modelResponse.FourthDayClickCount = 0
	modelResponse.FifthDayClickCount = 0
	modelResponse.SixthDayClickCount = 0
	modelResponse.SeventhDayClickCount = 0

	modelResponse.TotalClickDay = 1
	modelResponse.TotalClickCount = touchCount
	modelResponse.TotalClickSessionCount = 1

	modelResponse.TotalClickHour = 0
	modelResponse.TotalClickMinute = 1

	modelResponse.TotalStartXCor = float32(firstModel.StartXCor)
	modelResponse.TotalStartYCor = float32(firstModel.StartYCor)
	modelResponse.TotalFinishXCor = float32(firstModel.FinishXCor)
	modelResponse.TotalFinishYCor = float32(firstModel.FinishYCor)
	modelResponse.SessionBasedAvegareStartXCor = float32(firstModel.StartXCor)
	modelResponse.SessionBasedAvegareStartYCor = float32(firstModel.StartYCor)
	modelResponse.SessionBasedAvegareFinishXCor = float32(firstModel.FinishXCor)
	modelResponse.SessionBasedAvegareFinishYCor = float32(firstModel.FinishYCor)
	modelResponse.SessionBasedAvegareClickCount = float32(firstModel.TouchCount)
	modelResponse.DailyAvegareClickCount = float32(firstModel.TouchCount)
	modelResponse.LastTouchCountMinusSessionBasedAvegareClickCount = 0

	defer log.Print("ScreenClickManager", "ConvertRawModelToResponseModel",
		modelResponse.ClientId, modelResponse.ProjectId)
	oldModel, err := (*sc.IScreenClickDal).GetScreenClickById(modelResponse.ClientId)
	if err != nil && err.Error() != "null data error" {
		log.Fatal("ScreenClickManager", "ConvertRawModelToResponseModel",
			"ScreenClickDal_GetScreenClickById", err.Error())
	}
	switch {
	case err != nil && err.Error() == "null data error":

		logErr := (*sc.IScreenClickDal).Add(&modelResponse)
		if logErr != nil {
			log.Fatal("ScreenClickManager", "ConvertRawModelToResponseModel",
				"ScreenClickDal_Add", logErr.Error())
			return &modelResponse, false, logErr.Error()
		}
		return &modelResponse, true, "Added"

	case err == nil:
		updModel, updateResult, updateErr := sc.UpdateScreenClick(&modelResponse, oldModel)
		if updateErr != nil {
			return updModel, updateResult, updateErr.Error()
		}
		return updModel, updateResult, "Updated"

	default:

		return &modelResponse, false, err.Error()

	}
}

func (sc *screenClickManager) UpdateScreenClick(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (updatedModel *model.ScreenClickRespondModel, s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex

	oldModel.TotalClickDay = (int32(modelResponse.FirstClickSessionYearOfDay) - int32(oldModel.FirstClickSessionYearOfDay)) + 365*(int32(modelResponse.FirstClickSessionYear)-int32(oldModel.FirstClickSessionYear))
	oldModel.TotalClickCount = oldModel.TotalClickCount + modelResponse.TotalClickCount
	oldModel.TotalClickSessionCount = oldModel.TotalClickSessionCount + modelResponse.TotalClickSessionCount
	oldModel.TotalClickHour = ((int32(modelResponse.FirstClickSessionYearOfDay)+365*int32(modelResponse.FirstClickSessionYear))*24 + int32(modelResponse.FirstClickSessionHour)) - ((int32(oldModel.FirstClickSessionYearOfDay)+365*int32(oldModel.FirstClickSessionYear))*24 + int32(oldModel.FirstClickSessionHour))
	oldModel.TotalClickMinute = (((int32(modelResponse.FirstClickSessionYearOfDay)+365*int32(modelResponse.FirstClickSessionYear))*24+int32(modelResponse.FirstClickSessionHour))*60 + int32(modelResponse.FirstClickSessionMinute)) - (((int32(oldModel.FirstClickSessionYearOfDay)+365*int32(oldModel.FirstClickSessionYear))*24+int32(oldModel.FirstClickSessionHour))*60 + int32(oldModel.FirstClickSessionMinute))
	CalculateClickCount(modelResponse, oldModel)

	CalculateFirstDayClickCount(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateSecondDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	CalculateThirdDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	CalculateFourthDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	CalculateFifthDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	CalculateSixthDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	CalculateSeventhDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)

	oldModel.PenultimateClickSessionHour = oldModel.LastClickSessionHour
	oldModel.PenultimateClickSessionMinute = oldModel.LastClickSessionMinute
	oldModel.PenultimateTouchCount = oldModel.LastTouchCount
	oldModel.LastClickSessionYearOfDay = modelResponse.FirstClickSessionYearOfDay
	oldModel.LastClickSessionYear = modelResponse.FirstClickSessionYear
	oldModel.LastClickSessionHour = modelResponse.FirstClickSessionHour
	oldModel.LastClickSessionMinute = modelResponse.FirstClickSessionMinute
	oldModel.LastTouchCount = modelResponse.FirstTouchCount

	oldModel.PenultimateStartXCor = oldModel.LastStartXCor
	oldModel.PenultimateStartYCor = oldModel.LastStartYCor
	oldModel.PenultimateFinishXCor = oldModel.LastFinishXCor
	oldModel.PenultimateFinishYCor = oldModel.LastFinishYCor
	oldModel.LastStartXCor = modelResponse.FirstStartXCor
	oldModel.LastStartYCor = modelResponse.FirstStartYCor
	oldModel.LastFinishXCor = modelResponse.FirstFinishXCor
	oldModel.LastFinishYCor = modelResponse.FirstFinishYCor
	oldModel.FirstMinusLastTouchCount = int16(oldModel.FirstTouchCount) - int16(oldModel.LastTouchCount)

	oldModel.PenultimateFingerId = oldModel.LastFingerId
	oldModel.LastFingerId = modelResponse.FirstFingerId

	CalculateClickFiveMinutes(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickTenMinutes(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickQuarterHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickHalfHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickTwoHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickThreeHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickSixHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	CalculateClickTwelveHour(modelResponse, oldModel, oldModel.TotalClickMinute)

	oldModel.TotalStartXCor = oldModel.TotalStartXCor + modelResponse.TotalStartXCor
	oldModel.TotalStartYCor = oldModel.TotalStartYCor + modelResponse.TotalStartYCor
	oldModel.TotalFinishXCor = oldModel.TotalFinishXCor + modelResponse.TotalFinishXCor
	oldModel.TotalFinishYCor = oldModel.TotalFinishYCor + modelResponse.TotalFinishYCor
	oldModel.SessionBasedAvegareStartXCor = oldModel.TotalStartXCor / float32(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareStartYCor = oldModel.TotalStartYCor / float32(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareFinishXCor = oldModel.TotalFinishXCor / float32(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareFinishYCor = oldModel.TotalFinishYCor / float32(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareClickCount = float32(oldModel.TotalClickCount) / float32(oldModel.TotalClickSessionCount)
	oldModel.DailyAvegareClickCount = CalculateDailyAverageClickCount(oldModel)
	oldModel.LastTouchCountMinusSessionBasedAvegareClickCount = float32(oldModel.LastTouchCount) - float32(oldModel.SessionBasedAvegareClickCount)

	defer log.Fatal("ScreenClickManager", "UpdateScreenClick",
		oldModel.ClientId, oldModel.ProjectId)
	logErr := (*sc.IScreenClickDal).UpdateScreenClickById(oldModel.ClientId, oldModel)
	if logErr != nil {
		log.Fatal("ScreenClickManager", "UpdateScreenClick",
			"ScreenClickDal_UpdateScreenClickById", logErr.Error())
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateDailyAverageClickCount(oldModel *model.ScreenClickRespondModel) (count float32) {
	if oldModel.TotalClickDay == 0 {
		oldModel.DailyAvegareClickCount = float32(oldModel.TotalClickCount)
		return oldModel.DailyAvegareClickCount
	}
	return float32(oldModel.TotalClickCount) / float32(oldModel.TotalClickDay)
}

func CalculateClickFiveMinutes(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 5:
		oldModel.FirstFiveMinutesTouchCount = oldModel.FirstFiveMinutesTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickTenMinutes(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 10:
		oldModel.FirstTenMinutesTouchCount = oldModel.FirstTenMinutesTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickQuarterHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 15:
		oldModel.FirstQuarterHourTouchCount = oldModel.FirstQuarterHourTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickHalfHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 30:
		oldModel.FirstHalfHourTouchCount = oldModel.FirstHalfHourTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 60:
		oldModel.FirstHourTouchCount = oldModel.FirstHourTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickTwoHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 120:
		oldModel.FirstTwoHourTouchCount = oldModel.FirstTwoHourTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickThreeHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 180:
		oldModel.FirstThreeHourTouchCount = oldModel.FirstThreeHourTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickSixHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 360:
		oldModel.FirstSixHourTouchCount = oldModel.FirstSixHourTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickTwelveHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 720:
		oldModel.FirstTwelveHourTouchCount = oldModel.FirstTwelveHourTouchCount + int32(modelResponse.FirstTouchCount)
	}
}

func CalculateClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) {
	switch oldModel.TotalClickSessionCount {
	case 2:
		oldModel.SecondClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.SecondClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.SecondTouchCount = modelResponse.FirstTouchCount
		oldModel.SecondStartXCor = modelResponse.FirstStartXCor
		oldModel.SecondStartYCor = modelResponse.FirstStartYCor
		oldModel.SecondFinishXCor = modelResponse.FirstFinishXCor
		oldModel.SecondFinishYCor = modelResponse.FirstFinishYCor
	case 3:
		oldModel.ThirdClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.ThirdClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.ThirdTouchCount = modelResponse.FirstTouchCount
		oldModel.ThirdStartXCor = modelResponse.FirstStartXCor
		oldModel.ThirdStartYCor = modelResponse.FirstStartYCor
		oldModel.ThirdFinishXCor = modelResponse.FirstFinishXCor
		oldModel.ThirdFinishYCor = modelResponse.FirstFinishYCor
	case 4:
		oldModel.FourthClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.FourthClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.FourthTouchCount = modelResponse.FirstTouchCount
		oldModel.FourthStartXCor = modelResponse.FirstStartXCor
		oldModel.FourthStartYCor = modelResponse.FirstStartYCor
		oldModel.FourthFinishXCor = modelResponse.FirstFinishXCor
		oldModel.FourthFinishYCor = modelResponse.FirstFinishYCor
	case 5:
		oldModel.FifthClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.FifthClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.FifthTouchCount = modelResponse.FirstTouchCount
		oldModel.FifthStartXCor = modelResponse.FirstStartXCor
		oldModel.FifthStartYCor = modelResponse.FirstStartYCor
		oldModel.FifthFinishXCor = modelResponse.FirstFinishXCor
		oldModel.FifthFinishYCor = modelResponse.FirstFinishYCor
	case 6:
		oldModel.SixthClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.SixthClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.SixthTouchCount = modelResponse.FirstTouchCount
		oldModel.SixthStartXCor = modelResponse.FirstStartXCor
		oldModel.SixthStartYCor = modelResponse.FirstStartYCor
		oldModel.SixthFinishXCor = modelResponse.FirstFinishXCor
		oldModel.SixthFinishYCor = modelResponse.FirstFinishYCor
	case 7:
		oldModel.SeventhClickSessionHour = modelResponse.FirstClickSessionHour
		oldModel.SeventhClickSessionMinute = modelResponse.FirstClickSessionMinute
		oldModel.SeventhTouchCount = modelResponse.FirstTouchCount
		oldModel.SeventhStartXCor = modelResponse.FirstStartXCor
		oldModel.SeventhStartYCor = modelResponse.FirstStartYCor
		oldModel.SeventhFinishXCor = modelResponse.FirstFinishXCor
		oldModel.SeventhFinishYCor = modelResponse.FirstFinishYCor
	}
}

func CalculateFirstDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int32) {
	if total_session_minute <= 1440 {
		oldModel.FirstDayClickCount = oldModel.FirstDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateSecondDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int32) {
	if total_session_hour <= 48 && total_session_hour > 24 {
		oldModel.SecondDayClickCount = oldModel.SecondDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateThirdDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int32) {
	if total_session_hour <= 72 && total_session_hour > 48 {
		oldModel.ThirdDayClickCount = oldModel.ThirdDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateFourthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int32) {
	if total_session_hour <= 96 && total_session_hour > 72 {
		oldModel.FourthDayClickCount = oldModel.FourthDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateFifthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int32) {
	if total_session_hour <= 120 && total_session_hour > 96 {
		oldModel.FifthDayClickCount = oldModel.FifthDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateSixthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int32) {
	if total_session_hour <= 144 && total_session_hour > 120 {
		oldModel.SixthDayClickCount = oldModel.SixthDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateSeventhDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int32) {
	if total_session_hour <= 168 && total_session_hour > 144 {
		oldModel.SeventhDayClickCount = oldModel.SeventhDayClickCount + modelResponse.FirstDayClickCount
	}
}
