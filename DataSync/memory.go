package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
)


var lastSavingPoint = uint64(0)

var Records = make(map[string]DomainRecord)

type  DomainRecord struct{
	dns *Bas_Ethereum.DNSRecord
	asset *Bas_Ethereum.AssetRecord
}

