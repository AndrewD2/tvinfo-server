package controllers

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/AndrewD2/tvinfo-server/models"
	"github.com/AndrewD2/tvinfo-server/views"
)

func NewSeries(ss models.SeriesService) *Series {
	return &Series{
		NewView:       views.NewView("bootstrap", "series/new"),
		SeriesService: ss,
	}
}

type Series struct {
	NewView *views.View
	models.SeriesService
}

func (s *Series) New(w http.ResponseWriter, r *http.Request) {
	s.NewView.Render(w, nil)
}

type SeriesForm struct {
	Name            string          `schema:"series-name"`
	Description     string          `schema:"series-desc"`
	StartDate       time.Time       `schema:"start-date"`
	EndDate         time.Time       `schema:"end-date"`
	Creator         []models.People `schema:"series-creator"`
	Actors          []models.People `schema:"actors"`
	EpisodeDuration uint            `schema:"duration"`
	Network         uint            `schema:"network"`
	Genre           []models.Genre  `schema:"genre"`
}

func (s *Series) Create(w http.ResponseWriter, r *http.Request) {
	form := SeriesForm{}
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
	}
	series := &models.Series{
		Name:            form.Name,
		Description:     form.Description,
		StartDate:       form.StartDate,
		EndDate:         form.EndDate,
		Creator:         form.Creator,
		Actors:          form.Actors,
		EpisodeDuration: form.EpisodeDuration,
		NetworkID:       form.Network,
		Genre:           form.Genre,
	}
	if err := s.SeriesService.Create(series); err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, series)
}
