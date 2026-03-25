package controllers

import (
	databaseConfig "app_translator/config"
	"app_translator/models"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig(databaseConfig.DatabaseConnect())

	fmt.Println(config)

	tmpl := template.Must(template.ParseFiles("views/config.html"))
	tmpl.Execute(w, config)
}

func SetConfigHandler(w http.ResponseWriter, r *http.Request) {

	var req models.ConfigRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	models.SetConfig(req)

}
