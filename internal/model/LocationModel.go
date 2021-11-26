package model

type LocationResponseModel struct {
	ProjectId  string
	ClientId   string
	CustomerId string
	Continent  int64
	Country    int64
	City       int64
	Region     int64
	Org        int64
}

type LocationModel struct {
	ProjectId  string
	ClientId   string
	CustomerId string
	Continent  string
	Country    string
	City       string
	Region     string
	Org        string
}
