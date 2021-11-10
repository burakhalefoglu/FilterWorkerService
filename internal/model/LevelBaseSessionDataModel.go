package model

import "time"

type LevelBaseSessionRespondModel struct{
	ProjectId         string
	ClientId          string
	CustomerId        string
	LevelIndex        int64

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
