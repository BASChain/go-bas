package Notification

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func NotifyOwnershipUpdate(
	hash Bas_Ethereum.Hash,
	owner common.Address,
	expire big.Int,
	commitBlock uint64)  {
	for _,f := range OwnershipUpdateNotifications {
		f(hash,owner,expire,commitBlock)
	}
}

func NotifyOwnershipAdd(
	hash Bas_Ethereum.Hash,
	owner common.Address,
	expire big.Int,
	commitBlock uint64){
	for _,f := range OwnershipAddNotifications {
		f(hash,owner,expire,commitBlock)
	}
}

func NotifyOwnershipExtend (
	hash Bas_Ethereum.Hash,
	extend big.Int){
	for _,f := range OwnershipExtendNotifications {
		f(hash,extend)
	}
}

func NotifyOwnershipTakeover(
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address){
	for _,f :=range OwnershipTakeoverNotifications {
		f(hash,oldOwner,newOwner)
	}
}

func NotifyOwnershipTransfer(
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address){
	for _,f :=range OwnershipTransferNotifications {
		f(hash,oldOwner,newOwner)
	}
}

func NotifyOwnershipTransferFrom (
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address,
	by common.Address){
	for _,f := range OwnershipTransferFromNotifications {
		f(hash,oldOwner,newOwner,by)
	}
}

func NotifyOwnershipRemove(
	hash Bas_Ethereum.Hash){
	for _, f:=range OwnershipRemoveNotifications {
		f(hash)
	}
}

func NotifyAssetRootChanged(
	hash Bas_Ethereum.Hash,
	isOpen bool,
	isCustomized bool,
	isRare bool,
	customizedPrice big.Int){
	for _, f:=range AssetRootChangedNotifications {
		f(hash,isOpen,isCustomized,isRare,customizedPrice)
	}
}

func NotifyAssetSubChanged (
	hash Bas_Ethereum.Hash,
	rootHash Bas_Ethereum.Hash){
	for _,f:=range AssetSubChangedNotifications {
		f(hash,rootHash)
	}
}

func NotifyDNSChanged(
	hash Bas_Ethereum.Hash,
	IPv4 [4]byte,
	IPv6 [16]byte,
	Bca map[string]byte,
	OpData map[string]byte,
	AliasName string){
	for _,f :=range DNSChangedNotifications {
		f(hash,IPv4,IPv6,Bca,OpData,AliasName)
	}
}

func NotifyPaid (
	payer common.Address,
	name map[string]byte,
	option string,
	amount big.Int){
	for _,f :=range PaidNotifications {
		f(payer,name,option,amount)
	}
}

func NotifyMarketSoldBySell (
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address,
	price big.Int){
	for _,f :=range MarketSoldBySellNotifications {
		f(hash,oldOwner,newOwner,price)
	}
}

func NotifyMarketSoldByAsk (
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address,
	price big.Int){
	for _, f := range MarketSoldByAskNotifications {
		f(hash,oldOwner,newOwner,price)
	}
}

//func NotifyMarketSellAdded (
//	hash Bas_Ethereum.Hash,
//	seller common.Address,
//	price big.Int){
//	for _, f := range MarketSellAddedNotifications
//}

//type MarketSellChanged func(
//	hash Bas_Ethereum.Hash,
//	seller common.Address,
//	price big.Int)
//
//type MarketSellRemoved func(
//	hash Bas_Ethereum.Hash,
//	seller common.Address)
//
//type MarketAskAdded func(
//	hash Bas_Ethereum.Hash,
//	buyer common.Address,
//	price big.Int)
//
//type MarketAskChanged func(
//	hash Bas_Ethereum.Hash,
//	buyer common.Address,
//	price big.Int)
//
//type MarketAskRemove func(
//	hash Bas_Ethereum.Hash,
//	buyer common.Address)
