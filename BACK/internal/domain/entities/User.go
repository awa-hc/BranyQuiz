package entities

import (
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	if err := u.ValidateUserName(); err != nil {
		return err

	}

	if err := u.ValidateEmail(); err != nil {
		return err
	}

	if err := u.validatePassword(); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateUserName() error {
	usernameRegex := regexp.MustCompile("^[a-zA-Z0-9]{3,32}$")

	if !usernameRegex.MatchString(u.UserName) {
		return &ValidationError{Field: "username", Message: "username must be between 3 and 32 characters long"}
	}

	return nil
}

func (u *User) ValidateEmail() error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$")

	if !emailRegex.MatchString(u.Email) {
		return &ValidationError{Field: "Email", Message: "Invalid email format"}
	}

	return nil
}

func (u *User) validatePassword() error {
	minLength := 8
	hasLower := false
	hasUpper := false
	hasDigit := false
	hasSpecial := false

	for _, char := range u.Password {
		if char >= 'a' && char <= 'z' {
			hasLower = true
		} else if char >= 'A' && char <= 'Z' {
			hasUpper = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		} else {
			// Convert rune to byte for strings.IndexByte
			charByte := byte(char)
			if strings.IndexByte("~!@#$%^&*()_+-=[]{};':\"\\|,.<>/?", charByte) != -1 {
				hasSpecial = true
			}
		}
	}

	if len(u.Password) < minLength ||
		!hasLower || !hasUpper || !hasDigit || !hasSpecial {
		return &ValidationError{Field: "Password", Message: "Invalid password: must be at least 8 characters long and contain one lowercase, one uppercase, one number, and one special character"}
	}

	return nil
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}
