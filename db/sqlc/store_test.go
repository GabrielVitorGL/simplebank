package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	conta1 := criarContaAleatoria(t)
	conta2 := criarContaAleatoria(t)
	fmt.Println(">> antes:", conta1.Saldo, conta2.Saldo)

	// fazer n transações para teste
	n := 5
	quantia := int64(10) // valor que será transacionado

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				DeIDConta:   conta1.ID,
				ParaIDConta: conta2.ID,
				Quantia:     quantia,
			})

			errs <- err // mandando o valor de err para o canal errs, para que possamos checar esses erros fora do loop
			results <- result
		}()
	}

	// checar resultados
	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// o resultado contém diversos valores dentro dele, vamos checar cada um

		// tranferencia
		transferencia := result.Transferencia
		require.NotEmpty(t, transferencia)
		require.Equal(t, conta1.ID, transferencia.DeIDConta)
		require.Equal(t, conta2.ID, transferencia.ParaIDConta)
		require.Equal(t, quantia, transferencia.Quantia)
		require.NotZero(t, transferencia.ID)
		require.NotZero(t, transferencia.CriadaEm)

		// verificando que realmente a transferencia foi concluída pesquisando-a pelo seu ID
		_, err = store.ObterTransferencia(context.Background(), transferencia.ID)
		require.NoError(t, err)

		// mudanca
		de_mudanca := result.DeMudanca
		require.NotEmpty(t, de_mudanca)
		require.Equal(t, conta1.ID, de_mudanca.IDConta)
		require.Equal(t, -quantia, de_mudanca.Quantia) // tem que ser igual a -quantia pois o valor deverá ter ido embora
		require.NotZero(t, de_mudanca.ID)
		require.NotZero(t, de_mudanca.CriadaEm)
		_, err = store.ObterMudanca(context.Background(), de_mudanca.ID)
		require.NoError(t, err)

		para_mudanca := result.ParaMudanca
		require.NotEmpty(t, para_mudanca)
		require.Equal(t, conta2.ID, para_mudanca.IDConta)
		require.Equal(t, quantia, para_mudanca.Quantia)
		require.NotZero(t, para_mudanca.ID)
		require.NotZero(t, para_mudanca.CriadaEm)
		_, err = store.ObterMudanca(context.Background(), para_mudanca.ID)
		require.NoError(t, err)

		// contas
		de_conta := result.DeConta
		require.NotEmpty(t, de_conta)
		require.Equal(t, conta1.ID, de_conta.ID)

		para_conta := result.ParaConta
		require.NotEmpty(t, para_conta)
		require.Equal(t, conta2.ID, para_conta.ID)

		// checar o saldo das contas
		fmt.Println(">> tx:", de_conta.Saldo, para_conta.Saldo)
		diferenca1 := conta1.Saldo - de_conta.Saldo   // quantia que foi subtraída da conta 1
		diferenca2 := para_conta.Saldo - conta2.Saldo // quantia que foi somada da conta 2
		println(diferenca1, diferenca2)
		require.Equal(t, diferenca1, diferenca2)
		require.True(t, diferenca1 > 0)
		require.True(t, diferenca1%quantia == 0)

		k := int(diferenca1 / quantia)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	//checar o saldo final das duas contas
	Conta1Atualizada, err := testQueries.ObterConta(context.Background(), conta1.ID)
	require.NoError(t, err)

	Conta2Atualizada, err := testQueries.ObterConta(context.Background(), conta2.ID)
	require.NoError(t, err)

	fmt.Println(">> depois:", conta1.Saldo, conta2.Saldo)
	require.Equal(t, conta1.Saldo-int64(n)*quantia, Conta1Atualizada.Saldo)
	require.Equal(t, conta2.Saldo+int64(n)*quantia, Conta2Atualizada.Saldo)
}

func TestTransferTxDeadLock(t *testing.T) {
	store := NewStore(testDB)

	conta1 := criarContaAleatoria(t)
	conta2 := criarContaAleatoria(t)
	fmt.Println(">> antes:", conta1.Saldo, conta2.Saldo)

	// fazer n transações para teste
	n := 10
	quantia := int64(10) // valor que será transacionado
	errs := make(chan error)


	for i := 0; i < n; i++ {
		de_id_conta := conta1.ID
		para_id_conta := conta2.ID

		if i % 2 == 1 {
			de_id_conta = conta2.ID
			para_id_conta = conta1.ID
		}
		go func() {
			ctx := context.Background()
			_, err := store.TransferTx(ctx, TransferTxParams{
				DeIDConta:   de_id_conta,
				ParaIDConta: para_id_conta,
				Quantia:     quantia,
			})

			errs <- err // mandando o valor de err para o canal errs, para que possamos checar esses erros fora do loop
		}()
	}

	// checar resultados
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	//checar o saldo final das duas contas
	Conta1Atualizada, err := testQueries.ObterConta(context.Background(), conta1.ID)
	require.NoError(t, err)

	Conta2Atualizada, err := testQueries.ObterConta(context.Background(), conta2.ID)
	require.NoError(t, err)

	fmt.Println(">> depois:", conta1.Saldo, conta2.Saldo)
	require.Equal(t, conta1.Saldo, Conta1Atualizada.Saldo)
	require.Equal(t, conta2.Saldo, Conta2Atualizada.Saldo)
}
