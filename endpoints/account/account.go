package accountRoutes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/amina-b/project/database/accountDB"
)

// u novi folder modules i importujes a ne svaki put pravis
type User struct {
	User_id  int // UserID
	Name     string
	Username string
	Password string
}

func HandleUsersRegister(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is register route")
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var user User
	err = json.Unmarshal(reqBytes, &user)

	if err != nil {
		fmt.Println(err)
		//w.writeheader - status - jer po defaultu vraca 200
	} else {
		fmt.Println(user)
		accountDB.SaveUserInTable((accountDB.User)(user)) //saveuser zna se da je tabela -- package sa githuba
	}

	return
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
