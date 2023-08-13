package handler

import (
	"fmt"
	"regexp"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/go-playground/validator/v10"
)

func initValidator() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("customValidatePassword", ValidatePassword)
	return validate
}

type ValidationResult struct {
	IsValid     bool
	PhoneNumber string
	FullName    string
	Password    string
}

func (v *ValidationResult) ToHTTPResponse() *generated.ValidationResult {
	res := &generated.ValidationResult{}

	if v.FullName != "" {
		res.FullName = &v.FullName
	}
	if v.Password != "" {
		res.Password = &v.Password
	}
	if v.PhoneNumber != "" {
		res.PhoneNumber = &v.PhoneNumber
	}

	return res
}

func translateError(err error) ValidationResult {
	if err == nil {
		return ValidationResult{
			IsValid: true,
		}
	}

	res := ValidationResult{
		IsValid: false,
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		errMsg := generateMsgByTagAndParam(e.Tag(), e.Param())

		if e.Field() == "FullName" {
			res.FullName = errMsg
		}

		if e.Field() == "PhoneNumber" {
			res.PhoneNumber = errMsg
		}

		if e.Field() == "Password" {
			res.Password = errMsg
		}

	}

	return res
}

func generateMsgByTagAndParam(tag, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "max":
		return fmt.Sprintf("Should be less than %v characters", param)
	case "min":
		return fmt.Sprintf("Should be more than %v characters", param)
	case "startswith":
		return fmt.Sprintf("Should starts with %v", param)
	case "customValidatePassword":
		return "Must contain at least 1 capital characters and 1 number and 1 special (non alpha-numeric) characters"
	}
	return ""
}

func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	capitalRegex := regexp.MustCompile("[A-Z]")
	numberRegex := regexp.MustCompile("\\d")
	specialRegex := regexp.MustCompile("[^A-Za-z0-9]")

	hasCapital := capitalRegex.MatchString(password)
	hasNumber := numberRegex.MatchString(password)
	hasSpecial := specialRegex.MatchString(password)

	return hasCapital && hasNumber && hasSpecial
}
