package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func IsAdultValidator(field validator.FieldLevel) bool {
	fmt.Println(field.FieldName())
	return true
}
