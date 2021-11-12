package model

import "time"

type LevelBaseSessionRespondModel struct {
	ProjectId                            string
	ClientId                             string
	CustomerId                           string
	TotalLevelBaseSessionCount           int64
	FirstLevelSessionLevelIndex          int64
	FirstLevelSessionDuration            int64
	SecondLevelSessionLevelIndex         int64
	SecondLevelSessionDuration           int64
	ThirdLevelSessionLevelIndex          int64
	ThirdLevelSessionDuration            int64
	FourLevelSessionLevelIndex           int64
	FourLevelSessionDuration             int64
	FiveLevelSessionLevelIndex           int64
	FiveLevelSessionDuration             int64
	SixLevelSessionLevelIndex            int64
	SixLevelSessionDuration              int64
	SevenLevelSessionLevelIndex          int64
	SevenLevelSessionDuration            int64
	PenultimateLevelSessionLevelIndex    int64
	PenultimateLevelSessionLevelDuration int64
	LastLevelSessionLevelIndex           int64
	LastLevelSessionLevelDuration        int64
}

type LevelBaseSessionDataModel struct {
	ProjectId         string
	ClientId          string
	CustomerId        string
	SessionStartTime  time.Time
	SessionFinishTime time.Time
	SessionTimeMinute int
	LevelIndex        int
	LevelName         string
}
