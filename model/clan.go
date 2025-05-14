package model

type Clan struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	ClanTags     string `json:"clan_tags" gorm:"uniqueIndex"`
	ClanName     string `json:"clan_name"`
	ClanType     string `json:"clan_type"`
	ClanLocation string `json:"clan_location"`
}
