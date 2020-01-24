package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Service struct {
	db *sql.DB
}

func (sv *Service) SQL() *sql.DB {
	return sv.db
}

func (sv *Service) Setup() {
	db, err := sql.Open("mysql", "root:hoang810@/db_banhang")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("connected to db")
	}
	sv.db = db
	defer db.Close()

	var name string
	if err := db.QueryRow("SELECT name FROM users"); err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(name)
}
