package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
)

const (
	dbDriver         = "postgres"
	dbSource         = "postgresql://root:pass412@localhost:5432/simple_bank?sslmode=disable"
	enderecoServidor = "0.0.0.0:8080"
)

func main() {
	conec, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("nao foi possivel se conectar ao banco de dados:", err)
	}

	store := db.NewStore(conec)
	servidor := api.NovoServidor(store)

	err = servidor.Start(enderecoServidor)
	if err != nil {
		log.Fatal("nao foi possivel iniciar o servidor", err)
	}
}
