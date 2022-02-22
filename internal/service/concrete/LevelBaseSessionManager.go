package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	ILevelBaseSessionDal "FilterWorkerService/internal/repository/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"fmt"
	"log"

	logger "github.com/appneuroncompany/light-logger"
	"github.com/appneuroncompany/light-logger/clogger"
)

type levelBaseSessionManager struct {
	ILevelBaseSessionDal *ILevelBaseSessionDal.ILevelBaseSessionDal
	IJsonParser          *IJsonParser.IJsonParser
}

func LevelBaseSessionManagerConstructor() *levelBaseSessionManager {
	return &levelBaseSessionManager{
		ILevelBaseSessionDal: &IoC.LevelBaseSessionDal,
		IJsonParser:          &IoC.JsonParser,
	}
}

func (l *levelBaseSessionManager) ConvertRawModelToResponseModel(data *[]byte) (v interface{}, s bool, m string) {
	firstModel := model.LevelBaseSessionModel{}
	convertErr := (*l.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		clogger.Error(&logger.Messages{
			"Byte array to LevelBaseSessionModel LevelBaseSessionManager Json Parser Decode Err: ": convertErr.Error(),
		})
		// log.Fatal("LevelBaseSessionManager", "ConvertRawModelToResponseModel",
		// 	"byte array to LevelBaseSessionDataModel", "Json Parser Decode Err: ", convertErr.Error())
		return &model.LevelBaseSessionResponseModel{}, false, convertErr.Error()
	}
	hour := int16(firstModel.SessionFinishTime.Hour())
	yearOfDay := int16(firstModel.SessionFinishTime.YearDay())
	year := int16(firstModel.SessionFinishTime.Year())
	weekDay := int16(firstModel.SessionFinishTime.Weekday())
	minute := int16(firstModel.SessionFinishTime.Minute())
	modelResponse := model.LevelBaseSessionResponseModel{}
	modelResponse.Id = firstModel.Id
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.TotalLevelBaseSessionMinute = 1
	modelResponse.TotalLevelBaseSessionCount = 1
	modelResponse.FirstLevelSessionLevelIndex = int16(firstModel.LevelIndex)
	modelResponse.FirstLevelSessionDuration = int16(firstModel.SessionTimeMinute)
	modelResponse.FirstLevelSessionYearOfDay = yearOfDay
	modelResponse.FirstLevelSessionYear = year
	modelResponse.FirstLevelSessionWeekDay = weekDay
	modelResponse.FirstLevelSessionHour = hour
	modelResponse.FirstLevelSessionMinute = minute
	modelResponse.SecondLevelSessionLevelIndex = 0
	modelResponse.SecondLevelSessionDuration = 0
	modelResponse.ThirdLevelSessionLevelIndex = 0
	modelResponse.ThirdLevelSessionDuration = 0
	modelResponse.FourLevelSessionLevelIndex = 0
	modelResponse.FourLevelSessionDuration = 0
	modelResponse.FiveLevelSessionLevelIndex = 0
	modelResponse.FiveLevelSessionDuration = 0
	modelResponse.SixLevelSessionLevelIndex = 0
	modelResponse.SixLevelSessionDuration = 0
	modelResponse.SevenLevelSessionLevelIndex = 0
	modelResponse.SevenLevelSessionDuration = 0
	modelResponse.FirstFiveMinutesTotalLevelBaseSessionCount = 1
	modelResponse.FirstTenMinutesTotalLevelBaseSessionCount = 1
	modelResponse.FirstQuarterHourTotalLevelBaseSessionCount = 1
	modelResponse.FirstHalfHourTotalLevelBaseSessionCount = 1
	modelResponse.FirstHourTotalLevelBaseSessionCount = 1
	modelResponse.FirstTwoHourTotalLevelBaseSessionCount = 1
	modelResponse.FirstThreeHourTotalLevelBaseSessionCount = 1
	modelResponse.FirstSixHourTotalLevelBaseSessionCount = 1
	modelResponse.FirstTwelveHourTotalLevelBaseSessionCount = 1
	modelResponse.FirstDayTotalLevelBaseSessionCount = 1
	modelResponse.PenultimateLevelSessionLevelIndex = 0
	modelResponse.PenultimateLevelSessionLevelDuration = 0
	modelResponse.LastLevelSessionLevelIndex = 0
	modelResponse.LastLevelSessionLevelDuration = 0
	modelResponse.LastLevelSessionYearOfDay = 0
	modelResponse.LastLevelSessionYear = 0
	modelResponse.LastLevelSessionWeekDay = 0
	modelResponse.LastLevelSessionHour = 0
	modelResponse.LastLevelSessionMinute = 0
	defer log.Print("LevelBaseSessionManager", "ConvertRawModelToResponseModel",
		modelResponse.ClientId, modelResponse.ProjectId)

	oldModel, err := (*l.ILevelBaseSessionDal).GetById(modelResponse.ClientId, modelResponse.ProjectId)
	if err != nil && err.Error() != "null data error" {
		log.Fatal("LevelBaseSessionManager", "ConvertRawModelToResponseModel",
			"LevelBaseSessionDal_GetLevelBaseSessionById", err.Error())
	}
	switch {

	case err != nil && err.Error() != "not found":
		clogger.Error(&logger.Messages{
			fmt.Sprintf("Get clientId: %d, projectId: %d level_base_session_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): err.Error(),
		})
		
	case err != nil && err.Error() == "not found":

		logErr := (*l.ILevelBaseSessionDal).Add(&modelResponse)
		if logErr != nil {
			clogger.Error(&logger.Messages{
				fmt.Sprintf("Add clientId: %d, projectId: %d level_base_session_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): logErr.Error(),
			})
			return &modelResponse, false, logErr.Error()
		}
		clogger.Info(&logger.Messages{
			fmt.Sprintf("Add clientId: %d, projectId: %d level_base_session_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return &modelResponse, true, "Added"

	case err == nil:
		updatedModel, updateResult, updateErr := l.UpdateLevelBaseSession(&modelResponse, oldModel)
		if updateErr != nil {
			clogger.Error(&logger.Messages{
				fmt.Sprintf("Update clientId: %d, projectId: %d level_base_session_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): updateErr.Error(),
		})
			return updatedModel, updateResult, updateErr.Error()
		}
		clogger.Info(&logger.Messages{
			fmt.Sprintf("Update clientId: %d, projectId: %d level_base_session_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return updatedModel, updateResult, "Updated"
	default:
		return nil, false, ""
	}
	return nil, false, ""
}

func (l *levelBaseSessionManager) UpdateLevelBaseSession(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel) (updatedModel *model.LevelBaseSessionResponseModel, s bool, m error) {
	oldModel.Id = modelResponse.Id
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.TotalLevelBaseSessionMinute = (((int32(modelResponse.FirstLevelSessionYearOfDay)+365*int32(modelResponse.FirstLevelSessionYear))*24+int32(modelResponse.FirstLevelSessionHour))*60 + int32(modelResponse.FirstLevelSessionMinute)) - (((int32(oldModel.FirstLevelSessionYearOfDay)+365*int32(oldModel.FirstLevelSessionYear))*24+int32(oldModel.FirstLevelSessionHour))*60 + int32(oldModel.FirstLevelSessionMinute))
	oldModel.TotalLevelBaseSessionCount = modelResponse.TotalLevelBaseSessionCount + oldModel.TotalLevelBaseSessionCount
	CalculateLevelIndexBaseSession(modelResponse, oldModel)
	CalculateLevelBaseSessionFirstFiveMinutes(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstTenMinutes(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstQuarterHour(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstHalfHour(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstHour(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstTwoHour(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstThreeHour(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstSixHour(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstTwelveHour(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)
	CalculateLevelBaseSessionFirstDay(modelResponse, oldModel, oldModel.TotalLevelBaseSessionMinute)

	oldModel.PenultimateLevelSessionLevelIndex = oldModel.LastLevelSessionLevelIndex
	oldModel.PenultimateLevelSessionLevelDuration = oldModel.LastLevelSessionLevelDuration
	oldModel.LastLevelSessionLevelIndex = modelResponse.FirstLevelSessionLevelIndex
	oldModel.LastLevelSessionLevelDuration = modelResponse.FirstLevelSessionDuration
	oldModel.LastLevelSessionYearOfDay = modelResponse.FirstLevelSessionYearOfDay
	oldModel.LastLevelSessionYear = modelResponse.FirstLevelSessionYear
	oldModel.LastLevelSessionWeekDay = modelResponse.FirstLevelSessionWeekDay
	oldModel.LastLevelSessionHour = modelResponse.FirstLevelSessionHour
	oldModel.LastLevelSessionMinute = modelResponse.FirstLevelSessionMinute

	// defer log.Print("LevelBaseSessionManager", "UpdateLevelBaseSession",
	// 	oldModel.ClientId, oldModel.ProjectId)
	logErr := (*l.ILevelBaseSessionDal).UpdateById(oldModel.ClientId,oldModel.ProjectId, oldModel)
	if logErr != nil {
		// log.Fatal("LevelBaseSessionManager", "UpdateLevelBaseSession",
		// 	"LevelBaseSessionDal_UpdateLevelBaseSessionById", logErr.Error())
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateLevelBaseSessionFirstFiveMinutes(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 5:
		oldModel.FirstFiveMinutesTotalLevelBaseSessionCount = oldModel.FirstFiveMinutesTotalLevelBaseSessionCount + modelResponse.FirstFiveMinutesTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstTenMinutes(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 10:
		oldModel.FirstTenMinutesTotalLevelBaseSessionCount = oldModel.FirstTenMinutesTotalLevelBaseSessionCount + modelResponse.FirstTenMinutesTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstQuarterHour(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 15:
		oldModel.FirstQuarterHourTotalLevelBaseSessionCount = oldModel.FirstQuarterHourTotalLevelBaseSessionCount + modelResponse.FirstQuarterHourTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstHalfHour(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 30:
		oldModel.FirstHalfHourTotalLevelBaseSessionCount = oldModel.FirstHalfHourTotalLevelBaseSessionCount + modelResponse.FirstHalfHourTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstHour(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 60:
		oldModel.FirstHourTotalLevelBaseSessionCount = oldModel.FirstHourTotalLevelBaseSessionCount + modelResponse.FirstHourTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstTwoHour(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 120:
		oldModel.FirstTwoHourTotalLevelBaseSessionCount = oldModel.FirstTwoHourTotalLevelBaseSessionCount + modelResponse.FirstTwoHourTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstThreeHour(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 180:
		oldModel.FirstThreeHourTotalLevelBaseSessionCount = oldModel.FirstThreeHourTotalLevelBaseSessionCount + modelResponse.FirstThreeHourTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstSixHour(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 360:
		oldModel.FirstSixHourTotalLevelBaseSessionCount = oldModel.FirstSixHourTotalLevelBaseSessionCount + modelResponse.FirstSixHourTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstTwelveHour(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 720:
		oldModel.FirstTwelveHourTotalLevelBaseSessionCount = oldModel.FirstTwelveHourTotalLevelBaseSessionCount + modelResponse.FirstTwelveHourTotalLevelBaseSessionCount
	}
}

func CalculateLevelBaseSessionFirstDay(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel, total_session_minute int32) {
	switch {
	case total_session_minute <= 1440:
		oldModel.FirstDayTotalLevelBaseSessionCount = oldModel.FirstDayTotalLevelBaseSessionCount + modelResponse.FirstDayTotalLevelBaseSessionCount
	}
}

func CalculateLevelIndexBaseSession(modelResponse *model.LevelBaseSessionResponseModel, oldModel *model.LevelBaseSessionResponseModel) {
	switch oldModel.TotalLevelBaseSessionCount {
	case 2:
		oldModel.SecondLevelSessionLevelIndex = modelResponse.FirstLevelSessionLevelIndex
		oldModel.SecondLevelSessionDuration = modelResponse.FirstLevelSessionDuration
	case 3:
		oldModel.ThirdLevelSessionLevelIndex = modelResponse.FirstLevelSessionLevelIndex
		oldModel.ThirdLevelSessionLevelIndex = modelResponse.FirstLevelSessionDuration
	case 4:
		oldModel.FourLevelSessionLevelIndex = modelResponse.FirstLevelSessionLevelIndex
		oldModel.FourLevelSessionDuration = modelResponse.FirstLevelSessionDuration
	case 5:
		oldModel.FiveLevelSessionLevelIndex = modelResponse.FirstLevelSessionLevelIndex
		oldModel.FiveLevelSessionDuration = modelResponse.FirstLevelSessionDuration
	case 6:
		oldModel.SixLevelSessionLevelIndex = modelResponse.FirstLevelSessionLevelIndex
		oldModel.SixLevelSessionDuration = modelResponse.FirstLevelSessionDuration
	case 7:
		oldModel.SevenLevelSessionLevelIndex = modelResponse.FirstLevelSessionLevelIndex
		oldModel.SevenLevelSessionDuration = modelResponse.FirstLevelSessionDuration
	}
}
