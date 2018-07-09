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
	defer db.Close()
	if err != nil {
		fmt.Println("connect fail")
	} else {
		fmt.Println("connect success")
	}
	fmt.Println(addUser(db))
	fmt.Println(getUser(db))

}

func getUser(db *sql.DB) []Usersdata {
	result, _ := db.Query("SELECT * FROM testsck.user")
	var userDataList []Usersdata
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
		userDataList = append(userDataList, user)
	}
	return userDataList
}

func addUser(db *sql.DB) bool {
	statement, _ := db.Prepare(`INSERT INTO testsck.user (citizen_id, firstname, lastname,
			birthyear, firstname_father, lastname_father,
			firstname_mother, lastname_mother, soldier_id, address_id
		)VALUES (?,?,?,?,?,?,?,?,?,?)`)

	_, err := statement.Exec("5537478625189", "ผักกาด", "เล็ปขบ", 2538, "นนัย", "เล็ปขบ", "ออย", "เล็ปขบ", "449", "1")

	if err != nil {
		panic(err.Error())
		return false
	}
	return true
}
