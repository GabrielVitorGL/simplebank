package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// SenhaHash irá retornar a senha encriptada
func SenhaHash(senha string) (string, error) {
	senhaHash, err := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("falha em encriptar a senha: %w", err)
	}
	return string(senhaHash), nil
}

// ChecarSenha irá conferir se a senha digitada está correta ou não
func ChecarSenha(senha string, senhaHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senha))
}
