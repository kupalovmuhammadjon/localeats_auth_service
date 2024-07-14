package validations

import (
	"errors"
	"net/mail"
	"regexp"
	"unicode"
)

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func ValidatePassword(password string) error {
	var (
		hasMinLen      = false
		hasUpper       = false
		hasLower       = false
		hasNumber      = false
		hasSpecialChar = false
	)

	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasMinLen {
		return errors.New("password must be at least 8 characters long")
	}
	if !hasUpper {
		return errors.New("password must have at least one uppercase letter")
	}
	if !hasLower {
		return errors.New("password must have at least one lowercase letter")
	}
	if !hasNumber {
		return errors.New("password must have at least one digit")
	}
	if !hasSpecialChar {
		return errors.New("password must have at least one special character")
	}

	return nil
}

func ValidatePhoneNumber(number string) error {
	regexPattern := `^\+998(9[0-9])\d{7}$`

	regex := regexp.MustCompile(regexPattern)

	if !regex.MatchString(number) {
		return errors.New("invalid Uzbek phone number")
	}

	return nil
}
