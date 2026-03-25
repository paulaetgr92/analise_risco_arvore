package model

import (
	"time"
)

type RiskAssesmentRequest struct {
	TreeID int32  `json:"tree_id"`
	Score  int32  `json:"score"`
	Level  string `json:"level"`
}
type RiskAssessmentResponse struct {
	ID           int64     `json:"id"`
	TreeID       int64     `json:"tree_id"`
	Score        int32     `json:"score"`
	Level        string    `json:"level"`
	CalculatedAt time.Time `json:"calculated_at"`
}

type HighRiskTreeResponse struct {
	TreeID    int64     `json:"tree_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Species   string    `json:"species,omitempty"`
	Height    float64   `json:"height,omitempty"`
	Diameter  float64   `json:"diameter,omitempty"`
	Age       int32     `json:"age,omitempty"`
	Health    string    `json:"health,omitempty"`
	Status    string    `json:"status,omitempty"`
	Score     int32     `json:"score"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type MediumOrHighRiskTreeResponse struct {
	TreeID    int64     `json:"tree_id"`
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Species   string    `json:"species,omitempty"`
	Height    float64   `json:"height,omitempty"`
	Diameter  float64   `json:"diameter,omitempty"`
	Age       int32     `json:"age,omitempty"`
	Health    string    `json:"health,omitempty"`
	Status    string    `json:"status,omitempty"`
	Score     int32     `json:"score"`
	Level     string    `json:"level"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
