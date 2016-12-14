package controllers

import "github.com/AndrewD2/tvinfo-server/models"

func NewSeries(ss models.SeriesService) *Series {
	return &Series{
		SeriesService: ss,
	}
}

type Series struct {
	models.SeriesService
}
