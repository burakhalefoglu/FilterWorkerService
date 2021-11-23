package model

import "time"

type LevelBaseSessionRespondModel struct {
	ProjectId                                  string
	ClientId                                   string
	CustomerId                                 string
	TotalLevelBaseSessionMinute                int64
	TotalLevelBaseSessionCount                 int64
	FirstLevelSessionLevelIndex                int64
	FirstLevelSessionDuration                  int64
	FirstLevelSessionYearOfDay                 int64
	FirstLevelSessionYear                      int64
	FirstLevelSessionWeekDay                   int64
	FirstLevelSessionHour                      int64
	FirstLevelSessionMinute                    int64
	SecondLevelSessionLevelIndex               int64
	SecondLevelSessionDuration                 int64
	ThirdLevelSessionLevelIndex                int64
	ThirdLevelSessionDuration                  int64
	FourLevelSessionLevelIndex                 int64
	FourLevelSessionDuration                   int64
	FiveLevelSessionLevelIndex                 int64
	FiveLevelSessionDuration                   int64
	SixLevelSessionLevelIndex                  int64
	SixLevelSessionDuration                    int64
	SevenLevelSessionLevelIndex                int64
	SevenLevelSessionDuration                  int64
	FirstQuarterHourTotalLevelBaseSessionCount int64
	FirstHalfHourTotalLEvelBaseSessionCount    int64
	FirstHourTotalLevelBaseSessionCount        int64
	FirstTwoHourTotalLevelBaseSessionCount     int64
	FirstThreeHourTotalLevelBaseSessionCount   int64
	FirstSixHourTotalLevelBaseSessionCount     int64
	FirstTwelveHourTotalLevelBaseSessionCount  int64
	FirstDayTotalLevelBaseSessionCount         int64
	PenultimateLevelSessionLevelIndex          int64
	PenultimateLevelSessionLevelDuration       int64
	LastLevelSessionLevelIndex                 int64
	LastLevelSessionLevelDuration              int64
	LastLevelSessionYearOfDay                  int64
	LastLevelSessionYear                       int64
	LastLevelSessionWeekDay                    int64
	LastLevelSessionHour                       int64
	LastLevelSessionMinute                     int64
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
