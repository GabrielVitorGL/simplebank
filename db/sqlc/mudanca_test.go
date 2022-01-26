package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func criarMudancaAleatoria(t *testing.T, conta Conta) Mudanca {
	arg := CriarMudancaParams{
		IDConta: conta.ID,
		Quantia: util.RandomMoney(),
	}

	mudanca, err := testQueries.CriarMudanca(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, mudanca)

	require.Equal(t, arg.IDConta, mudanca.IDConta)
	require.Equal(t, arg.Quantia, mudanca.Quantia)

	require.NotZero(t, mudanca.ID)
	require.NotZero(t, mudanca.CriadaEm)

	return mudanca
}

func TestCriarMudanca(t *testing.T) {
	conta := criarContaAleatoria(t)
	criarMudancaAleatoria(t, conta)
}

func TestObterMudanca(t *testing.T) {
	conta := criarContaAleatoria(t)
	mudanca1 := criarMudancaAleatoria(t, conta)

	mudanca2, err := testQueries.ObterMudanca(context.Background(), mudanca1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, mudanca2)

	require.Equal(t, mudanca1.ID, mudanca2.ID)
	require.Equal(t, mudanca1.IDConta, mudanca2.IDConta)
	require.Equal(t, mudanca1.Quantia, mudanca2.Quantia)
	require.WithinDuration(t, mudanca1.CriadaEm, mudanca2.CriadaEm, time.Second)
}

func TestListarMudancas(t *testing.T) {
	conta := criarContaAleatoria(t)

	for i := 0; i < 10; i++ {
		criarMudancaAleatoria(t, conta)
	}

	arg := ListarMudancasParams{
		IDConta: conta.ID,
		Limit:   5, //Listará 5
		Offset:  5, //Comecará do 5
		// irá listar o 5, 6, 7, 8, 9
	}

	mudancas, err := testQueries.ListarMudancas(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, mudancas, 5) // garantindo que o slice retornado teve 5 mudancas

	for _, mudanca := range mudancas { // ignorando o índice e salvando o valor em "mudanca"
		require.NotEmpty(t, mudanca) // garantindo que as contas não estão vazias
		require.Equal(t, arg.IDConta, mudanca.IDConta)
	}
}
