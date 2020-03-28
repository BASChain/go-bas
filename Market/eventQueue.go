package Market

import "sort"

type EventQueue []EventWrapper

type EventWrapper struct {
	BlockNumber uint64
	TxIndex uint
	EventName string
	EventData interface{}
}

func (q EventQueue) Len() int {
	return len(q)
}

func (q EventQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q EventQueue) Less(i, j int) bool {
	return (q[i].BlockNumber < q[j].BlockNumber) || (q[i].BlockNumber == q[j].BlockNumber && q[i].TxIndex < q[j].TxIndex)
}

var eq =EventQueue{}

func insertEq(blockNumber uint64, txIndex uint, name string, data interface{}){
	eq = append(eq, EventWrapper{
		BlockNumber: blockNumber,
		TxIndex:     txIndex,
		EventName:   name,
		EventData:   data,
	})
}

func loopOverEventQueue()  {
	sort.Sort(eq)
	for _,e:=range eq {
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
			logger.Error("undefined type")
		}
	}
	eq = EventQueue{}
}