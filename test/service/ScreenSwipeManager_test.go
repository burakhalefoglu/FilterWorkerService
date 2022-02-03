package test

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var swipeModel = model.ScreenSwipeModel{
	ProjectId:       "Test",
	ClientId:        "Test",
	CustomerId:      "Test",
	SwipeDirection:  1,
	SwipeStartXCor:  180,
	SwipeStartYCor:  90,
	SwipeFinishXCor: 360,
	SwipeFinishYCor: 180,
	CreationAt: time.Date(
		2021, 11, 11, 18, 34, 36, 651387237, time.UTC),
	LevelIndex: 20,
	LevelName:  "20",
}

var swipeRespondModel = model.ScreenSwipeRespondModel{}

var firstswipe = time.Date(
	2021, 11, 6, 18, 34, 58, 651387237, time.UTC)

var lastswipe = time.Date(
	2021, 11, 7, 19, 34, 58, 651387237, time.UTC)

var swipeModel2 = model.ScreenSwipeModel{
	ProjectId:       "Test",
	ClientId:        "Test",
	CustomerId:      "Test",
	SwipeDirection:  4,
	SwipeStartXCor:  15,
	SwipeStartYCor:  30,
	SwipeFinishXCor: 45,
	SwipeFinishYCor: 60,
	CreationAt:      firstswipe,
	LevelIndex:      30,
	LevelName:       "30",
}

var swipeModel3 = model.ScreenSwipeModel{
	ProjectId:       "Test",
	ClientId:        "Test",
	CustomerId:      "Test",
	SwipeDirection:  3,
	SwipeStartXCor:  60,
	SwipeStartYCor:  160,
	SwipeFinishXCor: 260,
	SwipeFinishYCor: 360,
	CreationAt:      lastswipe,
	LevelIndex:      80,
	LevelName:       "80",
}

var totalSwipeOldHour = int64(((lastswipe.YearDay()+365*lastswipe.Year())*24 + lastswipe.Hour()) - ((firstswipe.YearDay()+365*firstswipe.Year())*24 + firstswipe.Hour()))

var swipeOldModel = model.ScreenSwipeRespondModel{
	ProjectId:                      "Test",
	ClientId:                       "Test",
	CustomerId:                     "Test",
	LevelIndex:                     int64(swipeModel3.LevelIndex),
	TotalSwipeSessionCount:         2,
	TotalSwipeHour:                 totalSwipeOldHour,
	FirstSwipeYearOfDay:            int64(firstswipe.YearDay()),
	FirstSwipeYear:                 int64(firstswipe.Year()),
	FirstSwipeHour:                 int64(firstswipe.Hour()),
	FirstSwipeWeekDay:              int64(firstswipe.Weekday()),
	FirstSwipeMinute:               int64(firstswipe.Minute()),
	FistSwipeDirection:             int64(swipeModel2.SwipeDirection),
	FirstSwipeStartXCor:            swipeModel2.SwipeStartXCor,
	FirstSwipeStartYCor:            swipeModel2.SwipeStartYCor,
	FirstSwipeFinishXCor:           swipeModel2.SwipeFinishXCor,
	FirstSwipeFinishYCor:           swipeModel2.SwipeFinishYCor,
	SecondSwipeDirection:           0,
	SecondSwipeStartXCor:           0,
	SecondSwipeStartYCor:           0,
	SecondSwipeFinishXCor:          0,
	SecondSwipeFinishYCor:          0,
	ThirdSwipeDirection:            0,
	ThirdSwipeStartXCor:            0,
	ThirdSwipeStartYCor:            0,
	ThirdSwipeFinishXCor:           0,
	ThirdSwipeFinishYCor:           0,
	FourthSwipeDirection:           0,
	FourthSwipeStartXCor:           0,
	FourthSwipeStartYCor:           0,
	FourthSwipeFinishXCor:          0,
	FourthSwipeFinishYCor:          0,
	FifthSwipeDirection:            0,
	FifthSwipeStartXCor:            0,
	FifthSwipeStartYCor:            0,
	FifthSwipeFinishXCor:           0,
	FifthSwipeFinishYCor:           0,
	PenultimateSwipeDirection:      0,
	PenultimateSwipeStartXCor:      0,
	PenultimateSwipeStartYCor:      0,
	PenultimateSwipeFinishXCor:     0,
	PenultimateSwipeFinishYCor:     0,
	PenultimateSwipeYearOfDay:      0,
	PenultimateSwipeYear:           0,
	PenultimateSwipeHour:           0,
	PenultimateSwipeWeekDay:        0,
	PenultimateSwipeMinute:         0,
	LastSwipeDirection:             int64(swipeModel3.SwipeDirection),
	LastSwipeStartXCor:             swipeModel3.SwipeStartXCor,
	LastSwipeStartYCor:             swipeModel3.SwipeStartYCor,
	LastSwipeFinishXCor:            swipeModel3.SwipeFinishXCor,
	LastSwipeFinishYCor:            swipeModel3.SwipeFinishYCor,
	LastSwipeYearOfDay:             int64(lastswipe.YearDay()),
	LastSwipeYear:                  int64(lastswipe.Year()),
	LastSwipeHour:                  int64(lastswipe.Hour()),
	LastSwipeWeekDay:               int64(lastswipe.Weekday()),
	LastSwipeMinute:                int64(lastswipe.Minute()),
	FirstDayTotalSwipeUpCount:      0,
	FirstDayTotalSwipeDownCount:    1,
	FirstDayTotalSwipeRightCount:   0,
	FirstDayTotalSwipeLeftCount:    0,
	FirstDaySwipeTotalStartXCor:    swipeModel2.SwipeStartXCor,
	FirstDaySwipeTotalStartYCor:    swipeModel2.SwipeStartYCor,
	FirstDaySwipeTotalFinishXCor:   swipeModel2.SwipeFinishXCor,
	FirstDaySwipeTotalFinishYCor:   swipeModel2.SwipeFinishYCor,
	SecondDayTotalSwipeUpCount:     1,
	SecondDayTotalSwipeDownCount:   0,
	SecondDayTotalSwipeRightCount:  0,
	SecondDayTotalSwipeLeftCount:   0,
	SecondDaySwipeTotalStartXCor:   swipeModel3.SwipeStartXCor,
	SecondDaySwipeTotalStartYCor:   swipeModel3.SwipeStartYCor,
	SecondDaySwipeTotalFinishXCor:  swipeModel3.SwipeFinishXCor,
	SecondDaySwipeTotalFinishYCor:  swipeModel3.SwipeFinishYCor,
	ThirdDayTotalSwipeUpCount:      0,
	ThirdDayTotalSwipeDownCount:    0,
	ThirdDayTotalSwipeRightCount:   0,
	ThirdDayTotalSwipeLeftCount:    0,
	ThirdDaySwipeTotalStartXCor:    0,
	ThirdDaySwipeTotalStartYCor:    0,
	ThirdDaySwipeTotalFinishXCor:   0,
	ThirdDaySwipeTotalFinishYCor:   0,
	FourthDayTotalSwipeUpCount:     0,
	FourthDayTotalSwipeDownCount:   0,
	FourthDayTotalSwipeRightCount:  0,
	FourthDayTotalSwipeLeftCount:   0,
	FourthDaySwipeTotalStartXCor:   0,
	FourthDaySwipeTotalStartYCor:   0,
	FourthDaySwipeTotalFinishXCor:  0,
	FourthDaySwipeTotalFinishYCor:  0,
	FifthDayTotalSwipeUpCount:      0,
	FifthDayTotalSwipeDownCount:    0,
	FifthDayTotalSwipeRightCount:   0,
	FifthDayTotalSwipeLeftCount:    0,
	FifthDaySwipeTotalStartXCor:    0,
	FifthDaySwipeTotalStartYCor:    0,
	FifthDaySwipeTotalFinishXCor:   0,
	FifthDaySwipeTotalFinishYCor:   0,
	SixthDayTotalSwipeUpCount:      0,
	SixthDayTotalSwipeDownCount:    0,
	SixthDayTotalSwipeRightCount:   0,
	SixthDayTotalSwipeLeftCount:    0,
	SixthDaySwipeTotalStartXCor:    0,
	SixthDaySwipeTotalStartYCor:    0,
	SixthDaySwipeTotalFinishXCor:   0,
	SixthDaySwipeTotalFinishYCor:   0,
	SeventhDayTotalSwipeUpCount:    0,
	SeventhDayTotalSwipeDownCount:  0,
	SeventhDayTotalSwipeRightCount: 0,
	SeventhDayTotalSwipeLeftCount:  0,
	SeventhDaySwipeTotalStartXCor:  0,
	SeventhDaySwipeTotalStartYCor:  0,
	SeventhDaySwipeTotalFinishXCor: 0,
	SeventhDaySwipeTotalFinishYCor: 0,
	TotalSwipeUpCount:              1,
	TotalSwipeDownCount:            1,
	TotalSwipeRightCount:           0,
	TotalSwipeLeftCount:            0,
	TotalSwipeStartXCor:            swipeModel2.SwipeStartXCor,
	TotalSwipeStartYCor:            swipeModel2.SwipeStartYCor,
	TotalSwipeFinishXCor:           swipeModel2.SwipeFinishXCor,
	TotalSwipeFinishYCor:           swipeModel2.SwipeFinishYCor,
}

var totalSwipeOldHour2 = ((swipeRespondModel.FirstSwipeYearOfDay+365*swipeRespondModel.FirstSwipeYear)*24 + swipeRespondModel.FirstSwipeHour) - ((swipeOldModel.FirstSwipeYearOfDay+365*swipeOldModel.FirstSwipeYear)*24 + swipeOldModel.FirstSwipeHour)

var swipeUpdatedModel = model.ScreenSwipeRespondModel{
	ProjectId:                      "Test",
	ClientId:                       "Test",
	CustomerId:                     "Test",
	LevelIndex:                     swipeRespondModel.LevelIndex,
	TotalSwipeSessionCount:         swipeOldModel.TotalSwipeSessionCount + swipeRespondModel.TotalSwipeSessionCount,
	TotalSwipeHour:                 totalSwipeOldHour2,
	FirstSwipeYearOfDay:            int64(firstswipe.YearDay()),
	FirstSwipeYear:                 int64(firstswipe.Year()),
	FirstSwipeHour:                 int64(firstswipe.Hour()),
	FirstSwipeWeekDay:              int64(firstswipe.Weekday()),
	FirstSwipeMinute:               int64(firstswipe.Minute()),
	FistSwipeDirection:             int64(swipeModel2.SwipeDirection),
	FirstSwipeStartXCor:            swipeModel2.SwipeStartXCor,
	FirstSwipeStartYCor:            swipeModel2.SwipeStartYCor,
	FirstSwipeFinishXCor:           swipeModel2.SwipeFinishXCor,
	FirstSwipeFinishYCor:           swipeModel2.SwipeFinishYCor,
	SecondSwipeDirection:           0,
	SecondSwipeStartXCor:           0,
	SecondSwipeStartYCor:           0,
	SecondSwipeFinishXCor:          0,
	SecondSwipeFinishYCor:          0,
	ThirdSwipeDirection:            0,
	ThirdSwipeStartXCor:            0,
	ThirdSwipeStartYCor:            0,
	ThirdSwipeFinishXCor:           0,
	ThirdSwipeFinishYCor:           0,
	FourthSwipeDirection:           0,
	FourthSwipeStartXCor:           0,
	FourthSwipeStartYCor:           0,
	FourthSwipeFinishXCor:          0,
	FourthSwipeFinishYCor:          0,
	FifthSwipeDirection:            0,
	FifthSwipeStartXCor:            0,
	FifthSwipeStartYCor:            0,
	FifthSwipeFinishXCor:           0,
	FifthSwipeFinishYCor:           0,
	PenultimateSwipeDirection:      int64(swipeModel3.SwipeDirection),
	PenultimateSwipeStartXCor:      swipeModel3.SwipeStartXCor,
	PenultimateSwipeStartYCor:      swipeModel3.SwipeStartYCor,
	PenultimateSwipeFinishXCor:     swipeModel3.SwipeFinishXCor,
	PenultimateSwipeFinishYCor:     swipeModel3.SwipeFinishYCor,
	PenultimateSwipeYearOfDay:      int64(lastswipe.YearDay()),
	PenultimateSwipeYear:           int64(lastswipe.Year()),
	PenultimateSwipeHour:           int64(lastswipe.Hour()),
	PenultimateSwipeWeekDay:        int64(lastswipe.Weekday()),
	PenultimateSwipeMinute:         int64(lastswipe.Minute()),
	LastSwipeDirection:             swipeRespondModel.FistSwipeDirection,
	LastSwipeStartXCor:             swipeRespondModel.FirstSwipeStartXCor,
	LastSwipeStartYCor:             swipeRespondModel.FirstSwipeStartYCor,
	LastSwipeFinishXCor:            swipeRespondModel.FirstSwipeFinishXCor,
	LastSwipeFinishYCor:            swipeRespondModel.FirstSwipeFinishYCor,
	LastSwipeYearOfDay:             swipeRespondModel.FirstSwipeYearOfDay,
	LastSwipeYear:                  swipeRespondModel.FirstSwipeYear,
	LastSwipeHour:                  swipeRespondModel.FirstSwipeHour,
	LastSwipeWeekDay:               swipeRespondModel.FirstSwipeWeekDay,
	LastSwipeMinute:                swipeRespondModel.FirstSwipeMinute,
	FirstDayTotalSwipeUpCount:      0,
	FirstDayTotalSwipeDownCount:    1,
	FirstDayTotalSwipeRightCount:   0,
	FirstDayTotalSwipeLeftCount:    0,
	FirstDaySwipeTotalStartXCor:    swipeModel2.SwipeStartXCor,
	FirstDaySwipeTotalStartYCor:    swipeModel2.SwipeStartYCor,
	FirstDaySwipeTotalFinishXCor:   swipeModel2.SwipeFinishXCor,
	FirstDaySwipeTotalFinishYCor:   swipeModel2.SwipeFinishYCor,
	SecondDayTotalSwipeUpCount:     1,
	SecondDayTotalSwipeDownCount:   0,
	SecondDayTotalSwipeRightCount:  0,
	SecondDayTotalSwipeLeftCount:   0,
	SecondDaySwipeTotalStartXCor:   0,
	SecondDaySwipeTotalStartYCor:   0,
	SecondDaySwipeTotalFinishXCor:  0,
	SecondDaySwipeTotalFinishYCor:  0,
	ThirdDayTotalSwipeUpCount:      0,
	ThirdDayTotalSwipeDownCount:    0,
	ThirdDayTotalSwipeRightCount:   0,
	ThirdDayTotalSwipeLeftCount:    0,
	ThirdDaySwipeTotalStartXCor:    0,
	ThirdDaySwipeTotalStartYCor:    0,
	ThirdDaySwipeTotalFinishXCor:   0,
	ThirdDaySwipeTotalFinishYCor:   0,
	FourthDayTotalSwipeUpCount:     0,
	FourthDayTotalSwipeDownCount:   0,
	FourthDayTotalSwipeRightCount:  0,
	FourthDayTotalSwipeLeftCount:   0,
	FourthDaySwipeTotalStartXCor:   0,
	FourthDaySwipeTotalStartYCor:   0,
	FourthDaySwipeTotalFinishXCor:  0,
	FourthDaySwipeTotalFinishYCor:  0,
	FifthDayTotalSwipeUpCount:      0,
	FifthDayTotalSwipeDownCount:    0,
	FifthDayTotalSwipeRightCount:   0,
	FifthDayTotalSwipeLeftCount:    0,
	FifthDaySwipeTotalStartXCor:    0,
	FifthDaySwipeTotalStartYCor:    0,
	FifthDaySwipeTotalFinishXCor:   0,
	FifthDaySwipeTotalFinishYCor:   0,
	SixthDayTotalSwipeUpCount:      0,
	SixthDayTotalSwipeDownCount:    0,
	SixthDayTotalSwipeRightCount:   0,
	SixthDayTotalSwipeLeftCount:    0,
	SixthDaySwipeTotalStartXCor:    0,
	SixthDaySwipeTotalStartYCor:    0,
	SixthDaySwipeTotalFinishXCor:   0,
	SixthDaySwipeTotalFinishYCor:   0,
	SeventhDayTotalSwipeUpCount:    0,
	SeventhDayTotalSwipeDownCount:  0,
	SeventhDayTotalSwipeRightCount: 0,
	SeventhDayTotalSwipeLeftCount:  0,
	SeventhDaySwipeTotalStartXCor:  0,
	SeventhDaySwipeTotalStartYCor:  0,
	SeventhDaySwipeTotalFinishXCor: 0,
	SeventhDaySwipeTotalFinishYCor: 0,
	TotalSwipeUpCount:              swipeRespondModel.TotalSwipeUpCount + swipeOldModel.TotalSwipeUpCount,
	TotalSwipeDownCount:            swipeRespondModel.TotalSwipeDownCount + swipeOldModel.TotalSwipeDownCount,
	TotalSwipeRightCount:           swipeRespondModel.TotalSwipeRightCount + swipeOldModel.TotalSwipeRightCount,
	TotalSwipeLeftCount:            swipeRespondModel.TotalSwipeLeftCount + swipeOldModel.TotalSwipeLeftCount,
	TotalSwipeStartXCor:            swipeRespondModel.TotalSwipeStartXCor + swipeOldModel.TotalSwipeStartXCor,
	TotalSwipeStartYCor:            swipeRespondModel.TotalSwipeStartYCor + swipeOldModel.TotalSwipeStartYCor,
	TotalSwipeFinishXCor:           swipeRespondModel.TotalSwipeFinishXCor + swipeOldModel.TotalSwipeFinishXCor,
	TotalSwipeFinishYCor:           swipeRespondModel.TotalSwipeFinishYCor + swipeOldModel.TotalSwipeFinishYCor,
}

func Test_UpdateScreenSwipe_Success(t *testing.T) {

	var testSwipeDal = new(repository.MockScreenSwipeDal)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.ScreenSwipeDal = testSwipeDal
	var manager = concrete.ScreenSwipeManagerConstructor()
	var swipeModel_test = swipeModel
	var swipeOldModel_test = swipeOldModel
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.ProjectId                   =   "Test"
	swipeRespondModel_test.ClientId                    =   "Test"
	swipeRespondModel_test.CustomerId                  =   "Test"
	swipeRespondModel_test.LevelIndex                  =   int64(swipeModel.LevelIndex)
	swipeRespondModel_test.TotalSwipeSessionCount      =   1
	swipeRespondModel_test.TotalSwipeHour              =   0
	swipeRespondModel_test.FirstSwipeYearOfDay         =   int64(swipeModel.CreationAt.YearDay())
	swipeRespondModel_test.FirstSwipeYear              =   int64(swipeModel.CreationAt.Year())
	swipeRespondModel_test.FirstSwipeHour              =   int64(swipeModel.CreationAt.Hour())
	swipeRespondModel_test.FirstSwipeWeekDay           =   int64(swipeModel.CreationAt.Weekday())
	swipeRespondModel_test.FirstSwipeMinute            =   int64(swipeModel.CreationAt.Minute())
	swipeRespondModel_test.FistSwipeDirection          =   int64(swipeModel.SwipeDirection)
	swipeRespondModel_test.FirstSwipeStartXCor         =   swipeModel.SwipeStartXCor
	swipeRespondModel_test.FirstSwipeStartYCor         =   swipeModel.SwipeStartYCor
	swipeRespondModel_test.FirstSwipeFinishXCor        =   swipeModel.SwipeFinishXCor
	swipeRespondModel_test.FirstSwipeFinishYCor        =   swipeModel.SwipeFinishYCor
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor      =  swipeModel.SwipeStartXCor
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor      = swipeModel.SwipeStartYCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor     =  swipeModel.SwipeFinishXCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor     =  swipeModel.SwipeFinishYCor
	swipeRespondModel_test.TotalSwipeRightCount          =  1
	swipeRespondModel_test.TotalSwipeLeftCount          =  0
	swipeRespondModel_test.FirstDayTotalSwipeRightCount  = 1
	swipeRespondModel_test.TotalSwipeStartXCor          =  swipeModel.SwipeStartXCor
	swipeRespondModel_test.TotalSwipeStartYCor          =  swipeModel.SwipeStartYCor
	swipeRespondModel_test.TotalSwipeFinishXCor          =  swipeModel.SwipeFinishXCor
	swipeRespondModel_test.TotalSwipeFinishYCor          =  swipeModel.SwipeFinishYCor

	swipeModel_test.CreationAt = time.Date(
		2021, 11, 5, 16, 11, 36, 651387237, time.UTC)

	swipeRespondModel_test.FirstSwipeYearOfDay = int64(swipeModel_test.CreationAt.YearDay())
	swipeRespondModel_test.FirstSwipeYear = int64(swipeModel_test.CreationAt.Year())
	swipeRespondModel_test.FirstSwipeHour = int64(swipeModel_test.CreationAt.Hour())
	swipeRespondModel_test.FirstSwipeWeekDay = int64(swipeModel_test.CreationAt.Weekday())
	swipeRespondModel_test.FirstSwipeMinute = int64(swipeModel_test.CreationAt.Minute())
	swipeRespondModel_test.FistSwipeDirection = int64(swipeModel_test.SwipeDirection)
	swipeRespondModel_test.FirstSwipeStartXCor = swipeModel_test.SwipeStartXCor
	swipeRespondModel_test.FirstSwipeStartYCor = swipeModel_test.SwipeStartYCor
	swipeRespondModel_test.FirstSwipeFinishXCor = swipeModel_test.SwipeFinishXCor
	swipeRespondModel_test.FirstSwipeFinishYCor = swipeModel_test.SwipeFinishYCor
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = swipeModel_test.SwipeStartXCor
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = swipeModel_test.SwipeStartYCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = swipeModel_test.SwipeFinishXCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = swipeModel_test.SwipeFinishYCor
	swipeRespondModel_test.TotalSwipeStartXCor = swipeModel_test.SwipeStartXCor
	swipeRespondModel_test.TotalSwipeStartYCor = swipeModel_test.SwipeStartYCor
	swipeRespondModel_test.TotalSwipeFinishXCor = swipeModel_test.SwipeFinishXCor
	swipeRespondModel_test.TotalSwipeFinishYCor = swipeModel_test.SwipeFinishYCor
	var swipeUpdatedModel_test = swipeUpdatedModel

	// fmt.Println(totalSwipeOldHour2)
	// fmt.Println(swipeUpdatedModel.TotalSwipeSessionCount)

	testSwipeDal.On("UpdateScreenSwipeById", swipeOldModel_test.ClientId, &swipeOldModel_test).Return(nil)
	v, s, m := manager.UpdateScreenSwipe(&swipeRespondModel_test, &swipeOldModel_test)
	assert.Equal(t, &swipeUpdatedModel_test, v)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
}

func Test_ConvertRawModelToResponseModel_Add(t *testing.T) {
	var testSwipeDal = new(repository.MockScreenSwipeDal)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.ScreenSwipeDal = testSwipeDal
	var manager = concrete.ScreenSwipeManagerConstructor()
	var swipeModel_test = swipeModel
	var swipeOldModel_test = swipeOldModel
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.ProjectId                   =   "Test"
	swipeRespondModel_test.ClientId                    =   "Test"
	swipeRespondModel_test.CustomerId                  =   "Test"
	swipeRespondModel_test.LevelIndex                  =   int64(swipeModel.LevelIndex)
	swipeRespondModel_test.TotalSwipeSessionCount      =   1
	swipeRespondModel_test.TotalSwipeHour              =   0
	swipeRespondModel_test.FirstSwipeYearOfDay         =   int64(swipeModel.CreationAt.YearDay())
	swipeRespondModel_test.FirstSwipeYear              =   int64(swipeModel.CreationAt.Year())
	swipeRespondModel_test.FirstSwipeHour              =   int64(swipeModel.CreationAt.Hour())
	swipeRespondModel_test.FirstSwipeWeekDay           =   int64(swipeModel.CreationAt.Weekday())
	swipeRespondModel_test.FirstSwipeMinute            =   int64(swipeModel.CreationAt.Minute())
	swipeRespondModel_test.FistSwipeDirection          =   int64(swipeModel.SwipeDirection)
	swipeRespondModel_test.FirstSwipeStartXCor         =   swipeModel.SwipeStartXCor
	swipeRespondModel_test.FirstSwipeStartYCor         =   swipeModel.SwipeStartYCor
	swipeRespondModel_test.FirstSwipeFinishXCor        =   swipeModel.SwipeFinishXCor
	swipeRespondModel_test.FirstSwipeFinishYCor        =   swipeModel.SwipeFinishYCor
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor      =  swipeModel.SwipeStartXCor
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor      = swipeModel.SwipeStartYCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor     =  swipeModel.SwipeFinishXCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor     =  swipeModel.SwipeFinishYCor
	swipeRespondModel_test.TotalSwipeRightCount          =  1
	swipeRespondModel_test.TotalSwipeLeftCount          =  0
	swipeRespondModel_test.FirstDayTotalSwipeRightCount  = 1
	swipeRespondModel_test.TotalSwipeStartXCor          =  swipeModel.SwipeStartXCor
	swipeRespondModel_test.TotalSwipeStartYCor          =  swipeModel.SwipeStartYCor
	swipeRespondModel_test.TotalSwipeFinishXCor          =  swipeModel.SwipeFinishXCor
	swipeRespondModel_test.TotalSwipeFinishYCor          =  swipeModel.SwipeFinishYCor

	swipeModel_test.CreationAt = time.Date(
		2021, 11, 5, 16, 11, 36, 651387237, time.UTC)

	swipeRespondModel_test.FirstSwipeYearOfDay = int64(swipeModel_test.CreationAt.YearDay())
	swipeRespondModel_test.FirstSwipeYear = int64(swipeModel_test.CreationAt.Year())
	swipeRespondModel_test.FirstSwipeHour = int64(swipeModel_test.CreationAt.Hour())
	swipeRespondModel_test.FirstSwipeWeekDay = int64(swipeModel_test.CreationAt.Weekday())
	swipeRespondModel_test.FirstSwipeMinute = int64(swipeModel_test.CreationAt.Minute())
	swipeRespondModel_test.FistSwipeDirection = int64(swipeModel_test.SwipeDirection)
	swipeRespondModel_test.FirstSwipeStartXCor = swipeModel_test.SwipeStartXCor
	swipeRespondModel_test.FirstSwipeStartYCor = swipeModel_test.SwipeStartYCor
	swipeRespondModel_test.FirstSwipeFinishXCor = swipeModel_test.SwipeFinishXCor
	swipeRespondModel_test.FirstSwipeFinishYCor = swipeModel_test.SwipeFinishYCor
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = swipeModel_test.SwipeStartXCor
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = swipeModel_test.SwipeStartYCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = swipeModel_test.SwipeFinishXCor
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = swipeModel_test.SwipeFinishYCor
	swipeRespondModel_test.TotalSwipeStartXCor = swipeModel_test.SwipeStartXCor
	swipeRespondModel_test.TotalSwipeStartYCor = swipeModel_test.SwipeStartYCor
	swipeRespondModel_test.TotalSwipeFinishXCor = swipeModel_test.SwipeFinishXCor
	swipeRespondModel_test.TotalSwipeFinishYCor = swipeModel_test.SwipeFinishYCor

	testSwipeDal.On("GetScreenSwipeById", swipeModel_test.ClientId).Return(&swipeOldModel_test,
		errors.New("null data error"))
	testSwipeDal.On("Add", &swipeRespondModel_test).Return(nil)
	var swipeModel_test_byte, _ = json.EncodeJson(swipeModel_test)
	var v, s, m = manager.ConvertRawModelToResponseModel(swipeModel_test_byte)
	var value, success = v.(model.BuyingEventRespondModel)
	if success == true {
		assert.Equal(t, &swipeRespondModel_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
}

func Test_CalculateSwipeNumber_SecondTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 2
	swipeOldModel_test.SecondSwipeDirection = 0
	swipeOldModel_test.SecondSwipeStartXCor = 0
	swipeOldModel_test.SecondSwipeStartYCor = 0
	swipeOldModel_test.SecondSwipeFinishXCor = 0
	swipeOldModel_test.SecondSwipeFinishYCor = 0
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 1
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 1
	var ExpSwipeStartXCor float64 = 0.8
	var ExpSwipeStartYCor float64 = 0.4
	var ExpSwipeFinishXCor float64 = 0.3
	var ExpSwipeFinishYCor float64 = 0.7

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.SecondSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.SecondSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.SecondSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.SecondSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.SecondSwipeFinishYCor)
}

func Test_CalculateSwipeNumber_NonSecondTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 3
	swipeOldModel_test.SecondSwipeDirection = 4
	swipeOldModel_test.SecondSwipeStartXCor = 0.9
	swipeOldModel_test.SecondSwipeStartYCor = 0.2
	swipeOldModel_test.SecondSwipeFinishXCor = 0.7
	swipeOldModel_test.SecondSwipeFinishYCor = 0.6
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 1
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 4
	var ExpSwipeStartXCor float64 = 0.9
	var ExpSwipeStartYCor float64 = 0.2
	var ExpSwipeFinishXCor float64 = 0.7
	var ExpSwipeFinishYCor float64 = 0.6

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.SecondSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.SecondSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.SecondSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.SecondSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.SecondSwipeFinishYCor)
}

func Test_CalculateSwipeNumber_ThirdTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 3
	swipeOldModel_test.ThirdSwipeDirection = 0
	swipeOldModel_test.ThirdSwipeStartXCor = 0
	swipeOldModel_test.ThirdSwipeStartYCor = 0
	swipeOldModel_test.ThirdSwipeFinishXCor = 0
	swipeOldModel_test.ThirdSwipeFinishYCor = 0
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 4
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 4
	var ExpSwipeStartXCor float64 = 0.8
	var ExpSwipeStartYCor float64 = 0.4
	var ExpSwipeFinishXCor float64 = 0.3
	var ExpSwipeFinishYCor float64 = 0.7

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.ThirdSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.ThirdSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.ThirdSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.ThirdSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.ThirdSwipeFinishYCor)
}

func Test_CalculateSwipeNumber_NonThirdTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 2
	swipeOldModel_test.ThirdSwipeDirection = 3
	swipeOldModel_test.ThirdSwipeStartXCor = 0.9
	swipeOldModel_test.ThirdSwipeStartYCor = 0.1
	swipeOldModel_test.ThirdSwipeFinishXCor = 0.6
	swipeOldModel_test.ThirdSwipeFinishYCor = 0.4
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 4
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 3
	var ExpSwipeStartXCor float64 = 0.9
	var ExpSwipeStartYCor float64 = 0.1
	var ExpSwipeFinishXCor float64 = 0.6
	var ExpSwipeFinishYCor float64 = 0.4

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.ThirdSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.ThirdSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.ThirdSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.ThirdSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.ThirdSwipeFinishYCor)
}

func Test_CalculateSwipeNumber_FourthTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 4
	swipeOldModel_test.FourthSwipeDirection = 3
	swipeOldModel_test.FourthSwipeStartXCor = 0.9
	swipeOldModel_test.FourthSwipeStartYCor = 0.1
	swipeOldModel_test.FourthSwipeFinishXCor = 0.6
	swipeOldModel_test.FourthSwipeFinishYCor = 0.4
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 4
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 4
	var ExpSwipeStartXCor float64 = 0.8
	var ExpSwipeStartYCor float64 = 0.4
	var ExpSwipeFinishXCor float64 = 0.3
	var ExpSwipeFinishYCor float64 = 0.7

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.FourthSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.FourthSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.FourthSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.FourthSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.FourthSwipeFinishYCor)
}

func Test_CalculateSwipeNumber_NonFourthTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 5
	swipeOldModel_test.FourthSwipeDirection = 3
	swipeOldModel_test.FourthSwipeStartXCor = 0.9
	swipeOldModel_test.FourthSwipeStartYCor = 0.1
	swipeOldModel_test.FourthSwipeFinishXCor = 0.6
	swipeOldModel_test.FourthSwipeFinishYCor = 0.4
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 4
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 3
	var ExpSwipeStartXCor float64 = 0.9
	var ExpSwipeStartYCor float64 = 0.1
	var ExpSwipeFinishXCor float64 = 0.6
	var ExpSwipeFinishYCor float64 = 0.4

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.FourthSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.FourthSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.FourthSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.FourthSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.FourthSwipeFinishYCor)
}

func Test_CalculateSwipeNumber_FifthTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 5
	swipeOldModel_test.FifthSwipeDirection = 3
	swipeOldModel_test.FifthSwipeStartXCor = 0.9
	swipeOldModel_test.FifthSwipeStartYCor = 0.1
	swipeOldModel_test.FifthSwipeFinishXCor = 0.6
	swipeOldModel_test.FifthSwipeFinishYCor = 0.4
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 4
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 4
	var ExpSwipeStartXCor float64 = 0.8
	var ExpSwipeStartYCor float64 = 0.4
	var ExpSwipeFinishXCor float64 = 0.3
	var ExpSwipeFinishYCor float64 = 0.7

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.FifthSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.FifthSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.FifthSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.FifthSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.FifthSwipeFinishYCor)
}

func Test_CalculateSwipeNumber_NonFifthTotalSwipeSessionCount(t *testing.T) {
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.TotalSwipeSessionCount = 6
	swipeOldModel_test.FifthSwipeDirection = 3
	swipeOldModel_test.FifthSwipeStartXCor = 0.9
	swipeOldModel_test.FifthSwipeStartYCor = 0.1
	swipeOldModel_test.FifthSwipeFinishXCor = 0.6
	swipeOldModel_test.FifthSwipeFinishYCor = 0.4
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FistSwipeDirection = 4
	swipeRespondModel_test.FirstSwipeStartXCor = 0.8
	swipeRespondModel_test.FirstSwipeStartYCor = 0.4
	swipeRespondModel_test.FirstSwipeFinishXCor = 0.3
	swipeRespondModel_test.FirstSwipeFinishYCor = 0.7
	concrete.CalculateSwipeNumber(&swipeRespondModel_test, &swipeOldModel_test)
	var ExpSwipeDirection int64 = 3
	var ExpSwipeStartXCor float64 = 0.9
	var ExpSwipeStartYCor float64 = 0.1
	var ExpSwipeFinishXCor float64 = 0.6
	var ExpSwipeFinishYCor float64 = 0.4

	assert.Equal(t, ExpSwipeDirection, swipeOldModel_test.FifthSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel_test.FifthSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel_test.FifthSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel_test.FifthSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel_test.FifthSwipeFinishYCor)
}

func Test_DetermineSwipeDirection_Right(t *testing.T) {
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.TotalSwipeUpCount = 0
	swipeRespondModel_test.TotalSwipeDownCount = 0
	swipeRespondModel_test.TotalSwipeRightCount = 0
	swipeRespondModel_test.TotalSwipeLeftCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 0
	var swipeModel_test = swipeModel
	swipeModel_test.SwipeDirection = 1
	swipeDirection := int64(swipeModel_test.SwipeDirection)
	concrete.DetermineSwipeDirection(&swipeRespondModel_test, swipeDirection)
	var ExpU int64 = 0
	var ExpD int64 = 0
	var ExpR int64 = 1
	var ExpL int64 = 0
	var ExpTotalU int64 = 0
	var ExpTotalD int64 = 0
	var ExpTotalR int64 = 1
	var ExpTotalL int64 = 0
	assert.Equal(t, ExpU, swipeRespondModel_test.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeRespondModel_test.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeRespondModel_test.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeRespondModel_test.FirstDayTotalSwipeLeftCount)
	assert.Equal(t, ExpTotalU, swipeRespondModel_test.TotalSwipeUpCount)
	assert.Equal(t, ExpTotalD, swipeRespondModel_test.TotalSwipeDownCount)
	assert.Equal(t, ExpTotalR, swipeRespondModel_test.TotalSwipeRightCount)
	assert.Equal(t, ExpTotalL, swipeRespondModel_test.TotalSwipeLeftCount)
}

func Test_DetermineSwipeDirection_Left(t *testing.T) {
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.TotalSwipeUpCount = 0
	swipeRespondModel_test.TotalSwipeDownCount = 0
	swipeRespondModel_test.TotalSwipeRightCount = 0
	swipeRespondModel_test.TotalSwipeLeftCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 0
	var swipeModel_test = swipeModel
	swipeModel_test.SwipeDirection = 2
	swipeDirection := int64(swipeModel_test.SwipeDirection)
	concrete.DetermineSwipeDirection(&swipeRespondModel_test, swipeDirection)
	var ExpU int64 = 0
	var ExpD int64 = 0
	var ExpR int64 = 0
	var ExpL int64 = 1
	var ExpTotalU int64 = 0
	var ExpTotalD int64 = 0
	var ExpTotalR int64 = 0
	var ExpTotalL int64 = 1
	assert.Equal(t, ExpU, swipeRespondModel_test.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeRespondModel_test.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeRespondModel_test.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeRespondModel_test.FirstDayTotalSwipeLeftCount)
	assert.Equal(t, ExpTotalU, swipeRespondModel_test.TotalSwipeUpCount)
	assert.Equal(t, ExpTotalD, swipeRespondModel_test.TotalSwipeDownCount)
	assert.Equal(t, ExpTotalR, swipeRespondModel_test.TotalSwipeRightCount)
	assert.Equal(t, ExpTotalL, swipeRespondModel_test.TotalSwipeLeftCount)
}

func Test_DetermineSwipeDirection_Up(t *testing.T) {
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.TotalSwipeUpCount = 0
	swipeRespondModel_test.TotalSwipeDownCount = 0
	swipeRespondModel_test.TotalSwipeRightCount = 0
	swipeRespondModel_test.TotalSwipeLeftCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 0
	var swipeModel_test = swipeModel
	swipeModel_test.SwipeDirection = 3
	swipeDirection := int64(swipeModel_test.SwipeDirection)
	concrete.DetermineSwipeDirection(&swipeRespondModel_test, swipeDirection)
	var ExpU int64 = 1
	var ExpD int64 = 0
	var ExpR int64 = 0
	var ExpL int64 = 0
	var ExpTotalU int64 = 1
	var ExpTotalD int64 = 0
	var ExpTotalR int64 = 0
	var ExpTotalL int64 = 0
	assert.Equal(t, ExpU, swipeRespondModel_test.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeRespondModel_test.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeRespondModel_test.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeRespondModel_test.FirstDayTotalSwipeLeftCount)
	assert.Equal(t, ExpTotalU, swipeRespondModel_test.TotalSwipeUpCount)
	assert.Equal(t, ExpTotalD, swipeRespondModel_test.TotalSwipeDownCount)
	assert.Equal(t, ExpTotalR, swipeRespondModel_test.TotalSwipeRightCount)
	assert.Equal(t, ExpTotalL, swipeRespondModel_test.TotalSwipeLeftCount)
}

func Test_DetermineSwipeDirection_Down(t *testing.T) {
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.TotalSwipeUpCount = 0
	swipeRespondModel_test.TotalSwipeDownCount = 0
	swipeRespondModel_test.TotalSwipeRightCount = 0
	swipeRespondModel_test.TotalSwipeLeftCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 0
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 0
	var swipeModel_test = swipeModel
	swipeModel_test.SwipeDirection = 4
	swipeDirection := int64(swipeModel_test.SwipeDirection)
	concrete.DetermineSwipeDirection(&swipeRespondModel_test, swipeDirection)
	var ExpU int64 = 0
	var ExpD int64 = 1
	var ExpR int64 = 0
	var ExpL int64 = 0
	var ExpTotalU int64 = 0
	var ExpTotalD int64 = 1
	var ExpTotalR int64 = 0
	var ExpTotalL int64 = 0
	assert.Equal(t, ExpU, swipeRespondModel_test.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeRespondModel_test.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeRespondModel_test.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeRespondModel_test.FirstDayTotalSwipeLeftCount)
	assert.Equal(t, ExpTotalU, swipeRespondModel_test.TotalSwipeUpCount)
	assert.Equal(t, ExpTotalD, swipeRespondModel_test.TotalSwipeDownCount)
	assert.Equal(t, ExpTotalR, swipeRespondModel_test.TotalSwipeRightCount)
	assert.Equal(t, ExpTotalL, swipeRespondModel_test.TotalSwipeLeftCount)
}

func Test_CalculateSwipeSeventhDay_In120To144Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 354
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 6
	swipeOldModel.SeventhDayTotalSwipeUpCount = 243
	swipeOldModel.SeventhDayTotalSwipeDownCount = 132
	swipeOldModel.SeventhDayTotalSwipeRightCount = 148
	swipeOldModel.SeventhDayTotalSwipeLeftCount = 987
	swipeOldModel.SeventhDaySwipeTotalStartXCor = 0.4
	swipeOldModel.SeventhDaySwipeTotalStartYCor = 0.9
	swipeOldModel.SeventhDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.SeventhDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeSeventhDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243 + 200
	var ExpD int64 = 132 + 300
	var ExpR int64 = 148 + 140
	var ExpL int64 = 987 + 150
	var ExpSX float64 = 0.4 + 0.9
	var ExpSY float64 = 0.9 + 0.1
	var ExpFX float64 = 0.3 + 0.2
	var ExpFY float64 = 0.7 + 0.3
	assert.Equal(t, ExpU, swipeOldModel.SeventhDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.SeventhDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.SeventhDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.SeventhDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.SeventhDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.SeventhDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.SeventhDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.SeventhDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeSeventhDay_Out120To144Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 354
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 7
	swipeOldModel.SeventhDayTotalSwipeUpCount = 243
	swipeOldModel.SeventhDayTotalSwipeDownCount = 132
	swipeOldModel.SeventhDayTotalSwipeRightCount = 148
	swipeOldModel.SeventhDayTotalSwipeLeftCount = 987
	swipeOldModel.SeventhDaySwipeTotalStartXCor = 0.4
	swipeOldModel.SeventhDaySwipeTotalStartYCor = 0.9
	swipeOldModel.SeventhDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.SeventhDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeSeventhDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243
	var ExpD int64 = 132
	var ExpR int64 = 148
	var ExpL int64 = 987
	var ExpSX float64 = 0.4
	var ExpSY float64 = 0.9
	var ExpFX float64 = 0.3
	var ExpFY float64 = 0.7
	assert.Equal(t, ExpU, swipeOldModel.SeventhDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.SeventhDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.SeventhDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.SeventhDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.SeventhDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.SeventhDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.SeventhDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.SeventhDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeSixthDay_In120To144Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 353
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 6
	swipeOldModel.SixthDayTotalSwipeUpCount = 243
	swipeOldModel.SixthDayTotalSwipeDownCount = 132
	swipeOldModel.SixthDayTotalSwipeRightCount = 148
	swipeOldModel.SixthDayTotalSwipeLeftCount = 987
	swipeOldModel.SixthDaySwipeTotalStartXCor = 0.4
	swipeOldModel.SixthDaySwipeTotalStartYCor = 0.9
	swipeOldModel.SixthDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.SixthDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeSixthDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243 + 200
	var ExpD int64 = 132 + 300
	var ExpR int64 = 148 + 140
	var ExpL int64 = 987 + 150
	var ExpSX float64 = 0.4 + 0.9
	var ExpSY float64 = 0.9 + 0.1
	var ExpFX float64 = 0.3 + 0.2
	var ExpFY float64 = 0.7 + 0.3
	assert.Equal(t, ExpU, swipeOldModel.SixthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.SixthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.SixthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.SixthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.SixthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.SixthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.SixthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.SixthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeSixthDay_Out120To144Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 353
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 7
	swipeOldModel.SixthDayTotalSwipeUpCount = 243
	swipeOldModel.SixthDayTotalSwipeDownCount = 132
	swipeOldModel.SixthDayTotalSwipeRightCount = 148
	swipeOldModel.SixthDayTotalSwipeLeftCount = 987
	swipeOldModel.SixthDaySwipeTotalStartXCor = 0.4
	swipeOldModel.SixthDaySwipeTotalStartYCor = 0.9
	swipeOldModel.SixthDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.SixthDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeSixthDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243
	var ExpD int64 = 132
	var ExpR int64 = 148
	var ExpL int64 = 987
	var ExpSX float64 = 0.4
	var ExpSY float64 = 0.9
	var ExpFX float64 = 0.3
	var ExpFY float64 = 0.7
	assert.Equal(t, ExpU, swipeOldModel.SixthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.SixthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.SixthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.SixthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.SixthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.SixthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.SixthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.SixthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeFifthDay_In96To120Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 352
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 6
	swipeOldModel.FifthDayTotalSwipeUpCount = 243
	swipeOldModel.FifthDayTotalSwipeDownCount = 132
	swipeOldModel.FifthDayTotalSwipeRightCount = 148
	swipeOldModel.FifthDayTotalSwipeLeftCount = 987
	swipeOldModel.FifthDaySwipeTotalStartXCor = 0.4
	swipeOldModel.FifthDaySwipeTotalStartYCor = 0.9
	swipeOldModel.FifthDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.FifthDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeFifthDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243 + 200
	var ExpD int64 = 132 + 300
	var ExpR int64 = 148 + 140
	var ExpL int64 = 987 + 150
	var ExpSX float64 = 0.4 + 0.9
	var ExpSY float64 = 0.9 + 0.1
	var ExpFX float64 = 0.3 + 0.2
	var ExpFY float64 = 0.7 + 0.3
	assert.Equal(t, ExpU, swipeOldModel.FifthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FifthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FifthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FifthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FifthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FifthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FifthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FifthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeFifthDay_Out96To120Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 352
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 7
	swipeOldModel.FifthDayTotalSwipeUpCount = 243
	swipeOldModel.FifthDayTotalSwipeDownCount = 132
	swipeOldModel.FifthDayTotalSwipeRightCount = 148
	swipeOldModel.FifthDayTotalSwipeLeftCount = 987
	swipeOldModel.FifthDaySwipeTotalStartXCor = 0.4
	swipeOldModel.FifthDaySwipeTotalStartYCor = 0.9
	swipeOldModel.FifthDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.FifthDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeFifthDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243
	var ExpD int64 = 132
	var ExpR int64 = 148
	var ExpL int64 = 987
	var ExpSX float64 = 0.4
	var ExpSY float64 = 0.9
	var ExpFX float64 = 0.3
	var ExpFY float64 = 0.7
	assert.Equal(t, ExpU, swipeOldModel.FifthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FifthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FifthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FifthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FifthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FifthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FifthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FifthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeFourthDay_In72To96Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 351
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 6
	swipeOldModel.FourthDayTotalSwipeUpCount = 243
	swipeOldModel.FourthDayTotalSwipeDownCount = 132
	swipeOldModel.FourthDayTotalSwipeRightCount = 148
	swipeOldModel.FourthDayTotalSwipeLeftCount = 987
	swipeOldModel.FourthDaySwipeTotalStartXCor = 0.4
	swipeOldModel.FourthDaySwipeTotalStartYCor = 0.9
	swipeOldModel.FourthDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.FourthDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeFourthDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243 + 200
	var ExpD int64 = 132 + 300
	var ExpR int64 = 148 + 140
	var ExpL int64 = 987 + 150
	var ExpSX float64 = 0.4 + 0.9
	var ExpSY float64 = 0.9 + 0.1
	var ExpFX float64 = 0.3 + 0.2
	var ExpFY float64 = 0.7 + 0.3
	assert.Equal(t, ExpU, swipeOldModel.FourthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FourthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FourthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FourthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FourthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FourthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FourthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FourthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeFourthDay_Out72To96Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 351
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 7
	swipeOldModel.FourthDayTotalSwipeUpCount = 243
	swipeOldModel.FourthDayTotalSwipeDownCount = 132
	swipeOldModel.FourthDayTotalSwipeRightCount = 148
	swipeOldModel.FourthDayTotalSwipeLeftCount = 987
	swipeOldModel.FourthDaySwipeTotalStartXCor = 0.4
	swipeOldModel.FourthDaySwipeTotalStartYCor = 0.9
	swipeOldModel.FourthDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.FourthDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeFourthDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243
	var ExpD int64 = 132
	var ExpR int64 = 148
	var ExpL int64 = 987
	var ExpSX float64 = 0.4
	var ExpSY float64 = 0.9
	var ExpFX float64 = 0.3
	var ExpFY float64 = 0.7
	assert.Equal(t, ExpU, swipeOldModel.FourthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FourthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FourthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FourthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FourthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FourthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FourthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FourthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeThirdDay_In48To72Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 350
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 6
	swipeOldModel.ThirdDayTotalSwipeUpCount = 243
	swipeOldModel.ThirdDayTotalSwipeDownCount = 132
	swipeOldModel.ThirdDayTotalSwipeRightCount = 148
	swipeOldModel.ThirdDayTotalSwipeLeftCount = 987
	swipeOldModel.ThirdDaySwipeTotalStartXCor = 0.4
	swipeOldModel.ThirdDaySwipeTotalStartYCor = 0.9
	swipeOldModel.ThirdDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.ThirdDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeThirdDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243 + 200
	var ExpD int64 = 132 + 300
	var ExpR int64 = 148 + 140
	var ExpL int64 = 987 + 150
	var ExpSX float64 = 0.4 + 0.9
	var ExpSY float64 = 0.9 + 0.1
	var ExpFX float64 = 0.3 + 0.2
	var ExpFY float64 = 0.7 + 0.3
	assert.Equal(t, ExpU, swipeOldModel.ThirdDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.ThirdDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.ThirdDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.ThirdDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.ThirdDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.ThirdDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.ThirdDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.ThirdDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeThirdDay_Out48To72Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 350
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 7
	swipeOldModel.ThirdDayTotalSwipeUpCount = 243
	swipeOldModel.ThirdDayTotalSwipeDownCount = 132
	swipeOldModel.ThirdDayTotalSwipeRightCount = 148
	swipeOldModel.ThirdDayTotalSwipeLeftCount = 987
	swipeOldModel.ThirdDaySwipeTotalStartXCor = 0.4
	swipeOldModel.ThirdDaySwipeTotalStartYCor = 0.9
	swipeOldModel.ThirdDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.ThirdDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeThirdDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243
	var ExpD int64 = 132
	var ExpR int64 = 148
	var ExpL int64 = 987
	var ExpSX float64 = 0.4
	var ExpSY float64 = 0.9
	var ExpFX float64 = 0.3
	var ExpFY float64 = 0.7
	assert.Equal(t, ExpU, swipeOldModel.ThirdDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.ThirdDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.ThirdDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.ThirdDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.ThirdDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.ThirdDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.ThirdDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.ThirdDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeSecondDay_In48To72Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 348
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 6
	swipeOldModel.SecondDayTotalSwipeUpCount = 243
	swipeOldModel.SecondDayTotalSwipeDownCount = 132
	swipeOldModel.SecondDayTotalSwipeRightCount = 148
	swipeOldModel.SecondDayTotalSwipeLeftCount = 987
	swipeOldModel.SecondDaySwipeTotalStartXCor = 0.4
	swipeOldModel.SecondDaySwipeTotalStartYCor = 0.9
	swipeOldModel.SecondDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.SecondDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeSecondDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243 + 200
	var ExpD int64 = 132 + 300
	var ExpR int64 = 148 + 140
	var ExpL int64 = 987 + 150
	var ExpSX float64 = 0.4 + 0.9
	var ExpSY float64 = 0.9 + 0.1
	var ExpFX float64 = 0.3 + 0.2
	var ExpFY float64 = 0.7 + 0.3

	assert.Equal(t, ExpU, swipeOldModel.SecondDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.SecondDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.SecondDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.SecondDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.SecondDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.SecondDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.SecondDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.SecondDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeSecondDay_Out48To72Hours(t *testing.T) {

	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 349
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 7
	swipeOldModel.SecondDayTotalSwipeUpCount = 243
	swipeOldModel.SecondDayTotalSwipeDownCount = 132
	swipeOldModel.SecondDayTotalSwipeRightCount = 148
	swipeOldModel.SecondDayTotalSwipeLeftCount = 987
	swipeOldModel.SecondDaySwipeTotalStartXCor = 0.4
	swipeOldModel.SecondDaySwipeTotalStartYCor = 0.9
	swipeOldModel.SecondDaySwipeTotalFinishXCor = 0.3
	swipeOldModel.SecondDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeSecondDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243
	var ExpD int64 = 132
	var ExpR int64 = 148
	var ExpL int64 = 987
	var ExpSX float64 = 0.4
	var ExpSY float64 = 0.9
	var ExpFX float64 = 0.3
	var ExpFY float64 = 0.7

	assert.Equal(t, ExpU, swipeOldModel.SecondDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.SecondDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.SecondDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.SecondDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.SecondDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.SecondDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.SecondDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.SecondDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeFirstDay_In24Hours(t *testing.T) {
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 348
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 6
	swipeOldModel_test.FirstDayTotalSwipeUpCount = 243
	swipeOldModel_test.FirstDayTotalSwipeDownCount = 132
	swipeOldModel_test.FirstDayTotalSwipeRightCount = 148
	swipeOldModel_test.FirstDayTotalSwipeLeftCount = 987
	swipeOldModel_test.FirstDaySwipeTotalStartXCor = 0.4
	swipeOldModel_test.FirstDaySwipeTotalStartYCor = 0.9
	swipeOldModel_test.FirstDaySwipeTotalFinishXCor = 0.3
	swipeOldModel_test.FirstDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeFirstDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 200 + 243
	var ExpD int64 = 300 + 132
	var ExpR int64 = 140 + 148
	var ExpL int64 = 150 + 987
	var ExpSX float64 = 0.9 + 0.4
	var ExpSY float64 = 0.1 + 0.9
	var ExpFX float64 = 0.2 + 0.3
	var ExpFY float64 = 0.3 + 0.7
	assert.Equal(t, ExpU, swipeOldModel.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FirstDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FirstDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FirstDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FirstDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FirstDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeFirstDay_Out24Hours(t *testing.T) {
	var swipeRespondModel_test = swipeRespondModel
	swipeRespondModel_test.FirstSwipeYearOfDay = 347
	swipeRespondModel_test.FirstSwipeYear = 2021
	swipeRespondModel_test.FirstSwipeHour = 6
	swipeRespondModel_test.FirstDayTotalSwipeUpCount = 200
	swipeRespondModel_test.FirstDayTotalSwipeDownCount = 300
	swipeRespondModel_test.FirstDayTotalSwipeRightCount = 140
	swipeRespondModel_test.FirstDayTotalSwipeLeftCount = 150
	swipeRespondModel_test.FirstDaySwipeTotalStartXCor = 0.9
	swipeRespondModel_test.FirstDaySwipeTotalStartYCor = 0.1
	swipeRespondModel_test.FirstDaySwipeTotalFinishXCor = 0.2
	swipeRespondModel_test.FirstDaySwipeTotalFinishYCor = 0.3
	var swipeOldModel_test = swipeOldModel
	swipeOldModel_test.FirstSwipeYearOfDay = 348
	swipeOldModel_test.FirstSwipeYear = 2021
	swipeOldModel_test.FirstSwipeHour = 7
	swipeOldModel_test.FirstDayTotalSwipeUpCount = 243
	swipeOldModel_test.FirstDayTotalSwipeDownCount = 132
	swipeOldModel_test.FirstDayTotalSwipeRightCount = 148
	swipeOldModel_test.FirstDayTotalSwipeLeftCount = 987
	swipeOldModel_test.FirstDaySwipeTotalStartXCor = 0.4
	swipeOldModel_test.FirstDaySwipeTotalStartYCor = 0.9
	swipeOldModel_test.FirstDaySwipeTotalFinishXCor = 0.3
	swipeOldModel_test.FirstDaySwipeTotalFinishYCor = 0.7
	var totalSwipeOldHour_test = ((swipeRespondModel_test.FirstSwipeYearOfDay+365*swipeRespondModel_test.FirstSwipeYear)*24 + swipeRespondModel_test.FirstSwipeHour) - ((swipeOldModel_test.FirstSwipeYearOfDay+365*swipeOldModel_test.FirstSwipeYear)*24 + swipeOldModel_test.FirstSwipeHour)
	concrete.CalculateSwipeFirstDay(&swipeRespondModel_test, &swipeOldModel_test, totalSwipeOldHour_test)
	var ExpU int64 = 243
	var ExpD int64 = 132
	var ExpR int64 = 148
	var ExpL int64 = 987
	var ExpSX float64 = 0.4
	var ExpSY float64 = 0.9
	var ExpFX float64 = 0.3
	var ExpFY float64 = 0.7
	assert.Equal(t, ExpU, swipeOldModel.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FirstDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FirstDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FirstDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FirstDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FirstDaySwipeTotalFinishYCor)
}
