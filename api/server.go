package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/db/sqlc"
)

// esse servidor será responsável por fazer os requests HTTP para nosso serviço bancário
type Servidor struct {
	store    db.Store
	roteador *gin.Engine
}

// Criará um servidor HTTP e irá prepará-lo
func NovoServidor(store db.Store) *Servidor {
	servidor := &Servidor{store: store}
	roteador := gin.Default()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok { // validando se a moeda é suportada
		v.RegisterValidation("moeda", validarMoeda)
	}

	roteador.POST("/usuarios", servidor.criarUsuario)
	
	roteador.POST("/contas", servidor.criarConta)
	roteador.GET("/contas/:id", servidor.obterConta)
	roteador.GET("/contas", servidor.listarContas)

	roteador.POST("/transferencias", servidor.criarTransferencia)

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
