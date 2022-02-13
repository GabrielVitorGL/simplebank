package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErroTokenInvalido = errors.New("o token e invalido")
	ErroTokenExpirado = errors.New("o token foi expirado")
)

type Payload struct {
	ID          uuid.UUID `json:"id"`
	NomeUsuario string    `json:"nome_usuario"`
	CriadoEm    time.Time `json:"criado_em"`
	ExpiradoEm  time.Time `json:"expirado_em"`
}

// NovoPayload irá criar um payload de um token com um nome de usuario e uma duração especificada
func NovoPayload(nome_usuario string, duracao time.Duration) (*Payload, error) {
	IDToken, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:          IDToken,
		NomeUsuario: nome_usuario,
		CriadoEm:    time.Now(),
		ExpiradoEm:  time.Now().Add(duracao),
	}
	return payload, nil
}

// Checa se o payload de um token é valido ou nao
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiradoEm) { // se o token já estiver expirado
		return ErroTokenExpirado
	}
	return nil
}
