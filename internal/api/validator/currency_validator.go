package custom_validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/kimhieu153255/first-go/pkg/utils"
)

var ValidatorCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return utils.IsCurrency(currency)
	}
	return false
}
