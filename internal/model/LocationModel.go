package model

type LocationResponseModel struct {
	ProjectId  string
	ClientId   string
	CustomerId string
	Continent  int16
	Country    int16
	City       int16
	Region     int16
	Org        int16
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
