package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const tamanhoMinChaveSecreta = 32

type JWTMaker struct {
	chaveSecreta string
}

func NewJWTMaker(chaveSecreta string) (Maker, error) {
	if len(chaveSecreta) < tamanhoMinChaveSecreta {
		return nil, fmt.Errorf("tamanho de chave invalida: deve ter pelo menos %d caracteres", tamanhoMinChaveSecreta)
	}
	return &JWTMaker{chaveSecreta}, nil
}

// CriarToken irá criar um novo token para um nome de usuario especifico
func (maker *JWTMaker) CriarToken(nome_usuario string, duracao time.Duration) (string, *Payload, error) {
	payload, err := NovoPayload(nome_usuario, duracao)
	if err != nil {
		return "", payload, err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.chaveSecreta))
	return token, payload, err
}

// VerificarToken irá checar se um token é valido ou não
func (maker *JWTMaker) VerificarToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErroTokenInvalido
		}
		return []byte(maker.chaveSecreta), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErroTokenExpirado) {
			return nil, ErroTokenExpirado
		}
		return nil, ErroTokenInvalido
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErroTokenInvalido
	}
	return payload, nil
}
