package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Starting server on port 8080")
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	fmt.Println("Starting go api service")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error("Error starting server: ", err)
		return
	}

}
