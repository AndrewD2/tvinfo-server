package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndrewD2/tvinfo-server/controllers"
	"github.com/AndrewD2/tvinfo-server/models"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "{placeholder}"
	dbname   = "tvinfo"
)

func main() {
	// Create a DB connection string and then use it to
	// create our model services.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	ug, err := models.NewUserGorm(psqlInfo)
	if err != nil {
		log.Println(err)
	}
	ug.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(ug)
	// episodeC := controllers.NewEpisode()

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	r.Handle("/login", usersC.LoginView).Methods("GET")
	r.HandleFunc("/login", usersC.Login).Methods("POST")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	http.ListenAndServe(":3000", r)
}
