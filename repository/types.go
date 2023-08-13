// This file contains types that are used in the repository layer.
package repository

import "github.com/golang-jwt/jwt"

type User struct {
	ID          int    `db:"id"`
	FullName    string `validate:"required,max=60,min=3"`
	Password    string `validate:"required,max=64,min=6,customValidatePassword"`
	PhoneNumber string `validate:"required,max=13,min=10,startswith=+62"`
}

type UserLoginRequest struct {
	PhoneNumber string `validate:"required"`
	Password    string `validate:"required"`
}

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type authClaims struct {
	Phone string `json:"phone_number"`
	jwt.StandardClaims
}
