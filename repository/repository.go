// This file contains the repository implementation layer.
package repository

import (
	"database/sql"

	jwtSvc "github.com/SawitProRecruitment/UserService/pkg/jwt"
	_ "github.com/lib/pq"
)

type Repository struct {
	Db         *sql.DB
	JwtService jwtSvc.JWTServiceInterface
}

type NewRepositoryOptions struct {
	Dsn        string
	JwtService jwtSvc.JWTServiceInterface
}

func NewRepository(opts NewRepositoryOptions) *Repository {
	db, err := sql.Open("postgres", opts.Dsn)
	if err != nil {
		panic(err)
	}
	return &Repository{
		Db:         db,
		JwtService: opts.JwtService,
	}
}
