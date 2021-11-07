package model

import "time"

type AdvEventRespondModel struct {
	ProjectId                                                 string
	ClientId                                                  string
	CustomerId                                                string
	LevelIndex                                                int64
	TotalAdvDay                                               int64
	TotalVideoAdvCount                                        int64
	TotalInterstitialAdvCount                                 int64
	LevelBasedAverageInterstitialAdvCount                     float64
	LevelBasedAverageVideoAdvCount                            float64
	AverageDailyVideoAdvClickCount                            float64
	FirstVideoClickMonth                                      int64
	FirstVideoClickWeek                                       int64
	FirstVideoClickDay                                        int64
	FirstVideoClickHour                                       int64
	FirstDayVideoClickCount                                   int64
	PenultimateDayVdeoClickCount                              int64
	LastDayVideoClickCount                                    int64
	LastMinusPenultimateDayVideoClickCount                    int64
	LastMinusFirstDayVideoClickCount                          int64
	LastDayVideoClickCountMinusAverageDailyVideoAdvClickCount int64
	IsdeadAndVideoClickCount                                  int64
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
