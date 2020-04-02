package Bas_Ethereum

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
	"math/big"
	"strconv"
	"sync"
	"time"
)

var logger, _ = logging.GetLogger("Bas_Ethereum")

type Conn struct {
	Client *ethclient.Client
	lock *sync.Mutex
}

func NewConn() *Conn {
	return &Conn{
		Client: nil,
		lock:   &sync.Mutex{},
	}
}

func (conn *Conn) GetClient() *ethclient.Client {
	conn.lock.Lock()
	defer conn.lock.Unlock()
	if conn.Client!=nil{
		return conn.Client
	}else{
		for _,s := range AccessPoints{
			c, err := ethclient.Dial(s)
			if err!=nil{
				logger.Error("can't get access through : " ,s)
				continue
			}else{
				conn.Client = c
				logger.Info("conn is ready")
				return conn.Client
			}
		}
	}
	logger.Fatal("can't get access to ethereum through any points")
	return nil
}

func (conn *Conn) Reset(){
	conn.Client.Close()
	conn.Client = nil
	conn.GetClient()
}

var blockTimeMapping = make(map[uint64]uint64)

var bLock = &sync.Mutex{}

func (conn *Conn) GetTimestamp(blockNumber uint64) (uint64,error){
	return conn._GetTimestamp(blockNumber,0)
}
func (conn *Conn)_GetTimestamp(blockNumber uint64, tryTimes int) (uint64,error){
	if blockTimeMapping[blockNumber] != 0 {
		return blockTimeMapping[blockNumber],nil
	}
	header,err:=conn.GetClient().HeaderByNumber(context.Background(),big.NewInt(int64(blockNumber)))
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return 0,err
		}else{
			logger.Error("query blockheader error : " , err , "retry in "+ strconv.Itoa(RetryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(RetryRule[tryTimes])*time.Second)
			return conn._GetTimestamp(blockNumber,tryTimes)
		}
	}else{
		bLock.Lock()
		defer bLock.Unlock()
		blockTimeMapping[blockNumber] = header.Time
		return header.Time,nil
	}
}

func (conn *Conn)GetLastBlockNumber() uint64{
	b,e:= conn.GetClient().BlockByNumber(context.Background(),nil)
	if e==nil{
		return  b.NumberU64()
	}else {
		logger.Error("can't get last block",e)
		return 0
	}
}

