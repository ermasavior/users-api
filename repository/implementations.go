package repository

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) InsertNewUser(ctx context.Context, input User) error {
	_, err := r.Db.ExecContext(ctx, queryInsertNewUser,
		input.FullName, input.PhoneNumber, input.Password)
	return err
}

func (r *Repository) GenerateHashedAndSaltedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error generating password", err)
		return "", err
	}
	return string(hash), err
}
