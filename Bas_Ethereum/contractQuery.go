package Bas_Ethereum

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/net/context"
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
	Ipv4   [4]byte
	Ipv6   [16]byte
	BcAddr string
	OpData []byte
	AName  string
}

func QueryDomain(key string)  (DNSRecord,error) {
	return Asset.DnsDetailsByHash(nil, Hash(key))
}


func QueryEvent(){
	var opts  = bind.FilterOpts{
		Start:0,
		End:nil,
		Context:context.Background(),
	}
	it,err:=Asset.FilterMintAsset(&opts)
	if err==nil{
		for it.Next() {
			fmt.Print(string(it.Event.Name),"\n")
		}
	}
}