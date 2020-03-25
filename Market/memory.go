package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type SellOrder struct {
	price *big.Int
	BlockNumber uint64
}

type AskOrder struct {
	price *big.Int
	protectiveRemainTime *big.Int
	BlockNumber uint64
}

var SellOrders = make(map[common.Address]map[Bas_Ethereum.Hash]*SellOrder)
var AskOrders = make(map[common.Address]map[Bas_Ethereum.Hash]*AskOrder)

//func updateSellOrdersByEvent(addr common.Address, hash Bas_Ethereum.Hash, price *big.Int)  {
//	if SellOrders[addr] == nil {
//		SellOrders[addr] = make(map[Bas_Ethereum.Hash]*big.Int)
//	}
//	SellOrders[addr][hash] = price
//}
//
//func updateAskOrdersByEvent(addr common.Address, hash Bas_Ethereum.Hash, price, protectiveRemainTime *big.Int)  {
//	if AskOrders[addr] == nil {
//		AskOrders[addr] = make(map[Bas_Ethereum.Hash]*AskOrder)
//	}
//	if AskOrders[addr][hash] == nil {
//		AskOrders[addr][hash] = &AskOrder{}
//	}
//	AskOrders[addr][hash].price = price
//	AskOrders[addr][hash].protectiveRemainTime = protectiveRemainTime
//}

