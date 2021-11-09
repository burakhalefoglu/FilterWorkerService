package model

import "time"

type ScreenClickRespondModel struct {
	ProjectId                                        string
	ClientId                                         string
	CustomerId                                       string
	LevelIndex                                       int64
	FirstClickSessionYearOfDay                       int64
	FirstClickSessionYear                            int64
	FirstClickSessionHour                            int64
	LastClickSessionYearOfDay                        int64
	LastClickSessionYear                             int64
	LastClickSessionHour                             int64
	FirstStartXCor                                   float64
	FirstStartYCor                                   float64
	FirstFinishXCor                                  float64
	FirstFinishYCor                                  float64
	LastStartXCor                                    float64
	LastStartYCor                                    float64
	LastFinishXCor                                   float64
	LastFinishYCor                                   float64
	FirstTouchCount                                  int64
	LastTouchCount                                   int64
	FirstMinusLastTouchCount                         int64
	FirstFingerId                                    int64
	LastFingerId                                     int64
	FirstDayClickCount                               int64
	TotalClickDay                                    int64
	TotalClickCount                                  int64
	TotalClickSessionCount                           int64
	TotalStartXCor                                   int64
	TotalStartYCor                                   int64
	TotalFinishXCor                                  int64
	TotalFinishYCor                                  int64
	SessionBasedAvegareStartXCor                     float64
	SessionBasedAvegareStartYCor                     float64
	SessionBasedAvegareFinishXCor                    float64
	SessionBasedAvegareFinishYCor                    float64
	SessionBasedAvegareClickCount                    float64
	DailyAvegareClickCount                           float64
	LastTouchCountMinusSessionBasedAvegareClickCount float64
}

type ScreenClickModel struct {
	ProjectId  string
	ClientId   string
	CustomerId string
	StartXCor  float64
	StartYCor  float64
	FinishXCor float64
	FinishYCor float64
	TouchCount int
	FingerId   int
	LevelIndex int
	LevelName  string
	CreationAt time.Time
}
