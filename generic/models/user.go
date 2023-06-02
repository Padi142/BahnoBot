package models

type User struct {
	ID                    uint
	Username			  string
	PreferredSubstanceID  string
	DiscordID  			  string
	Records 			  []Record	`gorm:"many2many:records;"`
}