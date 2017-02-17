package handler_test

import (
	"testing"

	"github.com/danield21/danield-space/server/handler"
	"github.com/stretchr/testify/assert"
)

func TestType_String(t *testing.T) {
	assert.Equal(t, "text/html", handler.HTML.String())
}

func TestType_AddCharset(t *testing.T) {
	assert.Equal(t, "text/html; charset=utf-8", handler.HTML.AddCharset("utf-8").String())
}
