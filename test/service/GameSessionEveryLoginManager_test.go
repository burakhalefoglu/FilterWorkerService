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

var sessionModel = model.GameSessionEveryLoginModel{
	ProjectId:  "Test",
	ClientId:   "Test",
	CustomerId: "Test",
	SessionStartTime: time.Date(
		2021, 11, 7, 22, 34, 36, 651387237, time.UTC),
	SessionFinishTime: time.Date(
		2021, 11, 8, 18, 34, 36, 651387237, time.UTC),
	SessionTimeMinute: 33,
}

var sessionRespondModel = model.GameSessionEveryLoginRespondModel{
	ProjectId:                                               "Test",
	ClientId:                                                "Test",
	CustomerId:                                              "Test",
	FirstSessionYearOfDay:                                   int64(sessionModel.SessionFinishTime.YearDay()),
	FirstSessionYear:                                        int64(sessionModel.SessionFinishTime.Year()),
	FirstSessionWeekDay:                                     int64(sessionModel.SessionFinishTime.Weekday()),
	FirstSessionHour:                                        int64(sessionModel.SessionFinishTime.Hour()),
	FirstSessionDuration:                                    int64(sessionModel.SessionTimeMinute),
	FirstSessionMinute:                                      int64(sessionModel.SessionFinishTime.Minute()),
	SecondSessionHour:                                       0,
	SecondSessionDuration:                                   0,
	SecondSessionMinute:                                     0,
	ThirdSessionHour:                                        0,
	ThirdSessionDuration:                                    0,
	ThirdSessinMinute:                                       0,
	PenultimateSessionHour:                                  0,
	PenultimateSessionDuration:                              0,
	PenultimateSessionMinute:                                0,
	LastSessionYearOfDay:                                    0,
	LastSessionYear:                                         0,
	LastSessionHour:                                         0,
	LastSessionDuration:                                     0,
	LastSessionMinute:                                       0,
	LastDurationMinusPenultimateDuration:                    0,
	FirstHalfHourTotalSessionCount:                          1,
	FirstHalfHourTotalSessionDuration:                       int64(sessionModel.SessionTimeMinute),
	FirstHourTotalSessionCount:                              1,
	FirstHourTotalSessionDuration:                           int64(sessionModel.SessionTimeMinute),
	FirstTwoHourTotalSessionCount:                           1,
	FirstTwoHourTotalSessionDuration:                        int64(sessionModel.SessionTimeMinute),
	FirstThreeHourTotalSessionCount:                         1,
	FirstThreeHourTotalSessionDuration:                      int64(sessionModel.SessionTimeMinute),
	FirstSixHourTotalSessionCount:                           1,
	FirstSixHourTotalSessionDuration:                        int64(sessionModel.SessionTimeMinute),
	FirstTwelveHourTotalSessionCount:                        1,
	FirstTwelveHourTotalSessionDuration:                     int64(sessionModel.SessionTimeMinute),
	TotalSessionDay:                                         1,
	TotalSessionHour:                                        1,
	TotalSessionMinute:                                      int64(sessionModel.SessionTimeMinute),
	TotalSessionDuration:                                    int64(sessionModel.SessionTimeMinute),
	TotalSessionCount:                                       1,
	FirstDayTotalSessionCount:                               1,
	FirstDayTotalSessionDuration:                            int64(sessionModel.SessionTimeMinute),
	SecondDayTotalSessionCount:                              0,
	SecondDayTotalSessionDuration:                           0,
	ThirdDayTotalSessionCount:                               0,
	ThirdDayTotalSessionDuration:                            0,
	FourthDayTotalSessionCount:                              0,
	FourthDayTotalSessionDuration:                           0,
	FifthDayTotalSessionCount:                               0,
	FifthDayTotalSessionDuration:                            0,
	SixthDayTotalSessionCount:                               0,
	SixthDayTotalSessionDuration:                            0,
	SeventhDayTotalSessionCount:                             0,
	SeventhDayTotalSessionDuration:                          0,
	MinSessionDuration:                                      int64(sessionModel.SessionTimeMinute),
	MaxSessionDuration:                                      int64(sessionModel.SessionTimeMinute),
	DailyAvegareSessionCount:                                1,
	DailyAverageSessionDuration:                             float64(sessionModel.SessionTimeMinute),
	SessionBasedAvegareSessionDuration:                      float64(sessionModel.SessionTimeMinute),
	DailyAvegareSessionCountMinusFirstDaySessionCount:       0,
	DailyAvegareSessionDurationMinusFirstDaySessionDuration: 0,
	SessionBasedAvegareSessionDurationMinusFirstSessionDuration: 0,
	SessionBasedAvegareSessionDurationMinusLastSessionDuration:  float64(sessionModel.SessionTimeMinute),
	SundaySessionCount:     0,
	MondaySessionCount:     1,
	TuesdaySessionCount:    0,
	WednesdaySessionCount:  0,
	ThursdaySessionCount:   0,
	FridaySessionCount:     0,
	SaturdaySessionCount:   0,
	AmSessionCount:         0,
	PmSessionCount:         1,
	Session0To5HourCount:   0,
	Session6To11HourCount:  0,
	Session12To17HourCount: 0,
	Session18To23HourCount: 1,
}

var first = time.Date(
	2021, 11, 6, 18, 34, 58, 651387237, time.UTC)

var first2 = time.Date(
	2021, 11, 7, 18, 34, 58, 651387237, time.UTC)

var sessionOldModel = model.GameSessionEveryLoginRespondModel{
	ProjectId:                                               "Test",
	ClientId:                                                "Test",
	CustomerId:                                              "Test",
	FirstSessionYearOfDay:                                   int64(first.YearDay()),
	FirstSessionYear:                                        int64(first.Year()),
	FirstSessionWeekDay:                                     int64(first.Weekday()),
	FirstSessionHour:                                        int64(first.Hour()),
	FirstSessionDuration:                                    int64(sessionModel.SessionTimeMinute),
	FirstSessionMinute:                                      int64(first.Minute()),
	SecondSessionHour:                                       0,
	SecondSessionDuration:                                   0,
	SecondSessionMinute:                                     0,
	ThirdSessionHour:                                        0,
	ThirdSessionDuration:                                    0,
	ThirdSessinMinute:                                       0,
	PenultimateSessionHour:                                  0,
	PenultimateSessionDuration:                              0,
	PenultimateSessionMinute:                                0,
	LastSessionYearOfDay:                                    int64(first2.YearDay()),
	LastSessionYear:                                         int64(first2.Year()),
	LastSessionHour:                                         int64(first2.Hour()),
	LastSessionDuration:                                     int64(sessionModel.SessionTimeMinute),
	LastSessionMinute:                                       int64(first2.Minute()),
	LastDurationMinusPenultimateDuration:                    0,
	FirstHalfHourTotalSessionCount:                          0,
	FirstHalfHourTotalSessionDuration:                       0,
	FirstHourTotalSessionCount:                              0,
	FirstHourTotalSessionDuration:                           0,
	FirstTwoHourTotalSessionCount:                           0,
	FirstTwoHourTotalSessionDuration:                        0,
	FirstThreeHourTotalSessionCount:                         0,
	FirstThreeHourTotalSessionDuration:                      0,
	FirstSixHourTotalSessionCount:                           0,
	FirstSixHourTotalSessionDuration:                        0,
	FirstTwelveHourTotalSessionCount:                        0,
	FirstTwelveHourTotalSessionDuration:                     0,
	TotalSessionDay:                                         1,
	TotalSessionHour:                                        24,
	TotalSessionMinute:                                      1440,
	TotalSessionDuration:                                    18,
	TotalSessionCount:                                       2,
	FirstDayTotalSessionCount:                               3,
	FirstDayTotalSessionDuration:                            18,
	SecondDayTotalSessionCount:                              0,
	SecondDayTotalSessionDuration:                           0,
	ThirdDayTotalSessionCount:                               0,
	ThirdDayTotalSessionDuration:                            0,
	FourthDayTotalSessionCount:                              0,
	FourthDayTotalSessionDuration:                           0,
	FifthDayTotalSessionCount:                               0,
	FifthDayTotalSessionDuration:                            0,
	SixthDayTotalSessionCount:                               0,
	SixthDayTotalSessionDuration:                            0,
	SeventhDayTotalSessionCount:                             0,
	SeventhDayTotalSessionDuration:                          0,
	MinSessionDuration:                                      19,
	MaxSessionDuration:                                      29,
	DailyAvegareSessionCount:                                1,
	DailyAverageSessionDuration:                             float64(sessionModel.SessionTimeMinute),
	SessionBasedAvegareSessionDuration:                      float64(sessionModel.SessionTimeMinute),
	DailyAvegareSessionCountMinusFirstDaySessionCount:       0,
	DailyAvegareSessionDurationMinusFirstDaySessionDuration: 0,
	SessionBasedAvegareSessionDurationMinusFirstSessionDuration: 0,
	SessionBasedAvegareSessionDurationMinusLastSessionDuration:  float64(sessionModel.SessionTimeMinute),
	SundaySessionCount:     2,
	MondaySessionCount:     9,
	TuesdaySessionCount:    8,
	WednesdaySessionCount:  0,
	ThursdaySessionCount:   0,
	FridaySessionCount:     1,
	SaturdaySessionCount:   0,
	AmSessionCount:         0,
	PmSessionCount:         1,
	Session0To5HourCount:   5,
	Session6To11HourCount:  0,
	Session12To17HourCount: 1,
	Session18To23HourCount: 87,
}

var TotalSessionHour int64 = ((sessionRespondModel.FirstSessionYearOfDay+365*sessionRespondModel.FirstSessionYear)*24 + sessionRespondModel.FirstSessionHour) - ((sessionOldModel.FirstSessionYearOfDay+365*sessionOldModel.FirstSessionYear)*24 + sessionOldModel.FirstSessionHour)
var TotalSessionMinute int64 = (((sessionRespondModel.FirstSessionYearOfDay+365*sessionRespondModel.FirstSessionYear)*24+sessionRespondModel.FirstSessionHour)*60 + sessionRespondModel.FirstSessionMinute) - (((sessionOldModel.FirstSessionYearOfDay+365*sessionOldModel.FirstSessionYear)*24+sessionOldModel.FirstSessionHour)*60 + sessionOldModel.FirstSessionMinute)
var TotalSessionDay = int64(sessionRespondModel.FirstSessionYearOfDay-sessionOldModel.FirstSessionYearOfDay) + 365*(sessionRespondModel.FirstSessionYear-sessionOldModel.FirstSessionYear) + 1
var TotalSessionDuration int64 = sessionRespondModel.TotalSessionDuration + sessionOldModel.TotalSessionDuration
var TotalSessionCount int64 = sessionRespondModel.TotalSessionCount + sessionOldModel.TotalSessionCount
var SessionBasedAvegareSessionDuration = float64(TotalSessionDuration) / float64(TotalSessionCount)

var sessionUpdateModel = model.GameSessionEveryLoginRespondModel{
	ProjectId:                                               "Test",
	ClientId:                                                "Test",
	CustomerId:                                              "Test",
	FirstSessionYearOfDay:                                   int64(first.YearDay()),
	FirstSessionYear:                                        int64(first.Year()),
	FirstSessionWeekDay:                                     int64(first.Weekday()),
	FirstSessionHour:                                        int64(first.Hour()),
	FirstSessionDuration:                                    int64(sessionModel.SessionTimeMinute),
	FirstSessionMinute:                                      int64(first.Minute()),
	SecondSessionHour:                                       0,
	SecondSessionDuration:                                   0,
	SecondSessionMinute:                                     0,
	ThirdSessionHour:                                        sessionRespondModel.FirstSessionHour,
	ThirdSessionDuration:                                    sessionRespondModel.FirstSessionDuration,
	ThirdSessinMinute:                                       sessionRespondModel.FirstSessionMinute,
	PenultimateSessionHour:                                  sessionOldModel.LastSessionHour,
	PenultimateSessionDuration:                              sessionOldModel.LastSessionDuration,
	PenultimateSessionMinute:                                sessionOldModel.LastSessionMinute,
	LastSessionYearOfDay:                                    sessionRespondModel.FirstSessionYearOfDay,
	LastSessionYear:                                         sessionRespondModel.FirstSessionYear,
	LastSessionHour:                                         sessionRespondModel.FirstSessionHour,
	LastSessionDuration:                                     sessionRespondModel.FirstSessionDuration,
	LastSessionMinute:                                       sessionRespondModel.FirstSessionMinute,
	LastDurationMinusPenultimateDuration:                    sessionRespondModel.FirstSessionDuration - sessionOldModel.LastSessionDuration,
	FirstHalfHourTotalSessionCount:                          0,
	FirstHalfHourTotalSessionDuration:                       0,
	FirstHourTotalSessionCount:                              0,
	FirstHourTotalSessionDuration:                           0,
	FirstTwoHourTotalSessionCount:                           0,
	FirstTwoHourTotalSessionDuration:                        0,
	FirstThreeHourTotalSessionCount:                         0,
	FirstThreeHourTotalSessionDuration:                      0,
	FirstSixHourTotalSessionCount:                           0,
	FirstSixHourTotalSessionDuration:                        0,
	FirstTwelveHourTotalSessionCount:                        0,
	FirstTwelveHourTotalSessionDuration:                     0,
	TotalSessionDay:                                         TotalSessionDay,
	TotalSessionHour:                                        TotalSessionHour,
	TotalSessionMinute:                                      TotalSessionMinute,
	TotalSessionDuration:                                    sessionRespondModel.TotalSessionDuration + sessionOldModel.TotalSessionDuration,
	TotalSessionCount:                                       sessionRespondModel.TotalSessionCount + sessionOldModel.TotalSessionCount,
	FirstDayTotalSessionCount:                               sessionOldModel.FirstDayTotalSessionCount,
	FirstDayTotalSessionDuration:                            sessionOldModel.FirstDayTotalSessionDuration,
	SecondDayTotalSessionCount:                              sessionRespondModel.FirstDayTotalSessionCount,
	SecondDayTotalSessionDuration:                           sessionRespondModel.FirstDayTotalSessionDuration,
	ThirdDayTotalSessionCount:                               0,
	ThirdDayTotalSessionDuration:                            0,
	FourthDayTotalSessionCount:                              0,
	FourthDayTotalSessionDuration:                           0,
	FifthDayTotalSessionCount:                               0,
	FifthDayTotalSessionDuration:                            0,
	SixthDayTotalSessionCount:                               0,
	SixthDayTotalSessionDuration:                            0,
	SeventhDayTotalSessionCount:                             0,
	SeventhDayTotalSessionDuration:                          0,
	MinSessionDuration:                                      int64(19),
	MaxSessionDuration:                                      int64(33),
	DailyAvegareSessionCount:                                float64(TotalSessionCount) / float64(TotalSessionDay),
	DailyAverageSessionDuration:                             float64(TotalSessionDuration) / float64(TotalSessionDay),
	SessionBasedAvegareSessionDuration:                      float64(TotalSessionDuration) / float64(TotalSessionCount),
	DailyAvegareSessionCountMinusFirstDaySessionCount:       float64(TotalSessionCount)/float64(TotalSessionDay) - float64(sessionOldModel.FirstDayTotalSessionCount),
	DailyAvegareSessionDurationMinusFirstDaySessionDuration: float64(TotalSessionDuration)/float64(TotalSessionDay) - float64(sessionOldModel.FirstDayTotalSessionDuration),
	SessionBasedAvegareSessionDurationMinusFirstSessionDuration: SessionBasedAvegareSessionDuration - float64(sessionOldModel.FirstSessionDuration),
	SessionBasedAvegareSessionDurationMinusLastSessionDuration:  SessionBasedAvegareSessionDuration - float64(sessionRespondModel.FirstSessionDuration),
	SundaySessionCount:     sessionOldModel.SundaySessionCount + sessionRespondModel.SundaySessionCount,
	MondaySessionCount:     sessionOldModel.MondaySessionCount + sessionRespondModel.MondaySessionCount,
	TuesdaySessionCount:    sessionOldModel.TuesdaySessionCount + sessionRespondModel.TuesdaySessionCount,
	WednesdaySessionCount:  sessionOldModel.WednesdaySessionCount + sessionRespondModel.WednesdaySessionCount,
	ThursdaySessionCount:   sessionOldModel.ThursdaySessionCount + sessionRespondModel.ThursdaySessionCount,
	FridaySessionCount:     sessionOldModel.FridaySessionCount + sessionRespondModel.FridaySessionCount,
	SaturdaySessionCount:   sessionOldModel.SaturdaySessionCount + sessionRespondModel.SaturdaySessionCount,
	AmSessionCount:         sessionOldModel.AmSessionCount + sessionRespondModel.AmSessionCount,
	PmSessionCount:         sessionOldModel.PmSessionCount + sessionRespondModel.PmSessionCount,
	Session0To5HourCount:   sessionOldModel.Session0To5HourCount + sessionRespondModel.Session0To5HourCount,
	Session6To11HourCount:  sessionOldModel.Session6To11HourCount + sessionRespondModel.Session6To11HourCount,
	Session12To17HourCount: sessionOldModel.Session12To17HourCount + sessionRespondModel.Session12To17HourCount,
	Session18To23HourCount: sessionOldModel.Session18To23HourCount + sessionRespondModel.Session18To23HourCount,
}

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
	var sessionOldModel_test = sessionOldModel
	var sessionUpdateModel_test = sessionUpdateModel
	var sessionRespondModel_test = sessionRespondModel
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
	//var sessionUpdateModel_test = sessionUpdateModel
	var sessionRespondModel_test = sessionRespondModel

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