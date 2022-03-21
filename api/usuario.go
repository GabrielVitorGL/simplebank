package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

type usuarioResponse struct {
	NomeUsuario  string    `json:"nome_usuario"`
	NomeCompleto string    `json:"nome_completo"`
	Email        string    `json:"email"`
	MudancaSenha time.Time `json:"mudanca_senha"`
	CriadaEm     time.Time `json:"criada_em"`
}

func newUsuarioResponse(usuario db.Usuario) usuarioResponse {
	return usuarioResponse{
		NomeUsuario:  usuario.NomeUsuario,
		NomeCompleto: usuario.NomeCompleto,
		Email:        usuario.Email,
		MudancaSenha: usuario.MudancaSenha,
		CriadaEm:     usuario.CriadaEm,
	}
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

	resp := newUsuarioResponse(usuario)
	// Se não houver erros
	ctx.JSON(http.StatusOK, resp)
}

type logarUsuarioRequest struct {
	NomeUsuario string `json:"nome_usuario" binding:"required,alphanum"`
	Senha       string `json:"senha" binding:"required,min=6"`
}

type logarUsuarioResponse struct {
	IDSecao              uuid.UUID       `json:"id_secao"`
	AccessToken          string          `json:"access_token"`
	AccessTokenExpiraEm  time.Time       `json:"access_token_expira_em"`
	RefreshToken         string          `json:"refresh_token"`
	RefreshTokenExpiraEm time.Time       `json:"refresh_token_expira_em"`
	Usuario              usuarioResponse `json:"usuario"`
}

func (servidor *Servidor) logarUsuario(ctx *gin.Context) {
	var req logarUsuarioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	usuario, err := servidor.store.ObterUsuario(ctx, req.NomeUsuario)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.ChecarSenha(req.Senha, usuario.SenhaHash)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, acessPayload, err := servidor.tokenMaker.CriarToken(
		usuario.NomeUsuario,
		servidor.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := servidor.tokenMaker.CriarToken(
		usuario.NomeUsuario,
		servidor.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	secao, err := servidor.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		NomeUsuario:  usuario.NomeUsuario,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiraEm:     refreshPayload.ExpiradoEm,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := logarUsuarioResponse{
		IDSecao: secao.ID,
		AccessToken: accessToken,
		AccessTokenExpiraEm: acessPayload.ExpiradoEm,
		RefreshToken: refreshToken,
		RefreshTokenExpiraEm: refreshPayload.ExpiradoEm,
		Usuario:     newUsuarioResponse(usuario),
	}
	ctx.JSON(http.StatusOK, rsp)
}
