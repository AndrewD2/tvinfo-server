package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AndrewD2/tvinfo-server/models"
	"github.com/AndrewD2/tvinfo-server/views"
)

func NewEpisode(es models.EpisodeService) *Episode {
	return &Episode{
		NewView:        views.NewView("bootstrap", "episodes/new"),
		EpisodeService: es,
	}
}

type Episode struct {
	NewView *views.View
	models.EpisodeService
}

func (e *Episode) New(w http.ResponseWriter, r *http.Request) {
	e.NewView.Render(w, nil)
}

type EpisodeForm struct {
	EpisodeNumber  uint   `schema:"episode-num"`
	Title          string `schema:"title"`
	FirstAired     string `schema:"first-aired"`
	ProductionCode string `schema:"prod-code"`
	Description    string `schema:"description"`
}

func (e *Episode) Create(w http.ResponseWriter, r *http.Request) {
	form := EpisodeForm{}
	if err := parseForm(r, &form); err != nil {
		log.Println(err)
	}
	episode := &models.Episode{
		EpisodeNumber: form.EpisodeNumber,
		Title:         form.Title,
		FirstAired:    form.FirstAired,
		//ProductionCode: form.ProductionCode,
		Description: form.Description,
	}
	if err := e.EpisodeService.Create(episode); err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, episode)
}
