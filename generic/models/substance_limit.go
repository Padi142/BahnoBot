package models

type SubstanceLimit struct {
	UserID      uint    `json:"user_id" gorm:"primary_key"`
	SubstanceID uint    `json:"substance_id" gorm:"primary_key"`
	DosageMin   float64 `json:"dosage_min"`
	DosageMax   float64 `json:"dosage_max"`
	PauseMin    float64 `json:"pause_min"`
	PauseMax    float64 `json:"pause_max"`
}
