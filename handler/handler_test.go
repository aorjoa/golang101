package handler_test

import (
	"net/http"
	"net/http/httptest"
	"ssm/mock"
	"ssm/model"
	"ssm/server"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCallPathParam_ShouldBeReturnThatParam(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/pinpinpin", nil)
	rec := httptest.NewRecorder()
	uls := mock.UserLoginLogDB{}

	server := server.Server{
		Echo:           e,
		UserLoginLogDB: uls,
	}
	server.ServerRoute()
	server.ServeHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello : pinpinpin", rec.Body.String())
}

func TestCallPathJson_ShouldReturnJson(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/json", nil)
	rec := httptest.NewRecorder()
	uls := mock.UserLoginLogDB{}

	server := server.Server{
		Echo:           e,
		UserLoginLogDB: uls,
	}
	server.ServerRoute()
	server.ServeHTTP(rec, req)

	expected := `{"Username":"aorjoa","Point":10.2}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
}

func TestCallPathLogin_ShouldBeAbleToLogin(t *testing.T) {
	data := `{"Username":"aorjoa","Password":"test"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	uls := mock.UserLoginLogDB{}

	ulm := model.UserLoginLog{Username: "aorjoa", Password: "test"}

	uls.On("Create", &ulm).Return(nil)

	server := server.Server{
		Echo:           e,
		UserLoginLogDB: uls,
	}
	server.ServerRoute()
	server.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, data, strings.TrimSpace(rec.Body.String()))
}
