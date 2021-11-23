package test

import (
	"FilterWorkerService/internal/model"
	"FilterWorkerService/internal/service/concrete"
	"FilterWorkerService/pkg/jsonParser/gojson"
	"FilterWorkerService/test/Mock/repository"
	"errors"

	//"fmt"
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
	var manager = concrete.GameSessionEveryLoginManager{
		IGameSessionEveryLoginDal: testSessionDal,
		IJsonParser:               &gojson.GoJson{},
	}
	testSessionDal.On("UpdateGameSessionEveryLoginById", sessionOldModel.ClientId, &sessionUpdateModel).Return(nil)
	var v, s, m = manager.UpdateGameSession(&sessionRespondModel, &sessionOldModel)
	assert.Equal(t, true, s)
	assert.Equal(t, nil, m)
	assert.Equal(t, &sessionUpdateModel, v)
}

func Test_ConvertRawModelToResponseModel_AddS(t *testing.T) {
	var testSessionDal = new(repository.MockGameSessionEveryLoginDal)
	var manager = concrete.GameSessionEveryLoginManager{
		IGameSessionEveryLoginDal: testSessionDal,
		IJsonParser:               &gojson.GoJson{},
	}
	gameByte, _ := manager.IJsonParser.EncodeJson(sessionModel)
	testSessionDal.On("GetGameSessionEveryLoginById", sessionRespondModel.ClientId).Return(&sessionOldModel,
		errors.New("mongo: no documents in result"))
	testSessionDal.On("Add", &sessionRespondModel).Return(nil)
	var v, s, m = manager.ConvertRawModelToResponseModel(gameByte)

	assert.Equal(t, true, s)
	assert.Equal(t, "Added", m)
	assert.Equal(t, &sessionRespondModel, v)
}


func Test_CalculateSecondDayTotalSessionCountAndDuration(t *testing.T) {
	var count, duration = concrete.CalculateSecondDayTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionHour)
	var ExpCount int64 = 1
	var ExpDuration int64 = 33
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateFirstTwentyFourTotalSessionCountAndDuration(t *testing.T) {
	var count, duration = concrete.CalculateFirstTwentyFourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute)
	var ExpCount int64 = 4
	var ExpDuration int64 = 51
	// var ExpCount int64= 0
	// var ExpDuration int64= 0
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateFirstTwelveHourTotalSessionCountAndDuration(t *testing.T) {
	var count, duration = concrete.CalculateFirstTwelveHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute)
	var ExpCount int64 = 1
	var ExpDuration int64 = 33
	// var ExpCount int64= 0
	// var ExpDuration int64= 0
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateFirstSixHourTotalSessionCountAndDuration(t *testing.T) {
	var count, duration = concrete.CalculateFirstSixHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute)
	// var ExpCount int64= 1
	// var ExpDuration int64= 33
	var ExpCount int64 = 0
	var ExpDuration int64 = 0
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateFirstThreeHourTotalSessionCountAndDuration(t *testing.T) {
	var count, duration = concrete.CalculateFirstThreeHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute)
	var ExpCount int64 = 1
	var ExpDuration int64 = 33
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateFirstTwoHourTotalSessionCountAndDuration(t *testing.T) {
	var count, duration = concrete.CalculateFirstTwoHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute)
	var ExpCount int64 = 1
	var ExpDuration int64 = 33
	// var ExpCount int64= 0
	// var ExpDuration int64= 0
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateFirstHourTotalSessionCountAndDuration(t *testing.T) {
	var count, duration = concrete.CalculateFirstHourTotalSessionCountAndDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute)
	var ExpCount int64 = 0
	var ExpDuration int64 = 0
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateFirstHalfHourTotalSessionCountAndSessionDuration(t *testing.T) {
	var count, duration = concrete.CalculateFirstHalfHourTotalSessionCountAndSessionDuration(&sessionRespondModel, &sessionOldModel, TotalSessionMinute)
	var ExpCount int64 = 1
	var ExpDuration int64 = 33
	assert.Equal(t, []int64{ExpCount, ExpDuration}, []int64{count, duration})
}

func Test_CalculateSecondGameSession(t *testing.T) {
	var hour, duration, minute int64 = concrete.CalculateSecondGameSession(&sessionRespondModel, &sessionOldModel)
	//var Exphour, Expduration, Expminute int64 = sessionRespondModel.FirstSessionHour, sessionRespondModel.FirstSessionDuration, sessionRespondModel.FirstSessionMinute
	var Exphour, Expduration, Expminute int64 = sessionOldModel.SecondSessionHour, sessionOldModel.SecondSessionDuration, sessionOldModel.SecondSessionMinute
	assert.Equal(t, []int64{Exphour, Expduration, Expminute}, []int64{hour, duration, minute})
}


func Test_CalculateMaxDuration(t *testing.T) {
	var duration = concrete.CalculateMaxDuration(&sessionRespondModel, &sessionOldModel)
	var ExpDuration int64 = 33
	assert.Equal(t, []int64{ExpDuration}, []int64{duration})
}

func Test_CalculateMinDuration(t *testing.T) {
	var duration = concrete.CalculateMinDuration(&sessionRespondModel, &sessionOldModel)
	var ExpDuration int64 = 19
	assert.Equal(t, []int64{ExpDuration}, []int64{duration})
}

func Test_CalculateThirdGameSession(t *testing.T) {
	var hour, duration, minute int64 = concrete.CalculateThirdGameSession(&sessionRespondModel, &sessionOldModel)
	//var Exphour, Expduration, Expminute int64 = sessionRespondModel.FirstSessionHour, sessionRespondModel.FirstSessionDuration, sessionRespondModel.FirstSessionMinute
	var Exphour, Expduration, Expminute int64 = sessionOldModel.ThirdSessionHour, sessionOldModel.ThirdSessionDuration, sessionOldModel.ThirdSessinMinute
	assert.Equal(t, []int64{Exphour, Expduration, Expminute}, []int64{hour, duration, minute})
}

func Test_DetermineGameSessionHour(t *testing.T){
	concrete.DetermineGameSessionHour(&sessionRespondModel, sessionRespondModel.FirstSessionHour)

	assert.Equal(t, int64(0), sessionRespondModel.Session0To5HourCount)
	assert.Equal(t, int64(0), sessionRespondModel.Session6To11HourCount )
	assert.Equal(t, int64(0), sessionRespondModel.Session12To17HourCount)
	assert.Equal(t, int64(1), sessionRespondModel.Session18To23HourCount)
}

func Test_DetermineGameSessionDay(t *testing.T){
	concrete.DetermineGameSessionDay(&sessionRespondModel, sessionRespondModel.FirstSessionWeekDay)

	assert.Equal(t, int64(0), sessionRespondModel.SundaySessionCount)
	assert.Equal(t, int64(1), sessionRespondModel.MondaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel.TuesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel.WednesdaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel.ThursdaySessionCount )
	assert.Equal(t, int64(0), sessionRespondModel.FridaySessionCount)
	assert.Equal(t, int64(0), sessionRespondModel.SaturdaySessionCount )
}


func Test_DetermineGameSessionAmPm(t *testing.T){
	concrete.DetermineGameSessionAmPm(&sessionRespondModel, sessionRespondModel.FirstSessionHour)

	assert.Equal(t, int64(0), sessionRespondModel.AmSessionCount)
	assert.Equal(t, int64(1), sessionRespondModel.PmSessionCount)

}