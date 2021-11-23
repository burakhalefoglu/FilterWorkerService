package concrete

import (
	model "FilterWorkerService/internal/model"
	IGameSessionEveryLoginDal "FilterWorkerService/internal/repository/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type GameSessionEveryLoginManager struct {
	IGameSessionEveryLoginDal IGameSessionEveryLoginDal.IGameSessionEveryLoginDal
	IJsonParser               IJsonParser.IJsonParser
}

func (g *GameSessionEveryLoginManager) ConvertRawModelToResponseModel(data *[]byte) (gameSession *model.GameSessionEveryLoginRespondModel, s bool, m string) {
	firstModel := model.GameSessionEveryLoginModel{}
	err := g.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return &model.GameSessionEveryLoginRespondModel{}, false, err.Error()
	}
	hour := int64(firstModel.SessionFinishTime.Hour())
	yearOfDay := int64(firstModel.SessionFinishTime.YearDay())
	year := int64(firstModel.SessionFinishTime.Year())
	weekDay := int64(firstModel.SessionFinishTime.Weekday())
	minute := int64(firstModel.SessionFinishTime.Minute())
	//minute := int64(firstModel.SessionTimeMinute)
	modelResponse := model.GameSessionEveryLoginRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.FirstSessionYearOfDay = yearOfDay
	modelResponse.FirstSessionYear = year
	modelResponse.FirstSessionWeekDay = weekDay
	modelResponse.FirstSessionHour = hour
	modelResponse.FirstSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.FirstSessionMinute = minute
	modelResponse.SecondSessionHour = 0
	modelResponse.SecondSessionDuration = 0
	modelResponse.SecondSessionMinute = 0
	modelResponse.ThirdSessionHour = 0
	modelResponse.ThirdSessionDuration = 0
	modelResponse.ThirdSessinMinute = 0
	modelResponse.PenultimateSessionHour = 0
	modelResponse.PenultimateSessionDuration = 0
	modelResponse.PenultimateSessionMinute = 0
	modelResponse.PenultimateSessionHour = 0
	modelResponse.PenultimateSessionDuration = 0
	modelResponse.PenultimateSessionMinute = 0
	modelResponse.LastSessionYearOfDay = 0
	modelResponse.LastSessionYear = 0
	modelResponse.LastSessionHour = 0
	modelResponse.LastSessionDuration = 0
	modelResponse.LastSessionMinute = 0
	modelResponse.LastDurationMinusPenultimateDuration = 0

	modelResponse.FirstHalfHourTotalSessionCount = 1
	modelResponse.FirstHalfHourTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.FirstHourTotalSessionCount = 1
	modelResponse.FirstHourTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.FirstTwoHourTotalSessionCount = 1
	modelResponse.FirstTwoHourTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.FirstThreeHourTotalSessionCount = 1
	modelResponse.FirstThreeHourTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.FirstSixHourTotalSessionCount = 1
	modelResponse.FirstSixHourTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.FirstTwelveHourTotalSessionCount = 1
	modelResponse.FirstTwelveHourTotalSessionDuration = int64(firstModel.SessionTimeMinute)

	modelResponse.TotalSessionDay = 1
	modelResponse.TotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.TotalSessionCount = 1
	modelResponse.TotalSessionHour = 1
	modelResponse.TotalSessionMinute = int64(firstModel.SessionTimeMinute)

	modelResponse.FirstDayTotalSessionCount = 1
	modelResponse.FirstDayTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.SecondDayTotalSessionCount = 0
	modelResponse.SecondDayTotalSessionDuration = 0
	modelResponse.ThirdDayTotalSessionCount = 0
	modelResponse.ThirdDayTotalSessionDuration = 0
	modelResponse.FourthDayTotalSessionCount = 0
	modelResponse.FourthDayTotalSessionDuration = 0
	modelResponse.FifthDayTotalSessionCount = 0
	modelResponse.FifthDayTotalSessionDuration = 0
	modelResponse.SixthDayTotalSessionCount = 0
	modelResponse.SixthDayTotalSessionDuration = 0
	modelResponse.SeventhDayTotalSessionCount = 0
	modelResponse.SeventhDayTotalSessionDuration = 0

	modelResponse.MinSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.MaxSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.DailyAvegareSessionCount = 1
	modelResponse.DailyAverageSessionDuration = float64(firstModel.SessionTimeMinute)
	modelResponse.SessionBasedAvegareSessionDuration = float64(firstModel.SessionTimeMinute)
	modelResponse.DailyAvegareSessionCountMinusFirstDaySessionCount = 0
	modelResponse.DailyAvegareSessionDurationMinusFirstDaySessionDuration = 0
	modelResponse.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = 0
	modelResponse.SessionBasedAvegareSessionDurationMinusLastSessionDuration = float64(firstModel.SessionTimeMinute)
	DetermineGameSessionDay(&modelResponse, weekDay)
	DetermineGameSessionHour(&modelResponse, hour)
	DetermineGameSessionAmPm(&modelResponse, hour)

	oldModel, err := g.IGameSessionEveryLoginDal.GetGameSessionEveryLoginById(modelResponse.ClientId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := g.IGameSessionEveryLoginDal.Add(&modelResponse)
		if logErr != nil {
			return &modelResponse, false, logErr.Error()
		}
		return &modelResponse, true, "Added"

	case err == nil:
		updatedModel, updateResult, updateErr := g.UpdateGameSession(&modelResponse, oldModel)
		if updateErr != nil {
			return updatedModel, updateResult, updateErr.Error()
		}
		return updatedModel, updateResult, "Updated"

	default:

		return &modelResponse, false, err.Error()

	}
}

func (g *GameSessionEveryLoginManager) UpdateGameSession(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (updatedModel *model.GameSessionEveryLoginRespondModel, s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId

	//oldModel.FirstSessionYearOfDay
	//oldModel.FirstSessionYear
	//oldModel.FirstSessionWeekDay
	//oldModel.FirstSessionHour
	//oldModel.FirstSessionDuration
	//oldModel.FirstSessionMinute

	oldModel.SecondSessionHour, oldModel.SecondSessionDuration, oldModel.SecondSessionMinute = CalculateSecondGameSession(modelResponse, oldModel)
	oldModel.ThirdSessionHour, oldModel.ThirdSessionDuration, oldModel.ThirdSessinMinute = CalculateThirdGameSession(modelResponse, oldModel)

	oldModel.TotalSessionDay = (modelResponse.FirstSessionYearOfDay - oldModel.FirstSessionYearOfDay) + 365*(modelResponse.FirstSessionYear-oldModel.FirstSessionYear)
	oldModel.TotalSessionDuration = oldModel.TotalSessionDuration + modelResponse.TotalSessionDuration
	oldModel.TotalSessionCount = oldModel.TotalSessionCount + modelResponse.TotalSessionCount

	oldModel.TotalSessionHour = ((modelResponse.FirstSessionYearOfDay+365*modelResponse.FirstSessionYear)*24 + modelResponse.FirstSessionHour) - ((oldModel.FirstSessionYearOfDay+365*oldModel.FirstSessionYear)*24 + oldModel.FirstSessionHour)
	oldModel.TotalSessionMinute = (((modelResponse.FirstSessionYearOfDay+365*modelResponse.FirstSessionYear)*24+modelResponse.FirstSessionHour)*60 + modelResponse.FirstSessionMinute) - (((oldModel.FirstSessionYearOfDay+365*oldModel.FirstSessionYear)*24+oldModel.FirstSessionHour)*60 + oldModel.FirstSessionMinute)

	oldModel.FirstDayTotalSessionCount, oldModel.FirstDayTotalSessionDuration = CalculateFirstTwentyFourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)

	oldModel.SecondDayTotalSessionCount, oldModel.SecondDayTotalSessionDuration = CalculateSecondDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	oldModel.ThirdDayTotalSessionCount, oldModel.ThirdDayTotalSessionDuration = CalculateThirdDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	oldModel.FourthDayTotalSessionCount, oldModel.FourthDayTotalSessionDuration = CalculateFourthDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	oldModel.FifthDayTotalSessionCount, oldModel.FifthDayTotalSessionDuration = CalculateFifthDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	oldModel.SixthDayTotalSessionCount, oldModel.SixthDayTotalSessionDuration = CalculateSixthDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	oldModel.SeventhDayTotalSessionCount, oldModel.SeventhDayTotalSessionDuration = CalculateSeventhDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)

	oldModel.FirstHalfHourTotalSessionCount, oldModel.FirstHalfHourTotalSessionDuration = CalculateFirstHalfHourTotalSessionCountAndSessionDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	oldModel.FirstHourTotalSessionCount, oldModel.FirstHourTotalSessionDuration = CalculateFirstHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	oldModel.FirstTwoHourTotalSessionCount, oldModel.FirstTwoHourTotalSessionDuration = CalculateFirstTwoHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	oldModel.FirstThreeHourTotalSessionCount, oldModel.FirstThreeHourTotalSessionDuration = CalculateFirstThreeHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	oldModel.FirstSixHourTotalSessionCount, oldModel.FirstSixHourTotalSessionDuration = CalculateFirstSixHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	oldModel.FirstTwelveHourTotalSessionCount, oldModel.FirstTwelveHourTotalSessionDuration = CalculateFirstTwelveHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)

	oldModel.PenultimateSessionDuration = oldModel.LastSessionDuration
	oldModel.PenultimateSessionHour = oldModel.LastSessionHour
	oldModel.PenultimateSessionMinute = oldModel.LastSessionMinute
	oldModel.LastSessionYearOfDay = modelResponse.FirstSessionYearOfDay
	oldModel.LastSessionYear = modelResponse.FirstSessionYear
	oldModel.LastSessionHour = modelResponse.FirstSessionHour
	oldModel.LastSessionDuration = modelResponse.FirstSessionDuration
	oldModel.LastDurationMinusPenultimateDuration = oldModel.LastSessionDuration - oldModel.PenultimateSessionDuration

	oldModel.MinSessionDuration = CalculateMinDuration(modelResponse, oldModel)
	oldModel.MaxSessionDuration = CalculateMaxDuration(modelResponse, oldModel)
	oldModel.DailyAvegareSessionCount, oldModel.DailyAverageSessionDuration = CalculateDailyAvegareSessionCountAndDuration(oldModel)
	oldModel.SessionBasedAvegareSessionDuration = float64(oldModel.TotalSessionDuration) / float64(oldModel.TotalSessionCount)
	oldModel.DailyAvegareSessionCountMinusFirstDaySessionCount = oldModel.DailyAvegareSessionCount - float64(oldModel.FirstDayTotalSessionCount)
	oldModel.DailyAvegareSessionDurationMinusFirstDaySessionDuration = oldModel.DailyAverageSessionDuration - float64(oldModel.FirstDayTotalSessionDuration)
	oldModel.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = oldModel.SessionBasedAvegareSessionDuration - float64(oldModel.FirstSessionDuration)
	oldModel.SessionBasedAvegareSessionDurationMinusLastSessionDuration = oldModel.SessionBasedAvegareSessionDuration - float64(oldModel.LastSessionDuration)
	oldModel.SundaySessionCount = oldModel.SundaySessionCount + modelResponse.SundaySessionCount
	oldModel.MondaySessionCount = oldModel.MondaySessionCount + modelResponse.MondaySessionCount
	oldModel.TuesdaySessionCount = oldModel.TuesdaySessionCount + modelResponse.TuesdaySessionCount
	oldModel.WednesdaySessionCount = oldModel.WednesdaySessionCount + modelResponse.WednesdaySessionCount
	oldModel.ThursdaySessionCount = oldModel.ThursdaySessionCount + modelResponse.ThursdaySessionCount
	oldModel.FridaySessionCount = oldModel.FridaySessionCount + modelResponse.FridaySessionCount
	oldModel.SaturdaySessionCount = oldModel.SaturdaySessionCount + modelResponse.SaturdaySessionCount
	oldModel.AmSessionCount = oldModel.AmSessionCount + modelResponse.AmSessionCount
	oldModel.PmSessionCount = oldModel.PmSessionCount + modelResponse.PmSessionCount
	oldModel.Session0To5HourCount = oldModel.Session0To5HourCount + modelResponse.Session0To5HourCount
	oldModel.Session6To11HourCount = oldModel.Session6To11HourCount + modelResponse.Session6To11HourCount
	oldModel.Session12To17HourCount = oldModel.Session12To17HourCount + modelResponse.Session12To17HourCount
	oldModel.Session18To23HourCount = oldModel.Session18To23HourCount + modelResponse.Session18To23HourCount

	logErr := g.IGameSessionEveryLoginDal.UpdateGameSessionEveryLoginById(oldModel.ClientId, oldModel)
	if logErr != nil {
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateDailyAvegareSessionCountAndDuration(oldModel *model.GameSessionEveryLoginRespondModel)(count float64, duration float64){
	if oldModel.TotalSessionDay == 0{
		return float64(oldModel.TotalSessionCount), float64(oldModel.TotalSessionDuration)
	}
	return float64(oldModel.TotalSessionCount) / float64(oldModel.TotalSessionDay), float64(oldModel.TotalSessionDuration) / float64(oldModel.TotalSessionDay)
}

func CalculateFirstHalfHourTotalSessionCountAndSessionDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_minute int64) (count int64, duration int64) {
	if total_session_minute <= 30 {
		oldModel.FirstHalfHourTotalSessionCount = oldModel.FirstHalfHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstHalfHourTotalSessionDuration = oldModel.FirstHalfHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstHalfHourTotalSessionCount, oldModel.FirstHalfHourTotalSessionDuration
	}
	return oldModel.FirstHalfHourTotalSessionCount, oldModel.FirstHalfHourTotalSessionDuration
}

func CalculateFirstHourTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_minute int64) (count int64, duration int64) {
	if total_session_minute <= 60 {
		oldModel.FirstHourTotalSessionCount = oldModel.FirstHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstHourTotalSessionDuration = oldModel.FirstHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstHourTotalSessionCount, oldModel.FirstHourTotalSessionDuration
	}
	return oldModel.FirstHourTotalSessionCount, oldModel.FirstHourTotalSessionDuration
}

func CalculateFirstTwoHourTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_minute int64) (count int64, duration int64) {
	if total_session_minute <= 120 {
		oldModel.FirstTwoHourTotalSessionCount = oldModel.FirstTwoHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstTwoHourTotalSessionDuration = oldModel.FirstTwoHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstTwoHourTotalSessionCount, oldModel.FirstTwoHourTotalSessionDuration
	}
	return oldModel.FirstTwoHourTotalSessionCount, oldModel.FirstTwoHourTotalSessionDuration
}

func CalculateFirstThreeHourTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_minute int64) (count int64, duration int64) {
	if total_session_minute <= 180 {
		oldModel.FirstThreeHourTotalSessionCount = oldModel.FirstThreeHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstThreeHourTotalSessionDuration = oldModel.FirstThreeHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstThreeHourTotalSessionCount, oldModel.FirstThreeHourTotalSessionDuration
	}
	return oldModel.FirstThreeHourTotalSessionCount, oldModel.FirstThreeHourTotalSessionDuration
}

func CalculateFirstSixHourTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_minute int64) (count int64, duration int64) {
	if total_session_minute <= 360 {
		oldModel.FirstSixHourTotalSessionCount = oldModel.FirstSixHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstSixHourTotalSessionDuration = oldModel.FirstSixHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstSixHourTotalSessionCount, oldModel.FirstSixHourTotalSessionDuration
	}
	return oldModel.FirstSixHourTotalSessionCount, oldModel.FirstSixHourTotalSessionDuration
}

func CalculateFirstTwelveHourTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_minute int64) (count int64, duration int64) {
	if total_session_minute <= 720 {
		oldModel.FirstTwelveHourTotalSessionCount = oldModel.FirstTwelveHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstTwelveHourTotalSessionDuration = oldModel.FirstTwelveHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstTwelveHourTotalSessionCount, oldModel.FirstTwelveHourTotalSessionDuration
	}
	return oldModel.FirstTwelveHourTotalSessionCount, oldModel.FirstTwelveHourTotalSessionDuration
}

func CalculateFirstTwentyFourTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_minute int64) (count int64, duration int64) {
	if total_session_minute <= 1440 {
		oldModel.FirstDayTotalSessionCount = oldModel.FirstDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstDayTotalSessionDuration = oldModel.FirstDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstDayTotalSessionCount, oldModel.FirstDayTotalSessionDuration
	}
	return oldModel.FirstDayTotalSessionCount, oldModel.FirstDayTotalSessionDuration
}

func CalculateSecondDayTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_hour int64) (count int64, duration int64) {
	if total_session_hour <= 48 && total_session_hour > 24 {
		oldModel.SecondDayTotalSessionCount = oldModel.SecondDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.SecondDayTotalSessionDuration = oldModel.SecondDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.SecondDayTotalSessionCount, oldModel.SecondDayTotalSessionDuration
	}
	return oldModel.SecondDayTotalSessionCount, oldModel.SecondDayTotalSessionDuration
}

func CalculateThirdDayTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_hour int64) (count int64, duration int64) {
	if total_session_hour <= 72 && total_session_hour > 48 {
		oldModel.ThirdDayTotalSessionCount = oldModel.ThirdDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.ThirdDayTotalSessionDuration = oldModel.ThirdDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.ThirdDayTotalSessionCount, oldModel.ThirdDayTotalSessionDuration
	}
	return oldModel.ThirdDayTotalSessionCount, oldModel.ThirdDayTotalSessionDuration
}

func CalculateFourthDayTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_hour int64) (count int64, duration int64) {
	if total_session_hour <= 96 && total_session_hour > 72 {
		oldModel.FourthDayTotalSessionCount = oldModel.FourthDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FourthDayTotalSessionDuration = oldModel.FourthDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FourthDayTotalSessionCount, oldModel.FourthDayTotalSessionDuration
	}
	return oldModel.FourthDayTotalSessionCount, oldModel.FourthDayTotalSessionDuration
}

func CalculateFifthDayTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_hour int64) (count int64, duration int64) {
	if total_session_hour <= 120 && total_session_hour > 96 {
		oldModel.FifthDayTotalSessionCount = oldModel.FifthDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FifthDayTotalSessionDuration = oldModel.FifthDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FifthDayTotalSessionCount, oldModel.FifthDayTotalSessionDuration
	}
	return oldModel.FifthDayTotalSessionCount, oldModel.FifthDayTotalSessionDuration
}

func CalculateSixthDayTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_hour int64) (count int64, duration int64) {
	if total_session_hour <= 144 && total_session_hour > 120 {
		oldModel.SixthDayTotalSessionCount = oldModel.SixthDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.SixthDayTotalSessionDuration = oldModel.SixthDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.SixthDayTotalSessionCount, oldModel.SixthDayTotalSessionDuration
	}
	return oldModel.SixthDayTotalSessionCount, oldModel.SixthDayTotalSessionDuration
}

func CalculateSeventhDayTotalSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel, total_session_hour int64) (count int64, duration int64) {
	if total_session_hour <= 168 && total_session_hour > 144 {
		oldModel.SeventhDayTotalSessionCount = oldModel.SeventhDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.SeventhDayTotalSessionDuration = oldModel.SeventhDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.SeventhDayTotalSessionCount, oldModel.SeventhDayTotalSessionDuration
	}
	return oldModel.SeventhDayTotalSessionCount, oldModel.SeventhDayTotalSessionDuration
}

func CalculateSecondGameSession(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (hour int64, duration int64, minute int64) {
	if oldModel.TotalSessionCount == 1 {
		oldModel.SecondSessionHour = modelResponse.FirstSessionHour
		oldModel.SecondSessionDuration = modelResponse.FirstSessionDuration
		oldModel.SecondSessionMinute = modelResponse.FirstSessionMinute
		return oldModel.SecondSessionHour, oldModel.SecondSessionDuration, oldModel.SecondSessionMinute
	}
	return oldModel.SecondSessionHour, oldModel.SecondSessionDuration, oldModel.SecondSessionMinute
}

func CalculateThirdGameSession(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (hour int64, duration int64, minute int64) {
	if oldModel.TotalSessionCount == 2 {
		oldModel.ThirdSessionHour = modelResponse.FirstSessionHour
		oldModel.ThirdSessionDuration = modelResponse.FirstSessionDuration
		oldModel.ThirdSessinMinute = modelResponse.FirstSessionMinute
		return oldModel.ThirdSessionHour, oldModel.ThirdSessionDuration, oldModel.ThirdSessinMinute
	}
	return oldModel.ThirdSessionHour, oldModel.ThirdSessionDuration, oldModel.ThirdSessinMinute
}


func CalculateMinDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (duration int64) {
	if oldModel.MinSessionDuration > modelResponse.MinSessionDuration {
		oldModel.MinSessionDuration = modelResponse.MinSessionDuration
		return oldModel.MinSessionDuration
	}
	return oldModel.MinSessionDuration
}

func CalculateMaxDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (duration int64) {
	if modelResponse.MaxSessionDuration > oldModel.MaxSessionDuration {
		oldModel.MaxSessionDuration = modelResponse.MaxSessionDuration
		return oldModel.MaxSessionDuration
	}
	return oldModel.MaxSessionDuration
}


func DetermineGameSessionDay(modelResponse *model.GameSessionEveryLoginRespondModel, day int64) {
	switch day {
	case 0:
		modelResponse.SundaySessionCount = 1
	case 1:
		modelResponse.MondaySessionCount = 1
	case 2:
		modelResponse.TuesdaySessionCount = 1
	case 3:
		modelResponse.WednesdaySessionCount = 1
	case 4:
		modelResponse.ThursdaySessionCount = 1
	case 5:
		modelResponse.FridaySessionCount = 1
	case 6:
		modelResponse.SaturdaySessionCount = 1
	}
}

func DetermineGameSessionHour(modelResponse *model.GameSessionEveryLoginRespondModel, hour int64) {
	switch {
	case hour <= 5:
		modelResponse.Session0To5HourCount = 1
	case (hour > 5) && (hour <= 11):
		modelResponse.Session6To11HourCount = 1
	case (hour > 11) && (hour <= 17):
		modelResponse.Session12To17HourCount = 1
	case (hour > 17) && (hour <= 23):
		modelResponse.Session18To23HourCount = 1
	}
}

func DetermineGameSessionAmPm(modelResponse *model.GameSessionEveryLoginRespondModel, hour int64) {
	switch {
	case hour <= 12:
		modelResponse.AmSessionCount = 1
	default:
		modelResponse.PmSessionCount = 1
	}
}
