package handler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danield21/danield-space/server"
	"github.com/danield21/danield-space/server/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestApplyHandler(t *testing.T) {
	tests := []struct {
		Handler handler.Handler
		Message string
	}{
		{mockHandler1, "World"},
		{mockHandler2, "Person"},
	}

	for _, test := range tests {
		e := server.TestingEnvironment{Templates: nil, Ctx: context.TODO()}
		h := handler.Apply(e, test.Handler)
		s := httptest.NewServer(h)
		defer s.Close()

		req, err := http.NewRequest(http.MethodGet, s.URL, nil)
		require.NoError(t, err, "Unable to create request")
		c := http.Client{}
		res, err := c.Do(req)
		require.NoError(t, err, "Unable to perform request")
		defer res.Body.Close()

		assert.NotZero(t, res.ContentLength, "Should have a message")

		bytes := make([]byte, res.ContentLength)
		_, err = res.Body.Read(bytes)

		require.Equal(t, io.EOF, err, "Did not get full message")
		assert.Equal(t, test.Message, string(bytes), "The correct message should pop up")
	}
}

func TestChainHandler(t *testing.T) {
	tests := []struct {
		Handler handler.Handler
		Message string
	}{
		{handler.Chain(mockHandler1, mockLink1), "Hello, World"},
		{handler.Chain(mockHandler1, mockLink2), "Goodbye, World"},
		{handler.Chain(mockHandler2, mockLink1), "Hello, Person"},
		{handler.Chain(mockHandler2, mockLink2), "Goodbye, Person"},
		{handler.Chain(mockHandler1, mockLink1, mockLink1), "Hello, Hello, World"},
		{handler.Chain(mockHandler1, mockLink1, mockLink2), "Goodbye, Hello, World"},
		{handler.Chain(mockHandler1, mockLink2, mockLink1), "Hello, Goodbye, World"},
	}

	for _, test := range tests {
		e := server.TestingEnvironment{Templates: nil, Ctx: context.TODO()}
		h := handler.Apply(e, test.Handler)
		s := httptest.NewServer(h)
		defer s.Close()

		req, err := http.NewRequest(http.MethodGet, s.URL, nil)
		require.NoError(t, err, "Unable to create request")
		c := http.Client{}
		res, err := c.Do(req)
		require.NoError(t, err, "Unable to perform request")
		defer res.Body.Close()

		assert.NotZero(t, res.ContentLength, "Should have a message")

		bytes := make([]byte, res.ContentLength)
		_, err = res.Body.Read(bytes)

		require.Equal(t, io.EOF, err, "Did not get full message")
		assert.Equal(t, test.Message, string(bytes), "The correct message should pop up")
	}
}

func TestPrepareHandler(t *testing.T) {
	e := server.TestingEnvironment{Templates: nil, Ctx: context.TODO()}

	tests := []struct {
		Handler http.HandlerFunc
		Message string
	}{
		{handler.Prepare(e, mockHandler1, mockLink1), "Hello, World"},
		{handler.Prepare(e, mockHandler1, mockLink2), "Goodbye, World"},
		{handler.Prepare(e, mockHandler2, mockLink1), "Hello, Person"},
		{handler.Prepare(e, mockHandler2, mockLink2), "Goodbye, Person"},
		{handler.Prepare(e, mockHandler1, mockLink1, mockLink1), "Hello, Hello, World"},
		{handler.Prepare(e, mockHandler1, mockLink1, mockLink2), "Goodbye, Hello, World"},
		{handler.Prepare(e, mockHandler1, mockLink2, mockLink1), "Hello, Goodbye, World"},
	}

	for _, test := range tests {
		s := httptest.NewServer(test.Handler)
		defer s.Close()

		req, err := http.NewRequest(http.MethodGet, s.URL, nil)
		require.NoError(t, err, "Unable to create request")
		c := http.Client{}
		res, err := c.Do(req)
		require.NoError(t, err, "Unable to perform request")
		defer res.Body.Close()

		assert.NotZero(t, res.ContentLength, "Should have a message")

		bytes := make([]byte, res.ContentLength)
		_, err = res.Body.Read(bytes)

		require.Equal(t, io.EOF, err, "Did not get full message")
		assert.Equal(t, test.Message, string(bytes), "The correct message should pop up")
	}
}

func mockHandler1(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Write([]byte("World"))
	return ctx, nil
}

func mockHandler2(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
	w.Write([]byte("Person"))
	return ctx, nil
}

func mockLink1(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		w.Write([]byte("Hello, "))
		return h(ctx, e, w)
	}
}

func mockLink2(h handler.Handler) handler.Handler {
	return func(ctx context.Context, e handler.Environment, w http.ResponseWriter) (context.Context, error) {
		w.Write([]byte("Goodbye, "))
		return h(ctx, e, w)
	}
}
