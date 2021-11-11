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

func (g *GameSessionEveryLoginManager) ConvertRawModelToResponseModel(data *[]byte) (respondModel *model.GameSessionEveryLoginRespondModel, s bool, m string) {
	firstModel := model.GameSessionEveryLoginModel{}
	err := g.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return &model.GameSessionEveryLoginRespondModel{}, false, err.Error()
	}
	hour := int64(firstModel.SessionFinishTime.Hour())
	yearOfDay := int64(firstModel.SessionFinishTime.YearDay())
	year := int64(firstModel.SessionFinishTime.Year())
	weekDay := int64(firstModel.SessionFinishTime.Weekday())
	modelResponse := model.GameSessionEveryLoginRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.FirstSessionYearOfDay = yearOfDay
	modelResponse.FirstSessionYear = year
	modelResponse.FirstSessionWeekDay = weekDay
	modelResponse.FirstSessionHour = hour
	modelResponse.FirstSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.PenultimateSessionHour = hour
	modelResponse.PenultimateSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.LastSessionYearOfDay = yearOfDay
	modelResponse.LastSessionYear = year
	modelResponse.LastSessionHour = hour
	modelResponse.LastSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.LastHourMinusPenultimateHour = modelResponse.LastSessionHour - modelResponse.PenultimateSessionHour
	modelResponse.LastDurationMinusPenultimateDuration = int64(modelResponse.LastSessionDuration - modelResponse.PenultimateSessionDuration)
	modelResponse.TotalSessionDay = 1
	modelResponse.TotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.TotalSessionCount = 1
	modelResponse.FirstDaySessionCount = 1
	modelResponse.FirstDayTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.PenultimateDayTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.PenultimateDayTotalSessionCount = 1
	modelResponse.LastDayTotalSessionCount = 1
	modelResponse.LastDayTotalSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.MinSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.MaxSessionDuration = int64(firstModel.SessionTimeMinute)
	modelResponse.DailyAvegareSessionCount = float64(modelResponse.TotalSessionCount) / float64(modelResponse.TotalSessionDay)
	modelResponse.DailyAverageSessionDuration = float64(modelResponse.TotalSessionDuration) / float64(modelResponse.TotalSessionDay)
	modelResponse.SessionBasedAvegareSessionDuration = float64(modelResponse.TotalSessionDuration) / float64(modelResponse.TotalSessionCount)
	modelResponse.DailyAvegareSessionCountMinusFirstDaySessionCount = modelResponse.DailyAvegareSessionCount - float64(modelResponse.FirstDaySessionCount)
	modelResponse.DailyAvegareSessionDurationMinusFirstDaySessionDuration = modelResponse.DailyAverageSessionDuration - float64(modelResponse.FirstDayTotalSessionDuration)
	modelResponse.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = modelResponse.SessionBasedAvegareSessionDuration - float64(modelResponse.FirstSessionDuration)
	modelResponse.SessionBasedAvegareSessionDurationMinusLastSessionDuration = modelResponse.SessionBasedAvegareSessionDuration - float64(modelResponse.LastSessionDuration)
	determineGameSessionDay(&modelResponse, weekDay)
	determineGameSessionHour(&modelResponse, hour)
	determineGameSessionAmPm(&modelResponse, hour)
	return &modelResponse, true, ""
}

func (g *GameSessionEveryLoginManager) AddGameSession(data *model.GameSessionEveryLoginRespondModel) (s bool, m string) {
	logErr := g.IGameSessionEveryLoginDal.Add(data)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func (g *GameSessionEveryLoginManager) UpdateGameSession(modelResponse *model.GameSessionEveryLoginRespondModel) (s bool, m string) {
	oldModel, err := g.IGameSessionEveryLoginDal.GetGameSessionEveryLoginByCustomerId(modelResponse.CustomerId)
	if err != nil {
		return false, err.Error()
	}
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	//oldModel.FirstSessionYearOfDay
	//oldModel.FirstSessionYear
	//oldModel.FirstSessionWeekDay
	//oldModel.FirstSessionHour
	//oldModel.FirstSessionDuration
	oldModel.PenultimateSessionDuration = oldModel.LastSessionDuration
	oldModel.PenultimateSessionHour = oldModel.LastSessionHour
	oldModel.LastSessionYearOfDay = modelResponse.LastSessionYearOfDay
	oldModel.LastSessionYear = modelResponse.LastSessionYear
	oldModel.LastSessionHour = modelResponse.LastSessionHour
	oldModel.LastSessionDuration = modelResponse.LastSessionDuration
	oldModel.LastHourMinusPenultimateHour = oldModel.LastSessionHour - oldModel.PenultimateSessionHour
	oldModel.LastDurationMinusPenultimateDuration = oldModel.LastSessionDuration - oldModel.PenultimateSessionDuration
	oldModel.TotalSessionDay = oldModel.LastSessionYearOfDay - oldModel.FirstSessionYearOfDay + 365*(oldModel.LastSessionYear-oldModel.FirstSessionYear)
	oldModel.TotalSessionDuration = oldModel.TotalSessionDuration + modelResponse.TotalSessionDuration
	oldModel.TotalSessionCount = oldModel.TotalSessionCount + modelResponse.TotalSessionCount
	oldModel.FirstDaySessionCount, oldModel.FirstDayTotalSessionDuration = calculateFirstDayGameSessionCountAndDuration(modelResponse, oldModel)
	oldModel.PenultimateDayTotalSessionCount, oldModel.PenultimateDayTotalSessionDuration = calculatePenultimateDay(modelResponse, oldModel)
	oldModel.LastDayTotalSessionCount, oldModel.LastDayTotalSessionDuration = calculateLastDaySessionCountAndDuration(modelResponse, oldModel)
	oldModel.MinSessionDuration = calculateMinDuration(modelResponse, oldModel)
	oldModel.MaxSessionDuration = calculateMaxDuration(modelResponse, oldModel)
	oldModel.DailyAvegareSessionCount = float64(oldModel.TotalSessionCount) / float64(oldModel.TotalSessionDay)
	oldModel.DailyAverageSessionDuration = float64(oldModel.TotalSessionDuration) / float64(oldModel.TotalSessionDay)
	oldModel.SessionBasedAvegareSessionDuration = float64(oldModel.TotalSessionDuration) / float64(oldModel.TotalSessionCount)
	oldModel.DailyAvegareSessionCountMinusFirstDaySessionCount = oldModel.DailyAvegareSessionCount - float64(oldModel.FirstDaySessionCount)
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

	logErr := g.IGameSessionEveryLoginDal.UpdateGameSessionEveryLoginByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
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
		oldModel.FirstDaySessionCount = oldModel.FirstDaySessionCount + modelResponse.FirstDaySessionCount
		oldModel.FirstDayTotalSessionDuration = oldModel.FirstDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
		return oldModel.FirstDaySessionCount, oldModel.FirstDayTotalSessionDuration
	}
	return oldModel.FirstDaySessionCount, oldModel.FirstDayTotalSessionDuration
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

	} else if (oldModel.LastSessionYearOfDay != modelResponse.LastSessionYearOfDay) && (oldModel.FirstSessionYearOfDay == modelResponse.FirstSessionYearOfDay) {
		return oldModel.LastDayTotalSessionCount, oldModel.LastDayTotalSessionDuration

	} else {
		return oldModel.LastDayTotalSessionCount, oldModel.LastDayTotalSessionDuration
	}
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
