package Bas_Ethereum

import (
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)
var logger, _ = logging.GetLogger("Bas_Ethereum")

const DATASYNC string = "data_sync"
const MARKET string = "market"

var (

	AccessPoints = []string{
		"wss://ropsten.infura.io/ws/v3/831ab04fa4964991b5fba5c52106d7b0",
		"wss://ropsten.infura.io/ws/v3/8b8db3cca50a4fcf97173b7619b1c4c3",
		"ws://75.135.96.248:3334",
	}


	_t = "0x9d0314f9Bacd569DCB22276867AAEeE1C8A87614"
	_o = "0x4b91b82bed39B1d946C9E3BC12ba09C2F22fd3ee"
	_a = "0x2B1110a13183A7045C7BCE3ba0092Ff0de4FD241"
	_d = "0x8951f6B80b880E8A47d0d18000A4c90F288F61a3"
	_m = "0xBC5Ef359eE25E8767bb11791ECe099909603DE58"
	_oann = "0xeC2218E1e40CB4b21a4142Df6ee8a6eAF4Bb4434"
	_market = "0xA32ccce4B7aB28d3Ce40BBa03A2748bCbe4544dB"
	_f = "0x6B80cA84B64e2Fd7957138eA7480FB4E775b997F"

	conns = make(map[string]*ethclient.Client)

	token *Contract.BasToken
	ownership *Contract.BasOwnership
	asset *Contract.BasAsset
	dns *Contract.BasDNS
	miner *Contract.BasMiner
	oann  *Contract.BasOANN
	market *Contract.BasMarket
	free  *Contract.SendFreeBas
)



func RestoreConnectionManually(entrance , service string) {
	c, err := ethclient.Dial(entrance)
	if err!=nil{
		logger.Error("can't connect to ethereum")
	}else{
		conns[service] = c
		logger.Info("conn is ready")
	}
}

func GetConn(service string) *ethclient.Client {
	if conns[service]!=nil{
		return conns[service]
	}else{
		for _,s := range AccessPoints{
			c, err := ethclient.Dial(s)
			if err!=nil{
				logger.Error("can't get access through : " ,s)
				continue
			}else{
				conns[service] = c
				logger.Info(service + " conn is ready")
				return conns[service]
			}
		}
	}
	logger.Fatal("can't get access to ethereum through any points")
	return nil
}


func ResetServiceSync(){
	conns[DATASYNC].Close()
	conns[DATASYNC] = nil
	token = nil
	ownership = nil
	asset = nil
	dns = nil
	miner = nil
	oann = nil
	free = nil
	GetConn(DATASYNC)
}

func ResetServiceMarket(){
	conns[MARKET].Close()
	conns[MARKET] = nil
	market = nil
	GetConn(MARKET)
}

func GetLastBlockNumber(service string ) uint64{
	b,e:= GetConn(service).BlockByNumber(context.Background(),nil)
	if e==nil{
		return  b.NumberU64()
	}else {
		logger.Error("can't get last block",e)
		return 0
	}
}

func BasToken() *Contract.BasToken{
	if token==nil{
		if t,err:=Contract.NewBasToken(common.HexToAddress(_t), GetConn(DATASYNC));err==nil{
			token = t
		}else{
			logger.Fatal("can't recover BasToken",err)
		}
	}
	return token
}

func BasOwnership() *Contract.BasOwnership  {
	if ownership==nil{
		if o,err:=Contract.NewBasOwnership(common.HexToAddress(_o), GetConn(DATASYNC));err==nil{
			ownership = o
		}else{
			logger.Fatal("can't recover ownership",err)
		}
	}
	return ownership
}

func BasAsset() *Contract.BasAsset{
	if asset==nil{
		if a,err:=Contract.NewBasAsset(common.HexToAddress(_a), GetConn(DATASYNC));err==nil{
			asset = a
		}else{
			logger.Fatal("can't recover Asset",err)
		}
	}
	return asset
}

func BasDNS() *Contract.BasDNS{
	if dns==nil{
		if d,err:=Contract.NewBasDNS(common.HexToAddress(_d), GetConn(DATASYNC));err==nil{
			dns = d
		}else{
			logger.Fatal("can't recover dns",err)
		}
	}
	return dns
}

func BasMiner() *Contract.BasMiner{
	if miner==nil{
		if m,err:=Contract.NewBasMiner(common.HexToAddress(_m), GetConn(DATASYNC));err==nil{
			miner = m
		}else{
			logger.Fatal("can't recover Miner",err)
		}
	}
	return miner
}

func BasOANN() *Contract.BasOANN{
	if oann==nil{
		if o,err:=Contract.NewBasOANN(common.HexToAddress(_oann), GetConn(DATASYNC));err==nil{
			oann = o
		}else{
			logger.Fatal("can't recover OANN",err)
		}
	}
	return oann
}

func BasMarket() *Contract.BasMarket{
	if market==nil{
		if m,err:=Contract.NewBasMarket(common.HexToAddress(_market), GetConn(MARKET));err==nil{
			market = m
		}else{
			logger.Fatal("can't recover market",err)
		}
	}
	return market
}

func FreeBas() *Contract.SendFreeBas{
	if free==nil{
		if o,err:=Contract.NewSendFreeBas(common.HexToAddress(_f),GetConn(DATASYNC));err==nil{
			free = o
		}else{
			logger.Fatal("can't recover Send Free Bas", err)
		}
	}
	return free
}















