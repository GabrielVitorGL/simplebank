package token

import "time"

// maker será uma interface para gerenciar toneks
type Maker interface {
	// CriarToken irá criar um novo token para um nome de usuario especifico
	CriarToken(nome_usuario string, duracao time.Duration) (string, error)

	// VerificarToken irá checar se um token é valido ou não
	VerificarToken(token string) (*Payload, error)
}
