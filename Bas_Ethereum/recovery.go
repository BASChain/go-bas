package Bas_Ethereum

import (
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
	"golang.org/x/net/context"
)
var logger, _ = logging.GetLogger("Bas_Ethereum")

var (

	AccessPoints = []string{
		"wss://ropsten.infura.io/ws/v3/831ab04fa4964991b5fba5c52106d7b0",
		"wss://ropsten.infura.io/ws/v3/8b8db3cca50a4fcf97173b7619b1c4c3",
		"ws://75.135.96.248:3334",
	}


	_t = "0x9d0314f9Bacd569DCB22276867AAEeE1C8A87614"
	_o = "0xB0f032d36461D8E29eD69647C23d4bb0015714b7"
	_a = "0x562660af1b94A5542Bc23Fdb20A6D996154de8Bb"
	_d = "0x6FA6f44aE5050b8aBf50145170226D193BB66d82"
	_m = "0xBC5Ef359eE25E8767bb11791ECe099909603DE58"
	_r = "0xc8138c752f15711e042839F0091FcC451A704442"
	_oann = "0x9542aD4e4B98a6050301cb10E3731E8FA2Fa3E39"
	_market = "0x9f306071D597Aa34F4223dB2dE408Af66aA6B67E"
	_f = "0x6B80cA84B64e2Fd7957138eA7480FB4E775b997F"

	conn *ethclient.Client

	token *Contract.BasToken
	miner *Contract.BasMiner
	asset *Contract.BasAsset
	dns *Contract.BasDNS
	oann  *Contract.BasOANN
	ownership *Contract.BasOwnership
	market *Contract.BasMarket
	free  *Contract.SendFreeBas
)



func RestoreConnectionManually(entrance string) {
	c, err := ethclient.Dial(entrance)
	if err!=nil{
		logger.Error("can't connect to ethereum")
	}else{
		conn = c
		logger.Info("conn is ready")
	}
}

func GetConn() *ethclient.Client {
	if conn!=nil{
		return conn
	}else{
		for _,s := range AccessPoints{
			c, err := ethclient.Dial(s)
			if err!=nil{
				logger.Error("can't get access through : " ,s)
				continue
			}else{
				conn = c
				logger.Info("conn is ready")
				return conn
			}
		}
	}
	logger.Fatal("can't get access to ethereum through any points")
	return nil
}

func ResetConnAndContracts(){
	conn.Close()
	conn = nil
	token = nil
	miner = nil
	asset = nil
	oann = nil
	GetConn()
}

func GetLastBlockNumber() uint64{
	b,e:= GetConn().BlockByNumber(context.Background(),nil)
	if e==nil{
		return  b.NumberU64()
	}else {
		logger.Error("can't get last block",e)
		return 0
	}
}

func BasToken() *Contract.BasToken{
	if token==nil{
		if t,err:=Contract.NewBasToken(common.HexToAddress(_t), GetConn());err==nil{
			token = t
		}else{
			logger.Error("can't recover BasToken",err)
		}
	}
	return token
}

func BasOwnership() *Contract.BasOwnership  {
	if oann==nil{
		if o,err:=Contract.NewBasOwnership(common.HexToAddress(_o), GetConn());err==nil{
			ownership = o
		}else{
			logger.Error("can't recover ownership",err)
		}
	}
	return ownership
}

func BasAsset() *Contract.BasAsset{
	if asset==nil{
		if a,err:=Contract.NewBasAsset(common.HexToAddress(_a), GetConn());err==nil{
			asset = a
		}else{
			logger.Error("can't recover Asset",err)
		}
	}
	return asset
}

func BasDNS() *Contract.BasDNS{
	if asset==nil{
		if d,err:=Contract.NewBasDNS(common.HexToAddress(_d), GetConn());err==nil{
			dns = d
		}else{
			logger.Error("can't recover dns",err)
		}
	}
	return dns
}

func BasMiner() *Contract.BasMiner{
	if miner==nil{
		if m,err:=Contract.NewBasMiner(common.HexToAddress(_m), GetConn());err==nil{
			miner = m
		}else{
			logger.Error("can't recover Miner",err)
		}
	}
	return miner
}

func BasOANN() *Contract.BasOANN{
	if oann==nil{
		if o,err:=Contract.NewBasOANN(common.HexToAddress(_oann), GetConn());err==nil{
			oann = o
		}else{
			logger.Error("can't recover OANN",err)
		}
	}
	return oann
}

func BasMarket() *Contract.BasMarket{
	if oann==nil{
		if m,err:=Contract.NewBasMarket(common.HexToAddress(_market), GetConn());err==nil{
			market = m
		}else{
			logger.Error("can't recover market",err)
		}
	}
	return market
}

func FreeBas() *Contract.SendFreeBas{
	if free==nil{
		if o,err:=Contract.NewSendFreeBas(common.HexToAddress(_f),GetConn());err==nil{
			free = o
		}else{
			logger.Error("can't recover Send Free Bas", err)
		}
	}
	return free
}















