package main

import (
	"app_translator/routes"
	"log"
	"net/http"
)

var port = "8081"

func main() {
	r := routes.SetupRoutes()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	log.Println("rodando em http://localhost:"+port)

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
