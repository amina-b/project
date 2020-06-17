package accountRoutes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../../database/accountDB"
)

type User struct {
	User_id  int
	Name     string
	Username string
	Password string
}

func HandleUsersRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is register route")
	fmt.Println("OVO JE REQUEST ", r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var user User
	err = json.Unmarshal(reqBytes, &user)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
		accountDB.SaveUserInTable((accountDB.User)(user))
	}
}

func HandleUsersLogin(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "this is login route")
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	err = json.Unmarshal(reqBytes, &user)
	if err != nil {
		fmt.Println(err)
	} else {
		accountDB.Login((accountDB.User)(user))
	}
}
