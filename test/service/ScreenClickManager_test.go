package test

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"errors"

	// "FilterWorkerService/pkg/jsonParser/gojson"
	// "FilterWorkerService/test/Mock/repository"
	// "errors"
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

var clickRespondModel = model.ScreenClickRespondModel{
	ProjectId:                     "Test",
	ClientId:                      "Test",
	CustomerId:                    "Test",
	LevelIndex:                    int64(clickModel.LevelIndex),
	FirstClickSessionYearOfDay:    int64(clickModel.CreationAt.YearDay()),
	FirstClickSessionYear:         int64(clickModel.CreationAt.Year()),
	FirstClickSessionHour:         int64(clickModel.CreationAt.Hour()),
	FirstClickSessionMinute:       int64(clickModel.CreationAt.Minute()),
	FirstTouchCount:               int64(clickModel.TouchCount),
	SecondClickSessionHour:        0,
	SecondClickSessionMinute:      0,
	SecondTouchCount:              0,
	ThirdClickSessionHour:         0,
	ThirdClickSessionMinute:       0,
	ThirdTouchCount:               0,
	FourthClickSessionHour:        0,
	FourthClickSessionMinute:      0,
	FourthTouchCount:              0,
	FifthClickSessionHour:         0,
	FifthClickSessionMinute:       0,
	FifthTouchCount:               0,
	PenultimateClickSessionHour:   0,
	PenultimateClickSessionMinute: 0,
	PenultimateTouchCount:         0,
	LastClickSessionYearOfDay:     0,
	LastClickSessionYear:          0,
	LastClickSessionHour:          0,
	LastClickSessionMinute:        0,
	LastTouchCount:                0,
	FirstStartXCor:                clickModel.StartXCor,
	FirstStartYCor:                clickModel.StartYCor,
	FirstFinishXCor:               clickModel.FinishXCor,
	FirstFinishYCor:               clickModel.FinishYCor,
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
	FifthStartXCor:                0,
	FifthStartYCor:                0,
	FifthFinishXCor:               0,
	FifthFinishYCor:               0,
	PenultimateStartXCor:          0,
	PenultimateStartYCor:          0,
	PenultimateFinishXCor:         0,
	PenultimateFinishYCor:         0,
	LastStartXCor:                 0,
	LastStartYCor:                 0,
	LastFinishXCor:                0,
	LastFinishYCor:                0,
	FirstHalfHourTouchCount:       int64(clickModel.TouchCount),
	FirstHourTouchCount:           int64(clickModel.TouchCount),
	FirstTwoHourTouchCount:        int64(clickModel.TouchCount),
	FirstThreeHourTouchCount:      int64(clickModel.TouchCount),
	FirstSixHourTouchCount:        int64(clickModel.TouchCount),
	FirstTwelveHourTouchCount:     int64(clickModel.TouchCount),
	FirstMinusLastTouchCount:      int64(clickModel.TouchCount),
	FirstFingerId:                 int64(clickModel.FingerId),
	PenultimateFingerId:           0,
	LastFingerId:                  0,
	FirstDayClickCount:            int64(clickModel.TouchCount),
	SecondDayClickCount:           0,
	ThirdDayClickCount:            0,
	FourthDayClickCount:           0,
	FifthDayClickCount:            0,
	SixthDayClickCount:            0,
	SeventhDayClickCount:          0,
	TotalClickDay:                 1,
	TotalClickCount:               int64(clickModel.TouchCount),
	TotalClickSessionCount:        1,
	TotalClickHour:                0,
	TotalClickMinute:              1,
	TotalStartXCor:                clickModel.StartXCor,
	TotalStartYCor:                clickModel.StartYCor,
	TotalFinishXCor:               clickModel.FinishXCor,
	TotalFinishYCor:               clickModel.FinishYCor,
	SessionBasedAvegareStartXCor:  clickModel.StartXCor,
	SessionBasedAvegareStartYCor:  clickModel.StartYCor,
	SessionBasedAvegareFinishXCor: clickModel.FinishXCor,
	SessionBasedAvegareFinishYCor: clickModel.FinishYCor,
	SessionBasedAvegareClickCount: float64(clickModel.TouchCount),
	DailyAvegareClickCount:        float64(clickModel.TouchCount),
	LastTouchCountMinusSessionBasedAvegareClickCount: 0,
}

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
var TotalClickHour = int64(((last.YearDay() + 365*last.Year())*24 + last.Hour()) - ((firstclick.YearDay() + 365*firstclick.Year())*24 + firstclick.Hour()))
var TotalClickMinute = int64((((last.YearDay() + 365*last.Year())*24 + last.Hour())*60 + last.Minute()) - (((firstclick.YearDay() + 365*firstclick.Year())*24 + firstclick.Hour())*60 + firstclick.Minute()))

var screenClickOldModel = model.ScreenClickRespondModel{
	ProjectId:                     "Test",
	ClientId:                      "Test",
	CustomerId:                    "Test",
	LevelIndex:                    int64(clickOldModel.LevelIndex),
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
	FifthClickSessionHour:         0,
	FifthClickSessionMinute:       0,
	FifthTouchCount:               0,
	PenultimateClickSessionHour:   0,
	PenultimateClickSessionMinute: 0,
	PenultimateTouchCount:         0,
	LastClickSessionYearOfDay:     int64(last.YearDay()),
	LastClickSessionYear:          int64(last.Year()),
	LastClickSessionHour:          int64(last.Hour()),
	LastClickSessionMinute:        int64(last.Minute()),
	LastTouchCount:                int64(clickOldModel2.TouchCount),
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
	FifthStartXCor:                0,
	FifthStartYCor:                0,
	FifthFinishXCor:               0,
	FifthFinishYCor:               0,
	PenultimateStartXCor:          0,
	PenultimateStartYCor:          0,
	PenultimateFinishXCor:         0,
	PenultimateFinishYCor:         0,
	LastStartXCor:                 clickOldModel2.StartXCor,
	LastStartYCor:                 clickOldModel2.StartYCor,
	LastFinishXCor:                clickOldModel2.FinishXCor,
	LastFinishYCor:                clickOldModel2.FinishYCor,
	FirstHalfHourTouchCount:       int64(clickOldModel.TouchCount),
	FirstHourTouchCount:           int64(clickOldModel.TouchCount),
	FirstTwoHourTouchCount:        int64(clickOldModel.TouchCount),
	FirstThreeHourTouchCount:      int64(clickOldModel.TouchCount),
	FirstSixHourTouchCount:        int64(clickOldModel.TouchCount),
	FirstTwelveHourTouchCount:     int64(clickOldModel.TouchCount),
	FirstMinusLastTouchCount:      int64(clickOldModel.TouchCount),
	FirstFingerId:                 int64(clickOldModel.FingerId),
	PenultimateFingerId:           0,
	LastFingerId:                  int64(clickOldModel2.FingerId),
	FirstDayClickCount:            int64(clickOldModel.TouchCount),
	SecondDayClickCount:           0,
	ThirdDayClickCount:            0,
	FourthDayClickCount:           0,
	FifthDayClickCount:            0,
	SixthDayClickCount:            0,
	SeventhDayClickCount:          0,
	TotalClickDay:                 TotalClickDay,
	TotalClickCount:               int64(clickOldModel.TouchCount) + int64(clickOldModel2.TouchCount),
	TotalClickSessionCount:        4,
	TotalClickHour:                TotalClickHour,
	TotalClickMinute:              TotalClickMinute,
	TotalStartXCor:                clickOldModel.StartXCor + clickOldModel2.StartXCor,
	TotalStartYCor:                clickOldModel.StartYCor + clickOldModel2.StartYCor,
	TotalFinishXCor:               clickOldModel.FinishXCor + clickOldModel2.FinishXCor,
	TotalFinishYCor:               clickOldModel.FinishYCor + clickOldModel2.FinishYCor,
	SessionBasedAvegareStartXCor:  clickOldModel.StartXCor,
	SessionBasedAvegareStartYCor:  clickOldModel.StartYCor,
	SessionBasedAvegareFinishXCor: clickOldModel.FinishXCor,
	SessionBasedAvegareFinishYCor: clickOldModel.FinishYCor,
	SessionBasedAvegareClickCount: float64(clickOldModel.TouchCount),
	DailyAvegareClickCount:        float64(clickOldModel.TouchCount),
	LastTouchCountMinusSessionBasedAvegareClickCount: 0,
}


var TotalClickDay2 = (clickRespondModel.FirstClickSessionYearOfDay - screenClickOldModel.FirstClickSessionYearOfDay) + 365*(clickRespondModel.FirstClickSessionYear-screenClickOldModel.FirstClickSessionYear)
var TotalClickCount2 = screenClickOldModel.TotalClickCount + clickRespondModel.TotalClickCount
var TotalClickSessionCount2 = screenClickOldModel.TotalClickSessionCount + clickRespondModel.TotalClickSessionCount
var TotalClickHour2 = ((clickRespondModel.FirstClickSessionYearOfDay + 365*clickRespondModel.FirstClickSessionYear)*24 + clickRespondModel.FirstClickSessionHour) - ((screenClickOldModel.FirstClickSessionYearOfDay + 365*screenClickOldModel.FirstClickSessionYear)*24 + screenClickOldModel.FirstClickSessionHour)
var TotalClickMinute2 = (((clickRespondModel.FirstClickSessionYearOfDay + 365*clickRespondModel.FirstClickSessionYear)*24 + clickRespondModel.FirstClickSessionHour)*60 + clickRespondModel.FirstClickSessionMinute) - (((screenClickOldModel.FirstClickSessionYearOfDay + 365*screenClickOldModel.FirstClickSessionYear)*24 + screenClickOldModel.FirstClickSessionHour)*60 + screenClickOldModel.FirstClickSessionMinute)

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
	SessionBasedAvegareStartXCor:  (clickOldModel.StartXCor + clickOldModel2.StartXCor + clickRespondModel.FirstStartXCor)/float64(TotalClickSessionCount2),
	SessionBasedAvegareStartYCor:  (clickOldModel.StartYCor + clickOldModel2.StartYCor + clickRespondModel.FirstStartYCor)/float64(TotalClickSessionCount2),
	SessionBasedAvegareFinishXCor: (clickOldModel.FinishXCor + clickOldModel2.FinishXCor + clickRespondModel.FirstFinishXCor)/float64(TotalClickSessionCount2),
	SessionBasedAvegareFinishYCor: (clickOldModel.FinishYCor + clickOldModel2.FinishYCor + clickRespondModel.FirstFinishYCor)/float64(TotalClickSessionCount2),
	SessionBasedAvegareClickCount: float64(TotalClickCount)/float64(TotalClickSessionCount2),
	DailyAvegareClickCount:        float64(TotalClickCount)/float64(TotalClickDay2),
	LastTouchCountMinusSessionBasedAvegareClickCount: float64(clickRespondModel.FirstTouchCount) - float64(TotalClickCount)/float64(TotalClickSessionCount2),
}



func Test_UpdateScreenClick_Success(t *testing.T){
	var testClickDal = new(repository.MockScreenClickDal)
	var manager = concrete.ScreenClickManager{
		IScreenClickDal: testClickDal,
		IJsonParser:     &gojson.goJson{},
	}
	testClickDal.On("UpdateScreenClickById", screenClickOldModel.ClientId ,&updatedClickModel).Return(nil)
	uptModel, s, e := manager.UpdateScreenClick(&clickRespondModel, &screenClickOldModel)
	assert.Equal(t, &updatedClickModel, uptModel)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, e)
}

func Test_ConvertRawModelToResponseModel_Add_Succes(t *testing.T){
	var testClickDal = new(repository.MockScreenClickDal)
	var manager = concrete.ScreenClickManager{
		IScreenClickDal: testClickDal,
		IJsonParser:     &gojson.goJson{},
	}
	testClickDal.On("Add", &clickRespondModel).Return(nil)
	testClickDal.On("GetScreenClickById", screenClickOldModel.ClientId).Return(&screenClickOldModel, 
		errors.New("mongo: no documents in result"))
	byteData, _ := manager.IJsonParser.EncodeJson(clickModel)
	v,s,m := manager.ConvertRawModelToResponseModel(byteData)
	assert.Equal(t, &clickRespondModel, v)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
}


func Test_CalculateSeventhDayClickCount(t *testing.T){
	var count = concrete.CalculateSeventhDayClickCount(&clickRespondModel, &screenClickOldModel, TotalClickHour2)
	//var Expcount int64 = 150
	var Expcount int64 = 0
	fmt.Println(TotalClickHour2)
	assert.Equal(t, Expcount,count)
}

func Test_CalculateSixthDayClickCount(t *testing.T){
	var count = concrete.CalculateSixthDayClickCount(&clickRespondModel, &screenClickOldModel, TotalClickHour2)
	//var Expcount int64 = 150
	var Expcount int64 = 0
	fmt.Println(TotalClickHour2)
	assert.Equal(t, Expcount,count)
}

func Test_CalculateFifthDayClickCount(t *testing.T){
	var count = concrete.CalculateFifthDayClickCount(&clickRespondModel, &screenClickOldModel, TotalClickHour2)
	var Expcount int64 = 150
	//var Expcount int64 = 0
	fmt.Println(TotalClickHour2)
	assert.Equal(t, Expcount,count)
}

func Test_CalculateFourthDayClickCount(t *testing.T){
	var count = concrete.CalculateFourthDayClickCount(&clickRespondModel, &screenClickOldModel, TotalClickHour2)
	var Expcount int64 = 150
	//var Expcount int64 = 0
	fmt.Println(TotalClickHour2)
	assert.Equal(t, Expcount,count)
}

func Test_calculateThirdDayClickCount(t *testing.T){
	var count = concrete.CalculateThirdDayClickCount(&clickRespondModel, &screenClickOldModel, TotalClickHour2)
	var Expcount int64 = 150
	//var Expcount int64 = 0
	fmt.Println(TotalClickHour2)
	assert.Equal(t, Expcount,count)
}

func Test_CalculateSecondDayClickCount(t *testing.T){
	var count = concrete.CalculateSecondDayClickCount(&clickRespondModel, &screenClickOldModel, TotalClickHour2)
	//var Expcount int64 = 150
	var Expcount int64 = 0
	fmt.Println(TotalClickHour2)
	assert.Equal(t, Expcount,count)
}


func Test_calculateFirstDayClickCount(t *testing.T){
	var count = concrete.CalculateFirstDayClickCount(&clickRespondModel, &screenClickOldModel, TotalClickMinute2)
	var Expcount int64 =  666 + 150
	fmt.Println(TotalClickMinute2)
	assert.Equal(t, Expcount,count)
}


func Test_CalculateClickCount_FourthClick(t *testing.T){
	concrete.CalculateClickCount(&clickRespondModel, &screenClickOldModel)
	assert.Equal(t, Hour, screenClickOldModel.FourthClickSessionHour)
	assert.Equal(t, Minute, screenClickOldModel.FourthClickSessionMinute)
	assert.Equal(t, Count, screenClickOldModel.FourthTouchCount)
	assert.Equal(t, StartXCor, screenClickOldModel.FourthStartXCor)
	assert.Equal(t, StartYCor, screenClickOldModel.FourthStartYCor)
	assert.Equal(t, FinishXCor, screenClickOldModel.FourthFinishXCor)
	assert.Equal(t, FinishYCor, screenClickOldModel.FourthFinishYCor)
}

func Test_CalculateClickCount_ThirdClick(t *testing.T){
	concrete.CalculateClickCount(&clickRespondModel, &screenClickOldModel)
	assert.Equal(t, Hour, screenClickOldModel.ThirdClickSessionHour)
	assert.Equal(t, Minute, screenClickOldModel.ThirdClickSessionMinute)
	assert.Equal(t, Count, screenClickOldModel.ThirdTouchCount)
	assert.Equal(t, StartXCor, screenClickOldModel.ThirdStartXCor)
	assert.Equal(t, StartYCor, screenClickOldModel.ThirdStartYCor)
	assert.Equal(t, FinishXCor, screenClickOldModel.ThirdFinishXCor)
	assert.Equal(t, FinishYCor, screenClickOldModel.ThirdFinishYCor)
}

func Test_CalculateClickCount_SecondClick(t *testing.T){
	concrete.CalculateClickCount(&clickRespondModel, &screenClickOldModel)
	assert.Equal(t, Hour, screenClickOldModel.SecondClickSessionHour)
	assert.Equal(t, Minute, screenClickOldModel.SecondClickSessionMinute)
	assert.Equal(t, Count, screenClickOldModel.SecondTouchCount)
	assert.Equal(t, StartXCor, screenClickOldModel.SecondStartXCor)
	assert.Equal(t, StartYCor, screenClickOldModel.SecondStartYCor)
	assert.Equal(t, FinishXCor, screenClickOldModel.SecondFinishXCor)
	assert.Equal(t, FinishYCor, screenClickOldModel.SecondFinishYCor)
}



func Test_CalculateClickTwentyHour(t *testing.T){
	var count = concrete.CalculateClickTwentyHour(&clickRespondModel, &screenClickOldModel, TotalClickMinute2)
	//var ExpCount int64 = (screenClickOldModel.FirstHalfHourTouchCount)
	fmt.Println(TotalClickMinute2)
	//var ExpCount int64 = 666 + 150
	var ExpCount int64 = 666
	assert.Equal(t, []int64{ExpCount}, []int64{count})
}

func Test_CalculateClickTwoHour(t *testing.T){
	var count = concrete.CalculateClickTwoHour(&clickRespondModel, &screenClickOldModel, TotalClickMinute2)
	//var ExpCount int64 = (screenClickOldModel.FirstHalfHourTouchCount)
	fmt.Println(TotalClickMinute2)
	var ExpCount int64 = 666 + 150
	//var ExpCount int64 = 666
	assert.Equal(t, []int64{ExpCount}, []int64{count})
}


func Test_CalculateClickThreeHour(t *testing.T){
	var count = concrete.CalculateClickThreeHour(&clickRespondModel, &screenClickOldModel, TotalClickMinute2)
	//var ExpCount int64 = (screenClickOldModel.FirstHalfHourTouchCount)
	fmt.Println(TotalClickMinute2)
	var ExpCount int64 = 666 + 150
	//var ExpCount int64 = 666
	assert.Equal(t, []int64{ExpCount}, []int64{count})
}

func Test_CalculateClickSixHour(t *testing.T){
	var count = concrete.CalculateClickSixHour(&clickRespondModel, &screenClickOldModel, TotalClickMinute2)
	//var ExpCount int64 = (screenClickOldModel.FirstHalfHourTouchCount)
	fmt.Println(TotalClickMinute2)
	var ExpCount int64 = 666 + 150
	//var ExpCount int64 = 666
	assert.Equal(t, []int64{ExpCount}, []int64{count})
}

func Test_calculateClickHour(t *testing.T){
	var count = concrete.CalculateClickHour(&clickRespondModel, &screenClickOldModel, TotalClickMinute2)
	//var ExpCount int64 = (screenClickOldModel.FirstHalfHourTouchCount)
	fmt.Println(TotalClickMinute2)
	//var ExpCount int64 = 666 + 150
	var ExpCount int64 = 666
	assert.Equal(t, []int64{ExpCount}, []int64{count})
}

func Test_CalculateClickHalfHour(t *testing.T){
	var count = concrete.CalculateClickHalfHour(&clickRespondModel, &screenClickOldModel, TotalClickMinute2)
	//var ExpCount int64 = (screenClickOldModel.FirstHalfHourTouchCount)
	fmt.Println(TotalClickMinute2)
	var ExpCount int64 = 666 + 150
	assert.Equal(t, []int64{ExpCount}, []int64{count})
}

	
func Test_CalculateDailyAverageClickCount(t *testing.T){
	var count = concrete.CalculateDailyAverageClickCount(&screenClickOldModel)
	var ExpCount float64= float64(screenClickOldModel.TotalClickCount) / float64(TotalClickDay)
	assert.Equal(t, []float64{ExpCount}, []float64{count})
}
