package controllers

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"app_translator/models"

	translator "github.com/Conight/go-googletrans"
)

type TextKey struct {
	Key   string
	Value string
}

type TranslateRequest struct {
	Engine string `json:"engine"`
	Text   string `json:"text"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig()

	var req TranslateRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "JSON invalido", http.StatusBadRequest)
		return
	}

	if req.Engine == "google" {
		t := translator.New()

		result, err := t.Translate(req.Text, "auto", config.Language)
		if err != nil {
			http.Error(w, "erro na tradução", 500)
			return
		}

		resp := map[string]interface{}{
			"engine":     req.Engine,
			"original":   result.Origin,
			"translated": result.Text,
			"source":     result.Src,
			"target":     result.Dest,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	http.Error(w, "engine não suportado", 400)
}

func LoadBaseFile(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig()
	file, err := os.Open(config.InputFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var textsSlice []map[string]string

	for scanner.Scan() {
		row := scanner.Text()

		parts := strings.SplitN(row, "=", 2)

		if len(parts) < 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		texts := map[string]string{
			"key":   key,
			"value": value,
		}

		textsSlice = append(textsSlice, texts)

	}

	if err := scanner.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(textsSlice)

}
