package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IGameSessionDal "FilterWorkerService/internal/repository/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type gameSessionManager struct {
	IGameSessionDal *IGameSessionDal.IGameSessionDal
	IJsonParser     *IJsonParser.IJsonParser
}

func GameSessionManagerConstructor() *gameSessionManager {
	return &gameSessionManager{
		IGameSessionDal: &IoC.GameSessionDal,
		IJsonParser:     &IoC.JsonParser,
	}
}

func (g *gameSessionManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.GameSessionModel{}
	convertErr := (*g.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		clogger.Error(&map[string]interface{}{"Byte array to GameSessionModel  GameSessionManager Json Parser Decode ERROR: ": convertErr.Error()})
		return false, convertErr.Error()
	}
	hour := int16(firstModel.SessionFinishTime.Hour())
	yearOfDay := int16(firstModel.SessionFinishTime.YearDay())
	year := int16(firstModel.SessionFinishTime.Year())
	weekDay := int16(firstModel.SessionFinishTime.Weekday())
	minute := int16(firstModel.SessionFinishTime.Minute())

	modelResponse := model.GameSessionResponseModel{}
	modelResponse.Id = firstModel.Id
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.FirstSessionYearOfDay = yearOfDay
	modelResponse.FirstSessionYear = year
	modelResponse.FirstSessionWeekDay = weekDay
	modelResponse.FirstSessionHour = hour
	modelResponse.FirstSessionDuration = int16(firstModel.SessionTime)
	modelResponse.FirstSessionMinute = minute
	modelResponse.SecondSessionHour = 0
	modelResponse.SecondSessionDuration = 0
	modelResponse.SecondSessionMinute = 0
	modelResponse.ThirdSessionHour = 0
	modelResponse.ThirdSessionDuration = 0
	modelResponse.ThirdSessionMinute = 0
	modelResponse.FourthSessionHour = 0
	modelResponse.FourthSessionDuration = 0
	modelResponse.FourthSessionMinute = 0
	modelResponse.FifthSessionHour = 0
	modelResponse.FifthSessionDuration = 0
	modelResponse.FifthSessionMinute = 0
	modelResponse.SixthSessionHour = 0
	modelResponse.SixthSessionDuration = 0
	modelResponse.SixthSessionMinute = 0
	modelResponse.SeventhSessionHour = 0
	modelResponse.SeventhSessionDuration = 0
	modelResponse.SeventhSessionMinute = 0
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
	modelResponse.FirstFiveMinutesTotalSessionCount = 1
	modelResponse.FirstFiveMinutesTotalSessionDuration = 1
	modelResponse.FirstTenMinutesTotalSessionCount = 1
	modelResponse.FirstTenMinutesTotalSessionDuration = 1
	modelResponse.FirstQuarterHourTotalSessionCount = 1
	modelResponse.FirstQuarterHourTotalSessionDuration = 1
	modelResponse.FirstHalfHourTotalSessionCount = 1
	modelResponse.FirstHalfHourTotalSessionDuration = int16(firstModel.SessionTime)
	modelResponse.FirstHourTotalSessionCount = 1
	modelResponse.FirstHourTotalSessionDuration = int16(firstModel.SessionTime)
	modelResponse.FirstTwoHourTotalSessionCount = 1
	modelResponse.FirstTwoHourTotalSessionDuration = int16(firstModel.SessionTime)
	modelResponse.FirstThreeHourTotalSessionCount = 1
	modelResponse.FirstThreeHourTotalSessionDuration = int16(firstModel.SessionTime)
	modelResponse.FirstSixHourTotalSessionCount = 1
	modelResponse.FirstSixHourTotalSessionDuration = int16(firstModel.SessionTime)
	modelResponse.FirstTwelveHourTotalSessionCount = 1
	modelResponse.FirstTwelveHourTotalSessionDuration = int16(firstModel.SessionTime)

	modelResponse.TotalSessionDay = 1
	modelResponse.TotalSessionDuration = int32(firstModel.SessionTime)
	modelResponse.TotalSessionCount = 1
	modelResponse.TotalSessionHour = 1
	modelResponse.TotalSessionMinute = int32(firstModel.SessionTime)

	modelResponse.FirstDayTotalSessionCount = 1
	modelResponse.FirstDayTotalSessionDuration = int16(firstModel.SessionTime)
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

	modelResponse.MinSessionDuration = int16(firstModel.SessionTime)
	modelResponse.MaxSessionDuration = int16(firstModel.SessionTime)
	modelResponse.DailyAvegareSessionCount = 1
	modelResponse.DailyAverageSessionDuration = float32(firstModel.SessionTime)
	modelResponse.SessionBasedAvegareSessionDuration = float32(firstModel.SessionTime)
	modelResponse.DailyAvegareSessionCountMinusFirstDaySessionCount = 0
	modelResponse.DailyAvegareSessionDurationMinusFirstDaySessionDuration = 0
	modelResponse.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = 0
	modelResponse.SessionBasedAvegareSessionDurationMinusLastSessionDuration = float32(firstModel.SessionTime)
	DetermineGameSessionDay(&modelResponse, weekDay)
	DetermineGameSessionHour(&modelResponse, hour)
	DetermineGameSessionAmPm(&modelResponse, hour)

	// defer log.Print("GameSessionEveryLoginManager", "ConvertRawModelToResponseModel",
	// 	modelResponse.ClientId, modelResponse.ProjectId)

	oldModel, err := (*g.IGameSessionDal).GetById(modelResponse.ClientId, modelResponse.ProjectId)
	// if err != nil && err.Error() != "null data error" {
	// 	log.Fatal("GameSessionEveryLoginManager", "ConvertRawModelToResponseModel",
	// 		"GameSessionEveryLoginDal_GetGameSessionEveryLoginById", err.Error())
	// }
	switch {

	case err != nil && err.Error() != "not found":
		clogger.Error(&map[string]interface{}{
			fmt.Sprintf("Get clientId: %d, projectId: %d game_session_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): err.Error(),
		})

	case err != nil && err.Error() == "not found":

		logErr := (*g.IGameSessionDal).Add(&modelResponse)
		if logErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Add clientId: %d, projectId: %d game_session_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): logErr.Error(),
			})

			// log.Fatal("GameSessionEveryLoginManager", "ConvertRawModelToResponseModel",
			// 	"GameSessionEveryLoginDal_Add", logErr.Error())
			return false, logErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Add clientId: %d, projectId: %d game_session_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return true, "Added"

	case err == nil:
		_, updateResult, updateErr := g.UpdateGameSession(&modelResponse, oldModel)
		if updateErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Update clientId: %d, projectId: %d game_session_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): updateErr.Error(),
			})
			return updateResult, updateErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Update clientId: %d, projectId: %d game_session_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return updateResult, "Updated"

	default:

		return false, ""

	}
	return false, ""
}

func (g *gameSessionManager) UpdateGameSession(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) (updatedModel *model.GameSessionResponseModel, s bool, m error) {
	oldModel.Id = modelResponse.Id
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.TotalSessionDay = (int32(modelResponse.FirstSessionYearOfDay) - int32(oldModel.FirstSessionYearOfDay)) + 365*(int32(modelResponse.FirstSessionYear)-int32(oldModel.FirstSessionYear))
	oldModel.TotalSessionDuration = oldModel.TotalSessionDuration + modelResponse.TotalSessionDuration
	oldModel.TotalSessionCount = oldModel.TotalSessionCount + modelResponse.TotalSessionCount

	oldModel.TotalSessionHour = ((int32(modelResponse.FirstSessionYearOfDay)+365*int32(modelResponse.FirstSessionYear))*24 + int32(modelResponse.FirstSessionHour)) - ((int32(oldModel.FirstSessionYearOfDay)+365*int32(oldModel.FirstSessionYear))*24 + int32(oldModel.FirstSessionHour))
	oldModel.TotalSessionMinute = (((int32(modelResponse.FirstSessionYearOfDay)+365*int32(modelResponse.FirstSessionYear))*24+int32(modelResponse.FirstSessionHour))*60 + int32(modelResponse.FirstSessionMinute)) - (((int32(oldModel.FirstSessionYearOfDay)+365*int32(oldModel.FirstSessionYear))*24+int32(oldModel.FirstSessionHour))*60 + int32(oldModel.FirstSessionMinute))
	CalculateSecondGameSession(modelResponse, oldModel)
	CalculateThirdGameSession(modelResponse, oldModel)
	CalculateFourthGameSession(modelResponse, oldModel)
	CalculateFifthGameSession(modelResponse, oldModel)
	CalculateSixthGameSession(modelResponse, oldModel)
	CalculateSixthGameSession(modelResponse, oldModel)

	CalculateFirstDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateSecondDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	CalculateThirdDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	CalculateFourthDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	CalculateFifthDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	CalculateSixthDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)
	CalculateSeventhDayTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionHour)

	CalculateFirstFiveMinutesTotalSessionCountAndSessionDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstTenMinutesTotalSessionCountAndSessionDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstQuarterHourTotalSessionCountAndSessionDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstHalfHourTotalSessionCountAndSessionDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstTwoHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstThreeHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstSixHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)
	CalculateFirstTwelveHourTotalSessionCountAndDuration(modelResponse, oldModel, oldModel.TotalSessionMinute)

	oldModel.PenultimateSessionDuration = oldModel.LastSessionDuration
	oldModel.PenultimateSessionHour = oldModel.LastSessionHour
	oldModel.PenultimateSessionMinute = oldModel.LastSessionMinute
	oldModel.LastSessionYearOfDay = modelResponse.FirstSessionYearOfDay
	oldModel.LastSessionYear = modelResponse.FirstSessionYear
	oldModel.LastSessionHour = modelResponse.FirstSessionHour
	oldModel.LastSessionDuration = modelResponse.FirstSessionDuration
	oldModel.LastDurationMinusPenultimateDuration = oldModel.LastSessionDuration - oldModel.PenultimateSessionDuration

	CalculateMinDuration(modelResponse, oldModel)
	CalculateMaxDuration(modelResponse, oldModel)
	oldModel.DailyAvegareSessionCount, oldModel.DailyAverageSessionDuration = CalculateDailyAvegareSessionCountAndDuration(oldModel)
	oldModel.SessionBasedAvegareSessionDuration = float32(oldModel.TotalSessionDuration) / float32(oldModel.TotalSessionCount)
	oldModel.DailyAvegareSessionCountMinusFirstDaySessionCount = oldModel.DailyAvegareSessionCount - float32(oldModel.FirstDayTotalSessionCount)
	oldModel.DailyAvegareSessionDurationMinusFirstDaySessionDuration = oldModel.DailyAverageSessionDuration - float32(oldModel.FirstDayTotalSessionDuration)
	oldModel.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = oldModel.SessionBasedAvegareSessionDuration - float32(oldModel.FirstSessionDuration)
	oldModel.SessionBasedAvegareSessionDurationMinusLastSessionDuration = oldModel.SessionBasedAvegareSessionDuration - float32(oldModel.LastSessionDuration)
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

	// defer log.Print("GameSessionEveryLoginManager", "UpdateLevelBaseSession",
	// 	oldModel.ClientId, oldModel.ProjectId)
	logErr := (*g.IGameSessionDal).UpdateById(oldModel.ClientId, oldModel.ProjectId, oldModel)
	if logErr != nil {

		// log.Fatal("GameSessionEveryLoginManager", "UpdateGameSession",
		// 	"GameSessionEveryLoginDal_UpdateGameSessionEveryLoginById", logErr.Error())
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateDailyAvegareSessionCountAndDuration(oldModel *model.GameSessionResponseModel) (count float32, duration float32) {
	if oldModel.TotalSessionDay == 0 {
		return float32(oldModel.TotalSessionCount), float32(oldModel.TotalSessionDuration)
	}
	return float32(oldModel.TotalSessionCount) / float32(oldModel.TotalSessionDay), float32(oldModel.TotalSessionDuration) / float32(oldModel.TotalSessionDay)
}

func CalculateFirstFiveMinutesTotalSessionCountAndSessionDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 5:
		oldModel.FirstFiveMinutesTotalSessionCount = oldModel.FirstFiveMinutesTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstFiveMinutesTotalSessionDuration = oldModel.FirstFiveMinutesTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstTenMinutesTotalSessionCountAndSessionDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 10:
		oldModel.FirstTenMinutesTotalSessionCount = oldModel.FirstTenMinutesTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstTenMinutesTotalSessionDuration = oldModel.FirstTenMinutesTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstQuarterHourTotalSessionCountAndSessionDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 15:
		oldModel.FirstQuarterHourTotalSessionCount = oldModel.FirstQuarterHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstQuarterHourTotalSessionDuration = oldModel.FirstQuarterHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstHalfHourTotalSessionCountAndSessionDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 30:
		oldModel.FirstHalfHourTotalSessionCount = oldModel.FirstHalfHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstHalfHourTotalSessionDuration = oldModel.FirstHalfHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstHourTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 60:
		oldModel.FirstHourTotalSessionCount = oldModel.FirstHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstHourTotalSessionDuration = oldModel.FirstHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstTwoHourTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 120:
		oldModel.FirstTwoHourTotalSessionCount = oldModel.FirstTwoHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstTwoHourTotalSessionDuration = oldModel.FirstTwoHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstThreeHourTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 180:
		oldModel.FirstThreeHourTotalSessionCount = oldModel.FirstThreeHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstThreeHourTotalSessionDuration = oldModel.FirstThreeHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstSixHourTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 360:
		oldModel.FirstSixHourTotalSessionCount = oldModel.FirstSixHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstSixHourTotalSessionDuration = oldModel.FirstSixHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstTwelveHourTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 720:
		oldModel.FirstTwelveHourTotalSessionCount = oldModel.FirstTwelveHourTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstTwelveHourTotalSessionDuration = oldModel.FirstTwelveHourTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFirstDayTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 1440:
		oldModel.FirstDayTotalSessionCount = oldModel.FirstDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FirstDayTotalSessionDuration = oldModel.FirstDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateSecondDayTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_hour int32) {
	switch {
	case total_session_hour <= 48 && total_session_hour > 24:
		oldModel.SecondDayTotalSessionCount = oldModel.SecondDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.SecondDayTotalSessionDuration = oldModel.SecondDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateThirdDayTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_hour int32) {
	switch {
	case total_session_hour <= 72 && total_session_hour > 48:
		oldModel.ThirdDayTotalSessionCount = oldModel.ThirdDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.ThirdDayTotalSessionDuration = oldModel.ThirdDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFourthDayTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_hour int32) {
	switch {
	case total_session_hour <= 96 && total_session_hour > 72:
		oldModel.FourthDayTotalSessionCount = oldModel.FourthDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FourthDayTotalSessionDuration = oldModel.FourthDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateFifthDayTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_hour int32) {
	switch {
	case total_session_hour <= 120 && total_session_hour > 96:
		oldModel.FifthDayTotalSessionCount = oldModel.FifthDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.FifthDayTotalSessionDuration = oldModel.FifthDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateSixthDayTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_hour int32) {
	switch {
	case total_session_hour <= 144 && total_session_hour > 120:
		oldModel.SixthDayTotalSessionCount = oldModel.SixthDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.SixthDayTotalSessionDuration = oldModel.SixthDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateSeventhDayTotalSessionCountAndDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel, total_session_hour int32) {
	switch {
	case total_session_hour <= 168 && total_session_hour > 144:
		oldModel.SeventhDayTotalSessionCount = oldModel.SeventhDayTotalSessionCount + modelResponse.FirstDayTotalSessionCount
		oldModel.SeventhDayTotalSessionDuration = oldModel.SeventhDayTotalSessionDuration + modelResponse.FirstDayTotalSessionDuration
	}
}

func CalculateSecondGameSession(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch oldModel.TotalSessionCount {
	case 2:
		oldModel.SecondSessionHour = modelResponse.FirstSessionHour
		oldModel.SecondSessionDuration = modelResponse.FirstSessionDuration
		oldModel.SecondSessionMinute = modelResponse.FirstSessionMinute
	}
}

func CalculateThirdGameSession(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch oldModel.TotalSessionCount {
	case 3:
		oldModel.ThirdSessionHour = modelResponse.FirstSessionHour
		oldModel.ThirdSessionDuration = modelResponse.FirstSessionDuration
		oldModel.ThirdSessionMinute = modelResponse.FirstSessionMinute
	}
}

func CalculateFourthGameSession(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch oldModel.TotalSessionCount {
	case 4:
		oldModel.FourthSessionHour = modelResponse.FirstSessionHour
		oldModel.FourthSessionDuration = modelResponse.FirstSessionDuration
		oldModel.FourthSessionMinute = modelResponse.FirstSessionMinute
	}
}

func CalculateFifthGameSession(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch oldModel.TotalSessionCount {
	case 5:
		oldModel.FifthSessionHour = modelResponse.FirstSessionHour
		oldModel.FifthSessionDuration = modelResponse.FirstSessionDuration
		oldModel.FifthSessionMinute = modelResponse.FirstSessionMinute
	}
}

func CalculateSixthGameSession(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch oldModel.TotalSessionCount {
	case 6:
		oldModel.SixthSessionHour = modelResponse.FirstSessionHour
		oldModel.SixthSessionDuration = modelResponse.FirstSessionDuration
		oldModel.SixthSessionMinute = modelResponse.FirstSessionMinute
	}
}

func CalculateSeventhGameSession(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch oldModel.TotalSessionCount {
	case 7:
		oldModel.SeventhSessionHour = modelResponse.FirstSessionHour
		oldModel.SeventhSessionDuration = modelResponse.FirstSessionDuration
		oldModel.SeventhSessionMinute = modelResponse.FirstSessionMinute
	}
}

func CalculateMinDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch {
	case oldModel.MinSessionDuration > modelResponse.MinSessionDuration:
		oldModel.MinSessionDuration = modelResponse.MinSessionDuration
	}
}

func CalculateMaxDuration(modelResponse *model.GameSessionResponseModel, oldModel *model.GameSessionResponseModel) {
	switch {
	case modelResponse.MaxSessionDuration > oldModel.MaxSessionDuration:
		oldModel.MaxSessionDuration = modelResponse.MaxSessionDuration
	}
}

func DetermineGameSessionDay(modelResponse *model.GameSessionResponseModel, day int16) {
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

func DetermineGameSessionHour(modelResponse *model.GameSessionResponseModel, hour int16) {
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

func DetermineGameSessionAmPm(modelResponse *model.GameSessionResponseModel, hour int16) {
	switch {
	case hour <= 12:
		modelResponse.AmSessionCount = 1
	default:
		modelResponse.PmSessionCount = 1
	}
}
