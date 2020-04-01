package Bas_Ethereum

import (
	"encoding/hex"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

type Hash [32]byte

func (hash Hash)String() string{
	return "0x"+hex.EncodeToString(hash[:])
}

func GetHash(key string) Hash{
	hash := solsha3.SoliditySHA3(solsha3.String(key))
	var ret [32]byte
	for i:=0;i<32;i++ {
		ret[i] = hash[i]
	}
	return ret
}

var AccessPoints = []string{
"wss://ropsten.infura.io/ws/v3/831ab04fa4964991b5fba5c52106d7b0",
"wss://ropsten.infura.io/ws/v3/8b8db3cca50a4fcf97173b7619b1c4c3",
"ws://75.135.96.248:3334",
}


var RetryRule = map[int]int{
	1:5,
	2:10,
	3:15,
}

