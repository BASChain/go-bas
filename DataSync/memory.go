package DataSync

import (
	"encoding/binary"
	"encoding/hex"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/Notification"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"net"
	"sync"
)

var DebugFlag = true

var lock = &sync.Mutex{}
var pLock = &sync.Mutex{}

var Records = make(map[Bas_Ethereum.Hash]*DomainRecord)
var Assets = make(map[common.Address][]Bas_Ethereum.Hash)
var PayRecords = make(map[Bas_Ethereum.Hash]Receipt)


type DomainRecord struct{
	Name          []byte
	Expire        big.Int
	Owner         common.Address
	IsRoot        bool
	ROpenToPublic bool
	RIsCustomed   bool
	RIsRare      bool
	RCustomPrice  big.Int
	SRootHash     [32]byte
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
	CommitBlock uint64
}

type Receipt struct {
	Payer  common.Address
	Name   []byte
	Option string
	Amount *big.Int
	CommitBlock uint64
}

func showMemory(hash Bas_Ethereum.Hash){
	if DebugFlag{
		logger.Info(string(Records[hash].Name),"0x"+hex.EncodeToString(hash[:]))
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
	if dr == nil{
		return 0
	}
	
	return binary.BigEndian.Uint32(dr.Ipv4[:])
}

func (dr *DomainRecord)GetIPv4Str() string  {
	if dr == nil{
		return ""
	}

	if dr.GetIPv4() == 0{
		return ""
	}

	ip:=net.IPv4(dr.Ipv4[0],dr.Ipv4[1],dr.Ipv4[2],dr.Ipv4[3])

	return ip.String()
}

func (dr *DomainRecord)GetIPv4Addr() [4]byte  {
	if dr == nil{
		return [4]byte{}
	}

	return dr.Ipv4
}

func (dr *DomainRecord)GetBCAddrStr() string {
	if dr == nil{
		return ""
	}
	if len(dr.Bca) == 0{
		return ""
	}
	return "0x"+hex.EncodeToString(dr.Bca)
}

func (dr *DomainRecord)GetIpv6Str() string  {
	if dr == nil{
		return ""
	}
	ip:=net.IPv6zero

	for i,b:=range dr.Ipv6[:]{
		ip[i] = b
	}

	s := ip.String()
	if s == "::"{
		return ""
	}else{
		return s
	}
}

func (dr *DomainRecord)GetAliasName() string  {
	if dr == nil{
		return ""
	}
	return dr.AliasName
}

func (dr *DomainRecord)GetExtraInfo() string  {
	if dr == nil{
		return ""
	}

	return string(dr.OpData)
}


func (dr *DomainRecord)GetName() string  {
	if dr == nil{
		return ""
	}else{
		return string(dr.Name)
	}
}

func (dr *DomainRecord)GetExpire() int64  {
	if dr == nil{
		return 0
	}

	return dr.Expire.Int64()
}

func (dr *DomainRecord)GetOpenStatus() bool  {
	if dr == nil{
		return false
	}
	return dr.ROpenToPublic
}

func (dr *DomainRecord)GetOwner() string  {
	if dr == nil{
		return ""
	}
	return  dr.Owner.String()

}

func (dr *DomainRecord)GetOwnerOrig() *common.Address {
	if dr == nil{
		return nil
	}

	return &dr.Owner
}


func (dr *DomainRecord)GetIsRoot() bool  {
	if dr == nil{
		return  false
	}
	return dr.IsRoot
}

func (dr *DomainRecord)GetIsCustomed() bool  {
	if dr == nil{
		return  false
	}
	return dr.RIsCustomed
}

func (dr *DomainRecord)GetIsRare() bool  {
	if dr == nil{
		return  false
	}
	return dr.RIsRare
}

func (dr *DomainRecord)GetCustomedPrice() string  {
	if dr == nil{
		return  ""
	}

	price := dr.RCustomPrice

	return price.String()
}

func (dr *DomainRecord)GetParentHash() Bas_Ethereum.Hash  {
	if dr == nil{
		return Bas_Ethereum.Hash{}
	}

	return dr.SRootHash
}

func (dr *DomainRecord)GetRegTime() int64  {
	t,_:=conn.GetTimestamp(dr.CommitBlock)
	return int64(t)
}



func updateByQueryOwnership(hash Bas_Ethereum.Hash, blockNumber uint64){
	newOwner,expire,err:=QueryOwnership(hash,blockNumber)
	if err!=nil{
		logger.Error("update Ownership error: " , err )
		return
	}else{
		lock.Lock()
		defer lock.Unlock()
		if Records[hash] == nil {
			Records[hash] = &DomainRecord{}
		}
		oldOwner:= Records[hash].Owner
		//new mint
		if oldOwner == *new(common.Address){
			Records[hash].CommitBlock = blockNumber
		}
		if oldOwner!= newOwner{
			Records[hash].Owner = newOwner
			Assets[newOwner] = append(Assets[newOwner], hash)
			Assets[oldOwner] = deleteFromStringList(Assets[oldOwner],hash)
		}
		Records[hash].Expire = *expire

		if(Records[hash].CommitBlock==0){
			Records[hash].CommitBlock = firstAppearInBlock[hash]
		}

		Notification.NotifyOwnershipUpdate(
			hash,
			newOwner,
			*expire,
			blockNumber)
	}
}

func updateByQueryRoot(hash Bas_Ethereum.Hash, blockNumber uint64)  {
	root,err:=QueryRoot(hash,blockNumber)
	if err!=nil{
		logger.Error("update root error : " , err)
		return
	}else{
		lock.Lock()
		defer lock.Unlock()
		if Records[hash] == nil {
			Records[hash] = &DomainRecord{}
		}
		Records[hash].RCustomPrice = *root.CustomedPrice
		Records[hash].RIsCustomed = root.IsCustomed
		Records[hash].Name = root.RootName
		Records[hash].IsRoot = true
		Records[hash].RIsRare = IsRare(string(root.RootName))
		Records[hash].ROpenToPublic = root.OpenToPublic

		Notification.NotifyAssetRootChanged(
			hash,
			root.OpenToPublic,
			root.IsCustomed,
			Records[hash].RIsRare,
			*root.CustomedPrice)
	}
}

func updateByQuerySub(hash Bas_Ethereum.Hash, blockNumber uint64) {
	sub,err:=QuerySub(hash,blockNumber)
	if err!=nil{
		logger.Error("update root error : " , err)
		return
	}else{
		lock.Lock()
		defer lock.Unlock()
		if Records[hash] == nil {
			Records[hash] = &DomainRecord{}
		}
		Records[hash].SRootHash = sub.RootHash
		Records[hash].Name = sub.TotalName

		Notification.NotifyAssetSubChanged(
			hash,
			sub.RootHash)
	}
}

func updateByQueryDNS(hash Bas_Ethereum.Hash,blockNumber uint64){
	dns,err:=QueryDNS(hash,blockNumber)
	if err!=nil{
		logger.Error("update dns error : " , err )
		return
	}else {
		lock.Lock()
		defer lock.Unlock()
		if Records[hash] == nil {
			Records[hash] = &DomainRecord{}
		}
		Records[hash].Ipv4 = dns.Ipv4
		Records[hash].Ipv6 = dns.Ipv6
		Records[hash].Bca = dns.Bca
		Records[hash].AliasName = dns.AliasName
		Records[hash].OpData = dns.OpData

		Notification.NotifyDNSChanged(
			hash,
			dns.Ipv4,
			dns.Ipv6,
			dns.Bca,
			dns.OpData,
			dns.AliasName)
	}
}

func updatePaid(receipt Receipt,ReceiptNumber Bas_Ethereum.Hash){
	pLock.Lock()
	defer pLock.Unlock()
	PayRecords[ReceiptNumber] = receipt

	Notification.NotifyPaid(
		receipt.Payer,
		receipt.Name,
		receipt.Option,
		*receipt.Amount,
		receipt.CommitBlock)
}