package accountDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	User_id  int
	Name     string
	Username string
	Password string
}

//enviroment varijable-- connection on DB
var driverName string = "mysql"
var connection string = "root:sifra123@@@@/"

func ConnectToDb() (db *sql.DB, err error) {

	db, err = sql.Open(driverName, connection)

	if err != nil {
		return nil, err
	} else {
		_, err = db.Exec("USE qamazing")
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB selected successfuly")
		}
		return db, nil
	}

}

func SaveUserInTable(user User) {

	db, dbErr := ConnectToDb()

	if dbErr == nil {
		_, err := db.Exec(`INSERT INTO users(user_id, name, username, password) VALUES(?, ?, ?, ?)`, user.User_id, user.Name, user.Username, user.Password)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(user)
			fmt.Println("Inserted User with username: " + user.Username)
		}
		defer db.Close()
	} else {
		fmt.Println(dbErr.Error())
		defer db.Close()
	}
}

func FindUserInTable(user User) (foundUser User, err error) {
	db, dbErr := ConnectToDb()
	nul := User{0, "", "", ""}

	if dbErr == nil {
		err := db.QueryRow("SELECT * FROM users WHERE username = ?", user.Username).Scan(&foundUser.User_id, &foundUser.Name, &foundUser.Username, &foundUser.Password)
		if err != nil {
			fmt.Println(err.Error())
			return nul, err
		} else {
			fmt.Println("nadjeni user ", foundUser)
			return foundUser, nil
		}
	}
	defer db.Close()
	return nul, dbErr
}

func Login(user User) {

	foundUser, err := FindUserInTable(user)
	if err != nil {
		fmt.Println(err)
	} else {
		if foundUser.Username == user.Username && foundUser.Password == user.Password {
			fmt.Println("uspjesan login")
		} else {
			fmt.Println("neki podatak nije tacan")
		}
	}
}
