package tool3contract

import (
	"testing"
)

var (
	ethRpcUrl    = "https://rpc.ankr.com/eth"
	baycContract = "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d"
)

func TestTotalSupply(t *testing.T) {
	e, err := NewErc721Contract(ethRpcUrl, baycContract)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(e.TotalSupply())
}

func TestNftMeta(t *testing.T) {
	e, err := NewErc721Contract(ethRpcUrl, baycContract)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(e.NftMeta(1))
}

func TestTokenMeta(t *testing.T) {
	e, err := NewErc721Contract(ethRpcUrl, baycContract)
	if err != nil {
		t.Fatal(err)
	}

	_, _, err = e.NftImage(2078)
	if err != nil {
		t.Fatal(err)
	}
}
