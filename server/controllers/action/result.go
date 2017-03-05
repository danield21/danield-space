package action

import "github.com/danield21/danield-space/server/handler/form"

type Result struct {
	Form     *form.Form
	Redirect URL
}
