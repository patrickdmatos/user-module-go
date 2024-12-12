package services

import (
	"errors"

	"github.com/patrickdmatos/api-shared-library-go/auth"
)

func AutenticateUser(email, password string) (string, error) {
	token, err := auth.AuthenticateUser(email, password)
	if err != nil {
		// Add more detailed logging to debug the error
		return "", errors.New("authentication failed: " + err.Error())
	}

	return token, nil
}
