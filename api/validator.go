package api

import (
	"github.com/Sciipia/simplebank/util"
	"github.com/go-playground/validator/v10"
)

var validCyrrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		//check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
