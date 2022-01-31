package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/techschool/simplebank/util"
)

var validarMoeda validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if moeda, ok := fieldLevel.Field().Interface().(string); ok {
		// checar se a moeda é suportada
		return util.MoedaSuportada(moeda)
	}
	return false
}
