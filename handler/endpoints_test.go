package handler

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/bytedance/sonic"
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
	rec, c := initHTTPCall(method, path)

	type args struct {
		ctx echo.Context
		req generated.AddUserRequest
	}
	tests := []struct {
		name           string
		args           args
		wantStatusCode int
		wantRes        generated.AddUserResponse
	}{
		{
			name: "success",
			args: args{
				req: generated.AddUserRequest{
					FullName:    "full name",
					Password:    "password",
					PhoneNumber: "+621343",
				},
				ctx: func() echo.Context {
					return c
				}(),
			},
			wantStatusCode: http.StatusOK,
			wantRes: generated.AddUserResponse{
				Success: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Server{}

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
	rec, c := initHTTPCall(method, path)

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
	rec, c := initHTTPCall(method, path)

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
