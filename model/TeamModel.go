package model

type Team struct {
	TeamId   int    `form:"team_id" json:"team_id"`
	TeamName string `form:"team_name" json:"team_name"`
	Date     string `form:"date" json:"date"`
	Player   []Player
}

type TeamData struct {
	Data []Team
}

type TeamDetail struct {
	TeamId   int    `form:"team_id" json:"team_id"`
	TeamName string `form:"team_name" json:"team_name"`
	Date     string `form:"date" json:"date"`
	Player   []Player
}

type TeamDetailData struct {
	Data TeamDetail
}

func (Team) TableName() string {
	return "team"
}
