package articles

import (
	"errors"
)

//ErrNoArticle appears when the article requested does not exist in the database
var ErrNoArticle = errors.New("unable to find article with category/url pair")

//ErrNilArticle appears when item parameter is nil
var ErrNilArticle = errors.New("err was nil")
