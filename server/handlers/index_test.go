package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danield21/danield-space/server/config"
	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	client := &http.Client{}

	view := template.New("page/index")
	view.Parse("Hello, World!")
	head := template.New("")
	head.Parse("Foo\n")
	view.AddParseTree("theme/balloon/head", head.Tree)
	foot := template.New("")
	foot.Parse("Foo\n")
	view.AddParseTree("theme/balloon/footer", foot.Tree)

	settings := config.MockConfig{Templates: view}

	server := httptest.NewServer(Index(settings))
	defer server.Close()

	request, err := http.NewRequest(http.MethodGet, server.URL, bytes.NewBuffer(nil))
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	request.Header.Add("Content-Type", "text/html")

	response, err := client.Do(request)
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode, "Expected response status 200, received %s", response.Status)
	assert.Equal(t, "text/html; charset=utf-8", response.Header.Get("Content-Type"))

	assert.NotEmpty(t, response.ContentLength)
}

func TestIndexHead(t *testing.T) {
	client := &http.Client{}

	view := template.New("pages/index")
	view.Parse("Hello, World!")

	settings := config.MockConfig{Templates: view}

	server := httptest.NewServer(IndexHeaders(settings))
	defer server.Close()

	request, err := http.NewRequest(http.MethodOptions, server.URL, bytes.NewBuffer(nil))
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	request.Header.Add("Content-Type", "text/html")

	response, err := client.Do(request)
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode, "Expected response status 200, received %s", response.Status)
	assert.Equal(t, "text/html; charset=utf-8", response.Header.Get("Content-Type"))

	assert.Empty(t, response.ContentLength)
}
