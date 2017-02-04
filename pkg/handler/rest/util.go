package rest

import (
	"net/http"
	"strconv"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

const unlimited = -1

func GetRedirect(r *http.Request) (redirect string) {
	if redirects, ok := r.Form["redirect"]; ok {
		redirect = redirects[0]
	}
	return
}

func GetLimit(r *http.Request) (limit int) {
	context := appengine.NewContext(r)
	limit = unlimited

	if limits, ok := r.Form["limit"]; ok {
		iLimit, err := strconv.Atoi(limits[0])
		if err != nil {
			log.Warningf(context, "article.getLimit - Received limit cannot be parsed into int\n%v", err)
			return
		}
		limit = iLimit
	}
	return
}
