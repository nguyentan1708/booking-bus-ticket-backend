package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID          string `json:"id"`
	Account     string `json:"account"`
	Password    string `json:"password"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CardID      string `json:"card_id"`
}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func getAllUsers(w http.ResponseWriter, r *http.Request) {

}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}

func HandleUser(router *mux.Router) {
	router.HandleFunc("/api/v1/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/v1/users", getAllUsers).Methods("GET")
}
