package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)

func criarUsuarioAleatorio(t *testing.T) Usuario {
	arg := CriarUsuarioParams{
		NomeUsuario:  util.RandomOwner(),
		SenhaHash: "senha",
		NomeCompleto:  util.RandomOwner() + " " + util.RandomOwner(),
		Email: util.EmailAleatorio(),
	}

	usuario, err := testQueries.CriarUsuario(context.Background(), arg)
	require.NoError(t, err)    // Certifica que não houve erros. Se houver irá falhar
	require.NotEmpty(t, usuario) // Certifica que "conta" não está vazio

	require.Equal(t, arg.NomeUsuario, usuario.NomeUsuario)
	require.Equal(t, arg.SenhaHash, usuario.SenhaHash)
	require.Equal(t, arg.NomeCompleto, usuario.NomeCompleto)
	require.Equal(t, arg.Email, usuario.Email)

	require.True(t, usuario.MudancaSenha.IsZero())
	require.NotZero(t, usuario.CriadaEm)

	return usuario
}
func TestCreateUsuario(t *testing.T) { // Sempre seguir essa sintaxe
	criarUsuarioAleatorio(t)
}

func TestObterUsuario(t *testing.T) {
	usuario1 := criarUsuarioAleatorio(t)
	usuario2, err := testQueries.ObterUsuario(context.Background(), usuario1.NomeUsuario) // o resultado será uma segunda conta ou um erro
	require.NoError(t, err)
	require.NotEmpty(t, usuario2)

	require.Equal(t, usuario1.NomeUsuario, usuario2.NomeUsuario)
	require.Equal(t, usuario1.SenhaHash, usuario2.SenhaHash)
	require.Equal(t, usuario1.NomeCompleto, usuario2.NomeCompleto)
	require.Equal(t, usuario1.Email, usuario2.Email)

	require.WithinDuration(t, usuario1.MudancaSenha, usuario2.MudancaSenha, time.Second)
	require.WithinDuration(t, usuario1.CriadaEm, usuario2.CriadaEm, time.Second) // Tem que ser criadas em até 1 segundo de diferença
}
