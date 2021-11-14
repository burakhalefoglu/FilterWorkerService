package test

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"FilterWorkerService/test/Mock/service"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var newModel2 = model.AdvEventRespondModel{
	ProjectId:                            "1",
	ClientId:                             "11",
	CustomerId:                           "111",
	LevelIndex:                           0,
	TotalAdvDay:                          43,
	TotalAdvCount:                        100,
	LevelBasedAverageAdvCount:            200,
	AverageAdvDailyClickCount:            300,
	FirstAdvYearOfDay:                    400,
	FirstAdvYear:                         500,
	FirstAdvClickHour:                    600,
	FirstADvClickMinute:                  700,
	FirstAdvType:                         800,
	SecondAdvYearOfDay:                   900,
	SecondAdvHour:                        1000,
	SecondAdvMinute:                      1100,
	ThirdAdvYearOfDay:                    1200,
	ThirdAdvHour:                         1300,
	ThirdAdvMinute:                       1400,
	PenultimateAdvYearOfDay:              1500,
	PenultimateAdvHour:                   1600,
	PenultimateAdvMinute:                 1700,
	LastAdvYearOfDay:                     1800,
	LastAdvYear:                          1900,
	LastAdvClickHour:                     2000,
	LastAdvClickMinute:                   2100,
	LastAdvType:                          2200,
	FirstDayAdvClickCount:                2300,
	PenultimateDayAdvClickCount:          2400,
	LastDayAdvClickCount:                 2500,
	LastMinusFirstDayAdvClickCount:       2600,
	LastMinusPenultimateDayAdvClickCount: 2700,
	LastDayAdvClickCountMinusAverageDailyAdvClickCount: 2800,
	SundayAdvClickCount:     0,
	MondayAdvClickCount:     0,
	TuesdayAdvClickCount:    0,
	WednesdayAdvClickCount:  0,
	ThursdayAdvClickCount:   0,
	FridayAdvClickCount:     0,
	SaturdayAdvClickCount:   0,
	AmAdvClickCount:         0,
	PmAdvClickCount:         0,
	AdvClick0To5HourCount:   0,
	AdvClick6To11HourCount:  0,
	AdvClick12To17HourCount: 0,
	AdvClick18To23HourCount: 0,
}
var oldModel = model.AdvEventRespondModel{
	ProjectId:                            "1",
	ClientId:                             "11",
	CustomerId:                           "111",
	LevelIndex:                           42,
	TotalAdvDay:                          43,
	TotalAdvCount:                        3,
	LevelBasedAverageAdvCount:            2,
	AverageAdvDailyClickCount:            3,
	FirstAdvYearOfDay:                    400,
	FirstAdvYear:                         5,
	FirstAdvClickHour:                    6,
	FirstADvClickMinute:                  7,
	FirstAdvType:                         8,
	SecondAdvYearOfDay:                   9,
	SecondAdvHour:                        10,
	SecondAdvMinute:                      11,
	ThirdAdvYearOfDay:                    12,
	ThirdAdvHour:                         13,
	ThirdAdvMinute:                       14,
	PenultimateAdvYearOfDay:              15,
	PenultimateAdvHour:                   16,
	PenultimateAdvMinute:                 17,
	LastAdvYearOfDay:                     18,
	LastAdvYear:                          19,
	LastAdvClickHour:                     20,
	LastAdvClickMinute:                   21,
	LastAdvType:                          22,
	FirstDayAdvClickCount:                23,
	PenultimateDayAdvClickCount:          24,
	LastDayAdvClickCount:                 25,
	LastMinusFirstDayAdvClickCount:       26,
	LastMinusPenultimateDayAdvClickCount: 27,
	LastDayAdvClickCountMinusAverageDailyAdvClickCount: 28,
	SundayAdvClickCount:     29,
	MondayAdvClickCount:     30,
	TuesdayAdvClickCount:    31,
	WednesdayAdvClickCount:  32,
	ThursdayAdvClickCount:   33,
	FridayAdvClickCount:     34,
	SaturdayAdvClickCount:   35,
	AmAdvClickCount:         36,
	PmAdvClickCount:         37,
	AdvClick0To5HourCount:   38,
	AdvClick6To11HourCount:  39,
	AdvClick12To17HourCount: 40,
	AdvClick18To23HourCount: 41,
}

func Test_ConvertRawModelToResponse_AddedSuccess(t *testing.T) {

	//Arrance
	var testAdv = new(repository.MockAdvEventDal)
	var testCache = new(service.MockCacheService)
	var manager = concrete.AdvEventManager{
		IAdvEventDal:  testAdv,
		IJsonParser:   &gojson.GoJson{},
		ICacheService: testCache,
	}
	var advModel = model.AdvEventModel{
		ProjectId:   "1",
		ClientId:    "11",
		CustomerId:  "111",
		LevelName:   "1",
		LevelIndex:  1,
		AdvType:     "test",
		InMinutes:   12,
		TrigerdTime: time.Now(),
	}
	var responseModel = model.AdvEventRespondModel{
		ProjectId:                            "1",
		ClientId:                             "11",
		CustomerId:                           "111",
		LevelIndex:                           1,
		TotalAdvDay:                          1,
		TotalAdvCount:                        1,
		LevelBasedAverageAdvCount:            1,
		AverageAdvDailyClickCount:            1,
		FirstAdvYearOfDay:                    318,
		FirstAdvYear:                         2021,
		FirstAdvClickHour:                    21,
		FirstADvClickMinute:                  int64(advModel.TrigerdTime.Minute()),
		FirstAdvType:                         1,
		SecondAdvYearOfDay:                   0,
		SecondAdvHour:                        0,
		SecondAdvMinute:                      0,
		ThirdAdvYearOfDay:                    0,
		ThirdAdvHour:                         0,
		ThirdAdvMinute:                       0,
		PenultimateAdvYearOfDay:              0,
		PenultimateAdvHour:                   0,
		PenultimateAdvMinute:                 0,
		LastAdvYearOfDay:                     318,
		LastAdvYear:                          2021,
		LastAdvClickHour:                     21,
		LastAdvClickMinute:                   int64(advModel.TrigerdTime.Minute()),
		LastAdvType:                          1,
		FirstDayAdvClickCount:                1,
		PenultimateDayAdvClickCount:          0,
		LastDayAdvClickCount:                 1,
		LastMinusFirstDayAdvClickCount:       0,
		LastMinusPenultimateDayAdvClickCount: 0,
		LastDayAdvClickCountMinusAverageDailyAdvClickCount: 0,
		SundayAdvClickCount:     1,
		MondayAdvClickCount:     0,
		TuesdayAdvClickCount:    0,
		WednesdayAdvClickCount:  0,
		ThursdayAdvClickCount:   0,
		FridayAdvClickCount:     0,
		SaturdayAdvClickCount:   0,
		AmAdvClickCount:         0,
		PmAdvClickCount:         1,
		AdvClick0To5HourCount:   0,
		AdvClick6To11HourCount:  0,
		AdvClick12To17HourCount: 0,
		AdvClick18To23HourCount: 1,
	}
	var message, _ = manager.IJsonParser.EncodeJson(&advModel)

	testCache.On("ManageCache", "AdvType", advModel.AdvType).Return(int64(1), true, "")
	testCache.On("ManageCache", "AdvType", advModel.AdvType).Return(int64(1), true, "")
	testAdv.On("GetAdvEventById", advModel.ClientId).Return(&responseModel,
		errors.New("mongo: no documents in result"))

	testAdv.On("Add", &responseModel).Return(nil)
	//Act
	var v, s, m = manager.ConvertRawModelToResponseModel(message)

	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
	assert.Equal(t, &responseModel, v)
}


func Test_CalculateSecondAdv_DateConversionSuccess(t *testing.T) {

	day, hour, minute := concrete.CalculateSecondAdv(&newModel2, &oldModel)

	var Expday int64 = 9
	var Exphour int64 = 10
	var ExpMinute int64 = 11

	assert.Equal(t, []int64{Expday, Exphour, ExpMinute}, []int64{day, hour, minute})

}

func TestCalculateThirdAdv(t *testing.T) {

	var Expday int64 = 400
	var Exphour int64 = 600
	var ExpMinute int64 = 700
	day, hour, minute := concrete.CalculateThirdAdv(&newModel2, &oldModel)

	assert.Equal(t, []int64{Expday, Exphour, ExpMinute}, []int64{day, hour, minute})

}

func TestCalculateFirstDayAdvClickCount(t *testing.T) {

	var exp int64 = 23
	num := concrete.CalculateFirstDayAdvClickCount(&newModel2, &oldModel)

	assert.Equal(t, []int64{exp}, []int64{num})
}

func TestCalculatePenultimateDayAdvDay(t *testing.T) {

	var exp int64 = 24
	num := concrete.CalculatePenultimateDayAdvDay(&newModel2, &oldModel)

	assert.Equal(t, []int64{exp}, []int64{num})
}

func TestCalculateLastDayAdvClickCount(t *testing.T) {
	var exp int64 = 25
	num := concrete.CalculateLastDayAdvClickCount(&newModel2, &oldModel)

	assert.Equal(t, []int64{exp}, []int64{num})
}

func TestDetermineAdvDay(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	//var exp int64 = 25
	concrete.DetermineAdvDay(&newModel, 6)

	assert.Equal(t, int64(0), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(0), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(0), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(1), newModel.SaturdayAdvClickCount)
}

func TestDetermineAdvHour(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	concrete.DetermineAdvHour(&newModel, 20)
	assert.Equal(t, int64(0), newModel.AdvClick0To5HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick6To11HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick12To17HourCount)
	assert.Equal(t, int64(1), newModel.AdvClick18To23HourCount)
}

func TestDetermineAdvAmPm(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	concrete.DetermineAdvAmPm(&newModel, 10)
	assert.Equal(t, int64(1), newModel.AmAdvClickCount)
	assert.Equal(t, int64(0), newModel.PmAdvClickCount)
}

func TestCalculateAdvLevelBasedAvgClickCount(t *testing.T) {
	concrete.CalculateAdvLevelBasedAvgClickCount(&newModel2)
	assert.Equal(t, float64(100), newModel2.LevelBasedAverageAdvCount)
}

// func Test_ConvertRawModelToResponseModelTrue(t *testing.T) {

// }
