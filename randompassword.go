package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "newuser"
	dbPass := "password"
	dbName := "locker"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {

	rand.Seed(time.Now().UnixNano())
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]byte, 25)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}
	newpassword := string(b)

	fmt.Println("password", newpassword)
	fmt.Println(b)
	db := dbConn()
	sql := fmt.Sprintf("INSERT INTO entry(user_name) VALUES ('%s')", newpassword)
	fmt.Println(sql)
	ll, err := db.Exec(sql)

	if err != nil {
		fmt.Println(err.Error())

	}
	id, err := ll.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println(id)
	defer db.Close()
}
