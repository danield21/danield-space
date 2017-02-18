package theme

import (
	"errors"

	"github.com/danield21/danield-space/server/repository"
	"github.com/danield21/danield-space/server/repository/bucket"
	"golang.org/x/net/context"
)

//DefaultAppTheme is the default theme for the public section of the site
const DefaultAppTheme = "balloon"

//DefaultAdminTheme is the default theme for the admin section of the site
const DefaultAdminTheme = "balloon-admin"

//ErrInvalidTheme is the error returned when theme doesn't pass validation
var ErrInvalidTheme = errors.New("Theme can only have letters and \"-\"")

const bucketPrefix = "theme-"
const appSuffix = "app"

//GetApp returns the default theme for the public section of the site
//If unable to get theme from database, it will default to DefaultAppTheme
func GetApp(c context.Context) string {
	return getWithDefault(c, appSuffix, DefaultAppTheme)
}

func getWithDefault(c context.Context, section, defaultTheme string) string {
	item := themeToItem(section, defaultTheme)

	field := bucket.Default(c, item)

	return itemsToTheme(field)
}

func set(c context.Context, section string, theme string) error {
	if !ValidTheme(theme) {
		return ErrInvalidTheme
	}

	item := themeToItem(section, theme)
	err := bucket.Set(c, item)
	return err
}

//ValidTheme is a helper function to determine if a entered theme can be valid
func ValidTheme(theme string) bool {
	return repository.ValidUrlPart(theme)
}

func themeToItem(section, theme string) *bucket.Item {
	return bucket.NewItem(bucketPrefix+section, theme, "string")
}

func itemsToTheme(item *bucket.Item) string {
	return item.Value
}
