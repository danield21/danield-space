package service

import (
	"net/http"

	"github.com/danield21/danield-space/server/envir"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

type uniqueKey string

const requestKey = uniqueKey("request")
const sessionKey = uniqueKey("session")

func SetupContext(ctx context.Context, req *http.Request) context.Context {
	ses := envir.GetSession(req)

	rCtx := context.WithValue(ctx, requestKey, req)
	sCtx := context.WithValue(rCtx, sessionKey, ses)

	return sCtx
}

func Request(ctx context.Context) *http.Request {
	iReq := ctx.Value(requestKey)
	if req, ok := iReq.(*http.Request); !ok {
		return nil
	} else {
		return req
	}
}

func Session(ctx context.Context) *sessions.Session {
	iSes := ctx.Value(sessionKey)
	if ses, ok := iSes.(*sessions.Session); !ok {
		return nil
	} else {
		return ses
	}
}
