package concrete

import (
	model "FilterWorkerService/internal/model"
	IBuyingEventDal "FilterWorkerService/internal/repository/abstract"
	//ICacheService "FilterWorkerService/internal/service/abstract"
	IJsonParser "FilterWorkerService/pkg/jsonParser"
)

type BuyingEventManager struct {
	IBuyingEventDal IBuyingEventDal.IBuyingEventDal
	IJsonParser             IJsonParser.IJsonParser
}

func (b *BuyingEventManager) AddBuyingEvent(data *[]byte) (s bool, m string) {
	// Todo : 1 Model karşılanacak
	firstModel := model.BuyingEventModel{}
	err := b.IJsonParser.DecodeJson(data, &firstModel)
	if err != nil {
		return false, err.Error()
	}
	
	//Geçiş ve video reklamı filtrelenecek




	// Todo: 2 Filtreler Buraya Yazılacak
	modelResponse := model.BuyingEventRespondModel{}
	modelResponse.ClientId = firstModel.ClientId
	modelResponse.ProjectId = firstModel.ProjectId
	modelResponse.CustomerId = firstModel.CustomerId
	modelResponse.LevelIndex = int64(firstModel.LevelIndex)
	modelResponse.TotalBuyingCount = 1
	modelResponse.TotalBuyingDay = calculateTotalDay(&modelResponse)
	modelResponse.TotalBuyingSession = 1
	//modelResponse.TotalSession
	//modelResponse.TotalDay
	modelResponse.FirstBuyingMonth = int64(firstModel.TrigerdTime.Month())
	modelResponse.FirstBuyingWeek = int64(firstModel.TrigerdTime.Weekday())
	modelResponse.FirstBuyingDay = int64(firstModel.TrigerdTime.Day())
	modelResponse.FirstBuyingHour = int64(firstModel.TrigerdTime.Hour())
	modelResponse.LastBuyingMonth = int64(firstModel.TrigerdTime.Month())
	modelResponse.LastBuyingWeek = int64(firstModel.TrigerdTime.Weekday())
	modelResponse.LastBuyingDay = int64(firstModel.TrigerdTime.Day())
	modelResponse.LastBuyingHour = int64(firstModel.TrigerdTime.Hour())
	modelResponse.FirstDayBuyingCount = 1
	modelResponse.PenultimateDayBuyingCount = modelResponse.FirstDayBuyingCount
	modelResponse.LastDayBuyingCount = 1
	modelResponse.LastMinusPenultimateDayBuyingCount = modelResponse.LastDayBuyingCount - modelResponse.PenultimateDayBuyingCount
	modelResponse.LastMinusFirstDayBuyingCount = modelResponse.LastDayBuyingCount - modelResponse.FirstDayBuyingCount
	determineDay(&firstModel, &modelResponse)
	determineHour(&firstModel, &modelResponse)
	determineAmPm(&firstModel, &modelResponse)
	//modelResponse.DailyAverageBuyingCount
	modelResponse.BuyingDayAverageBuyingCount = modelResponse.TotalBuyingCount/modelResponse.TotalBuyingDay
	//modelResponse.LevelBasedAverageBuyingCount = modelResponse.TotalBuyingCount/modelResponse.LevelIndex
	//modelResponse.SessionBasedAverageBuyingCount=modelResponse.TotalBuyingCount/modelResponse.TotalSession
	//modelResponse.FirstBuyingDayMinusFirstSessionDay = modelResponse.FirstBuyingDay - 
	//modelResponse.FirstBuyingMonthMinusFirstSessionMonth
	//modelResponse.TotalDifferenceBetweenFirstBuyingDayAndFirstSessionDay
	//modelResponse.IsDeadAndBuyingItemCount
	

	// Todo : 3 Model burada kayıt edilecek
	logErr := b.IBuyingEventDal.Add(&modelResponse)
	if logErr != nil {
		return false, logErr.Error()
	}
	return true, ""
}

func determineDay(firstModel *model.BuyingEventModel, modelResponse *model.BuyingEventRespondModel) {
	day := int64(firstModel.TrigerdTime.Day())
	if day==0 {
		modelResponse.SundayBuyingCount =  1
	}else if day == 1 {
		modelResponse.MondayBuyingCount = 1
	}else if day == 2 {
		modelResponse.TuesdayBuyingCount = 1
	}else if day == 3 {
		modelResponse.WednesdayBuyingCount = 1
	}else if day == 4 {
		modelResponse.ThursdayBuyingCount = 1
	}else if day == 5 {
		modelResponse.FridayBuyingCount = 1
	}else if day == 6 {
		modelResponse.SaturdayBuyingCount = 1
	}
}

func determineHour(firstModel *model.BuyingEventModel, modelResponse *model.BuyingEventRespondModel){
	hour := int64(firstModel.TrigerdTime.Hour())
	if (hour <= 5){
		modelResponse.Buying0To5HourCount = 1
	}else if (hour > 5) && (hour <= 11){
		modelResponse.Buying6To11HourCount = 1
	}else if (hour > 11) && (hour <= 17){
		modelResponse.Buying12To17HourCount = 1
	}else if (hour > 17) && (hour <= 23){
		modelResponse.Buying18To23HourCount = 1
	}
}

func determineAmPm(firstModel *model.BuyingEventModel, modelResponse *model.BuyingEventRespondModel) {
	hour := int64(firstModel.TrigerdTime.Hour())
	if (hour <= 12){
		modelResponse.AmBuyingCount = 1
	}else if (hour > 12){
		modelResponse.PmBuyingCount = 1
	}
}

func calculateTotalDay(modelResponse *model.BuyingEventRespondModel) int64 {
	modelResponse.TotalBuyingDay = (modelResponse.LastBuyingDay - modelResponse.FirstBuyingDay) + 30*(modelResponse.LastBuyingMonth-modelResponse.FirstBuyingMonth)
	return modelResponse.TotalBuyingDay
}