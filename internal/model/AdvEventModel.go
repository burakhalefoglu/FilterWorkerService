package model

import "time"

type AdvEventResponseModel struct {
	Id                                                 int64
	ClientId                                           int64
	ProjectId                                          int64
	CustomerId                                         int64
	LevelIndex                                         int16
	TotalAdvDay                                        int32
	TotalAdvCount                                      int32
	TotalAdvHour                                       int32
	TotalAdvMinute                                     int32
	LevelBasedAverageAdvCount                          float32
	AverageAdvDailyClickCount                          float32
	FirstAdvYearOfDay                                  int16
	FirstAdvYear                                       int16
	FirstWeekDay                                       int16
	FirstAdvClickHour                                  int16
	FirstAdvClickMinute                                int16
	FirstAdvType                                       byte
	SecondAdvYearOfDay                                 int16
	SecondAdvHour                                      int16
	SecondAdvMinute                                    int16
	SecondAdvType                                      byte
	ThirdAdvYearOfDay                                  int16
	ThirdAdvHour                                       int16
	ThirdAdvMinute                                     int16
	ThirdAdvType                                       byte
	FourthAdvYearOfDay                                 int16
	FourthAdvHour                                      int16
	FourthAdvMinute                                    int16
	FourthAdvType                                      byte
	FifthAdvYearOfDay                                  int16
	FifthAdvHour                                       int16
	FifthAdvMinute                                     int16
	FifthAdvType                                       byte
	SixthAdvYearOfDay                                  int16
	SixthAdvHour                                       int16
	SixthAdvMinute                                     int16
	SixthAdvType                                       byte
	SeventhAdvYearOfDay                                int16
	SeventhAdvHour                                     int16
	SeventhAdvMinute                                   int16
	SeventhAdvType                                     byte
	PenultimateAdvYearOfDay                            int16
	PenultimateAdvHour                                 int16
	PenultimateAdvMinute                               int16
	PenultimateAdvType                                 byte
	LastAdvYearOfDay                                   int16
	LastAdvYear                                        int16
	LastAdvClickHour                                   int16
	LastAdvClickMinute                                 int16
	LastAdvType                                        byte
	FirstFiveMinutesAdvClickCount                      int16
	FirstTenMinutesAdvClickCount                       int16
	FirstQuarterHourAdvClickCount                      int16
	FirstHalfHourAdvClickCount                         int16
	FirstHourAdvClickCount                             int16
	FirstTwoHourAdvClickCount                          int16
	FirstThreeHourAdvClickCount                        int16
	FirstSixHourAdvClickCount                          int16
	FirstTwelveHourAdvClickCount                       int16
	FirstDayAdvClickCount                              int16
	SecondDayAdvClickCount                             int16
	ThirdDayAdvClickCount                              int16
	FourthDayAdvClickCount                             int16
	FifthDayAdvClickCount                              int16
	SixthDayAdvClickCount                              int16
	SeventhDayAdvClickCount                            int16
	PenultimateDayAdvClickCount                        int16
	LastDayAdvClickCount                               int16
	LastMinusFirstDayAdvClickCount                     int16
	LastMinusPenultimateDayAdvClickCount               int16
	LastDayAdvClickCountMinusAverageDailyAdvClickCount float32
	SundayAdvClickCount                                int16
	MondayAdvClickCount                                int16
	TuesdayAdvClickCount                               int16
	WednesdayAdvClickCount                             int16
	ThursdayAdvClickCount                              int16
	FridayAdvClickCount                                int16
	SaturdayAdvClickCount                              int16
	AmAdvClickCount                                    int16
	PmAdvClickCount                                    int16
	AdvClick0To5HourCount                              int16
	AdvClick6To11HourCount                             int16
	AdvClick12To17HourCount                            int16
	AdvClick18To23HourCount                            int16
	Status                                             bool
}

//IsdeadAndVideoClickCount                                  int16

type AdvEventDataModel struct {
	Id            int64
	ClientId      int64
	ProjectId     int64
	CustomerId    int64
	LevelName     string
	LevelIndex    int32
	AdvType       string
	InMinutes     float32
	TriggeredTime time.Time
	Status        bool
}
