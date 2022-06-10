package tool3contract

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func Keccak256FuncHex(input string) string {
	return "0x" + hex.EncodeToString(crypto.Keccak256([]byte(input)))[:8]
}
