package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Usersdata struct {
	ID              int
	CitizenID       string
	Fristname       string
	Lastname        string
	BirthYear       int
	FristnameFather string
	LastnameFather  string
	FristnameMother string
	LastnameMother  string
	SoldierID       int
	AddressID       int
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testsck")
	if err != nil {
		fmt.Println("connect fail")
	} else {
		fmt.Println("connect success")
	}

	defer db.Close()

	result, _ := db.Query("SELECT * FROM testsck.user")

	for result.Next() {
		var user Usersdata
		err := result.Scan(
			&user.ID,
			&user.CitizenID,
			&user.Fristname,
			&user.Lastname,
			&user.BirthYear,
			&user.FristnameFather,
			&user.LastnameFather,
			&user.FristnameMother,
			&user.LastnameMother,
			&user.SoldierID,
			&user.AddressID,
		)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user)
	}
}
