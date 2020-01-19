package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Service struct {
	DB *sql.DB
}

func (sv *Service) Setup() {
	db, err := sql.Open("mysql", "root:hoang810@/db_banhang")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("connected to db")
	}
	sv.DB = db
	defer db.Close()
}
