package admin_test

import (
	"testing"

	"github.com/danield21/danield-space/server/controllers/admin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/appengine/aetest"
)

func TestIsAdmin(t *testing.T) {
	context, done, err := aetest.NewContext()
	require.NotNil(t, err, "Error occurred in creating mock context")
	defer done()

	isAdmin := admin.IsAdmin(context, "Root", []byte("ThisAVerySimplePassword"))
	assert.True(t, isAdmin, "Tests if we can enter using default account")

	accounts, err := admin.GetAll(context)
	if assert.NotNil(t, err, "No error should appear") {
		assert.Equal(t, 1, len(accounts), "Should add a root account incase there is none")
	}

	isAdmin = admin.IsAdmin(context, "r00t", []byte("ThisAVerySimplePassword"))
	assert.False(t, isAdmin, "Shouldn't allow a wrong username")

	isAdmin = admin.IsAdmin(context, "Root", []byte("WrongPassword"))
	assert.False(t, isAdmin, "Shouldn't allow a wrong password")
}
