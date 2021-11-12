package model

import "time"

type BuyingEventRespondModel struct {
	ProjectId                    string
	ClientId                     string
	CustomerId                   string
	LevelIndex                   int64
	TotalBuyingCount             int64
	TotalBuyingDay               int64
	FirstBuyingYearOfDay         int64
	FirstBuyingYear              int64
	FirstBuyingHour              int64
	SecondBuyingYearOfDay        int64
	SecondBuyingHour             int64
	ThirdBuyingYearOfDay         int64
	ThirdBuyingHour              int64
	PenultimateBuyingYearOfDay   int64
	PenultimateBuyingHour        int64
	LastBuyingYearOfDay          int64
	LastBuyingYear               int64
	LastBuyingHour               int64
	FirstDayBuyingCount          int64
	PenultimateDayBuyingCount    int64
	LastDayBuyingCount           int64
	LastMinusFirstDayBuyingCount int64
	SundayBuyingCount            int64
	MondayBuyingCount            int64
	TuesdayBuyingCount           int64
	WednesdayBuyingCount         int64
	ThursdayBuyingCount          int64
	FridayBuyingCount            int64
	SaturdayBuyingCount          int64
	AmBuyingCount                int64
	PmBuyingCount                int64
	Buying0To5HourCount          int64
	Buying6To11HourCount         int64
	Buying12To17HourCount        int64
	Buying18To23HourCount        int64
	BuyingDayAverageBuyingCount  float64
	LevelBasedAverageBuyingCount float64
}

//IsDeadAndBuyingItemCount                               int64

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
