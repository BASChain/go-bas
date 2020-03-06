package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)


var lastSavingPoint = uint64(0)

var lock = &sync.Mutex{}

var Records = make(map[Bas_Ethereum.Hash]DomainRecord)
var Assets = make(map[common.Address][]Bas_Ethereum.Hash)

type  DomainRecord struct{
	dns *Bas_Ethereum.DNSRecord
	asset *Bas_Ethereum.AssetRecord
}

func MemLock()  {
	lock.Lock()
}

func MemUnlock()  {
	lock.Unlock()
}

func (dr *DomainRecord)GetName() string  {
	if dr.asset == nil{
		return ""
	}
	return string(dr.asset.Name)
}

func (dr *DomainRecord)GetExpire() int64  {
	if dr.asset == nil{
		return 0
	}

	return dr.asset.Expire.Int64()
}

func (dr *DomainRecord)GetOpenStatus() bool  {
	if dr.asset == nil{
		return false
	}
	return dr.asset.ROpenToPublic
}

func updateAsset(hash Bas_Ethereum.Hash){
	lock.Lock()
	defer lock.Unlock()
	record,err:=Bas_Ethereum.QueryAssetInfo(hash)
	if err!=nil{
		logger.Error("query asset error : " , err )
		return
	}
	// not exist
	if Records[hash].asset==nil {
		rcd := DomainRecord{
			asset: &record,
			dns:  Records[hash].dns,
		}
		Records[hash] = rcd
		Assets[record.Owner] = append(Assets[rcd.asset.Owner], hash)
	}else{
		oldOwner:= Records[hash].asset.Owner
		rcd := DomainRecord{
			asset: &record,
			dns:  Records[hash].dns,
		}
		Records[hash] = rcd
		if oldOwner!=record.Owner {  // ownership changed
			Assets[record.Owner] = append(Assets[rcd.asset.Owner], hash)
			Assets[oldOwner] = deleteFromStringList(Assets[oldOwner],hash)
		}
	}
}

func updateDNS(hash Bas_Ethereum.Hash){
	lock.Lock()
	defer lock.Unlock()
	record,err:=Bas_Ethereum.QueryDNSInfo(hash)
	if err!=nil{
		logger.Error("query dns error : " , err )
		return
	}
	Records[hash] = DomainRecord{
		dns:   &record,
		asset: Records[hash].asset,
	}
}