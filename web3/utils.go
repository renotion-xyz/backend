package web3

import (
	"math/big"

	"github.com/holiman/uint256"
)

func uint256ToBytes32(tokenID string) [32]byte {
	x := big.NewInt(0)
	x.SetString(tokenID, 10)
	i, _ := uint256.FromBig(x)
	return i.Bytes32()
}
