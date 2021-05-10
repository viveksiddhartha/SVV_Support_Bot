package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
}

type RDatastore interface {
}

//var db sql.DB

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "svcrm:Pass#word1@tcp(localhost:3306)/sv_crm")
	db.SetConnMaxLifetime(20)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(5)
	db.Stats()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err.Error())
	}

	return db

}

/*
func RDBConn() (dbs *sql.DB) {
	dbDriver := "redis"
	dbs, err := sql.Open(dbDriver, "sv_crm:sv_crm@tcp(127.0.0.1:3306)/SV_CRM")
	if err != nil {
		panic(err.Error())
	}
	return dbs

}
*/
