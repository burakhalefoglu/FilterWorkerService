package concrete

import (
	model "FilterWorkerService/internal/model"
	IScreenSwipeDal "FilterWorkerService/internal/repository/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type ScreenSwipeManager struct {
	IScreenSwipeDal IScreenSwipeDal.IScreenSwipeDal
	IJsonParser     IJsonParser.IJsonParser
}

func (sc *ScreenSwipeManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.ScreenSwipeModel{}
	err := sc.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	hour := int64(firstModel.CreationAt.Hour())
	yearOfDay := int64(firstModel.CreationAt.YearDay())
	year := int64(firstModel.CreationAt.Year())
	swipeDirection := int64(firstModel.SwipeDirection)
	modelResponse := model.ScreenSwipeRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.TotalSwipeSessionCount = 1
	modelResponse.FirstSwipeYearOfDay = yearOfDay
	modelResponse.FirstSwipeYear = year
	modelResponse.FirstSwipeHour = hour
	modelResponse.FistSwipeDirection = int64(firstModel.SwipeDirection)
	modelResponse.FirstSwipeStartXCor = firstModel.SwipeStartXCor
	modelResponse.FirstSwipeStartYCor = firstModel.SwipeStartYCor
	modelResponse.FirstSwipeFinishXCor = firstModel.SwipeFinishXCor
	modelResponse.FirstSwipeFinishYCor = firstModel.SwipeFinishYCor
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
	modelResponse.PenultimateSwipeDirection = 0
	modelResponse.PenultimateSwipeStartXCor = 0
	modelResponse.PenultimateSwipeStartYCor = 0
	modelResponse.PenultimateSwipeFinishXCor = 0
	modelResponse.PenultimateSwipeFinishYCor = 0
	modelResponse.LastSwipeDirection = int64(firstModel.SwipeDirection)
	modelResponse.LastSwipeStartXCor = firstModel.SwipeStartXCor
	modelResponse.LastSwipeStartYCor = firstModel.SwipeStartYCor
	modelResponse.LastSwipeFinishXCor = firstModel.SwipeFinishXCor
	modelResponse.LastSwipeFinishYCor = firstModel.SwipeFinishYCor
	determineSwipeDirection(&modelResponse, swipeDirection)
	// modelResponse.FirstDayTotalSwipeUpCount
	// modelResponse.FirstDayTotalSwipeDownCount
	// modelResponse.FirstDayTotalSwipeRightCount
	// modelResponse.FirstDayTotalSwipeLeftCount
	modelResponse.FirstDaySwipeTotalStartXCor = firstModel.SwipeStartXCor
	modelResponse.FirstDaySwipeTotalStartYCor = firstModel.SwipeStartYCor
	modelResponse.FirstDaySwipeTotalFinishXCor = firstModel.SwipeFinishXCor
	modelResponse.FirstDaySwipeTotalFinishYCor = firstModel.SwipeFinishYCor
	//modelResponse.TotalSwipeUpCount
	//modelResponse.TotalSwipeDownCount
	//modelResponse.TotalSwipeRightCount
	//modelResponse.TotalSwipeLeftCount
	modelResponse.TotalSwipeStartXCor = firstModel.SwipeStartXCor
	modelResponse.TotalSwipeStartYCor = firstModel.SwipeStartYCor
	modelResponse.TotalSwipeFinishXCor = firstModel.SwipeFinishXCor
	modelResponse.TotalSwipeFinishYCor = firstModel.SwipeFinishYCor

	oldModel, err := sc.IScreenSwipeDal.GetScreenSwipeByCustomerId(modelResponse.CustomerId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := sc.IScreenSwipeDal.Add(&modelResponse)
		if logErr != nil {
			return false, logErr.Error()
		}
		return true, ""

	case err == nil:
		updateResult, updateErr := sc.updateScreenSwipe(&modelResponse, oldModel)
		if updateErr != nil {
			return updateResult, updateErr.Error()
		}
		return updateResult, ""

	default:

		return false, err.Error()

	}
}


func (sc *ScreenSwipeManager) updateScreenSwipe(modelResponse *model.ScreenSwipeRespondModel, oldModel *model.ScreenSwipeRespondModel) (s bool, m error) {
	
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	oldModel.TotalSwipeSessionCount = modelResponse.TotalSwipeSessionCount + oldModel.TotalSwipeSessionCount
	// modelResponse.FirstSwipeYearOfDay
	// modelResponse.FirstSwipeYear
	// modelResponse.FirstSwipeHour
	// oldModel.FistSwipeDirection
	// oldModel.FirstSwipeStartXCor
	// oldModel.FirstSwipeStartYCor
	// oldModel.FirstSwipeFinishXCor
	// oldModel.FirstSwipeFinishYCor
	modelResponse.SecondSwipeDirection, modelResponse.SecondSwipeStartXCor, modelResponse.SecondSwipeStartYCor, modelResponse.SecondSwipeFinishXCor, modelResponse.SecondSwipeFinishYCor = calculateSecondSwipe(modelResponse, oldModel)
	modelResponse.ThirdSwipeDirection, modelResponse.ThirdSwipeStartXCor, modelResponse.ThirdSwipeStartYCor, modelResponse.ThirdSwipeFinishXCor, modelResponse.ThirdSwipeFinishYCor = calculateThirdSwipe(modelResponse, oldModel)
	oldModel.PenultimateSwipeDirection = oldModel.LastSwipeDirection
	modelResponse.PenultimateSwipeStartXCor = oldModel.LastSwipeStartXCor
	modelResponse.PenultimateSwipeStartYCor = oldModel.LastSwipeStartYCor
	modelResponse.PenultimateSwipeFinishXCor = oldModel.LastSwipeFinishXCor
	modelResponse.PenultimateSwipeFinishYCor = oldModel.LastSwipeFinishYCor
	oldModel.LastSwipeDirection = modelResponse.LastSwipeDirection
	oldModel.LastSwipeStartXCor = modelResponse.LastSwipeStartXCor
	oldModel.LastSwipeStartYCor = modelResponse.LastSwipeStartYCor
	oldModel.LastSwipeFinishXCor = modelResponse.LastSwipeFinishXCor
	oldModel.LastSwipeFinishYCor = modelResponse.LastSwipeFinishYCor
	oldModel.FirstDayTotalSwipeUpCount, oldModel.FirstDayTotalSwipeDownCount, oldModel.FirstDayTotalSwipeRightCount, oldModel.FirstDayTotalSwipeLeftCount = calculateFirstDaySwipeDirectionCount(modelResponse, oldModel)
	oldModel.FirstDaySwipeTotalStartXCor, oldModel.FirstDaySwipeTotalStartYCor, oldModel.FirstDaySwipeTotalFinishXCor, oldModel.FirstDaySwipeTotalFinishYCor = calculateFirstDaySwipeCorCount(modelResponse, oldModel)
	oldModel.TotalSwipeUpCount = modelResponse.TotalSwipeUpCount + oldModel.TotalSwipeUpCount
	oldModel.TotalSwipeDownCount = modelResponse.TotalSwipeDownCount + oldModel.TotalSwipeDownCount
	oldModel.TotalSwipeRightCount = modelResponse.TotalSwipeRightCount + oldModel.TotalSwipeRightCount
	oldModel.TotalSwipeLeftCount = modelResponse.TotalSwipeLeftCount + oldModel.TotalSwipeLeftCount
	oldModel.TotalSwipeStartXCor = modelResponse.TotalSwipeStartXCor + oldModel.TotalSwipeStartXCor
	oldModel.TotalSwipeStartYCor = modelResponse.TotalSwipeStartYCor + oldModel.TotalSwipeStartYCor
	oldModel.TotalSwipeFinishXCor = modelResponse.TotalSwipeFinishXCor + oldModel.TotalSwipeFinishXCor
	oldModel.TotalSwipeFinishYCor = modelResponse.TotalSwipeFinishYCor + oldModel.TotalSwipeFinishYCor
	logErr := sc.IScreenSwipeDal.UpdateScreenSwipeByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr
	}
	return true, nil
}

func determineSwipeDirection(modelResponse *model.ScreenSwipeRespondModel, swipeDirection int64) {
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

func calculateSecondSwipe(modelResponse *model.ScreenSwipeRespondModel, oldModel *model.ScreenSwipeRespondModel) (direction int64, startx float64, starty float64, finishx float64, finishy float64) {
	if oldModel.TotalSwipeSessionCount == 2 {
		oldModel.SecondSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.SecondSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.SecondSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.SecondSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.SecondSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
		return oldModel.SecondSwipeDirection, oldModel.SecondSwipeStartXCor, oldModel.SecondSwipeStartYCor, oldModel.SecondSwipeFinishXCor, oldModel.SecondSwipeFinishYCor
	}
	return oldModel.SecondSwipeDirection, oldModel.SecondSwipeStartXCor, oldModel.SecondSwipeStartYCor, oldModel.SecondSwipeFinishXCor, oldModel.SecondSwipeFinishYCor
}

func calculateThirdSwipe(modelResponse *model.ScreenSwipeRespondModel, oldModel *model.ScreenSwipeRespondModel) (direction int64, startx float64, starty float64, finishx float64, finishy float64) {
	if oldModel.TotalSwipeSessionCount == 3 {
		oldModel.ThirdSwipeDirection = modelResponse.FistSwipeDirection
		oldModel.ThirdSwipeStartXCor = modelResponse.FirstSwipeStartXCor
		oldModel.ThirdSwipeStartYCor = modelResponse.FirstSwipeStartYCor
		oldModel.ThirdSwipeFinishXCor = modelResponse.FirstSwipeFinishXCor
		oldModel.ThirdSwipeFinishYCor = modelResponse.FirstSwipeFinishYCor
		return oldModel.SecondSwipeDirection, oldModel.SecondSwipeStartXCor, oldModel.SecondSwipeStartYCor, oldModel.SecondSwipeFinishXCor, oldModel.SecondSwipeFinishYCor
	}
	return oldModel.SecondSwipeDirection, oldModel.SecondSwipeStartXCor, oldModel.SecondSwipeStartYCor, oldModel.SecondSwipeFinishXCor, oldModel.SecondSwipeFinishYCor
}

func calculateFirstDaySwipeDirectionCount(modelResponse *model.ScreenSwipeRespondModel, oldModel *model.ScreenSwipeRespondModel) (up int64, down int64, right int64, left int64) {
	if (oldModel.FirstSwipeYearOfDay == modelResponse.FirstSwipeYearOfDay) && (oldModel.FirstSwipeYear == modelResponse.FirstSwipeYear) {
		oldModel.FirstDayTotalSwipeUpCount = modelResponse.FirstDayTotalSwipeUpCount + oldModel.FirstDayTotalSwipeUpCount
		oldModel.FirstDayTotalSwipeDownCount = modelResponse.FirstDayTotalSwipeDownCount + oldModel.FirstDayTotalSwipeDownCount
		oldModel.FirstDayTotalSwipeRightCount = modelResponse.FirstDayTotalSwipeRightCount + oldModel.FirstDayTotalSwipeRightCount
		oldModel.FirstDayTotalSwipeLeftCount = modelResponse.FirstDayTotalSwipeLeftCount + oldModel.FirstDayTotalSwipeLeftCount
		return oldModel.FirstDayTotalSwipeUpCount, oldModel.FirstDayTotalSwipeDownCount, oldModel.FirstDayTotalSwipeRightCount, oldModel.FirstDayTotalSwipeLeftCount
	}
	return oldModel.FirstDayTotalSwipeUpCount, oldModel.FirstDayTotalSwipeDownCount, oldModel.FirstDayTotalSwipeRightCount, oldModel.FirstDayTotalSwipeLeftCount
}

func calculateFirstDaySwipeCorCount(modelResponse *model.ScreenSwipeRespondModel, oldModel *model.ScreenSwipeRespondModel) (startx float64, starty float64, finishx float64, finishy float64) {
	if (oldModel.FirstSwipeYearOfDay == modelResponse.FirstSwipeYearOfDay) && (oldModel.FirstSwipeYear == modelResponse.FirstSwipeYear) {
		modelResponse.FirstDaySwipeTotalStartXCor = modelResponse.FirstDaySwipeTotalStartXCor + oldModel.FirstDaySwipeTotalStartXCor
		modelResponse.FirstDaySwipeTotalStartYCor = modelResponse.FirstDaySwipeTotalStartYCor + oldModel.FirstDaySwipeTotalStartYCor
		modelResponse.FirstDaySwipeTotalFinishXCor = modelResponse.FirstDaySwipeTotalFinishXCor + oldModel.FirstDaySwipeTotalFinishXCor
		modelResponse.FirstDaySwipeTotalFinishYCor = modelResponse.FirstDaySwipeTotalFinishYCor + oldModel.FirstDaySwipeTotalFinishYCor
		return modelResponse.FirstDaySwipeTotalStartXCor, modelResponse.FirstDaySwipeTotalStartYCor, modelResponse.FirstDaySwipeTotalFinishXCor, modelResponse.FirstDaySwipeTotalFinishYCor
	}
	return modelResponse.FirstDaySwipeTotalStartXCor, modelResponse.FirstDaySwipeTotalStartYCor, modelResponse.FirstDaySwipeTotalFinishXCor, modelResponse.FirstDaySwipeTotalFinishYCor
}
