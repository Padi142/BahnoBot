package models

type Substance struct {
	Id                string `bson:"id"`
	Name              string `bson:"name"`
	RecommendedDosage string `bson:"recommended_dosage"`
	RecommendedPause  string `bson:"recommended_pause"`
}
