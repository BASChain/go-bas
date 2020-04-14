package Miner

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	Miner uint8 = 0
	RootOwner uint8 = 1
)

type SimplifiedReceipt struct {
	BlockNumber   uint64
	TxIndex       uint
	ReceiptNumber Bas_Ethereum.Hash
	Amount        big.Int
	Allocation    uint8
}

type SimplifiedWithdraw struct {
	BlockNumber   uint64
	TxIndex       uint
	Drawer common.Address
	Amount big.Int
}

var Receipts  []SimplifiedReceipt
var Withdraws []SimplifiedWithdraw
