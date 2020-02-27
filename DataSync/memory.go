package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
)


var (
lastSavingPoint *uint64

)

var Records = make(map[string]DomainRecord)

type  DomainRecord struct{
	Bas_Ethereum.DNSRecord
	Bas_Ethereum.AssetRecord
}

