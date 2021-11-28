package test

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/Log"
	"FilterWorkerService/test/Mock/repository"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var buyingModel = model.BuyingEventModel{
	ProjectId:     "Test",
	ClientId:      "Test",
	CustomerId:    "Test",
	LevelName:     "",
	LevelIndex:    1,
	InWhatMinutes: 5,
	ProductType:   "TestProduct",
	TrigerdTime: time.Date(
		2021, 11, 5, 16, 11, 36, 651387237, time.UTC),
}

var newBuyingModel = model.BuyingEventRespondModel{
	ProjectId:                    "Test",
	ClientId:                     "Test",
	CustomerId:                   "Test",
	LevelIndex:                   int64(buyingModel.LevelIndex),
	TotalBuyingCount:             1,
	TotalBuyingDay:               1,
	FirstBuyingYearOfDay:         int64(buyingModel.TrigerdTime.YearDay()),
	FirstBuyingYear:              int64(buyingModel.TrigerdTime.Year()),
	FirstBuyingHour:              int64(buyingModel.TrigerdTime.Hour()),
	FirstBuyingMinute:            int64(buyingModel.InWhatMinutes),
	SecondBuyingYearOfDay:        0,
	SecondBuyingHour:             0,
	ThirdBuyingYearOfDay:         0,
	ThirdBuyingHour:              0,
	PenultimateBuyingYearOfDay:   0,
	PenultimateBuyingHour:        0,
	LastBuyingYearOfDay:          0,
	LastBuyingYear:               0,
	LastBuyingHour:               0,
	FirstDayBuyingCount:          1,
	PenultimateDayBuyingCount:    0,
	LastDayBuyingCount:           0,
	LastMinusFirstDayBuyingCount: -1,
	SundayBuyingCount:            0,
	MondayBuyingCount:            0,
	TuesdayBuyingCount:           1,
	WednesdayBuyingCount:         0,
	ThursdayBuyingCount:          0,
	FridayBuyingCount:            0,
	SaturdayBuyingCount:          0,
	AmBuyingCount:                0,
	PmBuyingCount:                1,
	Buying0To5HourCount:          0,
	Buying6To11HourCount:         0,
	Buying12To17HourCount:        1,
	Buying18To23HourCount:        0,
	BuyingDayAverageBuyingCount:  1,
	LevelBasedAverageBuyingCount: 1,
}

var then = time.Date(
	2021, 11, 5, 20, 34, 58, 651387237, time.UTC)

var then2 = time.Date(
	2021, 11, 7, 11, 45, 07, 651387237, time.UTC)

var oldBuyingModel = model.BuyingEventRespondModel{
	ProjectId:                    "Test",
	ClientId:                     "Test",
	CustomerId:                   "Test",
	LevelIndex:                   5,
	TotalBuyingCount:             2,
	TotalBuyingDay:               2,
	FirstBuyingYearOfDay:         int64(then.YearDay()),
	FirstBuyingYear:              int64(then.Year()),
	FirstBuyingHour:              int64(then.Hour()),
	FirstBuyingMinute:            int64(buyingModel.InWhatMinutes),
	SecondBuyingYearOfDay:        0,
	SecondBuyingHour:             0,
	ThirdBuyingYearOfDay:         0,
	ThirdBuyingHour:              0,
	PenultimateBuyingYearOfDay:   0,
	PenultimateBuyingHour:        0,
	LastBuyingYearOfDay:          int64(then2.YearDay()),
	LastBuyingYear:               int64(then2.Year()),
	LastBuyingHour:               int64(then2.Hour()),
	LastBuyingMinute:             int64(buyingModel.InWhatMinutes),
	FirstDayBuyingCount:          10,
	PenultimateDayBuyingCount:    45,
	LastDayBuyingCount:           30,
	LastMinusFirstDayBuyingCount: -20,
	SundayBuyingCount:            8,
	MondayBuyingCount:            5,
	TuesdayBuyingCount:           2,
	WednesdayBuyingCount:         3,
	ThursdayBuyingCount:          4,
	FridayBuyingCount:            1,
	SaturdayBuyingCount:          7,
	AmBuyingCount:                6,
	PmBuyingCount:                3,
	Buying0To5HourCount:          2,
	Buying6To11HourCount:         9,
	Buying12To17HourCount:        0,
	Buying18To23HourCount:        11,
	BuyingDayAverageBuyingCount:  37,
	LevelBasedAverageBuyingCount: 45,
}

var TotalBuyingCount = newBuyingModel.TotalBuyingCount + oldBuyingModel.TotalBuyingCount
var TotalBuyingDay = (newBuyingModel.FirstBuyingYearOfDay+365*newBuyingModel.FirstBuyingYear) - (oldBuyingModel.FirstBuyingYearOfDay+365*oldBuyingModel.FirstBuyingYear)
var FirstDayBuyingCount = oldBuyingModel.FirstDayBuyingCount+ newBuyingModel.FirstDayBuyingCount
var updateBuyingModel = model.BuyingEventRespondModel{
	ProjectId:                    "Test",
	ClientId:                     "Test",
	CustomerId:                   "Test",
	LevelIndex:                   newBuyingModel.LevelIndex,
	TotalBuyingCount:             newBuyingModel.TotalBuyingCount + oldBuyingModel.TotalBuyingCount,
	TotalBuyingDay:               (newBuyingModel.FirstBuyingYearOfDay+365*newBuyingModel.FirstBuyingYear) - (oldBuyingModel.FirstBuyingYearOfDay+365*oldBuyingModel.FirstBuyingYear),
	FirstBuyingYearOfDay:         int64(then.YearDay()),
	FirstBuyingYear:              int64(then.Year()),
	FirstBuyingHour:              int64(then.Hour()),
	FirstBuyingMinute:            int64(buyingModel.InWhatMinutes),
	SecondBuyingYearOfDay:        0,
	SecondBuyingHour:             0,
	ThirdBuyingYearOfDay:         newBuyingModel.FirstBuyingYearOfDay,
	ThirdBuyingHour:              newBuyingModel.FirstBuyingHour,
	PenultimateBuyingYearOfDay:   int64(then2.YearDay()),
	PenultimateBuyingHour:        int64(then2.Hour()),
	LastBuyingYearOfDay:          newBuyingModel.FirstBuyingYearOfDay,
	LastBuyingYear:               newBuyingModel.FirstBuyingYear,
	LastBuyingHour:               newBuyingModel.FirstBuyingHour,
	LastBuyingMinute:             int64(buyingModel.InWhatMinutes),
	FirstDayBuyingCount:          oldBuyingModel.FirstDayBuyingCount+ newBuyingModel.FirstDayBuyingCount,
	PenultimateDayBuyingCount:    oldBuyingModel.PenultimateDayBuyingCount,
	LastDayBuyingCount:           oldBuyingModel.LastDayBuyingCount,
	LastMinusFirstDayBuyingCount: oldBuyingModel.LastDayBuyingCount-FirstDayBuyingCount,
	SundayBuyingCount:            newBuyingModel.SundayBuyingCount + oldBuyingModel.SundayBuyingCount,
	MondayBuyingCount:            newBuyingModel.MondayBuyingCount + oldBuyingModel.MondayBuyingCount,
	TuesdayBuyingCount:           newBuyingModel.TuesdayBuyingCount + oldBuyingModel.TuesdayBuyingCount,
	WednesdayBuyingCount:         newBuyingModel.WednesdayBuyingCount + oldBuyingModel.WednesdayBuyingCount,
	ThursdayBuyingCount:          newBuyingModel.ThursdayBuyingCount + oldBuyingModel.ThursdayBuyingCount,
	FridayBuyingCount:            newBuyingModel.FridayBuyingCount + oldBuyingModel.FridayBuyingCount,
	SaturdayBuyingCount:          newBuyingModel.SaturdayBuyingCount + oldBuyingModel.SaturdayBuyingCount,
	AmBuyingCount:                newBuyingModel.AmBuyingCount + oldBuyingModel.AmBuyingCount,
	PmBuyingCount:                newBuyingModel.PmBuyingCount + oldBuyingModel.PmBuyingCount,
	Buying0To5HourCount:          newBuyingModel.Buying0To5HourCount + oldBuyingModel.Buying0To5HourCount,
	Buying6To11HourCount:         newBuyingModel.Buying6To11HourCount + oldBuyingModel.Buying6To11HourCount,
	Buying12To17HourCount:        newBuyingModel.Buying12To17HourCount + oldBuyingModel.Buying12To17HourCount,
	Buying18To23HourCount:        newBuyingModel.Buying18To23HourCount + oldBuyingModel.Buying18To23HourCount,
	BuyingDayAverageBuyingCount:  float64(TotalBuyingCount) / float64(TotalBuyingDay),
	LevelBasedAverageBuyingCount: float64(TotalBuyingCount)/float64(newBuyingModel.LevelIndex),
}


func Test_UpdateBuyingEvent_UpdateSuccess(t *testing.T) {
	var testBuyingDal = new(repository.MockBuyingEventDal)

	var manager = concrete.BuyingEventManager{
		IBuyingEventDal: testBuyingDal,
		IJsonParser:     &gojson.goJson{},
	}
	
	testBuyingDal.On("UpdateBuyingEventById", updateBuyingModel.ClientId, &updateBuyingModel).Return(nil)

	var v, s, m = manager.UpdateBuyingEvent(&newBuyingModel, &oldBuyingModel)

	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
	assert.Equal(t, &updateBuyingModel, v)
}

func Test_ConvertRawModelToResponseModel_AddSuccess(t *testing.T) {
	var testBuyingDal = new(repository.MockBuyingEventDal)
	var manager = concrete.BuyingEventManager{
		IBuyingEventDal: testBuyingDal,
		IJsonParser:     &gojson.goJson{},
	}
	buyingByte, _ := manager.IJsonParser.EncodeJson(buyingModel)

	testBuyingDal.On("GetBuyingEventById", buyingModel.ClientId).Return(&oldBuyingModel,
		errors.New("mongo: no documents in result"))
	testBuyingDal.On("Add", &newBuyingModel).Return(nil)

	var v, s, m = manager.ConvertRawModelToResponseModel(buyingByte)

	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
	assert.Equal(t, &newBuyingModel, v)

}

func Test_CalculateBuyingLevelBasedAvgBuyingCount(t *testing.T) {
	concrete.CalculateBuyingLevelBasedAvgBuyingCount(&newBuyingModel)
	//levelAvg := float64(newBuyingModel.TotalBuyingCount) / float64(newBuyingModel.LevelIndex)
	levelAvg := float64(newBuyingModel.TotalBuyingCount)
	assert.Equal(t, levelAvg, newBuyingModel.LevelBasedAverageBuyingCount)
}

func Test_DetermineBuyingAmPm(t *testing.T) {
	concrete.DetermineBuyingAmPm(&newBuyingModel, int64(buyingModel.TrigerdTime.Hour())-10)
	assert.Equal(t, int64(1), newBuyingModel.AmBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel.PmBuyingCount)
}

func Test_DetermineBuyingHour(t *testing.T) {
	concrete.DetermineBuyingHour(&newBuyingModel, int64(buyingModel.TrigerdTime.Hour())+2)
	assert.Equal(t, int64(0), newBuyingModel.Buying0To5HourCount)
	assert.Equal(t, int64(0), newBuyingModel.Buying6To11HourCount)
	assert.Equal(t, int64(0), newBuyingModel.Buying12To17HourCount)
	assert.Equal(t, int64(1), newBuyingModel.Buying18To23HourCount)
}

func Test_DetermineBuyingDay(t *testing.T) {
	concrete.DetermineBuyingDay(&newBuyingModel, int64(buyingModel.TrigerdTime.Weekday()))
	assert.Equal(t, int64(0), newBuyingModel.SundayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel.MondayBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel.TuesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel.WednesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel.ThursdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel.FridayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel.SaturdayBuyingCount)
}

func Test_CalculateLastDayBuyingCount(t *testing.T) {
	var count int64 = concrete.CalculateLastDayBuyingCount(&newBuyingModel, &oldBuyingModel)
	var expCount int64 = oldBuyingModel.LastDayBuyingCount + newBuyingModel.FirstDayBuyingCount
	//var expCount int64 = 1
	assert.Equal(t, []int64{count}, []int64{expCount})
}

func Test_CalculateFirstDayBuyingCount(t *testing.T) {
	var count int64 = concrete.CalculateFirstDayBuyingCount(&newBuyingModel, &oldBuyingModel)
	var expCount int64 = oldBuyingModel.FirstDayBuyingCount + newBuyingModel.FirstDayBuyingCount
	assert.Equal(t, []int64{count}, []int64{expCount})
}

func Test_CalculatePenultimateDayBuyingCount(t *testing.T) {
	var count int64 = concrete.CalculatePenultimateDayBuyingCount(&newBuyingModel, &oldBuyingModel)
	var expCount int64 = oldBuyingModel.LastDayBuyingCount
	assert.Equal(t, []int64{count}, []int64{expCount})
}

func Test_calculateThirdBuying(t *testing.T) {
	//var day, hour int64 = concrete.CalculateThirdBuying(&newBuyingModel, &oldBuyingModel)
	concrete.CalculateThirdBuying(&newBuyingModel, &oldBuyingModel)
	var Expday int64 = int64(buyingModel.TrigerdTime.YearDay())
	var Exphour int64 = int64(buyingModel.TrigerdTime.Hour())
	// var Expday int64 =  int64(0)
	// var Exphour int64 = int64(0)
	//assert.Equal(t, []int64{Expday, Exphour}, []int64{day, hour})
	assert.Equal(t, []int64{Expday, Exphour}, []int64{newBuyingModel.ThirdBuyingYearOfDay, newBuyingModel.ThirdBuyingHour})
}

func Test_CalculateSecondBuying(t *testing.T) {
	var day, hour int64 = concrete.CalculateSecondBuying(&newBuyingModel, &oldBuyingModel)

	var Expday int64 = int64(buyingModel.TrigerdTime.YearDay())
	var Exphour int64 = int64(buyingModel.TrigerdTime.Hour())

	assert.Equal(t, []int64{Expday, Exphour}, []int64{day, hour})

}
