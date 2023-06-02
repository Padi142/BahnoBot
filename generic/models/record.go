package models

import (
	"time"
)

type Record struct {
	ID            uint
	Time          time.Time
	Amount	      int
	SubstanceID   uint
	UserID 		  uint

}