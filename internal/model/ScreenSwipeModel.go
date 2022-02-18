package model

import "time"

type ScreenSwipeResponseModel struct {
	Id                             int64
	ClientId                       int64
	ProjectId                      int64
	CustomerId                     int64
	LevelIndex                     int16
	TotalSwipeSessionCount         int32
	TotalSwipeHour                 int32
	FirstSwipeYearOfDay            int16
	FirstSwipeYear                 int16
	FirstSwipeHour                 int16
	FirstSwipeWeekDay              int16
	FirstSwipeMinute               int16
	FistSwipeDirection             byte
	FirstSwipeStartXCor            float32
	FirstSwipeStartYCor            float32
	FirstSwipeFinishXCor           float32
	FirstSwipeFinishYCor           float32
	SecondSwipeDirection           byte
	SecondSwipeStartXCor           float32
	SecondSwipeStartYCor           float32
	SecondSwipeFinishXCor          float32
	SecondSwipeFinishYCor          float32
	ThirdSwipeDirection            byte
	ThirdSwipeStartXCor            float32
	ThirdSwipeStartYCor            float32
	ThirdSwipeFinishXCor           float32
	ThirdSwipeFinishYCor           float32
	FourthSwipeDirection           byte
	FourthSwipeStartXCor           float32
	FourthSwipeStartYCor           float32
	FourthSwipeFinishXCor          float32
	FourthSwipeFinishYCor          float32
	FifthSwipeDirection            byte
	FifthSwipeStartXCor            float32
	FifthSwipeStartYCor            float32
	FifthSwipeFinishXCor           float32
	FifthSwipeFinishYCor           float32
	SixthSwipeDirection            byte
	SixthSwipeStartXCor            float32
	SixthSwipeStartYCor            float32
	SixthSwipeFinishXCor           float32
	SixthSwipeFinishYCor           float32
	SeventhSwipeDirection          byte
	SeventhSwipeStartXCor          float32
	SeventhSwipeStartYCor          float32
	SeventhSwipeFinishXCor         float32
	SeventhSwipeFinishYCor         float32
	PenultimateSwipeDirection      byte
	PenultimateSwipeStartXCor      float32
	PenultimateSwipeStartYCor      float32
	PenultimateSwipeFinishXCor     float32
	PenultimateSwipeFinishYCor     float32
	PenultimateSwipeYearOfDay      int16
	PenultimateSwipeYear           int16
	PenultimateSwipeHour           int16
	PenultimateSwipeWeekDay        int16
	PenultimateSwipeMinute         int16
	LastSwipeDirection             byte
	LastSwipeStartXCor             float32
	LastSwipeStartYCor             float32
	LastSwipeFinishXCor            float32
	LastSwipeFinishYCor            float32
	LastSwipeYearOfDay             int16
	LastSwipeYear                  int16
	LastSwipeHour                  int16
	LastSwipeWeekDay               int16
	LastSwipeMinute                int16
	FirstDayTotalSwipeUpCount      int32
	FirstDayTotalSwipeDownCount    int32
	FirstDayTotalSwipeRightCount   int32
	FirstDayTotalSwipeLeftCount    int32
	FirstDaySwipeTotalStartXCor    float32
	FirstDaySwipeTotalStartYCor    float32
	FirstDaySwipeTotalFinishXCor   float32
	FirstDaySwipeTotalFinishYCor   float32
	SecondDayTotalSwipeUpCount     int32
	SecondDayTotalSwipeDownCount   int32
	SecondDayTotalSwipeRightCount  int32
	SecondDayTotalSwipeLeftCount   int32
	SecondDaySwipeTotalStartXCor   float32
	SecondDaySwipeTotalStartYCor   float32
	SecondDaySwipeTotalFinishXCor  float32
	SecondDaySwipeTotalFinishYCor  float32
	ThirdDayTotalSwipeUpCount      int32
	ThirdDayTotalSwipeDownCount    int32
	ThirdDayTotalSwipeRightCount   int32
	ThirdDayTotalSwipeLeftCount    int32
	ThirdDaySwipeTotalStartXCor    float32
	ThirdDaySwipeTotalStartYCor    float32
	ThirdDaySwipeTotalFinishXCor   float32
	ThirdDaySwipeTotalFinishYCor   float32
	FourthDayTotalSwipeUpCount     int32
	FourthDayTotalSwipeDownCount   int32
	FourthDayTotalSwipeRightCount  int32
	FourthDayTotalSwipeLeftCount   int32
	FourthDaySwipeTotalStartXCor   float32
	FourthDaySwipeTotalStartYCor   float32
	FourthDaySwipeTotalFinishXCor  float32
	FourthDaySwipeTotalFinishYCor  float32
	FifthDayTotalSwipeUpCount      int32
	FifthDayTotalSwipeDownCount    int32
	FifthDayTotalSwipeRightCount   int32
	FifthDayTotalSwipeLeftCount    int32
	FifthDaySwipeTotalStartXCor    float32
	FifthDaySwipeTotalStartYCor    float32
	FifthDaySwipeTotalFinishXCor   float32
	FifthDaySwipeTotalFinishYCor   float32
	SixthDayTotalSwipeUpCount      int32
	SixthDayTotalSwipeDownCount    int32
	SixthDayTotalSwipeRightCount   int32
	SixthDayTotalSwipeLeftCount    int32
	SixthDaySwipeTotalStartXCor    float32
	SixthDaySwipeTotalStartYCor    float32
	SixthDaySwipeTotalFinishXCor   float32
	SixthDaySwipeTotalFinishYCor   float32
	SeventhDayTotalSwipeUpCount    int32
	SeventhDayTotalSwipeDownCount  int32
	SeventhDayTotalSwipeRightCount int32
	SeventhDayTotalSwipeLeftCount  int32
	SeventhDaySwipeTotalStartXCor  float32
	SeventhDaySwipeTotalStartYCor  float32
	SeventhDaySwipeTotalFinishXCor float32
	SeventhDaySwipeTotalFinishYCor float32
	TotalSwipeUpCount              int32
	TotalSwipeDownCount            int32
	TotalSwipeRightCount           int32
	TotalSwipeLeftCount            int32
	TotalSwipeStartXCor            float32
	TotalSwipeStartYCor            float32
	TotalSwipeFinishXCor           float32
	TotalSwipeFinishYCor           float32
	Status                         bool
}

type ScreenSwipeModel struct {
	Id             int64
	ClientId       int64
	ProjectId      int64
	CustomerId     int64
	StartLocX      float32
	StartLocY      float32
	FinishLocX     float32
	FinishLocY     float32
	SwipeDirection int
	LevelName      string
	LevelIndex     int
	CreatedAt      time.Time
	Status         bool
}
