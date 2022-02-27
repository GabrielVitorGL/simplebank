package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NovoPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("tamanho da chave invalido: deve conter exatamente %d caracteres", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}
	return maker, nil
}

// CriarToken irá criar um novo token para um nome de usuario especifico
func (maker *PasetoMaker) CriarToken(nome_usuario string, duracao time.Duration) (string, *Payload, error) {
	payload, err := NovoPayload(nome_usuario, duracao)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// VerificarToken irá checar se um token é valido ou não
func (maker *PasetoMaker) VerificarToken(token string) (*Payload, error) {
	payload := &Payload{} // declarando um objeto do tipo payload vazio que irá armazenar os dados decriptados mais tarde

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErroTokenInvalido
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
