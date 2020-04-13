package Bas_Ethereum

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/net/context"
	"math/big"
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
"wss://mainnet.infura.io/ws/v3/831ab04fa4964991b5fba5c52106d7b0",
"wss://mainnet.infura.io/ws/v3/8b8db3cca50a4fcf97173b7619b1c4c3",
//"ws://75.135.96.248:3334",
}


var RetryRule = map[int]int{
	1:5,
	2:10,
	3:15,
}

func GetOpts(blockNumber uint64) *bind.CallOpts {
	var opts = new(bind.CallOpts)
	if blockNumber == 0 {
		opts = nil
	}else{
		opts.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}
	return opts
}

func GetLoopOpts(s uint64, e *uint64)  *bind.FilterOpts{
	var opts  = &bind.FilterOpts{
		Start:s,
		End: e,
		Context:context.Background(),
	}
	return opts
}

func GetWatchOpts(s uint64) *bind.WatchOpts{
	var opts = &bind.WatchOpts{
		Start:   &s,
		Context: context.Background(),
	}
	return opts
}