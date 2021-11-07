package model

import "time"

type BuyingEventRespondModel struct {
	ProjectId                                              string
	ClientId                                               string
	CustomerId                                             string
	LevelIndex                                             int64
	TotalBuyingCount                                       int64
	TotalBuyingDay                                         int64
	TotalBuyingSession                                     int64
	TotalSession                                           int64
	TotalDay                                               int64
	FirstBuyingMonth                                       int64
	FirstBuyingWeek                                        int64
	FirstBuyingDay                                         int64
	FirstBuyingHour                                        int64
	LastBuyingMonth                                        int64
	LastBuyingWeek                                         int64
	LastBuyingDay                                          int64
	LastBuyingHour                                         int64
	FirstDayBuyingCount                                    int64
	PenultimateDayBuyingCount                              int64
	LastDayBuyingCount                                     int64
	LastMinusPenultimateDayBuyingCount                     int64
	LastMinusFirstDayBuyingCount                           int64
	SundayBuyingCount                                      int64
	MondayBuyingCount                                      int64
	TuesdayBuyingCount                                     int64
	WednesdayBuyingCount                                   int64
	ThursdayBuyingCount                                    int64
	FridayBuyingCount                                      int64
	SaturdayBuyingCount                                    int64
	AmBuyingCount                                          int64
	PmBuyingCount                                          int64
	Buying0To5HourCount                                    int64
	Buying6To11HourCount                                   int64
	Buying12To17HourCount                                  int64
	Buying18To23HourCount                                  int64
	DailyAverageBuyingCount                                int64
	BuyingDayAverageBuyingCount                            int64
	LevelBasedAverageBuyingCount                           int64
	SessionBasedAverageBuyingCount                         int64
	FirstBuyingDayMinusFirstSessionDay                     int64
	FirstBuyingMonthMinusFirstSessionMonth                 int64
	TotalDifferenceBetweenFirstBuyingDayAndFirstSessionDay int64
	IsDeadAndBuyingItemCount                               int64
}

type BuyingEventModel struct {
	ProjectId     string
	ClientId      string
	CustomerId    string
	LevelName     string
	LevelIndex    int
	InWhatMinutes int
	ProductType   string
	TrigerdTime   time.Time
}
