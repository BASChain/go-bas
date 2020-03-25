package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
)

var lastSavingPoint = uint64(0)

type SellOrder struct {
	price big.Int
	BlockNumber uint64
}

type AskOrder struct {
	price big.Int
	protectiveRemainTime big.Int
	BlockNumber uint64
}

type SoldWay int

func (way SoldWay)String()string{
	if way==BuyFromSell{
		return "buy from sell"
	}else if way == SellToAsk{
		return "sell to ask"
	}
	return ""
}

const (
	BuyFromSell SoldWay = 0
	SellToAsk   SoldWay = 1
)


type Deal struct {
	nameHash Bas_Ethereum.Hash
	way SoldWay
	oldOwner common.Address
	newOwner common.Address
	agreedPrice big.Int
	BlockNumber uint64
}

var sLock = &sync.Mutex{}
var aLock = &sync.Mutex{}

var SellOrders = make(map[common.Address]map[Bas_Ethereum.Hash]*SellOrder)
var AskOrders = make(map[common.Address]map[Bas_Ethereum.Hash]*AskOrder)
var Sold = make([]Deal,0)

func echoSellOrder(addr common.Address, hash Bas_Ethereum.Hash) string {
	return "user : "+ addr.String() +
		" hash : " +  hash.String() +
		" price : "+ SellOrders[addr][hash].price.String()
}

func echoAskOrder(addr common.Address, hash Bas_Ethereum.Hash) string{
	return  "user : "+ addr.String() +
		" hash : " + hash.String() +
		" price : "+ AskOrders[addr][hash].price.String()
}

func echoSold(deal Deal) string {
	return "asset : " + deal.nameHash.String() +
		" has transfered from : " + deal.oldOwner.String() +
		" to : " + deal.newOwner.String() +
		"by soldway :" + deal.way.String()
}

func updateSellOrdersByEvent(addr common.Address, hash Bas_Ethereum.Hash, price big.Int, blockNumber *uint64)  {
	sLock.Lock()
	defer sLock.Unlock()
	if SellOrders[addr] == nil {
		SellOrders[addr] = make(map[Bas_Ethereum.Hash]*SellOrder)
	}
	if SellOrders[addr][hash] == nil {
		SellOrders[addr][hash] = &SellOrder{}
	}
	SellOrders[addr][hash].price = price
	if blockNumber !=nil {
		SellOrders[addr][hash].BlockNumber = *blockNumber
	}
}

func updateAskOrdersByEvent(addr common.Address, hash Bas_Ethereum.Hash, price, protectiveRemainTime big.Int, blockNumber *uint64)  {
	aLock.Lock()
	defer aLock.Unlock()
	if AskOrders[addr] == nil {
		AskOrders[addr] = make(map[Bas_Ethereum.Hash]*AskOrder)
	}
	if AskOrders[addr][hash] == nil {
		AskOrders[addr][hash] = &AskOrder{}
	}
	AskOrders[addr][hash].price = price
	AskOrders[addr][hash].protectiveRemainTime = protectiveRemainTime
	if blockNumber != nil {
		AskOrders[addr][hash].BlockNumber = *blockNumber
	}
}
