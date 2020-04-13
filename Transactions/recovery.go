package Transactions

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/BASChain/go-bas/DataSync"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("Transactions")

func BasDNS() (*Contract.BasDNS,*ethclient.Client){
	conn:=Bas_Ethereum.NewConn()
	//no watch, always get a new conn
	if o,err:=Contract.NewBasDNS(common.HexToAddress(DataSync.DNSAddr),conn.GetClient());err==nil{
		return o,conn.Client
	}else{
		logger.Error("can't recover dns", err)
		return nil,nil
	}
}

func BasOANN() (*Contract.BasOANN,*ethclient.Client){
	conn:=Bas_Ethereum.NewConn()
	//no watch, always get a new conn
	if o,err:=Contract.NewBasOANN(common.HexToAddress(DataSync.OANNAddr),conn.GetClient());err==nil{
		return o,conn.Client
	}else{
		logger.Error("can't recover oann", err)
		return nil,nil
	}
}