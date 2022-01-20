package service

import (
	"encoding/json"
	"fmt"
	"myapp/config"
	"myapp/model"
	"net/http"
	"strconv"
	"time"
)

func ReturnTeam(w http.ResponseWriter, r *http.Request) {
	var teams []model.Team
	var result []model.Team
	var page int

	db := config.GormConnect()
	defer db.Close()
	team_id := r.URL.Query()["team_id"]
	team_name := r.URL.Query()["team_name"]
	pages := r.URL.Query()["page"]
	if pages != nil {
		page, _ = strconv.Atoi(pages[0])
	} else {
		page = 1
	}
	if team_id != nil {
		db.Table("team").Where("team_id = ?", team_id[0]).Find(&teams)
	} else if team_name != nil {
		db.Table("team").Where("team_id = ?", "%"+team_name[0]+"%").Find(&teams)
	} else {
		db.Table("team").Find(&teams).Limit(5).Offset((page - 1) * 3)
	}
	for _, val := range teams {
		val.Player, _ = ReturnPlayerByTeam(&val.TeamId)
		result = append(result, val)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func ReturnTeamDetail(w http.ResponseWriter, r *http.Request) {
	var team model.Team
	var result model.TeamDetail

	db := config.GormConnect()
	defer db.Close()
	team_id := r.URL.Query()["team_id"]
	team_name := r.URL.Query()["team_name"]
	fmt.Println(team_id)

	if team_id == nil && team_name == nil {
		result := "No team id or team name provided"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	} else if team_id != nil && team_name != nil {
		result := "You should provide only one either team id or team name, could not provide both"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	} else {
		query := db.Table("team")
		id, _ := strconv.Atoi(team_id[0])

		if team_id != nil {
			query = query.Where("team_id = ?", id)
		}

		if team_name != nil {
			query = query.Where("team_name = ?", team_name[0])
		}
		query.First(&team)
		if team.TeamName == "" {
			result := "can not find team"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(result)
		} else {
			result.Player, _ = ReturnPlayerByTeam(&id)
			result.TeamId = team.TeamId
			result.TeamName = team.TeamName
			result.Date = team.Date
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
		}
	}
}

func InsertTeam(w http.ResponseWriter, r *http.Request) {
	var team []*model.Team
	db := config.GormConnect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	team_id, _ := strconv.Atoi(r.FormValue("team_id"))
	team_name := r.FormValue("team_name")
	db.Table("team").Where("team_name = ?", team_name).First(&team)
	if len(team) > 0 {
		result := "the team name that you inserted has been used, please insert another name"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	} else {
		t := time.Now()
		team := model.Team{
			TeamId:   team_id,
			TeamName: team_name,
			Date:     t.String(),
		}
		db.Create(&team)
		result := "success"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
