package about

import (
	"html/template"

	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/bucket"
	"golang.org/x/net/context"
)

var DefaultAbout = []byte(
	"<p>My name is Daniel Juan Dominguez, and I am a developer.</p>",
)

const bucketKey = "about-page"

func Get(c context.Context) (template.HTML, error) {
	item, err := aboutToItem(DefaultAbout)
	if err != nil {
		return "", err
	}

	field := bucket.Default(c, item)

	return itemToAbout(field)
}

func Set(c context.Context, about []byte) error {
	item, err := aboutToItem(about)
	if err != nil {
		return err
	}
	err = bucket.Set(c, item)
	return err
}

func aboutToItem(about []byte) (*bucket.Item, error) {
	clean, err := repository.CleanHTML(about)
	return bucket.NewItem(bucketKey, string(clean), "[]byte"), err
}

func itemToAbout(item *bucket.Item) (template.HTML, error) {
	return repository.CleanHTML([]byte(item.Value))
}
