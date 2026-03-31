package model

import (
	"time"
)

type Tree struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Height    float64 `json:"height"`
	Species   string  `json:"species"`
}

type TreeResponse struct {
	Id        int64     `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"` // ✅ CORRIGIDO
	Species   string    `json:"species"`
	Height    float64   `json:"height"`
	Diameter  float64   `json:"diameter"`
	Age       int       `json:"age"`
	Health    string    `json:"health"`
	CreatedAt time.Time `json:"created_at"`
}

type ListTreesByBoundingBoxRequest struct {
	Id          int64   `json:"id"`
	Latitude    float64 `json:"latitude"`
	Latitude_2  float64 `json:"latitude_2"`
	Longitude   float64 `json:"longitude"`
	Longitude_2 float64 `json:"longitude_2"`
	Species     string  `json:"species"`
}
type ListTreesByBoundingBoxResponse struct {
	Id        int64   `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Height    float64 `json:"height"`
	Species   string  `json:"species"`
}
type UpdateTreeRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Species   string  `json:"species"`
	ID        int64   `json:"id"`
}
type UpdateTreeResponse struct {
	ID        int64     `json:"id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Species   string    `json:"species"`
	Height    float64   `json:"height"`
	Diameter  float64   `json:"diameter"`
	Age       int32     `json:"age"`
	Health    string    `json:"health"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
