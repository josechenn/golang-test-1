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

func ReturnPlayerByTeam(team_id *int) ([]model.Player, error) {
	var player []model.Player

	db := config.GormConnect()
	defer db.Close()

	db.Table("player").Where("team_id = ?", team_id).Find(&player)

	return player, nil
}

func ReturnPlayer(w http.ResponseWriter, r *http.Request) {
	var player []model.Player
	var result model.PlayerData

	team_id := r.URL.Query()["team_id"]
	player_name := r.URL.Query()["player_name"]
	player_id := r.URL.Query()["player_id"]

	db := config.GormConnect()
	defer db.Close()

	if team_id == nil && player_name == nil && player_id == nil {
		exception := "no team_id or player_name or player_id provided"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(exception)
	} else {
		query := db.Table("player")
		if team_id != nil {
			query = query.Where("team_id = ?", team_id)
		}
		if player_id != nil {
			query = query.Where("player_id = ?", player_id)
		}
		if player_name != nil {
			query = query.Where("player_name LIKE ?", "%"+player_name[0]+"%")
		}
		query.Find(&player)

		if len(player) == 0 {
			result := "no player found"
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(result)
		} else {
			data := model.PlayerData{
				Data: player,
			}

			result = data
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(result)
		}
	}
}

func InsertPlayer(w http.ResponseWriter, r *http.Request) {
	var player []*model.Player
	var team []*model.Team
	var result string
	db := config.GormConnect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	player_number, _ := strconv.Atoi(r.FormValue("player_number"))
	player_name := r.FormValue("player_name")
	position := r.FormValue("position")
	team_id := r.FormValue("team_id")

	if err := db.Table("player").Where("player_number = ?", player_number).Where("team_id = ?", team_id).Find(&player).Error; err != nil {
		panic(err)
	}
	if len(player) != 0 {
		result = "the player that you insert probably duplicate because there is player who use that number on the team"
	} else {
		db.Table("team").Select("team_name").Where("team_id = ?", team_id).Find(&team)
		if team[0] != nil {
			team_id, _ := strconv.Atoi(team_id)
			t := time.Now()
			player := model.Player{
				PlayerNumber: player_number,
				PlayerName:   player_name,
				Position:     position,
				TeamId:       team_id,
				Date:         t.String(),
			}
			db.Create(&player)
			result = "success"

		} else {
			result = "the team that you inserted is not exist"
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	var player []*model.Player
	var result string
	var old_team []*model.Team
	var new_team []*model.Team
	db := config.GormConnect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	player_number, _ := strconv.Atoi(r.FormValue("player_number"))
	player_id, _ := strconv.Atoi(r.FormValue("player_id"))
	player_name := r.FormValue("player_name")
	position := r.FormValue("position")
	team_id := r.FormValue("team_id")

	db.Table("player").Where("player_id = ?", player_id).Find(&player)
	if len(player) == 0 {
		result = "the player is not exist"
	} else {
		db.Table("team").Select("team_name").Where("team_id = ?", player[0].TeamId).Find(&old_team)
		team_id, _ := strconv.Atoi(team_id)
		db.Table("team").Select("team_name").Where("team_id = ?", team_id).Find(&new_team)
		if new_team[0] != nil {
			t := time.Now()
			player := model.Player{
				PlayerNumber: player_number,
				PlayerName:   player_name,
				Position:     position,
				TeamId:       team_id,
				Date:         t.String(),
			}
			db.Model(&player).Where("player_id = ?", player_id).Update(&player)
			if new_team[0].TeamName != old_team[0].TeamName {
				result = fmt.Sprintf("Success update player team from %s to %s", old_team[0].TeamName, new_team[0].TeamName)
			} else {
				result = "Success update player"
			}

		} else {
			result = "the team that you inserted is not exist"
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	var result string
	var deleted_player []*model.Player
	db := config.GormConnect()
	player_id, _ := strconv.Atoi(r.FormValue("player_id"))
	player := model.Player{
		PlayerId: player_id,
	}
	db.Table("player").Where("player_id = ?", player_id).Find(&deleted_player)
	if len(deleted_player) > 0 {
		db.Table("player").Where("player_id = ?", player_id).Delete(&player)
		result = "success delete player"
	} else {
		result = "the player that you want to delete is not exist"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
