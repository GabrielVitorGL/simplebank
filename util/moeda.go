package util

// moedas suportadas
const (
	USD = "USD"
	GBP = "GBP"
	EUR = "EUR"
	LTC = "LTC"
	BRL = "BRL"
	RUB = "RUB"
	CAD = "CAD"
	CST = "CST"
	CHE = "CHE"
	CHW = "CHW"
	BTN = "BTN"
	CAT = "CAT"
	NZDT = "NZDT"
	HKT = "HKT"
	CNY = "CNY"
	HKD = "HKD"
	COP = "COP"
	COU = "COU"
	CRC = "CRC"
)

func MoedaSuportada (moeda string) bool { // retornará True se a moeda for suportada
	switch moeda {
	case USD, GBP, EUR, LTC, BRL, RUB, CAD, CHE, CHW, BTN, CAT, NZDT, CST, HKT, CNY, HKD, COP, COU, CRC:
		return true
	}
	return false
}