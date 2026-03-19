package controllers

import (
	"html/template"
	"net/http"
	"app_translator/models"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig()

	tmpl := template.Must(template.ParseFiles("views/main.html"))
	tmpl.Execute(w, config)
}