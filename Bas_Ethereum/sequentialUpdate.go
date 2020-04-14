package Bas_Ethereum

func PushQueryThenUpdate(prior chan bool,
						query func(hash Hash, blockNumber uint64) chan interface{},
						update func(data interface{}),
						hash Hash,
						blockNumber uint64) chan bool{
						//query should be in go routine
						r := query(hash,blockNumber)
						next := make(chan bool)
						defer func() {
							go func() {
								d := <- r
								//update should be sequential
								_ = <- prior
								update(d)
								next <- true
							}()
						}()
						return next
}