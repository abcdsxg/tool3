package tool3storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	SetIpfsGateway("https://ipfs.eternum.io/ipfs/")
}

func TestGetIPFSContent(t *testing.T) {
	content, err := GetIPFSContent("bafybeifx7yeb55armcsxwwitkymga5xf53dxiarykms3ygqic223w5sk3m")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "Hello from IPFS Gateway Checker\n", string(content))
}
