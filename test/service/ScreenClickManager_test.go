package test

import (
	"FilterWorkerService/internal/IoC"
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

var clickModel = model.ScreenClickModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	StartXCor:  100,
	StartYCor:  200,
	FinishXCor: 300,
	FinishYCor: 400,
	TouchCount: 150,
	FingerId:   1,
	LevelIndex: 5,
	LevelName:  "5",
	CreationAt: time.Date(
		2021, 11, 13, 19, 34, 36, 651387237, time.UTC),
}

var clickRespondModel = model.ScreenClickRespondModel{}

var firstclick = time.Date(
	2021, 11, 6, 18, 34, 58, 651387237, time.UTC)

var last = time.Date(
	2021, 11, 7, 18, 34, 58, 651387237, time.UTC)

var clickOldModel = model.ScreenClickModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	StartXCor:  555,
	StartYCor:  222,
	FinishXCor: 333,
	FinishYCor: 444,
	TouchCount: 666,
	FingerId:   99,
	LevelIndex: 20,
	LevelName:  "20",
	CreationAt: firstclick,
}

var clickOldModel2 = model.ScreenClickModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	StartXCor:  99,
	StartYCor:  999,
	FinishXCor: 9,
	FinishYCor: 10,
	TouchCount: 20,
	FingerId:   30,
	LevelIndex: 40,
	LevelName:  "40",
	CreationAt: last,
}

var TotalClickDay = int64((last.YearDay() - firstclick.YearDay()) + 365*(last.Year()-firstclick.Year()))
var TotalClickHour = int64(((last.YearDay()+365*last.Year())*24 + last.Hour()) - ((firstclick.YearDay()+365*firstclick.Year())*24 + firstclick.Hour()))
var TotalClickMinute = int64((((last.YearDay()+365*last.Year())*24+last.Hour())*60 + last.Minute()) - (((firstclick.YearDay()+365*firstclick.Year())*24+firstclick.Hour())*60 + firstclick.Minute()))

var screenClickOldModel = model.ScreenClickRespondModel{}

var TotalClickDay2 = (clickRespondModel.FirstClickSessionYearOfDay - screenClickOldModel.FirstClickSessionYearOfDay) + 365*(clickRespondModel.FirstClickSessionYear-screenClickOldModel.FirstClickSessionYear)
var TotalClickCount2 = screenClickOldModel.TotalClickCount + clickRespondModel.TotalClickCount
var TotalClickSessionCount2 = screenClickOldModel.TotalClickSessionCount + clickRespondModel.TotalClickSessionCount
var TotalClickHour2 = ((clickRespondModel.FirstClickSessionYearOfDay+365*clickRespondModel.FirstClickSessionYear)*24 + clickRespondModel.FirstClickSessionHour) - ((screenClickOldModel.FirstClickSessionYearOfDay+365*screenClickOldModel.FirstClickSessionYear)*24 + screenClickOldModel.FirstClickSessionHour)
var TotalClickMinute2 = (((clickRespondModel.FirstClickSessionYearOfDay+365*clickRespondModel.FirstClickSessionYear)*24+clickRespondModel.FirstClickSessionHour)*60 + clickRespondModel.FirstClickSessionMinute) - (((screenClickOldModel.FirstClickSessionYearOfDay+365*screenClickOldModel.FirstClickSessionYear)*24+screenClickOldModel.FirstClickSessionHour)*60 + screenClickOldModel.FirstClickSessionMinute)

var Hour = clickRespondModel.FirstClickSessionHour
var Minute = clickRespondModel.FirstClickSessionMinute
var Count = clickRespondModel.FirstTouchCount
var StartXCor = clickRespondModel.FirstStartXCor
var StartYCor = clickRespondModel.FirstStartYCor
var FinishXCor = clickRespondModel.FirstFinishXCor
var FinishYCor = clickRespondModel.FirstFinishYCor

var TotalClickCount = int64(clickOldModel.TouchCount) + int64(clickOldModel2.TouchCount) + clickRespondModel.FirstTouchCount

var updatedClickModel = model.ScreenClickRespondModel{
	ProjectId:                     "Test",
	ClientId:                      "Test",
	CustomerId:                    "Test",
	LevelIndex:                    clickRespondModel.LevelIndex,
	FirstClickSessionYearOfDay:    int64(clickOldModel.CreationAt.YearDay()),
	FirstClickSessionYear:         int64(clickOldModel.CreationAt.Year()),
	FirstClickSessionHour:         int64(clickOldModel.CreationAt.Hour()),
	FirstClickSessionMinute:       int64(clickOldModel.CreationAt.Minute()),
	FirstTouchCount:               int64(clickOldModel.TouchCount),
	SecondClickSessionHour:        0,
	SecondClickSessionMinute:      0,
	SecondTouchCount:              0,
	ThirdClickSessionHour:         0,
	ThirdClickSessionMinute:       0,
	ThirdTouchCount:               0,
	FourthClickSessionHour:        0,
	FourthClickSessionMinute:      0,
	FourthTouchCount:              0,
	FifthClickSessionHour:         Hour,
	FifthClickSessionMinute:       Minute,
	FifthTouchCount:               Count,
	PenultimateClickSessionHour:   int64(last.Hour()),
	PenultimateClickSessionMinute: int64(last.Minute()),
	PenultimateTouchCount:         int64(clickOldModel2.TouchCount),
	LastClickSessionYearOfDay:     clickRespondModel.FirstClickSessionYearOfDay,
	LastClickSessionYear:          clickRespondModel.FirstClickSessionYear,
	LastClickSessionHour:          clickRespondModel.FirstClickSessionHour,
	LastClickSessionMinute:        clickRespondModel.FirstClickSessionMinute,
	LastTouchCount:                clickRespondModel.FirstTouchCount,
	FirstStartXCor:                clickOldModel.StartXCor,
	FirstStartYCor:                clickOldModel.StartYCor,
	FirstFinishXCor:               clickOldModel.FinishXCor,
	FirstFinishYCor:               clickOldModel.FinishYCor,
	SecondStartXCor:               0,
	SecondStartYCor:               0,
	SecondFinishXCor:              0,
	SecondFinishYCor:              0,
	ThirdStartXCor:                0,
	ThirdStartYCor:                0,
	ThirdFinishXCor:               0,
	ThirdFinishYCor:               0,
	FourthStartXCor:               0,
	FourthStartYCor:               0,
	FourthFinishXCor:              0,
	FourthFinishYCor:              0,
	FifthStartXCor:                StartXCor,
	FifthStartYCor:                StartYCor,
	FifthFinishXCor:               FinishXCor,
	FifthFinishYCor:               FinishYCor,
	PenultimateStartXCor:          clickOldModel2.StartXCor,
	PenultimateStartYCor:          clickOldModel2.StartYCor,
	PenultimateFinishXCor:         clickOldModel2.FinishXCor,
	PenultimateFinishYCor:         clickOldModel2.FinishYCor,
	LastStartXCor:                 clickRespondModel.FirstStartXCor,
	LastStartYCor:                 clickRespondModel.FirstStartYCor,
	LastFinishXCor:                clickRespondModel.FirstFinishXCor,
	LastFinishYCor:                clickRespondModel.FirstFinishYCor,
	FirstHalfHourTouchCount:       int64(clickOldModel.TouchCount),
	FirstHourTouchCount:           int64(clickOldModel.TouchCount),
	FirstTwoHourTouchCount:        int64(clickOldModel.TouchCount),
	FirstThreeHourTouchCount:      int64(clickOldModel.TouchCount),
	FirstSixHourTouchCount:        int64(clickOldModel.TouchCount),
	FirstTwelveHourTouchCount:     int64(clickOldModel.TouchCount),
	FirstMinusLastTouchCount:      int64(clickOldModel.TouchCount) - clickRespondModel.FirstTouchCount,
	FirstFingerId:                 int64(clickOldModel.FingerId),
	PenultimateFingerId:           int64(clickOldModel2.FingerId),
	LastFingerId:                  clickRespondModel.FirstFingerId,
	FirstDayClickCount:            int64(clickOldModel.TouchCount),
	SecondDayClickCount:           clickRespondModel.FirstTouchCount,
	ThirdDayClickCount:            0,
	FourthDayClickCount:           0,
	FifthDayClickCount:            0,
	SixthDayClickCount:            0,
	SeventhDayClickCount:          0,
	TotalClickDay:                 TotalClickDay2,
	TotalClickCount:               int64(clickOldModel.TouchCount) + int64(clickOldModel2.TouchCount) + clickRespondModel.FirstTouchCount,
	TotalClickSessionCount:        TotalClickSessionCount2,
	TotalClickHour:                TotalClickHour2,
	TotalClickMinute:              TotalClickMinute2,
	TotalStartXCor:                clickOldModel.StartXCor + clickOldModel2.StartXCor + clickRespondModel.FirstStartXCor,
	TotalStartYCor:                clickOldModel.StartYCor + clickOldModel2.StartYCor + clickRespondModel.FirstStartYCor,
	TotalFinishXCor:               clickOldModel.FinishXCor + clickOldModel2.FinishXCor + clickRespondModel.FirstFinishXCor,
	TotalFinishYCor:               clickOldModel.FinishYCor + clickOldModel2.FinishYCor + clickRespondModel.FirstFinishYCor,
	SessionBasedAvegareStartXCor:  (clickOldModel.StartXCor + clickOldModel2.StartXCor + clickRespondModel.FirstStartXCor) / float64(TotalClickSessionCount2),
	SessionBasedAvegareStartYCor:  (clickOldModel.StartYCor + clickOldModel2.StartYCor + clickRespondModel.FirstStartYCor) / float64(TotalClickSessionCount2),
	SessionBasedAvegareFinishXCor: (clickOldModel.FinishXCor + clickOldModel2.FinishXCor + clickRespondModel.FirstFinishXCor) / float64(TotalClickSessionCount2),
	SessionBasedAvegareFinishYCor: (clickOldModel.FinishYCor + clickOldModel2.FinishYCor + clickRespondModel.FirstFinishYCor) / float64(TotalClickSessionCount2),
	SessionBasedAvegareClickCount: float64(TotalClickCount) / float64(TotalClickSessionCount2),
	DailyAvegareClickCount:        float64(TotalClickCount) / float64(TotalClickDay2),
	LastTouchCountMinusSessionBasedAvegareClickCount: float64(clickRespondModel.FirstTouchCount) - float64(TotalClickCount)/float64(TotalClickSessionCount2),
}

func Test_UpdateScreenClick_Success(t *testing.T) {

	var testClickDal = new(repository.MockScreenClickDal)
	var testCache = new(service.MockCacheService)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.ScreenClickDal = testClickDal
	IoC.CacheService = testCache
	var manager = concrete.ScreenClickManagerConstructor()
	var clickModel_test = clickModel
	clickModel_test.CreationAt = time.Date(
		2021, 11, 5, 16, 11, 36, 651387237, time.UTC)

	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.ProjectId                     =  "Test"
	screenClickOldModel_test.ClientId                      =  "Test"
	screenClickOldModel_test.CustomerId                    =  "Test"
	screenClickOldModel_test.LevelIndex                    =  int64(clickOldModel.LevelIndex)
	screenClickOldModel_test.FirstClickSessionYearOfDay    =  int64(clickOldModel.CreationAt.YearDay())
	screenClickOldModel_test.FirstClickSessionYear         =  int64(clickOldModel.CreationAt.Year())
	screenClickOldModel_test.FirstClickSessionHour         =  int64(clickOldModel.CreationAt.Hour())
	screenClickOldModel_test.FirstClickSessionMinute       =  int64(clickOldModel.CreationAt.Minute())
	screenClickOldModel_test.FirstTouchCount               =  int64(clickOldModel.TouchCount)
	screenClickOldModel_test.LastClickSessionYearOfDay    =  int64(last.YearDay())
	screenClickOldModel_test.LastClickSessionYear         =  int64(last.Year())
	screenClickOldModel_test.LastClickSessionHour         =  int64(last.Hour())
	screenClickOldModel_test.LastClickSessionMinute       =  int64(last.Minute())
	screenClickOldModel_test.LastTouchCount               =  int64(clickOldModel2.TouchCount)
	screenClickOldModel_test.FirstStartXCor               =  clickOldModel.StartXCor
	screenClickOldModel_test.FirstStartYCor               =  clickOldModel.StartYCor
	screenClickOldModel_test.FirstFinishXCor              =  clickOldModel.FinishXCor
	screenClickOldModel_test.FirstFinishYCor              =  clickOldModel.FinishYCor
	screenClickOldModel_test.LastStartXCor                   = clickOldModel2.StartXCor
	screenClickOldModel_test.LastStartYCor                   = clickOldModel2.StartYCor
	screenClickOldModel_test.LastFinishXCor                  = clickOldModel2.FinishXCor
	screenClickOldModel_test.LastFinishYCor                  = clickOldModel2.FinishYCor
	screenClickOldModel_test.FirstHalfHourTouchCount         = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.FirstHourTouchCount             = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.FirstTwoHourTouchCount          = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.FirstThreeHourTouchCount        = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.FirstSixHourTouchCount          = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.FirstTwelveHourTouchCount       = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.FirstMinusLastTouchCount        = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.FirstFingerId                   = int64(clickOldModel.FingerId)
	screenClickOldModel_test.PenultimateFingerId             = 0
	screenClickOldModel_test.LastFingerId                    = int64(clickOldModel2.FingerId)
	screenClickOldModel_test.FirstDayClickCount              = int64(clickOldModel.TouchCount)
	screenClickOldModel_test.SecondDayClickCount             = 0
	screenClickOldModel_test.ThirdDayClickCount              = 0
	screenClickOldModel_test.FourthDayClickCount             = 0
	screenClickOldModel_test.FifthDayClickCount              = 0
	screenClickOldModel_test.SixthDayClickCount              = 0
	screenClickOldModel_test.SeventhDayClickCount            = 0
	screenClickOldModel_test.TotalClickDay                   = TotalClickDay
	screenClickOldModel_test.TotalClickCount                 = int64(clickOldModel.TouchCount) + int64(clickOldModel2.TouchCount)
	screenClickOldModel_test.TotalClickSessionCount          = 4
	screenClickOldModel_test.TotalClickHour                  = TotalClickHour
	screenClickOldModel_test.TotalClickMinute                = TotalClickMinute
	screenClickOldModel_test.TotalStartXCor                  = clickOldModel.StartXCor + clickOldModel2.StartXCor
	screenClickOldModel_test.TotalStartYCor                  = clickOldModel.StartYCor + clickOldModel2.StartYCor
	screenClickOldModel_test.TotalFinishXCor                 = clickOldModel.FinishXCor + clickOldModel2.FinishXCor
	screenClickOldModel_test.TotalFinishYCor                 = clickOldModel.FinishYCor + clickOldModel2.FinishYCor
	screenClickOldModel_test.SessionBasedAvegareStartXCor    = clickOldModel.StartXCor
	screenClickOldModel_test.SessionBasedAvegareStartYCor    = clickOldModel.StartYCor
	screenClickOldModel_test.SessionBasedAvegareFinishXCor   = clickOldModel.FinishXCor
	screenClickOldModel_test.SessionBasedAvegareFinishYCor   = clickOldModel.FinishYCor
	screenClickOldModel_test.SessionBasedAvegareClickCount   = float64(clickOldModel.TouchCount)
	screenClickOldModel_test.DailyAvegareClickCount          = float64(clickOldModel.TouchCount)

	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.ProjectId                =     "Test"
	clickRespondModel_test.ClientId                 =     "Test"
	clickRespondModel_test.CustomerId               =     "Test"
	clickRespondModel_test.LevelIndex               =     int64(clickOldModel.LevelIndex)
	clickRespondModel_test.FirstClickSessionYearOfDay = int64(clickModel_test.CreationAt.YearDay())
	clickRespondModel_test.FirstClickSessionYear = int64(clickModel_test.CreationAt.Year())
	clickRespondModel_test.FirstClickSessionHour = int64(clickModel_test.CreationAt.Hour())
	clickRespondModel_test.FirstClickSessionMinute = int64(clickModel_test.CreationAt.Minute())
	clickRespondModel_test.FirstTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstStartXCor = clickModel_test.StartXCor
	clickRespondModel_test.FirstStartYCor = clickModel_test.StartYCor
	clickRespondModel_test.FirstFinishXCor = clickModel_test.FinishXCor
	clickRespondModel_test.FirstFinishYCor = clickModel_test.FinishYCor
	clickRespondModel_test.FirstHalfHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstTwoHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstThreeHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstSixHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstTwelveHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstMinusLastTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstFingerId = int64(clickModel_test.FingerId)
	clickRespondModel_test.FirstDayClickCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.TotalClickCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.TotalStartXCor = clickModel_test.StartXCor
	clickRespondModel_test.TotalStartYCor = clickModel_test.StartYCor
	clickRespondModel_test.TotalFinishXCor = clickModel_test.FinishXCor
	clickRespondModel_test.TotalFinishYCor = clickModel_test.FinishYCor
	clickRespondModel_test.SessionBasedAvegareStartXCor = (clickModel_test.StartXCor)
	clickRespondModel_test.SessionBasedAvegareStartYCor = (clickModel_test.StartYCor)
	clickRespondModel_test.SessionBasedAvegareFinishXCor = (clickModel_test.FinishXCor)
	clickRespondModel_test.SessionBasedAvegareFinishYCor = (clickModel_test.FinishYCor)
	clickRespondModel_test.SessionBasedAvegareClickCount = float64(clickModel_test.TouchCount)
	clickRespondModel_test.DailyAvegareClickCount = float64(clickModel_test.TouchCount)

	// var TotalClickDay2 = (clickRespondModel_test.FirstClickSessionYearOfDay - screenClickOldModel_test.FirstClickSessionYearOfDay) + 365*(clickRespondModel_test.FirstClickSessionYear-screenClickOldModel_test.FirstClickSessionYear)
	// var TotalClickCount2 = screenClickOldModel_test.TotalClickCount + clickRespondModel_test.TotalClickCount
	// var TotalClickSessionCount2 = screenClickOldModel_test.TotalClickSessionCount + clickRespondModel_test.TotalClickSessionCount
	// var TotalClickHour2 = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	// var TotalClickMinute2 = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)

	// var Hour = clickRespondModel.FirstClickSessionHour
	// var Minute = clickRespondModel.FirstClickSessionMinute
	// var Count = clickRespondModel.FirstTouchCount
	// var StartXCor = clickRespondModel.FirstStartXCor
	// var StartYCor = clickRespondModel.FirstStartYCor
	// var FinishXCor = clickRespondModel.FirstFinishXCor
	// var FinishYCor = clickRespondModel.FirstFinishYCor

	var updatedClickModel_test = updatedClickModel
	// updatedClickModel_test.ProjectId                   =  "Test"
	// updatedClickModel_test.ClientId                    =  "Test"
	// updatedClickModel_test.CustomerId                  =  "Test"
	// updatedClickModel_test.LevelIndex                  =  clickRespondModel.LevelIndex
	// updatedClickModel_test.FirstClickSessionYearOfDay  =  int64(clickOldModel.CreationAt.YearDay())
	// updatedClickModel_test.FirstClickSessionYear       =  int64(clickOldModel.CreationAt.Year())
	// updatedClickModel_test.FirstClickSessionHour       =  int64(clickOldModel.CreationAt.Hour())
	// updatedClickModel_test.FirstClickSessionMinute     =  int64(clickOldModel.CreationAt.Minute())
	// updatedClickModel_test.FirstTouchCount             =  int64(clickOldModel.TouchCount)

	testClickDal.On("UpdateScreenClickById", screenClickOldModel_test.ClientId, &updatedClickModel_test).Return(nil)
	uptModel, s, e := manager.UpdateScreenClick(&clickRespondModel_test, &screenClickOldModel_test)
	assert.Equal(t, &updatedClickModel_test, uptModel)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, e)
}

func Test_ConvertRawModelToResponseModel_Add_Succes(t *testing.T) {

	var testClickDal = new(repository.MockScreenClickDal)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.ScreenClickDal = testClickDal
	var manager = concrete.ScreenClickManagerConstructor()
	var clickModel_test = clickModel
	var screenClickOldModel_test = screenClickOldModel
	var clickRespondModel_test = clickRespondModel

	clickModel_test.CreationAt = time.Date(
		2021, 11, 5, 16, 11, 36, 651387237, time.UTC)

	clickRespondModel_test.ProjectId                =     "Test"
	clickRespondModel_test.ClientId                 =     "Test"
	clickRespondModel_test.CustomerId               =     "Test"
	clickRespondModel_test.LevelIndex               =     int64(clickOldModel.LevelIndex)
	clickRespondModel_test.FirstClickSessionYearOfDay = int64(clickModel_test.CreationAt.YearDay())
	clickRespondModel_test.FirstClickSessionYear = int64(clickModel_test.CreationAt.Year())
	clickRespondModel_test.FirstClickSessionHour = int64(clickModel_test.CreationAt.Hour())
	clickRespondModel_test.FirstClickSessionMinute = int64(clickModel_test.CreationAt.Minute())
	clickRespondModel_test.FirstTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstStartXCor = clickModel_test.StartXCor
	clickRespondModel_test.FirstStartYCor = clickModel_test.StartYCor
	clickRespondModel_test.FirstFinishXCor = clickModel_test.FinishXCor
	clickRespondModel_test.FirstFinishYCor = clickModel_test.FinishYCor
	clickRespondModel_test.FirstHalfHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstTwoHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstThreeHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstSixHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstTwelveHourTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstMinusLastTouchCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.FirstFingerId = int64(clickModel_test.FingerId)
	clickRespondModel_test.FirstDayClickCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.TotalClickCount = int64(clickModel_test.TouchCount)
	clickRespondModel_test.TotalStartXCor = clickModel_test.StartXCor
	clickRespondModel_test.TotalStartYCor = clickModel_test.StartYCor
	clickRespondModel_test.TotalFinishXCor = clickModel_test.FinishXCor
	clickRespondModel_test.TotalFinishYCor = clickModel_test.FinishYCor
	clickRespondModel_test.SessionBasedAvegareStartXCor = (clickModel_test.StartXCor)
	clickRespondModel_test.SessionBasedAvegareStartYCor = (clickModel_test.StartYCor)
	clickRespondModel_test.SessionBasedAvegareFinishXCor = (clickModel_test.FinishXCor)
	clickRespondModel_test.SessionBasedAvegareFinishYCor = (clickModel_test.FinishYCor)
	clickRespondModel_test.SessionBasedAvegareClickCount = float64(clickModel_test.TouchCount)
	clickRespondModel_test.DailyAvegareClickCount = float64(clickModel_test.TouchCount)

	testClickDal.On("Add", &clickRespondModel_test).Return(nil)
	testClickDal.On("GetScreenClickById", screenClickOldModel_test.ClientId).Return(&screenClickOldModel_test,
		errors.New("null data error"))
	var clickModel_test_byte, _ = json.EncodeJson(clickModel_test)
	var v, s, m = manager.ConvertRawModelToResponseModel(clickModel_test_byte)
	var value, success = v.(model.BuyingEventRespondModel)
	if success == true {
		assert.Equal(t, &clickRespondModel_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
}

func Test_CalculateSeventhDayClickCount_In144To168Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 17
	screenClickOldModel_test.SeventhDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 209
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 16
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateSeventhDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324 + 123

	assert.Equal(t, Expcount, screenClickOldModel_test.SeventhDayClickCount)
}

func Test_CalculateSeventhDayClickCount_Out144To168Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 17
	screenClickOldModel_test.SeventhDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 209
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 18
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateSeventhDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324

	assert.Equal(t, Expcount, screenClickOldModel_test.SeventhDayClickCount)
}

func Test_CalculateSixthDayClickCount_In120To144Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 17
	screenClickOldModel_test.SixthDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 208
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 17
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateSixthDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324 + 123
	assert.Equal(t, Expcount, screenClickOldModel_test.SixthDayClickCount)
}

func Test_CalculateSixthDayClickCount_Out120To144Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 17
	screenClickOldModel_test.SixthDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 208
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 18
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateSixthDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324
	assert.Equal(t, Expcount, screenClickOldModel_test.SixthDayClickCount)
}

func Test_CalculateFifthDayClickCount_In96To120Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 22
	screenClickOldModel_test.FifthDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 207
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 22
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateFifthDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324 + 123
	assert.Equal(t, Expcount, screenClickOldModel_test.FifthDayClickCount)
}

func Test_CalculateFifthDayClickCount_Out96To120Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 17
	screenClickOldModel_test.FifthDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 207
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 18
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateFifthDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324
	assert.Equal(t, Expcount, screenClickOldModel_test.FifthDayClickCount)
}

func Test_CalculateFourthDayClickCount_In72To96Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.FourthDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 206
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 18
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateFourthDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324 + 123
	assert.Equal(t, Expcount, screenClickOldModel_test.FourthDayClickCount)
}

func Test_CalculateFourthDayClickCount_Out72To96Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.FourthDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 206
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 19
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateFourthDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324
	assert.Equal(t, Expcount, screenClickOldModel_test.FourthDayClickCount)
}

func Test_calculateThirdDayClickCount_In48To72Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.ThirdDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 205
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 18
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateThirdDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324 + 123
	assert.Equal(t, Expcount, screenClickOldModel_test.ThirdDayClickCount)
}

func Test_calculateThirdDayClickCount_Out48To72Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.ThirdDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 205
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 19
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateThirdDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324
	assert.Equal(t, Expcount, screenClickOldModel_test.ThirdDayClickCount)
}

func Test_CalculateSecondDayClickCount_In24To48Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.SecondDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 204
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 18
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateSecondDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324 + 123
	assert.Equal(t, Expcount, screenClickOldModel_test.SecondDayClickCount)
}

func Test_CalculateSecondDayClickCount_Out24To48Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.SecondDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 204
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 19
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateSecondDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324
	assert.Equal(t, Expcount, screenClickOldModel_test.SecondDayClickCount)
}

func Test_calculateFirstDayClickCount_In24Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.FirstDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 203
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 18
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateFirstDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324 + 123
	assert.Equal(t, Expcount, screenClickOldModel_test.FirstDayClickCount)
}

func Test_calculateFirstDayClickCount_Out24Hours(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 18
	screenClickOldModel_test.FirstDayClickCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 203
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 19
	clickRespondModel_test.FirstDayClickCount = 123
	var TotalClickHour_test = ((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24 + clickRespondModel_test.FirstClickSessionHour) - ((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24 + screenClickOldModel_test.FirstClickSessionHour)
	concrete.CalculateFirstDayClickCount(&clickRespondModel_test, &screenClickOldModel_test, TotalClickHour_test)
	var Expcount int64 = 324
	assert.Equal(t, Expcount, screenClickOldModel_test.FirstDayClickCount)
}

func Test_CalculateClickCount_FifthClick(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.TotalClickSessionCount = 5
	screenClickOldModel_test.FifthClickSessionHour = 13
	screenClickOldModel_test.FifthClickSessionMinute = 48
	screenClickOldModel_test.FifthTouchCount = 176
	screenClickOldModel_test.FifthStartXCor = 0.3
	screenClickOldModel_test.FifthStartYCor = 0.7
	screenClickOldModel_test.FifthFinishXCor = 0.4
	screenClickOldModel_test.FifthFinishYCor = 0.1
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionHour = 21
	clickRespondModel_test.FirstClickSessionMinute = 39
	clickRespondModel_test.FirstTouchCount = 430
	clickRespondModel_test.FirstStartXCor = 0.4
	clickRespondModel_test.FirstStartYCor = 0.5
	clickRespondModel_test.FirstFinishXCor = 0.2
	clickRespondModel_test.FirstFinishYCor = 0.9
	concrete.CalculateClickCount(&clickRespondModel_test, &screenClickOldModel_test)
	assert.Equal(t, 21, screenClickOldModel_test.FifthClickSessionHour)
	assert.Equal(t, 39, screenClickOldModel_test.FifthClickSessionMinute)
	assert.Equal(t, 430, screenClickOldModel_test.FifthTouchCount)
	assert.Equal(t, 0.4, screenClickOldModel_test.FifthStartXCor)
	assert.Equal(t, 0.5, screenClickOldModel_test.FifthStartYCor)
	assert.Equal(t, 0.2, screenClickOldModel_test.FifthFinishXCor)
	assert.Equal(t, 0.9, screenClickOldModel_test.FifthFinishYCor)
}

func Test_CalculateClickCount_FourthClick(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.TotalClickSessionCount = 4
	screenClickOldModel_test.FourthClickSessionHour = 13
	screenClickOldModel_test.FourthClickSessionMinute = 48
	screenClickOldModel_test.FourthTouchCount = 176
	screenClickOldModel_test.FourthStartXCor = 0.3
	screenClickOldModel_test.FourthStartYCor = 0.7
	screenClickOldModel_test.FourthFinishXCor = 0.4
	screenClickOldModel_test.FourthFinishYCor = 0.1
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionHour = 21
	clickRespondModel_test.FirstClickSessionMinute = 39
	clickRespondModel_test.FirstTouchCount = 430
	clickRespondModel_test.FirstStartXCor = 0.4
	clickRespondModel_test.FirstStartYCor = 0.5
	clickRespondModel_test.FirstFinishXCor = 0.2
	clickRespondModel_test.FirstFinishYCor = 0.9
	concrete.CalculateClickCount(&clickRespondModel_test, &screenClickOldModel_test)
	assert.Equal(t, 21, screenClickOldModel.FourthClickSessionHour)
	assert.Equal(t, 39, screenClickOldModel.FourthClickSessionMinute)
	assert.Equal(t, 430, screenClickOldModel.FourthTouchCount)
	assert.Equal(t, 0.4, screenClickOldModel.FourthStartXCor)
	assert.Equal(t, 0.5, screenClickOldModel.FourthStartYCor)
	assert.Equal(t, 0.2, screenClickOldModel.FourthFinishXCor)
	assert.Equal(t, 0.9, screenClickOldModel.FourthFinishYCor)
}

func Test_CalculateClickCount_ThirdClick(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.TotalClickSessionCount = 3
	screenClickOldModel_test.ThirdClickSessionHour = 13
	screenClickOldModel_test.ThirdClickSessionMinute = 48
	screenClickOldModel_test.ThirdTouchCount = 176
	screenClickOldModel_test.ThirdStartXCor = 0.3
	screenClickOldModel_test.ThirdStartYCor = 0.7
	screenClickOldModel_test.ThirdFinishXCor = 0.4
	screenClickOldModel_test.ThirdFinishYCor = 0.1
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionHour = 21
	clickRespondModel_test.FirstClickSessionMinute = 39
	clickRespondModel_test.FirstTouchCount = 430
	clickRespondModel_test.FirstStartXCor = 0.4
	clickRespondModel_test.FirstStartYCor = 0.5
	clickRespondModel_test.FirstFinishXCor = 0.2
	clickRespondModel_test.FirstFinishYCor = 0.9
	concrete.CalculateClickCount(&clickRespondModel_test, &screenClickOldModel_test)
	assert.Equal(t, 21, screenClickOldModel_test.ThirdClickSessionHour)
	assert.Equal(t, 39, screenClickOldModel_test.ThirdClickSessionMinute)
	assert.Equal(t, 430, screenClickOldModel_test.ThirdTouchCount)
	assert.Equal(t, 0.4, screenClickOldModel_test.ThirdStartXCor)
	assert.Equal(t, 0.5, screenClickOldModel_test.ThirdStartYCor)
	assert.Equal(t, 0.2, screenClickOldModel_test.ThirdFinishXCor)
	assert.Equal(t, 0.9, screenClickOldModel_test.ThirdFinishYCor)
}

func Test_CalculateClickCount_SecondClick(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.TotalClickSessionCount = 3
	screenClickOldModel_test.SecondClickSessionHour = 13
	screenClickOldModel_test.SecondClickSessionMinute = 48
	screenClickOldModel_test.SecondTouchCount = 176
	screenClickOldModel_test.SecondStartXCor = 0.3
	screenClickOldModel_test.SecondStartYCor = 0.7
	screenClickOldModel_test.SecondFinishXCor = 0.4
	screenClickOldModel_test.SecondFinishYCor = 0.1
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionHour = 21
	clickRespondModel_test.FirstClickSessionMinute = 39
	clickRespondModel_test.FirstTouchCount = 430
	clickRespondModel_test.FirstStartXCor = 0.4
	clickRespondModel_test.FirstStartYCor = 0.5
	clickRespondModel_test.FirstFinishXCor = 0.2
	clickRespondModel_test.FirstFinishYCor = 0.9
	concrete.CalculateClickCount(&clickRespondModel_test, &screenClickOldModel_test)
	assert.Equal(t, 21, screenClickOldModel_test.SecondClickSessionHour)
	assert.Equal(t, 39, screenClickOldModel_test.SecondClickSessionMinute)
	assert.Equal(t, 430, screenClickOldModel_test.SecondTouchCount)
	assert.Equal(t, 0.4, screenClickOldModel_test.SecondStartXCor)
	assert.Equal(t, 0.5, screenClickOldModel_test.SecondStartYCor)
	assert.Equal(t, 0.2, screenClickOldModel_test.SecondFinishXCor)
	assert.Equal(t, 0.9, screenClickOldModel_test.SecondFinishYCor)
}

func Test_CalculateClickTwentyHour_In720Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstTwelveHourTouchCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 22
	clickRespondModel_test.FirstClickSessionMinute = 50
	clickRespondModel_test.FirstTouchCount = 123
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickTwelveHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)

	var ExpCount int64 = 324 + 123
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstTwelveHourTouchCount)
}

func Test_CalculateClickTwentyHour_Out720Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstTwelveHourTouchCount = 324
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 22
	clickRespondModel_test.FirstClickSessionMinute = 51
	clickRespondModel_test.FirstTouchCount = 123
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickTwelveHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)

	var ExpCount int64 = 324
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstTwelveHourTouchCount)
}

func Test_CalculateClickTwoHour_In120Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstTwoHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 12
	clickRespondModel_test.FirstClickSessionMinute = 50
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickTwoHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666 + 150
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstTwoHourTouchCount)
}

func Test_CalculateClickTwoHour_Out120Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstTwoHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 12
	clickRespondModel_test.FirstClickSessionMinute = 51
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickTwoHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstTwoHourTouchCount)
}

func Test_CalculateClickThreeHour_In180Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstThreeHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 13
	clickRespondModel_test.FirstClickSessionMinute = 50
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickThreeHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	fmt.Println(TotalClickMinute2)
	var ExpCount int64 = 666 + 150

	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstThreeHourTouchCount)
}

func Test_CalculateClickThreeHour_Out180Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstThreeHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 13
	clickRespondModel_test.FirstClickSessionMinute = 51
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickThreeHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666 + 150

	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstThreeHourTouchCount)
}

func Test_CalculateClickSixHour_In360Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstSixHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 16
	clickRespondModel_test.FirstClickSessionMinute = 50
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickSixHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666 + 150
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstSixHourTouchCount)
}

func Test_CalculateClickSixHour_Out360Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstSixHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 16
	clickRespondModel_test.FirstClickSessionMinute = 51
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickSixHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstSixHourTouchCount)
}

func Test_calculateClickHour_In60Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 11
	clickRespondModel_test.FirstClickSessionMinute = 50
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstHourTouchCount)
}

func Test_calculateClickHour_Out60Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 50
	screenClickOldModel_test.FirstHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 11
	clickRespondModel_test.FirstClickSessionMinute = 51
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstHourTouchCount)
}

func Test_CalculateClickHalfHour_In30Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 20
	screenClickOldModel_test.FirstHalfHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 10
	clickRespondModel_test.FirstClickSessionMinute = 50
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickHalfHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666 + 150
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstHalfHourTouchCount)
}

func Test_CalculateClickHalfHour_Out30Minutes(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.FirstClickSessionYearOfDay = 202
	screenClickOldModel_test.FirstClickSessionYear = 2021
	screenClickOldModel_test.FirstClickSessionHour = 10
	screenClickOldModel_test.FirstClickSessionMinute = 20
	screenClickOldModel_test.FirstHalfHourTouchCount = 666
	var clickRespondModel_test = clickRespondModel
	clickRespondModel_test.FirstClickSessionYearOfDay = 202
	clickRespondModel_test.FirstClickSessionYear = 2021
	clickRespondModel_test.FirstClickSessionHour = 10
	clickRespondModel_test.FirstClickSessionMinute = 51
	clickRespondModel_test.FirstTouchCount = 150
	var TotalClickMinute_test = (((clickRespondModel_test.FirstClickSessionYearOfDay+365*clickRespondModel_test.FirstClickSessionYear)*24+clickRespondModel_test.FirstClickSessionHour)*60 + clickRespondModel_test.FirstClickSessionMinute) - (((screenClickOldModel_test.FirstClickSessionYearOfDay+365*screenClickOldModel_test.FirstClickSessionYear)*24+screenClickOldModel_test.FirstClickSessionHour)*60 + screenClickOldModel_test.FirstClickSessionMinute)
	concrete.CalculateClickHalfHour(&clickRespondModel_test, &screenClickOldModel_test, TotalClickMinute_test)
	var ExpCount int64 = 666
	assert.Equal(t, ExpCount, screenClickOldModel_test.FirstHalfHourTouchCount)
}

func Test_CalculateDailyAverageClickCount_ZeroTotalClickDay(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.TotalClickDay = 0
	screenClickOldModel_test.TotalClickCount = 150
	var count = concrete.CalculateDailyAverageClickCount(&screenClickOldModel_test)
	var ExpCount float64 = float64(150)
	assert.Equal(t, ExpCount, count)
}

func Test_CalculateDailyAverageClickCount_NonZeroTotalClickDay(t *testing.T) {
	var screenClickOldModel_test = screenClickOldModel
	screenClickOldModel_test.TotalClickDay = 5
	screenClickOldModel_test.TotalClickCount = 150
	var count = concrete.CalculateDailyAverageClickCount(&screenClickOldModel_test)
	var ExpCount float64 = float64(150) / float64(5)
	assert.Equal(t, ExpCount, count)
}
