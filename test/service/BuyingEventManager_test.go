package test

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/Log"
	"FilterWorkerService/test/Mock/repository"
	"FilterWorkerService/test/Mock/service"
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
	TotalBuyingHour:              0,
	FirstBuyingYearOfDay:         int64(buyingModel.TrigerdTime.YearDay()),
	FirstBuyingYear:              int64(buyingModel.TrigerdTime.Year()),
	FirstBuyingHour:              int64(buyingModel.TrigerdTime.Hour()),
	FirstBuyingMinute:            int64(buyingModel.InWhatMinutes),
	FirstBuyingProductType:       0,
	SecondBuyingYearOfDay:        0,
	SecondBuyingHour:             0,
	SecondBuyingMinute:           0,
	SecondBuyingProductType:      0,
	ThirdBuyingYearOfDay:         0,
	ThirdBuyingHour:              0,
	ThirdBuyingMinute:            0,
	ThirdBuyingProductType:       0,
	FourthBuyingYearOfDay:        0,
	FourthBuyingHour:             0,
	FourthBuyingMinute:           0,
	FourthBuyingProductType:      0,
	FifthBuyingYearOfDay:         0,
	FifthBuyingHour:              0,
	FifthBuyingMinute:            0,
	FifthBuyingProductType:       0,
	PenultimateBuyingYearOfDay:   0,
	PenultimateBuyingHour:        0,
	PenultimateBuyingMinute:      0,
	PenultimateBuyingProductType: 0,
	LastBuyingYearOfDay:          0,
	LastBuyingYear:               0,
	LastBuyingHour:               0,
	LastBuyingMinute:             0,
	LastBuyingProductType:        0,
	FirstDayBuyingCount:          1,
	SecondDayBuyingCount:         0,
	ThirdDayBuyingCount:          0,
	FourthDayBuyingCount:         0,
	FifthDayBuyingCount:          0,
	SixthDayBuyingCount:          0,
	SeventhDayBuyingCount:        0,
	PenultimateDayBuyingCount:    0,
	LastDayBuyingCount:           0,
	LastMinusFirstDayBuyingCount: -1,
	SundayBuyingCount:            0,
	MondayBuyingCount:            0,
	TuesdayBuyingCount:           0,
	WednesdayBuyingCount:         0,
	ThursdayBuyingCount:          0,
	FridayBuyingCount:            0,
	SaturdayBuyingCount:          0,
	AmBuyingCount:                0,
	PmBuyingCount:                0,
	Buying0To5HourCount:          0,
	Buying6To11HourCount:         0,
	Buying12To17HourCount:        0,
	Buying18To23HourCount:        0,
	BuyingDayAverageBuyingCount:  1,
	LevelBasedAverageBuyingCount: 0,
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
	TotalBuyingHour:              0,
	FirstBuyingYearOfDay:         int64(then.YearDay()),
	FirstBuyingYear:              int64(then.Year()),
	FirstBuyingHour:              int64(then.Hour()),
	FirstBuyingMinute:            int64(buyingModel.InWhatMinutes),
	FirstBuyingProductType:       0,
	SecondBuyingYearOfDay:        0,
	SecondBuyingHour:             0,
	SecondBuyingMinute:           0,
	SecondBuyingProductType:      0,
	ThirdBuyingYearOfDay:         0,
	ThirdBuyingHour:              0,
	ThirdBuyingMinute:            0,
	ThirdBuyingProductType:       0,
	FourthBuyingYearOfDay:        0,
	FourthBuyingHour:             0,
	FourthBuyingMinute:           0,
	FourthBuyingProductType:      0,
	FifthBuyingYearOfDay:         0,
	FifthBuyingHour:              0,
	FifthBuyingMinute:            0,
	FifthBuyingProductType:       0,
	PenultimateBuyingYearOfDay:   0,
	PenultimateBuyingHour:        0,
	PenultimateBuyingMinute:      0,
	PenultimateBuyingProductType: 0,
	LastBuyingYearOfDay:          int64(then2.YearDay()),
	LastBuyingYear:               int64(then2.Year()),
	LastBuyingHour:               int64(then2.Hour()),
	LastBuyingMinute:             int64(buyingModel.InWhatMinutes),
	LastBuyingProductType:        0,
	FirstDayBuyingCount:          10,
	SecondDayBuyingCount:         0,
	ThirdDayBuyingCount:          0,
	FourthDayBuyingCount:         0,
	FifthDayBuyingCount:          0,
	SixthDayBuyingCount:          0,
	SeventhDayBuyingCount:        0,
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
var TotalBuyingDay = (newBuyingModel.FirstBuyingYearOfDay + 365*newBuyingModel.FirstBuyingYear) - (oldBuyingModel.FirstBuyingYearOfDay + 365*oldBuyingModel.FirstBuyingYear)
var FirstDayBuyingCount = oldBuyingModel.FirstDayBuyingCount + newBuyingModel.FirstDayBuyingCount

var updateBuyingModel = model.BuyingEventRespondModel{
	ProjectId:                    "Test",
	ClientId:                     "Test",
	CustomerId:                   "Test",
	LevelIndex:                   newBuyingModel.LevelIndex,
	TotalBuyingCount:             newBuyingModel.TotalBuyingCount + oldBuyingModel.TotalBuyingCount,
	TotalBuyingDay:               (newBuyingModel.FirstBuyingYearOfDay + 365*newBuyingModel.FirstBuyingYear) - (oldBuyingModel.FirstBuyingYearOfDay + 365*oldBuyingModel.FirstBuyingYear),
	TotalBuyingHour:              0,
	FirstBuyingYearOfDay:         int64(then.YearDay()),
	FirstBuyingYear:              int64(then.Year()),
	FirstBuyingHour:              int64(then.Hour()),
	FirstBuyingMinute:            int64(buyingModel.InWhatMinutes),
	FirstBuyingProductType:       0,
	SecondBuyingYearOfDay:        0,
	SecondBuyingHour:             0,
	SecondBuyingMinute:           0,
	SecondBuyingProductType:      0,
	ThirdBuyingYearOfDay:         newBuyingModel.FirstBuyingYearOfDay,
	ThirdBuyingHour:              newBuyingModel.FirstBuyingHour,
	ThirdBuyingMinute:            0,
	ThirdBuyingProductType:       0,
	FourthBuyingYearOfDay:        0,
	FourthBuyingHour:             0,
	FourthBuyingMinute:           0,
	FourthBuyingProductType:      0,
	FifthBuyingYearOfDay:         0,
	FifthBuyingHour:              0,
	FifthBuyingMinute:            0,
	FifthBuyingProductType:       0,
	PenultimateBuyingYearOfDay:   int64(then2.YearDay()),
	PenultimateBuyingHour:        int64(then2.Hour()),
	PenultimateBuyingMinute:      0,
	PenultimateBuyingProductType: 0,
	LastBuyingYearOfDay:          newBuyingModel.FirstBuyingYearOfDay,
	LastBuyingYear:               newBuyingModel.FirstBuyingYear,
	LastBuyingHour:               newBuyingModel.FirstBuyingHour,
	LastBuyingMinute:             int64(buyingModel.InWhatMinutes),
	LastBuyingProductType:        0,
	FirstDayBuyingCount:          oldBuyingModel.FirstDayBuyingCount + newBuyingModel.FirstDayBuyingCount,
	SecondDayBuyingCount:         0,
	ThirdDayBuyingCount:          0,
	FourthDayBuyingCount:         0,
	FifthDayBuyingCount:          0,
	SixthDayBuyingCount:          0,
	SeventhDayBuyingCount:        0,
	PenultimateDayBuyingCount:    oldBuyingModel.PenultimateDayBuyingCount,
	LastDayBuyingCount:           oldBuyingModel.LastDayBuyingCount,
	LastMinusFirstDayBuyingCount: oldBuyingModel.LastDayBuyingCount - FirstDayBuyingCount,
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
	LevelBasedAverageBuyingCount: float64(TotalBuyingCount) / float64(newBuyingModel.LevelIndex),
}

func Test_UpdateBuyingEvent_UpdateSuccess(t *testing.T) {

	var testBuyingDal = new(repository.MockBuyingEventDal)
	var testCache = new(service.MockCacheService)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.BuyingEventDal = testBuyingDal
	IoC.Logger = testLog
	IoC.CacheService = testCache
	var manager = concrete.BuyingEventManagerConstructor()
	//var buyingModel_test = buyingModel
	var oldBuyingModel_test = oldBuyingModel
	var newBuyingModel_test = newBuyingModel
	var updateBuyingModel_test = updateBuyingModel

	testBuyingDal.On("UpdateBuyingEventById", updateBuyingModel_test.ClientId, &updateBuyingModel_test).Return(nil)

	var v, s, m = manager.UpdateBuyingEvent(&newBuyingModel_test, &oldBuyingModel_test)

	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
	assert.Equal(t, &updateBuyingModel_test, v)
}


func Test_ConvertRawModelToResponseModel_AddSuccess(t *testing.T) {
	var testBuyingDal = new(repository.MockBuyingEventDal)
	var testCache = new(service.MockCacheService)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.BuyingEventDal = testBuyingDal
	IoC.Logger = testLog
	IoC.CacheService = testCache
	var manager = concrete.BuyingEventManagerConstructor()
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 5, 16, 11, 36, 651387237, time.UTC)
	var oldBuyingModel_test = oldBuyingModel
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FridayBuyingCount = 1
	newBuyingModel_test.PmBuyingCount = 1
	newBuyingModel_test.Buying12To17HourCount = 1
	newBuyingModel_test.FirstBuyingProductType = 19
	newBuyingModel_test.FirstBuyingYearOfDay  =       int64(buyingModel_test.TrigerdTime.YearDay())
	newBuyingModel_test.FirstBuyingYear       =       int64(buyingModel_test.TrigerdTime.Year())
	newBuyingModel_test.FirstBuyingHour       =       int64(buyingModel_test.TrigerdTime.Hour())
	newBuyingModel_test.FirstBuyingMinute     =       int64(buyingModel_test.InWhatMinutes)
	testBuyingDal.On("GetBuyingEventById", newBuyingModel_test.ClientId).Return(&oldBuyingModel_test,
		errors.New("null data error"))
	testCache.On("ManageCache", "ProductType", buyingModel_test.ProductType).Return(int64(19), true, "")
	testBuyingDal.On("Add", &newBuyingModel_test).Return(nil)
	var buyingModel_byte, _ = json.EncodeJson(buyingModel_test)
	var v, s, m = manager.ConvertRawModelToResponseModel(buyingModel_byte)

	var value, success = v.(model.BuyingEventRespondModel)
	if success == true {
		assert.Equal(t, &newBuyingModel_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
	//assert.Equal(t, &newBuyingModel, v)

}

func Test_CalculateBuyingLevelBasedAvgBuyingCount_ZeroLevelIndex(t *testing.T) {
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.LevelIndex = 0
	concrete.CalculateBuyingLevelBasedAvgBuyingCount(&newBuyingModel_test)
	levelAvg := float64(newBuyingModel.TotalBuyingCount)
	assert.Equal(t, levelAvg, newBuyingModel.LevelBasedAverageBuyingCount)
}

func Test_CalculateBuyingLevelBasedAvgBuyingCount_DifferZeroLevelIndex(t *testing.T) {
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 20
	newBuyingModel_test.LevelIndex = 5
	concrete.CalculateBuyingLevelBasedAvgBuyingCount(&newBuyingModel_test)
	levelAvg := float64(newBuyingModel.TotalBuyingCount) / float64(newBuyingModel.LevelIndex)
	assert.Equal(t, levelAvg, newBuyingModel.LevelBasedAverageBuyingCount)
}

func Test_DetermineBuyingAmPm_Am(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 12, 9, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingAmPm(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Hour()))
	assert.Equal(t, int64(1), newBuyingModel_test.AmBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.PmBuyingCount)
}

func Test_DetermineBuyingAmPm_Pm(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 12, 15, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingAmPm(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Hour()))
	assert.Equal(t, int64(0), newBuyingModel_test.AmBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel_test.PmBuyingCount)
}

func Test_DetermineBuyingHour_05(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 12, 4, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingHour(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Hour()))
	assert.Equal(t, int64(1), newBuyingModel_test.Buying0To5HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying6To11HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying12To17HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying18To23HourCount)
}

func Test_DetermineBuyingHour_611(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 12, 8, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingHour(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Hour()))
	assert.Equal(t, int64(0), newBuyingModel_test.Buying0To5HourCount)
	assert.Equal(t, int64(1), newBuyingModel_test.Buying6To11HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying12To17HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying18To23HourCount)
}

func Test_DetermineBuyingHour_1217(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 12, 16, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingHour(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Hour()))
	assert.Equal(t, int64(0), newBuyingModel_test.Buying0To5HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying6To11HourCount)
	assert.Equal(t, int64(1), newBuyingModel_test.Buying12To17HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying18To23HourCount)
}

func Test_DetermineBuyingHour_1823(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 21, 22, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingHour(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Hour()))
	assert.Equal(t, int64(0), newBuyingModel_test.Buying0To5HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying6To11HourCount)
	assert.Equal(t, int64(0), newBuyingModel_test.Buying12To17HourCount)
	assert.Equal(t, int64(1), newBuyingModel_test.Buying18To23HourCount)
}

func Test_DetermineBuyingDay_Sunday(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 21, 16, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingDay(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Weekday()))
	assert.Equal(t, int64(1), newBuyingModel_test.SundayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.MondayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.TuesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.WednesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.ThursdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.FridayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.SaturdayBuyingCount)
}

func Test_DetermineBuyingDay_Monday(t *testing.T) {
	var buyingModel_test = buyingModel
	var newBuyingModel_test = newBuyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 22, 16, 11, 36, 651387237, time.UTC)
	concrete.DetermineBuyingDay(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Weekday()))
	assert.Equal(t, int64(0), newBuyingModel_test.SundayBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel_test.MondayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.TuesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.WednesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.ThursdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.FridayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.SaturdayBuyingCount)
}

func Test_DetermineBuyingDay_Tuesday(t *testing.T) {
	var buyingModel_test = buyingModel
	var newBuyingModel_test = newBuyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 23, 16, 11, 36, 651387237, time.UTC)
	concrete.DetermineBuyingDay(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Weekday()))
	assert.Equal(t, int64(0), newBuyingModel_test.SundayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.MondayBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel_test.TuesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.WednesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.ThursdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.FridayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.SaturdayBuyingCount)
}

func Test_DetermineBuyingDay_Wednesday(t *testing.T) {
	var buyingModel_test = buyingModel
	var newBuyingModel_test = newBuyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 24, 16, 11, 36, 651387237, time.UTC)
	concrete.DetermineBuyingDay(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Weekday()))
	assert.Equal(t, int64(0), newBuyingModel_test.SundayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.MondayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.TuesdayBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel_test.WednesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.ThursdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.FridayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.SaturdayBuyingCount)
}

func Test_DetermineBuyingDay_Thursday(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 25, 16, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingDay(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Weekday()))
	assert.Equal(t, int64(0), newBuyingModel_test.SundayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.MondayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.TuesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.WednesdayBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel_test.ThursdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.FridayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.SaturdayBuyingCount)
}

func Test_DetermineBuyingDay_Friday(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 26, 16, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingDay(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Weekday()))
	assert.Equal(t, int64(0), newBuyingModel_test.SundayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.MondayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.TuesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.WednesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.ThursdayBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel_test.FridayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.SaturdayBuyingCount)
}

func Test_DetermineBuyingDay_Saturday(t *testing.T) {
	var buyingModel_test = buyingModel
	buyingModel_test.TrigerdTime = time.Date(
		2021, 11, 27, 16, 11, 36, 651387237, time.UTC)
	var newBuyingModel_test = newBuyingModel
	concrete.DetermineBuyingDay(&newBuyingModel_test, int64(buyingModel_test.TrigerdTime.Weekday()))
	assert.Equal(t, int64(0), newBuyingModel_test.SundayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.MondayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.TuesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.WednesdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.ThursdayBuyingCount)
	assert.Equal(t, int64(0), newBuyingModel_test.FridayBuyingCount)
	assert.Equal(t, int64(1), newBuyingModel_test.SaturdayBuyingCount)
}
func Test_CalculateLastDayBuyingCount_NewModeldayEqualOldModelLastDay(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.LastBuyingYearOfDay = 120
	oldBuyingModel_test.LastBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingYearOfDay = 19
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.LastDayBuyingCount = 2
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 120
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstDayBuyingCount = 1
	var count int64 = concrete.CalculateLastDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test)
	var expCount int64 = oldBuyingModel_test.LastDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	assert.Equal(t, []int64{count}, []int64{expCount})
}

func Test_CalculateLastDayBuyingCount_NewModeldayBiggerOldModelLastDay(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.LastBuyingYearOfDay = 120
	oldBuyingModel_test.LastBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingYearOfDay = 19
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.LastDayBuyingCount = 2
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 122
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstDayBuyingCount = 1
	var count int64 = concrete.CalculateLastDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test)
	var expCount int64 = newBuyingModel_test.FirstDayBuyingCount
	assert.Equal(t, []int64{count}, []int64{expCount})
}

func Test_CalculateLastDayBuyingCount_NewModeldayLowerOldModelLastDay(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.LastBuyingYearOfDay = 120
	oldBuyingModel_test.LastBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingYearOfDay = 19
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.LastDayBuyingCount = 2
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 119
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstDayBuyingCount = 1
	var count int64 = concrete.CalculateLastDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test)
	var expCount int64 = oldBuyingModel_test.LastDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateFirstDayBuyingCount_Lower24Hours(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 301
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 13
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateFirstDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.FirstDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	assert.Equal(t, oldBuyingModel_test.FirstDayBuyingCount, expCount)
}

func Test_CalculateFirstDayBuyingCount_Bigger24Hours(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 301
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateFirstDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.FirstDayBuyingCount
	assert.Equal(t, oldBuyingModel_test.FirstDayBuyingCount, expCount)
}

func Test_CalculateSecondDayBuyingCount_In24Between48(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 301
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateSecondDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.SecondDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	assert.Equal(t, oldBuyingModel_test.SecondDayBuyingCount, expCount)
}

func Test_CalculateSecondDayBuyingCount_Out24Between48(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 302
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateSecondDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.SecondDayBuyingCount
	assert.Equal(t, oldBuyingModel_test.SecondDayBuyingCount, expCount)
}

func Test_CalculateThirdDayBuyingCount_In48Between72(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 302
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 20
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateThirdDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.ThirdDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.ThirdDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateThirdDayBuyingCount_Out48Between72(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 303
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateThirdDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.ThirdDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.ThirdDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateFourthDayBuyingCount_In72Between96(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 303
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateFourthDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.FourthDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.FourthDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateFourthDayBuyingCount_Out72Between96(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 304
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateFourthDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.FourthDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.FourthDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateFifthDayBuyingCount_In96Between120(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 305
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 9
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateFifthDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.FifthDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.FifthDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateFifthDayBuyingCount_Out96Between120(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 305
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 16
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateFifthDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.FifthDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.FifthDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateSixthDayBuyingCount_In120Between144(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 305
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 16
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateSixthDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.SixthDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.SixthDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateSixthDayBuyingCount_Out120Between144(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 306
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 16
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateSixthDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.SixthDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.SixthDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateSevenhDayBuyingCount_In144Between168(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 306
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateSeventhDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.SeventhDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.SeventhDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateSeventhDayBuyingCount_Out144Between168(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.FirstBuyingYearOfDay = 300
	oldBuyingModel_test.FirstBuyingYear = 2021
	oldBuyingModel_test.FirstBuyingHour = 14
	oldBuyingModel_test.FirstDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 307
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 15
	newBuyingModel_test.FirstDayBuyingCount = 1
	var TotalBuyingHour = ((newBuyingModel_test.FirstBuyingYearOfDay+365*newBuyingModel_test.FirstBuyingYear)*24 + newBuyingModel_test.FirstBuyingHour) - ((oldBuyingModel_test.FirstBuyingYearOfDay+365*oldBuyingModel_test.FirstBuyingYear)*24 + oldBuyingModel_test.FirstBuyingHour)
	concrete.CalculateSeventhDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test, TotalBuyingHour)
	var expCount int64 = oldBuyingModel_test.SeventhDayBuyingCount + newBuyingModel_test.FirstDayBuyingCount
	var count = oldBuyingModel_test.SeventhDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculatePenultimateDayBuyingCount_NewModeldayEqualOldModelLastDay(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.LastBuyingYearOfDay = 300
	oldBuyingModel_test.LastBuyingYear = 2021
	oldBuyingModel_test.LastBuyingHour = 14
	oldBuyingModel_test.LastDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 300
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 19
	newBuyingModel_test.FirstDayBuyingCount = 1
	var count int64 = concrete.CalculatePenultimateDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test)
	//var expCount int64 = oldBuyingModel_test.LastDayBuyingCount
	var expCount int64 = oldBuyingModel_test.PenultimateDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculatePenultimateDayBuyingCount_NewModeldayBiggerOldModelLastDay(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.LastBuyingYearOfDay = 300
	oldBuyingModel_test.LastBuyingYear = 2021
	oldBuyingModel_test.LastBuyingHour = 14
	oldBuyingModel_test.LastDayBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.FirstBuyingYearOfDay = 301
	newBuyingModel_test.FirstBuyingYear = 2021
	newBuyingModel_test.FirstBuyingHour = 19
	newBuyingModel_test.FirstDayBuyingCount = 1
	var count int64 = concrete.CalculatePenultimateDayBuyingCount(&newBuyingModel_test, &oldBuyingModel_test)
	var expCount int64 = oldBuyingModel_test.LastDayBuyingCount
	//var expCount int64 = oldBuyingModel_test.PenultimateDayBuyingCount
	assert.Equal(t, count, expCount)
}

func Test_CalculateSecondBuying_TotalBuyingCountEqual2(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 1
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 355
	newBuyingModel_test.FirstBuyingHour = 23
	newBuyingModel_test.FirstBuyingMinute = 19
	newBuyingModel_test.FirstBuyingProductType = 3
	concrete.CalculateSecondBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.SecondBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.SecondBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.SecondBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.SecondBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 3558, oldBuyingModel_test.SecondBuyingYearOfDay)
	assert.Equal(t, 238, oldBuyingModel_test.SecondBuyingHour)
	assert.Equal(t, 198, oldBuyingModel_test.SecondBuyingMinute)
	assert.Equal(t, 38, oldBuyingModel_test.SecondBuyingProductType)

}

func Test_CalculateSecondBuying_TotalBuyingCountNotEqual2(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 2
	oldBuyingModel_test.SecondBuyingYearOfDay = 248
	oldBuyingModel_test.SecondBuyingHour = 16
	oldBuyingModel_test.SecondBuyingMinute = 36
	oldBuyingModel_test.SecondBuyingProductType = 7
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 355
	newBuyingModel_test.FirstBuyingHour = 23
	newBuyingModel_test.FirstBuyingMinute = 48
	newBuyingModel_test.FirstBuyingProductType = 3

	concrete.CalculateSecondBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.SecondBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.SecondBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.SecondBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.SecondBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 248, oldBuyingModel_test.SecondBuyingYearOfDay)
	assert.Equal(t, 16, oldBuyingModel_test.SecondBuyingHour)
	assert.Equal(t, 36, oldBuyingModel_test.SecondBuyingMinute)
	assert.Equal(t, 7, oldBuyingModel_test.SecondBuyingProductType)

}

func Test_CalculateThirdBuying_TotalBuyingCountEqual3(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 2
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 146
	newBuyingModel_test.FirstBuyingHour = 8
	newBuyingModel_test.FirstBuyingMinute = 48
	newBuyingModel_test.FirstBuyingProductType = 5

	concrete.CalculateThirdBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.ThirdBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.ThirdBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.ThirdBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.ThirdBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 146, oldBuyingModel_test.ThirdBuyingYearOfDay)
	assert.Equal(t, 8, oldBuyingModel_test.ThirdBuyingHour)
	assert.Equal(t, 48, oldBuyingModel_test.ThirdBuyingMinute)
	assert.Equal(t, 5, oldBuyingModel_test.ThirdBuyingProductType)

}

func Test_CalculateThirdBuying_TotalBuyingCountNotEqual3(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 1
	oldBuyingModel_test.ThirdBuyingYearOfDay = 333
	oldBuyingModel_test.ThirdBuyingHour = 21
	oldBuyingModel_test.ThirdBuyingMinute = 36
	oldBuyingModel_test.ThirdBuyingProductType = 7
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 355
	newBuyingModel_test.FirstBuyingHour = 23
	newBuyingModel_test.FirstBuyingMinute = 48
	newBuyingModel_test.FirstBuyingProductType = 3

	concrete.CalculateThirdBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.ThirdBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.ThirdBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.ThirdBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.ThirdBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 333, oldBuyingModel_test.ThirdBuyingYearOfDay)
	assert.Equal(t, 21, oldBuyingModel_test.ThirdBuyingHour)
	assert.Equal(t, 36, oldBuyingModel_test.ThirdBuyingMinute)
	assert.Equal(t, 7, oldBuyingModel_test.ThirdBuyingProductType)

}

func Test_CalculateFourthBuying_TotalBuyingCountEqual4(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 3
	oldBuyingModel_test.FourthBuyingYearOfDay = 248
	oldBuyingModel_test.FourthBuyingHour = 6
	oldBuyingModel_test.FourthBuyingMinute = 36
	oldBuyingModel_test.FourthBuyingProductType = 7
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 355
	newBuyingModel_test.FirstBuyingHour = 23
	newBuyingModel_test.FirstBuyingMinute = 48
	newBuyingModel_test.FirstBuyingProductType = 3
	concrete.CalculateFourthBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.FourthBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.FourthBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.FourthBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.FourthBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 355, oldBuyingModel_test.FourthBuyingYearOfDay)
	assert.Equal(t, 23, oldBuyingModel_test.FourthBuyingHour)
	assert.Equal(t, 48, oldBuyingModel_test.FourthBuyingMinute)
	assert.Equal(t, 3, oldBuyingModel_test.FourthBuyingProductType)

}

func Test_CalculateFourthBuying_TotalBuyingCountNotEqual4(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 2
	oldBuyingModel_test.FifthBuyingYearOfDay = 248
	oldBuyingModel_test.FifthBuyingHour = 6
	oldBuyingModel_test.FifthBuyingMinute = 36
	oldBuyingModel_test.FifthBuyingProductType = 7
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 355
	newBuyingModel_test.FirstBuyingHour = 23
	newBuyingModel_test.FirstBuyingMinute = 48
	newBuyingModel_test.FirstBuyingProductType = 3

	concrete.CalculateFourthBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.FourthBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.FourthBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.FourthBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.FourthBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 248, oldBuyingModel_test.FourthBuyingYearOfDay)
	assert.Equal(t, 6, oldBuyingModel_test.FourthBuyingHour)
	assert.Equal(t, 36, oldBuyingModel_test.FourthBuyingMinute)
	assert.Equal(t, 7, oldBuyingModel_test.FourthBuyingProductType)

}

func Test_CalculateFifthBuying_TotalBuyingCountEqual5(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 4
	oldBuyingModel_test.FourthBuyingYearOfDay = 248
	oldBuyingModel_test.FourthBuyingHour = 6
	oldBuyingModel_test.FourthBuyingMinute = 36
	oldBuyingModel_test.FourthBuyingProductType = 7
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 355
	newBuyingModel_test.FirstBuyingHour = 23
	newBuyingModel_test.FirstBuyingMinute = 48
	newBuyingModel_test.FirstBuyingProductType = 3

	concrete.CalculateFifthBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.FifthBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.FifthBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.FifthBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.FifthBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 355, oldBuyingModel_test.FifthBuyingYearOfDay)
	assert.Equal(t, 23, oldBuyingModel_test.FifthBuyingHour)
	assert.Equal(t, 48, oldBuyingModel_test.FifthBuyingMinute)
	assert.Equal(t, 3, oldBuyingModel_test.FifthBuyingProductType)

}

func Test_CalculateFifthBuying_TotalBuyingCountNotEqual5(t *testing.T) {
	var oldBuyingModel_test = oldBuyingModel
	oldBuyingModel_test.TotalBuyingCount = 5
	oldBuyingModel_test.FifthBuyingYearOfDay =   248
	oldBuyingModel_test.FifthBuyingHour =         6
	oldBuyingModel_test.FifthBuyingMinute =       36
	oldBuyingModel_test.FifthBuyingProductType = 7
	var newBuyingModel_test = newBuyingModel
	newBuyingModel_test.TotalBuyingCount = 1
	newBuyingModel_test.FirstBuyingYearOfDay = 355
	newBuyingModel_test.FirstBuyingHour = 23
	newBuyingModel_test.FirstBuyingMinute = 48
	newBuyingModel_test.FirstBuyingProductType = 3

	concrete.CalculateFifthBuying(&newBuyingModel_test, &oldBuyingModel_test)

	// oldBuyingModel_test.FourthBuyingYearOfDay =   newBuyingModel_test.FirstBuyingYearOfDay
	// oldBuyingModel_test.FourthBuyingHour =        newBuyingModel_test.FirstBuyingHour
	// oldBuyingModel_test.FourthBuyingMinute =      newBuyingModel_test.FirstBuyingMinute
	// oldBuyingModel_test.FourthBuyingProductType = newBuyingModel_test.FirstBuyingProductType
	assert.Equal(t, 248, oldBuyingModel_test.FifthBuyingYearOfDay)
	assert.Equal(t,  6, oldBuyingModel_test.FifthBuyingHour)
	assert.Equal(t,  36, oldBuyingModel_test.FifthBuyingMinute)
	assert.Equal(t, 7, oldBuyingModel_test.FifthBuyingProductType)

}
