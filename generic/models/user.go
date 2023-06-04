package models

type User struct {
	ID                   uint      `json:"id"`
	Username             string    `json:"username"`
	PreferredSubstanceID uint      `json:"preferred_substance_id"`
	PreferredSubstance   Substance `json:"preferred_substance"`
	DiscordID            string    `json:"discord_id"`
	Records              []Record  `gorm:"many2many:records;" json:"records"`
}
