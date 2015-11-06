package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

//This is the table scheme for creating your own user table.
// CREATE TABLE `userinfo` (
//     `sessionToken` VARCHAR(64) NULL,
//     `username` VARCHAR(64) NULL,
//     `password` VARCHAR(64) NULL,
//     `created` DATE NULL
// );

func Database() *sql.DB {
	db, err := sql.Open("sqlite3", "./users.db")
	checkErr(err)
	return db
}

func (s DB) InsertUser(u *User) {

	stmt, err := s.db.Prepare("INSERT INTO userinfo(username, password, sessionToken) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec(u.Username, u.Password, sessionId())
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	log.Println(id)
}

func (s DB) FindUser(u *User) {
	// query
	rows, err := s.db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var sessionToken []byte
		var username []byte
		var password []byte
		var created []byte
		err = rows.Scan(&sessionToken, &username, &password, &created)
		checkErr(err)
		log.Print(string(sessionToken))
		log.Print(string(username))
		log.Print(string(password))
		log.Print(string(created))
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
