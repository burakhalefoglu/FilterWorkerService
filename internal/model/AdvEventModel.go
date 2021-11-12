package model

import "time"

type AdvEventRespondModel struct {
	ProjectId                                          string
	ClientId                                           string
	CustomerId                                         string
	LevelIndex                                         int64
	TotalAdvDay                                        int64
	TotalAdvCount                                      int64
	LevelBasedAverageAdvCount                          float64
	AverageAdvDailyClickCount                          float64
	FirstAdvYearOfDay                                  int64
	FirstAdvYear                                       int64
	FirstAdvClickHour                                  int64
	FirstADvClickMinute                                int64
	FirstAdvType                                       int64
	SecondAdvYearOfDay                                 int64
	SecondAdvHour                                      int64
	SecondAdvMinute                                    int64
	ThirdAdvYearOfDay                                  int64
	ThirdAdvHour                                       int64
	ThirdAdvMinute                                     int64
	PenultimateAdvYearOfDay                            int64
	PenultimateAdvHour                                 int64
	PenultimateAdvMinute                               int64
	LastAdvYearOfDay                                   int64
	LastAdvYear                                        int64
	LastAdvClickHour                                   int64
	LastAdvClickMinute                                 int64
	LastAdvType                                        int64
	FirstDayAdvClickCount                              int64
	PenultimateDayAdvClickCount                        int64
	LastDayAdvClickCount                               int64
	LastMinusFirstDayAdvClickCount                     int64
	LastMinusPenultimateDayAdvClickCount               int64
	LastDayAdvClickCountMinusAverageDailyAdvClickCount int64
	SundayAdvClickCount                                int64
	MondayAdvClickCount                                int64
	TuesdayAdvClickCount                               int64
	WednesdayAdvClickCount                             int64
	ThursdayAdvClickCount                              int64
	FridayAdvClickCount                                int64
	SaturdayAdvClickCount                              int64
	AmAdvClickCount                                    int64
	PmAdvClickCount                                    int64
	AdvClick0To5HourCount                              int64
	AdvClick6To11HourCount                             int64
	AdvClick12To17HourCount                            int64
	AdvClick18To23HourCount                            int64
}

//IsdeadAndVideoClickCount                                  int64

type AdvEventModel struct {
	ProjectId   string
	ClientId    string
	CustomerId  string
	LevelName   string
	LevelIndex  int
	AdvType     string
	InMinutes   int
	TrigerdTime time.Time
}
