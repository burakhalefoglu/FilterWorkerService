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

func (sc *ScreenSwipeManager) ConvertRawModelToResponseModel(data *[]byte) (respondModel *model.ScreenSwipeRespondModel, s bool, m string){
	firstModel := model.ScreenSwipeModel{}
	err := sc.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return &model.ScreenSwipeRespondModel{}, false, err.Error()
	}
	modelResponse := model.ScreenSwipeRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.FistSwipeDirection = int64(firstModel.SwipeDirection)
	modelResponse.FirstSwipeStartXCor = firstModel.SwipeStartXCor
	modelResponse.FirstSwipeStartYCor = firstModel.SwipeStartYCor
	modelResponse.FirstSwipeFinishXCor = firstModel.SwipeFinishXCor
	modelResponse.FirstSwipeFinishYCor = firstModel.SwipeFinishYCor
	modelResponse.LastSwipeDirection = int64(firstModel.SwipeDirection)
	modelResponse.LastSwipeStartXCor = firstModel.SwipeStartXCor
	modelResponse.LastSwipeStartYCor = firstModel.SwipeStartYCor
	modelResponse.LastSwipeFinishXCor = firstModel.SwipeFinishXCor
	modelResponse.LastSwipeFinishYCor = firstModel.SwipeFinishYCor
	//modelResponse.TotalSwipeUpCount
	//modelResponse.TotalSwipeDownCount
	//modelResponse.TotalSwipeRightCount
	//modelResponse.TotalSwipeLeftCount
	// modelResponse.TotalSwipeStartXCor
	// modelResponse.TotalSwipeStartYCor
	// modelResponse.TotalSwipeFinishXCor
	// modelResponse.TotalSwipeFinishYCor
	modelResponse.TotalSwipeSessionCount = 1
	return &modelResponse, true, ""
}

func (sc *ScreenSwipeManager) AddScreenSwipe(data *model.ScreenSwipeRespondModel) (s bool, m string){
	logErr := sc.IScreenSwipeDal.Add(data)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func (sc *ScreenSwipeManager) UpdateScreenSwipe(modelResponse *model.ScreenSwipeRespondModel) (s bool, m string){
	oldModel, err := sc.IScreenSwipeDal.GetScreenSwipeByCustomerId(modelResponse.CustomerId)
	if err != nil {
		return false, err.Error()
	}
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.LevelIndex = modelResponse.LevelIndex
	// oldModel.FistSwipeDirection 
	// oldModel.FirstSwipeStartXCor
	// oldModel.FirstSwipeStartYCor
	// oldModel.FirstSwipeFinishXCor
	// oldModel.FirstSwipeFinishYCor
	oldModel.LastSwipeDirection = modelResponse.LastSwipeDirection
	oldModel.LastSwipeStartXCor = modelResponse.LastSwipeStartXCor
	oldModel.LastSwipeStartYCor = modelResponse.LastSwipeStartYCor
	oldModel.LastSwipeFinishXCor = modelResponse.LastSwipeFinishXCor
	oldModel.LastSwipeFinishYCor = modelResponse.LastSwipeFinishYCor
	oldModel.TotalSwipeUpCount = modelResponse.TotalSwipeUpCount +oldModel.TotalSwipeUpCount
	oldModel.TotalSwipeDownCount = modelResponse.TotalSwipeDownCount + oldModel.TotalSwipeDownCount
	oldModel.TotalSwipeRightCount = modelResponse.TotalSwipeRightCount + oldModel.TotalSwipeRightCount
	oldModel.TotalSwipeLeftCount = modelResponse.TotalSwipeLeftCount + oldModel.TotalSwipeLeftCount
	oldModel.TotalSwipeStartXCor = modelResponse.TotalSwipeStartXCor + oldModel.TotalSwipeStartXCor
	oldModel.TotalSwipeStartYCor = modelResponse.TotalSwipeStartYCor + oldModel.TotalSwipeStartYCor
	oldModel.TotalSwipeFinishXCor = modelResponse.TotalSwipeFinishXCor + oldModel.TotalSwipeFinishXCor
	oldModel.TotalSwipeFinishYCor = modelResponse.TotalSwipeFinishYCor + oldModel.TotalSwipeFinishYCor
	oldModel.TotalSwipeSessionCount = modelResponse.TotalSwipeSessionCount + oldModel.TotalSwipeSessionCount
	logErr := sc.IScreenSwipeDal.UpdateScreenSwipeByCustomerId(oldModel.CustomerId, oldModel)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}