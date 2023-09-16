package email_verify

import (
	"errors"
)

type Email struct {
	email string
}

func NewEmail(email string) (*Email, error) {
	if email == "" {
		return nil, errors.New("empty email address")
	}

	return &Email{
		email: email,
	}, nil
}
