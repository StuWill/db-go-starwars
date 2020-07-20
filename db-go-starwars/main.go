package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "anakin"
	password = "pa55word"
	dbname   = "holocron"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// sqlCreateTable := `CREATE TABLE cats (
	// 	age INT,
	// 	first_name TEXT,
	// 	last_name TEXT
	// );`

	// _, err = db.Exec(sqlCreateTable)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Successfully created table!")

	sqlInsert := `INSERT INTO cats (age, first_name, last_name)
	VALUES (9, 'Bijou', 'Metcalf');`

	_, err = db.Exec(sqlInsert)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully inserted into table!")
}
