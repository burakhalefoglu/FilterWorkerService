package model

import "time"

type AdvEventRespondModel struct {
	ProjectId                                                 string
	ClientId                                                  string
	CustomerId                                                string
	LevelIndex                                                int64
	TotalAdvDay                                               int64
	TotalAdvCount                                             int64
	TotalVideoAdvCount                                        int64
	TotalInterstitialAdvCount                                 int64
	LevelBasedAverageInterstitialAdvCount                     float64
	LevelBasedAverageVideoAdvCount                            float64
	AverageAdvDailyVideoClickCount                            float64
	FirstAdvYearOfDay                                         int64
	FirstAdvYear                                              int64
	FirstAdvClickHour                                         int64
	FirstVideoClickYearOfDay                                  int64
	FirstVideoClickHour                                       int64
	FirstAdvType                                              int64
	LastAdvYearOfDay                                          int64
	LastAdvYear                                               int64
	LastVideoClickYearOfDay                                   int64
	LastAdvClickHour                                          int64
	LastAdvType                                               int64
	FirstDayVideoClickCount                                   int64
	LastDayVideoClickCount                                    int64
	LastMinusFirstDayVideoClickCount                          int64
	LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount int64
	SundayVideoAdvClickCount                                  int64
	MondayVideoAdvClickCount                                  int64
	TuesdayVideoAdvClickCount                                 int64
	WednesdayVideoAdvClickCount                               int64
	ThursdayVideoAdvClickCount                                int64
	FridayVideoAdvClickCount                                  int64
	SaturdayVideoAdvClickCount                                int64
	AmVideoAdvClickCount                                      int64
	PmVideoAdvClickCount                                      int64
	VideoAdvClick0To5HourCount                                int64
	VideoAdvClick6To11HourCount                               int64
	VideoAdvClick12To17HourCount                              int64
	VideoAdvClick18To23HourCount                              int64
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
