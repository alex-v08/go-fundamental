package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var users []User
var MaxID uint64


func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status":%d, "data":%s}`, status, value)
}

func init() {
	users = []User{
		{ID: 1, FirstName: "John", LastName: "Doe", Email: ""},
		{ID: 2, FirstName: "Jane", LastName: "Doe", Email: ""},
		{ID: 3, FirstName: "John", LastName: "Smith", Email: ""},
	}
	MaxID = 3


}

func main() {
	http.HandleFunc("/users", UserServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllUsers(w)
	case http.MethodPost:
		SaveUser(w, r)
	default:
		MsgResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return

	}

}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	var user User


	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, "Invalid Request")
		return
	}
    MaxID++
	user.ID = MaxID
	users = append(users, user)
	DataResponse(w, http.StatusCreated, user)
}


func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status":%d, "message":"%s"}`, status, message)
}
