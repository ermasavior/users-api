package handler

import (
	"io"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/bytedance/sonic"
	"github.com/labstack/echo/v4"
)

func (s *Server) AddUser(ctx echo.Context) error {
	var (
		resp    generated.AddUserResponse
		userReq generated.AddUserJSONRequestBody
	)

	reqBody, _ := io.ReadAll(ctx.Request().Body)
	_ = sonic.Unmarshal(reqBody, &userReq)

	userInput := repository.User{
		FullName:    userReq.FullName,
		Password:    userReq.Password,
		PhoneNumber: userReq.PhoneNumber,
	}

	err := s.validate.Struct(userInput)
	if err != nil {
		validateRes := translateError(err)
		resp.Validation = validateRes.ToHTTPResponse()
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	userInput.Password, err = s.Repository.GenerateHashedAndSaltedPassword(userInput.Password)
	if err != nil {
		resp.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	err = s.Repository.InsertNewUser(ctx.Request().Context(), userInput)
	if err != nil {
		resp.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp = generated.AddUserResponse{
		Success: true,
	}
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) LoginUser(ctx echo.Context) error {
	var (
		resp    generated.UserLoginResponse
		userReq generated.LoginUserJSONRequestBody
	)

	c := ctx.Request().Context()
	reqBody, _ := io.ReadAll(ctx.Request().Body)
	_ = sonic.Unmarshal(reqBody, &userReq)

	user, err := s.Repository.GetUserByPhoneNumber(c, userReq.PhoneNumber)
	if err != nil {
		resp.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, resp)
	}
	if user.ID == 0 {
		resp.Error = &generated.ErrorResponse{
			Message: repository.FailedUserLoginMismatchedPhoneNumberPassword,
		}
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	isMatched, err := s.Repository.ComparePasswords(user.Password, userReq.Password)
	if err != nil {
		resp.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	if !isMatched {
		resp.Error = &generated.ErrorResponse{
			Message: repository.FailedUserLoginMismatchedPhoneNumberPassword,
		}
		return ctx.JSON(http.StatusBadRequest, resp)
	}

	token, err := s.Repository.GenerateToken(user)
	if err != nil {
		resp.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	err = s.Repository.IncrementSuccessLoginCount(c, user.ID)
	if err != nil {
		resp.Error = &generated.ErrorResponse{
			Message: err.Error(),
		}
		return ctx.JSON(http.StatusInternalServerError, resp)
	}

	resp.Success = true
	resp.UserId = &user.ID
	resp.AuthToken = &token

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
