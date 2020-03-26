package Notification

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var OwnershipUpdateNotifications = make([]OwnershipUpdate,0)

type OwnershipUpdate func(
	hash Bas_Ethereum.Hash,
	owner common.Address,
	expire big.Int,
	commitBlock uint64)

type OwnershipAdd func(
	
	)



