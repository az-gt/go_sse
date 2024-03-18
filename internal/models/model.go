package models

import "time"

type Record struct {
	ID     string `json:"id"`
	Data   Ride   `json:"data"`
	Action string `json:"action"`
}

type Ride struct {
	PassengerId string `json:"passenger_id"`
	Origin      struct {
		Long int `json:"long"`
		Lat  int `json:"lat"`
	} `json:"origin"`
	Destination struct {
		Long int `json:"long"`
		Lat  int `json:"lat"`
	} `json:"destination"`
	DepartureTime  time.Time `json:"departure_time"`
	ArrivalTime    time.Time `json:"arrival_time"`
	Fare           int       `json:"fare"`
	Distance       float64   `json:"distance"`
	Duration       int       `json:"duration"`
	VehicleType    string    `json:"vehicle_type"`
	RecomendedFare int       `json:"recomended_fare"`
	Id             string    `json:"id"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedAt      time.Time `json:"created_at"`
}
