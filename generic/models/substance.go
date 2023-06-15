package models

type Substance struct {
	ID                   uint    `json:"id"`
	Label                string  `json:"label"`
	Value                string  `json:"value"`
	Description          string  `json:"description"`
	RecommendedDosageMin float64 `json:"recommended_dosage_min"`
	RecommendedDosageMax float64 `json:"recommended_dosage_max"`
	RecommendedPauseMin  float64 `json:"recommended_pause_min"`
	RecommendedPauseMax  float64 `json:"recommended_pause_max"`
}
