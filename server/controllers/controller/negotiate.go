package controller

import (
	"errors"
	"mime"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type MIMEType struct {
	MIME       string
	Parameters map[string]string
	value      float64
}

type mimeTypes []MIMEType

func (m mimeTypes) Len() int {
	return len(m)
}
func (m mimeTypes) Less(i, j int) bool {
	return m[i].value > m[j].value
}
func (m mimeTypes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func Negotiate(r *http.Request, types ...string) (MIMEType, error) {
	validTypes := make([]string, 0, len(types))

	for _, t := range types {
		mediaType, _, err := mime.ParseMediaType(t)
		if err != nil {
			return MIMEType{}, errors.New("\"" + t + "\" is not a proper MIME type")
		}
		validTypes = append(validTypes, mediaType)
	}

	if len(validTypes) == 0 {
		return MIMEType{}, errors.New("Valid types is empty")
	}

	accept := strings.TrimSpace(r.Header.Get("Accept"))

	if accept == "" {
		return MIMEType{}, errors.New("Accept header is empty")
	}

	acceptTypes := strings.Split(accept, ",")

	combos := make([]MIMEType, 0, len(acceptTypes))
	for _, acceptType := range acceptTypes {
		mediaType, params, err := mime.ParseMediaType(acceptType)

		if err != nil {
			return MIMEType{}, errors.New("Badly formed Accept header")
		}

		mime := MIMEType{
			MIME:       mediaType,
			Parameters: params,
		}

		if params["q"] == "" {
			mime.value = 1
		} else {
			var err error
			mime.value, err = strconv.ParseFloat(params["q"], 64)
			if err != nil {
				continue
			}
		}

		combos = append(combos, mime)
	}

	sort.Stable(mimeTypes(combos))

	for _, combo := range combos {
		if combo.MIME == "*/*" {
			combo.MIME = validTypes[0]
			return combo, nil
		}
		for _, match := range types {
			if combo.MIME == match {
				combo.MIME = match
				return combo, nil
			} else if strings.HasPrefix(combo.MIME, "*/") {
				parts := strings.Split(combo.MIME, "/")
				matches := strings.Split(match, "/")

				if len(parts) == 1 || parts[1] == matches[1] {
					combo.MIME = match
					return combo, nil
				}
			} else if strings.HasSuffix(combo.MIME, "/*") {
				parts := strings.Split(combo.MIME, "/")
				matches := strings.Split(match, "/")

				if parts[0] == matches[0] {
					combo.MIME = match
					return combo, nil
				}
			}
		}
	}

	return MIMEType{}, errors.New("Unable to find matching MIME type")
}
