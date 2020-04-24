package handlers

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/hunter1085/notary/tuf/data"
	"github.com/hunter1085/notary/tuf/signed"
	"github.com/hunter1085/notary/tuf/testutils"
)

func mustCopyKeys(t *testing.T, from signed.CryptoService, roles ...data.RoleName) signed.CryptoService {
	cs, err := testutils.CopyKeys(from, roles...)
	require.NoError(t, err)
	return cs
}
