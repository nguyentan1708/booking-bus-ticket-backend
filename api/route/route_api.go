package route

import (
	"booking-bus-ticket/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	ID         string `json:"id"`
	StartPoint string `json:"start_point"`
	EndPoint   string `json:"end_point"`
	Date       string `json:"date"`
}

func getRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var route Route
	db := config.Connect()
	defer db.Close()
	parameter := mux.Vars(r)["id"]
	result, err := db.Query("select * from route where id = ?", parameter)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&route.ID, &route.StartPoint, &route.EndPoint, &route.Date)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(route)
}

func getAllRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var routes []Route
	db := config.Connect()
	defer db.Close()
	result, err := db.Query("select * from route")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var route Route
		err := result.Scan(&route.ID, &route.StartPoint, &route.EndPoint, &route.Date)
		if err != nil {
			panic(err.Error())
		}
		routes = append(routes, route)
	}
	json.NewEncoder(w).Encode(routes)
}

func createRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("insert into route (id, start_point, end_point, date) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyValue := make(map[string]string)
	json.Unmarshal(body, &keyValue)
	idRoute := keyValue["id"]
	startPoint := keyValue["start_point"]
	endPoint := keyValue["end_point"]
	date := keyValue["date"]
	_, err = stmt.Exec(idRoute, startPoint, endPoint, date)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "You have successfully added a route!")
}

func updateRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)["id"]
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("update route set id=?, start_point=?, end_point=?, date=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyValue := make(map[string]string)
	json.Unmarshal(body, &keyValue)
	id := keyValue["id"]
	startPoint := keyValue["start_point"]
	endPoint := keyValue["end_point"]
	date := keyValue["date"]
	_, err = stmt.Exec(id, startPoint, endPoint, date, parameter)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Route have id = %s have been updated", parameter)
}

func deleteRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parameter := mux.Vars(r)["id"]
	db := config.Connect()
	defer db.Close()
	stmt, err := db.Prepare("delete from route where id = %s")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(parameter)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Route has id = %s has been deleted", parameter)
}

func HandleRoute(router *mux.Router) {
	router.HandleFunc("/api/v1/route/{id}", getRoute).Methods("GET")
	router.HandleFunc("/api/v1/routes", getAllRoutes).Methods("GET")
	router.HandleFunc("/api/v1/route", createRoute).Methods("POST")
	router.HandleFunc("/api/v1/route/{id}", updateRoute).Methods("PATCH")
	router.HandleFunc("/api/v1/route/{id}", deleteRoute).Methods("DELETE")
}
