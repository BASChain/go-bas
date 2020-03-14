package DataSync

import (
	"encoding/hex"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"sync"
	"encoding/binary"
	"time"
)

var DebugFlag = true

var lastSavingPoint = uint64(0)

var lock = &sync.Mutex{}

var Records = make(map[Bas_Ethereum.Hash]DomainRecord)
var Assets = make(map[common.Address][]Bas_Ethereum.Hash)

type  DomainRecord struct{
	dns *Bas_Ethereum.DNSRecord
	asset *Bas_Ethereum.AssetRecord
}

func showMemeory(hash Bas_Ethereum.Hash){
	if DebugFlag{
		logger.Info(string(Records[hash].asset.Name),"0x"+hex.EncodeToString(hash[:]))
	}
}

func SetLastSavingPoint(bn uint64){
	lastSavingPoint = bn
}

func MemLock()  {
	lock.Lock()
}

func MemUnlock()  {
	lock.Unlock()
}

func (dr *DomainRecord)GetIPv4() uint32  {
	if dr.dns == nil{
		return 0
	}

	return binary.BigEndian.Uint32(dr.dns.Ipv4[:])
}

func (dr *DomainRecord)GetIPv4Addr() [4]byte  {
	if dr.dns == nil{
		return [4]byte{}
	}

	return dr.dns.Ipv4
}

func (dr *DomainRecord)GetBCAddr() string {
	return ""
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

func (dr *DomainRecord)GetOwner() string  {
	if dr.asset == nil{
		return ""
	}
	return  dr.asset.Owner.String()

}

func (dr *DomainRecord)GetIsRoot() bool  {
	if dr.asset == nil{
		return  false
	}
	return dr.asset.IsRoot
}

func (dr *DomainRecord)GetIsCustomed() bool  {
	if dr.asset == nil{
		return  false
	}
	return dr.asset.RIsCustomed
}

func (dr *DomainRecord)GetIsPureA() bool  {
	if dr.asset == nil{
		return  false
	}
	return dr.asset.RIsPureA
}

func (dr *DomainRecord)GetCustomedPrice() string  {
	if dr.asset == nil{
		return  ""
	}

	price := dr.asset.RCustomPrice

	return price.String()
}

func (dr *DomainRecord)GetParentHash() Bas_Ethereum.Hash  {
	if dr.asset == nil{
		return Bas_Ethereum.Hash{}
	}

	return dr.asset.SRootHash
}

func updateAsset(hash Bas_Ethereum.Hash,blockNumber uint64){
	record,err:=Bas_Ethereum.QueryAssetInfo(hash,blockNumber)
	if err!=nil{
		logger.Error("query asset error : " , err , "retry in 30 seconds")
		time.Sleep(time.Duration(30)*time.Second)
		updateAsset(hash,blockNumber)
		return
	}
	lock.Lock()
	defer lock.Unlock()
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

func updateDNS(hash Bas_Ethereum.Hash,blockNumber uint64){
	record,err:=Bas_Ethereum.QueryDNSInfo(hash,blockNumber)
	if err!=nil{
		logger.Error("query dns error : " , err , "retry in 30 seconds")
		time.Sleep(time.Duration(30)*time.Second)
		updateDNS(hash,blockNumber)
		return
	}
	lock.Lock()
	defer lock.Unlock()
	Records[hash] = DomainRecord{
		dns:   &record,
		asset: Records[hash].asset,
	}
}
