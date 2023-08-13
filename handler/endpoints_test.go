package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/bytedance/sonic"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
)

func initHTTPCall(method, path string) (*httptest.ResponseRecorder, echo.Context) {
	e := echo.New()
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	return rec, c
}

func buildHTTPRequestBody(c echo.Context, method, path string, bodyRequest interface{}) echo.Context {
	reqBytes, _ := sonic.Marshal(bodyRequest)
	reqBody := string(reqBytes)
	request, _ := http.NewRequest(method, path, strings.NewReader(reqBody))
	request.Header.Add("Content-Type", "application/json")
	c.SetRequest(request)

	return c
}

func TestAddUser(t *testing.T) {
	method, path := http.MethodPost, "/users"

	ctrl := gomock.NewController(t)
	mockRepo := repository.NewMockRepositoryInterface(ctrl)

	emptyReq := "This field is required"
	hashedPassword := "hashed_password"

	validReq := generated.AddUserRequest{
		FullName:    "full name",
		PhoneNumber: "+62818426881",
		Password:    "Pa$$w0rd",
	}

	userInput := repository.User{
		FullName:    "full name",
		PhoneNumber: "+62818426881",
		Password:    hashedPassword,
	}

	type args struct {
		req generated.AddUserRequest
	}
	tests := []struct {
		name           string
		args           args
		mockFunc       func()
		wantStatusCode int
		wantRes        generated.AddUserResponse
	}{
		{
			name: "bad request - invalid params",
			args: args{
				req: generated.AddUserRequest{
					FullName:    "",
					Password:    "",
					PhoneNumber: "",
				},
			},
			wantStatusCode: http.StatusBadRequest,
			wantRes: generated.AddUserResponse{
				Success: false,
				Validation: &generated.ValidationResult{
					PhoneNumber: &emptyReq,
					FullName:    &emptyReq,
					Password:    &emptyReq,
				},
			},
		},
		{
			name: "success - valid params",
			args: args{
				req: validReq,
			},
			mockFunc: func() {
				mockRepo.EXPECT().GenerateHashedAndSaltedPassword(validReq.Password).
					Return(hashedPassword, nil).Times(1)
				mockRepo.EXPECT().InsertNewUser(gomock.Any(), userInput).
					Return(nil)
			},
			wantStatusCode: http.StatusOK,
			wantRes: generated.AddUserResponse{
				Success: true,
			},
		},
		{
			name: "failed - error generating password",
			args: args{
				req: validReq,
			},
			mockFunc: func() {
				mockRepo.EXPECT().GenerateHashedAndSaltedPassword(validReq.Password).
					Return(hashedPassword, errors.New("error password")).Times(1)
			},
			wantStatusCode: http.StatusInternalServerError,
			wantRes: generated.AddUserResponse{
				Success: false,
				Error: &generated.ErrorResponse{
					Message: "error password",
				},
			},
		},
		{
			name: "failed - error inserting new user",
			args: args{
				req: validReq,
			},
			mockFunc: func() {
				mockRepo.EXPECT().GenerateHashedAndSaltedPassword(validReq.Password).
					Return(hashedPassword, nil).Times(1)
				mockRepo.EXPECT().InsertNewUser(gomock.Any(), userInput).
					Return(errors.New("error db"))
			},
			wantStatusCode: http.StatusInternalServerError,
			wantRes: generated.AddUserResponse{
				Success: false,
				Error: &generated.ErrorResponse{
					Message: "error db",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{
				Repository: mockRepo,
				validate:   initValidator(),
			}
			if tt.mockFunc != nil {
				tt.mockFunc()
			}

			rec, c := initHTTPCall(method, path)
			c = buildHTTPRequestBody(c, method, path, tt.args.req)
			s.AddUser(c)

			var got generated.AddUserResponse
			sonic.Unmarshal(rec.Body.Bytes(), &got)

			if rec.Code != tt.wantStatusCode {
				t.Errorf("invalid status code, got: %v, want: %v", rec.Code, tt.wantStatusCode)
			}

			if !reflect.DeepEqual(got, tt.wantRes) {
				t.Errorf("invalid body response, got: %v, want: %v", got, tt.wantRes)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	method, path := http.MethodPatch, "/users"

	type args struct {
		req generated.UpdateUserRequest
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantRes        generated.UpdateUserResponse
	}{
		{
			name: "success",
			args: args{
				req: generated.UpdateUserRequest{},
			},
			wantStatusCode: http.StatusOK,
			wantRes: generated.UpdateUserResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{}

			rec, c := initHTTPCall(method, path)
			c = buildHTTPRequestBody(c, method, path, tt.args.req)
			s.UpdateUser(c)

			var got generated.UpdateUserResponse
			sonic.Unmarshal(rec.Body.Bytes(), &got)

			if rec.Code != tt.wantStatusCode {
				t.Errorf("invalid status code, got: %v, want: %v", rec.Code, tt.wantStatusCode)
			}

			if !reflect.DeepEqual(got, tt.wantRes) {
				t.Errorf("invalid body response, got: %v, want: %v", got, tt.wantRes)
			}
		})
	}
}

func TestLoginUser(t *testing.T) {
	method, path := http.MethodPost, "/users/login"

	var (
		token  = "token-dummy"
		userID = 1
	)

	type args struct {
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantRes        generated.UserLoginResponse
	}{
		{
			name:           "success",
			args:           args{},
			wantStatusCode: http.StatusOK,
			wantRes: generated.UserLoginResponse{
				AuthToken: &token,
				UserId:    &userID,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{}

			rec, c := initHTTPCall(method, path)
			s.LoginUser(c)

			var got generated.UserLoginResponse
			sonic.Unmarshal(rec.Body.Bytes(), &got)

			if rec.Code != tt.wantStatusCode {
				t.Errorf("invalid status code, got: %v, want: %v", rec.Code, tt.wantStatusCode)
			}

			if !reflect.DeepEqual(got, tt.wantRes) {
				t.Errorf("invalid body response, got: %v, want: %v", got, tt.wantRes)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	method, path := http.MethodPost, "/users/login"
	rec, c := initHTTPCall(method, path)

	type args struct {
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantRes        generated.GetUserResponse
	}{
		{
			name:           "success",
			args:           args{},
			wantStatusCode: http.StatusOK,
			wantRes: generated.GetUserResponse{
				Success: true,
				Data: &generated.User{
					FullName:    "Dummy",
					PhoneNumber: "+62123",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{}

			s.GetUser(c)

			var got generated.GetUserResponse
			sonic.Unmarshal(rec.Body.Bytes(), &got)

			if rec.Code != tt.wantStatusCode {
				t.Errorf("invalid status code, got: %v, want: %v", rec.Code, tt.wantStatusCode)
			}

			if !reflect.DeepEqual(got, tt.wantRes) {
				t.Errorf("invalid body response, got: %v, want: %v", got, tt.wantRes)
			}
		})
	}
}
