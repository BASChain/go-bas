package Bas_Ethereum

func PushQueryThenUpdate(prior chan bool,
						q func(hash Hash, blockNumber uint64) chan interface{},
						u func(data interface{}),
						hash Hash,
						blockNumber uint64) chan bool{
						r := q(hash,blockNumber)
						next := make(chan bool)
						defer func() {
						go func() {
							d := <- r
							_ = <- prior
							u(d)
							next <- true
						}()
						}()
						return next
}