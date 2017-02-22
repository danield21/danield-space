package link

import (
	"net/http"

	"google.golang.org/appengine/log"

	"github.com/danield21/danield-space/server/handler"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

const UserKey = "user"

func User(s *sessions.Session) (string, bool) {
	iUser, ok := s.Values[UserKey]

	if !ok {
		return "", false
	}

	user, ok := iUser.(string)
	if !ok {
		delete(s.Values, UserKey)
		return "", false
	}

	return user, true
}

func SetUser(s *sessions.Session, user string) {
	s.Values[UserKey] = user
}

func SaveSession(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		req := handler.Request(ctx)
		ses := handler.Session(ctx)
		err := ses.Save(req, w)
		if err != nil {
			log.Warningf(ctx, "link.SaveSession - Unable to save session\n%v", err)
		}
		return h(ctx, e, w)
	}
}
