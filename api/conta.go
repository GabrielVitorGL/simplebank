package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type criarContaRequerimentos struct {
	Dono  string `json:"dono" binding:"required"`
	Moeda string `json:"moeda" binding:"required,oneof=USD GBP EUR LTC BRL RUB CAD CST CHE CHW BTN CAT NZDT CST HKT CNY HKD COP COU CRC"`
	// Não colocamos o saldo pois ele será 0, quando a conta for criada
}

func (servidor *Servidor) criarConta(ctx *gin.Context) {
	var req criarContaRequerimentos
	if err := ctx.ShouldBindJSON(&req); err != nil { // Se houver erros
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Caso nao haja erros, iremos criar a conta:
	arg := db.CriarContaParams{
		Dono:  req.Dono,
		Saldo: 0, // O saldo é 0 pois uma conta deve começar com 0 de saldo, quando for criada
		Moeda: req.Moeda,
	}

	conta, err := servidor.store.CriarConta(ctx, arg)
	if err != nil { // Caso dê erro
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Se não houver erros
	ctx.JSON(http.StatusOK, conta)
}

type ObterContaRequerimentos struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (servidor *Servidor) obterConta(ctx *gin.Context) {
	var req ObterContaRequerimentos
	if err := ctx.ShouldBindUri(&req); err != nil { // Se houver erros
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	conta, err := servidor.store.ObterConta(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		return
	}

	ctx.JSON(http.StatusOK, conta)
}

type ListarContasRequerimentos struct {
	IDPag  int32 `form:"id_pag" binding:"required,min=1"`
	TamPag int32 `form:"tam_pag" binding:"required,min=5,max=10"`
}

func (servidor *Servidor) listarContas(ctx *gin.Context) {
	var req ListarContasRequerimentos
	if err := ctx.ShouldBindQuery(&req); err != nil { // Se houver erros
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListarContasParams{
		Limit:  req.TamPag,
		Offset: (req.IDPag - 1) * req.TamPag,
	}

	contas, err := servidor.store.ListarContas(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, contas)
}
