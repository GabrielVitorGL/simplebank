package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/techschool/simplebank/db/sqlc"
	"github.com/techschool/simplebank/util"
)

type criarUsuarioRequerimentos struct {
	NomeUsuario  string `json:"nome_usuario" binding:"required,alphanum"`
	Senha        string `json:"senha" binding:"required,min=6"`
	NomeCompleto string `json:"nome_completo" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
}

type criarUsuarioResponse struct {
	NomeUsuario  string    `json:"nome_usuario"`
	NomeCompleto string    `json:"nome_completo"`
	Email        string    `json:"email"`
	MudancaSenha time.Time `json:"mudanca_senha"`
	CriadaEm     time.Time `json:"criada_em"`
}

func (servidor *Servidor) criarUsuario(ctx *gin.Context) {
	var req criarUsuarioRequerimentos
	if err := ctx.ShouldBindJSON(&req); err != nil { // Se houver erros
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	senhaHash, err := util.SenhaHash(req.Senha)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Caso nao haja erros, iremos criar a conta:
	arg := db.CriarUsuarioParams{
		NomeUsuario:  req.NomeUsuario,
		SenhaHash:    senhaHash, // O saldo é 0 pois uma conta deve começar com 0 de saldo, quando for criada
		NomeCompleto: req.NomeCompleto,
		Email:        req.Email,
	}

	usuario, err := servidor.store.CriarUsuario(ctx, arg)
	if err != nil { // Caso dê erro
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := criarUsuarioResponse{
		NomeUsuario: usuario.NomeUsuario,
		NomeCompleto: usuario.NomeCompleto,
		Email: usuario.Email,
		MudancaSenha: usuario.MudancaSenha,
		CriadaEm: usuario.CriadaEm,
	}
	// Se não houver erros
	ctx.JSON(http.StatusOK, resp)
}
