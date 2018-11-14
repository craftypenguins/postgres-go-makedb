package main

import (
	"database/sql"
	"os"
	"log"

	//Extends database/sql to support postgres
	_ "github.com/lib/pq"
)

func handleErr(err error){
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	connStr := os.Getenv("POSTGRES_URL")
	if len(connStr) == 0 {
		log.Fatal("POSTGRES_URL not set, exiting")
	}

	dbToMake := os.Getenv("DB_TO_MAKE")
	if len(dbToMake) == 0 {
		log.Fatal("DB_TO_MAKE not set, exiting")
	}

	dbOwner := os.Getenv("DB_OWNER")
	dbOwnerPwd := os.Getenv("DB_OWNER_PWD")

	if len(dbOwner) != 0 && len(dbOwnerPwd) == 0 {
		log.Fatal("DB_OWNER is set but not DB_OWNER_PWD, exiting")
	}

	if len(dbOwnerPwd) != 0 && len(dbOwner) == 0 {
		log.Fatal("DB_OWNER_PWD is set but not DB_OWNER, exiting")
	}

	db, err := sql.Open("postgres", connStr)
	handleErr(err)

	rows, err := db.Query("SELECT 1 FROM pg_database WHERE datname='" + dbToMake +"';")
	handleErr(err)

	//If there is a row returned then the database exists
	if rows.Next() {
		log.Println("Database " + dbToMake +" exists")
	} else {
		log.Println("Database " + dbToMake +" does not exist, creating")
		_, err = db.Query("CREATE DATABASE "+ dbToMake +";")
		handleErr(err)
	}

	rows, err = db.Query("SELECT 1 FROM pg_roles WHERE rolname='" + dbOwner +"';")
	handleErr(err)

	//If there is a row returned then the user exists
	if rows.Next() {
		log.Println("User " + dbOwner +" exists, setting password and ownership")
		_, err = db.Query("ALTER USER "+ dbOwner +" WITH ENCRYPTED PASSWORD '"+ dbOwnerPwd +"';")
		handleErr(err)
	} else {
		log.Println("User " + dbOwner +" does not exist, creating")
		_, err = db.Query("CREATE USER "+ dbOwner +" WITH ENCRYPTED PASSWORD '"+ dbOwnerPwd +"';")
		handleErr(err)
	}

	log.Println("Setting user " + dbOwner +" as owner of database "+dbToMake)
	_, err = db.Query("ALTER DATABASE "+ dbToMake +" OWNER TO "+ dbOwner +";")
	handleErr(err)

	db.Close()
}
