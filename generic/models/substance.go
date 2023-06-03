package models

type Substance struct {
	ID                       uint 
	Label				     string
	Value                    string
	RecommendedDosageMin     uint 
	RecommendedDosageMax     uint 
	RecommendedPauseMin 	 uint 
	RecommendedPauseMax      uint 
}