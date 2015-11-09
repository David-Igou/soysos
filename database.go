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

	_, err = db.Exec(
		"CREATE TABLE IF NOT EXISTS users ( id integer PRIMARY KEY, username varchar(64) NOT NULL, password varchar(64), sessionToken varchar(64))")

	checkErr(err)

	return db
}

func (s DB) InsertUser(u *User) (string, int, error) {

	stmt, err := s.db.Prepare("INSERT INTO users(id, username, password, sessionToken) values(?,?,?,?)")
	checkErr(err)

	session := sessionId()
	res, err := stmt.Exec(nil, u.Username, u.Password, session)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	log.Println(id)
	return session, int(id), err
}

func (s DB) FindUser(u *User) (bool, error) {
	// query
	ident := u.ID
	//var username string
	row := s.db.QueryRow("SELECT username FROM users WHERE id=?", ident)
	//row, err := s.db.QueryRow("SELECT username FROM userinfo WHERE id=?", id)
	//checkErr(err)

	var id []byte
	var username []byte
	var password []byte
	var sessionToken []byte

	err := row.Scan(&id, &username, &password, &sessionToken)
	checkErr(err)

	t := User{
		0,
		string(username),
		string(password),
		string(sessionToken),
	}

	if t.SessionToken == u.SessionToken {
		return true, nil
	}
	// for row.Next() {
	// var sessionToken []byte
	// var username []byte
	// var password []byte
	// var created []byte
	// 	err = row.Scan(&sessionToken, &username, &password, &created)
	// 	checkErr(err)
	// 	log.Print(string(sessionToken))
	// 	log.Print(string(username))
	// 	log.Print(string(password))
	// 	log.Print(string(created))
	// }

	return false, err

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (s DB) FindToken(token string) (bool, error) {

	row := s.db.QueryRow("SELECT sessionToken FROM users WHERE sessionToken=?", token)

	var sessionToken []byte

	err := row.Scan(&sessionToken)
	//checkErr(err)
	//log.Print(string(sessionToken))

	if err != nil {
		return false, err
	}
	if token == string(sessionToken) {
		return true, nil
	}
	return false, nil
}
