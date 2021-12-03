package test

import (
	"FilterWorkerService/internal/IoC"
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/Log"
	"FilterWorkerService/test/Mock/repository"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var levelBaseSession = model.LevelBaseSessionDataModel{}

var levelBaseRespondSession = model.LevelBaseSessionRespondModel{
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

var levelBaseOldSession = model.LevelBaseSessionRespondModel{}

var TotalLevelBaseSessionMinute = int64((((levelBaseRespondSession.FirstLevelSessionYearOfDay+365*levelBaseRespondSession.FirstLevelSessionYear)*24+levelBaseRespondSession.FirstLevelSessionHour)*60 + levelBaseRespondSession.FirstLevelSessionMinute) - (((levelBaseOldSession.FirstLevelSessionYearOfDay+365*levelBaseOldSession.FirstLevelSessionYear)*24+levelBaseOldSession.FirstLevelSessionHour)*60 + levelBaseOldSession.FirstLevelSessionMinute))



var levelBaseUpdateSession = model.LevelBaseSessionRespondModel{}

func Test_UpdateLevelBaseSession_UpdatedSuccess(t *testing.T) {

	var testLevelBaseDal = new(repository.MockLevelBaseSessionDal)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.LevelBaseSessionDal = testLevelBaseDal
	IoC.Logger = testLog
	var manager = concrete.LevelBaseSessionManagerConstructor()
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.ProjectId                               =   "Test"
	levelBaseOldSession_test.ClientId                                =   "Test"
	levelBaseOldSession_test.CustomerId                              =   "Test"
	levelBaseOldSession_test.TotalLevelBaseSessionMinute             =   TotalLevelBaseSessionOldMinute
	levelBaseOldSession_test.TotalLevelBaseSessionCount              =   3
	levelBaseOldSession_test.FirstLevelSessionLevelIndex             =   int64(levelBaseOldSession1.LevelIndex)
	levelBaseOldSession_test.FirstLevelSessionDuration               =   int64(levelBaseOldSession1.SessionTimeMinute)
	levelBaseOldSession_test.FirstLevelSessionYearOfDay              =   int64(levelBaseOldSession1.SessionFinishTime.YearDay())
	levelBaseOldSession_test.FirstLevelSessionYear                   =   int64(levelBaseOldSession1.SessionFinishTime.Year())
	levelBaseOldSession_test.FirstLevelSessionWeekDay                =   int64(levelBaseOldSession1.SessionFinishTime.Weekday())
	levelBaseOldSession_test.FirstLevelSessionHour                   =   int64(levelBaseOldSession1.SessionFinishTime.Hour())
	levelBaseOldSession_test.FirstLevelSessionMinute                 =   int64(levelBaseOldSession1.SessionFinishTime.Minute())
	levelBaseOldSession_test.FirstQuarterHourTotalLevelBaseSessionCount   =  2
	levelBaseOldSession_test.FirstHalfHourTotalLEvelBaseSessionCount      =  3
	levelBaseOldSession_test.FirstHourTotalLevelBaseSessionCount          =  1
	levelBaseOldSession_test.FirstTwoHourTotalLevelBaseSessionCount       =  1
	levelBaseOldSession_test.FirstThreeHourTotalLevelBaseSessionCount     =  4
	levelBaseOldSession_test.FirstSixHourTotalLevelBaseSessionCount       =  1
	levelBaseOldSession_test.FirstTwelveHourTotalLevelBaseSessionCount    =  6
	levelBaseOldSession_test.FirstDayTotalLevelBaseSessionCount           =  1
	levelBaseOldSession_test.PenultimateLevelSessionLevelIndex            =  0
	levelBaseOldSession_test.PenultimateLevelSessionLevelDuration         =  0
	levelBaseOldSession_test.LastLevelSessionLevelIndex                   =  int64(levelBaseOldSession2.LevelIndex)
	levelBaseOldSession_test.LastLevelSessionLevelDuration                =  int64(levelBaseOldSession2.SessionTimeMinute)
	levelBaseOldSession_test.LastLevelSessionYearOfDay                    =  int64(levelBaseOldSession2.SessionFinishTime.YearDay())
	levelBaseOldSession_test.LastLevelSessionYear                         =  int64(levelBaseOldSession2.SessionFinishTime.Year())
	levelBaseOldSession_test.LastLevelSessionWeekDay                      =  int64(levelBaseOldSession2.SessionFinishTime.Weekday())
	levelBaseOldSession_test.LastLevelSessionHour                         =  int64(levelBaseOldSession2.SessionFinishTime.Hour())
	levelBaseOldSession_test.LastLevelSessionMinute                       =  int64(levelBaseOldSession2.SessionFinishTime.Minute())

	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.ProjectId                              =    "Test"
	levelBaseRespondSession_test.ClientId                               =    "Test"
	levelBaseRespondSession_test.CustomerId                             =    "Test"
	levelBaseRespondSession_test.TotalLevelBaseSessionMinute            =    1
	levelBaseRespondSession_test.TotalLevelBaseSessionCount             =    1
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex            =    int64(levelBaseSession.LevelIndex)
	levelBaseRespondSession_test.FirstLevelSessionDuration              =    int64(levelBaseSession.SessionTimeMinute)
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay             =    int64(levelBaseSession.SessionFinishTime.YearDay())
	levelBaseRespondSession_test.FirstLevelSessionYear                  =    int64(levelBaseSession.SessionFinishTime.Year())
	levelBaseRespondSession_test.FirstLevelSessionWeekDay               =    int64(levelBaseSession.SessionFinishTime.Weekday())
	levelBaseRespondSession_test.FirstLevelSessionHour                  =    int64(levelBaseSession.SessionFinishTime.Hour())
	levelBaseRespondSession_test.FirstLevelSessionMinute                =    int64(levelBaseSession.SessionFinishTime.Minute())
	levelBaseRespondSession_test.FirstQuarterHourTotalLevelBaseSessionCount  = 1
	levelBaseRespondSession_test.FirstHalfHourTotalLEvelBaseSessionCount     = 1
	levelBaseRespondSession_test.FirstHourTotalLevelBaseSessionCount         = 1
	levelBaseRespondSession_test.FirstTwoHourTotalLevelBaseSessionCount      = 1
	levelBaseRespondSession_test.FirstThreeHourTotalLevelBaseSessionCount    = 1
	levelBaseRespondSession_test.FirstSixHourTotalLevelBaseSessionCount      = 1
	levelBaseRespondSession_test.FirstTwelveHourTotalLevelBaseSessionCount   = 1
	levelBaseRespondSession_test.FirstDayTotalLevelBaseSessionCount          = 1

	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	var levelBaseUpdateSession_test = levelBaseUpdateSession
	levelBaseUpdateSession_test.ProjectId                               =   "Test"
	levelBaseUpdateSession_test.ClientId                                =   "Test"
	levelBaseUpdateSession_test.CustomerId                              =   "Test"
	levelBaseUpdateSession_test.TotalLevelBaseSessionMinute             =   TotalLevelBaseSessionMinute_test
	levelBaseUpdateSession_test.TotalLevelBaseSessionCount              =   levelBaseOldSession.TotalLevelBaseSessionCount +levelBaseRespondSession.TotalLevelBaseSessionMinute
	levelBaseUpdateSession_test.FirstLevelSessionLevelIndex             =   int64(levelBaseOldSession1.LevelIndex)
	levelBaseUpdateSession_test.FirstLevelSessionDuration               =   int64(levelBaseOldSession1.SessionTimeMinute)
	levelBaseUpdateSession_test.FirstLevelSessionYearOfDay              =   int64(levelBaseOldSession1.SessionFinishTime.YearDay())
	levelBaseUpdateSession_test.FirstLevelSessionYear                   =   int64(levelBaseOldSession1.SessionFinishTime.Year())
	levelBaseUpdateSession_test.FirstLevelSessionWeekDay                =   int64(levelBaseOldSession1.SessionFinishTime.Weekday())
	levelBaseUpdateSession_test.FirstLevelSessionHour                   =   int64(levelBaseOldSession1.SessionFinishTime.Hour())
	levelBaseUpdateSession_test.FirstLevelSessionMinute                 =   int64(levelBaseOldSession1.SessionFinishTime.Minute())
	levelBaseUpdateSession_test.FourLevelSessionLevelIndex       =          levelBaseRespondSession.FirstLevelSessionLevelIndex
	levelBaseUpdateSession_test.FourLevelSessionDuration         =          levelBaseRespondSession.FirstLevelSessionDuration
	levelBaseUpdateSession_test.FirstQuarterHourTotalLevelBaseSessionCount       =   2
	levelBaseUpdateSession_test.FirstHalfHourTotalLEvelBaseSessionCount          =   3
	levelBaseUpdateSession_test.FirstHourTotalLevelBaseSessionCount              =   1 + 1
	levelBaseUpdateSession_test.FirstTwoHourTotalLevelBaseSessionCount           =   1 + 1
	levelBaseUpdateSession_test.FirstThreeHourTotalLevelBaseSessionCount         =   4 + 1
	levelBaseUpdateSession_test.FirstSixHourTotalLevelBaseSessionCount           =   1 + 1
	levelBaseUpdateSession_test.FirstTwelveHourTotalLevelBaseSessionCount        =   6 + 1
	levelBaseUpdateSession_test.FirstDayTotalLevelBaseSessionCount               =   1 + 1
	levelBaseUpdateSession_test.PenultimateLevelSessionLevelIndex                =   int64(levelBaseOldSession2.LevelIndex)
	levelBaseUpdateSession_test.PenultimateLevelSessionLevelDuration             =   int64(levelBaseOldSession2.SessionTimeMinute)
	levelBaseUpdateSession_test.LastLevelSessionLevelIndex                       =   levelBaseRespondSession.FirstLevelSessionLevelIndex
	levelBaseUpdateSession_test.LastLevelSessionLevelDuration                    =   levelBaseRespondSession.FirstLevelSessionDuration
	levelBaseUpdateSession_test.LastLevelSessionYearOfDay                        =   levelBaseRespondSession.FirstLevelSessionYearOfDay
	levelBaseUpdateSession_test.LastLevelSessionYear                             =   levelBaseRespondSession.FirstLevelSessionYear   
	levelBaseUpdateSession_test.LastLevelSessionWeekDay                          =   levelBaseRespondSession.FirstLevelSessionWeekDay 
	levelBaseUpdateSession_test.LastLevelSessionHour                             =   levelBaseRespondSession.FirstLevelSessionHour   
	levelBaseUpdateSession_test.LastLevelSessionMinute                           =   levelBaseRespondSession.FirstLevelSessionMinute 

	fmt.Println(TotalLevelBaseSessionMinute)
	testLevelBaseDal.On("UpdateLevelBaseSessionById", levelBaseOldSession_test.ClientId, &levelBaseOldSession_test).Return(nil)
	v,s,m:= manager.UpdateLevelBaseSession(&levelBaseRespondSession_test, &levelBaseOldSession_test)
	assert.Equal(t, &levelBaseUpdateSession_test, v)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
}

func Test_ConvertRawModelToResponseModel_AddSucces(t *testing.T){
	var testLevelBaseDal = new(repository.MockLevelBaseSessionDal)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.LevelBaseSessionDal = testLevelBaseDal
	IoC.Logger = testLog
	var manager = concrete.LevelBaseSessionManagerConstructor()

	var levelBaseModel_test = levelBaseSession
	levelBaseModel_test.ProjectId  = "Test"
	levelBaseModel_test.ClientId   = "Test"
	levelBaseModel_test.CustomerId = "Test"
	levelBaseModel_test.SessionStartTime = time.Date(
		2021, 11, 6, 18, 33, 58, 651387237, time.UTC)
	levelBaseModel_test.SessionFinishTime = time.Date(
		2021, 11, 6, 19, 34, 58, 651387237, time.UTC)
	levelBaseModel_test.SessionTimeMinute = 15
	levelBaseModel_test.LevelIndex    =    25
	levelBaseModel_test.LevelName     =   "25"

	var levelBaseOldSession_test = levelBaseOldSession

	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.ProjectId                              =    "Test"
	levelBaseRespondSession_test.ClientId                               =    "Test"
	levelBaseRespondSession_test.CustomerId                             =    "Test"
	levelBaseRespondSession_test.TotalLevelBaseSessionMinute            =    1
	levelBaseRespondSession_test.TotalLevelBaseSessionCount             =    1
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex            =    int64(levelBaseSession.LevelIndex)
	levelBaseRespondSession_test.FirstLevelSessionDuration              =    int64(levelBaseSession.SessionTimeMinute)
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay             =    int64(levelBaseSession.SessionFinishTime.YearDay())
	levelBaseRespondSession_test.FirstLevelSessionYear                  =    int64(levelBaseSession.SessionFinishTime.Year())
	levelBaseRespondSession_test.FirstLevelSessionWeekDay               =    int64(levelBaseSession.SessionFinishTime.Weekday())
	levelBaseRespondSession_test.FirstLevelSessionHour                  =    int64(levelBaseSession.SessionFinishTime.Hour())
	levelBaseRespondSession_test.FirstLevelSessionMinute                =    int64(levelBaseSession.SessionFinishTime.Minute())
	levelBaseRespondSession_test.FirstQuarterHourTotalLevelBaseSessionCount  = 1
	levelBaseRespondSession_test.FirstHalfHourTotalLEvelBaseSessionCount     = 1
	levelBaseRespondSession_test.FirstHourTotalLevelBaseSessionCount         = 1
	levelBaseRespondSession_test.FirstTwoHourTotalLevelBaseSessionCount      = 1
	levelBaseRespondSession_test.FirstThreeHourTotalLevelBaseSessionCount    = 1
	levelBaseRespondSession_test.FirstSixHourTotalLevelBaseSessionCount      = 1
	levelBaseRespondSession_test.FirstTwelveHourTotalLevelBaseSessionCount   = 1
	levelBaseRespondSession_test.FirstDayTotalLevelBaseSessionCount          = 1
	
	var levelBaseModel_test_byte, _ = json.EncodeJson(levelBaseModel_test)

	
	testLevelBaseDal.On("GetLevelBaseSessionById", levelBaseRespondSession_test.ClientId).Return(&levelBaseOldSession_test, 
		errors.New("null data error"))

	testLevelBaseDal.On("Add", &levelBaseRespondSession_test).Return(nil)

	v, s, m := manager.ConvertRawModelToResponseModel(levelBaseModel_test_byte)
	var value, success = v.(model.LevelBaseSessionRespondModel)
	if success == true {
		assert.Equal(t, &levelBaseRespondSession_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
	
}


func Test_CalculateLevelBaseSessionFirstQuarterHour_In15Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 48
	levelBaseOldSession_test.FirstQuarterHourTotalLevelBaseSessionCount = 2
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 16
	levelBaseRespondSession_test.FirstLevelSessionMinute = 55
	levelBaseRespondSession_test.FirstQuarterHourTotalLevelBaseSessionCount = 3
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstQuarterHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstQuarterHour int64 = 5
	assert.Equal(t, ExpFirstQuarterHour, levelBaseOldSession_test.FirstQuarterHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstQuarterHour_Out15Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 45
	levelBaseOldSession_test.FirstQuarterHourTotalLevelBaseSessionCount = 2
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 17
	levelBaseRespondSession_test.FirstLevelSessionMinute = 01
	levelBaseRespondSession_test.FirstQuarterHourTotalLevelBaseSessionCount = 3
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstQuarterHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstQuarterHour int64 = 2
	assert.Equal(t, ExpFirstQuarterHour, levelBaseOldSession_test.FirstQuarterHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstHalfHour_In30Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstHalfHourTotalLEvelBaseSessionCount = 4
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 17
	levelBaseRespondSession_test.FirstLevelSessionMinute = 20
	levelBaseRespondSession_test.FirstHalfHourTotalLEvelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstHalfHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute)
	var ExpFirstHalfHour int64 = 5
	assert.Equal(t, ExpFirstHalfHour, levelBaseOldSession_test.FirstHalfHourTotalLEvelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstHalfHour_Out30Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstHalfHourTotalLEvelBaseSessionCount = 4
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 17
	levelBaseRespondSession_test.FirstLevelSessionMinute = 21
	levelBaseRespondSession_test.FirstHalfHourTotalLEvelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstHalfHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute)
	var ExpFirstHalfHour int64 = 4
	assert.Equal(t, ExpFirstHalfHour, levelBaseOldSession_test.FirstHalfHourTotalLEvelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstHour_In60Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstHourTotalLevelBaseSessionCount = 2
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 350
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 17
	levelBaseRespondSession_test.FirstLevelSessionMinute = 50
	levelBaseRespondSession_test.FirstHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstHour int64 = 2 + 1
	assert.Equal(t, ExpFirstHour, levelBaseOldSession.FirstHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstHour_Out60Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstHourTotalLevelBaseSessionCount = 2
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 17
	levelBaseRespondSession_test.FirstLevelSessionMinute = 51
	levelBaseRespondSession_test.FirstHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstHour int64 = 2
	assert.Equal(t, ExpFirstHour, levelBaseOldSession.FirstHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstTwoHour_In120Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstHourTotalLevelBaseSessionCount = 9
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 18
	levelBaseRespondSession_test.FirstLevelSessionMinute = 50
	levelBaseRespondSession_test.FirstHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstTwoHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstTwoHour int64 = 9 + 1
	assert.Equal(t, ExpFirstTwoHour, levelBaseOldSession.FirstTwoHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstTwoHour_Out120Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstTwoHourTotalLevelBaseSessionCount = 9
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 18
	levelBaseRespondSession_test.FirstLevelSessionMinute = 51
	levelBaseRespondSession_test.FirstTwoHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstTwoHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstTwoHour int64 = 9
	assert.Equal(t, ExpFirstTwoHour, levelBaseOldSession.FirstTwoHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstThreeHour_In180Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstThreeHourTotalLevelBaseSessionCount = 12
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 19
	levelBaseRespondSession_test.FirstLevelSessionMinute = 30
	levelBaseRespondSession_test.FirstThreeHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstThreeHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstThreeHour int64 =12+1
	assert.Equal(t, ExpFirstThreeHour, levelBaseOldSession.FirstThreeHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstThreeHour_Out180Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 16
	levelBaseOldSession_test.FirstLevelSessionMinute = 50
	levelBaseOldSession_test.FirstThreeHourTotalLevelBaseSessionCount = 12
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 19
	levelBaseRespondSession_test.FirstLevelSessionMinute = 51
	levelBaseRespondSession_test.FirstThreeHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstThreeHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstThreeHour int64 =12
	assert.Equal(t, ExpFirstThreeHour, levelBaseOldSession.FirstThreeHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstSixHour_In360Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 12
	levelBaseOldSession_test.FirstLevelSessionMinute = 00
	levelBaseOldSession_test.FirstSixHourTotalLevelBaseSessionCount = 18
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 18
	levelBaseRespondSession_test.FirstLevelSessionMinute = 00
	levelBaseRespondSession_test.FirstSixHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstSixHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute)
	var ExpFirstSixHour int64 = 18 + 1
	assert.Equal(t, ExpFirstSixHour, levelBaseOldSession.FirstSixHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstSixHour_Out360Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 12
	levelBaseOldSession_test.FirstLevelSessionMinute = 00
	levelBaseOldSession_test.FirstSixHourTotalLevelBaseSessionCount = 18
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 18
	levelBaseRespondSession_test.FirstLevelSessionMinute = 01
	levelBaseRespondSession_test.FirstSixHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstSixHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute)
	var ExpFirstSixHour int64 = 18
	assert.Equal(t, ExpFirstSixHour, levelBaseOldSession.FirstSixHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstTwelveHour_In720Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 12
	levelBaseOldSession_test.FirstLevelSessionMinute = 01
	levelBaseOldSession_test.FirstTwelveHourTotalLevelBaseSessionCount = 20
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 347
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 00
	levelBaseRespondSession_test.FirstLevelSessionMinute = 01
	levelBaseRespondSession_test.FirstTwelveHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstTwelveHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstTwelveHour int64 = 20+1
	assert.Equal(t, ExpFirstTwelveHour, levelBaseOldSession.FirstTwelveHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstTwelveHour_Out720Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 12
	levelBaseOldSession_test.FirstLevelSessionMinute = 01
	levelBaseOldSession_test.FirstTwelveHourTotalLevelBaseSessionCount = 20
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 347
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 00
	levelBaseRespondSession_test.FirstLevelSessionMinute = 02
	levelBaseRespondSession_test.FirstTwelveHourTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstTwelveHour(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstTwelveHour int64 = 20
	assert.Equal(t, ExpFirstTwelveHour, levelBaseOldSession.FirstTwelveHourTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstDay_In1440Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 12
	levelBaseOldSession_test.FirstLevelSessionMinute = 01
	levelBaseOldSession_test.FirstDayTotalLevelBaseSessionCount = 30
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 347
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 12
	levelBaseRespondSession_test.FirstLevelSessionMinute = 01
	levelBaseRespondSession_test.FirstDayTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstDay(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstDay int64 = 30 + 1
	assert.Equal(t, ExpFirstDay, levelBaseOldSession.FirstDayTotalLevelBaseSessionCount)
}

func Test_CalculateLevelBaseSessionFirstDay_Out1440Minutes(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.FirstLevelSessionYearOfDay = 346
	levelBaseOldSession_test.FirstLevelSessionYear = 2021
	levelBaseOldSession_test.FirstLevelSessionHour = 12
	levelBaseOldSession_test.FirstLevelSessionMinute = 01
	levelBaseOldSession_test.FirstDayTotalLevelBaseSessionCount = 30
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionYearOfDay = 347
	levelBaseRespondSession_test.FirstLevelSessionYear = 2021
	levelBaseRespondSession_test.FirstLevelSessionHour = 12
	levelBaseRespondSession_test.FirstLevelSessionMinute = 02
	levelBaseRespondSession_test.FirstDayTotalLevelBaseSessionCount = 1
	var TotalLevelBaseSessionMinute_test = int64((((levelBaseRespondSession_test.FirstLevelSessionYearOfDay+365*levelBaseRespondSession_test.FirstLevelSessionYear)*24+levelBaseRespondSession_test.FirstLevelSessionHour)*60 + levelBaseRespondSession_test.FirstLevelSessionMinute) - (((levelBaseOldSession_test.FirstLevelSessionYearOfDay+365*levelBaseOldSession_test.FirstLevelSessionYear)*24+levelBaseOldSession_test.FirstLevelSessionHour)*60 + levelBaseOldSession_test.FirstLevelSessionMinute))
	concrete.CalculateLevelBaseSessionFirstDay(&levelBaseRespondSession_test, &levelBaseOldSession_test, TotalLevelBaseSessionMinute_test)
	fmt.Println(TotalLevelBaseSessionMinute_test)
	var ExpFirstDay int64 = 30
	assert.Equal(t, ExpFirstDay, levelBaseOldSession.FirstDayTotalLevelBaseSessionCount)
}

func Test_CalculateLevelIndexBaseSession_SecondSession(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.TotalLevelBaseSessionCount = 2
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex = 35
	levelBaseRespondSession_test.FirstLevelSessionDuration   = 23
	concrete.CalculateLevelIndexBaseSession(&levelBaseRespondSession_test, &levelBaseOldSession_test)
	fmt.Println(levelBaseOldSession.TotalLevelBaseSessionCount)
	var ExpSecondLevelindex int64 =    levelBaseRespondSession_test.FirstLevelSessionLevelIndex         
	var ExpSecondDuration int64 =      levelBaseRespondSession_test.FirstLevelSessionDuration       
	var ExpThirdindex int64 =          0   
	var ExpThirdDuration int64 =       0      
	var ExpFourLevelindex int64 =      0       
	var ExpFourDuration int64 =        0     
	var ExpFiveLevelindex int64 =      0       
	var ExpFiveDuration int64 =        0     
	var ExpSixLevelindex int64 =       0      
	var ExpSixDuration int64 =         0    
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

func Test_CalculateLevelIndexBaseSession_ThirdSession(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.TotalLevelBaseSessionCount = 3
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex = 35
	levelBaseRespondSession_test.FirstLevelSessionDuration   = 23
	concrete.CalculateLevelIndexBaseSession(&levelBaseRespondSession_test, &levelBaseOldSession_test)
	fmt.Println(levelBaseOldSession.TotalLevelBaseSessionCount)
	var ExpSecondLevelindex int64 =    0         
	var ExpSecondDuration int64 =      0     
	var ExpThirdindex int64 =          levelBaseRespondSession_test.FirstLevelSessionLevelIndex   
	var ExpThirdDuration int64 =       levelBaseRespondSession_test.FirstLevelSessionDuration        
	var ExpFourLevelindex int64 =      0       
	var ExpFourDuration int64 =        0     
	var ExpFiveLevelindex int64 =      0       
	var ExpFiveDuration int64 =        0     
	var ExpSixLevelindex int64 =       0      
	var ExpSixDuration int64 =         0    
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

func Test_CalculateLevelIndexBaseSession_FourthSession(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.TotalLevelBaseSessionCount = 4
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex = 35
	levelBaseRespondSession_test.FirstLevelSessionDuration   = 23
	concrete.CalculateLevelIndexBaseSession(&levelBaseRespondSession_test, &levelBaseOldSession_test)
	fmt.Println(levelBaseOldSession.TotalLevelBaseSessionCount)
	var ExpSecondLevelindex int64 =    0         
	var ExpSecondDuration int64 =      0     
	var ExpThirdindex int64 =          0   
	var ExpThirdDuration int64 =       0      
	var ExpFourLevelindex int64 =      levelBaseRespondSession_test.FirstLevelSessionLevelIndex       
	var ExpFourDuration int64 =        levelBaseRespondSession_test.FirstLevelSessionDuration       
	var ExpFiveLevelindex int64 =      0       
	var ExpFiveDuration int64 =        0     
	var ExpSixLevelindex int64 =       0      
	var ExpSixDuration int64 =         0    
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

func Test_CalculateLevelIndexBaseSession_FifthSession(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.TotalLevelBaseSessionCount = 5
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex = 35
	levelBaseRespondSession_test.FirstLevelSessionDuration   = 23
	concrete.CalculateLevelIndexBaseSession(&levelBaseRespondSession_test, &levelBaseOldSession_test)
	fmt.Println(levelBaseOldSession.TotalLevelBaseSessionCount)
	var ExpSecondLevelindex int64 =    0         
	var ExpSecondDuration int64 =      0     
	var ExpThirdindex int64 =          0   
	var ExpThirdDuration int64 =       0      
	var ExpFourLevelindex int64 =      0       
	var ExpFourDuration int64 =        0     
	var ExpFiveLevelindex int64 =      levelBaseRespondSession_test.FirstLevelSessionLevelIndex       
	var ExpFiveDuration int64 =        levelBaseRespondSession_test.FirstLevelSessionDuration       
	var ExpSixLevelindex int64 =       0      
	var ExpSixDuration int64 =         0    
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

func Test_CalculateLevelIndexBaseSession_SixthSession(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.TotalLevelBaseSessionCount = 6
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex = 35
	levelBaseRespondSession_test.FirstLevelSessionDuration   = 23
	concrete.CalculateLevelIndexBaseSession(&levelBaseRespondSession_test, &levelBaseOldSession_test)
	fmt.Println(levelBaseOldSession.TotalLevelBaseSessionCount)
	var ExpSecondLevelindex int64 =    0         
	var ExpSecondDuration int64 =      0     
	var ExpThirdindex int64 =          0   
	var ExpThirdDuration int64 =       0      
	var ExpFourLevelindex int64 =      0       
	var ExpFourDuration int64 =        0     
	var ExpFiveLevelindex int64 =      0       
	var ExpFiveDuration int64 =        0     
	var ExpSixLevelindex int64 =       levelBaseRespondSession_test.FirstLevelSessionLevelIndex      
	var ExpSixDuration int64 =         levelBaseRespondSession_test.FirstLevelSessionDuration      
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

func Test_CalculateLevelIndexBaseSession_SeventhSession(t *testing.T) {
	var levelBaseOldSession_test = levelBaseOldSession
	levelBaseOldSession_test.TotalLevelBaseSessionCount = 7
	var levelBaseRespondSession_test = levelBaseRespondSession
	levelBaseRespondSession_test.FirstLevelSessionLevelIndex = 35
	levelBaseRespondSession_test.FirstLevelSessionDuration   = 23
	concrete.CalculateLevelIndexBaseSession(&levelBaseRespondSession_test, &levelBaseOldSession_test)
	fmt.Println(levelBaseOldSession.TotalLevelBaseSessionCount)
	var ExpSecondLevelindex int64 =    0         
	var ExpSecondDuration int64 =      0     
	var ExpThirdindex int64 =          0   
	var ExpThirdDuration int64 =       0      
	var ExpFourLevelindex int64 =      0       
	var ExpFourDuration int64 =        0     
	var ExpFiveLevelindex int64 =      0       
	var ExpFiveDuration int64 =        0     
	var ExpSixLevelindex int64 =       0      
	var ExpSixDuration int64 =         0    
	var ExpSevenLevelindex int64 =     levelBaseRespondSession_test.FirstLevelSessionLevelIndex        
	var ExpSevenDuration int64 =       levelBaseRespondSession_test.FirstLevelSessionDuration        
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