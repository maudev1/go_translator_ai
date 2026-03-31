package controllers

import (
	databaseConfig "app_translator/config"
	"app_translator/models"
	"encoding/json"
	"html/template"
	"net/http"
)

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig(databaseConfig.DatabaseConnect())

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(config)
}

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig(databaseConfig.DatabaseConnect())

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

func SetBaseFile(w http.ResponseWriter, r *http.Request) {

	// file
	file, handler, err := r.FormFile("baseFile")

	if err != nil {
		http.Error(w, "Failed to get base file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	models.SetBaseFileConfig("files/input/" + handler.Filename)

}
