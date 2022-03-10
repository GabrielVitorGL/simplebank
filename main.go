package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/api"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

func main() {
	config, err := util.CarregarConfig(".")
	if err != nil {
		log.Fatal("nao foi possivel carregar as configuracoes:", err)
	}
	conec, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("nao foi possivel se conectar ao banco de dados:", err)
	}

	store := db.NewStore(conec)
	servidor, err := api.NovoServidor(config, store)
	if err != nil {
		log.Fatal("nao foi possivel criar o servidor")
	}

	err = servidor.Start(config.ServerAdress)
	if err != nil {
		log.Fatal("nao foi possivel iniciar o servidor", err)
	}
}
