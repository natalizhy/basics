package model

import (
	"github.com/go-ozzo/ozzo-validation"
)

func required(cond bool) validation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return validation.Validate(value, validation.Required)
		}

		return nil
	}
}
