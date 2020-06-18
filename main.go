package main

import (
	"log"
	"net/http"

	"./config"

	"./endpoints/accountRoutes"
	"github.com/amina-b/project/config/config.go"
)

func main() {
	config.Initialize()

	http.HandleFunc("/users/register", accountRoutes.HandleUsersRegister)

	http.HandleFunc("/users/login", accountRoutes.HandleUsersLogin)

	log.Fatal(http.ListenAndServe(":"+config.Config.Port, nil))
}
