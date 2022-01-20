package model

type Player struct {
	PlayerId     int    `form:"player_id" json:"player_id"`
	PlayerNumber int    `form:"player_number" json:"player_number"`
	PlayerName   string `form:"player_name" json:"player_name"`
	TeamId       int    `form:"team_id" json:"team_id"`
	Date         string `form:"date" json:"date"`
}

type PlayerData struct {
	Data []Player
}

func (Player) TableName() string {
	return "player"
}
