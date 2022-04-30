package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/go?parseTime=true")

	query := `
	CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`

	_, err = db.Exec(query)
	fmt.Printf("%+v\n", err)

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

	userId, err := result.LastInsertId()
	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", userId)

	var (
		id int
		//username  string
		//password  string
		//createdAt time.Time
	)

	query = `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err = db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	fmt.Printf("%+v\n", err)
	fmt.Printf("%d %s %s %s\n", id, username, password, createdAt)

	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`) // check err
	defer rows.Close()

	var users []user
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.username, &u.password, &u.createdAt) // check err
		users = append(users, u)
	}
	err = rows.Err() // check err
	fmt.Printf("%+v\n", err)
	fmt.Printf("%v\n", users)
}
