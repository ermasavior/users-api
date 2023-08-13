// This file contains types that are used in the repository layer.
package repository

type User struct {
	FullName    string `validate:"required,max=60,min=3"`
	Password    string `validate:"required,max=64,min=6,customValidatePassword"`
	PhoneNumber string `validate:"required,max=13,min=10,startswith=+62"`
}

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}
