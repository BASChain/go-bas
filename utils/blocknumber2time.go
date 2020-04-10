package utils

import (
	"sync"
	"errors"
	"github.com/BASChain/go-bas/Bas_Ethereum"
)

var(
	conn = Bas_Ethereum.NewConn()
	bn2tInst *BlockNum2Time
	bn2tInstLock sync.Mutex
)

type BlockNum2Time struct {
	c chan uint64

	quit chan int
	wg *sync.WaitGroup
}

func GetBlockNum2Time() *BlockNum2Time {
	if bn2tInst != nil{
		return bn2tInst
	}

	bn2tInstLock.Lock()
	defer bn2tInstLock.Unlock()
	if bn2tInst != nil{
		return bn2tInst
	}

	bn2tInst = NewBlockNum2Time()

	return bn2tInst

}

func NewBlockNum2Time() *BlockNum2Time  {
	bn2t:=&BlockNum2Time{}

	bn2t.c = make(chan uint64,20480)
	bn2t.quit = make(chan int ,1)
	bn2t.wg = &sync.WaitGroup{}

	return bn2t
}

func (bn2t *BlockNum2Time)Run()  {
	bn2t.wg.Add(1)
	defer bn2t.wg.Done()


	for{
		select {
		case bn:=<-bn2t.c:
			go BlockNumnber2TimeStamp(bn)
		case <-bn2t.quit:
			return
		}
	}

	return

}

func (bn2t *BlockNum2Time)Stop()  {
	bn2t.quit <- 0

	bn2t.wg.Wait()
	close(bn2t.c)
	//close(bn2t.quit)
}

func (bn2t *BlockNum2Time)Push(blockNumber uint64) error {
	select {
	case bn2t.c<-blockNumber:
	default:
		return errors.New("push failed")
	}

	return nil
}


func BlockNumnber2TimeStamp(blockNumber uint64) int64  {
	t,_:=conn.GetTimestamp(blockNumber)

	return int64(t)
}