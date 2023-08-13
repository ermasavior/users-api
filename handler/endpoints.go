package handler

import (
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/labstack/echo/v4"
)

func (s *Server) AddUser(ctx echo.Context) error {
	var resp = generated.AddUserResponse{
		Success: true,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) LoginUser(ctx echo.Context) error {
	token := "token-dummy"
	userID := 1

	var resp = generated.UserLoginResponse{
		AuthToken: &token,
		UserId:    &userID,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetUser(ctx echo.Context) error {
	var resp = generated.GetUserResponse{
		Success: true,
		Data: &generated.User{
			FullName:    "Dummy",
			PhoneNumber: "+62123",
		},
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) UpdateUser(ctx echo.Context) error {
	var resp = generated.UpdateUserResponse{
		Success: true,
		UserId:  nil,
	}
	return ctx.JSON(http.StatusOK, resp)
}
