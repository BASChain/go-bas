package Notification

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var OwnershipUpdateNotifications = make([]OwnershipUpdate,0)
var OwnershipAddNotifications = make([]OwnershipAdd,0)


/*
please check if hash exists, and act accordingly,
if hash exists, you should drop any record of old owner and insert this record,
and don't update commitBlock because hash is mint already.
if hash does not exist, simple add record is ok
*/
type OwnershipUpdate func(
	hash Bas_Ethereum.Hash,
	owner common.Address,
	expire big.Int,
	commitBlock uint64)

type OwnershipAdd func(
	hash Bas_Ethereum.Hash,
	owner common.Address,
	expire big.Int,
	commitBlock uint64)


