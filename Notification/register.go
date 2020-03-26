package Notification

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var OwnershipUpdateNotifications = make(map[string]OwnershipUpdate)
var OwnershipAddNotifications = make(map[string]OwnershipAdd)
var OwnershipExtendNotifications = make(map[string]OwnershipExtend)
var OwnershipTakeoverNotifications = make(map[string]OwnershipTakeover)
var OwnershipTransferNotifications = make(map[string]OwnershipTransfer)
var OwnershipTransferFromNotifications = make(map[string]OwnershipTransferFrom)
var OwnershipRemoveNotifications = make(map[string]OwnershipRemove)

var AssetRootChangedNotifications = make(map[string]AssetRootChanged)
var AssetSubChangedNotifications = make(map[string]AssetSubChanged)

var DNSChangedNotifications = make(map[string]DNSChanged)

var PaidNotifications = make(map[string]Paid)

var MarketSoldBySellNotifications = make(map[string]MarketSoldBySell)
var MarketSoldByAskNotifications = make(map[string]MarketSoldByAsk)
var MarketSellAddedNotifications = make(map[string]MarketSellAdded)
var MarketSellChangedNotifications = make(map[string]MarketSellChanged)
var MarketSellRemovedNotifications = make(map[string]MarketSellRemoved)
var MarketAskAddedNotifications = make(map[string]MarketAskAdded)
var MarketAskChangedNotifications = make(map[string]MarketAskChanged)
var MarketAskRemovedNotifications = make(map[string]MarketAskRemove)



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

/*
extend is time in seconds added to current time (when hash expired) or expire time (hash not expired)
*/
type OwnershipExtend func(
	hash Bas_Ethereum.Hash,
	extend big.Int)

/*
takeover, transfer, transferFrom are similar functions, meaning transfer hash from old owner to new owner
*/
type OwnershipTakeover func(
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address)

type OwnershipTransfer func(
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address)

//by is the approved logic contract address, not very useful to backend server
type OwnershipTransferFrom func(
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address,
	by common.Address)

type OwnershipRemove func(
	hash Bas_Ethereum.Hash)

//there we do not specify which variable is changed, rather offer a whole record
type AssetRootChanged func(
	hash Bas_Ethereum.Hash,
	isOpen bool,
	isCustomized bool,
	isRare bool,
	customizedPrice big.Int)

type AssetSubChanged func(
	hash Bas_Ethereum.Hash,
	rootHash Bas_Ethereum.Hash)

type DNSChanged func(
	hash Bas_Ethereum.Hash,
	IPv4 [4]byte,
	IPv6 [16]byte,
	Bca map[string]byte,
	OpData map[string]byte,
	AliasName string)

//name is domain in ascii, option is pay reason
type Paid func(
	payer common.Address,
	name map[string]byte,
	option string,
	amount big.Int)

type MarketSoldBySell func(
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address,
	price big.Int)

type MarketSoldByAsk func(
	hash Bas_Ethereum.Hash,
	oldOwner common.Address,
	newOwner common.Address,
	price big.Int)

type MarketSellAdded func(
	hash Bas_Ethereum.Hash,
	seller common.Address,
	price big.Int)

type MarketSellChanged func(
	hash Bas_Ethereum.Hash,
	seller common.Address,
	price big.Int)

type MarketSellRemoved func(
	hash Bas_Ethereum.Hash,
	seller common.Address)

type MarketAskAdded func(
	hash Bas_Ethereum.Hash,
	buyer common.Address,
	price big.Int)

type MarketAskChanged func(
	hash Bas_Ethereum.Hash,
	buyer common.Address,
	price big.Int)

type MarketAskRemove func(
	hash Bas_Ethereum.Hash,
	buyer common.Address)

