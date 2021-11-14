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

func (g *GameSessionEveryLoginManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.GameSessionEveryLoginModel{}
	err := g.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	hour := int64(firstModel.SessionFinishTime.Hour())
	yearOfDay := int64(firstModel.SessionFinishTime.YearDay())
	year := int64(firstModel.SessionFinishTime.Year())
	weekDay := int64(firstModel.SessionFinishTime.Weekday())
	minute := int64(firstModel.SessionFinishTime.Minute())
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
	modelResponse.LastSessionYearOfDay = yearOfDay
	modelResponse.LastSessionYear = year
	modelResponse.LastSessionHour = hour
	modelResponse.LastSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.LastSessionMinute = minute
	modelResponse.LastDurationMinusPenultimateDuration = int64(modelResponse.LastSessionDuration)
	modelResponse.TotalSessionDay = 1
	modelResponse.TotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.TotalSessionCount = 1
	modelResponse.FirstDayTotalSessionCount = 1
	modelResponse.FirstDayTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.PenultimateDayTotalSessionDuration = 0
	modelResponse.PenultimateDayTotalSessionCount = 0
	modelResponse.LastDayTotalSessionCount = 1
	modelResponse.LastDayTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.MinSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.MaxSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.DailyAvegareSessionCount = 1
	modelResponse.DailyAverageSessionDuration = float64(firstModel.SessionTimeMinute)
	modelResponse.SessionBasedAvegareSessionDuration = float64(firstModel.SessionTimeMinute)
	modelResponse.DailyAvegareSessionCountMinusFirstDaySessionCount = 0
	modelResponse.DailyAvegareSessionDurationMinusFirstDaySessionDuration = 0
	modelResponse.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = 0
	modelResponse.SessionBasedAvegareSessionDurationMinusLastSessionDuration = 0
	determineGameSessionDay(&modelResponse, weekDay)
	determineGameSessionHour(&modelResponse, hour)
	determineGameSessionAmPm(&modelResponse, hour)

	oldModel, err := g.IGameSessionEveryLoginDal.GetGameSessionEveryLoginById(modelResponse.ClientId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := g.IGameSessionEveryLoginDal.Add(&modelResponse)
		if logErr != nil {
			return false, logErr.Error()
		}
		return true, ""

	case err == nil:
		updateResult, updateErr :=  g.updateGameSession(&modelResponse, oldModel)
		if updateErr != nil {
		return updateResult, updateErr.Error()
	}
		return updateResult, ""

	default:

		return false, err.Error()

	}
}


func (g *GameSessionEveryLoginManager) updateGameSession(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (s bool, m error) {
	
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	//oldModel.FirstSessionYearOfDay
	//oldModel.FirstSessionYear
	//oldModel.FirstSessionWeekDay
	//oldModel.FirstSessionHour
	//oldModel.FirstSessionDuration
	//oldModel.FirstSessionMinute
	oldModel.SecondSessionHour, oldModel.SecondSessionDuration, oldModel.SecondSessionMinute = calculateSecondGameSession(modelResponse, oldModel)
	modelResponse.ThirdSessionHour, modelResponse.ThirdSessionDuration, modelResponse.ThirdSessinMinute = calculateThirdGameSession(modelResponse, oldModel)
	oldModel.PenultimateSessionDuration = oldModel.LastSessionDuration
	oldModel.PenultimateSessionHour = oldModel.LastSessionHour
	oldModel.PenultimateSessionMinute = oldModel.LastSessionMinute
	oldModel.LastSessionYearOfDay = modelResponse.LastSessionYearOfDay
	oldModel.LastSessionYear = modelResponse.LastSessionYear
	oldModel.LastSessionHour = modelResponse.LastSessionHour
	oldModel.LastSessionDuration = modelResponse.LastSessionDuration
	oldModel.LastDurationMinusPenultimateDuration = oldModel.LastSessionDuration - oldModel.PenultimateSessionDuration
	oldModel.TotalSessionDay = oldModel.LastSessionYearOfDay - oldModel.FirstSessionYearOfDay + 365*(oldModel.LastSessionYear-oldModel.FirstSessionYear)
	oldModel.TotalSessionDuration = oldModel.TotalSessionDuration + modelResponse.TotalSessionDuration
	oldModel.TotalSessionCount = oldModel.TotalSessionCount + modelResponse.TotalSessionCount
	oldModel.FirstDayTotalSessionCount, oldModel.FirstDayTotalSessionDuration = calculateFirstDayGameSessionCountAndDuration(modelResponse, oldModel)
	oldModel.PenultimateDayTotalSessionCount, oldModel.PenultimateDayTotalSessionDuration = calculatePenultimateDay(modelResponse, oldModel)
	oldModel.LastDayTotalSessionCount, oldModel.LastDayTotalSessionDuration = calculateLastDaySessionCountAndDuration(modelResponse, oldModel)
	oldModel.MinSessionDuration = calculateMinDuration(modelResponse, oldModel)
	oldModel.MaxSessionDuration = calculateMaxDuration(modelResponse, oldModel)
	oldModel.DailyAvegareSessionCount = float64(oldModel.TotalSessionCount) / float64(oldModel.TotalSessionDay)
	oldModel.DailyAverageSessionDuration = float64(oldModel.TotalSessionDuration) / float64(oldModel.TotalSessionDay)
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
		return false, logErr
	}
	return true, nil
}

func calculateSecondGameSession(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (hour int64, duration int64, minute int64) {
	if oldModel.TotalSessionCount == 1 {
		oldModel.SecondSessionHour = modelResponse.FirstSessionHour
		oldModel.SecondSessionDuration = modelResponse.FirstSessionDuration
		oldModel.SecondSessionMinute = modelResponse.FirstSessionMinute
		return oldModel.SecondSessionHour, oldModel.SecondSessionDuration, oldModel.SecondSessionMinute
	}
	return oldModel.SecondSessionHour, oldModel.SecondSessionDuration, oldModel.SecondSessionMinute
}

func calculateThirdGameSession(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (hour int64, duration int64, minute int64) {
	if oldModel.TotalSessionCount == 2 {
		oldModel.ThirdSessionHour = modelResponse.FirstSessionHour
		oldModel.ThirdSessionDuration = modelResponse.FirstSessionDuration
		oldModel.ThirdSessinMinute = modelResponse.FirstSessionMinute
		return oldModel.ThirdSessionHour, oldModel.ThirdSessionDuration, oldModel.ThirdSessinMinute
	}
	return oldModel.ThirdSessionHour, oldModel.ThirdSessionDuration, oldModel.ThirdSessinMinute
}

func calculatePenultimateDay(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (count int64, duration int64) {
	if oldModel.LastSessionYearOfDay != modelResponse.LastSessionYearOfDay {
		oldModel.PenultimateDayTotalSessionCount = oldModel.LastDayTotalSessionCount
		oldModel.PenultimateDayTotalSessionDuration = oldModel.LastDayTotalSessionDuration
		return oldModel.PenultimateDayTotalSessionCount, oldModel.PenultimateDayTotalSessionDuration
	}
	return oldModel.PenultimateDayTotalSessionCount, oldModel.PenultimateDayTotalSessionDuration
}

func calculateFirstDayGameSessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (count int64, duration int64) {
	if (oldModel.FirstSessionYearOfDay == modelResponse.FirstSessionYearOfDay) && (oldModel.FirstSessionYear == modelResponse.FirstSessionYear) {
		oldModel.FirstDayTotalSessionCount = oldModel.FirstDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstDayTotalSessionDuration = oldModel.FirstDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstDayTotalSessionCount, oldModel.FirstDayTotalSessionDuration
	}
	return oldModel.FirstDayTotalSessionCount, oldModel.FirstDayTotalSessionDuration
}

func calculateMinDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (duration int64) {
	if oldModel.MinSessionDuration > modelResponse.MinSessionDuration {
		oldModel.MinSessionDuration = modelResponse.MinSessionDuration
		return oldModel.MinSessionDuration
	}
	return oldModel.MinSessionDuration
}

func calculateMaxDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (duration int64) {
	if modelResponse.MaxSessionDuration > oldModel.MaxSessionDuration {
		oldModel.MaxSessionDuration = modelResponse.MaxSessionDuration
		return oldModel.MaxSessionDuration
	}
	return oldModel.MaxSessionDuration
}

func calculateLastDaySessionCountAndDuration(modelResponse *model.GameSessionEveryLoginRespondModel, oldModel *model.GameSessionEveryLoginRespondModel) (count int64, duration int64) {
	if (oldModel.LastSessionYearOfDay == modelResponse.LastSessionYearOfDay) && (oldModel.FirstSessionYearOfDay != modelResponse.FirstSessionYearOfDay) {
		oldModel.LastDayTotalSessionCount = oldModel.LastDayTotalSessionCount + modelResponse.LastDayTotalSessionCount
		oldModel.LastDayTotalSessionDuration = oldModel.LastDayTotalSessionDuration + modelResponse.LastDayTotalSessionDuration
		return oldModel.LastDayTotalSessionCount, oldModel.LastDayTotalSessionDuration

	} else if (oldModel.LastSessionYearOfDay != modelResponse.LastSessionYearOfDay) && (oldModel.FirstSessionYearOfDay != modelResponse.FirstSessionYearOfDay) {
		return modelResponse.LastDayTotalSessionCount, modelResponse.LastDayTotalSessionDuration
	}

	return oldModel.LastDayTotalSessionCount, oldModel.LastDayTotalSessionDuration
}

func determineGameSessionDay(modelResponse *model.GameSessionEveryLoginRespondModel, day int64) {
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

func determineGameSessionHour(modelResponse *model.GameSessionEveryLoginRespondModel, hour int64) {
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

func determineGameSessionAmPm(modelResponse *model.GameSessionEveryLoginRespondModel, hour int64) {
	switch {
	case hour <= 12:
		modelResponse.AmSessionCount = 1
	default:
		modelResponse.PmSessionCount = 1
	}
}
