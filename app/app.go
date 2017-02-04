package server

import (
	"net/http"

	server "github.com/danield21/danield-space/pkg"
)

func init() {
	s := server.New()
	http.Handle("/", s)
}
