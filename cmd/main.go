package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"timestamps/app/services"
	"timestamps/persistence"
	"timestamps/presentation"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	serveOnPort := getEnv("SERVE_ON_PORT", "8000")

	// router
	r := mux.NewRouter()

	timestampsRepository := persistence.NewTimestampsRepository()
	timestampsService := services.NewTimestampsService(timestampsRepository)

	timestampsRoutes := presentation.CreateRoutes(timestampsService)
	for routePath, routeMethods := range timestampsRoutes {
		fmt.Printf("adding %s route with methods %v\n", routePath, routeMethods.Methods)
		r.Handle(routePath, routeMethods.HandlerFunc).Methods(routeMethods.Methods...)
	}

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+serveOnPort, r))
}
