package main

import (
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	jwtSvc "github.com/SawitProRecruitment/UserService/pkg/jwt"
	"github.com/SawitProRecruitment/UserService/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")
	jwtService := jwtSvc.NewJWTService()

	var repo repository.RepositoryInterface = repository.NewRepository(repository.NewRepositoryOptions{
		Dsn:        dbDsn,
		JwtService: jwtService,
	})
	opts := handler.NewServerOptions{
		Repository: repo,
	}
	return handler.NewServer(opts)
}
