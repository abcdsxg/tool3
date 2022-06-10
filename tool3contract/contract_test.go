package tool3contract

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeccak256FuncHex(t *testing.T) {
	res := Keccak256FuncHex(`transfer(address,uint256)`)
	assert.Equal(t, "0xa9059cbb", res)

	res = Keccak256FuncHex(`greet2(uint256)`)
	assert.Equal(t, "0xf9220889", res)
}
