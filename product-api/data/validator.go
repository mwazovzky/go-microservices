// https://github.com/go-playground/validator
package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

// Validator
type Validator struct {
	validate *validator.Validate
}

// NewValidation creates a new Validation type
func NewValidator() *Validator {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return &Validator{validate}
}

// Validate SKU
// sku format must be qwe-asdf-zxcvb
func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
}
