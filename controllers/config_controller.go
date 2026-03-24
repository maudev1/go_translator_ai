package controllers

import (
	"app_translator/models"
	"fmt"
	"html/template"
	"net/http"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig()

	fmt.Println(config)

	tmpl := template.Must(template.ParseFiles("views/config.html"))
	tmpl.Execute(w, config)
}
