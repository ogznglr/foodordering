package helpers

import (
	"net/mail"
)

func ValidEmail(email string) (string, error) {
	addr, err := mail.ParseAddress(email)
	if err != nil {
		return "", err
	}
	return addr.Address, nil
}
