package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken         string    `json:"access_token"`
	AccessTokenExpiraEm time.Time `json:"access_token_expira_em"`
}

func (servidor *Servidor) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshPayload, err := servidor.tokenMaker.VerificarToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	secao, err := servidor.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if secao.IsBlocked {
		err := fmt.Errorf("secao bloqueada")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if secao.NomeUsuario != refreshPayload.NomeUsuario {
		err := fmt.Errorf("nome de usuario incorreto")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if secao.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("session token nao e valido")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	if time.Now().After(secao.ExpiraEm) {
		err := fmt.Errorf("sessao expirada")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := servidor.tokenMaker.CriarToken(
		refreshPayload.NomeUsuario,
		servidor.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:         accessToken,
		AccessTokenExpiraEm: accessPayload.ExpiradoEm,
	}
	ctx.JSON(http.StatusOK, rsp)
}
