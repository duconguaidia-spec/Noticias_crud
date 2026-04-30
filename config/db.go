package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "AgroCampo"
	schema := "Noticias"
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error al conectar:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("No se puede conectar:", err)
	}
	fmt.Println("Conexion exitosa a:", dbname, "esquema:", schema)
	DB = db
}