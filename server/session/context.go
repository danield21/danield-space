package session

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

type sessionKeyType string

const requestKey = sessionKeyType("request")
const sessionKey = sessionKeyType("sessions")

func NewContext(r *http.Request) context.Context {
	ctx := appengine.NewContext(r)
	ses := getSession(r)

	return context.WithValue(
		context.WithValue(ctx, requestKey, r),
		sessionKey, ses,
	)
}

func Request(ctx context.Context) *http.Request {
	r := ctx.Value(requestKey)
	req, ok := r.(*http.Request)
	if !ok {
		return nil
	}
	return req
}
