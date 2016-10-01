package content

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestType_String(t *testing.T) {
	assert.Equal(t, "text/html", Html.String())
}

func TestType_AddCharset(t *testing.T) {
	assert.Equal(t, "text/html; charset=utf-8", Html.AddCharset("utf-8").String());
}
