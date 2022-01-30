package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/techschool/simplebank/util"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.CarregarConfig("../..")
	if err != nil {
		log.Fatal("nao foi possivel carregar as configuracoes:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("nao foi possivel se conectar ao banco de dados:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
