package controllers

import (
	"assignment-3/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetStatus(ctx *gin.Context) {
	helpers.UpdateJSONData()

	data := helpers.GetDataJson()

	waterStatus, windStatus := helpers.GetWaterStatus(data.Status.Water), helpers.GetWindStatus(data.Status.Wind)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"waterValue":  data.Status.Water,
		"waterStatus": waterStatus,
		"waterClass":  strings.ToLower(waterStatus),
		"windValue":   data.Status.Wind,
		"windStatus":  windStatus,
		"windClass":   strings.ToLower(windStatus),
	})
}
