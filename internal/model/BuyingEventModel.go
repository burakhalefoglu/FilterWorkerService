package model

import "time"

type BuyingEventRespondModel struct {
	ProjectId                    string
	ClientId                     string
	CustomerId                   string
	LevelIndex                   int16
	TotalBuyingCount             int32
	TotalBuyingDay               int32
	TotalBuyingHour              int32
	FirstBuyingYearOfDay         int16
	FirstBuyingYear              int16
	FirstBuyingHour              int16
	FirstBuyingMinute            int16
	FirstBuyingProductType       int16
	SecondBuyingYearOfDay        int16
	SecondBuyingHour             int16
	SecondBuyingMinute           int16
	SecondBuyingProductType      int16
	ThirdBuyingYearOfDay         int16
	ThirdBuyingHour              int16
	ThirdBuyingMinute            int16
	ThirdBuyingProductType       int16
	FourthBuyingYearOfDay        int16
	FourthBuyingHour             int16
	FourthBuyingMinute           int16
	FourthBuyingProductType      int16
	FifthBuyingYearOfDay         int16
	FifthBuyingHour              int16
	FifthBuyingMinute            int16
	FifthBuyingProductType       int16
	PenultimateBuyingYearOfDay   int16
	PenultimateBuyingHour        int16
	PenultimateBuyingMinute      int16
	PenultimateBuyingProductType int16
	LastBuyingYearOfDay          int16
	LastBuyingYear               int16
	LastBuyingHour               int16
	LastBuyingMinute             int16
	LastBuyingProductType        int16
	FirstDayBuyingCount          int16
	SecondDayBuyingCount         int16
	ThirdDayBuyingCount          int16
	FourthDayBuyingCount         int16
	FifthDayBuyingCount          int16
	SixthDayBuyingCount          int16
	SeventhDayBuyingCount        int16
	SundayBuyingCount            int16
	MondayBuyingCount            int16
	TuesdayBuyingCount           int16
	WednesdayBuyingCount         int16
	ThursdayBuyingCount          int16
	FridayBuyingCount            int16
	SaturdayBuyingCount          int16
	AmBuyingCount                int16
	PmBuyingCount                int16
	Buying0To5HourCount          int16
	Buying6To11HourCount         int16
	Buying12To17HourCount        int16
	Buying18To23HourCount        int16
	BuyingDayAverageBuyingCount  float32
	LevelBasedAverageBuyingCount float32
}

// PenultimateDayBuyingCount    int16
// LastDayBuyingCount           int16
// LastMinusFirstDayBuyingCount int16

//IsDeadAndBuyingItemCount                               int16

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
