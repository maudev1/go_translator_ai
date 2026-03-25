package controllers

import (
	databaseConfig "app_translator/config"
	"app_translator/models"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig(databaseConfig.DatabaseConnect())

	tmpl := template.Must(template.ParseFiles("views/main.html"))
	tmpl.Execute(w, config)
}
