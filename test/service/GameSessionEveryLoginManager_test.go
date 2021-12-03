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

var sessionModel = model.GameSessionEveryLoginModel{}

var sessionRespondModel = model.GameSessionEveryLoginRespondModel{}

var sessionOldModel = model.GameSessionEveryLoginRespondModel{}

var sessionUpdateModel = model.GameSessionEveryLoginRespondModel{}

func Test_UpdateGameSession_UpdateSuccess(t *testing.T) {
	var testSessionDal = new(repository.MockGameSessionEveryLoginDal)
	var testCache = new(service.MockCacheService)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.GameSessionEveryLoginDal = testSessionDal
	IoC.Logger = testLog
	IoC.CacheService = testCache
	var manager = concrete.GameSessionEveryLoginManagerConstructor()
	var first = time.Date(
	2021, 11, 6, 18, 34, 58, 651387237, time.UTC)
    var first2 = time.Date(
	2021, 11, 7, 18, 34, 58, 651387237, time.UTC)
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.ProjectId                                 =              "Test"
	sessionOldModel_test.ClientId                                  =              "Test"
	sessionOldModel_test.CustomerId                                =              "Test"
	sessionOldModel_test.FirstSessionYearOfDay                     =              int64(first.YearDay())
	sessionOldModel_test.FirstSessionYear                          =              int64(first.Year())
	sessionOldModel_test.FirstSessionWeekDay                       =              int64(first.Weekday())
	sessionOldModel_test.FirstSessionHour                          =              int64(first.Hour())
	sessionOldModel_test.FirstSessionDuration                      =              int64(sessionModel.SessionTimeMinute)
	sessionOldModel_test.FirstSessionMinute                        =              int64(first.Minute())
	sessionOldModel_test.LastSessionYearOfDay                       =             int64(first2.YearDay())
	sessionOldModel_test.LastSessionYear                            =             int64(first2.Year())
	sessionOldModel_test.LastSessionHour                            =             int64(first2.Hour())
	sessionOldModel_test.LastSessionDuration                        =             int64(sessionModel.SessionTimeMinute)
	sessionOldModel_test.LastSessionMinute                          =             int64(first2.Minute())
	sessionOldModel_test.TotalSessionDay                            =             1
	sessionOldModel_test.TotalSessionHour                           =             24
	sessionOldModel_test.TotalSessionMinute                         =             1440
	sessionOldModel_test.TotalSessionDuration                       =             18
	sessionOldModel_test.TotalSessionCount                          =             2
	sessionOldModel_test.FirstDayTotalSessionCount                  =             3
	sessionOldModel_test.FirstDayTotalSessionDuration               =             18
	sessionOldModel_test.MinSessionDuration                         =             19
	sessionOldModel_test.MaxSessionDuration                        =             29
	sessionOldModel_test.DailyAvegareSessionCount                   =             1
	sessionOldModel_test.DailyAverageSessionDuration               =             float64(sessionModel.SessionTimeMinute)
	sessionOldModel_test.SessionBasedAvegareSessionDuration         =             float64(sessionModel.SessionTimeMinute)
	sessionOldModel_test.DailyAvegareSessionCountMinusFirstDaySessionCount =       0
	sessionOldModel_test.DailyAvegareSessionDurationMinusFirstDaySessionDuration = 0
	sessionOldModel_test.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = 0
	sessionOldModel_test.SessionBasedAvegareSessionDurationMinusLastSessionDuration = float64(sessionModel.SessionTimeMinute)
	sessionOldModel_test.SundaySessionCount     =  2
	sessionOldModel_test.MondaySessionCount     =  9
	sessionOldModel_test.TuesdaySessionCount    =  8
	sessionOldModel_test.WednesdaySessionCount  =  0
	sessionOldModel_test.ThursdaySessionCount   =  0
	sessionOldModel_test.FridaySessionCount     =  1
	sessionOldModel_test.SaturdaySessionCount   =  0
	sessionOldModel_test.AmSessionCount         =  0
	sessionOldModel_test.PmSessionCount         =  1
	sessionOldModel_test.Session0To5HourCount    = 5
	sessionOldModel_test.Session6To11HourCount   = 0
	sessionOldModel_test.Session12To17HourCount  = 1
	sessionOldModel_test.Session18To23HourCount  = 87

	var sessionModel_test = sessionModel
	sessionModel_test.ProjectId  = "Test"
	sessionModel_test.ClientId   = "Test"
	sessionModel_test.CustomerId = "Test"
	sessionModel_test.SessionStartTime = time.Date(
		2021, 11, 7, 22, 34, 36, 651387237, time.UTC)
	sessionModel_test.SessionFinishTime = time.Date(
		2021, 11, 8, 18, 34, 36, 651387237, time.UTC)
	sessionModel_test.SessionTimeMinute = 33
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.ProjectId                                =               "Test"
	sessionRespondModel_test.ClientId                                 =               "Test"
	sessionRespondModel_test.CustomerId                               =               "Test"
	sessionRespondModel_test.FirstSessionYearOfDay                    =               int64(sessionModel_test.SessionFinishTime.YearDay())
	sessionRespondModel_test.FirstSessionYear                         =               int64(sessionModel_test.SessionFinishTime.Year())
	sessionRespondModel_test.FirstSessionWeekDay                      =               int64(sessionModel_test.SessionFinishTime.Weekday())
	sessionRespondModel_test.FirstSessionHour                         =               int64(sessionModel_test.SessionFinishTime.Hour())
	sessionRespondModel_test.FirstSessionDuration                     =               int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstSessionMinute                       =               int64(sessionModel_test.SessionFinishTime.Minute())
	sessionRespondModel_test.FirstHalfHourTotalSessionCount                   =       1
	sessionRespondModel_test.FirstHalfHourTotalSessionDuration                =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstHourTotalSessionCount                       =       1
	sessionRespondModel_test.FirstHourTotalSessionDuration                    =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstTwoHourTotalSessionCount                    =       1
	sessionRespondModel_test.FirstTwoHourTotalSessionDuration                 =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstThreeHourTotalSessionCount                  =       1
	sessionRespondModel_test.FirstThreeHourTotalSessionDuration               =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstSixHourTotalSessionCount                    =       1
	sessionRespondModel_test.FirstSixHourTotalSessionDuration                 =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstTwelveHourTotalSessionCount                 =       1
	sessionRespondModel_test.FirstTwelveHourTotalSessionDuration              =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.TotalSessionDay                                  =       1
	sessionRespondModel_test.TotalSessionHour                                 =       1
	sessionRespondModel_test.TotalSessionMinute                               =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.TotalSessionDuration                             =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.TotalSessionCount                                =       1
	sessionRespondModel_test.FirstDayTotalSessionCount                        =       1
	sessionRespondModel_test.FirstDayTotalSessionDuration                     =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.MinSessionDuration                                =      int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.MaxSessionDuration                                =      int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.DailyAvegareSessionCount                          =      1
	sessionRespondModel_test.DailyAverageSessionDuration                       =     float64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.SessionBasedAvegareSessionDuration                =      float64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.DailyAvegareSessionCountMinusFirstDaySessionCount =       0
	sessionRespondModel_test.DailyAvegareSessionDurationMinusFirstDaySessionDuration = 0
	sessionRespondModel_test.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = 0
	sessionRespondModel_test.SessionBasedAvegareSessionDurationMinusLastSessionDuration  = float64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.SundaySessionCount     =  0
	sessionRespondModel_test.MondaySessionCount     =  1
	sessionRespondModel_test.TuesdaySessionCount    =  0
	sessionRespondModel_test.WednesdaySessionCount  =  0
	sessionRespondModel_test.ThursdaySessionCount   =  0
	sessionRespondModel_test.FridaySessionCount     =  0
	sessionRespondModel_test.SaturdaySessionCount   =  0
	sessionRespondModel_test.AmSessionCount        =  0
	sessionRespondModel_test.PmSessionCount         =  1
	sessionRespondModel_test.Session0To5HourCount   =  0
	sessionRespondModel_test.Session6To11HourCount  =  0
	sessionRespondModel_test.Session12To17HourCount  =  0
	sessionRespondModel_test.Session18To23HourCount  =  1

	var TotalSessionHour int64 = ((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour) - ((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)
	var TotalSessionMinute int64 = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24+sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24+sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	var TotalSessionDay = int64(sessionRespondModel_test.FirstSessionYearOfDay-sessionOldModel_test.FirstSessionYearOfDay) + 365*(sessionRespondModel_test.FirstSessionYear-sessionOldModel_test.FirstSessionYear) + 1
	var TotalSessionDuration int64 = sessionRespondModel_test.TotalSessionDuration + sessionOldModel_test.TotalSessionDuration
	var TotalSessionCount int64 = sessionRespondModel_test.TotalSessionCount + sessionOldModel_test.TotalSessionCount
	var SessionBasedAvegareSessionDuration = float64(TotalSessionDuration) / float64(TotalSessionCount)

	var sessionUpdateModel_test = sessionUpdateModel
	sessionUpdateModel_test.ProjectId                             =                  "Test"
	sessionUpdateModel_test.ClientId                              =                  "Test"
	sessionUpdateModel_test.CustomerId                            =                  "Test"
	sessionUpdateModel_test.FirstSessionYearOfDay                 =                  int64(first.YearDay())
	sessionUpdateModel_test.FirstSessionYear                      =                  int64(first.Year())
	sessionUpdateModel_test.FirstSessionWeekDay                   =                  int64(first.Weekday())
	sessionUpdateModel_test.FirstSessionHour                      =                  int64(first.Hour())
	sessionUpdateModel_test.FirstSessionDuration                  =                  int64(sessionModel.SessionTimeMinute)
	sessionUpdateModel_test.FirstSessionMinute                    =                  int64(first.Minute())
	sessionUpdateModel_test.ThirdSessionHour                       =                sessionRespondModel_test.FirstSessionHour
	sessionUpdateModel_test.ThirdSessionDuration                   =                sessionRespondModel_test.FirstSessionDuration
	sessionUpdateModel_test.ThirdSessinMinute                      =                sessionRespondModel_test.FirstSessionMinute
	sessionUpdateModel_test.PenultimateSessionHour                 =                sessionOldModel_test.LastSessionHour
	sessionUpdateModel_test.PenultimateSessionDuration             =                sessionOldModel_test.LastSessionDuration
	sessionUpdateModel_test.PenultimateSessionMinute               =                sessionOldModel_test.LastSessionMinute
	sessionUpdateModel_test.LastSessionYearOfDay                   =                sessionRespondModel_test.FirstSessionYearOfDay
	sessionUpdateModel_test.LastSessionYear                        =                sessionRespondModel_test.FirstSessionYear
	sessionUpdateModel_test.LastSessionHour                        =                sessionRespondModel_test.FirstSessionHour
	sessionUpdateModel_test.LastSessionDuration                    =                sessionRespondModel_test.FirstSessionDuration
	sessionUpdateModel_test.LastSessionMinute                      =                sessionRespondModel_test.FirstSessionMinute
	sessionUpdateModel_test.LastDurationMinusPenultimateDuration   =                sessionRespondModel_test.FirstSessionDuration - sessionOldModel_test.LastSessionDuration
	sessionUpdateModel_test.TotalSessionDay                        =                 TotalSessionDay
	sessionUpdateModel_test.TotalSessionHour                       =                 TotalSessionHour
	sessionUpdateModel_test.TotalSessionMinute                     =                 TotalSessionMinute
	sessionUpdateModel_test.TotalSessionDuration                   =                 sessionRespondModel_test.TotalSessionDuration + sessionOldModel_test.TotalSessionDuration
	sessionUpdateModel_test.TotalSessionCount                      =                 sessionRespondModel_test.TotalSessionCount + sessionOldModel_test.TotalSessionCount
	sessionUpdateModel_test.FirstDayTotalSessionCount              =                 sessionOldModel_test.FirstDayTotalSessionCount
	sessionUpdateModel_test.FirstDayTotalSessionDuration           =                 sessionOldModel_test.FirstDayTotalSessionDuration
	sessionUpdateModel_test.SecondDayTotalSessionCount             =                 sessionRespondModel_test.FirstDayTotalSessionCount
	sessionUpdateModel_test.SecondDayTotalSessionDuration          =                 sessionRespondModel_test.FirstDayTotalSessionDuration
	sessionUpdateModel_test.MinSessionDuration                                   =   int64(19)
	sessionUpdateModel_test.MaxSessionDuration                                   =   int64(33)
	sessionUpdateModel_test.DailyAvegareSessionCount                             =   float64(TotalSessionCount) / float64(TotalSessionDay)
	sessionUpdateModel_test.DailyAverageSessionDuration                          =   float64(TotalSessionDuration) / float64(TotalSessionDay)
	sessionUpdateModel_test.SessionBasedAvegareSessionDuration                   =   float64(TotalSessionDuration) / float64(TotalSessionCount)
	sessionUpdateModel_test.DailyAvegareSessionCountMinusFirstDaySessionCount    =   float64(TotalSessionCount)/float64(TotalSessionDay) - float64(sessionOldModel_test.FirstDayTotalSessionCount)
	sessionUpdateModel_test.DailyAvegareSessionDurationMinusFirstDaySessionDuration = float64(TotalSessionDuration)/float64(TotalSessionDay) - float64(sessionOldModel_test.FirstDayTotalSessionDuration)
	sessionUpdateModel_test.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = SessionBasedAvegareSessionDuration - float64(sessionOldModel_test.FirstSessionDuration)
	sessionUpdateModel_test.SessionBasedAvegareSessionDurationMinusLastSessionDuration =  SessionBasedAvegareSessionDuration - float64(sessionRespondModel_test.FirstSessionDuration)
	sessionUpdateModel_test.SundaySessionCount       = sessionOldModel_test.SundaySessionCount + sessionRespondModel_test.SundaySessionCount
	sessionUpdateModel_test.MondaySessionCount       =  sessionOldModel_test.MondaySessionCount + sessionRespondModel_test.MondaySessionCount
	sessionUpdateModel_test.TuesdaySessionCount      =     sessionOldModel_test.TuesdaySessionCount + sessionRespondModel_test.TuesdaySessionCount
	sessionUpdateModel_test.WednesdaySessionCount    =     sessionOldModel_test.WednesdaySessionCount + sessionRespondModel_test.WednesdaySessionCount
	sessionUpdateModel_test.ThursdaySessionCount     =     sessionOldModel_test.ThursdaySessionCount + sessionRespondModel_test.ThursdaySessionCount
	sessionUpdateModel_test.FridaySessionCount       =     sessionOldModel_test.FridaySessionCount + sessionRespondModel_test.FridaySessionCount
	sessionUpdateModel_test.SaturdaySessionCount     =     sessionOldModel_test.SaturdaySessionCount + sessionRespondModel_test.SaturdaySessionCount
	sessionUpdateModel_test.AmSessionCount          =      sessionOldModel_test.AmSessionCount + sessionRespondModel_test.AmSessionCount
	sessionUpdateModel_test.PmSessionCount          =      sessionOldModel_test.PmSessionCount + sessionRespondModel_test.PmSessionCount
	sessionUpdateModel_test.Session0To5HourCount    =      sessionOldModel_test.Session0To5HourCount + sessionRespondModel_test.Session0To5HourCount
	sessionUpdateModel_test.Session6To11HourCount   =      sessionOldModel_test.Session6To11HourCount + sessionRespondModel_test.Session6To11HourCount
	sessionUpdateModel_test.Session12To17HourCount  =   sessionOldModel_test.Session12To17HourCount + sessionRespondModel_test.Session12To17HourCount
	sessionUpdateModel_test.Session18To23HourCount  =  sessionOldModel_test.Session18To23HourCount + sessionRespondModel_test.Session18To23HourCount
    testSessionDal.On("UpdateGameSessionEveryLoginById", sessionOldModel_test.ClientId, &sessionUpdateModel_test).Return(nil)
	var v, s, m = manager.UpdateGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
	assert.Equal(t, &sessionUpdateModel_test, v)
}

func Test_ConvertRawModelToResponseModel_AddS(t *testing.T) {
	var testSessionDal = new(repository.MockGameSessionEveryLoginDal)
	var testCache = new(service.MockCacheService)
	var testLog = new(Log.MockLogger)
	var json = gojson.GoJsonConstructor()
	IoC.JsonParser = json
	IoC.GameSessionEveryLoginDal = testSessionDal
	IoC.Logger = testLog
	IoC.CacheService = testCache
	var manager = concrete.GameSessionEveryLoginManagerConstructor()
	var sessionOldModel_test = sessionOldModel
	var sessionModel_test = sessionModel
	sessionModel_test.ProjectId  = "Test"
	sessionModel_test.ClientId   = "Test"
	sessionModel_test.CustomerId = "Test"
	sessionModel_test.SessionStartTime = time.Date(
		2021, 11, 7, 22, 34, 36, 651387237, time.UTC)
	sessionModel_test.SessionFinishTime = time.Date(
		2021, 11, 8, 18, 34, 36, 651387237, time.UTC)
	sessionModel_test.SessionTimeMinute = 33
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.ProjectId                                =               "Test"
	sessionRespondModel_test.ClientId                                 =               "Test"
	sessionRespondModel_test.CustomerId                               =               "Test"
	sessionRespondModel_test.FirstSessionYearOfDay                    =               int64(sessionModel_test.SessionFinishTime.YearDay())
	sessionRespondModel_test.FirstSessionYear                         =               int64(sessionModel_test.SessionFinishTime.Year())
	sessionRespondModel_test.FirstSessionWeekDay                      =               int64(sessionModel_test.SessionFinishTime.Weekday())
	sessionRespondModel_test.FirstSessionHour                         =               int64(sessionModel_test.SessionFinishTime.Hour())
	sessionRespondModel_test.FirstSessionDuration                     =               int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstSessionMinute                       =               int64(sessionModel_test.SessionFinishTime.Minute())
	sessionRespondModel_test.FirstHalfHourTotalSessionCount                   =       1
	sessionRespondModel_test.FirstHalfHourTotalSessionDuration                =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstHourTotalSessionCount                       =       1
	sessionRespondModel_test.FirstHourTotalSessionDuration                    =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstTwoHourTotalSessionCount                    =       1
	sessionRespondModel_test.FirstTwoHourTotalSessionDuration                 =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstThreeHourTotalSessionCount                  =       1
	sessionRespondModel_test.FirstThreeHourTotalSessionDuration               =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstSixHourTotalSessionCount                    =       1
	sessionRespondModel_test.FirstSixHourTotalSessionDuration                 =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.FirstTwelveHourTotalSessionCount                 =       1
	sessionRespondModel_test.FirstTwelveHourTotalSessionDuration              =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.TotalSessionDay                                  =       1
	sessionRespondModel_test.TotalSessionHour                                 =       1
	sessionRespondModel_test.TotalSessionMinute                               =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.TotalSessionDuration                             =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.TotalSessionCount                                =       1
	sessionRespondModel_test.FirstDayTotalSessionCount                        =       1
	sessionRespondModel_test.FirstDayTotalSessionDuration                     =       int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.MinSessionDuration                                =      int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.MaxSessionDuration                                =      int64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.DailyAvegareSessionCount                          =      1
	sessionRespondModel_test.DailyAverageSessionDuration                       =     float64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.SessionBasedAvegareSessionDuration                =      float64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.DailyAvegareSessionCountMinusFirstDaySessionCount =       0
	sessionRespondModel_test.DailyAvegareSessionDurationMinusFirstDaySessionDuration = 0
	sessionRespondModel_test.SessionBasedAvegareSessionDurationMinusFirstSessionDuration = 0
	sessionRespondModel_test.SessionBasedAvegareSessionDurationMinusLastSessionDuration  = float64(sessionModel_test.SessionTimeMinute)
	sessionRespondModel_test.SundaySessionCount     =  0
	sessionRespondModel_test.MondaySessionCount     =  1
	sessionRespondModel_test.TuesdaySessionCount    =  0
	sessionRespondModel_test.WednesdaySessionCount  =  0
	sessionRespondModel_test.ThursdaySessionCount   =  0
	sessionRespondModel_test.FridaySessionCount     =  0
	sessionRespondModel_test.SaturdaySessionCount   =  0
	sessionRespondModel_test.AmSessionCount        =  0
	sessionRespondModel_test.PmSessionCount         =  1
	sessionRespondModel_test.Session0To5HourCount   =  0
	sessionRespondModel_test.Session6To11HourCount  =  0
	sessionRespondModel_test.Session12To17HourCount  =  0
	sessionRespondModel_test.Session18To23HourCount  =  1
	gameByte, _ := json.EncodeJson(sessionModel_test)
	testSessionDal.On("GetGameSessionEveryLoginById", sessionRespondModel_test.ClientId).Return(&sessionOldModel_test,
		errors.New("null data error"))
	testSessionDal.On("Add", &sessionRespondModel_test).Return(nil)
	var v, s, m = manager.ConvertRawModelToResponseModel(gameByte)
	var value, success = v.(model.BuyingEventRespondModel)
	if success == true {
		assert.Equal(t, &sessionRespondModel_test, value)
	}
	assert.Equal(t, true, success)
	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
	assert.Equal(t, &sessionRespondModel, v)
}


func Test_CalculateSecondDayTotalSessionCountAndDuration_In24To48(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstDayTotalSessionCount = 6
	sessionOldModel_test.FirstDayTotalSessionDuration = 10
	sessionOldModel_test.FirstSessionMinute = 49
	sessionOldModel_test.SecondDayTotalSessionCount = 25
	sessionOldModel_test.SecondDayTotalSessionDuration = 33
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 15
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 35
	sessionRespondModel_test.FirstSessionMinute = 22
	//var TotalSessionHour = ((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour) - ((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)
	var TotalSessionMinute = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateSecondDayTotalSessionCountAndDuration(&sessionRespondModel_test, &sessionOldModel_test, TotalSessionMinute)
	var ExpCount int64 = 26
	var ExpDuration int64 = 68
	assert.Equal(t, ExpCount , sessionOldModel_test.SecondDayTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.SecondDayTotalSessionDuration)
}

func Test_CalculateSecondDayTotalSessionCountAndDuration_Out24To48(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstDayTotalSessionCount = 6
	sessionOldModel_test.FirstDayTotalSessionDuration = 10
	sessionOldModel_test.FirstSessionMinute = 49
	sessionOldModel_test.SecondDayTotalSessionCount = 25
	sessionOldModel_test.SecondDayTotalSessionDuration = 33
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 302
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 15
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 35
	sessionRespondModel_test.FirstSessionMinute = 22
	var TotalSessionHour_session = ((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour) - ((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)
	concrete.CalculateSecondDayTotalSessionCountAndDuration(&sessionRespondModel_test, &sessionOldModel_test, TotalSessionHour_session)
	var ExpCount int64 = 25
	var ExpDuration int64 = 33
	assert.Equal(t, ExpCount , sessionOldModel_test.SecondDayTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.SecondDayTotalSessionDuration)
}

func Test_CalculateFirstDayTotalSessionCountAndDuration_In00To24(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstDayTotalSessionCount = 6
	sessionOldModel_test.FirstDayTotalSessionDuration = 10
	sessionOldModel_test.FirstSessionMinute = 51
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 300
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 23
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 78
	sessionRespondModel_test.FirstSessionMinute = 18
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstDayTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 6 + 1
	var ExpDuration int64 = 78 + 10
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstDayTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstDayTotalSessionDuration)
}

func Test_CalculateFirstDayTotalSessionCountAndDuration_Out00To24(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstDayTotalSessionCount = 6
	sessionOldModel_test.FirstDayTotalSessionDuration = 10
	sessionOldModel_test.FirstSessionMinute = 51
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 15
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 78
	sessionRespondModel_test.FirstSessionMinute = 18
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstDayTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 6
	var ExpDuration int64 = 10
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstDayTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstDayTotalSessionDuration)
}

func Test_CalculateFirstTwelveHourTotalSessionCountAndDuration_In00To12(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstTwelveHourTotalSessionCount = 2
	sessionOldModel_test.FirstTwelveHourTotalSessionDuration = 17
	sessionOldModel_test.FirstSessionMinute = 51
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 02
	sessionRespondModel_test.FirstSessionMinute = 14
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 18
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstTwelveHourTotalSessionCountAndDuration(&sessionRespondModel_test, &sessionOldModel_test, TotalSessionMinute_session)
	var ExpCount int64 = 2 + 1
	var ExpDuration int64 = 17 + 18
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstTwelveHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstTwelveHourTotalSessionDuration)
}

func Test_CalculateFirstTwelveHourTotalSessionCountAndDuration_Out00To12(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstTwelveHourTotalSessionCount = 2
	sessionOldModel_test.FirstTwelveHourTotalSessionDuration = 17
	sessionOldModel_test.FirstSessionMinute = 51
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 02
	sessionRespondModel_test.FirstSessionMinute = 14
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 45
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstTwelveHourTotalSessionCountAndDuration(&sessionRespondModel_test, &sessionOldModel_test, TotalSessionMinute_session)
	var ExpCount int64 = 2
	var ExpDuration int64 = 17
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstTwelveHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstTwelveHourTotalSessionDuration)
}

func Test_CalculateFirstSixHourTotalSessionCountAndDuration_In00To06(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstSixHourTotalSessionCount = 2
	sessionOldModel_test.FirstSixHourTotalSessionDuration = 17
	sessionOldModel_test.FirstSessionMinute = 51
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 20
	sessionRespondModel_test.FirstSessionMinute = 13
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 11
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstSixHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 2+1
	var ExpDuration int64 = 17+11
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstSixHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstSixHourTotalSessionDuration)
}

func Test_CalculateFirstThreeHourTotalSessionCountAndDuration_In00To03(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstThreeHourTotalSessionCount = 5
	sessionOldModel_test.FirstThreeHourTotalSessionDuration = 83
	sessionOldModel_test.FirstSessionMinute = 31
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 17
	sessionRespondModel_test.FirstSessionMinute = 31
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 7
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstThreeHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 5+1
	var ExpDuration int64 = 83+7
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstThreeHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstThreeHourTotalSessionDuration)
}

func Test_CalculateFirstThreeHourTotalSessionCountAndDuration_Out00To03(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstThreeHourTotalSessionCount = 5
	sessionOldModel_test.FirstThreeHourTotalSessionDuration = 83
	sessionOldModel_test.FirstSessionMinute = 31
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 17
	sessionRespondModel_test.FirstSessionMinute = 33
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 7
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstThreeHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 5
	var ExpDuration int64 = 83
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstThreeHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstThreeHourTotalSessionDuration)
}

func Test_CalculateFirstTwoHourTotalSessionCountAndDuration_In00To02(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstTwoHourTotalSessionCount = 12
	sessionOldModel_test.FirstTwoHourTotalSessionDuration = 83
	sessionOldModel_test.FirstSessionMinute = 31
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 15
	sessionRespondModel_test.FirstSessionMinute = 58
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 6
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstTwoHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 12+1
	var ExpDuration int64 = 83+6
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstTwoHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstTwoHourTotalSessionDuration)
}

func Test_CalculateFirstTwoHourTotalSessionCountAndDuration_Out00To02(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 14
	sessionOldModel_test.FirstTwoHourTotalSessionCount = 12
	sessionOldModel_test.FirstTwoHourTotalSessionDuration = 83
	sessionOldModel_test.FirstSessionMinute = 31
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 16
	sessionRespondModel_test.FirstSessionMinute = 32
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 6
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstTwoHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 12
	var ExpDuration int64 = 83
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstTwoHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstTwoHourTotalSessionDuration)
}

func Test_CalculateFirstHourTotalSessionCountAndDuration_In00To01(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 15
	sessionOldModel_test.FirstHourTotalSessionCount = 3
	sessionOldModel_test.FirstHourTotalSessionDuration = 5
	sessionOldModel_test.FirstSessionMinute = 46
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 16
	sessionRespondModel_test.FirstSessionMinute = 46
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 9
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 3+1
	var ExpDuration int64 = 5+9
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstHourTotalSessionDuration)
}

func Test_CalculateFirstHourTotalSessionCountAndDuration_Out00To01(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 15
	sessionOldModel_test.FirstHourTotalSessionCount = 1
	sessionOldModel_test.FirstHourTotalSessionDuration = 8
	sessionOldModel_test.FirstSessionMinute = 27
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 16
	sessionRespondModel_test.FirstSessionMinute = 28
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 46
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 1
	var ExpDuration int64 = 8
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstTwoHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstTwoHourTotalSessionDuration)
}

func Test_CalculateFirstHalfHourTotalSessionCountAndSessionDuration_In00To30(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 15
	sessionOldModel_test.FirstHalfHourTotalSessionCount = 0
	sessionOldModel_test.FirstHalfHourTotalSessionDuration = 0
	sessionOldModel_test.FirstSessionMinute = 55
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 16
	sessionRespondModel_test.FirstSessionMinute = 25
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 53
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstHalfHourTotalSessionCountAndSessionDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 0+1
	var ExpDuration int64 = 0+53
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstHalfHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstHalfHourTotalSessionDuration)
}

func Test_CalculateFirstHalfHourTotalSessionCountAndSessionDuration_Out00To30(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.FirstSessionYearOfDay = 300
	sessionOldModel_test.FirstSessionYear = 2021
	sessionOldModel_test.FirstSessionHour = 15
	sessionOldModel_test.FirstHalfHourTotalSessionCount = 0
	sessionOldModel_test.FirstHalfHourTotalSessionDuration = 0
	sessionOldModel_test.FirstSessionMinute = 55
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.FirstSessionYearOfDay = 301
	sessionRespondModel_test.FirstSessionYear = 2021
	sessionRespondModel_test.FirstSessionHour = 16
	sessionRespondModel_test.FirstSessionMinute = 26
	sessionRespondModel_test.FirstDayTotalSessionCount = 1
	sessionRespondModel_test.FirstDayTotalSessionDuration = 53
	var TotalSessionMinute_session = (((sessionRespondModel_test.FirstSessionYearOfDay+365*sessionRespondModel_test.FirstSessionYear)*24 + sessionRespondModel_test.FirstSessionHour)*60 + sessionRespondModel_test.FirstSessionMinute) - (((sessionOldModel_test.FirstSessionYearOfDay+365*sessionOldModel_test.FirstSessionYear)*24 + sessionOldModel_test.FirstSessionHour)*60 + sessionOldModel_test.FirstSessionMinute)
	concrete.CalculateFirstHalfHourTotalSessionCountAndSessionDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute_session)
	var ExpCount int64 = 0
	var ExpDuration int64 = 0
	assert.Equal(t, ExpCount , sessionOldModel_test.FirstHalfHourTotalSessionCount)
	assert.Equal(t, ExpDuration, sessionOldModel_test.FirstHalfHourTotalSessionDuration)
}

func Test_CalculateSecondGameSession_EqualTotalSessionCount2(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 2
	sessionOldModel_test.SecondSessionHour      = 0
	sessionOldModel_test.SecondSessionDuration = 0
	sessionOldModel_test.SecondSessionMinute = 0
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 22
	sessionRespondModel_test.FirstSessionDuration = 22
	sessionRespondModel_test.FirstSessionMinute = 22
	concrete.CalculateSecondGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 22,22,22
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateSecondGameSession_NotEqualTotalSessionCount2(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 3
	sessionOldModel_test.SecondSessionHour      = 0
	sessionOldModel_test.SecondSessionDuration = 0
	sessionOldModel_test.SecondSessionMinute = 0
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 23
	sessionRespondModel_test.FirstSessionDuration = 48
	sessionRespondModel_test.FirstSessionMinute = 39
	concrete.CalculateSecondGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 0,0,0
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateThirdGameSession_EqualTotalSessionCount3(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 3
	sessionOldModel_test.ThirdSessionHour      = 0
	sessionOldModel_test.ThirdSessionDuration = 0
	sessionOldModel_test.ThirdSessinMinute = 0
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 23
	sessionRespondModel_test.FirstSessionDuration = 48
	sessionRespondModel_test.FirstSessionMinute = 39
	concrete.CalculateThirdGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 23,48,39
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateThirdGameSession_NotEqualTotalSessionCount3(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 4
	sessionOldModel_test.ThirdSessionHour      = 5
	sessionOldModel_test.ThirdSessionDuration = 8
	sessionOldModel_test.ThirdSessinMinute = 10
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 23
	sessionRespondModel_test.FirstSessionDuration = 48
	sessionRespondModel_test.FirstSessionMinute = 39
	concrete.CalculateThirdGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 5,8,10
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateFourthGameSession_EqualTotalSessionCount4(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 4
	sessionOldModel_test.FourthSessionHour      = 1
	sessionOldModel_test.FourthSessionDuration = 2
	sessionOldModel_test.FourthSessinMinute = 9
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 11
	sessionRespondModel_test.FirstSessionDuration = 48
	sessionRespondModel_test.FirstSessionMinute = 39
	concrete.CalculateFourthGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 11,48,39
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateFourthGameSession_NotEqualTotalSessionCount4(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 5
	sessionOldModel_test.FourthSessionHour      = 19
	sessionOldModel_test.FourthSessionDuration = 8
	sessionOldModel_test.FourthSessinMinute = 7
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 23
	sessionRespondModel_test.FirstSessionDuration = 48
	sessionRespondModel_test.FirstSessionMinute = 39
	concrete.CalculateFourthGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 19,8,10
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateFifthGameSession_EqualTotalSessionCount5(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 5
	sessionOldModel_test.FifthSessionHour      = 13
	sessionOldModel_test.FifthSessionDuration = 28
	sessionOldModel_test.FifthSessinMinute = 79
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 14
	sessionRespondModel_test.FirstSessionDuration = 48
	sessionRespondModel_test.FirstSessionMinute = 55
	concrete.CalculateFifthGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 14,48,55
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateFifthGameSession_NotEqualTotalSessionCount5(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.TotalSessionCount = 4
	sessionOldModel_test.FifthSessionHour      = 6
	sessionOldModel_test.FifthSessionDuration = 8
	sessionOldModel_test.FifthSessinMinute = 7
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.TotalSessionCount = 1
	sessionRespondModel_test.FirstSessionHour      = 23
	sessionRespondModel_test.FirstSessionDuration = 48
	sessionRespondModel_test.FirstSessionMinute = 39
	concrete.CalculateFifthGameSession(&sessionRespondModel_test, &sessionOldModel_test)
	var Exphour, Expduration, Expminute int64 = 6,8,7
	assert.Equal(t, Exphour, sessionOldModel_test.SecondSessionHour )
	assert.Equal(t, Expduration , sessionOldModel_test.SecondSessionDuration)
	assert.Equal(t, Expminute, sessionOldModel_test.SecondSessionMinute)
}

func Test_CalculateMaxDuration_ChangeMaxValue(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.MaxSessionDuration = 4
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.MaxSessionDuration = 6
	concrete.CalculateMaxDuration(&sessionRespondModel, &sessionOldModel)
	var ExpDuration int64 = 6
	assert.Equal(t, ExpDuration, sessionOldModel_test.MaxSessionDuration)
}

func Test_CalculateMaxDuration_NotChangeMaxValue(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.MaxSessionDuration = 4
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.MaxSessionDuration = 3
	concrete.CalculateMaxDuration(&sessionRespondModel_test, &sessionOldModel_test)
	var ExpDuration int64 = 4
	assert.Equal(t, ExpDuration, sessionOldModel_test.MaxSessionDuration)
}

func Test_CalculateMinDuration_ChangedMinValue(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.MinSessionDuration = 7
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.MinSessionDuration = 5
	concrete.CalculateMinDuration(&sessionRespondModel, &sessionOldModel)
	var ExpDuration int64 = 5
	assert.Equal(t, ExpDuration, sessionOldModel_test.MinSessionDuration)
}



func Test_CalculateMinDuration_NotChangedMinValue(t *testing.T) {
	var sessionOldModel_test = sessionOldModel
	sessionOldModel_test.MinSessionDuration = 5
	var sessionRespondModel_test = sessionRespondModel
	sessionRespondModel_test.MinSessionDuration = 9
	concrete.CalculateMinDuration(&sessionRespondModel, &sessionOldModel)
	var ExpDuration int64 = 5
	assert.Equal(t, ExpDuration, sessionOldModel_test.MinSessionDuration)
}


func Test_DetermineGameSessionHour_00To05(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 11, 7, 5, 34, 36, 651387237, time.UTC)
	var hour int64 = int64(sessionModel_test.SessionFinishTime.Hour())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionHour(&sessionRespondModel_test, hour)
	assert.Equal(t, int64(1), sessionRespondModel_test.Session0To5HourCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.Session6To11HourCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.Session12To17HourCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.Session18To23HourCount)
}

func Test_DetermineGameSessionHour_06To11(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 11, 7, 11, 34, 36, 651387237, time.UTC)
	var hour int64 = int64(sessionModel_test.SessionFinishTime.Hour())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionHour(&sessionRespondModel_test, hour)

	assert.Equal(t, int64(0), sessionRespondModel_test.Session0To5HourCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.Session6To11HourCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.Session12To17HourCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.Session18To23HourCount)
}

func Test_DetermineGameSessionHour_12To17(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 11, 7, 15, 34, 36, 651387237, time.UTC)
	var hour int64 = int64(sessionModel_test.SessionFinishTime.Hour())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionHour(&sessionRespondModel_test, hour)

	assert.Equal(t, int64(0), sessionRespondModel_test.Session0To5HourCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.Session6To11HourCount )
	assert.Equal(t, int64(1), sessionRespondModel_test.Session12To17HourCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.Session18To23HourCount)
}

func Test_DetermineGameSessionHour_18To23(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 11, 7, 23, 34, 36, 651387237, time.UTC)
	var hour int64 = int64(sessionModel_test.SessionFinishTime.Hour())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionHour(&sessionRespondModel_test, hour)

	assert.Equal(t, int64(0), sessionRespondModel_test.Session0To5HourCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.Session6To11HourCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.Session12To17HourCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.Session18To23HourCount)
}


func Test_DetermineGameSessionDay_Sunday(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 5, 23, 34, 36, 651387237, time.UTC)
	var day int64 = int64(sessionModel_test.SessionFinishTime.Weekday())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionDay(&sessionRespondModel_test, day)
	assert.Equal(t, int64(1), sessionRespondModel_test.SundaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.MondaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.TuesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.WednesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.ThursdaySessionCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.FridaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.SaturdaySessionCount )
}

func Test_DetermineGameSessionDay_Monday(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 6, 23, 34, 36, 651387237, time.UTC)
	var day int64 = int64(sessionModel_test.SessionFinishTime.Weekday())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionDay(&sessionRespondModel_test, day)
	assert.Equal(t, int64(0), sessionRespondModel_test.SundaySessionCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.MondaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.TuesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.WednesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.ThursdaySessionCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.FridaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.SaturdaySessionCount )
}

func Test_DetermineGameSessionDay_Tuesday(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 7, 23, 34, 36, 651387237, time.UTC)
	var day int64 = int64(sessionModel_test.SessionFinishTime.Weekday())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionDay(&sessionRespondModel_test, day)
	assert.Equal(t, int64(0), sessionRespondModel_test.SundaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.MondaySessionCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.TuesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.WednesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.ThursdaySessionCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.FridaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.SaturdaySessionCount )
}

func Test_DetermineGameSessionDay_Wednesday(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 8, 23, 34, 36, 651387237, time.UTC)
	var day int64 = int64(sessionModel_test.SessionFinishTime.Weekday())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionDay(&sessionRespondModel_test, day)
	assert.Equal(t, int64(0), sessionRespondModel_test.SundaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.MondaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.TuesdaySessionCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.WednesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.ThursdaySessionCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.FridaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.SaturdaySessionCount )
}

func Test_DetermineGameSessionDay_Thursday(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 9, 23, 34, 36, 651387237, time.UTC)
	var day int64 = int64(sessionModel_test.SessionFinishTime.Weekday())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionDay(&sessionRespondModel_test, day)
	assert.Equal(t, int64(0), sessionRespondModel_test.SundaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.MondaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.TuesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.WednesdaySessionCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.ThursdaySessionCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.FridaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.SaturdaySessionCount )
}

func Test_DetermineGameSessionDay_Friday(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 10, 23, 34, 36, 651387237, time.UTC)
	var day int64 = int64(sessionModel_test.SessionFinishTime.Weekday())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionDay(&sessionRespondModel_test, day)
	assert.Equal(t, int64(0), sessionRespondModel_test.SundaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.MondaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.TuesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.WednesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.ThursdaySessionCount )
	assert.Equal(t, int64(1), sessionRespondModel_test.FridaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.SaturdaySessionCount )
}

func Test_DetermineGameSessionDay_Saturday(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 11, 23, 34, 36, 651387237, time.UTC)
	var day int64 = int64(sessionModel_test.SessionFinishTime.Weekday())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionDay(&sessionRespondModel_test, day)
	assert.Equal(t, int64(0), sessionRespondModel_test.SundaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.MondaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.TuesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.WednesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.ThursdaySessionCount )
	assert.Equal(t, int64(0), sessionRespondModel_test.FridaySessionCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.SaturdaySessionCount )
}

func Test_DetermineGameSessionAmPm_Am(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 11, 12, 34, 36, 651387237, time.UTC)
	var hour int64 = int64(sessionModel_test.SessionFinishTime.Hour())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionAmPm(&sessionRespondModel_test, hour)

	assert.Equal(t, int64(1), sessionRespondModel_test.AmSessionCount)
	assert.Equal(t, int64(0), sessionRespondModel_test.PmSessionCount)

}

func Test_DetermineGameSessionAmPm_Pm(t *testing.T){
	var sessionModel_test = sessionModel
	sessionModel_test.SessionFinishTime =  time.Date(
		2021, 12, 11, 13, 34, 36, 651387237, time.UTC)
	var hour int64 = int64(sessionModel_test.SessionFinishTime.Hour())
	var sessionRespondModel_test = sessionRespondModel
	concrete.DetermineGameSessionAmPm(&sessionRespondModel_test, hour)

	assert.Equal(t, int64(0), sessionRespondModel_test.AmSessionCount)
	assert.Equal(t, int64(1), sessionRespondModel_test.PmSessionCount)

}