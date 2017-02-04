// +build integration

package account_test

import (
	"testing"

	"github.com/danield21/danield-space/pkg/controllers/account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/appengine/aetest"
)

func TestIsAdmin(t *testing.T) {
	context, done, err := aetest.NewContext()
	require.NoError(t, err, "Error occurred in creating mock context")
	defer done()

	isAdmin := account.IsAdmin(context, "Root", []byte("ThisAVerySimplePassword!"))
	assert.True(t, isAdmin, "Tests if we can enter using default account")

	accounts, err := account.GetAll(context)
	if assert.NoError(t, err, "Error in getting all accounts") {
		assert.Equal(t, 1, len(accounts), "Should add a root account incase there is none")
	}

	isAdmin = account.IsAdmin(context, "r00t", []byte("ThisAVerySimplePassword"))
	assert.False(t, isAdmin, "Shouldn't allow a wrong username")

	isAdmin = account.IsAdmin(context, "Root", []byte("WrongPassword"))
	assert.False(t, isAdmin, "Shouldn't allow a wrong password")
}
