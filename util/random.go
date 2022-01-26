package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() { // Executará primeiro e pegará um número aleatório todas vezes que o código for executado
	rand.Seed(time.Now().UnixNano())
}

//Gera um número aleatório entre o mínimo e o máximo
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // gerará um valor maior que 0, e entre o máximo e o mínimo
}

//Gera um texto aleatório de tamanho n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)] // pegará um valor de "alphabet" entre 0 e k-1
		sb.WriteByte(c)
	}

	return sb.String()
}

// Gera um nome aleatório de um suposto dono para uma conta
func RandomOwner() string {
	n := RandomInt(0, 5)
	switch n {
	case 0:
		return RandomString(3)
	case 1:
		return RandomString(4)
	case 2:
		return RandomString(5)
	case 3:
		return RandomString(6)
	case 4:
		return RandomString(7)
	default:
		return RandomString(8)
	}
}

// Gera uma quantia aleatória de dinheiro para essa pessoa
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// Gera um código de moeda aleatório
func RandomCurrency() string {
	currencies := []string{"USD", "GBP", "EUR", "LTC", "BRL", "RUB", "CAD", "CST", "CHE", "CHW", "BTN", "CAT", "NZDT", "CST", "HKT", "CNY", "HKD", "COP", "COU", "CRC"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
