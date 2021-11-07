package model

import "time"

type AdvEventRespondModel struct {
	ProjectId   string
	ClientId    string
	CustomerId string
	LevelIndex  int64
	VideoAdvCount int64
	InterstitialAdvCount int64
	VideoClickMonth int64
	VideoClickWeek int64
	VideoClickDay int64
	VideoClickHour int64
}

type AdvEventModel struct {
	ProjectId   string
	ClientId    string
	CustomerId string
	LevelName string
	LevelIndex  int
	AdvType     string
	InMinutes   int
	TrigerdTime time.Time
}