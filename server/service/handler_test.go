package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"io"

	"github.com/danield21/danield-space/server/envir"
	"github.com/danield21/danield-space/server/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
)

func TestApplyHandler(t *testing.T) {
	tests := []struct {
		Handler service.Handler
		Message string
	}{
		{mockHandler1, "World"},
		{mockHandler2, "Person"},
	}

	for _, test := range tests {
		e := envir.TestingEnvironment{Templates: nil, Ctx: context.TODO()}
		h := service.Apply(e, test.Handler)
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
		Handler service.Handler
		Message string
	}{
		{service.Chain(mockHandler1, mockLink1), "Hello, World"},
		{service.Chain(mockHandler1, mockLink2), "Goodbye, World"},
		{service.Chain(mockHandler2, mockLink1), "Hello, Person"},
		{service.Chain(mockHandler2, mockLink2), "Goodbye, Person"},
		{service.Chain(mockHandler1, mockLink1, mockLink1), "Hello, Hello, World"},
		{service.Chain(mockHandler1, mockLink1, mockLink2), "Goodbye, Hello, World"},
		{service.Chain(mockHandler1, mockLink2, mockLink1), "Hello, Goodbye, World"},
	}

	for _, test := range tests {
		e := envir.TestingEnvironment{Templates: nil, Ctx: context.TODO()}
		h := service.Apply(e, test.Handler)
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
	e := envir.TestingEnvironment{Templates: nil, Ctx: context.TODO()}

	tests := []struct {
		Handler http.HandlerFunc
		Message string
	}{
		{service.Prepare(e, mockHandler1, mockLink1), "Hello, World"},
		{service.Prepare(e, mockHandler1, mockLink2), "Goodbye, World"},
		{service.Prepare(e, mockHandler2, mockLink1), "Hello, Person"},
		{service.Prepare(e, mockHandler2, mockLink2), "Goodbye, Person"},
		{service.Prepare(e, mockHandler1, mockLink1, mockLink1), "Hello, Hello, World"},
		{service.Prepare(e, mockHandler1, mockLink1, mockLink2), "Goodbye, Hello, World"},
		{service.Prepare(e, mockHandler1, mockLink2, mockLink1), "Hello, Goodbye, World"},
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

func mockHandler1(ctx context.Context, e envir.Environment, w http.ResponseWriter) error {
	w.Write([]byte("World"))
	return nil
}

func mockHandler2(ctx context.Context, e envir.Environment, w http.ResponseWriter) error {
	w.Write([]byte("Person"))
	return nil
}

func mockLink1(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) error {
		w.Write([]byte("Hello, "))
		return h(ctx, e, w)
	}
}

func mockLink2(h service.Handler) service.Handler {
	return func(ctx context.Context, e envir.Environment, w http.ResponseWriter) error {
		w.Write([]byte("Goodbye, "))
		return h(ctx, e, w)
	}
}
