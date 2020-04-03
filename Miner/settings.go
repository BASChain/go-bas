package Miner

import (
	"math/big"
)

type Allocation struct {
	ToAdmin     *big.Int
	ToBurn      *big.Int
	ToMiner     *big.Int
	ToRootOwner *big.Int
}

type AllocationPack struct {
	RootSetting  *Allocation
	DefaultSubSetting *Allocation
	SelfSubSetting *Allocation
	CustomedSubSetting *Allocation
}

//var AllocationData  = &Bas_Ethereum.BBDQ{}
//
//func UpdateAllocSetting(blockNumber uint64) {
//
//	root,err1:=BasMiner().RootSetting(Bas_Ethereum.GetOpts(blockNumber))
//	sub, err2:=BasMiner().DefaultSubSetting(Bas_Ethereum.GetOpts(blockNumber))
//	self, err3:=BasMiner().SelfSubSetting(Bas_Ethereum.GetOpts(blockNumber))
//	customed, err4:=BasMiner().CustomedSubSetting(Bas_Ethereum.GetOpts(blockNumber))
//
//	if err1!=nil || err2!=nil || err3 !=nil || err4 !=nil{
//		logger.Fatal("can't get allocation settings")
//	}
//
//	data:=AllocationPack{
//		RootSetting:        &Allocation{
//			ToAdmin:     root.ToAdmin,
//			ToBurn:      root.ToBurn,
//			ToMiner:     root.ToMiner,
//			ToRootOwner: root.ToRootOwner,
//		},
//		DefaultSubSetting:  &Allocation{
//			ToAdmin:     sub.ToAdmin,
//			ToBurn:      sub.ToBurn,
//			ToMiner:     sub.ToMiner,
//			ToRootOwner: sub.ToRootOwner,
//		},
//		SelfSubSetting:    &Allocation{
//			ToAdmin:     self.ToAdmin,
//			ToBurn:      self.ToBurn,
//			ToMiner:     self.ToMiner,
//			ToRootOwner: self.ToRootOwner,
//		},
//		CustomedSubSetting: &Allocation{
//			ToAdmin:     customed.ToAdmin,
//			ToBurn:      customed.ToBurn,
//			ToMiner:     customed.ToMiner,
//			ToRootOwner: customed.ToRootOwner,
//		},
//	}
//
//	AllocationData.Update(blockNumber,data)
//}
//
//func GetAllocSetting(blockNumber uint64)  AllocationPack{
//	err,data := AllocationData.GetClosestData(blockNumber)
//	if err!=nil{
//		logger.Error(err)
//		return  AllocationPack{}
//	}else{
//		return data.(AllocationPack)
//	}
//}
//
//func InitSetting(){
//	UpdateAllocSetting(7637159)
//	                             //7637159
//}




