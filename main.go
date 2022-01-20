package main

import (
	"fmt"
	"log"
	"myapp/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
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

//
//func router() *mux.Router {
//	router := mux.NewRouter().StrictSlash(true)
//	return router
//}
