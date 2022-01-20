package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"myapp/seeds"
	"myapp/service"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	HandleArgs()
	router := mux.NewRouter().StrictSlash(true)

	//Player Route
	//route to get player, we can add query param such as team_id to get a player who is in the team x,
	//or we can add player name to get a player who has that name, or we can add player_id to get specify player,
	//and if there're no query param it will get all player
	router.HandleFunc("/get_player", service.ReturnPlayer).Methods("GET")
	//route to insert new player, and whenever a new player is inserted, it will check if there're same player_number on the team, because in soccer player_number should be unique
	router.HandleFunc("/insert_player", service.InsertPlayer).Methods("POST")
	//route to update player, if a player change team it will return message to inform us from where to where the player update
	router.HandleFunc("/update_player", service.UpdatePlayer).Methods("PUT")
	//route to delete player
	router.HandleFunc("/delete_player", service.DeletePlayer).Methods("DELETE")

	//Team Route
	//route to get all team, and player list already included to this route
	router.HandleFunc("/get_all_team", service.ReturnTeam).Methods("GET")
	//route to get team detail by giving the team id
	router.HandleFunc("/get_team_detail", service.ReturnTeamDetail).Methods("GET")
	//route to insert new team, same name team could not be stored, because team name in soccer should be unique (small similaraty is acceptable)
	router.HandleFunc("/insert_team", service.InsertTeam).Methods("POST")

	router.HandleFunc("/count_total", service.CountTotalBox).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}

func HandleArgs() {
	flag.Parse()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			connString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&multiStatements=true", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
			// connect DB
			db, err := sql.Open("mysql", connString)
			if err != nil {
				log.Fatalf("Error opening DB: %v", err)
			}
			seeds.Execute(db, args[1:]...)
			os.Exit(0)
		}
	}
}
