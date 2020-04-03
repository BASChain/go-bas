package Miner

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)


/* ---About Allocation---
0  RootSetting.ToAdmin
1  RootSetting.ToBurn
2  RootSetting.ToMiner

3  DefaultSubSetting.ToAdmin
4  DefaultSubSetting.ToBurn
5  DefaultSubSetting.ToMiner
6  DefaultSubSetting.ToRootOwner

7  SelfSubSetting.ToAdmin
8  SelfSubSetting.ToBurn
9  SelfSubSetting.ToMiner
10 SelfSubSetting.ToRootOwner

11 CustomedSubSetting.ToAdmin
12 CustomedSubSetting.ToBurn
13 CustomedSubSetting.ToMiner
14 CustomedSubSetting.ToRootOwner
*/
type Setting struct {
	Allocation [15]big.Int
	Miners []common.Address
}

var SettingRecords = &Bas_Ethereum.Gallery{}
