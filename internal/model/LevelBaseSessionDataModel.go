package model

import "time"

type LevelBaseSessionResponseModel struct {
	Id                                         int64
	ClientId                                   int64
	ProjectId                                  int64
	CustomerId                                 int64
	TotalLevelBaseSessionMinute                int32
	TotalLevelBaseSessionCount                 int32
	FirstLevelSessionLevelIndex                int16
	FirstLevelSessionDuration                  int16
	FirstLevelSessionYearOfDay                 int16
	FirstLevelSessionYear                      int16
	FirstLevelSessionWeekDay                   int16
	FirstLevelSessionHour                      int16
	FirstLevelSessionMinute                    int16
	SecondLevelSessionLevelIndex               int16
	SecondLevelSessionDuration                 int16
	ThirdLevelSessionLevelIndex                int16
	ThirdLevelSessionDuration                  int16
	FourLevelSessionLevelIndex                 int16
	FourLevelSessionDuration                   int16
	FiveLevelSessionLevelIndex                 int16
	FiveLevelSessionDuration                   int16
	SixLevelSessionLevelIndex                  int16
	SixLevelSessionDuration                    int16
	SevenLevelSessionLevelIndex                int16
	SevenLevelSessionDuration                  int16
	FirstFiveMinutesTotalLevelBaseSessionCount int16
	FirstTenMinutesTotalLevelBaseSessionCount  int16
	FirstQuarterHourTotalLevelBaseSessionCount int16
	FirstHalfHourTotalLevelBaseSessionCount    int16
	FirstHourTotalLevelBaseSessionCount        int16
	FirstTwoHourTotalLevelBaseSessionCount     int16
	FirstThreeHourTotalLevelBaseSessionCount   int16
	FirstSixHourTotalLevelBaseSessionCount     int16
	FirstTwelveHourTotalLevelBaseSessionCount  int16
	FirstDayTotalLevelBaseSessionCount         int16
	PenultimateLevelSessionLevelIndex          int16
	PenultimateLevelSessionLevelDuration       int16
	LastLevelSessionLevelIndex                 int16
	LastLevelSessionLevelDuration              int16
	LastLevelSessionYearOfDay                  int16
	LastLevelSessionYear                       int16
	LastLevelSessionWeekDay                    int16
	LastLevelSessionHour                       int16
	LastLevelSessionMinute                     int16
	Status                                     bool
}

type LevelBaseSessionModel struct {
	Id                int64
	ClientId          int64
	ProjectId         int64
	CustomerId        int64
	LevelName         string
	LevelIndex        int
	SessionTimeMinute float32
	SessionStartTime  time.Time
	SessionFinishTime time.Time
	Status            bool
}
