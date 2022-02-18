package model

import "time"

type LocationResponseModel struct {
	Id         int64
	ClientId   int64
	ProjectId  int64
	CustomerId int64
	Continent  int16
	Country    int16
	City       int16
	Region     int16
	Org        int16
	Status     bool
}

type LocationModel struct {
	Id         int64
	ClientId   int64
	ProjectId  int64
	CustomerId int64
	Continent  string
	Country    string
	City       string
	Query      string
	Region     string
	Org        string
	CreatedAt  time.Time
	Status     bool
}