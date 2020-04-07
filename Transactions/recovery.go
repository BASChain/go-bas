package Transactions

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/BASChain/go-bas/DataSync"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	FreeBasAddr = "0x6B80cA84B64e2Fd7957138eA7480FB4E775b997F"
)

func FreeBas() (*Contract.SendFreeBas,*ethclient.Client){
	conn:=Bas_Ethereum.NewConn()
	//no watch, always get a new conn
	if o,err:=Contract.NewSendFreeBas(common.HexToAddress(FreeBasAddr),conn.GetClient());err==nil{
		return o,conn.Client
	}else{
		logger.Error("can't recover Send Free Bas", err)
		return nil,nil
	}
}

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