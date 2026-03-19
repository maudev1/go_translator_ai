package main

import (
	"app_translator/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.SetupRoutes()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("rodando em http://localhost:8081")

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err)
	}
}
