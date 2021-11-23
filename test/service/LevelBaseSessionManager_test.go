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

var levelBaseSession = model.LevelBaseSessionDataModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	SessionStartTime: time.Date(
		2021, 11, 6, 18, 33, 58, 651387237, time.UTC),
	SessionFinishTime: time.Date(
		2021, 11, 6, 19, 34, 58, 651387237, time.UTC),
	SessionTimeMinute: 15,
	LevelIndex:        25,
	LevelName:         "25",
}

var levelBaseRespondSession = model.LevelBaseSessionRespondModel{
	ProjectId:                                  "Test",
	ClientId:                                   "Test",
	CustomerId:                                 "Test",
	TotalLevelBaseSessionMinute:                1,
	TotalLevelBaseSessionCount:                 1,
	FirstLevelSessionLevelIndex:                int64(levelBaseSession.LevelIndex),
	FirstLevelSessionDuration:                  int64(levelBaseSession.SessionTimeMinute),
	FirstLevelSessionYearOfDay:                 int64(levelBaseSession.SessionFinishTime.YearDay()),
	FirstLevelSessionYear:                      int64(levelBaseSession.SessionFinishTime.Year()),
	FirstLevelSessionWeekDay:                   int64(levelBaseSession.SessionFinishTime.Weekday()),
	FirstLevelSessionHour:                      int64(levelBaseSession.SessionFinishTime.Hour()),
	FirstLevelSessionMinute:                    int64(levelBaseSession.SessionFinishTime.Minute()),
	SecondLevelSessionLevelIndex:               0,
	SecondLevelSessionDuration:                 0,
	ThirdLevelSessionLevelIndex:                0,
	ThirdLevelSessionDuration:                  0,
	FourLevelSessionLevelIndex:                 0,
	FourLevelSessionDuration:                   0,
	FiveLevelSessionLevelIndex:                 0,
	FiveLevelSessionDuration:                   0,
	SixLevelSessionLevelIndex:                  0,
	SixLevelSessionDuration:                    0,
	SevenLevelSessionLevelIndex:                0,
	SevenLevelSessionDuration:                  0,
	FirstQuarterHourTotalLevelBaseSessionCount: 1,
	FirstHalfHourTotalLEvelBaseSessionCount:    1,
	FirstHourTotalLevelBaseSessionCount:        1,
	FirstTwoHourTotalLevelBaseSessionCount:     1,
	FirstThreeHourTotalLevelBaseSessionCount:   1,
	FirstSixHourTotalLevelBaseSessionCount:     1,
	FirstTwelveHourTotalLevelBaseSessionCount:  1,
	FirstDayTotalLevelBaseSessionCount:         1,
	PenultimateLevelSessionLevelIndex:          0,
	PenultimateLevelSessionLevelDuration:       0,
	LastLevelSessionLevelIndex:                 0,
	LastLevelSessionLevelDuration:              0,
	LastLevelSessionYearOfDay:                  0,
	LastLevelSessionYear:                       0,
	LastLevelSessionWeekDay:                    0,
	LastLevelSessionHour:                       0,
	LastLevelSessionMinute:                     0,
}

var levelBaseOldSession1 = model.LevelBaseSessionDataModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	SessionStartTime: time.Date(
		2021, 11, 6, 18, 33, 58, 651387237, time.UTC),
	SessionFinishTime: time.Date(
		2021, 11, 6, 18, 34, 58, 651387237, time.UTC),
	SessionTimeMinute: 15,
	LevelIndex:        25,
	LevelName:         "25",
}

var levelBaseOldSession2 = model.LevelBaseSessionDataModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	SessionStartTime: time.Date(
		2021, 11, 7, 18, 33, 58, 651387237, time.UTC),
	SessionFinishTime: time.Date(
		2021, 11, 7, 18, 34, 58, 651387237, time.UTC),
	SessionTimeMinute: 15,
	LevelIndex:        25,
	LevelName:         "25",
}

var TotalLevelBaseSessionOldMinute = int64((((levelBaseOldSession2.SessionFinishTime.YearDay()+365*levelBaseOldSession2.SessionFinishTime.Year())*24+levelBaseOldSession2.SessionFinishTime.Hour())*60 + levelBaseOldSession2.SessionFinishTime.Minute()) - (((levelBaseOldSession1.SessionFinishTime.YearDay()+365*levelBaseOldSession1.SessionFinishTime.Year())*24+levelBaseOldSession1.SessionFinishTime.Hour())*60 + levelBaseOldSession1.SessionFinishTime.Minute()))

var levelBaseOldSession = model.LevelBaseSessionRespondModel{
	ProjectId:                                  "Test",
	ClientId:                                   "Test",
	CustomerId:                                 "Test",
	TotalLevelBaseSessionMinute:                TotalLevelBaseSessionOldMinute,
	TotalLevelBaseSessionCount:                 3,
	FirstLevelSessionLevelIndex:                int64(levelBaseOldSession1.LevelIndex),
	FirstLevelSessionDuration:                  int64(levelBaseOldSession1.SessionTimeMinute),
	FirstLevelSessionYearOfDay:                 int64(levelBaseOldSession1.SessionFinishTime.YearDay()),
	FirstLevelSessionYear:                      int64(levelBaseOldSession1.SessionFinishTime.Year()),
	FirstLevelSessionWeekDay:                   int64(levelBaseOldSession1.SessionFinishTime.Weekday()),
	FirstLevelSessionHour:                      int64(levelBaseOldSession1.SessionFinishTime.Hour()),
	FirstLevelSessionMinute:                    int64(levelBaseOldSession1.SessionFinishTime.Minute()),
	SecondLevelSessionLevelIndex:               0,
	SecondLevelSessionDuration:                 0,
	ThirdLevelSessionLevelIndex:                0,
	ThirdLevelSessionDuration:                  0,
	FourLevelSessionLevelIndex:                 0,
	FourLevelSessionDuration:                   0,
	FiveLevelSessionLevelIndex:                 0,
	FiveLevelSessionDuration:                   0,
	SixLevelSessionLevelIndex:                  0,
	SixLevelSessionDuration:                    0,
	SevenLevelSessionLevelIndex:                0,
	SevenLevelSessionDuration:                  0,
	FirstQuarterHourTotalLevelBaseSessionCount: 2,
	FirstHalfHourTotalLEvelBaseSessionCount:    3,
	FirstHourTotalLevelBaseSessionCount:        1,
	FirstTwoHourTotalLevelBaseSessionCount:     1,
	FirstThreeHourTotalLevelBaseSessionCount:   4,
	FirstSixHourTotalLevelBaseSessionCount:     1,
	FirstTwelveHourTotalLevelBaseSessionCount:  6,
	FirstDayTotalLevelBaseSessionCount:         1,
	PenultimateLevelSessionLevelIndex:          0,
	PenultimateLevelSessionLevelDuration:       0,
	LastLevelSessionLevelIndex:                 int64(levelBaseOldSession2.LevelIndex),
	LastLevelSessionLevelDuration:              int64(levelBaseOldSession2.SessionTimeMinute),
	LastLevelSessionYearOfDay:                  int64(levelBaseOldSession2.SessionFinishTime.YearDay()),
	LastLevelSessionYear:                       int64(levelBaseOldSession2.SessionFinishTime.Year()),
	LastLevelSessionWeekDay:                    int64(levelBaseOldSession2.SessionFinishTime.Weekday()),
	LastLevelSessionHour:                       int64(levelBaseOldSession2.SessionFinishTime.Hour()),
	LastLevelSessionMinute:                     int64(levelBaseOldSession2.SessionFinishTime.Minute()),
}

var TotalLevelBaseSessionMinute = int64((((levelBaseRespondSession.FirstLevelSessionYearOfDay+365*levelBaseRespondSession.FirstLevelSessionYear)*24+levelBaseRespondSession.FirstLevelSessionHour)*60 + levelBaseRespondSession.FirstLevelSessionMinute) - (((levelBaseOldSession.FirstLevelSessionYearOfDay+365*levelBaseOldSession.FirstLevelSessionYear)*24+levelBaseOldSession.FirstLevelSessionHour)*60 + levelBaseOldSession.FirstLevelSessionMinute))



var levelBaseUpdateSession = model.LevelBaseSessionRespondModel{
	ProjectId:                                  "Test",
	ClientId:                                   "Test",
	CustomerId:                                 "Test",
	TotalLevelBaseSessionMinute:                TotalLevelBaseSessionMinute,
	TotalLevelBaseSessionCount:                 levelBaseOldSession.TotalLevelBaseSessionCount +levelBaseRespondSession.TotalLevelBaseSessionMinute,
	FirstLevelSessionLevelIndex:                int64(levelBaseOldSession1.LevelIndex),
	FirstLevelSessionDuration:                  int64(levelBaseOldSession1.SessionTimeMinute),
	FirstLevelSessionYearOfDay:                 int64(levelBaseOldSession1.SessionFinishTime.YearDay()),
	FirstLevelSessionYear:                      int64(levelBaseOldSession1.SessionFinishTime.Year()),
	FirstLevelSessionWeekDay:                   int64(levelBaseOldSession1.SessionFinishTime.Weekday()),
	FirstLevelSessionHour:                      int64(levelBaseOldSession1.SessionFinishTime.Hour()),
	FirstLevelSessionMinute:                    int64(levelBaseOldSession1.SessionFinishTime.Minute()),
	SecondLevelSessionLevelIndex:               0,
	SecondLevelSessionDuration:                 0,
	ThirdLevelSessionLevelIndex:                0,
	ThirdLevelSessionDuration:                  0,
	FourLevelSessionLevelIndex:                 levelBaseRespondSession.FirstLevelSessionLevelIndex,
	FourLevelSessionDuration:                   levelBaseRespondSession.FirstLevelSessionDuration,
	FiveLevelSessionLevelIndex:                 0,
	FiveLevelSessionDuration:                   0,
	SixLevelSessionLevelIndex:                  0,
	SixLevelSessionDuration:                    0,
	SevenLevelSessionLevelIndex:                0,
	SevenLevelSessionDuration:                  0,
	FirstQuarterHourTotalLevelBaseSessionCount: 2,
	FirstHalfHourTotalLEvelBaseSessionCount:    3,
	FirstHourTotalLevelBaseSessionCount:        1 + 1,
	FirstTwoHourTotalLevelBaseSessionCount:     1 + 1,
	FirstThreeHourTotalLevelBaseSessionCount:   4 + 1,
	FirstSixHourTotalLevelBaseSessionCount:     1 + 1,
	FirstTwelveHourTotalLevelBaseSessionCount:  6 + 1,
	FirstDayTotalLevelBaseSessionCount:         1 + 1,
	PenultimateLevelSessionLevelIndex:          int64(levelBaseOldSession2.LevelIndex),
	PenultimateLevelSessionLevelDuration:       int64(levelBaseOldSession2.SessionTimeMinute),
	LastLevelSessionLevelIndex:                 levelBaseRespondSession.FirstLevelSessionLevelIndex,
	LastLevelSessionLevelDuration:              levelBaseRespondSession.FirstLevelSessionDuration,
	LastLevelSessionYearOfDay:                  levelBaseRespondSession.FirstLevelSessionYearOfDay,
	LastLevelSessionYear:                       levelBaseRespondSession.FirstLevelSessionYear   ,
	LastLevelSessionWeekDay:                    levelBaseRespondSession.FirstLevelSessionWeekDay ,
	LastLevelSessionHour:                       levelBaseRespondSession.FirstLevelSessionHour   ,
	LastLevelSessionMinute:                     levelBaseRespondSession.FirstLevelSessionMinute ,
}

func Test_UpdateLevelBaseSession_Updated(t *testing.T) {
	var testLevelBaseDal = new(repository.MockLevelBaseSessionDal)
	var manager = concrete.LevelBaseSessionManager{
		ILevelBaseSessionDal: testLevelBaseDal,
		IJsonParser:          &gojson.GoJson{},
	}
	fmt.Println(TotalLevelBaseSessionMinute)
	testLevelBaseDal.On("UpdateLevelBaseSessionById", levelBaseOldSession.ClientId, &levelBaseOldSession).Return(nil)
	v,s,m:= manager.UpdateLevelBaseSession(&levelBaseRespondSession, &levelBaseOldSession)
	assert.Equal(t, &levelBaseUpdateSession, v)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
}

func Test_ConvertRawModelToResponseModel_AddSucces(t *testing.T){
	var testLevelBaseDal = new(repository.MockLevelBaseSessionDal)
	var manager = concrete.LevelBaseSessionManager{
		ILevelBaseSessionDal: testLevelBaseDal,
		IJsonParser:          &gojson.GoJson{},
	}
	bytData, _ := manager.IJsonParser.EncodeJson(levelBaseSession)
	testLevelBaseDal.On("GetLevelBaseSessionById", levelBaseRespondSession.ClientId).Return(&levelBaseOldSession, 
		errors.New("mongo: no documents in result"))
	testLevelBaseDal.On("Add", &levelBaseRespondSession).Return(nil)
	v, s, m := manager.ConvertRawModelToResponseModel(bytData)
	assert.Equal(t, &levelBaseRespondSession, v)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
	
}

func Test_CalculateLevelBaseSessionFirstQuarterHour(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstQuarterHour(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstQuarterHour int64 = 0 + 2
	var ExpFirstQuarterHour int64 = 1 + 2
	assert.Equal(t, ExpFirstQuarterHour, levelBaseOldSession.FirstQuarterHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstHalfHour(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstHalfHour(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstHalfHour    int64 = 1 + 3
	var ExpFirstHalfHour int64 = 0 + 3
	assert.Equal(t, ExpFirstHalfHour, levelBaseOldSession.FirstHalfHourTotalLEvelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstHour(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstHour(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstHour        int64 = 1 + 1
	var ExpFirstHour int64 = 0 + 1
	assert.Equal(t, ExpFirstHour, levelBaseOldSession.FirstHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstTwoHour(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstTwoHour(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstTwoHour     int64 = 1 + 1
	var ExpFirstTwoHour int64 = 0 + 1
	assert.Equal(t, ExpFirstTwoHour, levelBaseOldSession.FirstTwoHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstThreeHour(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstThreeHour(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstThreeHour   int64 = 1 + 4
	var ExpFirstThreeHour int64 = 0 + 4
	assert.Equal(t, ExpFirstThreeHour, levelBaseOldSession.FirstThreeHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstSixHour(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstSixHour(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstSixHour     int64 = 1 + 1
	var ExpFirstSixHour int64 = 0 + 1
	assert.Equal(t, ExpFirstSixHour, levelBaseOldSession.FirstSixHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstTwelveHour(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstTwelveHour(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstTwelveHour  int64 = 1 + 6
	var ExpFirstTwelveHour int64 = 0 + 6
	assert.Equal(t, ExpFirstTwelveHour, levelBaseOldSession.FirstTwelveHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstDay(t *testing.T) {
	concrete.CalculateLevelBaseSessionFirstDay(&levelBaseRespondSession, &levelBaseOldSession, TotalLevelBaseSessionMinute)
	fmt.Println(TotalLevelBaseSessionMinute)
	//var ExpFirstDay         int64 = 1 + 1
	var ExpFirstDay int64 = 0 + 1
	assert.Equal(t, ExpFirstDay, levelBaseOldSession.FirstDayTotalLevelBaseSessionCount)
}

func Test_CalculateLevelIndexBaseSession(t *testing.T) {
	concrete.CalculateLevelIndexBaseSession(&levelBaseRespondSession, &levelBaseOldSession)
	fmt.Println(levelBaseOldSession.TotalLevelBaseSessionCount)
	var ExpSecondLevelindex int64 =    0         
	var ExpSecondDuration int64 =      0     
	var ExpThirdindex int64 =          0   
	var ExpThirdDuration int64 =       0      
	var ExpFourLevelindex int64 =      0       
	var ExpFourDuration int64 =        0     
	var ExpFiveLevelindex int64 =      0       
	var ExpFiveDuration int64 =        0     
	var ExpSixLevelindex int64 =       levelBaseRespondSession.FirstLevelSessionLevelIndex      
	var ExpSixDuration int64 =         levelBaseRespondSession.FirstLevelSessionDuration      
	var ExpSevenLevelindex int64 =     0        
	var ExpSevenDuration int64 =       0      
	assert.Equal(t, ExpSecondLevelindex, levelBaseOldSession.SecondLevelSessionLevelIndex)
	assert.Equal(t, ExpSecondDuration, levelBaseOldSession.SecondLevelSessionDuration)
	assert.Equal(t, ExpThirdindex, levelBaseOldSession.ThirdLevelSessionLevelIndex)
	assert.Equal(t, ExpThirdDuration, levelBaseOldSession.ThirdLevelSessionLevelIndex)
	assert.Equal(t, ExpFourLevelindex, levelBaseOldSession.FourLevelSessionLevelIndex)
	assert.Equal(t, ExpFourDuration, levelBaseOldSession.FourLevelSessionDuration)
	assert.Equal(t, ExpFiveLevelindex, levelBaseOldSession.FiveLevelSessionLevelIndex)
	assert.Equal(t, ExpFiveDuration, levelBaseOldSession.FiveLevelSessionDuration)
	assert.Equal(t, ExpSixLevelindex, levelBaseOldSession.SixLevelSessionLevelIndex)
	assert.Equal(t, ExpSixDuration, levelBaseOldSession.SixLevelSessionDuration)
	assert.Equal(t, ExpSevenLevelindex, levelBaseOldSession.SevenLevelSessionLevelIndex)
	assert.Equal(t, ExpSevenDuration, levelBaseOldSession.SevenLevelSessionDuration)
}
