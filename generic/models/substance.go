package models

type Substance struct {
	ID                   uint
	Label                string
	Value                string
	RecommendedDosageMin float64
	RecommendedDosageMax float64
	RecommendedPauseMin  float64
	RecommendedPauseMax  float64
}
