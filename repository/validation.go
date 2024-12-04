package repository

import (
	"errors"
	"regexp"
)

func ValidatePhone(phone string) error {
	regex := `^\+?[0-9]{10,15}$`
	matched, _ := regexp.MatchString(regex, phone)
	if !matched {
		return errors.New("nomor telepon tidak valid")
	}
	return nil
}

func ValidateEmail(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	if !matched {
		return errors.New("email tidak valid")
	}
	return nil
}