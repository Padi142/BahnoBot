package models

import (
	"time"
)

type Record struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Amount      float64   `json:"amount"`
	SubstanceID uint      `json:"substance_id"`
	Substance   Substance `json:"substance"`
	UserID      uint      `json:"user_id"`
	User        User      `json:"user"`
}
