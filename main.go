package main

import (
	"log"
	"net/http"

	"./endpoints/accountRoutes"
)

func main() {

	http.HandleFunc("/users/register", accountRoutes.HandleUsersRegister)

	http.HandleFunc("/users/login", accountRoutes.HandleUsersLogin)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
