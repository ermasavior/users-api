// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

//go:generate mockgen -package=repository -source=interfaces.go -destination=interfaces.mock.gen.go
type RepositoryInterface interface {
	GenerateHashedAndSaltedPassword(password string) (string, error)

	InsertNewUser(ctx context.Context, input User) error
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
}
