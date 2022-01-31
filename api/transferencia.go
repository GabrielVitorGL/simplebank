package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/techschool/simplebank/db/sqlc"
)

type transferenciaRequerimentos struct {
	DeIDConta   int64  `json:"de_id_conta" binding:"required,min=1"`
	ParaIDConta int64  `json:"para_id_conta" binding:"required,min=1"`
	Quantia     int64  `json:"quantia" binding:"required,gt=0"` // (gt=0: maior que 0) iremos usar int64 para ser mais simples, então as transações devem ser valores inteiros
	Moeda       string `json:"moeda" binding:"required,moeda"`
	// Não colocamos o saldo pois ele será 0, quando a conta for criada
}

func (servidor *Servidor) criarTransferencia(ctx *gin.Context) {
	var req transferenciaRequerimentos
	if err := ctx.ShouldBindJSON(&req); err != nil { // Se houver erros
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !servidor.validarConta(ctx, req.DeIDConta, req.Moeda) { // if simplificado, apenas o if e a função esperaria que ela fosse True, nesse caso como colocamos o ! irá esperar que ela seja False para executar o return
		return
	}

	if !servidor.validarConta(ctx, req.ParaIDConta, req.Moeda) {
		return
	}

	// Caso nao haja erros, iremos fazer a transferencia:
	arg := db.TransferTxParams{
		DeIDConta:   req.DeIDConta,
		ParaIDConta: req.ParaIDConta,
		Quantia:     req.Quantia,
	}

	resultado, err := servidor.store.TransferTx(ctx, arg)
	if err != nil { // Caso dê erro
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Se não houver erros
	ctx.JSON(http.StatusOK, resultado)
}

func (servidor *Servidor) validarConta(ctx *gin.Context, IDConta int64, moeda string) bool { // Irá checar se uma conta realmente existe e se a moeda é a mesma da especificada no input
	conta, err := servidor.store.ObterConta(ctx, IDConta)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	// checando se a moeda da conta bate com a inserida
	if conta.Moeda != moeda {
		err := fmt.Errorf("a moeda da conta [%d](%s) nao corresponde a digitada (%s)", conta.ID, conta.Moeda, moeda)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return false
	}
	
	// se tudo der certo
	return true
}
