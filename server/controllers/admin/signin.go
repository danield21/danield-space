package admin

import (
	"errors"
	"html/template"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/controllers/link"
	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

const authUsrKey = "username"
const authPwdKey = "password"

var (
	ErrUnableAuth = errors.New("Unable to authenicate")
)

type SignInHandler struct {
	Context          handler.ContextGenerator
	Session          handler.SessionGenerator
	Renderer         handler.Renderer
	SiteInfo         store.SiteInfoRepository
	Account          store.AccountRepository
	MethodNotAllowed http.Handler
}

func (hnd SignInHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodGet {
		hnd.MethodNotAllowed.ServeHTTP(w, r)
	}

	ctx := hnd.Context.Generate(r)
	ses := hnd.Session.Generate(ctx, r)
	pg := handler.NewPage()

	info := hnd.SiteInfo.Get(ctx)

	pg.Title = info.Title
	pg.Meta["description"] = info.ShortDescription()
	pg.Meta["author"] = info.Owner

	var frm form.Form

	if r.Method == http.MethodPost {
		frm = hnd.ProcessForm(ctx, r, ses)
	}

	if frm.IsSuccessful() {
		pg.Status = 303
		pg.Header["Location"] = "."
	}

	cnt, err := hnd.Renderer.Render(ctx, "page/admin/signin", struct {
		Form form.Form
	}{
		frm,
	})

	if err != nil {
		log.Errorf(ctx, "admin.SignInHandler - Unable to render content\n%v", err)
		return
	}

	pg.Content = template.HTML(cnt)

	ses.Save(r, w)
	hnd.Renderer.Send(w, r, pg)
}

func (hnd SignInHandler) ProcessForm(ctx context.Context, r *http.Request, s *sessions.Session) form.Form {
	err := r.ParseForm()

	if err != nil {
		return form.Form{Error: errors.New("Unable to parse form")}
	}

	username, password, frm := unpackAuth(r.Form)
	if frm.HasErrors() {
		return frm
	}

	if !hnd.Account.CanLogIn(ctx, username, password) {
		log.Infof(ctx, "%s attempted to login with incorrect password", username)
		frm.Error = ErrUnableAuth
		return frm
	}

	link.SetUser(s, username)

	return frm
}

func unpackAuth(values url.Values) (string, []byte, form.Form) {
	frm := form.MakeForm()
	frm.Submitted = true

	username := frm.AddFieldFromValue(authUsrKey, values)
	form.NotEmpty(username, "Username is required")

	password := frm.AddFieldFromValue(authPwdKey, values)
	form.NotEmpty(password, "Password is required")

	return username.Get(), []byte(password.Get()), frm
}
