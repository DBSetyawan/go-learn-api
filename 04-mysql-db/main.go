package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jacky-htg/api-go/04-mysql-db/config"
	"github.com/jacky-htg/api-go/04-mysql-db/controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/users", controllers.GetUsersEndPoint).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUserEndPoint).Methods("POST")
	router.HandleFunc("/api/users/{id}", controllers.GetUserEndPoint).Methods("GET")

	http.ListenAndServe(config.GetString("server.address"), router)
}
