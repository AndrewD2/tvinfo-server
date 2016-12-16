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
	password = ""
	dbname   = "tvinfo"
)

func main() {
	// Create a DB connection string and then use it to
	// create our model services.
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	all, err := models.AllModels(psqlInfo)
	if err != nil {
		log.Println(err)
	}

	ug, ig, pg, rg, gg, ng, mg, sg, seg, eg := all.UserGorm, all.ImageGorm, all.PeopleGorm, all.RoleGorm, all.GenreGorm,
		all.NetworkGorm, all.MacroSeriesGorm, all.SeriesGorm, all.SeasonGorm, all.EpisodeGorm

	destructiveReset(ug, ig, pg, rg, gg, ng, mg, sg, seg, eg)

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(ug)
	episodeC := controllers.NewEpisode(eg)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods(http.MethodGet)
	r.Handle("/contact", staticC.Contact).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.New).Methods(http.MethodGet)
	r.HandleFunc("/signup", usersC.Create).Methods(http.MethodPost)
	r.Handle("/login", usersC.LoginView).Methods(http.MethodGet)
	r.HandleFunc("/login", usersC.Login).Methods(http.MethodPost)
	r.Handle("/faq", staticC.Faq).Methods(http.MethodGet)
	r.HandleFunc("/episode/new", episodeC.New).Methods(http.MethodGet)
	r.HandleFunc("/episode/new", episodeC.Create).Methods(http.MethodPost)
	http.ListenAndServe(":3000", r)
}

func destructiveReset(all ...interface {
	models.DBReset
}) {

	ug, ig, pg, rg, gg, ng, mg, sg, seg, eg := all[0], all[1], all[2], all[3], all[4], all[5], all[6], all[7], all[8], all[9]
	ug.DestructiveReset()
	eg.DestructiveReset()
	mg.DestructiveReset()
	sg.DestructiveReset()
	seg.DestructiveReset()
	ig.DestructiveReset()
	pg.DestructiveReset()
	gg.DestructiveReset()
	ng.DestructiveReset()
	rg.DestructiveReset()
}

func autoMigrate(all ...interface {
	models.DBReset
}) {
	ug, ig, pg, rg, gg, ng, mg, sg, seg, eg := all[0], all[1], all[2], all[3], all[4], all[5], all[6], all[7], all[8], all[9]

	ug.AutoMigrate()
	eg.AutoMigrate()
	mg.AutoMigrate()
	sg.AutoMigrate()
	seg.AutoMigrate()
	ig.AutoMigrate()
	pg.AutoMigrate()
	gg.AutoMigrate()
	ng.AutoMigrate()
	rg.AutoMigrate()
}
