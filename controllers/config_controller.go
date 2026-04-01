package controllers

import (
	databaseConfig "app_translator/config"
	"app_translator/models"
	"encoding/json"
	"html/template"
	"net/http"
)

type Error struct {
	Message string
	Code    int
}

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

	config := models.SetConfig(req)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(config)

}

func SetBaseFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	r.ParseMultipartForm(10 << 20)

	// file
	file, handler, err := r.FormFile("baseFile")

	if err != nil {
		http.Error(w, "Failed to set base file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	if handler.Header.Get("Content-Type") != "application/octet-stream" {
		http.Error(w, "Failed to set base file: Incompatible File", http.StatusBadRequest)

		// error := &Error{
		// 	Message: "Incompatible File",
		// 	Code:    404,
		// }

		// e, err := json.Marshal(error)
		// if err != nil {
		// 	fmt.Printf(err.Error())
		// }

		// json.NewEncoder(w).Encode(error).Error("sdfsd")

		// json.NewEncoder(w).Encode(error)

		return

	}

	config := models.SetBaseFileConfig("files/input/" + handler.Filename)

	json.NewEncoder(w).Encode(config)

}
