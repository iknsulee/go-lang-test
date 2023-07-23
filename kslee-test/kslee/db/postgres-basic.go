package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "simple_db"
)

func TestPostgres() {
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	// insert
	// hardcoded
	insertStmt := `insert into students(name, roll_number) values('Jacob', 20)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	//// dynamic
	insertDynStmt := `insert into students(name, roll_number) values($1, $2)`
	_, e = db.Exec(insertDynStmt, "Jack", 21)
	CheckError(e)

	// update
	updateStmt := `update students set name=$1, roll_number=$2 where roll_number=$3`
	_, e = db.Exec(updateStmt, "Rachel", 24, 8)
	CheckError(e)

	// select
	rows, err := db.Query(`SELECT name, roll_number FROM students`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var name string
		var roll_number int

		err = rows.Scan(&name, &roll_number)
		CheckError(err)

		fmt.Println(name, roll_number)
	}

	CheckError(err)

	fmt.Println("Connected!")
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
