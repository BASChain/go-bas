package Bas_Ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
)

type Hash [32]byte

func GetHash(key string) [32]byte{
	hash := solsha3.SoliditySHA3(solsha3.String(key))
	var ret [32]byte
	for i:=0;i<32;i++ {
		ret[i] = hash[i]
	}
	return ret
}

type  DNSRecord struct {
	Name   []byte
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
}

type AssetRecord struct{
	Name          []byte
	Expire        *big.Int
	Owner         common.Address
	IsRoot        bool
	ROpenToPublic bool
	RIsCustomed   bool
	RIsPureA      bool
	RCustomPrice  *big.Int
	SRootHash     [32]byte
}

func QueryDNSInfo(key Hash)  (DNSRecord,error) {
	return BasAsset().DnsDetailsByHash(nil, key)
}

func QueryAssetInfo(key Hash) (AssetRecord,error) {
	return BasAsset().AssetDetailsByHash(nil, key)
}
