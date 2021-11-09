package model

import "time"

type ScreenSwipeRespondModel struct {
	ProjectId              string
	ClientId               string
	CustomerId             string
	LevelIndex             int64
	FistSwipeDirection     int64
	FirstSwipeStartXCor    float64
	FirstSwipeStartYCor    float64
	FirstSwipeFinishXCor   float64
	FirstSwipeFinishYCor   float64
	LastSwipeDirection     int64
	LastSwipeStartXCor     float64
	LastSwipeStartYCor     float64
	LastSwipeFinishXCor    float64
	LastSwipeFinishYCor    float64
	TotalSwipeUpCount      int64
	TotalSwipeDownCount    int64
	TotalSwipeRightCount   int64
	TotalSwipeLeftCount    int64
	TotalSwipeStartXCor    int64
	TotalSwipeStartYCor    int64
	TotalSwipeFinishXCor   int64
	TotalSwipeFinishYCor   int64
	TotalSwipeSessionCount int64
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
