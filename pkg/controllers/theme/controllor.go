package siteInfo

import (
	"github.com/danield21/danield-space/pkg/controllers/bucket"
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//DefaultAppTheme is the default theme for the public section of the site
const DefaultAppTheme = "balloon"

//DefaultAdminTheme is the default theme for the admin section of the site
const DefaultAdminTheme = "balloon-admin"

const bucketPrefix = "theme-"
const appSuffix = "app"
const adminSuffix = "admin"

func GetApp(c context.Context) (theme string) {
	var ok bool
	theme, ok = get(c, appSuffix)
	if ok {
		return
	}

	theme = DefaultAppTheme
	err := set(c, appSuffix, theme)
	if err != nil {
		log.Warningf(c, "theme.GetApp - Unable to put theme for app into database\n%v", err)
	}
	return
}

func GetAdmin(c context.Context) (theme string) {
	var ok bool
	theme, ok = get(c, adminSuffix)
	if ok {
		return
	}

	theme = DefaultAdminTheme
	err := set(c, adminSuffix, theme)
	if err != nil {
		log.Warningf(c, "theme.GetAdmin - Unable to put theme for admin into database\n%v", err)
	}
	return
}

//Get gets all information about the site
func get(c context.Context, section string) (theme string, ok bool) {
	field, err := bucket.Get(c, bucketPrefix+section)

	if err != nil {
		log.Warningf(c, "theme.get - Unable to get default from database\n%v", err)
		return
	}

	theme = field.Value
	ok = true
	return
}

func set(c context.Context, section string, theme string) (err error) {
	item := bucket.Item{
		Field: bucketPrefix + section,
		Value: theme,
		Type:  "string",
	}
	err = bucket.Set(c, item)
	return
}
