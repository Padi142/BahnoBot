package models

type Substance struct {
	Id                string `bson:"id"`
	Name              string `bson:"name"`
	Value             string `bson:"value"`
	RecommendedDosage string `bson:"recommended_dosage"`
	RecommendedPause  string `bson:"recommended_pause"`
}

func GetSubstances() []*Substance {
	return []*Substance{
		{
			Id:                "1",
			Name:              "Bahno",
			Value:             "bahno",
			RecommendedDosage: "0.5-3 g",
			RecommendedPause:  "1-2 days",
		},
		{
			Id:                "2",
			Name:              "Maté",
			Value:             "mate",
			RecommendedDosage: "10-50 g",
			RecommendedPause:  "1-2 days",
		},
		{
			Id:                "3",
			Name:              "Ďáblovo zelí",
			Value:             "zeli",
			RecommendedDosage: "0.1-0.5 g",
			RecommendedPause:  "4-7 days",
		},
		{
			Id:                "4",
			Name:              "Nikotin",
			Value:             "nikotin",
			RecommendedDosage: "0.1-0.5 g",
			RecommendedPause:  "idk",
		},
		{
			Id:                "5",
			Name:              "Kofein",
			Value:             "kofein",
			RecommendedDosage: "0.1-0.5 g",
			RecommendedPause:  "1 day",
		},
	}
}
