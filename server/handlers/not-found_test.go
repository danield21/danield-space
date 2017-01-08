package handlers

import (
	"bytes"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(t *testing.T) {
	client := &http.Client{}

	view := template.New("pages/not-found")
	view.Parse("Hello, World!")

	settings := config{Templates: view}

	server := httptest.NewServer(NotFound(settings))
	defer server.Close()

	request, err := http.NewRequest(http.MethodGet, server.URL, bytes.NewBuffer(nil))
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	request.Header.Add("Content-Type", "text/html")

	response, err := client.Do(request)
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusNotFound, response.StatusCode, "Expected response status 404, received %s", response.Status)
	assert.Equal(t, "text/html; charset=utf-8", response.Header.Get("Content-Type"))

	assert.NotEmpty(t, response.ContentLength)
}
