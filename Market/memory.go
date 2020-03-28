package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/Notification"
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

func (so *SellOrder)GetPrice() *big.Int  {
	return &so.price
}

func (d *Deal)GetFromOwner() string  {
	return d.oldOwner.String()
}

func (d *Deal)GetOwner() string  {
	return d.newOwner.String()
}

func (d *Deal)GetHash() Bas_Ethereum.Hash  {
	return d.nameHash
}

func (d *Deal)GetAGreedPrice() *big.Int  {
	return &d.agreedPrice
}

func (d *Deal)GetTime() int64  {
	t,_:=Bas_Ethereum.GetTimestamp(d.BlockNumber)
	return int64(t)
}


func (so *SellOrder)GetPriceStr() string  {
	return so.price.String()
}
func (so *SellOrder)GetTime() int64  {
	t,_:=Bas_Ethereum.GetTimestamp(so.BlockNumber)
	return int64(t)
}



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

func insertSellOrders(addr common.Address, hash Bas_Ethereum.Hash, price big.Int, blockNumber *uint64)  {
	sLock.Lock()
	defer sLock.Unlock()
	if SellOrders[addr] == nil {
		SellOrders[addr] = make(map[Bas_Ethereum.Hash]*SellOrder)
	}

	SellOrders[addr][hash] = &SellOrder{
		price:       price,
		BlockNumber: *blockNumber,
	}

	Notification.NotifyMarketSellAdded(
		hash,
		addr,
		price,
		*blockNumber)
}


func updateSellOrders(addr common.Address, hash Bas_Ethereum.Hash, price big.Int)  {
	sLock.Lock()
	defer sLock.Unlock()
	SellOrders[addr][hash].price = price
	Notification.NotifyMarketSellChanged(
		hash,
		addr,
		price)
}

func removeSellOrders(addr common.Address, hash Bas_Ethereum.Hash){
	sLock.Lock()
	defer sLock.Unlock()
	delete(SellOrders[addr],hash)
	
	Notification.NotifyMarketSellRemoved(
		hash,
		addr)
}

func insertAskOrders(addr common.Address, hash Bas_Ethereum.Hash, price, protectiveRemainTime big.Int, blockNumber *uint64)  {
	aLock.Lock()
	defer aLock.Unlock()
	if AskOrders[addr] == nil {
		AskOrders[addr] = make(map[Bas_Ethereum.Hash]*AskOrder)
	}
	AskOrders[addr][hash] = &AskOrder{
		price:                price,
		protectiveRemainTime: protectiveRemainTime,
		BlockNumber:          *blockNumber,
	}

	Notification.NotifyMarketAskAdded(
		hash,
		addr,
		price,
		protectiveRemainTime,
		*blockNumber)
}

func updateAskOrders(addr common.Address, hash Bas_Ethereum.Hash, price, protectiveRemainTime big.Int)  {
	aLock.Lock()
	defer aLock.Unlock()
	
	AskOrders[addr][hash].price = price
	AskOrders[addr][hash].protectiveRemainTime = protectiveRemainTime

	Notification.NotifyMarketAskChanged(
		hash,
		addr,
		price,
		protectiveRemainTime)
}

func removeAskOrders(addr common.Address, hash Bas_Ethereum.Hash)  {
	aLock.Lock()
	defer aLock.Unlock()
	delete(AskOrders[addr],hash)

	Notification.NotifyMarketAskRemove(
		hash,
		addr)
}
