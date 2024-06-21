package common

// import (
// 	//"net/http"
// 	"regexp"
// )

// type PasswordValidationResponse struct {
// 	IsValid bool
// 	Message string
// }

// func ValidatePasswordString(password string) (bool, *PasswordValidationResponse) {
// 	regex := `^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,20}$`
// 	matched, _ := regexp.MatchString(regex, password)

// 	if !matched {
// 		return false, &PasswordValidationResponse{
// 			Message: "Password must contain a capital letter, number, special character, and be between 8 and 20 characters long.",
// 		}
// 	}
// 	return true, nil
// }

// package common

import (
	"errors"
	"regexp"
)

type PasswordValidationResponse struct {
	IsValid bool
}

func ValidatePasswordString(password string) (*PasswordValidationResponse, error) {
	if len(password) < 8 {
		return &PasswordValidationResponse{IsValid: false}, errors.New("password must be at least 8 characters long")
	}
	if len(password) > 20 {
		return &PasswordValidationResponse{IsValid: false}, errors.New("password must be no more than 20 characters long")
	}

	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)

	if !hasDigit {
		return &PasswordValidationResponse{IsValid: false}, errors.New("password must contain at least one digit")
	}
	if !hasLower {
		return &PasswordValidationResponse{IsValid: false}, errors.New("password must contain at least one lowercase letter")
	}
	if !hasUpper {
		return &PasswordValidationResponse{IsValid: false}, errors.New("password must contain at least one uppercase letter")
	}

	return &PasswordValidationResponse{IsValid: true}, nil
}
