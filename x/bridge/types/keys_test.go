package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOperatorLastClaimNonceKey(t *testing.T) {
	k := GetOperatorLastClaimNonceKey([]byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1})
	expected := []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	require.EqualValues(t, expected, k)
}
