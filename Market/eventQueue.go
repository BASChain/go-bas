package Market

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"sort"
)



var eq = &Bas_Ethereum.EventQueue{}

func loopOverEventQueue()  {
	sort.Sort(*eq)
	for _,e:=range *eq {
		//logger.Info("@@@@@order :" ,e.BlockNumber, e.TxIndex)
		switch e.EventName {
		case "SellAdded":
			handleSellAdded(e.EventData)
		case "SellChanged":
			handleSellChanged(e.EventData)
		case "SellRemoved":
			handleSellRemoved(e.EventData)
		case "AskAdded":
			handleAskAdded(e.EventData)
		case "AskChanged":
			handleAskChanged(e.EventData)
		case "AskRemoved":
			handleAskRemoved(e.EventData)
		case "SoldBySell":
			handleSoldBySell(e.EventData)
		case "SoldByAsk":
			handleSoldByAsk(e.EventData)
		default:
			logger.Error("undefined type : ", e.EventName)
		}
	}
	eq = &Bas_Ethereum.EventQueue{}
}