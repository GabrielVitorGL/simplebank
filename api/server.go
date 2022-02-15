package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/token"
	"github.com/techschool/simplebank/util"
)

// esse servidor será responsável por fazer os requests HTTP para nosso serviço bancário
type Servidor struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	roteador   *gin.Engine
}

// Criará um servidor HTTP e irá prepará-lo
func NovoServidor(config util.Config, store db.Store) (*Servidor, error) {
	tokenMaker, err := token.NovoPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("nao foi possivel criar o token: %w", err)
	}
	servidor := &Servidor{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok { // validando se a moeda é suportada
		v.RegisterValidation("moeda", validarMoeda)
	}

	servidor.configurarRoteador()
	return servidor, nil
}

func (servidor *Servidor) configurarRoteador() {
	roteador := gin.Default()

	roteador.POST("/usuarios", servidor.criarUsuario)
	roteador.POST("/usuarios/login", servidor.logarUsuario)

	roteador.POST("/contas", servidor.criarConta)
	roteador.GET("/contas/:id", servidor.obterConta)
	roteador.GET("/contas", servidor.listarContas)

	roteador.POST("/transferencias", servidor.criarTransferencia)

	servidor.roteador = roteador

}

// Começa a rodar o server HTTP em um endereço específico
func (servidor *Servidor) Start(endereco string) error {
	return servidor.roteador.Run(endereco)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
