package concrete

import (
	"FilterWorkerService/internal/IoC"
	model "FilterWorkerService/internal/model"
	IScreenSwipeDal "FilterWorkerService/internal/repository/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
	"fmt"

	"github.com/appneuroncompany/light-logger/clogger"
)

type screenSwipeManager struct {
	IScreenSwipeDal *IScreenSwipeDal.IScreenSwipeDal
	IJsonParser     *IJsonParser.IJsonParser
}

func ScreenSwipeManagerConstructor() *screenSwipeManager {
	return &screenSwipeManager{
		IScreenSwipeDal: &IoC.ScreenSwipeDal,
		IJsonParser:     &IoC.JsonParser,
	}
}

func (sc *screenSwipeManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.ScreenSwipeModel{}
	convertErr := (*sc.IJsonParser).DecodeJson(data, &firstModel)
	if convertErr != nil {
		clogger.Error(&map[string]interface{}{"Byte array to ScreenSwipeModel  ScreenSwipeManager Json Parser Decode ERROR: ": convertErr.Error()})
		return false, convertErr.Error()
	}
	hour := int16(firstModel.CreatedAt.Hour())
	yearOfDay := int16(firstModel.CreatedAt.YearDay())
	year := int16(firstModel.CreatedAt.Year())
	weekDay := int16(firstModel.CreatedAt.Weekday())
	minute := int16(firstModel.CreatedAt.Minute())
	swipeDirection := byte(firstModel.SwipeDirection)
	modelResponse := model.ScreenSwipeResponseModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int16(firstModel.LevelIndex)
	modelResponse.TotalSwipeSessionCount = 1
	modelResponse.TotalSwipeHour = 0
	modelResponse.FirstSwipeYearOfDay = yearOfDay
	modelResponse.FirstSwipeYear = year
	modelResponse.FirstSwipeHour = hour
	modelResponse.FirstSwipeWeekDay = weekDay
	modelResponse.FirstSwipeMinute = minute
	modelResponse.FistSwipeDirection = byte(firstModel.SwipeDirection)
	modelResponse.FirstSwipeStartXCor = firstModel.StartLocX
	modelResponse.FirstSwipeStartYCor = firstModel.StartLocY
	modelResponse.FirstSwipeFinishXCor = firstModel.FinishLocX
	modelResponse.FirstSwipeFinishYCor = firstModel.FinishLocY
	modelResponse.SecondSwipeDirection = 0
	modelResponse.SecondSwipeStartXCor = 0
	modelResponse.SecondSwipeStartYCor = 0
	modelResponse.SecondSwipeFinishXCor = 0
	modelResponse.SecondSwipeFinishYCor = 0
	modelResponse.ThirdSwipeDirection = 0
	modelResponse.ThirdSwipeStartXCor = 0
	modelResponse.ThirdSwipeStartYCor = 0
	modelResponse.ThirdSwipeFinishXCor = 0
	modelResponse.ThirdSwipeFinishYCor = 0

	modelResponse.FourthSwipeDirection = 0
	modelResponse.FourthSwipeStartXCor = 0
	modelResponse.FourthSwipeStartYCor = 0
	modelResponse.FourthSwipeFinishXCor = 0
	modelResponse.FourthSwipeFinishYCor = 0
	modelResponse.FifthSwipeDirection = 0
	modelResponse.FifthSwipeStartXCor = 0
	modelResponse.FifthSwipeStartYCor = 0
	modelResponse.FifthSwipeFinishXCor = 0
	modelResponse.FifthSwipeFinishYCor = 0

	modelResponse.SixthSwipeDirection = 0
	modelResponse.SixthSwipeStartXCor = 0
	modelResponse.SixthSwipeStartYCor = 0
	modelResponse.SixthSwipeFinishXCor = 0
	modelResponse.SixthSwipeFinishYCor = 0
	modelResponse.SeventhSwipeDirection = 0
	modelResponse.SeventhSwipeStartXCor = 0
	modelResponse.SeventhSwipeStartYCor = 0
	modelResponse.SeventhSwipeFinishXCor = 0
	modelResponse.SeventhSwipeFinishYCor = 0

	modelResponse.PenultimateSwipeDirection = 0
	modelResponse.PenultimateSwipeStartXCor = 0
	modelResponse.PenultimateSwipeStartYCor = 0
	modelResponse.PenultimateSwipeFinishXCor = 0
	modelResponse.PenultimateSwipeFinishYCor = 0

	modelResponse.PenultimateSwipeYearOfDay = 0
	modelResponse.PenultimateSwipeYear = 0
	modelResponse.PenultimateSwipeHour = 0
	modelResponse.PenultimateSwipeWeekDay = 0
	modelResponse.PenultimateSwipeMinute = 0

	modelResponse.LastSwipeDirection = 0
	modelResponse.LastSwipeStartXCor = 0
	modelResponse.LastSwipeStartYCor = 0
	modelResponse.LastSwipeFinishXCor = 0
	modelResponse.LastSwipeFinishYCor = 0

	modelResponse.LastSwipeYearOfDay = 0
	modelResponse.LastSwipeYear = 0
	modelResponse.LastSwipeHour = 0
	modelResponse.LastSwipeWeekDay = 0
	modelResponse.LastSwipeMinute = 0

	DetermineSwipeDirection(&modelResponse, swipeDirection)

	modelResponse.FirstDaySwipeTotalStartXCor = firstModel.StartLocX
	modelResponse.FirstDaySwipeTotalStartYCor = firstModel.StartLocY
	modelResponse.FirstDaySwipeTotalFinishXCor = firstModel.FinishLocX
	modelResponse.FirstDaySwipeTotalFinishYCor = firstModel.FinishLocY

	modelResponse.SecondDayTotalSwipeUpCount = 0
	modelResponse.SecondDayTotalSwipeDownCount = 0
	modelResponse.SecondDayTotalSwipeRightCount = 0
	modelResponse.SecondDayTotalSwipeLeftCount = 0
	modelResponse.SecondDaySwipeTotalStartXCor = 0
	modelResponse.SecondDaySwipeTotalStartYCor = 0
	modelResponse.SecondDaySwipeTotalFinishXCor = 0
	modelResponse.SecondDaySwipeTotalFinishYCor = 0

	modelResponse.ThirdDayTotalSwipeUpCount = 0
	modelResponse.ThirdDayTotalSwipeDownCount = 0
	modelResponse.ThirdDayTotalSwipeRightCount = 0
	modelResponse.ThirdDayTotalSwipeLeftCount = 0
	modelResponse.ThirdDaySwipeTotalStartXCor = 0
	modelResponse.ThirdDaySwipeTotalStartYCor = 0
	modelResponse.ThirdDaySwipeTotalFinishXCor = 0
	modelResponse.ThirdDaySwipeTotalFinishYCor = 0

	modelResponse.FourthDayTotalSwipeUpCount = 0
	modelResponse.FourthDayTotalSwipeDownCount = 0
	modelResponse.FourthDayTotalSwipeRightCount = 0
	modelResponse.FourthDayTotalSwipeLeftCount = 0
	modelResponse.FourthDaySwipeTotalStartXCor = 0
	modelResponse.FourthDaySwipeTotalStartYCor = 0
	modelResponse.FourthDaySwipeTotalFinishXCor = 0
	modelResponse.FourthDaySwipeTotalFinishYCor = 0

	modelResponse.FifthDayTotalSwipeUpCount = 0
	modelResponse.FifthDayTotalSwipeDownCount = 0
	modelResponse.FifthDayTotalSwipeRightCount = 0
	modelResponse.FifthDayTotalSwipeLeftCount = 0
	modelResponse.FifthDaySwipeTotalStartXCor = 0
	modelResponse.FifthDaySwipeTotalStartYCor = 0
	modelResponse.FifthDaySwipeTotalFinishXCor = 0
	modelResponse.FifthDaySwipeTotalFinishYCor = 0

	modelResponse.SixthDayTotalSwipeUpCount = 0
	modelResponse.SixthDayTotalSwipeDownCount = 0
	modelResponse.SixthDayTotalSwipeRightCount = 0
	modelResponse.SixthDayTotalSwipeLeftCount = 0
	modelResponse.SixthDaySwipeTotalStartXCor = 0
	modelResponse.SixthDaySwipeTotalStartYCor = 0
	modelResponse.SixthDaySwipeTotalFinishXCor = 0
	modelResponse.SixthDaySwipeTotalFinishYCor = 0

	modelResponse.SeventhDayTotalSwipeUpCount = 0
	modelResponse.SeventhDayTotalSwipeDownCount = 0
	modelResponse.SeventhDayTotalSwipeRightCount = 0
	modelResponse.SeventhDayTotalSwipeLeftCount = 0
	modelResponse.SeventhDaySwipeTotalStartXCor = 0
	modelResponse.SeventhDaySwipeTotalStartYCor = 0
	modelResponse.SeventhDaySwipeTotalFinishXCor = 0
	modelResponse.SeventhDaySwipeTotalFinishYCor = 0

	modelResponse.TotalSwipeStartXCor = firstModel.StartLocX
	modelResponse.TotalSwipeStartYCor = firstModel.StartLocY
	modelResponse.TotalSwipeFinishXCor = firstModel.FinishLocX
	modelResponse.TotalSwipeFinishYCor = firstModel.FinishLocY

	// defer log.Print("ScreenSwipeManager", "ConvertRawModelToResponseModel",
	// 	modelResponse.ClientId, modelResponse.ProjectId)
	oldModel, err := (*sc.IScreenSwipeDal).GetById(modelResponse.ClientId, modelResponse.ProjectId)
	// if err != nil && err.Error() != "null data error" {
	// 	log.Fatal("ScreenSwipeManager", "ConvertRawModelToResponseModel",
	// 		"ScreenSwipeDal_GetScreenSwipeById", err.Error())
	// }
	switch {

	case err != nil && err.Error() != "not found":
		clogger.Error(&map[string]interface{}{
			fmt.Sprintf("Get clientId: %d, projectId: %d screen_swipe_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): err.Error(),
		})

	case err != nil && err.Error() == "not found":

		logErr := (*sc.IScreenSwipeDal).Add(&modelResponse)
		if logErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Add clientId: %d, projectId: %d screen_swipe_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): logErr.Error(),
			})
			return false, logErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Add clientId: %d, projectId: %d screen_swipe_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return true, "Added"

	case err == nil:
		_, updateResult, updateErr := sc.UpdateScreenSwipe(&modelResponse, oldModel)
		if updateErr != nil {
			clogger.Error(&map[string]interface{}{
				fmt.Sprintf("Update clientId: %d, projectId: %d screen_swipe_data ERROR: ", modelResponse.ClientId, modelResponse.ProjectId): updateErr.Error(),
			})
			return updateResult, updateErr.Error()
		}
		clogger.Info(&map[string]interface{}{
			fmt.Sprintf("Update clientId: %d, projectId: %d screen_swipe_data  : ", modelResponse.ClientId, modelResponse.ProjectId): "SUCCESS",
		})
		return updateResult, "Updated"

	default:

		return false, ""

	}
	return false, ""
}

func (sc *screenSwipeManager) UpdateScreenSwipe(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel) (updatedModel *model.ScreenSwipeResponseModel, s bool, m error) {

	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalSwipeSessionCount = modelResponse.TotalSwipeSessionCount + oldModel.TotalSwipeSessionCount
	oldModel.TotalSwipeHour = ((int32(modelResponse.FirstSwipeYearOfDay)+365*int32(modelResponse.FirstSwipeYear))*24 + int32(modelResponse.FirstSwipeHour)) - ((int32(oldModel.FirstSwipeYearOfDay)+365*int32(oldModel.FirstSwipeYear))*24 + int32(oldModel.FirstSwipeHour))

	CalculateSwipeNumber(modelResponse, oldModel)
	oldModel.PenultimateSwipeDirection = oldModel.LastSwipeDirection
	oldModel.PenultimateSwipeStartXCor = oldModel.LastSwipeStartXCor
	oldModel.PenultimateSwipeStartYCor = oldModel.LastSwipeStartYCor
	oldModel.PenultimateSwipeFinishXCor = oldModel.LastSwipeFinishXCor
	oldModel.PenultimateSwipeFinishYCor = oldModel.LastSwipeFinishYCor

	oldModel.PenultimateSwipeYearOfDay = oldModel.LastSwipeYearOfDay
	oldModel.PenultimateSwipeYear = oldModel.LastSwipeYear
	oldModel.PenultimateSwipeHour = oldModel.LastSwipeHour
	oldModel.PenultimateSwipeWeekDay = oldModel.LastSwipeWeekDay
	oldModel.PenultimateSwipeMinute = oldModel.LastSwipeMinute

	oldModel.LastSwipeDirection = modelResponse.FistSwipeDirection
	oldModel.LastSwipeStartXCor = modelResponse.FirstSwipeStartXCor
	oldModel.LastSwipeStartYCor = modelResponse.FirstSwipeStartYCor
	oldModel.LastSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
	oldModel.LastSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor

	oldModel.LastSwipeYearOfDay = modelResponse.FirstSwipeYearOfDay
	oldModel.LastSwipeYear = modelResponse.FirstSwipeYear
	oldModel.LastSwipeHour = modelResponse.FirstSwipeHour
	oldModel.LastSwipeWeekDay = modelResponse.FirstSwipeWeekDay
	oldModel.LastSwipeMinute = modelResponse.FirstSwipeMinute
	CalculateSwipeFirstDay(modelResponse, oldModel, oldModel.TotalSwipeHour)
	CalculateSwipeSecondDay(modelResponse, oldModel, oldModel.TotalSwipeHour)
	CalculateSwipeThirdDay(modelResponse, oldModel, oldModel.TotalSwipeHour)
	CalculateSwipeFourthDay(modelResponse, oldModel, oldModel.TotalSwipeHour)
	CalculateSwipeFifthDay(modelResponse, oldModel, oldModel.TotalSwipeHour)
	CalculateSwipeSixthDay(modelResponse, oldModel, oldModel.TotalSwipeHour)
	CalculateSwipeSeventhDay(modelResponse, oldModel, oldModel.TotalSwipeHour)
	oldModel.TotalSwipeUpCount = modelResponse.TotalSwipeUpCount + oldModel.TotalSwipeUpCount
	oldModel.TotalSwipeDownCount = modelResponse.TotalSwipeDownCount + oldModel.TotalSwipeDownCount
	oldModel.TotalSwipeRightCount = modelResponse.TotalSwipeRightCount + oldModel.TotalSwipeRightCount
	oldModel.TotalSwipeLeftCount = modelResponse.TotalSwipeLeftCount + oldModel.TotalSwipeLeftCount
	oldModel.TotalSwipeStartXCor = modelResponse.TotalSwipeStartXCor + oldModel.TotalSwipeStartXCor
	oldModel.TotalSwipeStartYCor = modelResponse.TotalSwipeStartYCor + oldModel.TotalSwipeStartYCor
	oldModel.TotalSwipeFinishXCor = modelResponse.TotalSwipeFinishXCor + oldModel.TotalSwipeFinishXCor
	oldModel.TotalSwipeFinishYCor = modelResponse.TotalSwipeFinishYCor + oldModel.TotalSwipeFinishYCor

	// defer log.Print("ScreenSwipeManager", "UpdateScreenSwipe",
	// 	oldModel.ClientId, oldModel.ProjectId)
	logErr := (*sc.IScreenSwipeDal).UpdateById(oldModel.ClientId, oldModel.ProjectId, oldModel)
	if logErr != nil {
		// log.Fatal("ScreenSwipeManager", "UpdateScreenSwipe",
		// 	"ScreenSwipeDal_UpdateScreenSwipeById", logErr.Error())
		return oldModel, false, logErr
	}
	return oldModel, true, nil
}

func CalculateSwipeFirstDay(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel, total_swipe_hour int32) {
	switch {
	case total_swipe_hour <= 24:
		oldModel.FirstDayTotalSwipeUpCount = oldModel.FirstDayTotalSwipeUpCount + modelResponse.FirstDayTotalSwipeUpCount
		oldModel.FirstDayTotalSwipeDownCount = oldModel.FirstDayTotalSwipeDownCount + modelResponse.FirstDayTotalSwipeDownCount
		oldModel.FirstDayTotalSwipeRightCount = oldModel.FirstDayTotalSwipeRightCount + modelResponse.FirstDayTotalSwipeRightCount
		oldModel.FirstDayTotalSwipeLeftCount = oldModel.FirstDayTotalSwipeLeftCount + modelResponse.FirstDayTotalSwipeLeftCount

		oldModel.FirstDaySwipeTotalStartXCor = oldModel.FirstDaySwipeTotalStartXCor + modelResponse.FirstDaySwipeTotalStartXCor
		oldModel.FirstDaySwipeTotalStartYCor = oldModel.FirstDaySwipeTotalStartYCor + modelResponse.FirstDaySwipeTotalStartYCor
		oldModel.FirstDaySwipeTotalFinishXCor = oldModel.FirstDaySwipeTotalFinishXCor + modelResponse.FirstDaySwipeTotalFinishXCor
		oldModel.FirstDaySwipeTotalFinishYCor = oldModel.FirstDaySwipeTotalFinishYCor + modelResponse.FirstDaySwipeTotalFinishYCor
	}
}

func CalculateSwipeSecondDay(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel, total_swipe_hour int32) {
	switch {
	case total_swipe_hour <= 48 && total_swipe_hour > 24:
		oldModel.SecondDayTotalSwipeUpCount = oldModel.SecondDayTotalSwipeUpCount + modelResponse.FirstDayTotalSwipeUpCount
		oldModel.SecondDayTotalSwipeDownCount = oldModel.SecondDayTotalSwipeDownCount + modelResponse.FirstDayTotalSwipeDownCount
		oldModel.SecondDayTotalSwipeRightCount = oldModel.SecondDayTotalSwipeRightCount + modelResponse.FirstDayTotalSwipeRightCount
		oldModel.SecondDayTotalSwipeLeftCount = oldModel.SecondDayTotalSwipeLeftCount + modelResponse.FirstDayTotalSwipeLeftCount

		oldModel.SecondDaySwipeTotalStartXCor = oldModel.SecondDaySwipeTotalStartXCor + modelResponse.FirstDaySwipeTotalStartXCor
		oldModel.SecondDaySwipeTotalStartYCor = oldModel.SecondDaySwipeTotalStartYCor + modelResponse.FirstDaySwipeTotalStartYCor
		oldModel.SecondDaySwipeTotalFinishXCor = oldModel.SecondDaySwipeTotalFinishXCor + modelResponse.FirstDaySwipeTotalFinishXCor
		oldModel.SecondDaySwipeTotalFinishYCor = oldModel.SecondDaySwipeTotalFinishYCor + modelResponse.FirstDaySwipeTotalFinishYCor
	}
}

func CalculateSwipeThirdDay(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel, total_swipe_hour int32) {
	switch {
	case total_swipe_hour <= 72 && total_swipe_hour > 48:
		oldModel.ThirdDayTotalSwipeUpCount = oldModel.ThirdDayTotalSwipeUpCount + modelResponse.FirstDayTotalSwipeUpCount
		oldModel.ThirdDayTotalSwipeDownCount = oldModel.ThirdDayTotalSwipeDownCount + modelResponse.FirstDayTotalSwipeDownCount
		oldModel.ThirdDayTotalSwipeRightCount = oldModel.ThirdDayTotalSwipeRightCount + modelResponse.FirstDayTotalSwipeRightCount
		oldModel.ThirdDayTotalSwipeLeftCount = oldModel.ThirdDayTotalSwipeLeftCount + modelResponse.FirstDayTotalSwipeLeftCount

		oldModel.ThirdDaySwipeTotalStartXCor = oldModel.ThirdDaySwipeTotalStartXCor + modelResponse.FirstDaySwipeTotalStartXCor
		oldModel.ThirdDaySwipeTotalStartYCor = oldModel.ThirdDaySwipeTotalStartYCor + modelResponse.FirstDaySwipeTotalStartYCor
		oldModel.ThirdDaySwipeTotalFinishXCor = oldModel.ThirdDaySwipeTotalFinishXCor + modelResponse.FirstDaySwipeTotalFinishXCor
		oldModel.ThirdDaySwipeTotalFinishYCor = oldModel.ThirdDaySwipeTotalFinishYCor + modelResponse.FirstDaySwipeTotalFinishYCor
	}
}

func CalculateSwipeFourthDay(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel, total_swipe_hour int32) {
	switch {
	case total_swipe_hour <= 96 && total_swipe_hour > 72:
		oldModel.FourthDayTotalSwipeUpCount = oldModel.FourthDayTotalSwipeUpCount + modelResponse.FirstDayTotalSwipeUpCount
		oldModel.FourthDayTotalSwipeDownCount = oldModel.FourthDayTotalSwipeDownCount + modelResponse.FirstDayTotalSwipeDownCount
		oldModel.FourthDayTotalSwipeRightCount = oldModel.FourthDayTotalSwipeRightCount + modelResponse.FirstDayTotalSwipeRightCount
		oldModel.FourthDayTotalSwipeLeftCount = oldModel.FourthDayTotalSwipeLeftCount + modelResponse.FirstDayTotalSwipeLeftCount

		oldModel.FourthDaySwipeTotalStartXCor = oldModel.FourthDaySwipeTotalStartXCor + modelResponse.FirstDaySwipeTotalStartXCor
		oldModel.FourthDaySwipeTotalStartYCor = oldModel.FourthDaySwipeTotalStartYCor + modelResponse.FirstDaySwipeTotalStartYCor
		oldModel.FourthDaySwipeTotalFinishXCor = oldModel.FourthDaySwipeTotalFinishXCor + modelResponse.FirstDaySwipeTotalFinishXCor
		oldModel.FourthDaySwipeTotalFinishYCor = oldModel.FourthDaySwipeTotalFinishYCor + modelResponse.FirstDaySwipeTotalFinishYCor
	}
}

func CalculateSwipeFifthDay(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel, total_swipe_hour int32) {
	switch {
	case total_swipe_hour <= 120 && total_swipe_hour > 96:
		oldModel.FifthDayTotalSwipeUpCount = oldModel.FifthDayTotalSwipeUpCount + modelResponse.FirstDayTotalSwipeUpCount
		oldModel.FifthDayTotalSwipeDownCount = oldModel.FifthDayTotalSwipeDownCount + modelResponse.FirstDayTotalSwipeDownCount
		oldModel.FifthDayTotalSwipeRightCount = oldModel.FifthDayTotalSwipeRightCount + modelResponse.FirstDayTotalSwipeRightCount
		oldModel.FifthDayTotalSwipeLeftCount = oldModel.FifthDayTotalSwipeLeftCount + modelResponse.FirstDayTotalSwipeLeftCount

		oldModel.FifthDaySwipeTotalStartXCor = oldModel.FifthDaySwipeTotalStartXCor + modelResponse.FirstDaySwipeTotalStartXCor
		oldModel.FifthDaySwipeTotalStartYCor = oldModel.FifthDaySwipeTotalStartYCor + modelResponse.FirstDaySwipeTotalStartYCor
		oldModel.FifthDaySwipeTotalFinishXCor = oldModel.FifthDaySwipeTotalFinishXCor + modelResponse.FirstDaySwipeTotalFinishXCor
		oldModel.FifthDaySwipeTotalFinishYCor = oldModel.FifthDaySwipeTotalFinishYCor + modelResponse.FirstDaySwipeTotalFinishYCor
	}
}

func CalculateSwipeSixthDay(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel, total_swipe_hour int32) {
	switch {
	case total_swipe_hour <= 144 && total_swipe_hour > 120:
		oldModel.SixthDayTotalSwipeUpCount = oldModel.SixthDayTotalSwipeUpCount + modelResponse.FirstDayTotalSwipeUpCount
		oldModel.SixthDayTotalSwipeDownCount = oldModel.SixthDayTotalSwipeDownCount + modelResponse.FirstDayTotalSwipeDownCount
		oldModel.SixthDayTotalSwipeRightCount = oldModel.SixthDayTotalSwipeRightCount + modelResponse.FirstDayTotalSwipeRightCount
		oldModel.SixthDayTotalSwipeLeftCount = oldModel.SixthDayTotalSwipeLeftCount + modelResponse.FirstDayTotalSwipeLeftCount

		oldModel.SixthDaySwipeTotalStartXCor = oldModel.SixthDaySwipeTotalStartXCor + modelResponse.FirstDaySwipeTotalStartXCor
		oldModel.SixthDaySwipeTotalStartYCor = oldModel.SixthDaySwipeTotalStartYCor + modelResponse.FirstDaySwipeTotalStartYCor
		oldModel.SixthDaySwipeTotalFinishXCor = oldModel.SixthDaySwipeTotalFinishXCor + modelResponse.FirstDaySwipeTotalFinishXCor
		oldModel.SixthDaySwipeTotalFinishYCor = oldModel.SixthDaySwipeTotalFinishYCor + modelResponse.FirstDaySwipeTotalFinishYCor
	}
}

func CalculateSwipeSeventhDay(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel, total_swipe_hour int32) {
	switch {
	case total_swipe_hour <= 168 && total_swipe_hour > 144:
		oldModel.SeventhDayTotalSwipeUpCount = oldModel.SeventhDayTotalSwipeUpCount + modelResponse.FirstDayTotalSwipeUpCount
		oldModel.SeventhDayTotalSwipeDownCount = oldModel.SeventhDayTotalSwipeDownCount + modelResponse.FirstDayTotalSwipeDownCount
		oldModel.SeventhDayTotalSwipeRightCount = oldModel.SeventhDayTotalSwipeRightCount + modelResponse.FirstDayTotalSwipeRightCount
		oldModel.SeventhDayTotalSwipeLeftCount = oldModel.SeventhDayTotalSwipeLeftCount + modelResponse.FirstDayTotalSwipeLeftCount

		oldModel.SeventhDaySwipeTotalStartXCor = oldModel.SeventhDaySwipeTotalStartXCor + modelResponse.FirstDaySwipeTotalStartXCor
		oldModel.SeventhDaySwipeTotalStartYCor = oldModel.SeventhDaySwipeTotalStartYCor + modelResponse.FirstDaySwipeTotalStartYCor
		oldModel.SeventhDaySwipeTotalFinishXCor = oldModel.SeventhDaySwipeTotalFinishXCor + modelResponse.FirstDaySwipeTotalFinishXCor
		oldModel.SeventhDaySwipeTotalFinishYCor = oldModel.SeventhDaySwipeTotalFinishYCor + modelResponse.FirstDaySwipeTotalFinishYCor
	}
}

func DetermineSwipeDirection(modelResponse *model.ScreenSwipeResponseModel, swipeDirection byte) {
	switch swipeDirection {
	case 1:
		modelResponse.FirstDayTotalSwipeRightCount = 1
		modelResponse.TotalSwipeRightCount = 1
	case 2:
		modelResponse.FirstDayTotalSwipeLeftCount = 1
		modelResponse.TotalSwipeLeftCount = 1
	case 3:
		modelResponse.FirstDayTotalSwipeUpCount = 1
		modelResponse.TotalSwipeUpCount = 1
	case 4:
		modelResponse.FirstDayTotalSwipeDownCount = 1
		modelResponse.TotalSwipeDownCount = 1
	}
}

func CalculateSwipeNumber(modelResponse *model.ScreenSwipeResponseModel, oldModel *model.ScreenSwipeResponseModel) {
	switch oldModel.TotalSwipeSessionCount {
	case 2:
		oldModel.SecondSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.SecondSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.SecondSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.SecondSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.SecondSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
	case 3:
		oldModel.ThirdSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.ThirdSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.ThirdSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.ThirdSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.ThirdSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
	case 4:
		oldModel.FourthSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.FourthSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.FourthSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.FourthSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.FourthSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
	case 5:
		oldModel.FifthSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.FifthSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.FifthSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.FifthSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.FifthSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
	case 6:
		oldModel.SixthSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.SixthSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.SixthSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.SixthSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.SixthSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
	case 7:
		oldModel.SeventhSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.SeventhSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.SeventhSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.SeventhSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.SeventhSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
	}
}
