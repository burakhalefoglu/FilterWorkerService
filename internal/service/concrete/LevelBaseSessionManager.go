package concrete

import (
	model "FilterWorkerService/internal/model"
	ILevelBaseSessionDal "FilterWorkerService/internal/repository/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type LevelBaseSessionManager struct {
	ILevelBaseSessionDal ILevelBaseSessionDal.ILevelBaseSessionDal
	IJsonParser          IJsonParser.IJsonParser
}

func (l *LevelBaseSessionManager) ConvertRawModelToResponseModel(data *[]byte) (s bool, m string) {
	firstModel := model.LevelBaseSessionDataModel{}
	err := l.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	modelResponse := model.LevelBaseSessionRespondModel{}
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.TotalLevelBaseSessionCount = 1
	modelResponse.FirstLevelSessionLevelIndex = int64(firstModel.LevelIndex)
	modelResponse.FirstLevelSessionDuration = int64(firstModel.SessionTimeMinute)
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
	modelResponse.PenultimateLevelSessionLevelIndex = 0
	modelResponse.PenultimateLevelSessionLevelDuration = 0
	modelResponse.LastLevelSessionLevelIndex = int64(firstModel.LevelIndex)
	modelResponse.LastLevelSessionLevelDuration = int64(firstModel.SessionTimeMinute)

	oldModel, err := l.ILevelBaseSessionDal.GetLevelBaseSessionById(modelResponse.ClientId)
	switch {
	case err.Error() == "mongo: no documents in result":

		logErr := l.ILevelBaseSessionDal.Add(&modelResponse)
		if logErr != nil {
			return false, logErr.Error()
		}
		return true, ""

	case err == nil:
		updateResult, updateErr := l.updateLevelBaseSession(&modelResponse, oldModel)
		if updateErr != nil {
			return updateResult, updateErr.Error()
		}
		return updateResult, ""

	default:

		return false, err.Error()

	}
}


func (l *LevelBaseSessionManager) updateLevelBaseSession(modelResponse *model.LevelBaseSessionRespondModel, oldModel *model.LevelBaseSessionRespondModel) (s bool, m error) {
	
	oldModel.ProjectId = modelResponse.ProjectId
	oldModel.ClientId = modelResponse.ClientId
	oldModel.CustomerId = modelResponse.CustomerId
	oldModel.TotalLevelBaseSessionCount = modelResponse.TotalLevelBaseSessionCount + oldModel.TotalLevelBaseSessionCount
	calculateLevelIndexBaseSession(modelResponse, oldModel)
	// oldModel.FirstLevelSessionLevelIndex
	// oldModel.FirstLevelSessionDuration
	// oldModel.SecondLevelSessionLevelIndex
	// oldModel.SecondLevelSessionDuration
	// modelResponse.ThirdLevelSessionLevelIndex
	// modelResponse.ThirdLevelSessionDuration
	// modelResponse.FourLevelSessionLevelIndex
	// modelResponse.FourLevelSessionDuration
	// modelResponse.FiveLevelSessionLevelIndex
	// modelResponse.FiveLevelSessionDuration
	// modelResponse.SixLevelSessionLevelIndex
	// modelResponse.SixLevelSessionDuration
	// modelResponse.SevenLevelSessionLevelIndex
	// modelResponse.SevenLevelSessionDuration
	oldModel.PenultimateLevelSessionLevelIndex = oldModel.LastLevelSessionLevelIndex
	oldModel.PenultimateLevelSessionLevelDuration = oldModel.LastLevelSessionLevelDuration
	oldModel.LastLevelSessionLevelIndex = modelResponse.LastLevelSessionLevelIndex
	oldModel.LastLevelSessionLevelDuration = modelResponse.LastLevelSessionLevelDuration
	logErr := l.ILevelBaseSessionDal.UpdateLevelBaseSessionById(oldModel.ClientId, oldModel)
	if logErr != nil {
		return false, logErr
	}
	return true, nil
}

func calculateLevelIndexBaseSession(modelResponse *model.LevelBaseSessionRespondModel, oldModel *model.LevelBaseSessionRespondModel) {
	switch oldModel.TotalLevelBaseSessionCount {
	case 2:
		oldModel.SecondLevelSessionLevelIndex = modelResponse.SecondLevelSessionLevelIndex
		oldModel.SecondLevelSessionDuration = modelResponse.SecondLevelSessionDuration
	case 3:
		oldModel.ThirdLevelSessionLevelIndex = modelResponse.ThirdLevelSessionLevelIndex
		oldModel.ThirdLevelSessionLevelIndex = modelResponse.ThirdLevelSessionLevelIndex
	case 4:
		oldModel.FourLevelSessionLevelIndex = modelResponse.FourLevelSessionLevelIndex
		oldModel.FourLevelSessionDuration = modelResponse.FourLevelSessionDuration
	case 5:
		oldModel.FiveLevelSessionLevelIndex = modelResponse.FiveLevelSessionLevelIndex
		oldModel.FiveLevelSessionDuration = modelResponse.FiveLevelSessionDuration
	case 6:
		oldModel.SixLevelSessionLevelIndex = modelResponse.SixLevelSessionLevelIndex
		oldModel.SixLevelSessionDuration = modelResponse.SixLevelSessionDuration
	case 7:
		oldModel.SevenLevelSessionLevelIndex = modelResponse.SevenLevelSessionLevelIndex
		oldModel.SevenLevelSessionDuration = modelResponse.SevenLevelSessionDuration
	}
}
