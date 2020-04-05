package handler_test

import (
	"net/http"
	"net/http/httptest"
	"ssm/handler"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCallPathParam_ShouldBeReturnThatParam(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/:name")
	c.SetParamNames("name")
	c.SetParamValues("pinpinpin")

	handler.ServeHomePath(c)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello : pinpinpin", rec.Body.String())
}

func TestCallPathJson_ShouldReturnJson(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("/json")

	handler.ServeHomePathReturnJson(c)
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

	c := e.NewContext(req, rec)
	handler.ServeHomePathReceivePostBody(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, data, strings.TrimSpace(rec.Body.String()))
}
