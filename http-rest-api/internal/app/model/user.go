package model

import (
	"github.com/x/crypto/bcrypt"

	"github.com/go-ozzo/ozzo-validation/v4"
)

type User struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (u *User) Validate() error {
	return  validation.ValidateStruct(u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return  nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword{[]byte, bcrypt.MinCost}
	if err != nil {
		return "", err
	}
	return  string(b), nil
}
