package controllers

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"app_translator/models"

	databaseConfig "app_translator/config"

	translator "github.com/Conight/go-googletrans"

	groq "app_translator/config"
)

type TextKey struct {
	Key   string
	Value string
}

type TranslateRequest struct {
	Engine string `json:"engine"`
	Text   string `json:"text"`
}

type TranslateReponse struct {
	Engine     string `json:"engine"`
	Original   string `json:"original"`
	Translated string `json:"translated"`
	Source     string `json:"source"`
	Target     string `json:"target"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig(databaseConfig.DatabaseConnect())

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

		resp := TranslateReponse{
			Engine:     req.Engine,
			Original:   result.Origin,
			Translated: result.Text,
			Source:     result.Src,
			Target:     result.Dest,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	if req.Engine == "ai" {
		response := groq.Chat(req.Text, config.GroqToken)

		var data map[string]interface{}
		err = json.Unmarshal([]byte(response), &data)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		var translated = data["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

		resp := TranslateReponse{
			Engine:     "ai",
			Original:   req.Text,
			Translated: translated,
			Source:     "",
			Target:     "",
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}

	http.Error(w, "engine não suportado", 400)
}

func LoadBaseFile(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig(databaseConfig.DatabaseConnect())
	file, err := os.Open(config.BaseFile)
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
