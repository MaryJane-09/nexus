package validate

import (
	"errors"
	"log"
	"net/mail"
	"strings"
	"unicode"

	"github.com/MaryJane-09/nexus/backend/internal/user"
)

func ValidateRegister(info user.User) error {
	if info.Name == "" {
		return errors.New("Name cannot be empty")
	}
	for _, char := range info.Name {
		if unicode.IsDigit(char) {
			return errors.New("Name cannot contain numbers")
		}
		if !unicode.IsLetter(char) {
			if !unicode.IsSpace(char) {
				return errors.New("Name can only contain letters")
			}
		}
	}

	if info.Email == "" {
		return errors.New("Email cannot be empty")
	}
	_, err := mail.ParseAddressList(info.Email)
	if err != nil{
		return errors.New("Invalid email input")
	}

	if info.Password == ""{
		return errors.New("Password cannot be empty")
	}
	if len(info.Password) < 8 {
		return errors.New("Password is too short")
	}

	return nil
}
