package Miner

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sort"
)

var eq = &Bas_Ethereum.EventQueue{}

func loopOverEventQueue() {
	sort.Sort(eq)
	for _,e := range *eq {
		//logger.Info("@@@@@order :" ,e.BlockNumber, e.TxIndex)
		switch e.EventName {
		case "AllocationChanged":
			handleAllocation(e.EventData)
		case "MinerAdd":
			handleMinerAdd(e.EventData)
		case "MinerRemove":
			handleMinerReplace(e.EventData)
		case "MinerReplace":
			handleMinerReplace(e.EventData)
		default:
			logger.Error("undefined type : ", e.EventName)
		}
	}
	eq = &Bas_Ethereum.EventQueue{}
}