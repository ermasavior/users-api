package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) InsertNewUser(ctx context.Context, input User) error {
	_, err := r.Db.ExecContext(ctx, queryInsertNewUser,
		input.FullName, input.PhoneNumber, input.Password)
	return err
}

func (r *Repository) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (user User, err error) {
	err = r.Db.QueryRowContext(ctx, queryGetUserByPhoneNumber, phoneNumber).
		Scan(&user.ID, &user.Password)
	if err == sql.ErrNoRows {
		return User{}, nil
	}
	if err != nil {
		return
	}
	return
}

func (r *Repository) IncrementSuccessLoginCount(ctx context.Context, id int) error {
	_, err := r.Db.ExecContext(ctx, queryIncrementSuccessLoginCount, id)
	return err
}

func (r *Repository) GenerateHashedAndSaltedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error GenerateHashedAndSaltedPassword", err)
		return "", err
	}
	return string(hash), err
}

func (r *Repository) ComparePasswords(hashedPwd, plainPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	if err != nil {
		log.Println("error ComparePasswords", err)
		return false, err
	}

	return true, nil
}

func (r *Repository) GenerateToken(user User) (string, error) {
	c := authClaims{
		Phone: user.PhoneNumber,
	}

	token, err := r.JwtService.GenerateToken(c)
	if err != nil {
		return "", err
	}

	return token, nil
}
