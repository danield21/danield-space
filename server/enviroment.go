package server

import (
	"net/http"

	"github.com/danield21/danield-space/server/handler"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

//ProductionEnvironment contains all the data required to run the server
type ProductionEnvironment struct {
	Connections handler.Repositories
}

//Session gets the session using a secure key
func (p *ProductionEnvironment) Session(ctx context.Context, r *http.Request) (session *sessions.Session) {
	session = GetSession(ctx, p, r)
	return
}

//Context gets a new app engine context
func (p *ProductionEnvironment) Context(r *http.Request) (ctx context.Context) {
	ctx = appengine.NewContext(r)
	return
}

func (p *ProductionEnvironment) Repository() handler.Repositories {
	return p.Connections
}
