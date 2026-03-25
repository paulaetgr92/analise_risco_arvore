package model

import "time"

type WeatherEventRequest struct {
	Region    string  `json:"region"`
	WindSpeed float64 `json:"wind_speed"`
	Rainfall  bool    `json:"rainfall"`
}

type WeatherEventResponse struct {
	ID         int64     `json:"id"`
	Region     string    `json:"region"`
	WindSpeed  float64   `json:"wind_speed"`
	Rainfall   bool      `json:"rainfall"`
	OccurredAt time.Time `json:"occurred_at"`
}
