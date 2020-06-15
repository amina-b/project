package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("connect to mysql")

	db, err := sql.Open("mysql", "root:sifra123@@@@/qamazing")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected")

	defer db.Close()

	
}