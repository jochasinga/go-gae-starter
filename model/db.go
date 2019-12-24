package model

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

const connectionName = "go-test-app-255121:us-west2:testy-db"
var (
	db *sql.DB
	usrTable = `
CREATE TABLE usr (
  id       INT GENERATED ALWAYS AS IDENTITY,
  name     VARCHAR (40),
  email    VARCHAR (320) UNIQUE NOT NULL,
  phone    VARCHAR (20),
  created  TIMESTAMP NOT NULL,
  PRIMARY KEY (id)
);
`
	tables = [...]struct{
		name string
		stmt string
	}{
		{"usr", usrTable},
	}
)

func Initialize() {
	log.Println("initializing db")
	initDB()
	createTables()
	createSampleData()
}

func createSampleData() {
	createSampleUsers()
}

func tableAlreadyExists(name string) bool {
	schemaname := "public"
	var exists bool
	err := db.QueryRow(`
SELECT EXISTS (
  SELECT 1
  FROM     pg_tables
  WHERE    schemaname = $1
  AND      tablename = $2
);`, schemaname, name).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}
	return exists
}

func createTables() {
	for _, table := range tables {
		if tableAlreadyExists(table.name) {
			log.Printf("%s table already exists\n", table.name)
			continue
		}
		
		_, err := db.Exec(table.stmt)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func dropTables() {
	for _, table := range tables {
		if !tableAlreadyExists(table.name) {
			continue
		} else {
			stmt := fmt.Sprintf("DROP TABLE %s CASCADE", table.name)
			if _, err := db.Exec(stmt); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func setupDB() {
	_, err := db.Exec("CREATE DATABASE boo;")
	if err != nil {
		log.Fatal(err)
	}
}

func teardownDB() {
	_, err := db.Exec("DROP DATABASE boo;")
	if err != nil {
		log.Fatal(err)
	}
}


func initDB() {
	socketname := "/cloudsql/" + connectionName
	sslmode := "verify-full"
	dbURI := fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s sslmode=%s",
		"postgres",
		"postgres",
		socketname,
		"postgres",
		sslmode,
	)

	fmt.Printf("dbURI: %s\n", dbURI)

	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB successfully connected")
	db = conn
}
