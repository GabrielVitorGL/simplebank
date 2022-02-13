package token
import (
	"testing"
	"time"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/util"
)
func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	
	nome_usuario := util.RandomOwner()
	duracao := time.Minute

	criado_em := time.Now()
	expirado_em := criado_em.Add(duracao)

	token, err := maker.CriarToken(nome_usuario, duracao)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerificarToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	require.NotZero(t, payload.ID)
	require.Equal(t, nome_usuario, payload.NomeUsuario)
	require.WithinDuration(t, criado_em, payload.CriadoEm, time.Second)
	require.WithinDuration(t, expirado_em, payload.ExpiradoEm, time.Second)
}
func TestJWTTokenExpirado(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	token, err := maker.CriarToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := maker.VerificarToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErroTokenExpirado.Error())
	require.Nil(t, payload)
}

func TestJWTTokenAlgNoneInvalido(t *testing.T) {
	payload, err := NovoPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerificarToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErroTokenInvalido.Error())
	require.Nil(t, payload)
}