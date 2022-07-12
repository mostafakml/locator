package request

import (
	"github.com/go-playground/validator/v10"
	"strconv"
	"strings"
)

func ValidateIsConvertible(field validator.FieldLevel) bool {

	str_field:=field.Field().String()

	_,err:=strconv.ParseFloat(str_field,64)

	if err!=nil {
		return false
	}

	return true


	return strings.Contains(field.Field().String(),"cool")
}


