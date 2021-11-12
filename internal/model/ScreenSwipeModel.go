package model

import "time"

type ScreenSwipeRespondModel struct {
	ProjectId                    string
	ClientId                     string
	CustomerId                   string
	LevelIndex                   int64
	TotalSwipeSessionCount       int64
	FirstSwipeYearOfDay          int64
	FirstSwipeYear               int64
	FirstSwipeHour               int64
	FistSwipeDirection           int64
	FirstSwipeStartXCor          float64
	FirstSwipeStartYCor          float64
	FirstSwipeFinishXCor         float64
	FirstSwipeFinishYCor         float64
	SecondSwipeDirection         int64
	SecondSwipeStartXCor         float64
	SecondSwipeStartYCor         float64
	SecondSwipeFinishXCor        float64
	SecondSwipeFinishYCor        float64
	ThirdSwipeDirection          int64
	ThirdSwipeStartXCor          float64
	ThirdSwipeStartYCor          float64
	ThirdSwipeFinishXCor         float64
	ThirdSwipeFinishYCor         float64
	PenultimateSwipeDirection    int64
	PenultimateSwipeStartXCor    float64
	PenultimateSwipeStartYCor    float64
	PenultimateSwipeFinishXCor   float64
	PenultimateSwipeFinishYCor   float64
	LastSwipeDirection           int64
	LastSwipeStartXCor           float64
	LastSwipeStartYCor           float64
	LastSwipeFinishXCor          float64
	LastSwipeFinishYCor          float64
	FirstDayTotalSwipeUpCount    int64
	FirstDayTotalSwipeDownCount  int64
	FirstDayTotalSwipeRightCount int64
	FirstDayTotalSwipeLeftCount  int64
	FirstDaySwipeTotalStartXCor  float64
	FirstDaySwipeTotalStartYCor  float64
	FirstDaySwipeTotalFinishXCor float64
	FirstDaySwipeTotalFinishYCor float64
	TotalSwipeUpCount            int64
	TotalSwipeDownCount          int64
	TotalSwipeRightCount         int64
	TotalSwipeLeftCount          int64
	TotalSwipeStartXCor          float64
	TotalSwipeStartYCor          float64
	TotalSwipeFinishXCor         float64
	TotalSwipeFinishYCor         float64
}

type ScreenSwipeModel struct {
	ProjectId       string
	ClientId        string
	CustomerId      string
	SwipeDirection  int
	SwipeStartXCor  float64
	SwipeStartYCor  float64
	SwipeFinishXCor float64
	SwipeFinishYCor float64
	CreationAt      time.Time
	LevelIndex      int
	LevelName       string
}
