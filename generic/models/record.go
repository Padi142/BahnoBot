package models

import (
	"time"
)

type Record struct {
	ID            uint 			`json:"id"`
	CreatedAt     time.Time		`json:"created_at"`
	Amount	      int			`json:"amount"`
	SubstanceID   uint			`json:"substance_id"`
	Substance 	  Substance		
	UserID 		  uint			`json:"user_id"`
	User 	      User
}