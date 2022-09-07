package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./nee.db")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("DROP TABLE IF EXISTS User;")
	_, _ = db.Exec("CREATE TABLE User(Name text,Age integer);")
	result, err := db.Exec("INSERT INTO User(`Name`, `Age`) VALUES (?,?), (?,?)", "Tom", 18, "Sam", 20)
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}

	row := db.QueryRow("SELECT Name, Age FROM User Limit 1")
	var name string
	var age int
	if err = row.Scan(&name, &age); err == nil {
		log.Println(name, age)
	} else {
		log.Println("err:", err)
	}

}
