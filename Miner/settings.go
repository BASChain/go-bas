package Miner

import "math/big"

type RootAllocation struct {
	ToAdmin *big.Int
	ToBurn  *big.Int
	ToMiner *big.Int
}

type SubAllocation struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}




func Settings() {
	BasMiner().CustomedSubSetting()
}
