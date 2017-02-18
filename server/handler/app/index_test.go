package app_test

import (
	"bytes"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"google.golang.org/appengine/aetest"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/handler"
	"github.com/danield21/danield-space/server/handler/app"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndex(t *testing.T) {
	client := http.Client{}
	options := aetest.Options{
		AppID: "danield-space",
		StronglyConsistentDatastore: true,
	}
	instance, err := aetest.NewInstance(&options)
	require.NoError(t, err, "Error in creating instance: %v", err)
	defer instance.Close()

	view := template.New("page/app/index")
	view.Parse("Hello, World!")
	head := template.New("")
	head.Parse("Foo")
	view.AddParseTree("theme/balloon/head", head.Tree)
	foot := template.New("")
	foot.Parse("Bar")
	view.AddParseTree("theme/balloon/footer", foot.Tree)

	ctx, done, err := aetest.NewContext()
	require.NoError(t, err, "Error in creating context")
	defer done()

	e := envir.TestingEnvironment{Templates: view, Ctx: ctx}

	server := httptest.NewServer(handler.Prepare(e, app.Index))
	defer server.Close()

	request, err := instance.NewRequest(http.MethodGet, server.URL, nil)
	require.NoError(t, err, "Error in creating GET request for Index: %v", err)
	request.Header.Add("Content-Type", "text/html")

	response, err := client.Do(request)
	require.NoError(t, err, "Error in sending GET request to Index: %v", err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode, "Expected response status 200, received %s", response.Status)
	assert.Equal(t, "text/html; charset=utf-8", response.Header.Get("Content-Type"))

	require.NotZero(t, response.ContentLength)

	bytes := make([]byte, response.ContentLength)
	_, err = response.Body.Read(bytes)

	require.Equal(t, io.EOF, err, "Did not get full message")
	assert.Equal(t, "FooHello, World!Bar", string(bytes), "The correct message should pop up")
}

func TestIndexHead(t *testing.T) {
	client := &http.Client{}

	view := template.New("pages/app/index")
	view.Parse("Hello, World!")

	ctx, done, err := aetest.NewContext()
	require.NoError(t, err, "TestIndexHead.TestIndex - Error in creating context")
	e := envir.TestingEnvironment{Templates: view, Ctx: ctx}
	defer done()

	server := httptest.NewServer(handler.Prepare(e, app.IndexHeaders))
	defer server.Close()

	request, err := http.NewRequest(http.MethodOptions, server.URL, bytes.NewBuffer(nil))
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	request.Header.Add("Content-Type", "text/html")

	response, err := client.Do(request)
	assert.NoError(t, err, "Error in creating GET request for Index: %v", err)
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode, "Expected response status 200, received %s", response.Status)
	assert.Equal(t, "text/html; charset=utf-8", response.Header.Get("Content-Type"))

	assert.Zero(t, response.ContentLength)
}
