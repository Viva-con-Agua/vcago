package vcago

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Test struct {
	Server Server
}

func NewTest(server *Server) *Test {
	return &Test{
		Server: *server,
	}

}

func (i *Test) POSTContext(data string, rec *httptest.ResponseRecorder, token *jwt.Token) Context {
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cc := i.Server.NewContext(req, rec)
	if token != nil {
		cc.Set("token", token)
	}
	return Context{Model: "test", Context: cc}
}

func (i *Test) GETByIDContext(data string, rec *httptest.ResponseRecorder, token *jwt.Token) Context {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cc := i.Server.NewContext(req, rec)
	if token != nil {
		cc.Set("token", token)
	}
	cc.SetParamNames("id")
	cc.SetParamValues(data)
	return Context{Model: "test", Context: cc}
}

func (i *Test) PUTContext(data string, rec *httptest.ResponseRecorder, token *jwt.Token) Context {
	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cc := i.Server.NewContext(req, rec)
	if token != nil {
		cc.Set("token", token)
	}
	return Context{Model: "test", Context: cc}
}

func (i *Test) GETContext(data string, rec *httptest.ResponseRecorder, token *jwt.Token) Context {
	req := httptest.NewRequest(http.MethodGet, "/"+data, nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cc := i.Server.NewContext(req, rec)
	if token != nil {
		cc.Set("token", token)
	}
	return Context{Model: "test", Context: cc}
}

func (i *Test) DELETEContext(data string, rec *httptest.ResponseRecorder, token *jwt.Token) Context {
	req := httptest.NewRequest(http.MethodDelete, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	cc := i.Server.NewContext(req, rec)
	if token != nil {
		cc.Set("token", token)
	}
	cc.SetParamNames("id")
	cc.SetParamValues(data)
	return Context{Model: "test", Context: cc}
}
