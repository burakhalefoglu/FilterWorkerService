package concrete

import (
	model "FilterWorkerService/internal/model"
	IScreenClickDal "FilterWorkerService/internal/repository/abstract"

	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type ScreenClickManager struct {
	IScreenClickDal IScreenClickDal.IScreenClickDal
	IJsonParser     IJsonParser.IJsonParser
}

func (sc *ScreenClickManager) ConvertRawModelToResponseModel(data *[]byte) (respondModel *model.ScreenClickRespondModel, s bool, m string) {
	firstModel := model.ScreenClickModel{}
	err := sc.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return &model.ScreenClickRespondModel{},false, err.Error()
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

	oldModel, err := sc.IScreenClickDal.GetScreenClickById(modelResponse.ClientId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := sc.IScreenClickDal.Add(&modelResponse)
		if logErr != nil {
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

func (sc *ScreenClickManager) UpdateScreenClick(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel) (updatedModel *model.ScreenClickRespondModel, s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	//oldModel.FirstClickSessionYearOfDay
	//oldModel.FirstClickSessionYear
	//oldModel.FirstClickSessionHour
	//oldModel.FirstClickSessionMinute
	//oldModel.FirstTouchCount
	oldModel.TotalClickDay = (modelResponse.FirstClickSessionYearOfDay - oldModel.FirstClickSessionYearOfDay) + 365*(modelResponse.FirstClickSessionYear-oldModel.FirstClickSessionYear)
	oldModel.TotalClickCount = oldModel.TotalClickCount + modelResponse.TotalClickCount
	oldModel.TotalClickSessionCount = oldModel.TotalClickSessionCount + modelResponse.TotalClickSessionCount
	oldModel.TotalClickHour = ((modelResponse.FirstClickSessionYearOfDay + 365*modelResponse.FirstClickSessionYear)*24 + modelResponse.FirstClickSessionHour) - ((oldModel.FirstClickSessionYearOfDay + 365*oldModel.FirstClickSessionYear)*24 + oldModel.FirstClickSessionHour)
	oldModel.TotalClickMinute = (((modelResponse.FirstClickSessionYearOfDay + 365*modelResponse.FirstClickSessionYear)*24 + modelResponse.FirstClickSessionHour)*60 + modelResponse.FirstClickSessionMinute) - (((oldModel.FirstClickSessionYearOfDay + 365*oldModel.FirstClickSessionYear)*24 + oldModel.FirstClickSessionHour)*60 + oldModel.FirstClickSessionMinute)

	CalculateClickCount(modelResponse, oldModel)

	oldModel.FirstDayClickCount = CalculateFirstDayClickCount(modelResponse, oldModel, oldModel.TotalClickMinute)
	oldModel.SecondDayClickCount = CalculateSecondDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	oldModel.ThirdDayClickCount = CalculateThirdDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	oldModel.FourthDayClickCount = CalculateFourthDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	oldModel.FifthDayClickCount = CalculateFifthDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	oldModel.SixthDayClickCount = CalculateSixthDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)
	oldModel.SeventhDayClickCount = CalculateSeventhDayClickCount(modelResponse, oldModel, oldModel.TotalClickHour)

	oldModel.PenultimateClickSessionHour = oldModel.LastClickSessionHour
	oldModel.PenultimateClickSessionMinute = oldModel.LastClickSessionMinute
	oldModel.PenultimateTouchCount = oldModel.LastTouchCount
	oldModel.LastClickSessionYearOfDay = modelResponse.FirstClickSessionYearOfDay
	oldModel.LastClickSessionYear = modelResponse.FirstClickSessionYear
	oldModel.LastClickSessionHour = modelResponse.FirstClickSessionHour
	oldModel.LastClickSessionMinute = modelResponse.FirstClickSessionMinute
	oldModel.LastTouchCount = modelResponse.FirstTouchCount
	//oldModel.FirstStartXCor
	//oldModel.FirstStartYCor
	//oldModel.FirstFinishXCor
	//oldModel.FirstFinishYCor

	oldModel.PenultimateStartXCor = oldModel.LastStartXCor
	oldModel.PenultimateStartYCor = oldModel.LastStartYCor
	oldModel.PenultimateFinishXCor = oldModel.LastFinishXCor
	oldModel.PenultimateFinishYCor = oldModel.LastFinishYCor
	oldModel.LastStartXCor = modelResponse.FirstStartXCor
	oldModel.LastStartYCor = modelResponse.FirstStartYCor
	oldModel.LastFinishXCor = modelResponse.FirstFinishXCor
	oldModel.LastFinishYCor = modelResponse.FirstFinishYCor
	oldModel.FirstMinusLastTouchCount = oldModel.FirstTouchCount - oldModel.LastTouchCount
	//oldModel.FirstFingerId
	oldModel.PenultimateFingerId = oldModel.LastFingerId
	oldModel.LastFingerId = modelResponse.FirstFingerId
	oldModel.FirstHalfHourTouchCount = CalculateClickHalfHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	oldModel.FirstHourTouchCount = CalculateClickHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	oldModel.FirstTwoHourTouchCount = CalculateClickTwoHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	oldModel.FirstThreeHourTouchCount = CalculateClickThreeHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	oldModel.FirstSixHourTouchCount = CalculateClickSixHour(modelResponse, oldModel, oldModel.TotalClickMinute)
	oldModel.FirstTwelveHourTouchCount = CalculateClickTwentyHour(modelResponse, oldModel, oldModel.TotalClickMinute)

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
	logErr := sc.IScreenClickDal.UpdateScreenClickById(oldModel.ClientId, oldModel)
	if logErr != nil {
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

func CalculateClickHalfHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64)(count int64){
	switch {
	case total_session_minute<=30:
		oldModel.FirstHalfHourTouchCount = oldModel.FirstHalfHourTouchCount + modelResponse.FirstTouchCount
		return oldModel.FirstHalfHourTouchCount
	}
	return oldModel.FirstHalfHourTouchCount
}

func CalculateClickHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) (count int64){
	switch  {
	case total_session_minute<=60:
		oldModel.FirstHourTouchCount = oldModel.FirstHourTouchCount + modelResponse.FirstTouchCount
		return oldModel.FirstHourTouchCount
	}
	return oldModel.FirstHourTouchCount
}

func CalculateClickTwoHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) (count int64){
	switch  {
	case total_session_minute<=120:
		oldModel.FirstTwoHourTouchCount = oldModel.FirstTwoHourTouchCount + modelResponse.FirstTouchCount
		return oldModel.FirstTwoHourTouchCount
	}
	return oldModel.FirstTwoHourTouchCount
}

func CalculateClickThreeHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) (count int64){
	switch  {
	case total_session_minute<=180:
		oldModel.FirstThreeHourTouchCount = oldModel.FirstThreeHourTouchCount + modelResponse.FirstTouchCount
		return oldModel.FirstThreeHourTouchCount
	}
	return oldModel.FirstThreeHourTouchCount
}

func CalculateClickSixHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) (count int64){
	switch  {
	case total_session_minute<=360:
		oldModel.FirstSixHourTouchCount = oldModel.FirstSixHourTouchCount + modelResponse.FirstTouchCount
		return oldModel.FirstSixHourTouchCount
	}
	return oldModel.FirstSixHourTouchCount
}

func CalculateClickTwentyHour(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) (count int64){
	switch  {
	case total_session_minute<=720:
		oldModel.FirstTwelveHourTouchCount = oldModel.FirstTwelveHourTouchCount + modelResponse.FirstTouchCount
		return oldModel.FirstTwelveHourTouchCount
	}
	return oldModel.FirstTwelveHourTouchCount
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


func CalculateFirstDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_minute int64) (count int64) {
	if  total_session_minute <= 1440{
		oldModel.FirstDayClickCount = oldModel.FirstDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.FirstDayClickCount
	}
	return oldModel.FirstDayClickCount
}

func CalculateSecondDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) (count int64) {
	if  total_session_hour <= 48 && total_session_hour > 24{
		oldModel.SecondDayClickCount = oldModel.SecondDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.SecondDayClickCount
	}
	return oldModel.SecondDayClickCount
}

func CalculateThirdDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) (count int64) {
	if  total_session_hour <= 72 && total_session_hour > 48{
		oldModel.ThirdDayClickCount = oldModel.ThirdDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.ThirdDayClickCount
	}
	return oldModel.ThirdDayClickCount
}

func CalculateFourthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) (count int64) {
	if  total_session_hour <= 96 && total_session_hour > 72{
		oldModel.FourthDayClickCount = oldModel.FourthDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.FourthDayClickCount
	}
	return oldModel.FourthDayClickCount
}

func CalculateFifthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) (count int64) {
	if  total_session_hour <= 120 && total_session_hour > 96{
		oldModel.FifthDayClickCount = oldModel.FifthDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.FifthDayClickCount
	}
	return oldModel.FifthDayClickCount
}

func CalculateSixthDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) (count int64) {
	if  total_session_hour <= 144 && total_session_hour > 120{
		oldModel.SixthDayClickCount = oldModel.SixthDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.SixthDayClickCount
	}
	return oldModel.SixthDayClickCount
}

func CalculateSeventhDayClickCount(modelResponse *model.ScreenClickRespondModel, oldModel *model.ScreenClickRespondModel, total_session_hour int64) (count int64) {
	if  total_session_hour <= 168 && total_session_hour > 144{
		oldModel.SeventhDayClickCount = oldModel.SeventhDayClickCount + modelResponse.FirstDayClickCount
		return oldModel.SeventhDayClickCount
	}
	return oldModel.SeventhDayClickCount
}

