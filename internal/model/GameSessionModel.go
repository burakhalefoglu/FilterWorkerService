package model

import "time"

type GameSessionResponseModel struct {
	Id                                                          int64
	ClientId                                                    int64
	ProjectId                                                   int64
	CustomerId                                                  int64
	FirstSessionYearOfDay                                       int16
	FirstSessionYear                                            int16
	FirstSessionWeekDay                                         int16
	FirstSessionHour                                            int16
	FirstSessionDuration                                        int16
	FirstSessionMinute                                          int16
	SecondSessionHour                                           int16
	SecondSessionDuration                                       int16
	SecondSessionMinute                                         int16
	ThirdSessionHour                                            int16
	ThirdSessionDuration                                        int16
	ThirdSessionMinute                                          int16
	FourthSessionHour                                           int16
	FourthSessionDuration                                       int16
	FourthSessionMinute                                         int16
	FifthSessionHour                                            int16
	FifthSessionDuration                                        int16
	FifthSessionMinute                                          int16
	SixthSessionHour                                            int16
	SixthSessionDuration                                        int16
	SixthSessionMinute                                          int16
	SeventhSessionHour                                          int16
	SeventhSessionDuration                                      int16
	SeventhSessionMinute                                        int16
	PenultimateSessionHour                                      int16
	PenultimateSessionDuration                                  int16
	PenultimateSessionMinute                                    int16
	LastSessionYearOfDay                                        int16
	LastSessionYear                                             int16
	LastSessionHour                                             int16
	LastSessionDuration                                         int16
	LastSessionMinute                                           int16
	LastDurationMinusPenultimateDuration                        int16
	FirstFiveMinutesTotalSessionCount                           int16
	FirstFiveMinutesTotalSessionDuration                        int16
	FirstTenMinutesTotalSessionCount                            int16
	FirstTenMinutesTotalSessionDuration                         int16
	FirstQuarterHourTotalSessionCount                           int16
	FirstQuarterHourTotalSessionDuration                        int16
	FirstHalfHourTotalSessionCount                              int16
	FirstHalfHourTotalSessionDuration                           int16
	FirstHourTotalSessionCount                                  int16
	FirstHourTotalSessionDuration                               int16
	FirstTwoHourTotalSessionCount                               int16
	FirstTwoHourTotalSessionDuration                            int16
	FirstThreeHourTotalSessionCount                             int16
	FirstThreeHourTotalSessionDuration                          int16
	FirstSixHourTotalSessionCount                               int16
	FirstSixHourTotalSessionDuration                            int16
	FirstTwelveHourTotalSessionCount                            int16
	FirstTwelveHourTotalSessionDuration                         int16
	TotalSessionDay                                             int32
	TotalSessionHour                                            int32
	TotalSessionMinute                                          int32
	TotalSessionDuration                                        int32
	TotalSessionCount                                           int32
	FirstDayTotalSessionCount                                   int16
	FirstDayTotalSessionDuration                                int16
	SecondDayTotalSessionCount                                  int16
	SecondDayTotalSessionDuration                               int16
	ThirdDayTotalSessionCount                                   int16
	ThirdDayTotalSessionDuration                                int16
	FourthDayTotalSessionCount                                  int16
	FourthDayTotalSessionDuration                               int16
	FifthDayTotalSessionCount                                   int16
	FifthDayTotalSessionDuration                                int16
	SixthDayTotalSessionCount                                   int16
	SixthDayTotalSessionDuration                                int16
	SeventhDayTotalSessionCount                                 int16
	SeventhDayTotalSessionDuration                              int16
	MinSessionDuration                                          int16
	MaxSessionDuration                                          int16
	DailyAvegareSessionCount                                    float32
	DailyAverageSessionDuration                                 float32
	SessionBasedAvegareSessionDuration                          float32
	DailyAvegareSessionCountMinusFirstDaySessionCount           float32
	DailyAvegareSessionDurationMinusFirstDaySessionDuration     float32
	SessionBasedAvegareSessionDurationMinusFirstSessionDuration float32
	SessionBasedAvegareSessionDurationMinusLastSessionDuration  float32
	SundaySessionCount                                          int16
	MondaySessionCount                                          int16
	TuesdaySessionCount                                         int16
	WednesdaySessionCount                                       int16
	ThursdaySessionCount                                        int16
	FridaySessionCount                                          int16
	SaturdaySessionCount                                        int16
	AmSessionCount                                              int16
	PmSessionCount                                              int16
	Session0To5HourCount                                        int16
	Session6To11HourCount                                       int16
	Session12To17HourCount                                      int16
	Session18To23HourCount                                      int16
	Status                                                      bool
}

type GameSessionModel struct {
	Id                int64
	ClientId          int64
	ProjectId         int64
	CustomerId        int64
	CreatedAt         time.Time
	SessionTime       float32
	SessionStartTime  time.Time
	SessionFinishTime time.Time
	Status            bool
}
