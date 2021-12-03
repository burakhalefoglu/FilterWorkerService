package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IScreenClickDal "FilterWorkerService/internal/repository/abstract"

	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"FilterWorkerService/pkg/logger"
)

type screenClickManager struct {
	IScreenClickDal *IScreenClickDal.IScreenClickDal
	IJsonParser     *IJsonParser.IJsonParser
	ILog          *logger.ILog
}

func ScreenClickManagerConstructor() *screenClickManager {
	return &screenClickManager{
		IScreenClickDal: &IoC.ScreenClickDal,
		IJsonParser:     &IoC.JsonParser,
		ILog:          &IoC.Logger,
	}
}

func (sc *screenClickManager) ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string) {
	firstModel := model.ScreenClickModel{}
	convertErr := (*sc.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		(*sc.ILog).SendErrorLog("ScreenClickManager", "ConvertRawModelToResponseModel",
			"byte array to ScreenClickModel", "Json Parser Decode Err: ", convertErr.Error())
		return &model.ScreenClickRespondModel{},false, convertErr.Error()
	}
	hour := int64(firstModel.CreationAt.Hour())
	yearOfDay := int64(firstModel.CreationAt.YearDay())
	year := int64(firstModel.CreationAt.Year())
	minute := int64(firstModel.CreationAt.Minute())
	modelResponse := model.ScreenClickRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.FirstClickSessionYearOfDay = yearOfDay
	modelResponse.FirstClickSessionYear = year
	modelResponse.FirstClickSessionHour = hour
	modelResponse.FirstClickSessionMinute = minute
	modelResponse.FirstTouchCount = int64(firstModel.TouchCount)
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

	modelResponse.PenultimateClickSessionHour = 0
	modelResponse.PenultimateClickSessionMinute = 0
	modelResponse.PenultimateTouchCount = 0
	modelResponse.LastClickSessionYearOfDay = 0
	modelResponse.LastClickSessionYear = 0
	modelResponse.LastClickSessionHour = 0
	modelResponse.LastClickSessionMinute = 0
	modelResponse.LastTouchCount = 0

	modelResponse.FirstStartXCor = firstModel.StartXCor
	modelResponse.FirstStartYCor = firstModel.StartYCor
	modelResponse.FirstFinishXCor = firstModel.FinishXCor
	modelResponse.FirstFinishYCor = firstModel.FinishYCor
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

	modelResponse.PenultimateStartXCor = 0
	modelResponse.PenultimateStartYCor = 0
	modelResponse.PenultimateFinishXCor = 0
	modelResponse.PenultimateFinishYCor = 0
	modelResponse.LastStartXCor = 0
	modelResponse.LastStartYCor = 0
	modelResponse.LastFinishXCor = 0
	modelResponse.LastFinishYCor = 0

	modelResponse.FirstHalfHourTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstHourTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstTwoHourTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstThreeHourTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstSixHourTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstTwelveHourTouchCount = int64(firstModel.TouchCount)

	modelResponse.FirstMinusLastTouchCount = int64(firstModel.TouchCount)
	modelResponse.FirstFingerId = int64(firstModel.FingerId)
	modelResponse.PenultimateFingerId = 0
	modelResponse.LastFingerId = 0

	modelResponse.FirstDayClickCount = int64(firstModel.TouchCount)
	modelResponse.SecondDayClickCount = 0
	modelResponse.ThirdDayClickCount = 0
	modelResponse.FourthDayClickCount = 0
	modelResponse.FifthDayClickCount = 0
	modelResponse.SixthDayClickCount = 0
	modelResponse.SeventhDayClickCount = 0

	modelResponse.TotalClickDay = 1
	modelResponse.TotalClickCount = int64(firstModel.TouchCount)
	modelResponse.TotalClickSessionCount = 1

	modelResponse.TotalClickHour = 0
	modelResponse.TotalClickMinute = 1

	modelResponse.TotalStartXCor = firstModel.StartXCor
	modelResponse.TotalStartYCor = firstModel.StartYCor
	modelResponse.TotalFinishXCor = firstModel.FinishXCor
	modelResponse.TotalFinishYCor = firstModel.FinishYCor
	modelResponse.SessionBasedAvegareStartXCor = (firstModel.StartXCor)
	modelResponse.SessionBasedAvegareStartYCor = (firstModel.StartYCor)
	modelResponse.SessionBasedAvegareFinishXCor = (firstModel.FinishXCor)
	modelResponse.SessionBasedAvegareFinishYCor = (firstModel.FinishYCor)
	modelResponse.SessionBasedAvegareClickCount = float64(firstModel.TouchCount)
	modelResponse.DailyAvegareClickCount = float64(firstModel.TouchCount)
	modelResponse.LastTouchCountMinusSessionBasedAvegareClickCount = 0

	defer (*sc.ILog).SendInfoLog("ScreenClickManager", "ConvertRawModelToResponseModel",
		modelResponse.ClientId, modelResponse.ProjectId)
	oldModel, err := (*sc.IScreenClickDal).GetScreenClickById(modelResponse.ClientId)
	if err != nil && err.Error() != "null data error" {
		(*sc.ILog).SendErrorLog("ScreenClickManager", "ConvertRawModelToResponseModel",
			"ScreenClickDal_GetScreenClickById", err.Error())
	}
	switch {
	case err != nil && err.Error() == "null data error":

		logErr := (*sc.IScreenClickDal).Add(&modelResponse)
		if logErr != nil {
			(*sc.ILog).SendErrorLog("ScreenClickManager", "ConvertRawModelToResponseModel",
				"ScreenClickDal_Add", logErr.Error())
			return &modelResponse,false, logErr.Error()
		}
		return &modelResponse, true, "Added"

	case err == nil:
		updModel, updateResult, updateErr := sc.UpdateScreenClick(&modelResponse, oldModel)
		if updateErr != nil {
			return updModel,updateResult, updateErr.Error()
		}
		return updModel,updateResult, "Updated"

	default:

		return &modelResponse,false, err.Error()

	}
}

func (sc *screenClickManager) UpdateScreenClick(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (updatedModel *model.ScreenClickRespondModel, s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	
	oldModel.TotalClickDay = (modelResponse.FirstClickSessionYearOfDay - oldModel.FirstClickSessionYearOfDay) + 365*(modelResponse.FirstClickSessionYear-oldModel.FirstClickSessionYear)
	oldModel.TotalClickCount = oldModel.TotalClickCount + modelResponse.TotalClickCount
	oldModel.TotalClickSessionCount = oldModel.TotalClickSessionCount + modelResponse.TotalClickSessionCount
	oldModel.TotalClickHour = ((modelResponse.FirstClickSessionYearOfDay + 365*modelResponse.FirstClickSessionYear)*24 + modelResponse.FirstClickSessionHour) - ((oldModel.FirstClickSessionYearOfDay + 365*oldModel.FirstClickSessionYear)*24 + oldModel.FirstClickSessionHour)
	oldModel.TotalClickMinute = (((modelResponse.FirstClickSessionYearOfDay + 365*modelResponse.FirstClickSessionYear)*24 + modelResponse.FirstClickSessionHour)*60 + modelResponse.FirstClickSessionMinute) - (((oldModel.FirstClickSessionYearOfDay + 365*oldModel.FirstClickSessionYear)*24 + oldModel.FirstClickSessionHour)*60 + oldModel.FirstClickSessionMinute)
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
	oldModel.FirstMinusLastTouchCount = oldModel.FirstTouchCount - oldModel.LastTouchCount
	
	oldModel.PenultimateFingerId = oldModel.LastFingerId
	oldModel.LastFingerId = modelResponse.FirstFingerId
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
	oldModel.SessionBasedAvegareStartXCor = oldModel.TotalStartXCor / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareStartYCor = oldModel.TotalStartYCor / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareFinishXCor = oldModel.TotalFinishXCor / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareFinishYCor = oldModel.TotalFinishYCor / float64(oldModel.TotalClickSessionCount)
	oldModel.SessionBasedAvegareClickCount = float64(oldModel.TotalClickCount) / float64(oldModel.TotalClickSessionCount)
	oldModel.DailyAvegareClickCount = CalculateDailyAverageClickCount(oldModel)
	oldModel.LastTouchCountMinusSessionBasedAvegareClickCount = float64(oldModel.LastTouchCount) - float64(oldModel.SessionBasedAvegareClickCount)

	defer (*sc.ILog).SendInfoLog("ScreenClickManager", "UpdateScreenClick",
		oldModel.ClientId, oldModel.ProjectId)
	logErr := (*sc.IScreenClickDal).UpdateScreenClickById(oldModel.ClientId, oldModel)
	if logErr != nil {
		(*sc.ILog).SendErrorLog("ScreenClickManager", "UpdateScreenClick",
			"ScreenClickDal_UpdateScreenClickById", logErr.Error())
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateDailyAverageClickCount(oldModel *model.ScreenClickRespondModel)(count float64){
	if oldModel.TotalClickDay == 0 {
		oldModel.DailyAvegareClickCount = float64(oldModel.TotalClickCount)
		return oldModel.DailyAvegareClickCount
	}
	return float64(oldModel.TotalClickCount) / float64(oldModel.TotalClickDay)
}

func CalculateClickHalfHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) {
	switch {
	case total_session_minute<=30:
		oldModel.FirstHalfHourTouchCount = oldModel.FirstHalfHourTouchCount + modelResponse.FirstTouchCount
	}
}

func CalculateClickHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) {
	switch  {
	case total_session_minute<=60:
		oldModel.FirstHourTouchCount = oldModel.FirstHourTouchCount + modelResponse.FirstTouchCount
	}
}

func CalculateClickTwoHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) {
	switch  {
	case total_session_minute<=120:
		oldModel.FirstTwoHourTouchCount = oldModel.FirstTwoHourTouchCount + modelResponse.FirstTouchCount
	}
}

func CalculateClickThreeHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) {
	switch  {
	case total_session_minute<=180:
		oldModel.FirstThreeHourTouchCount = oldModel.FirstThreeHourTouchCount + modelResponse.FirstTouchCount
	}
}

func CalculateClickSixHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) {
	switch  {
	case total_session_minute<=360:
		oldModel.FirstSixHourTouchCount = oldModel.FirstSixHourTouchCount + modelResponse.FirstTouchCount
	}
}

func CalculateClickTwelveHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) {
	switch  {
	case total_session_minute<=720:
		oldModel.FirstTwelveHourTouchCount = oldModel.FirstTwelveHourTouchCount + modelResponse.FirstTouchCount
	}
}

func CalculateClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel){
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
	}
}


func CalculateFirstDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) {
	if  total_session_minute <= 1440{
		oldModel.FirstDayClickCount = oldModel.FirstDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateSecondDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) {
	if  total_session_hour <= 48 && total_session_hour > 24{
		oldModel.SecondDayClickCount = oldModel.SecondDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateThirdDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) {
	if  total_session_hour <= 72 && total_session_hour > 48{
		oldModel.ThirdDayClickCount = oldModel.ThirdDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateFourthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) {
	if  total_session_hour <= 96 && total_session_hour > 72{
		oldModel.FourthDayClickCount = oldModel.FourthDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateFifthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) {
	if  total_session_hour <= 120 && total_session_hour > 96{
		oldModel.FifthDayClickCount = oldModel.FifthDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateSixthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) {
	if  total_session_hour <= 144 && total_session_hour > 120{
		oldModel.SixthDayClickCount = oldModel.SixthDayClickCount + modelResponse.FirstDayClickCount
	}
}

func CalculateSeventhDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) {
	if  total_session_hour <= 168 && total_session_hour > 144{
		oldModel.SeventhDayClickCount = oldModel.SeventhDayClickCount + modelResponse.FirstDayClickCount
	}
}

