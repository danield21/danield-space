package articles

import (
	"errors"
)

var ErrNoArticle = errors.New("Unable to find article with type/key pair")
