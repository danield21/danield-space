package controller_test

import (
	"net/http"
	"testing"

	"github.com/danield21/danield-space/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBasicHTMLNegotiate(t *testing.T) {
	rqs := new(http.Request)

	rqs.Header = make(map[string][]string)

	rqs.Header.Set("Accept", "text/html")
	mime, err := server.Negotiate(rqs, "text/html")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "text/html", mime.MIME, "Should pick HTML")
}

func TestBadNegotiate(t *testing.T) {
	rqs := new(http.Request)

	rqs.Header = make(map[string][]string)

	rqs.Header.Set("Accept", "application")
	_, err := server.Negotiate(rqs)
	assert.Error(t, err, "Should not allow empty valid types")

	rqs.Header.Set("Accept", "application/json")
	_, err = server.Negotiate(rqs, "application/$%^%$^")
	assert.Error(t, err, "Should not allow bad MIME as a valid type ")

	rqs.Header.Set("Accept", "")
	_, err = server.Negotiate(rqs, "text/html")
	assert.Error(t, err, "Should not allow no accept headers")

	rqs.Header.Set("Accept", "application/$%^%$^")
	_, err = server.Negotiate(rqs)
	assert.Error(t, err, "Should not allow bad accept header")

	rqs.Header.Set("Accept", "application/json")
	_, err = server.Negotiate(rqs, "text/html")
	assert.Error(t, err, "Should not find a MIME type")
}

func TestStarNegotiate(t *testing.T) {
	rqs := new(http.Request)

	rqs.Header = make(map[string][]string)

	rqs.Header.Set("Accept", "*/json")
	mime, err := server.Negotiate(rqs, "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/json", mime.MIME, "Should pick JSON")

	rqs.Header.Set("Accept", "image/*")
	mime, err = server.Negotiate(rqs, "image/png")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "image/png", mime.MIME, "Should pick PNG")

	rqs.Header.Set("Accept", "*/*")
	mime, err = server.Negotiate(rqs, "text/html")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "text/html", mime.MIME, "Should pick HTML")

	rqs.Header.Set("Accept", "*/json, */xml; q=0.4")
	mime, err = server.Negotiate(rqs, "appliction/xml")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "appliction/xml", mime.MIME, "Should pick HTML")
	assert.Equal(t, "0.4", mime.Parameters["q"], "Should maintain parameter")

	rqs.Header.Set("Accept", "image/png, image/*; q=0.5")
	mime, err = server.Negotiate(rqs, "image/jpg")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "image/jpg", mime.MIME, "Should pick JPEG")
	assert.Equal(t, "0.5", mime.Parameters["q"], "Should maintain parameter")

	rqs.Header.Set("Accept", "text/html, */*; q=0.3")
	mime, err = server.Negotiate(rqs, "application/xml")
	assert.Equal(t, "application/xml", mime.MIME, "Should pick XML")
	assert.Equal(t, "0.3", mime.Parameters["q"], "Should maintain parameter")
}

func TestTwoNegotiate(t *testing.T) {
	rqs := new(http.Request)

	rqs.Header = make(map[string][]string)

	rqs.Header.Set("Accept", "text/html")
	mime, err := server.Negotiate(rqs, "text/html", "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "text/html", mime.MIME, "Should pick HTML")

	rqs.Header.Set("Accept", "application/json")
	mime, err = server.Negotiate(rqs, "text/html", "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/json", mime.MIME, "Should pick JSON")
}

func TestFirefoxNegotiate(t *testing.T) {
	rqs := new(http.Request)

	rqs.Header = make(map[string][]string)

	rqs.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")

	mime, err := server.Negotiate(rqs, "text/html", "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "text/html", mime.MIME, "Should pick HTML")

	mime, err = server.Negotiate(rqs, "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/json", mime.MIME, "Should pick JSON")

	mime, err = server.Negotiate(rqs, "application/xml")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/xml", mime.MIME, "Should pick XML")
	assert.Equal(t, "0.9", mime.Parameters["q"], "Should maintain parameter")

	mime, err = server.Negotiate(rqs, "application/vnd.space.danield.article+json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/vnd.space.danield.article+json", mime.MIME, "Should maintain optional params even with any")
	assert.Equal(t, "0.8", mime.Parameters["q"], "Should maintain parameter")
}

func TestWebkitNegotiate(t *testing.T) {
	rqs := new(http.Request)

	rqs.Header = make(map[string][]string)

	rqs.Header.Set("Accept", "application/xml,application/xhtml+xml,text/html;q=0.9,\ntext/plain;q=0.8,image/png,*/*;q=0.5")

	mime, err := server.Negotiate(rqs, "text/html", "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "text/html", mime.MIME, "Should pick HTML")

	mime, err = server.Negotiate(rqs, "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/json", mime.MIME, "Should pick JSON")

	mime, err = server.Negotiate(rqs, "application/xml")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/xml", mime.MIME, "Should pick XML")

	mime, err = server.Negotiate(rqs, "application/vnd.space.danield.article+json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/vnd.space.danield.article+json", mime.MIME, "Should maintain optional params even with any")
	assert.Equal(t, "0.5", mime.Parameters["q"], "Should maintain parameter")
}

func TestIENegotiate(t *testing.T) {
	rqs := new(http.Request)

	rqs.Header = make(map[string][]string)

	rqs.Header.Set("Accept", "image/jpeg, application/x-ms-application, image/gif,\napplication/xaml+xml, image/pjpeg, application/x-ms-xbap,\napplication/x-shockwave-flash, application/msword, */*")

	mime, err := server.Negotiate(rqs, "text/html", "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "text/html", mime.MIME, "Should pick HTML")

	mime, err = server.Negotiate(rqs, "application/json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/json", mime.MIME, "Should pick JSON")

	mime, err = server.Negotiate(rqs, "application/xml")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/xml", mime.MIME, "Should pick XML")

	mime, err = server.Negotiate(rqs, "application/vnd.space.danield.article+json")
	require.NoError(t, err, "No error should occur")
	assert.Equal(t, "application/vnd.space.danield.article+json", mime.MIME, "Should maintain optional params even with any")
}
