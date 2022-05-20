package main

import (
	"booking-bus-ticket/api/bus"
	"booking-bus-ticket/api/route"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	router := mux.NewRouter()
	bus.HandleBus(router)
	route.HandleRoute(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
