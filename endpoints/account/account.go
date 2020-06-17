package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../../database/account"
)

type User struct {
	User_id  int
	Name     string
	Username string
	Password string
}

func main() {
	http.HandleFunc("/users/register", HandleUsersRegister)

	http.HandleFunc("/users/login", HandleUsersLogin)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HandleUsersRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is register route")
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var User User
	err = json.Unmarshal(reqBytes, &User)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(User)
		account.SaveUserInTable((account.User)(User))
	}
}

func HandleUsersLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "this is login route")
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var User User
	err = json.Unmarshal(reqBytes, &User)
	if err != nil {
		fmt.Println(err)
	} else {
		account.Login((account.User)(User))
	}

}
