package models

import (
	"database/sql"
	"github.com/bangweiz/blog/pkg"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	DB *sql.DB
)

type BaseModel struct {
	ID int `json:"id"`
	CreateOn string `json:"created_on"`
	ModifiedOn string `json:"modified_on"`
}

func init() {
	db, err := sql.Open("mysql", pkg.URI)
	if err != nil {
		log.Fatalf("cannot connect to database due to err: %v", err)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	DB = db
}
