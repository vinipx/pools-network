package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testEthereumAddress() EthereumAddress {
	return [20]byte{1}
}

func testConsensusAddress() ConsensusAddress {
	return []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
}

func TestConsensusAddress_Size(t *testing.T) {
	require.EqualValues(t, 20, testConsensusAddress().Size())
}

func TestConsensusAddress_Unmarshal(t *testing.T) {
	add := ConsensusAddress{}
	require.NoError(t, add.Unmarshal(testConsensusAddress()))
	require.EqualValues(t, testConsensusAddress(), add)
}

func TestConsensusAddress_MarshalTo(t *testing.T) {
	byts := make([]byte, len(testConsensusAddress()))
	size, err := testConsensusAddress().MarshalTo(byts)
	require.EqualValues(t, 20, size)
	require.NoError(t, err)
	require.EqualValues(t, testConsensusAddress(), byts)
}

func TestEthereumAddress_Size(t *testing.T) {
	require.EqualValues(t, 20, testEthereumAddress().Size())
}

func TestEthereumAddress_Unmarshal(t *testing.T) {
	add := EthereumAddress{}
	toMarshal := testEthereumAddress()
	require.NoError(t, add.Unmarshal(toMarshal[:]))
	require.EqualValues(t, testEthereumAddress(), add)
}

func TestEthereumAddress_MarshalTo(t *testing.T) {
	byts := EthereumAddress{}
	size, err := testEthereumAddress().MarshalTo(byts[:])
	require.EqualValues(t, 20, size)
	require.NoError(t, err)
	require.EqualValues(t, testEthereumAddress(), byts)
}
