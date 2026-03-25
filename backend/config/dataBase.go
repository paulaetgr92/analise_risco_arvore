package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDatabase() *sql.DB {

	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "postgres"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Erro ao abrir conexão com banco:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	fmt.Println("Banco conectado com sucesso")
	fmt.Println("Banco conectado:", dbname)

	return db
}
