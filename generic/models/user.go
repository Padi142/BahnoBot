package models

type User struct {
	ID                   uint
	Username             string
	PreferredSubstanceID uint
	PreferredSubstance   Substance
	DiscordID            string
	Records              []Record `gorm:"many2many:records;"`
}
