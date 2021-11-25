package test

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"FilterWorkerService/test/Mock/service"
	"errors"
	"fmt"
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
	var manager = concrete.advEventManager{
		IAdvEventDal:  testAdv,
		IJsonParser:   &gojson.goJson{},
		ICacheService: testCache,
	}
	var advModel = model.AdvEventModel{
		ProjectId:   "Test",
		ClientId:    "Test",
		CustomerId:  "Test",
		LevelName:   "1",
		LevelIndex:  1,
		AdvType:     "test",
		InMinutes:   12,
		TrigerdTime: time.Now(),
	}

	var responseModel = model.AdvEventRespondModel{
		ProjectId:                            "Test",
		ClientId:                             "Test",
		CustomerId:                           "Test",
		LevelIndex:                           1,
		TotalAdvDay:                          1,
		TotalAdvCount:                        1,
		TotalAdvHour:                         0,
		TotalAdvMinute:                       0,
		LevelBasedAverageAdvCount:            1,
		AverageAdvDailyClickCount:            1,
		FirstAdvYearOfDay:                    int64(advModel.TrigerdTime.YearDay()),
		FirstAdvYear:                         int64(advModel.TrigerdTime.Year()),
		FirstWeekDay:                         int64(advModel.TrigerdTime.Weekday()),
		FirstAdvClickHour:                    int64(advModel.TrigerdTime.Hour()),
		FirstADvClickMinute:                  int64(advModel.TrigerdTime.Minute()),
		FirstAdvType:                         1,
		SecondAdvYearOfDay:                   0,
		SecondAdvHour:                        0,
		SecondAdvMinute:                      0,
		SecondAdvType:                        0,
		ThirdAdvYearOfDay:                    0,
		ThirdAdvHour:                         0,
		ThirdAdvMinute:                       0,
		ThirdAdvType:                         0,
		FourthAdvYearOfDay:                   0,
		FourthAdvHour:                        0,
		FourthAdvMinute:                      0,
		FourthAdvType:                        0,
		FifthAdvYearOfDay:                    0,
		FifthAdvHour:                         0,
		FifthAdvMinute:                       0,
		FifthAdvType:                         0,
		PenultimateAdvYearOfDay:              0,
		PenultimateAdvHour:                   0,
		PenultimateAdvMinute:                 0,
		PenultimateAdvType:                   0,
		LastAdvYearOfDay:                     0,
		LastAdvYear:                          0,
		LastAdvClickHour:                     0,
		LastAdvClickMinute:                   0,
		LastAdvType:                          0,
		FirstHalfHourAdvClickCount:           1,
		FirstHourAdvClickCount:               1,
		FirstTwoHourAdvClickCount:            1,
		FirstThreeHourAdvClickCount:          1,
		FirstSixHourAdvClickCount:            1,
		FirstTwelveHourAdvClickCount:         1,
		FirstDayAdvClickCount:                1,
		SecondDayAdvClickCount:               0,
		ThirdDayAdvClickCount:                0,
		FourthDayAdvClickCount:               0,
		FifthDayAdvClickCount:                0,
		SixthDayAdvClickCount:                0,
		SeventhDayAdvClickCount:              0,
		PenultimateDayAdvClickCount:          0,
		LastDayAdvClickCount:                 0,
		LastMinusFirstDayAdvClickCount:       -1,
		LastMinusPenultimateDayAdvClickCount: 0,
		LastDayAdvClickCountMinusAverageDailyAdvClickCount: -1,
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



func Test_UpdateAdvEvent(t *testing.T){
	var testADvDal = new(repository.MockAdvEventDal)
	var testCache = new(service.MockCacheService)
	var manager = concrete.advEventManager{
		IAdvEventDal:  testADvDal,
		IJsonParser:   &gojson.goJson{},
		ICacheService: testCache,
	}
	var advModel = model.AdvEventModel{
		ProjectId:   "Test",
		ClientId:    "Test",
		CustomerId:  "Test",
		LevelName:   "Level",
		LevelIndex:  5,
		AdvType:     "test",
		InMinutes:   12,
		TrigerdTime: time.Date(
        2021, 11, 8, 16, 11, 36, 651387237, time.UTC),
	}
	var responseModel = model.AdvEventRespondModel{
		ProjectId:                            "Test",
		ClientId:                             "Test",
		CustomerId:                           "Test",
		LevelIndex:                           int64(advModel.LevelIndex),
		TotalAdvDay:                          1,
		TotalAdvCount:                        1,
		TotalAdvHour:                         0,
		TotalAdvMinute:                       1,
		LevelBasedAverageAdvCount:            0.5,
		AverageAdvDailyClickCount:            1,
		FirstAdvYearOfDay:                    int64(advModel.TrigerdTime.YearDay()),
		FirstAdvYear:                         int64(advModel.TrigerdTime.Year()),
		FirstWeekDay:                         int64(advModel.TrigerdTime.Weekday()),
		FirstAdvClickHour:                    int64(advModel.TrigerdTime.Hour()),
		FirstADvClickMinute:                  int64(advModel.TrigerdTime.Minute()),
		FirstAdvType:                         12,
		SecondAdvYearOfDay:                   0,
		SecondAdvHour:                        0,
		SecondAdvMinute:                      0,
		SecondAdvType:                        0,
		ThirdAdvYearOfDay:                    0,
		ThirdAdvHour:                         0,
		ThirdAdvMinute:                       0,
		ThirdAdvType:                         0,
		FourthAdvYearOfDay:                   0,
		FourthAdvHour:                        0,
		FourthAdvMinute:                      0,
		FourthAdvType:                        0,
		FifthAdvYearOfDay:                    0,
		FifthAdvHour:                         0,
		FifthAdvMinute:                       0,
		FifthAdvType:                         0,
		PenultimateAdvYearOfDay:              0,
		PenultimateAdvHour:                   0,
		PenultimateAdvMinute:                 0,
		PenultimateAdvType:                   0,
		LastAdvYearOfDay:                     0,
		LastAdvYear:                          0,
		LastAdvClickHour:                     0,
		LastAdvClickMinute:                   0,
		LastAdvType:                          0,
		FirstHalfHourAdvClickCount:           1,
		FirstHourAdvClickCount:               1,
		FirstTwoHourAdvClickCount:            1,
		FirstThreeHourAdvClickCount:          1,
		FirstSixHourAdvClickCount:            1,
		FirstTwelveHourAdvClickCount:         1,
		FirstDayAdvClickCount:                1,
		SecondDayAdvClickCount:               0,
		ThirdDayAdvClickCount:                0,
		FourthDayAdvClickCount:               0,
		FifthDayAdvClickCount:                0,
		SixthDayAdvClickCount:                0,
		SeventhDayAdvClickCount:              0,
		PenultimateDayAdvClickCount:          0,
		LastDayAdvClickCount:                 0,
		LastMinusFirstDayAdvClickCount:       -1,
		LastMinusPenultimateDayAdvClickCount: 0,
		LastDayAdvClickCountMinusAverageDailyAdvClickCount: -1,
		SundayAdvClickCount:     0,
		MondayAdvClickCount:     1,
		TuesdayAdvClickCount:    0,
		WednesdayAdvClickCount:  0,
		ThursdayAdvClickCount:   0,
		FridayAdvClickCount:     0,
		SaturdayAdvClickCount:   0,
		AmAdvClickCount:         0,
		PmAdvClickCount:         1,
		AdvClick0To5HourCount:   0,
		AdvClick6To11HourCount:  0,
		AdvClick12To17HourCount: 1,
		AdvClick18To23HourCount: 0,
	}
	then := time.Date(
        2021, 11, 5, 20, 34, 58, 651387237, time.UTC)

	then2 := time.Date(
        2021, 11, 7, 11, 45, 07, 651387237, time.UTC)

	var oldModel = model.AdvEventRespondModel{
		ProjectId:                            "Test",
		ClientId:                             "Test",
		CustomerId:                           "Test",
		LevelIndex:                           0,
		TotalAdvDay:                          int64(then2.YearDay()) - int64(then.YearDay()),
		TotalAdvCount:                        2,
		TotalAdvHour:                         0,
		TotalAdvMinute:                       0,
		LevelBasedAverageAdvCount:            2,
		AverageAdvDailyClickCount:            1,
		FirstAdvYearOfDay:                    int64(then.YearDay()),
		FirstAdvYear:                         int64(then.Year()),
		FirstWeekDay:                         int64(then.Weekday()),
		FirstAdvClickHour:                    int64(then.Hour()),
		FirstADvClickMinute:                  int64(then.Minute()),
		FirstAdvType:                         8,
		SecondAdvYearOfDay:                   0,
		SecondAdvHour:                        0,
		SecondAdvMinute:                      0,
		SecondAdvType:                        0,
		ThirdAdvYearOfDay:                    0,
		ThirdAdvHour:                         0,
		ThirdAdvMinute:                       0,
		ThirdAdvType:                         0,
		FourthAdvYearOfDay:                   0,
		FourthAdvHour:                        0,
		FourthAdvMinute:                      0,
		FourthAdvType:                        0,
		FifthAdvYearOfDay:                    0,
		FifthAdvHour:                         0,
		FifthAdvMinute:                       0,
		FifthAdvType:                         0,
		PenultimateAdvYearOfDay:              0,
		PenultimateAdvHour:                   0,
		PenultimateAdvMinute:                 0,
		PenultimateAdvType:                   0,
		LastAdvYearOfDay:                     int64(then2.YearDay()),
		LastAdvYear:                          int64(then2.Year()),
		LastAdvClickHour:                     int64(then2.Hour()),
		LastAdvClickMinute:                   int64(then2.Minute()),
		LastAdvType:                          5,
		FirstHalfHourAdvClickCount:           0,
		FirstHourAdvClickCount:               0,
		FirstTwoHourAdvClickCount:            0,
		FirstThreeHourAdvClickCount:          0,
		FirstSixHourAdvClickCount:            0,
		FirstTwelveHourAdvClickCount:         0,
		FirstDayAdvClickCount:                3,
		SecondDayAdvClickCount:               0,
		ThirdDayAdvClickCount:                0,
		FourthDayAdvClickCount:               0,
		FifthDayAdvClickCount:                0,
		SixthDayAdvClickCount:                0,
		SeventhDayAdvClickCount:              0,
		PenultimateDayAdvClickCount:          9,
		LastDayAdvClickCount:                 4,
		LastMinusFirstDayAdvClickCount:       3,
		LastMinusPenultimateDayAdvClickCount: 4,
		LastDayAdvClickCountMinusAverageDailyAdvClickCount: 0,
		SundayAdvClickCount:     3,
		MondayAdvClickCount:     0,
		TuesdayAdvClickCount:    5,
		WednesdayAdvClickCount:  0,
		ThursdayAdvClickCount:   3,
		FridayAdvClickCount:     1,
		SaturdayAdvClickCount:   1,
		AmAdvClickCount:         2,
		PmAdvClickCount:         3,
		AdvClick0To5HourCount:   0,
		AdvClick6To11HourCount:  1,
		AdvClick12To17HourCount: 0,
		AdvClick18To23HourCount: 2,
	}

	var totalAdvCount int64 = oldModel.TotalAdvCount + responseModel.TotalAdvCount
	var totalAdvDay int64 = (responseModel.FirstAdvYearOfDay-oldModel.FirstAdvYearOfDay) + 365*(responseModel.FirstAdvYear-oldModel.FirstAdvYear)
	var avegareDailyVideoClickCount float64 =  float64(totalAdvCount) / float64(totalAdvDay)
	var totalAdvHour int64 = ((responseModel.FirstAdvYearOfDay+365*responseModel.FirstAdvYear)*24 + responseModel.FirstAdvClickHour) - ((oldModel.FirstAdvYearOfDay+365*oldModel.FirstAdvYear)*24 + oldModel.FirstAdvClickHour)
	var totalAdvMinute int64 = (((responseModel.FirstAdvYearOfDay+365*responseModel.FirstAdvYear)*24 + responseModel.FirstAdvClickHour)*60 + responseModel.FirstADvClickMinute) - (((oldModel.FirstAdvYearOfDay+365*oldModel.FirstAdvYear)*24 + oldModel.FirstAdvClickHour)*60 + oldModel.FirstADvClickMinute)

	var updatedModel = model.AdvEventRespondModel{
		ProjectId:                            "Test",
		ClientId:                             "Test",
		CustomerId:                           "Test",
		LevelIndex:                           responseModel.LevelIndex,
		TotalAdvDay:                          totalAdvDay,
		TotalAdvCount:                        oldModel.TotalAdvCount + responseModel.TotalAdvCount,
		TotalAdvHour:                         totalAdvHour,
		TotalAdvMinute:                       totalAdvMinute,
		LevelBasedAverageAdvCount:            float64(totalAdvCount) / float64(responseModel.LevelIndex),
		AverageAdvDailyClickCount:            avegareDailyVideoClickCount,
		FirstAdvYearOfDay:                    oldModel.FirstAdvYearOfDay,
		FirstAdvYear:                         oldModel.FirstAdvYear,
		FirstWeekDay:                         oldModel.FirstWeekDay,
		FirstAdvClickHour:                    oldModel.FirstAdvClickHour,
		FirstADvClickMinute:                  oldModel.FirstADvClickMinute,
		FirstAdvType:                         oldModel.FirstAdvType,
		SecondAdvYearOfDay:                   0,
		SecondAdvHour:                        0,
		SecondAdvMinute:                      0,
		SecondAdvType:                        0,
		ThirdAdvYearOfDay:                    responseModel.FirstAdvYearOfDay,
		ThirdAdvHour:                         responseModel.FirstAdvClickHour,
		ThirdAdvMinute:                       responseModel.FirstADvClickMinute,
		ThirdAdvType:                         responseModel.FirstAdvType,
		FourthAdvYearOfDay:                   0,
		FourthAdvHour:                        0,
		FourthAdvMinute:                      0,
		FourthAdvType:                        0,
		FifthAdvYearOfDay:                    0,
		FifthAdvHour:                         0,
		FifthAdvMinute:                       0,
		FifthAdvType:                         0,
		PenultimateAdvYearOfDay:              oldModel.LastAdvYearOfDay,
		PenultimateAdvHour:                   oldModel.LastAdvClickHour,
		PenultimateAdvMinute:                 oldModel.LastAdvClickMinute,
		PenultimateAdvType:                   0,
		LastAdvYearOfDay:                     responseModel.FirstAdvYearOfDay,
		LastAdvYear:                          responseModel.FirstAdvYear,
		LastAdvClickHour:                     responseModel.FirstAdvClickHour,
		LastAdvClickMinute:                   responseModel.FirstADvClickMinute,
		LastAdvType:                          responseModel.FirstAdvType,
		FirstHalfHourAdvClickCount:           0,
		FirstHourAdvClickCount:               0,
		FirstTwoHourAdvClickCount:            0,
		FirstThreeHourAdvClickCount:          0,
		FirstSixHourAdvClickCount:            0,
		FirstTwelveHourAdvClickCount:         0,
		FirstDayAdvClickCount:                oldModel.FirstDayAdvClickCount,
		SecondDayAdvClickCount:               0,
		ThirdDayAdvClickCount:                1,
		FourthDayAdvClickCount:               0,
		FifthDayAdvClickCount:                0,
		SixthDayAdvClickCount:                0,
		SeventhDayAdvClickCount:              0,
		PenultimateDayAdvClickCount:          oldModel.LastDayAdvClickCount,
		LastDayAdvClickCount:                 responseModel.FirstDayAdvClickCount,
		LastMinusFirstDayAdvClickCount:       responseModel.FirstDayAdvClickCount  - oldModel.FirstDayAdvClickCount,
		LastMinusPenultimateDayAdvClickCount: responseModel.FirstDayAdvClickCount - oldModel.LastDayAdvClickCount,
		LastDayAdvClickCountMinusAverageDailyAdvClickCount: float64(responseModel.FirstDayAdvClickCount) - avegareDailyVideoClickCount,
		SundayAdvClickCount:     responseModel.SundayAdvClickCount + oldModel.SundayAdvClickCount,
		MondayAdvClickCount:     responseModel.MondayAdvClickCount + oldModel.MondayAdvClickCount,
		TuesdayAdvClickCount:    responseModel.TuesdayAdvClickCount + oldModel.TuesdayAdvClickCount,
		WednesdayAdvClickCount:  responseModel.WednesdayAdvClickCount + oldModel.WednesdayAdvClickCount,
		ThursdayAdvClickCount:   responseModel.ThursdayAdvClickCount + oldModel.ThursdayAdvClickCount,
		FridayAdvClickCount:     responseModel.FridayAdvClickCount + oldModel.FridayAdvClickCount,
		SaturdayAdvClickCount:   responseModel.SaturdayAdvClickCount + oldModel.SaturdayAdvClickCount,
		AmAdvClickCount:         responseModel.AmAdvClickCount + oldModel.AmAdvClickCount,
		PmAdvClickCount:         responseModel.PmAdvClickCount + oldModel.PmAdvClickCount,
		AdvClick0To5HourCount:   responseModel.AdvClick0To5HourCount + oldModel.AdvClick0To5HourCount,
		AdvClick6To11HourCount:  responseModel.AdvClick6To11HourCount + oldModel.AdvClick6To11HourCount,
		AdvClick12To17HourCount: responseModel.AdvClick12To17HourCount + oldModel.AdvClick12To17HourCount,
		AdvClick18To23HourCount: responseModel.AdvClick18To23HourCount + oldModel.AdvClick18To23HourCount,
	}
	testADvDal.On("UpdateAdvEventById", advModel.ClientId, &updatedModel).Return(nil)

	var v, s, m = manager.UpdateAdvEvent(&responseModel, &oldModel)
	assert.Equal(t, true, s)
	assert.Equal(t, &updatedModel, v)
	assert.Equal(t, nil, m)
	fmt.Println(v)

}




func Test_CalculateSecondAdv_DateConversionSuccess(t *testing.T) {

	day, hour, minute, advType := concrete.CalculateSecondAdv(&newModel2, &oldModel)

	var Expday int64 = 9
	var Exphour int64 = 10
	var ExpMinute int64 = 11
	var ExpAddvType int64= 12

	assert.Equal(t, []int64{Expday, Exphour, ExpMinute, ExpAddvType}, []int64{day, hour, minute, advType})

}

func TestCalculateThirdAdv(t *testing.T) {

	var Expday int64 = 400
	var Exphour int64 = 600
	var ExpMinute int64 = 700
	var ExpAdvType int64 = 12
	day, hour, minute, advType:= concrete.CalculateThirdAdv(&newModel2, &oldModel)

	assert.Equal(t, []int64{Expday, Exphour, ExpMinute, ExpAdvType}, []int64{day, hour, minute, advType})

}

// func TestCalculateFirstDayAdvClickCount(t *testing.T) {

// 	var exp int64 = 23
// 	num := concrete.CalculateFirstDayAdvClickCount(&newModel2, &oldModel, totalAdvMinute)

// 	assert.Equal(t, []int64{exp}, []int64{num})
// }

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

