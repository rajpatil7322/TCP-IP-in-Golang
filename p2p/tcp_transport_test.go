package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransaport(t *testing.T) {
	tr := NewTCPTransport(":9000")
	assert.Equal(t, tr.listenAddress, ":9000")

	assert.Nil(t, tr.ListenAndAccept())

}
