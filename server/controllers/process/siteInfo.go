package process

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/danield21/danield-space/server/form"
	"github.com/danield21/danield-space/server/store"
	"github.com/gorilla/sessions"
	"golang.org/x/net/context"
)

const siteTitleKey = "title"
const siteOwnerKey = "owner"
const siteDescriptionKey = "description"
const siteLinkKey = "link"

func UnpackSiteInfo(ctx context.Context, values url.Values) (*store.SiteInfo, form.Form) {
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

	s := new(store.SiteInfo)
	*s = store.SiteInfo{
		Title:       titleFld.Get(),
		Owner:       ownerFld.Get(),
		Description: dscFld.Get(),
		Link:        linkFld.Get(),
	}

	return s, frm
}

func RepackSiteInfo(info store.SiteInfo) form.Form {
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

type PutSiteInfoProcessor struct {
	SiteInfo store.SiteInfoRepository
}

func (prc PutSiteInfoProcessor) Process(ctx context.Context, rqs *http.Request, ses *sessions.Session) form.Form {
	err := rqs.ParseForm()
	if err != nil {
		return form.NewErrorForm(errors.New("Unable to parse form"))
	}

	info, frm := UnpackSiteInfo(ctx, rqs.Form)
	if info == nil {
		return frm
	}

	err = prc.SiteInfo.Set(ctx, *info)
	if err != nil {
		frm.Error = errors.New("Unable to put into database")
		return frm
	}

	return frm
}
