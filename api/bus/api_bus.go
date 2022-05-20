package bus

import (
	"booking-bus-ticket/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Bus struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	SeatNumber   int    `json:"seat_number"`
	FloorsNumber int    `json:"floors_number"`
}

func getBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var bus Bus
	db := config.Connect()
	defer db.Close()
	parameter := mux.Vars(r)["id"]
	result, err := db.Query("select * from bus where id=?", parameter)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&bus.ID, &bus.Type, &bus.SeatNumber, &bus.FloorsNumber)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(bus)
}

func getAllBuses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var buses []Bus
	db := config.Connect()
	defer db.Close()
	result, err := db.Query("select * from bus")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var bus Bus
		err := result.Scan(&bus.ID, &bus.Type, &bus.SeatNumber, &bus.FloorsNumber)
		if err != nil {
			panic(err.Error())
		}
		buses = append(buses, bus)
	}
	json.NewEncoder(w).Encode(buses)
}

func createBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("insert into bus(id, type, seat_number, floors_number) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idBus := keyVal["id"]
	typeBus := keyVal["type"]
	seatNumber := keyVal["seat_number"]
	floorsNumber := keyVal["floors_number"]
	_, err = stmt.Exec(idBus, typeBus, seatNumber, floorsNumber)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Inserted a new bus")
}

func updateBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)["id"]
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("update bus set id=?,type=?,seat_number=?,floors_number=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	idBus := keyVal["id"]
	typeBus := keyVal["type"]
	seatNumber := keyVal["seat_number"]
	floorsNumber := keyVal["floors_number"]
	_, err = stmt.Exec(idBus, typeBus, seatNumber, floorsNumber, parameter)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Bus with id = %s has been updated", parameter)
}

func deleteBus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)["id"]
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("delete from bus where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(parameter)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Bus with id = %s has been deleted", parameter)
}

func HandleBus(router *mux.Router) {
	router.HandleFunc("/api/v1/bus/{id}", getBus).Methods("GET")
	router.HandleFunc("/api/v1/buses", getAllBuses).Methods("GET")
	router.HandleFunc("/api/v1/bus", createBus).Methods("POST")
	router.HandleFunc("/api/v1/bus/{id}", updateBus).Methods("PATCH")
	router.HandleFunc("/api/v1/bus/{id}", deleteBus).Methods("DELETE")
}
