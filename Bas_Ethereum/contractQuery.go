package Bas_Ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)


func Hash(key string) [32]byte{
	hash := solsha3.SoliditySHA3(solsha3.String(key))
	var ret [32]byte
	for i:=0;i<32;i++ {
		ret[i] = hash[i]
	}
	return ret
}

type  DNSRecord struct {
	Name   []byte
	Expire *big.Int
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
	Owner  common.Address
}

func QueryDomain(key string)  (DNSRecord,error) {
	return Asset.DnsDetailsByHash(nil, Hash(key))
}

