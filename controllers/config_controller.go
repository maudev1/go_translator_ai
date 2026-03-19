package controllers

import (
	"html/template"
	"net/http"
	"app_translator/models"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig()

	tmpl := template.Must(template.ParseFiles("views/config.html"))
	tmpl.Execute(w, config)
}