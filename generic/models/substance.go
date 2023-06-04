package models

type Substance struct {
	ID                   uint   `json:"id"`
	Label                string `json:"label"`
	Value                string `json:"value"`
	RecommendedDosageMin uint   `json:"recommended_dosage_min"`
	RecommendedDosageMax uint   `json:"recommended_dosage_max"`
	RecommendedPauseMin  uint   `json:"recommended_pause_min"`
	RecommendedPauseMax  uint   `json:"recommended_pause_max"`
}
