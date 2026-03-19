package routes

import (
	"app_translator/controllers"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", controllers.MainHandler)
	mux.HandleFunc("/config", controllers.ConfigHandler)
	mux.HandleFunc("/translate", controllers.TranslateHandler)
	mux.HandleFunc("/load-base-text", controllers.LoadBaseFile)
	return mux
}
