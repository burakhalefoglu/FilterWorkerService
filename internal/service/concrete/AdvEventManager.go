package concrete

import (
	model "FilterWorkerService/internal/model"
	IAdvEventDal "FilterWorkerService/internal/repository/abstract"
	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type AdvEventManager struct {
	IAdvEventDal IAdvEventDal.IAdvEventDal
	IJsonParser             IJsonParser.IJsonParser
}

func (a *AdvEventManager) AddAdvEvent(data *[]byte) (s bool, m string) {
	// Todo : 1 Model karşılanacak
	firstModel := model.AdvEventModel{}
	err := a.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	
	//Geçiş ve video reklamı filtrelenecek




	// Todo: 2 Filtreler Buraya Yazılacak
	modelResponse := model.AdvEventRespondModel{}
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	//modelResponse.VideoAdvCount = 1 video veya geçiş reklamı çıkmaasına göre belirlenecek
	//modelResponse.InterstitialAdvCount = 1
	modelResponse.VideoClickMonth = int64(firstModel.TrigerdTime.Month())
	modelResponse.VideoClickWeek = int64(firstModel.TrigerdTime.Weekday())
	modelResponse.VideoClickDay = int64(firstModel.TrigerdTime.Day())
	modelResponse.VideoClickHour = int64(firstModel.TrigerdTime.Hour())
	
	// Todo : 3 Model burada kayıt edilecek
	logErr := a.IAdvEventDal.Add(&modelResponse)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}