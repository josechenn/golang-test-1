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
	router.HandleFunc("/get_player", service.ReturnPlayer).Methods("GET")
	router.HandleFunc("/insert_player", service.InsertPlayer).Methods("POST")
	router.HandleFunc("/update_player", service.UpdatePlayer).Methods("PUT")
	router.HandleFunc("/delete_player", service.DeletePlayer).Methods("DELETE")

	//Team Route
	router.HandleFunc("/get_all_team", service.ReturnTeam).Methods("GET")
	router.HandleFunc("/get_team_detail", service.ReturnTeamDetail).Methods("GET")
	router.HandleFunc("/insert_team", service.InsertTeam).Methods("POST")

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
