package controllers

import (
	"bufio"
	"encoding/json"
	"fmt"
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

func TranslateHandler(w http.ResponseWriter, r *http.Request) {

	t := translator.New()
	result, err := t.Translate("Hello, World!", "auto", "pt_br")
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// text := result.Text

	// tmpl := template.Must(template.ParseFiles("views/translate.html"))
	// tmpl.Execute(w, result)

}

func LoadBaseFile(w http.ResponseWriter, r *http.Request) {
	config := models.GetConfig()
	file, err := os.Open(config.InputFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		// panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var textsSlice []map[string]string

	for scanner.Scan() {
		row := scanner.Text()

		fmt.Println("linha:", row)

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

		// fmt.Println(texts)

		textsSlice = append(textsSlice, texts)

	}

	// fmt.Println(textsSlice)

	if err := scanner.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(textsSlice)

}
