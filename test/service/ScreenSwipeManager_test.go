package test

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"errors"
	"fmt"
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

var swipeRespondModel = model.ScreenSwipeRespondModel{
	ProjectId:                      "Test",
	ClientId:                       "Test",
	CustomerId:                     "Test",
	LevelIndex:                     int64(swipeModel.LevelIndex),
	TotalSwipeSessionCount:         1,
	TotalSwipeHour:                 0,
	FirstSwipeYearOfDay:            int64(swipeModel.CreationAt.YearDay()),
	FirstSwipeYear:                 int64(swipeModel.CreationAt.Year()),
	FirstSwipeHour:                 int64(swipeModel.CreationAt.Hour()),
	FirstSwipeWeekDay:              int64(swipeModel.CreationAt.Weekday()),
	FirstSwipeMinute:               int64(swipeModel.CreationAt.Minute()),
	FistSwipeDirection:             int64(swipeModel.SwipeDirection),
	FirstSwipeStartXCor:            swipeModel.SwipeStartXCor,
	FirstSwipeStartYCor:            swipeModel.SwipeStartYCor,
	FirstSwipeFinishXCor:           swipeModel.SwipeFinishXCor,
	FirstSwipeFinishYCor:           swipeModel.SwipeFinishYCor,
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
	LastSwipeDirection:             0,
	LastSwipeStartXCor:             0,
	LastSwipeStartYCor:             0,
	LastSwipeFinishXCor:            0,
	LastSwipeFinishYCor:            0,
	LastSwipeYearOfDay:             0,
	LastSwipeYear:                  0,
	LastSwipeHour:                  0,
	LastSwipeWeekDay:               0,
	LastSwipeMinute:                0,
	FirstDayTotalSwipeUpCount:      0,
	FirstDayTotalSwipeDownCount:    0,
	FirstDayTotalSwipeRightCount:   1,
	FirstDayTotalSwipeLeftCount:    0,
	FirstDaySwipeTotalStartXCor:    swipeModel.SwipeStartXCor,
	FirstDaySwipeTotalStartYCor:    swipeModel.SwipeStartYCor,
	FirstDaySwipeTotalFinishXCor:   swipeModel.SwipeFinishXCor,
	FirstDaySwipeTotalFinishYCor:   swipeModel.SwipeFinishYCor,
	SecondDayTotalSwipeUpCount:     0,
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
	TotalSwipeUpCount:              0,
	TotalSwipeDownCount:            0,
	TotalSwipeRightCount:           1,
	TotalSwipeLeftCount:            0,
	TotalSwipeStartXCor:            swipeModel.SwipeStartXCor,
	TotalSwipeStartYCor:            swipeModel.SwipeStartYCor,
	TotalSwipeFinishXCor:           swipeModel.SwipeFinishXCor,
	TotalSwipeFinishYCor:           swipeModel.SwipeFinishYCor,
}

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
	ThirdSwipeDirection:            swipeRespondModel.FistSwipeDirection,
	ThirdSwipeStartXCor:            swipeRespondModel.FirstSwipeStartXCor,
	ThirdSwipeStartYCor:            swipeRespondModel.FirstSwipeStartYCor,
	ThirdSwipeFinishXCor:           swipeRespondModel.FirstSwipeFinishXCor,
	ThirdSwipeFinishYCor:           swipeRespondModel.FirstSwipeFinishYCor,
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
	FifthDayTotalSwipeUpCount:      swipeRespondModel.TotalSwipeUpCount,
	FifthDayTotalSwipeDownCount:    swipeRespondModel.TotalSwipeDownCount,
	FifthDayTotalSwipeRightCount:   swipeRespondModel.TotalSwipeRightCount,
	FifthDayTotalSwipeLeftCount:    swipeRespondModel.TotalSwipeLeftCount,
	FifthDaySwipeTotalStartXCor:    swipeRespondModel.TotalSwipeStartXCor,
	FifthDaySwipeTotalStartYCor:    swipeRespondModel.TotalSwipeStartYCor,
	FifthDaySwipeTotalFinishXCor:   swipeRespondModel.TotalSwipeFinishXCor,
	FifthDaySwipeTotalFinishYCor:   swipeRespondModel.TotalSwipeFinishYCor,
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
	fmt.Println(totalSwipeOldHour2)
	fmt.Println(swipeUpdatedModel.TotalSwipeSessionCount)
	var testSwipeDal = new(repository.MockScreenSwipeDal)
	var manager = concrete.ScreenSwipeManager{
		IScreenSwipeDal: testSwipeDal,
		IJsonParser:     &gojson.GoJson{},
	}
	testSwipeDal.On("UpdateScreenSwipeById", swipeOldModel.ClientId, &swipeOldModel).Return(nil)
	v, s, m := manager.UpdateScreenSwipe(&swipeRespondModel, &swipeOldModel)
	assert.Equal(t, &swipeUpdatedModel, v)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
}

func Test_ConvertRawModelToResponseModel_Add(t *testing.T) {
	var testSwipeDal = new(repository.MockScreenSwipeDal)
	var manager = concrete.ScreenSwipeManager{
		IScreenSwipeDal: testSwipeDal,
		IJsonParser:     &gojson.GoJson{},
	}
	testSwipeDal.On("GetScreenSwipeById", swipeModel.ClientId).Return(&swipeOldModel,
		errors.New("mongo: no documents in result"))
	testSwipeDal.On("Add", &swipeRespondModel).Return(nil)
	byte_data, _ := manager.IJsonParser.EncodeJson(swipeModel)
	v, s, m := manager.ConvertRawModelToResponseModel(byte_data)
	assert.Equal(t, &swipeRespondModel, v)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
}

func Test_CalculateSwipeNumber(t *testing.T) {
	concrete.CalculateSwipeNumber(&swipeRespondModel, &swipeOldModel)
	var ExpSwipeDirection int64 = 1
	var ExpSwipeStartXCor float64 = 180
	var ExpSwipeStartYCor float64 = 90
	var ExpSwipeFinishXCor float64 = 360
	var ExpSwipeFinishYCor float64 = 180

	assert.Equal(t, ExpSwipeDirection, swipeOldModel.FifthSwipeDirection)
	assert.Equal(t, ExpSwipeStartXCor, swipeOldModel.FifthSwipeStartXCor)
	assert.Equal(t, ExpSwipeStartYCor, swipeOldModel.FifthSwipeStartYCor)
	assert.Equal(t, ExpSwipeFinishXCor, swipeOldModel.FifthSwipeFinishXCor)
	assert.Equal(t, ExpSwipeFinishYCor, swipeOldModel.FifthSwipeFinishYCor)

	// assert.Equal(t, ExpSwipeDirection ,  swipeOldModel.FourthSwipeDirection   )
	// assert.Equal(t, ExpSwipeStartXCor ,  swipeOldModel.FourthSwipeStartXCor   )
	// assert.Equal(t, ExpSwipeStartYCor ,  swipeOldModel.FourthSwipeStartYCor   )
	// assert.Equal(t, ExpSwipeFinishXCor,  swipeOldModel.FourthSwipeFinishXCor  )
	// assert.Equal(t, ExpSwipeFinishYCor,  swipeOldModel.FourthSwipeFinishYCor  )

	// assert.Equal(t, ExpSwipeDirection ,  swipeOldModel.ThirdSwipeDirection   )
	// assert.Equal(t, ExpSwipeStartXCor ,  swipeOldModel.ThirdSwipeStartXCor   )
	// assert.Equal(t, ExpSwipeStartYCor ,  swipeOldModel.ThirdSwipeStartYCor   )
	// assert.Equal(t, ExpSwipeFinishXCor,  swipeOldModel.ThirdSwipeFinishXCor  )
	// assert.Equal(t, ExpSwipeFinishYCor,  swipeOldModel.ThirdSwipeFinishYCor  )

	// assert.Equal(t, ExpSwipeDirection ,  swipeOldModel.SecondSwipeDirection   )
	// assert.Equal(t, ExpSwipeStartXCor ,  swipeOldModel.SecondSwipeStartXCor   )
	// assert.Equal(t, ExpSwipeStartYCor ,  swipeOldModel.SecondSwipeStartYCor   )
	// assert.Equal(t, ExpSwipeFinishXCor,  swipeOldModel.SecondSwipeFinishXCor  )
	// assert.Equal(t, ExpSwipeFinishYCor,  swipeOldModel.SecondSwipeFinishYCor  )
}

func Test_DetermineSwipeDirection(t *testing.T) {
	concrete.DetermineSwipeDirection(&swipeRespondModel, 2)
	var ExpU int64 = 0
	var ExpD int64 = 0
	var ExpR int64 = 1
	var ExpL int64 = 1

	var ExpTotalU int64 = 0
	var ExpTotalD int64 = 0
	var ExpTotalR int64 = 1
	var ExpTotalL int64 = 1
	assert.Equal(t, ExpU, swipeRespondModel.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeRespondModel.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeRespondModel.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeRespondModel.FirstDayTotalSwipeLeftCount)

	assert.Equal(t, ExpTotalU, swipeRespondModel.TotalSwipeUpCount)
	assert.Equal(t, ExpTotalD, swipeRespondModel.TotalSwipeDownCount)
	assert.Equal(t, ExpTotalR, swipeRespondModel.TotalSwipeRightCount)
	assert.Equal(t, ExpTotalL, swipeRespondModel.TotalSwipeLeftCount)
}

func Test_CalculateSwipeFifthDay(t *testing.T) {
	concrete.CalculateSwipeFifthDay(&swipeRespondModel, &swipeOldModel, totalSwipeOldHour2)
	var ExpU int64 = 0
	var ExpD int64 = 0
	var ExpR int64 = 1
	var ExpL int64 = 0

	// var ExpU int64 = 0
	// var ExpD int64 = 0
	// var ExpR int64 = 0
	// var ExpL int64 = 0
	fmt.Println(totalSwipeOldHour2)
	// var ExpSX float64 = 0
	// var ExpSY float64 = 0
	// var ExpFX float64 = 0
	// var ExpFY float64 = 0
	var ExpSX float64 = 180
	var ExpSY float64 = 90
	var ExpFX float64 = 360
	var ExpFY float64 = 180
	assert.Equal(t, ExpU, swipeOldModel.FifthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FifthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FifthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FifthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FifthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FifthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FifthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FifthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeFourthDay(t *testing.T) {
	concrete.CalculateSwipeFourthDay(&swipeRespondModel, &swipeOldModel, totalSwipeOldHour2)
	// var ExpU int64 = 0
	// var ExpD int64 = 0
	// var ExpR int64 = 1
	// var ExpL int64 = 0

	var ExpU int64 = 0
	var ExpD int64 = 0
	var ExpR int64 = 0
	var ExpL int64 = 0
	fmt.Println(totalSwipeOldHour2)
	var ExpSX float64 = 0
	var ExpSY float64 = 0
	var ExpFX float64 = 0
	var ExpFY float64 = 0
	// var ExpSX float64 = 180
	// var ExpSY float64 = 90
	// var ExpFX float64 = 360
	// var ExpFY float64 = 180
	assert.Equal(t, ExpU, swipeOldModel.FourthDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FourthDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FourthDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FourthDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FourthDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FourthDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FourthDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FourthDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeThirdDay(t *testing.T) {
	concrete.CalculateSwipeThirdDay(&swipeRespondModel, &swipeOldModel, totalSwipeOldHour2)
	var ExpU int64 = 0
	var ExpD int64 = 0
	var ExpR int64 = 1
	var ExpL int64 = 0
	fmt.Println(totalSwipeOldHour2)
	// var ExpSX float64 = 0
	// var ExpSY float64 = 0
	// var ExpFX float64 = 0
	// var ExpFY float64 = 0
	var ExpSX float64 = 180
	var ExpSY float64 = 90
	var ExpFX float64 = 360
	var ExpFY float64 = 180
	assert.Equal(t, ExpU, swipeOldModel.ThirdDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.ThirdDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.ThirdDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.ThirdDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.ThirdDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.ThirdDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.ThirdDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.ThirdDaySwipeTotalFinishYCor)
}

func Test_CalculateSwipeSecondDay(t *testing.T) {
	concrete.CalculateSwipeSecondDay(&swipeRespondModel, &swipeOldModel, totalSwipeOldHour2)
	var ExpU int64 = 1
	var ExpD int64 = 0
	var ExpR int64 = 0
	var ExpL int64 = 0
	fmt.Println(totalSwipeOldHour2)
	// var ExpSX float64 = 60  + 180
	// var ExpSY float64 = 160+ 90
	// var ExpFX float64 = 260+ 360
	// var ExpFY float64 = 360+ 180
	var ExpSX float64 = 60
	var ExpSY float64 = 160
	var ExpFX float64 = 260
	var ExpFY float64 = 360
	assert.Equal(t, ExpU, swipeOldModel.SecondDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.SecondDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.SecondDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.SecondDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.SecondDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.SecondDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.SecondDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.SecondDaySwipeTotalFinishYCor)
}

func Test_CalculateFirstSwipeDay(t *testing.T) {
	concrete.CalculateFirstSwipeDay(&swipeRespondModel, &swipeOldModel, totalSwipeOldHour2)
	var ExpU int64 = 0
	var ExpD int64 = 1
	var ExpR int64 = 1
	var ExpL int64 = 0
	fmt.Println(totalSwipeOldHour2)
	var ExpSX float64 = 15 + 180
	var ExpSY float64 = 30 + 90
	var ExpFX float64 = 45 + 360
	var ExpFY float64 = 60 + 180
	assert.Equal(t, ExpU, swipeOldModel.FirstDayTotalSwipeUpCount)
	assert.Equal(t, ExpD, swipeOldModel.FirstDayTotalSwipeDownCount)
	assert.Equal(t, ExpR, swipeOldModel.FirstDayTotalSwipeRightCount)
	assert.Equal(t, ExpL, swipeOldModel.FirstDayTotalSwipeLeftCount)
	assert.Equal(t, ExpSX, swipeOldModel.FirstDaySwipeTotalStartXCor)
	assert.Equal(t, ExpSY, swipeOldModel.FirstDaySwipeTotalStartYCor)
	assert.Equal(t, ExpFX, swipeOldModel.FirstDaySwipeTotalFinishXCor)
	assert.Equal(t, ExpFY, swipeOldModel.FirstDaySwipeTotalFinishYCor)
}
