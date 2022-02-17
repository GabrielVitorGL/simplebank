// colocar o arquivo de teste na mesma pasta do resto do código
//! O nome do arquivo de teste deve terminar com o sufixo "text"

package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/techschool/simplebank/util"

	"github.com/stretchr/testify/require"
)

func criarContaAleatoria(t *testing.T) Conta {
	usuario := criarUsuarioAleatorio(t)

	arg := CriarContaParams{
		Dono:  usuario.NomeUsuario,
		Saldo: util.RandomMoney(),
		Moeda: util.RandomCurrency(),
	}

	conta, err := testQueries.CriarConta(context.Background(), arg)
	require.NoError(t, err)    // Certifica que não houve erros. Se houver irá falhar
	require.NotEmpty(t, conta) // Certifica que "conta" não está vazio

	require.Equal(t, arg.Dono, conta.Dono)
	require.Equal(t, arg.Saldo, conta.Saldo)
	require.Equal(t, arg.Moeda, conta.Moeda)

	require.NotZero(t, conta.ID) // Certifica que o ID da conta foi preenchido automaticamente pelo postgres
	require.NotZero(t, conta.CriadaEm)

	return conta
}
func TestCreateAccount(t *testing.T) { // Sempre seguir essa sintaxe
	criarContaAleatoria(t)
}

func TestObterConta(t *testing.T) {
	conta1 := criarContaAleatoria(t)
	conta2, err := testQueries.ObterConta(context.Background(), conta1.ID) // o resultado será uma segunda conta ou um erro
	require.NoError(t, err)
	require.NotEmpty(t, conta2)

	require.Equal(t, conta1.ID, conta2.ID)
	require.Equal(t, conta1.Dono, conta2.Dono)
	require.Equal(t, conta1.Saldo, conta2.Saldo)
	require.Equal(t, conta1.Moeda, conta2.Moeda)

	require.WithinDuration(t, conta1.CriadaEm, conta2.CriadaEm, time.Second) // Tem que ser criadas em até 1 segundo de diferença
}

func TestAtualizarConta(t *testing.T) {
	conta1 := criarContaAleatoria(t) // criando uma nova conta1

	arg := AtualizarContaParams{
		ID:    conta1.ID,
		Saldo: util.RandomMoney(), // mudando o saldo
	}

	conta2, err := testQueries.AtualizarConta(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, conta2)

	require.Equal(t, conta1.ID, conta2.ID)
	require.Equal(t, conta1.Dono, conta2.Dono)
	require.Equal(t, arg.Saldo, conta2.Saldo)
	require.Equal(t, conta1.Moeda, conta2.Moeda)

	require.WithinDuration(t, conta1.CriadaEm, conta2.CriadaEm, time.Second)
}

func TestDeletarConta(t *testing.T) {
	conta1 := criarContaAleatoria(t) // criando uma nova conta1
	err := testQueries.DeletarConta(context.Background(), conta1.ID)
	require.NoError(t, err)

	//Certificando que realmente foi deletada
	conta2, err := testQueries.ObterConta(context.Background(), conta1.ID)
	require.Error(t, err) // Precisa ser um erro pois a conta não pode ser encontrada para termos certeza que foi deletada
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, conta2) // Conta 2 precisa ser vazio
}

func TestListarContas(t *testing.T) {
	var ultimaConta Conta
	for i := 0; i < 10; i++ {
		ultimaConta = criarContaAleatoria(t) // Criará 10 contas aleatórias
	}

	arg := ListarContasParams{
		Dono: ultimaConta.Dono,
		Limit:  5, //Listará 5
		Offset: 0, //Comecará do 5
		// irá listar o 5, 6, 7, 8, 9
	}

	contas, err := testQueries.ListarContas(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, contas) // garantindo que o slice retornado teve 5 contas

	for _, conta := range contas { // ignorando o índice e salvando o valor em "conta"
		require.NotEmpty(t, conta) // garantindo que as contas não estão vazias
		require.Equal(t, ultimaConta.Dono, conta.Dono)
	}
}
