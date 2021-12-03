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
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var oldAdvModel = model.AdvEventRespondModel{}

var advModel = model.AdvEventModel{}

var responseAdvModel = model.AdvEventRespondModel{}



func Test_ConvertRawModelToResponse_AddedSuccess(t *testing.T) {
	//Arrance
	var testAdv = new(repository.MockAdvEventDal)
	var testCache = new(service.MockCacheService)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.AdvEventDal = testAdv
	IoC.Logger = testLog
	IoC.CacheService = testCache
	var advModel_test = advModel
	advModel_test.ProjectId = "Test"
	advModel_test.ClientId = "Test"
	advModel_test.CustomerId = "Test"
	advModel_test.LevelName = "1"
	advModel_test.LevelIndex = 1
	advModel_test.AdvType = "test"
	advModel_test.InMinutes = 12
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 5, 16, 11, 36, 651387237, time.UTC)
	var oldAdvModel_test = oldAdvModel
	var firstTime = time.Date(
		2021, 11, 5, 20, 34, 58, 651387237, time.UTC)
	var lastTime = time.Date(
		2021, 11, 7, 11, 45, 07, 651387237, time.UTC)
	oldAdvModel_test.ProjectId = "Test"
	oldAdvModel_test.ClientId = "Test"
	oldAdvModel_test.CustomerId = "Test"
	oldAdvModel_test.TotalAdvDay = int64(lastTime.YearDay()) - int64(firstTime.YearDay())
	oldAdvModel_test.FirstAdvYearOfDay = int64(firstTime.YearDay())
	oldAdvModel_test.FirstAdvYear = int64(firstTime.Year())
	oldAdvModel_test.FirstWeekDay = int64(firstTime.Weekday())
	oldAdvModel_test.FirstAdvClickHour = int64(firstTime.Hour())
	oldAdvModel_test.FirstADvClickMinute = int64(firstTime.Minute())
	oldAdvModel_test.FirstAdvType = 8
	oldAdvModel_test.FirstDayAdvClickCount = 3
	oldAdvModel_test.PenultimateDayAdvClickCount = 9
	oldAdvModel_test.LastDayAdvClickCount = 4
	oldAdvModel_test.LastMinusFirstDayAdvClickCount = 3
	oldAdvModel_test.LastMinusPenultimateDayAdvClickCount = 4
	oldAdvModel_test.TotalAdvCount = 2
	oldAdvModel_test.TotalAdvHour = 0
	oldAdvModel_test.TotalAdvMinute = 0
	oldAdvModel_test.LevelBasedAverageAdvCount = 2
	oldAdvModel_test.AverageAdvDailyClickCount = 1
	oldAdvModel_test.LastAdvYearOfDay = int64(lastTime.YearDay())
	oldAdvModel_test.LastAdvYear = int64(lastTime.Year())
	oldAdvModel_test.LastAdvClickHour = int64(lastTime.Hour())
	oldAdvModel_test.LastAdvClickMinute = int64(lastTime.Minute())
	oldAdvModel_test.LastAdvType = 5
	oldAdvModel_test.SundayAdvClickCount = 3
	oldAdvModel_test.MondayAdvClickCount = 0
	oldAdvModel_test.TuesdayAdvClickCount = 5
	oldAdvModel_test.WednesdayAdvClickCount = 0
	oldAdvModel_test.ThursdayAdvClickCount = 3
	oldAdvModel_test.FridayAdvClickCount = 1
	oldAdvModel_test.SaturdayAdvClickCount = 1
	oldAdvModel_test.AmAdvClickCount = 2
	oldAdvModel_test.PmAdvClickCount = 3
	oldAdvModel_test.AdvClick0To5HourCount = 0
	oldAdvModel_test.AdvClick6To11HourCount = 1
	oldAdvModel_test.AdvClick12To17HourCount = 0
	oldAdvModel_test.AdvClick18To23HourCount = 2
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.ProjectId = "Test"
	responseAdvModel_test.ClientId = "Test"
	responseAdvModel_test.CustomerId = "Test"
	responseAdvModel_test.LevelIndex = int64(advModel_test.LevelIndex)
	responseAdvModel_test.TotalAdvDay = 1
	responseAdvModel_test.TotalAdvCount = 1
	responseAdvModel_test.TotalAdvHour = 0
	responseAdvModel_test.TotalAdvMinute = 1
	responseAdvModel_test.LevelBasedAverageAdvCount = 1
	responseAdvModel_test.AverageAdvDailyClickCount = 1
	responseAdvModel_test.FirstAdvYearOfDay = int64(advModel.TrigerdTime.YearDay())
	responseAdvModel_test.FirstAdvYear = int64(advModel.TrigerdTime.Year())
	responseAdvModel_test.FirstWeekDay = int64(advModel.TrigerdTime.Weekday())
	responseAdvModel_test.FirstAdvClickHour = int64(advModel.TrigerdTime.Hour())
	responseAdvModel_test.FirstADvClickMinute = int64(advModel.TrigerdTime.Minute())
	responseAdvModel_test.SundayAdvClickCount = 1
	responseAdvModel_test.PmAdvClickCount = 1
	responseAdvModel_test.AdvClick12To17HourCount = 1
	responseAdvModel_test.FirstAdvType = 1
	responseAdvModel_test.FirstHalfHourAdvClickCount = 1
	responseAdvModel_test.FirstHourAdvClickCount = 1
	responseAdvModel_test.FirstTwoHourAdvClickCount = 1
	responseAdvModel_test.FirstThreeHourAdvClickCount = 1
	responseAdvModel_test.FirstSixHourAdvClickCount = 1
	responseAdvModel_test.FirstTwelveHourAdvClickCount = 1
	responseAdvModel_test.FirstDayAdvClickCount = 1
	responseAdvModel_test.LastMinusFirstDayAdvClickCount = -1
	responseAdvModel_test.LastDayAdvClickCountMinusAverageDailyAdvClickCount = -1
	responseAdvModel_test.FirstAdvType = 1

	var manager = concrete.AdvEventManagerConstructor()
	var message, _ = (*manager.IJsonParser).EncodeJson(&advModel_test)
	testCache.On("ManageCache", "AdvType", advModel_test.AdvType).Return(int64(1), true, "")
	testAdv.On("GetAdvEventById", advModel_test.ClientId).Return(&oldAdvModel_test,
		errors.New("null data error"))
	testAdv.On("Add", &responseAdvModel_test).Return(nil)
	//Act
	var v, s, m = manager.ConvertRawModelToResponseModel(message)
	var value, success = v.(model.AdvEventRespondModel)
	if success == true {
		assert.Equal(t, &responseAdvModel_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
}

func Test_UpdateAdvEvent(t *testing.T) {

	//Arrance
	var testAdv = new(repository.MockAdvEventDal)
	var testCache = new(service.MockCacheService)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.AdvEventDal = testAdv
	IoC.Logger = testLog
	IoC.CacheService = testCache
	var manager = concrete.AdvEventManagerConstructor()
	var advModel_test = advModel
	advModel_test.LevelIndex = 5
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 1, 16, 11, 36, 651387237, time.UTC)
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.LevelBasedAverageAdvCount = 0.2
	responseAdvModel_test.FirstAdvYearOfDay = int64(advModel_test.TrigerdTime.YearDay())
	responseAdvModel_test.FirstAdvYear = int64(advModel_test.TrigerdTime.Year())
	responseAdvModel_test.FirstWeekDay = int64(advModel_test.TrigerdTime.Weekday())
	responseAdvModel_test.FirstAdvClickHour = int64(advModel_test.TrigerdTime.Hour())
	responseAdvModel_test.FirstADvClickMinute = int64(advModel_test.TrigerdTime.Minute())
	responseAdvModel_test.FirstAdvType = 12
	responseAdvModel_test.PmAdvClickCount = 1
	responseAdvModel_test.WednesdayAdvClickCount = 1
	var oldAdvModel_test = oldAdvModel
	var firstTime = time.Date(
		2021, 11, 5, 20, 34, 58, 651387237, time.UTC)
	var lastTime = time.Date(
		2021, 11, 7, 11, 45, 07, 651387237, time.UTC)

	oldAdvModel_test.ProjectId = "Test"
	oldAdvModel_test.ClientId = "Test"
	oldAdvModel_test.CustomerId = "Test"
	oldAdvModel_test.TotalAdvDay = int64(lastTime.YearDay()) - int64(firstTime.YearDay())
	oldAdvModel_test.FirstAdvYearOfDay = int64(firstTime.YearDay())
	oldAdvModel_test.FirstAdvYear = int64(firstTime.Year())
	oldAdvModel_test.FirstWeekDay = int64(firstTime.Weekday())
	oldAdvModel_test.FirstAdvClickHour = int64(firstTime.Hour())
	oldAdvModel_test.FirstADvClickMinute = int64(firstTime.Minute())
	oldAdvModel_test.FirstAdvType = 8
	oldAdvModel_test.FirstDayAdvClickCount = 3
	oldAdvModel_test.PenultimateDayAdvClickCount = 9
	oldAdvModel_test.LastDayAdvClickCount = 4
	oldAdvModel_test.LastMinusFirstDayAdvClickCount = 3
	oldAdvModel_test.LastMinusPenultimateDayAdvClickCount = 4
	oldAdvModel_test.TotalAdvCount = 2
	oldAdvModel_test.TotalAdvHour = 0
	oldAdvModel_test.TotalAdvMinute = 0
	oldAdvModel_test.LevelBasedAverageAdvCount = 2
	oldAdvModel_test.AverageAdvDailyClickCount = 1
	oldAdvModel_test.LastAdvYearOfDay = int64(lastTime.YearDay())
	oldAdvModel_test.LastAdvYear = int64(lastTime.Year())
	oldAdvModel_test.LastAdvClickHour = int64(lastTime.Hour())
	oldAdvModel_test.LastAdvClickMinute = int64(lastTime.Minute())
	oldAdvModel_test.LastAdvType = 5
	oldAdvModel_test.SundayAdvClickCount = 3
	oldAdvModel_test.MondayAdvClickCount = 0
	oldAdvModel_test.TuesdayAdvClickCount = 5
	oldAdvModel_test.WednesdayAdvClickCount = 0
	oldAdvModel_test.ThursdayAdvClickCount = 3
	oldAdvModel_test.FridayAdvClickCount = 1
	oldAdvModel_test.SaturdayAdvClickCount = 1
	oldAdvModel_test.AmAdvClickCount = 2
	oldAdvModel_test.PmAdvClickCount = 3
	oldAdvModel_test.AdvClick0To5HourCount = 0
	oldAdvModel_test.AdvClick6To11HourCount = 1
	oldAdvModel_test.AdvClick12To17HourCount = 0
	oldAdvModel_test.AdvClick18To23HourCount = 2

	var totalAdvCount int64 = oldAdvModel_test.TotalAdvCount + responseAdvModel_test.TotalAdvCount
	var totalAdvDay int64 = (responseAdvModel_test.FirstAdvYearOfDay - oldAdvModel_test.FirstAdvYearOfDay) + 365*(responseAdvModel_test.FirstAdvYear-oldAdvModel_test.FirstAdvYear)
	var avegareDailyVideoClickCount float64 = float64(totalAdvCount) / float64(totalAdvDay)
	var totalAdvHour int64 = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var totalAdvMinute int64 = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)

	var updatedModel = model.AdvEventRespondModel{
		ProjectId:                            "Test",
		ClientId:                             "Test",
		CustomerId:                           "Test",
		LevelIndex:                           responseAdvModel_test.LevelIndex,
		TotalAdvDay:                          totalAdvDay,
		TotalAdvCount:                        oldAdvModel_test.TotalAdvCount + responseAdvModel_test.TotalAdvCount,
		TotalAdvHour:                         totalAdvHour,
		TotalAdvMinute:                       totalAdvMinute,
		LevelBasedAverageAdvCount:            float64(totalAdvCount) / float64(responseAdvModel_test.LevelIndex),
		AverageAdvDailyClickCount:            avegareDailyVideoClickCount,
		FirstAdvYearOfDay:                    oldAdvModel_test.FirstAdvYearOfDay,
		FirstAdvYear:                         oldAdvModel_test.FirstAdvYear,
		FirstWeekDay:                         oldAdvModel_test.FirstWeekDay,
		FirstAdvClickHour:                    oldAdvModel_test.FirstAdvClickHour,
		FirstADvClickMinute:                  oldAdvModel_test.FirstADvClickMinute,
		FirstAdvType:                         oldAdvModel_test.FirstAdvType,
		SecondAdvYearOfDay:                   0,
		SecondAdvHour:                        0,
		SecondAdvMinute:                      0,
		SecondAdvType:                        0,
		ThirdAdvYearOfDay:                    responseAdvModel_test.FirstAdvYearOfDay,
		ThirdAdvHour:                         responseAdvModel_test.FirstAdvClickHour,
		ThirdAdvMinute:                       responseAdvModel_test.FirstADvClickMinute,
		ThirdAdvType:                         responseAdvModel_test.FirstAdvType,
		FourthAdvYearOfDay:                   0,
		FourthAdvHour:                        0,
		FourthAdvMinute:                      0,
		FourthAdvType:                        0,
		FifthAdvYearOfDay:                    0,
		FifthAdvHour:                         0,
		FifthAdvMinute:                       0,
		FifthAdvType:                         0,
		PenultimateAdvYearOfDay:              oldAdvModel_test.LastAdvYearOfDay,
		PenultimateAdvHour:                   oldAdvModel_test.LastAdvClickHour,
		PenultimateAdvMinute:                 oldAdvModel_test.LastAdvClickMinute,
		PenultimateAdvType:                   0,
		LastAdvYearOfDay:                     responseAdvModel_test.FirstAdvYearOfDay,
		LastAdvYear:                          responseAdvModel_test.FirstAdvYear,
		LastAdvClickHour:                     responseAdvModel_test.FirstAdvClickHour,
		LastAdvClickMinute:                   responseAdvModel_test.FirstADvClickMinute,
		LastAdvType:                          responseAdvModel_test.FirstAdvType,
		FirstHalfHourAdvClickCount:           0,
		FirstHourAdvClickCount:               0,
		FirstTwoHourAdvClickCount:            0,
		FirstThreeHourAdvClickCount:          0,
		FirstSixHourAdvClickCount:            0,
		FirstTwelveHourAdvClickCount:         0,
		FirstDayAdvClickCount:                oldAdvModel_test.FirstDayAdvClickCount,
		SecondDayAdvClickCount:               0,
		ThirdDayAdvClickCount:                1,
		FourthDayAdvClickCount:               0,
		FifthDayAdvClickCount:                0,
		SixthDayAdvClickCount:                0,
		SeventhDayAdvClickCount:              0,
		PenultimateDayAdvClickCount:          oldAdvModel_test.LastDayAdvClickCount,
		LastDayAdvClickCount:                 responseAdvModel_test.FirstDayAdvClickCount,
		LastMinusFirstDayAdvClickCount:       responseAdvModel_test.FirstDayAdvClickCount - oldAdvModel_test.FirstDayAdvClickCount,
		LastMinusPenultimateDayAdvClickCount: responseAdvModel_test.FirstDayAdvClickCount - oldAdvModel_test.LastDayAdvClickCount,
		LastDayAdvClickCountMinusAverageDailyAdvClickCount: float64(responseAdvModel_test.FirstDayAdvClickCount) - avegareDailyVideoClickCount,
		SundayAdvClickCount:     responseAdvModel_test.SundayAdvClickCount + oldAdvModel_test.SundayAdvClickCount,
		MondayAdvClickCount:     responseAdvModel_test.MondayAdvClickCount + oldAdvModel_test.MondayAdvClickCount,
		TuesdayAdvClickCount:    responseAdvModel_test.TuesdayAdvClickCount + oldAdvModel_test.TuesdayAdvClickCount,
		WednesdayAdvClickCount:  responseAdvModel_test.WednesdayAdvClickCount + oldAdvModel_test.WednesdayAdvClickCount,
		ThursdayAdvClickCount:   responseAdvModel_test.ThursdayAdvClickCount + oldAdvModel_test.ThursdayAdvClickCount,
		FridayAdvClickCount:     responseAdvModel_test.FridayAdvClickCount + oldAdvModel_test.FridayAdvClickCount,
		SaturdayAdvClickCount:   responseAdvModel_test.SaturdayAdvClickCount + oldAdvModel_test.SaturdayAdvClickCount,
		AmAdvClickCount:         responseAdvModel_test.AmAdvClickCount + oldAdvModel_test.AmAdvClickCount,
		PmAdvClickCount:         responseAdvModel_test.PmAdvClickCount + oldAdvModel_test.PmAdvClickCount,
		AdvClick0To5HourCount:   responseAdvModel_test.AdvClick0To5HourCount + oldAdvModel_test.AdvClick0To5HourCount,
		AdvClick6To11HourCount:  responseAdvModel_test.AdvClick6To11HourCount + oldAdvModel_test.AdvClick6To11HourCount,
		AdvClick12To17HourCount: responseAdvModel_test.AdvClick12To17HourCount + oldAdvModel_test.AdvClick12To17HourCount,
		AdvClick18To23HourCount: responseAdvModel_test.AdvClick18To23HourCount + oldAdvModel_test.AdvClick18To23HourCount,
	}
	testAdv.On("UpdateAdvEventById", advModel.ClientId, &updatedModel).Return(nil)

	var v, s, m = manager.UpdateAdvEvent(&responseAdvModel_test, &oldAdvModel_test)
	assert.Equal(t, true, s)
	assert.Equal(t, &updatedModel, v)
	assert.Equal(t, nil, m)
	fmt.Println(v)

}

func Test_CalculateSecondAdv_DateConversion_EqualTotalAdvCount2(t *testing.T) {
	var oldAdvModel_test = oldAdvModel

	oldAdvModel_test.TotalAdvCount = 2
	oldAdvModel_test.SecondAdvYearOfDay = 0
	oldAdvModel_test.SecondAdvHour = 0
	oldAdvModel_test.SecondAdvMinute = 0
	oldAdvModel_test.SecondAdvType = 0
	var responseAdvModel_test = responseAdvModel

	responseAdvModel_test.FirstAdvYearOfDay = 104
	responseAdvModel_test.FirstAdvClickHour = 15
	responseAdvModel_test.FirstADvClickMinute = 34
	responseAdvModel_test.FirstAdvType = 8
	var Expday int64 = 104
	var Exphour int64 = 15
	var ExpMinute int64 = 34
	var ExpAddvType int64 = 8
	concrete.CalculateSecondAdv(&responseAdvModel_test, &oldAdvModel_test)
	assert.Equal(t, Expday, oldAdvModel_test.SecondAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.SecondAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.SecondAdvMinute)
	assert.Equal(t, ExpAddvType, oldAdvModel_test.SecondAdvType)

}

func Test_CalculateSecondAdv_DateConversion_NotEqualTotalAdvCount2(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	var responseAdvModel_test = responseAdvModel

	oldAdvModel_test.TotalAdvCount = 1
	oldAdvModel_test.SecondAdvYearOfDay = 0
	oldAdvModel_test.SecondAdvHour = 0
	oldAdvModel_test.SecondAdvMinute = 0
	oldAdvModel_test.SecondAdvType = 0

	responseAdvModel_test.FirstAdvYearOfDay = 104
	responseAdvModel_test.FirstAdvClickHour = 15
	responseAdvModel_test.FirstADvClickMinute = 34
	responseAdvModel_test.FirstAdvType = 8

	var Expday int64 = 0
	var Exphour int64 = 0
	var ExpMinute int64 = 0
	var ExpAddvType int64 = 0
	concrete.CalculateSecondAdv(&responseAdvModel_test, &oldAdvModel_test)
	assert.Equal(t, Expday, oldAdvModel_test.SecondAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.SecondAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.SecondAdvMinute)
	assert.Equal(t, ExpAddvType, oldAdvModel_test.SecondAdvType)

}

func Test_CalculateThirdAdv_EqualTotalAdvCount3(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	var responseAdvModel_test = responseAdvModel

	oldAdvModel_test.TotalAdvCount = 3
	oldAdvModel_test.ThirdAdvYearOfDay = 0
	oldAdvModel_test.ThirdAdvHour = 0
	oldAdvModel_test.ThirdAdvMinute = 0
	oldAdvModel_test.ThirdAdvType = 0

	responseAdvModel_test.FirstAdvYearOfDay = 165
	responseAdvModel_test.FirstAdvClickHour = 13
	responseAdvModel_test.FirstADvClickMinute = 48
	responseAdvModel_test.FirstAdvType = 7
	var Expday int64 = 165
	var Exphour int64 = 13
	var ExpMinute int64 = 48
	var ExpAdvType int64 = 7

	concrete.CalculateThirdAdv(&responseAdvModel_test, &oldAdvModel_test)

	assert.Equal(t, Expday, oldAdvModel_test.ThirdAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.ThirdAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.ThirdAdvMinute)
	assert.Equal(t, ExpAdvType, oldAdvModel_test.ThirdAdvType)
}

func Test_CalculateThirdAdv_NotEqualTotalAdvCount3(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	var responseAdvModel_test = responseAdvModel

	oldAdvModel_test.TotalAdvCount = 5
	oldAdvModel_test.ThirdAdvYearOfDay = 0
	oldAdvModel_test.ThirdAdvHour = 0
	oldAdvModel_test.ThirdAdvMinute = 0
	oldAdvModel_test.ThirdAdvType = 0

	responseAdvModel_test.FirstAdvYearOfDay = 165
	responseAdvModel_test.FirstAdvClickHour = 13
	responseAdvModel_test.FirstADvClickMinute = 48
	responseAdvModel_test.FirstAdvType = 7
	var Expday int64 = 0
	var Exphour int64 = 0
	var ExpMinute int64 = 0
	var ExpAdvType int64 = 0
	concrete.CalculateThirdAdv(&responseAdvModel_test, &oldAdvModel_test)

	assert.Equal(t, Expday, oldAdvModel_test.ThirdAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.ThirdAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.ThirdAdvMinute)
	assert.Equal(t, ExpAdvType, oldAdvModel_test.ThirdAdvType)
}

func Test_CalculateFourthAdv_EqualTotalAdvCount4(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	var responseAdvModel_test = responseAdvModel

	oldAdvModel_test.TotalAdvCount = 4
	oldAdvModel_test.FourthAdvYearOfDay = 0
	oldAdvModel_test.FourthAdvHour = 0
	oldAdvModel_test.FourthAdvMinute = 0
	oldAdvModel_test.FourthAdvType = 0

	responseAdvModel_test.FirstAdvYearOfDay = 239
	responseAdvModel_test.FirstAdvClickHour = 18
	responseAdvModel_test.FirstADvClickMinute = 30
	responseAdvModel_test.FirstAdvType = 7
	var Expday int64 = 239
	var Exphour int64 = 18
	var ExpMinute int64 = 30
	var ExpAdvType int64 = 7
	concrete.CalculateFourthAdv(&responseAdvModel_test, &oldAdvModel_test)

	assert.Equal(t, Expday, oldAdvModel_test.FourthAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.FourthAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.FourthAdvMinute)
	assert.Equal(t, ExpAdvType, oldAdvModel_test.FourthAdvType)
}

func Test_CalculateFourthAdv_NotEqualTotalAdvCount4(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	var responseAdvModel_test = responseAdvModel

	oldAdvModel_test.TotalAdvCount = 5
	oldAdvModel_test.FourthAdvYearOfDay = 0
	oldAdvModel_test.FourthAdvHour = 0
	oldAdvModel_test.FourthAdvMinute = 0
	oldAdvModel_test.FourthAdvType = 0

	responseAdvModel_test.FirstAdvYearOfDay = 165
	responseAdvModel_test.FirstAdvClickHour = 13
	responseAdvModel_test.FirstADvClickMinute = 48
	responseAdvModel_test.FirstAdvType = 7
	var Expday int64 = 0
	var Exphour int64 = 0
	var ExpMinute int64 = 0
	var ExpAdvType int64 = 0
	concrete.CalculateFourthAdv(&responseAdvModel_test, &oldAdvModel_test)

	assert.Equal(t, Expday, oldAdvModel_test.FourthAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.FourthAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.FourthAdvMinute)
	assert.Equal(t, ExpAdvType, oldAdvModel_test.FourthAdvType)
}

func Test_CalculateFifthAdv_EqualTotalAdvCount5(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	var responseAdvModel_test = responseAdvModel

	oldAdvModel_test.TotalAdvCount = 5
	oldAdvModel_test.FifthAdvYearOfDay = 0
	oldAdvModel_test.FifthAdvHour = 0
	oldAdvModel_test.FifthAdvMinute = 0
	oldAdvModel_test.FifthAdvType = 0

	responseAdvModel_test.FirstAdvYearOfDay = 350
	responseAdvModel_test.FirstAdvClickHour = 23
	responseAdvModel_test.FirstADvClickMinute = 43
	responseAdvModel_test.FirstAdvType = 7
	var Expday int64 = 350
	var Exphour int64 = 23
	var ExpMinute int64 = 43
	var ExpAdvType int64 = 7
	concrete.CalculateFifthAdv(&responseAdvModel_test, &oldAdvModel_test)

	assert.Equal(t, Expday, oldAdvModel_test.FifthAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.FifthAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.FifthAdvMinute)
	assert.Equal(t, ExpAdvType, oldAdvModel_test.FifthAdvType)
}

func Test_CalculateFifthAdv_NotEqualTotalAdvCount5(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	var responseAdvModel_test = responseAdvModel

	oldAdvModel_test.TotalAdvCount = 9
	oldAdvModel_test.FifthAdvYearOfDay = 0
	oldAdvModel_test.FifthAdvHour = 0
	oldAdvModel_test.FifthAdvMinute = 0
	oldAdvModel_test.FifthAdvType = 0

	responseAdvModel_test.FirstAdvYearOfDay = 165
	responseAdvModel_test.FirstAdvClickHour = 13
	responseAdvModel_test.FirstADvClickMinute = 48
	responseAdvModel_test.FirstAdvType = 7
	var Expday int64 = 0
	var Exphour int64 = 0
	var ExpMinute int64 = 0
	var ExpAdvType int64 = 0
	concrete.CalculateFourthAdv(&responseAdvModel_test, &oldAdvModel_test)

	assert.Equal(t, Expday, oldAdvModel_test.FifthAdvYearOfDay)
	assert.Equal(t, Exphour, oldAdvModel_test.FifthAdvHour)
	assert.Equal(t, ExpMinute, oldAdvModel_test.FifthAdvMinute)
	assert.Equal(t, ExpAdvType, oldAdvModel_test.FifthAdvType)
}

func Test_CalculateAverageAdvDailyClickCount_TotalAdvDayZero(t *testing.T) {

	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.TotalAdvDay = 0
	oldAdvModel_test.TotalAdvCount = 23
	var ExpCount float64 = 23
	var count float64 = concrete.CalculateAverageAdvDailyClickCount(&oldAdvModel_test)
	assert.Equal(t, ExpCount, count)
}

func Test_CalculateAverageAdvDailyClickCount_NotTotalAdvDayZero(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.TotalAdvDay = 10
	oldAdvModel_test.TotalAdvCount = 23
	var ExpCount float64 = float64(23) / float64(10)
	var count float64 = concrete.CalculateAverageAdvDailyClickCount(&oldAdvModel_test)
	assert.Equal(t, ExpCount, count)
}

func Test_CalculateFirstHalfHourTotalAdvCount_In30Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 300
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstHalfHourAdvClickCount = 5
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 300
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 20
	responseAdvModel_test.FirstADvClickMinute = 30
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstHalfHourTotalAdvCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 5 + 1
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstHalfHourAdvClickCount)
}

func Test_CalculateFirstHalfHourTotalAdvCount_Out30Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 300
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstHalfHourAdvClickCount = 5
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 300
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 20
	responseAdvModel_test.FirstADvClickMinute = 31
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstHalfHourTotalAdvCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 5
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstHalfHourAdvClickCount)
}

func Test_CalculateFirstHourTotalAdvCount_In60Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 300
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstHourAdvClickCount = 9
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 300
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 21
	responseAdvModel_test.FirstADvClickMinute = 00
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstHourTotalAdvCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 9 + 1
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstHourAdvClickCount)
}

func Test_CalculateFirstHourTotalAdvCount_Out60Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 300
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstHourAdvClickCount = 7
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 300
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 21
	responseAdvModel_test.FirstADvClickMinute = 01
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstHourTotalAdvCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 7
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstHourAdvClickCount)
}

func Test_CalculateFirstTwoHourTotalAdvCount_In120Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 307
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstTwoHourAdvClickCount = 9
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 307
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 22
	responseAdvModel_test.FirstADvClickMinute = 00
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstTwoHourTotalAdvCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 9 + 1
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstTwoHourAdvClickCount)
}

func Test_CalculateFirstTwoHourTotalAdvCount_Out120Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 300
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstTwoHourAdvClickCount = 16
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 300
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 22
	responseAdvModel_test.FirstADvClickMinute = 01
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstTwoHourTotalAdvCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 16
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstTwoHourAdvClickCount)
}

func Test_CalculateFirstThreeHourAdvClickCount_In180Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 307
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstThreeHourAdvClickCount = 12
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 307
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 23
	responseAdvModel_test.FirstADvClickMinute = 00
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstThreeHourAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 12 + 1
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstThreeHourAdvClickCount)
}

func Test_CalculateFirstThreeHourAdvClickCount_Out180Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 300
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 20
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstThreeHourAdvClickCount = 16
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 300
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 23
	responseAdvModel_test.FirstADvClickMinute = 01
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstThreeHourAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 16
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstThreeHourAdvClickCount)
}

func Test_CalculateFirstSixHourAdvClickCount_In360Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 353
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 14
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstSixHourAdvClickCount = 42
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 353
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 20
	responseAdvModel_test.FirstADvClickMinute = 00
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstSixHourAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 42 + 1
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstSixHourAdvClickCount)
}

func Test_CalculateFirstSixHourAdvClickCount_Out360Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 214
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 12
	oldAdvModel_test.FirstADvClickMinute = 30
	oldAdvModel_test.FirstSixHourAdvClickCount = 23
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 214
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 18
	responseAdvModel_test.FirstADvClickMinute = 31
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstSixHourAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 23
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstSixHourAdvClickCount)
}

func Test_CalculateFirstTwelveHourAdvClickCount_In720Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 111
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 14
	oldAdvModel_test.FirstADvClickMinute = 00
	oldAdvModel_test.FirstTwelveHourAdvClickCount = 42
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 112
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 2
	responseAdvModel_test.FirstADvClickMinute = 00
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstSixHourAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 42 + 1
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstTwelveHourAdvClickCount)
}

func Test_CalculateFirstTwelveHourAdvClickCount_Out720Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 111
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 13
	oldAdvModel_test.FirstADvClickMinute = 30
	oldAdvModel_test.FirstTwelveHourAdvClickCount = 49
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 112
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 1
	responseAdvModel_test.FirstADvClickMinute = 31
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	concrete.CalculateFirstSixHourAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)
	var ExpCount int64 = 49
	assert.Equal(t, ExpCount, oldAdvModel_test.FirstTwelveHourAdvClickCount)
}

func Test_CalculateFirstDayAdvClickCount_In1440Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 70
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 13
	oldAdvModel_test.FirstADvClickMinute = 30
	oldAdvModel_test.FirstDayAdvClickCount = 120
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 71
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 13
	responseAdvModel_test.FirstADvClickMinute = 30
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	var expCount int64 = 120 + 1
	concrete.CalculateFirstDayAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)

	assert.Equal(t, expCount, oldAdvModel_test.FirstDayAdvClickCount)
}

func Test_CalculateFirstDayAdvClickCount_Out1440Minutes(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 70
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 13
	oldAdvModel_test.FirstADvClickMinute = 30
	oldAdvModel_test.FirstDayAdvClickCount = 120
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 71
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 13
	responseAdvModel_test.FirstADvClickMinute = 31
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvMinute_test = (((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24+responseAdvModel_test.FirstAdvClickHour)*60 + responseAdvModel_test.FirstADvClickMinute) - (((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24+oldAdvModel_test.FirstAdvClickHour)*60 + oldAdvModel_test.FirstADvClickMinute)
	var expCount int64 = 120
	concrete.CalculateFirstDayAdvClickCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvMinute_test)

	assert.Equal(t, expCount, oldAdvModel_test.FirstDayAdvClickCount)
}

func Test_CalculateSecondDayTotalSessionCount_In24To48Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 70
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 13
	oldAdvModel_test.FirstADvClickMinute = 30
	oldAdvModel_test.SecondDayAdvClickCount = 120
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 72
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 13
	responseAdvModel_test.FirstADvClickMinute = 30
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 120 + 1
	concrete.CalculateSecondDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.SecondDayAdvClickCount)
}

func Test_CalculateSecondDayTotalSessionCount_Out24To48Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 70
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 13
	oldAdvModel_test.FirstADvClickMinute = 30
	oldAdvModel_test.SecondDayAdvClickCount = 120
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 72
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 16
	responseAdvModel_test.FirstADvClickMinute = 31
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 120
	concrete.CalculateSecondDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.SecondDayAdvClickCount)
}

func Test_CalculateThirdDayTotalSessionCount_In48To72Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 270
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 13
	oldAdvModel_test.ThirdDayAdvClickCount = 246
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 273
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 12
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 246 + 1
	concrete.CalculateThirdDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.ThirdDayAdvClickCount)
}

func Test_CalculateThirdDayTotalSessionCount_Out48To72Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 270
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 13
	oldAdvModel_test.ThirdDayAdvClickCount = 230
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 273
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 14
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 230
	concrete.CalculateThirdDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.ThirdDayAdvClickCount)
}

func Test_CalculateFourthDayTotalSessionCount_In72To96Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 333
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 22
	oldAdvModel_test.FourthDayAdvClickCount = 246
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 337
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 22
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 246 + 1
	concrete.CalculateFourthDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.FourthDayAdvClickCount)
}

func Test_CalculateFourthDayTotalSessionCount_Out72To96Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 333
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 22
	oldAdvModel_test.FourthDayAdvClickCount = 230
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 337
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 23
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 230
	concrete.CalculateFourthDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.FourthDayAdvClickCount)
}

func Test_CalculateFifthDayTotalSessionCount_In96To120Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 333
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 21
	oldAdvModel_test.FifthDayAdvClickCount = 246
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 338
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 21
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 246 + 1
	concrete.CalculateFifthDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.FifthDayAdvClickCount)
}

func Test_CalculateFifthDayTotalSessionCount_Out96To120Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 333
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 22
	oldAdvModel_test.FifthDayAdvClickCount = 230
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 338
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 23
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 230
	concrete.CalculateFifthDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.FifthDayAdvClickCount)
}

func Test_CalculateSixthDayTotalSessionCount_In120To144Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 330
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 21
	oldAdvModel_test.SixthDayAdvClickCount = 246
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 336
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 21
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 246 + 1
	concrete.CalculateSixthDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.SixthDayAdvClickCount)
}

func Test_CalculateSixthDayTotalSessionCount_Out120To144Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 330
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 22
	oldAdvModel_test.SixthDayAdvClickCount = 230
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 336
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 23
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 230
	concrete.CalculateSeventhDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.SixthDayAdvClickCount)
}

func Test_CalculateSeventhDayTotalSessionCount_In144To168Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 330
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 21
	oldAdvModel_test.SixthDayAdvClickCount = 246
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 337
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 21
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 246 + 1
	concrete.CalculateSeventhDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.SixthDayAdvClickCount)
}

func Test_CalculateSeventhDayTotalSessionCount_Out144To168Hours(t *testing.T) {
	var oldAdvModel_test = oldAdvModel
	oldAdvModel_test.FirstAdvYearOfDay = 330
	oldAdvModel_test.FirstAdvYear = 2021
	oldAdvModel_test.FirstAdvClickHour = 22
	oldAdvModel_test.SixthDayAdvClickCount = 230
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.FirstAdvYearOfDay = 337
	responseAdvModel_test.FirstAdvYear = 2021
	responseAdvModel_test.FirstAdvClickHour = 23
	responseAdvModel_test.FirstDayAdvClickCount = 1
	var TotalAdvHours_test = ((responseAdvModel_test.FirstAdvYearOfDay+365*responseAdvModel_test.FirstAdvYear)*24 + responseAdvModel_test.FirstAdvClickHour) - ((oldAdvModel_test.FirstAdvYearOfDay+365*oldAdvModel_test.FirstAdvYear)*24 + oldAdvModel_test.FirstAdvClickHour)
	var expCount int64 = 230
	concrete.CalculateSixthDayTotalSessionCount(&responseAdvModel_test, &oldAdvModel_test, TotalAdvHours_test)

	assert.Equal(t, expCount, oldAdvModel_test.SixthDayAdvClickCount)
}

// func TestCalculatePenultimateDayAdvDay(t *testing.T) {

// 	var exp int64 = 24
// 	num := concrete.CalculatePenultimateDayAdvDay(&newModel2, &oldModel)

// 	assert.Equal(t, []int64{exp}, []int64{num})
// }

// func TestCalculateLastDayAdvClickCount(t *testing.T) {
// 	var exp int64 = 25
// 	num := concrete.CalculateLastDayAdvClickCount(&newModel2, &oldModel)

// 	assert.Equal(t, []int64{exp}, []int64{num})
// }

func Test_DetermineAdvDay_Sunday(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 5, 17, 11, 36, 651387237, time.UTC)
	var day = int64(advModel_test.TrigerdTime.Weekday())
	concrete.DetermineAdvDay(&newModel, day)

	assert.Equal(t, int64(1), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(0), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(0), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(0), newModel.SaturdayAdvClickCount)
}

func Test_DetermineAdvDay_Monday(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 6, 17, 11, 36, 651387237, time.UTC)
	var day = int64(advModel_test.TrigerdTime.Weekday())
	concrete.DetermineAdvDay(&newModel, day)

	assert.Equal(t, int64(0), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(1), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(0), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(0), newModel.SaturdayAdvClickCount)
}

func Test_DetermineAdvDay_Tuesday(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 7, 17, 11, 36, 651387237, time.UTC)
	var day = int64(advModel_test.TrigerdTime.Weekday())
	concrete.DetermineAdvDay(&newModel, day)

	assert.Equal(t, int64(0), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(0), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(1), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(0), newModel.SaturdayAdvClickCount)
}

func Test_DetermineAdvDay_Wednesday(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 8, 17, 11, 36, 651387237, time.UTC)
	var day = int64(advModel_test.TrigerdTime.Weekday())
	concrete.DetermineAdvDay(&newModel, day)

	assert.Equal(t, int64(0), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(0), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(0), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(1), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(0), newModel.SaturdayAdvClickCount)
}

func Test_DetermineAdvDay_Thursday(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 9, 17, 11, 36, 651387237, time.UTC)
	var day = int64(advModel_test.TrigerdTime.Weekday())
	concrete.DetermineAdvDay(&newModel, day)

	assert.Equal(t, int64(0), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(0), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(0), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(1), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(0), newModel.SaturdayAdvClickCount)
}

func Test_DetermineAdvDay_Friday(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 10, 17, 11, 36, 651387237, time.UTC)
	var day = int64(advModel_test.TrigerdTime.Weekday())
	concrete.DetermineAdvDay(&newModel, day)

	assert.Equal(t, int64(0), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(0), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(0), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(1), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(0), newModel.SaturdayAdvClickCount)
}

func Test_DetermineAdvDay_Saturday(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 11, 17, 11, 36, 651387237, time.UTC)
	var day = int64(advModel_test.TrigerdTime.Weekday())
	concrete.DetermineAdvDay(&newModel, day)

	assert.Equal(t, int64(0), newModel.SundayAdvClickCount)
	assert.Equal(t, int64(0), newModel.MondayAdvClickCount)
	assert.Equal(t, int64(0), newModel.TuesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.WednesdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.ThursdayAdvClickCount)
	assert.Equal(t, int64(0), newModel.FridayAdvClickCount)
	assert.Equal(t, int64(1), newModel.SaturdayAdvClickCount)
}

func Test_DetermineAdvHour_00To05Hours(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 11, 05, 11, 36, 651387237, time.UTC)
	var hour = int64(advModel_test.TrigerdTime.Hour())
	concrete.DetermineAdvHour(&newModel, hour)
	assert.Equal(t, int64(1), newModel.AdvClick0To5HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick6To11HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick12To17HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick18To23HourCount)
}

func Test_DetermineAdvHour_06To11Hours(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 11, 10, 11, 36, 651387237, time.UTC)
	var hour = int64(advModel_test.TrigerdTime.Hour())
	concrete.DetermineAdvHour(&newModel, hour)
	assert.Equal(t, int64(0), newModel.AdvClick0To5HourCount)
	assert.Equal(t, int64(1), newModel.AdvClick6To11HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick12To17HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick18To23HourCount)
}

func Test_DetermineAdvHour_12To17Hours(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 11, 17, 11, 36, 651387237, time.UTC)
	var hour = int64(advModel_test.TrigerdTime.Hour())
	concrete.DetermineAdvHour(&newModel, hour)
	assert.Equal(t, int64(0), newModel.AdvClick0To5HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick6To11HourCount)
	assert.Equal(t, int64(1), newModel.AdvClick12To17HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick18To23HourCount)
}

func Test_DetermineAdvHour_18To23Hours(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 11, 23, 11, 36, 651387237, time.UTC)
	var hour = int64(advModel_test.TrigerdTime.Hour())
	concrete.DetermineAdvHour(&newModel, hour)
	assert.Equal(t, int64(0), newModel.AdvClick0To5HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick6To11HourCount)
	assert.Equal(t, int64(0), newModel.AdvClick12To17HourCount)
	assert.Equal(t, int64(1), newModel.AdvClick18To23HourCount)
}

func TestDetermineAdvAmPm_Am(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 11, 12, 11, 36, 651387237, time.UTC)
	var hour = int64(advModel_test.TrigerdTime.Hour())
	concrete.DetermineAdvAmPm(&newModel, hour)
	assert.Equal(t, int64(1), newModel.AmAdvClickCount)
	assert.Equal(t, int64(0), newModel.PmAdvClickCount)
}

func TestDetermineAdvAmPm_Pm(t *testing.T) {
	newModel := model.AdvEventRespondModel{}
	var advModel_test = advModel
	advModel_test.TrigerdTime = time.Date(
		2021, 12, 11, 13, 11, 36, 651387237, time.UTC)
	var hour = int64(advModel_test.TrigerdTime.Hour())
	concrete.DetermineAdvAmPm(&newModel, hour)
	assert.Equal(t, int64(0), newModel.AmAdvClickCount)
	assert.Equal(t, int64(1), newModel.PmAdvClickCount)
}

func TestCalculateAdvLevelBasedAvgClickCount_ZeroLevelIndex(t *testing.T) {
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.LevelIndex = 0
	responseAdvModel_test.TotalAdvCount = 100
	var expCount int64 = 100
	concrete.CalculateAdvLevelBasedAvgClickCount(&responseAdvModel_test)
	assert.Equal(t, expCount, responseAdvModel_test.LevelBasedAverageAdvCount)
}

func TestCalculateAdvLevelBasedAvgClickCount_NotZeroLevelIndex(t *testing.T) {
	var responseAdvModel_test = responseAdvModel
	responseAdvModel_test.LevelIndex = 10
	responseAdvModel_test.TotalAdvCount = 100
	var expCount int64 = 10
	concrete.CalculateAdvLevelBasedAvgClickCount(&responseAdvModel_test)
	assert.Equal(t, expCount, responseAdvModel_test.LevelBasedAverageAdvCount)
}
