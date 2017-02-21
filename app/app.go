package app

import (
	"net/http"

	"github.com/danield21/danield-space/server"
)

func init() {
	s := server.New()
	http.Handle("/", s)
}
