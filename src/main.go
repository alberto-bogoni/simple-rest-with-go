package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Surname string `json:"surname"`
}

var users = []user{
    {"1", "Alberto", "Rossi"},
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var newUser user

    json.NewDecoder(r.Body).Decode(&newUser)
    users = append(users, newUser)

    w.WriteHeader(http.StatusCreated)
    fmt.Printf("Created user %s", newUser.Name)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Get Users!")

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("POST /users", createUser)
	http.HandleFunc("GET /users", getUsers)
    http.ListenAndServe(":8080", nil)
}
