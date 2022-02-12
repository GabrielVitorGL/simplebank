package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestSenha(t *testing.T) {
	n := RandomInt(6, 12)
	senha := RandomString(int(n))

	// encriptando a senha
	senhaHash1, err := SenhaHash(senha)
	require.NoError(t, err)
	require.NotEmpty(t, senhaHash1)

	// conferindo a senha
	err = ChecarSenha(senha, senhaHash1)
	require.NoError(t, err)

	// testando com uma senha errada
	n = RandomInt(6, 12)
	senhaErrada := RandomString(int(n))
	err = ChecarSenha(senhaErrada, senhaHash1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	senhaHash2, err := SenhaHash(senha)
	require.NoError(t, err)
	require.NotEmpty(t, senhaHash2)
	require.NotEqual(t, senhaHash1, senhaHash2)
}
