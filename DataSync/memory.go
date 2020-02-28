package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
)


var lastSavingPoint = uint64(0)

var Records = make(map[string]DomainRecord)
var Assets = make(map[common.Address][]string)

type  DomainRecord struct{
	dns *Bas_Ethereum.DNSRecord
	asset *Bas_Ethereum.AssetRecord
}

//todo func update dns record
//todo func update asset record