package service_test

import (
	"testing"

	"github.com/danield21/danield-space/server/service"
	"github.com/stretchr/testify/assert"
)

func TestType_String(t *testing.T) {
	assert.Equal(t, "text/html", service.HTML.String())
}

func TestType_AddCharset(t *testing.T) {
	assert.Equal(t, "text/html; charset=utf-8", service.HTML.AddCharset("utf-8").String())
}
