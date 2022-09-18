package config

import (
	"database/sql"
	"log"
)

func DbConn() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Mysqlpassword"
	dbName := "capitals"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println("Couldnt connect to the DB due to error:", err.Error())
		return nil
	}
	log.Println("DB Connectoin successful")
	return db
}
