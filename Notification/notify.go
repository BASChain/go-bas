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

//func OwnershipRemove(
//	hash Bas_Ethereum.Hash){
//	for _, f:=range OwnershipRemoveNotifications {
//
//	}
//}
//
////there we do not specify which variable is changed, rather offer a whole record
//type AssetRootChanged func(
//	hash Bas_Ethereum.Hash,
//	isOpen bool,
//	isCustomized bool,
//	isRare bool,
//	customizedPrice big.Int)
//
//type AssetSubChanged func(
//	hash Bas_Ethereum.Hash,
//	rootHash Bas_Ethereum.Hash)
//
//type DNSChanged func(
//	hash Bas_Ethereum.Hash,
//	IPv4 [4]byte,
//	IPv6 [16]byte,
//	Bca map[string]byte,
//	OpData map[string]byte,
//	AliasName string)
//
////name is domain in ascii, option is pay reason
//type Paid func(
//	payer common.Address,
//	name map[string]byte,
//	option string,
//	amount big.Int)
//
//type MarketSoldBySell func(
//	hash Bas_Ethereum.Hash,
//	oldOwner common.Address,
//	newOwner common.Address,
//	price big.Int)
//
//type MarketSoldByAsk func(
//	hash Bas_Ethereum.Hash,
//	oldOwner common.Address,
//	newOwner common.Address,
//	price big.Int)
//
//type MarketSellAdded func(
//	hash Bas_Ethereum.Hash,
//	seller common.Address,
//	price big.Int)
//
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
