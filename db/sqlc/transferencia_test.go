package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func criarTransferenciaAleatoria(t *testing.T, conta1, conta2 Conta) Transferencia {
	arg := CriarTransferenciaParams{
		DeIDConta: conta1.ID,
		ParaIDConta:   conta2.ID,
		Quantia:        util.RandomMoney(),
	}

	transferencia, err := testQueries.CriarTransferencia(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transferencia)

	require.Equal(t, arg.DeIDConta, transferencia.DeIDConta)
	require.Equal(t, arg.ParaIDConta, transferencia.ParaIDConta)
	require.Equal(t, arg.Quantia, transferencia.Quantia)

	require.NotZero(t, transferencia.ID)
	require.NotZero(t, transferencia.CriadaEm)

	return transferencia
}

func TestCriarTransferencia(t *testing.T) {
	conta1 := criarContaAleatoria(t)
	conta2 := criarContaAleatoria(t)
	criarTransferenciaAleatoria(t, conta1, conta2)
}

func TestObterTransferencia(t *testing.T) {
	conta1 := criarContaAleatoria(t)
	conta2 := criarContaAleatoria(t)
	transferencia1 := criarTransferenciaAleatoria(t, conta1, conta2)

	transferencia2, err := testQueries.ObterTransferencia(context.Background(), transferencia1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transferencia2)

	require.Equal(t, transferencia1.ID, transferencia2.ID)
	require.Equal(t, transferencia1.DeIDConta, transferencia2.DeIDConta)
	require.Equal(t, transferencia1.ParaIDConta, transferencia2.ParaIDConta)
	require.Equal(t, transferencia1.Quantia, transferencia2.Quantia)
	require.WithinDuration(t, transferencia1.CriadaEm, transferencia2.CriadaEm, time.Second)
}

func TestListarTransferencia(t *testing.T) {
	conta1 := criarContaAleatoria(t)
	conta2 := criarContaAleatoria(t)

	for i := 0; i < 5; i++ {
		criarTransferenciaAleatoria(t, conta1, conta2)
		criarTransferenciaAleatoria(t, conta2, conta1)
	}

	arg := ListarTransferenciasParams{
		DeIDConta: conta1.ID,
		ParaIDConta:   conta1.ID,
		Limit:         5,
		Offset:        5,
	}

	transferencias, err := testQueries.ListarTransferencias(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transferencias, 5)

	for _, transferencia := range transferencias {
		require.NotEmpty(t, transferencia)
		require.True(t, transferencia.DeIDConta == conta1.ID || transferencia.ParaIDConta == conta1.ID)
	}
}