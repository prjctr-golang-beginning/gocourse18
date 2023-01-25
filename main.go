package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type User struct {
	ID int `json:"id"`
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   `root`,
		Passwd: `password`,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "my-app",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Select
	rows, err := db.Query(`SELECT id FROM users`)
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	var us []User

	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID)
		if err != nil {
			log.Fatal(err)
		}
		us = append(us, u)
	}
	rows.Close()
	fmt.Println(us)

	// Insert
	res, err := db.Exec(`INSERT INTO courses (price, name, description) VALUES (127, 'Some new', 'Details about some new')`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())
}
