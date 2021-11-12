package model

import "time"

type GameSessionEveryLoginRespondModel struct {
	ProjectId                                                   string
	ClientId                                                    string
	CustomerId                                                  string
	FirstSessionYearOfDay                                       int64
	FirstSessionYear                                            int64
	FirstSessionWeekDay                                         int64
	FirstSessionHour                                            int64
	FirstSessionDuration                                        int64
	FirstSessionMinute                                          int64
	SecondSessionHour                                           int64
	SecondSessionDuration                                       int64
	SecondSessionMinute                                         int64
	ThirdSessionHour                                            int64
	ThirdSessionDuration                                        int64
	ThirdSessinMinute                                           int64
	PenultimateSessionHour                                      int64
	PenultimateSessionDuration                                  int64
	PenultimateSessionMinute                                    int64
	LastSessionYearOfDay                                        int64
	LastSessionYear                                             int64
	LastSessionHour                                             int64
	LastSessionDuration                                         int64
	LastSessionMinute                                           int64
	LastDurationMinusPenultimateDuration                        int64
	TotalSessionDay                                             int64
	TotalSessionDuration                                        int64
	TotalSessionCount                                           int64
	FirstDayTotalSessionCount                                   int64
	FirstDayTotalSessionDuration                                int64
	PenultimateDayTotalSessionDuration                          int64
	PenultimateDayTotalSessionCount                             int64
	LastDayTotalSessionCount                                    int64
	LastDayTotalSessionDuration                                 int64
	MinSessionDuration                                          int64
	MaxSessionDuration                                          int64
	DailyAvegareSessionCount                                    float64
	DailyAverageSessionDuration                                 float64
	SessionBasedAvegareSessionDuration                          float64
	DailyAvegareSessionCountMinusFirstDaySessionCount           float64
	DailyAvegareSessionDurationMinusFirstDaySessionDuration     float64
	SessionBasedAvegareSessionDurationMinusFirstSessionDuration float64
	SessionBasedAvegareSessionDurationMinusLastSessionDuration  float64
	SundaySessionCount                                          int64
	MondaySessionCount                                          int64
	TuesdaySessionCount                                         int64
	WednesdaySessionCount                                       int64
	ThursdaySessionCount                                        int64
	FridaySessionCount                                          int64
	SaturdaySessionCount                                        int64
	AmSessionCount                                              int64
	PmSessionCount                                              int64
	Session0To5HourCount                                        int64
	Session6To11HourCount                                       int64
	Session12To17HourCount                                      int64
	Session18To23HourCount                                      int64
}

type GameSessionEveryLoginModel struct {
	ProjectId         string
	ClientId          string
	CustomerId        string
	SessionStartTime  time.Time
	SessionFinishTime time.Time
	SessionTimeMinute int
}
