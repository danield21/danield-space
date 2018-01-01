package store

import (
	"context"
	"regexp"
	"strings"
)

//Default has default information about the site
var DefaultSiteInfo = SiteInfo{
	Title:       "Ballooneer's Code",
	Link:        "http://danield.space",
	Owner:       "Daniel J Dominguez",
	Description: "Sometimes, having a lofty head is necessary. This is a site is dedicated to having an overview discussion to code, without worrying too much about implementation. Of course, we will be touching the ground to get a better view of what may need to happen, but for the most part we will care about the overall look and feel.",
}

type SiteInfoRepository interface {
	Get(ctx context.Context) SiteInfo
	Set(ctx context.Context, info SiteInfo) error
}

//SiteInfo is a struct containing information about the website
type SiteInfo struct {
	Title       string
	Link        string
	Owner       string
	Description string
}

//ShortDescription generates a shortened description
func (s SiteInfo) ShortDescription() string {
	const MaxLength = 200
	sentenceRegex := regexp.MustCompile("[.!?]+")
	sentences := sentenceRegex.Split(s.Description, -1)

	if len(sentences[0]) < MaxLength {
		return generateShortDescriptionUsing(sentences, ". ", MaxLength)
	}

	wordsRegex := regexp.MustCompile("\\s+")
	words := wordsRegex.Split(s.Description, -1)

	if len(sentences[0]) < MaxLength-3 {
		return generateShortDescriptionUsing(words, " ", MaxLength-3) + "..."
	}

	return strings.Join(strings.Split(s.Description, "")[:MaxLength-3], "") + "..."
}

func generateShortDescriptionUsing(breaks []string, seperator string, max int) string {
	short := ""
	for _, b := range breaks {
		tempShort := short + b + seperator
		if len(tempShort) > max {
			break
		}
		short = tempShort
	}
	return strings.TrimSpace(short)
}
