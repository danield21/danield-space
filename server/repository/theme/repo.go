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
func GetApp(c context.Context) (theme string) {
	theme = getWithDefault(c, appSuffix, DefaultAppTheme)
	return
}

func getWithDefault(c context.Context, section, defaultTheme string) (theme string) {
	item := themeToItem(section, defaultTheme)

	field := bucket.Default(c, item)

	theme = itemsToTheme(field)

	return
}

func set(c context.Context, section string, theme string) (err error) {
	if !ValidTheme(theme) {
		err = ErrInvalidTheme
		return
	}

	item := themeToItem(section, theme)
	err = bucket.Set(c, item)
	return
}

//ValidTheme is a helper function to determine if a entered theme can be valid
func ValidTheme(theme string) bool {
	return repository.ValidUrlPart(theme)
}

func themeToItem(section, theme string) bucket.Item {
	return bucket.Item{
		Field: bucketPrefix + section,
		Value: theme,
		Type:  "string",
	}
}

func itemsToTheme(item bucket.Item) (theme string) {
	theme = item.Value
	return
}
