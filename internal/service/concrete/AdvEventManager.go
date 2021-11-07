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
	//modelResponse.TotalAdvDay
	//modelResponse.TotalVideoAdvCount > 1 video veya geçiş reklamı çıkmaasına göre belirlenecek
	//modelResponse.TotalInterstitialAdvCount
	//modelResponse.LevelBasedAverageInterstitialAdvCount
	//modelResponse.LevelBasedAverageVideoAdvCount
	//modelResponse.AverageDailyVideoAdvClickCount
	modelResponse.FirstVideoClickMonth = int64(firstModel.TrigerdTime.Month())
	modelResponse.FirstVideoClickWeek = int64(firstModel.TrigerdTime.Weekday())
	modelResponse.FirstVideoClickDay = int64(firstModel.TrigerdTime.Day())
	modelResponse.FirstVideoClickHour = int64(firstModel.TrigerdTime.Hour())
	//modelResponse.FirstDayVideoClickCount
	//modelResponse.PenultimateDayVdeoClickCount
	//modelResponse.LastDayVideoClickCount
	//modelResponse.LastDayVideoClickCount
	//modelResponse.LastMinusPenultimateDayVideoClickCount
	//modelResponse.LastMinusFirstDayVideoClickCount
	//modelResponse.LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount
	//modelResponse.SundayVideoAdvClickCount
	//modelResponse.MondayVideoAdvClickCount
	//modelResponse.TuesdayVideoAdvClickCount
	//modelResponse.WednesdayVideoAdvClickCount
	//modelResponse.ThursdayVideoAdvClickCount
	//modelResponse.FridayVideoAdvClickCount
	//modelResponse.SaturdayVideoAdvClickCount
	//modelResponse.AmVideoAdvClickCount
	//modelResponse.PmVideoAdvClickCount
	//modelResponse.VideoAdvClick0To5HourCount
	//modelResponse.VideoAdvClick6To11HourCount
	//modelResponse.VideoAdvClick12To17HourCount
	//modelResponse.VideoAdvClick18To23HourCount


	// Todo : 3 Model burada kayıt edilecek
	logErr := a.IAdvEventDal.Add(&modelResponse)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}