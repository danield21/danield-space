package service

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/gorilla/sessions"
)

type Scope struct {
	request *http.Request
	session *sessions.Session
}

func (s Scope) Request() *http.Request {
	return s.request
}

func (s Scope) Session() *sessions.Session {
	return s.session
}

func NewScope(req *http.Request) Scope {
	ses := envir.GetSession(req)
	return Scope{req, ses}
}
