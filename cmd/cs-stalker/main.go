package main

import (
	"fmt"
	"github.com/finnbechinka/cs-stalker/internal/routes"
	"log"
	"net/http"
)

func main() {
	log.Println("cs-stalker started")

	router := routes.NewRouter()
	port := ":8085"

	s := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println(fmt.Sprintf("server listening on http://localhost%s", port))
	log.Fatal(s.ListenAndServe())
}
