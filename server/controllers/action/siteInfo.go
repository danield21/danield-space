package action

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/models"
	"golang.org/x/net/context"
)

const siteTitleKey = "title"
const siteOwnerKey = "owner"
const siteDescriptionKey = "description"
const siteLinkKey = "link"

func UnpackSiteInfo(ctx context.Context, values url.Values) (*models.SiteInfo, form.Form) {
	frm := form.MakeForm()
	frm.Submitted = true

	titleFld := frm.AddFieldFromValue(siteTitleKey, values)
	form.NotEmpty(titleFld, "title is required")

	ownerFld := frm.AddFieldFromValue(siteOwnerKey, values)
	form.NotEmpty(ownerFld, "author is required")

	dscFld := frm.AddFieldFromValue(siteDescriptionKey, values)
	form.NotEmpty(dscFld, "author is required")

	linkFld := frm.AddFieldFromValue(siteLinkKey, values)
	form.NotEmpty(linkFld, "url is required")

	if frm.HasErrors() {
		return nil, frm
	}

	s := new(models.SiteInfo)
	*s = models.SiteInfo{
		Title:       titleFld.Get(),
		Owner:       ownerFld.Get(),
		Description: dscFld.Get(),
		Link:        linkFld.Get(),
	}

	return s, frm
}

func RepackSiteInfo(info models.SiteInfo) form.Form {
	frm := form.MakeForm()

	titleFld := new(form.Field)
	titleFld.Values = []string{info.Title}
	frm.Fields[siteTitleKey] = titleFld

	ownerFld := new(form.Field)
	ownerFld.Values = []string{info.Owner}
	frm.Fields[siteOwnerKey] = ownerFld

	dscFld := new(form.Field)
	dscFld.Values = []string{info.Description}
	frm.Fields[siteDescriptionKey] = dscFld

	linkFld := new(form.Field)
	linkFld.Values = []string{info.Link}
	frm.Fields[siteLinkKey] = linkFld

	return frm
}

func PutSiteInfoLink(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		r := handler.Request(ctx)
		err := r.ParseForm()
		if err != nil {
			return h(WithForm(ctx, form.Form{Error: errors.New("Unable to parse form")}), e, w)
		}

		info, frm := UnpackSiteInfo(ctx, r.Form)
		if info == nil {
			return h(WithForm(ctx, frm), e, w)
		}

		err = e.Repository().SiteInfo().Set(ctx, *info)
		if err != nil {
			frm.Error = errors.New("Unable to put into database")
			return h(WithForm(ctx, frm), e, w)
		}

		return h(WithForm(ctx, frm), e, w)
	}
}
