package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

// esse servidor será responsável por fazer os requests HTTP para nosso serviço bancário
type Servidor struct {
	store    *db.Store
	roteador *gin.Engine
}

// Criará um servidor HTTP e irá prepará-lo
func NovoServidor(store *db.Store) *Servidor {
	servidor := &Servidor{store: store}
	roteador := gin.Default()

	roteador.POST("/contas", servidor.criarConta)
	roteador.GET("/contas/:id", servidor.obterConta)
	roteador.GET("/contas", servidor.listarContas)

	servidor.roteador = roteador
	return servidor
}

// Começa a rodar o server HTTP em um endereço específico
func (servidor *Servidor) Start(endereco string) error {
	return servidor.roteador.Run(endereco)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
