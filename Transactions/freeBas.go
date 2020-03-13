package Transactions

import (
	"context"
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func CheckIfApplied(addr common.Address) (bool,error){
	return Bas_Ethereum.FreeBas().ApplyRecord(nil,addr)
}


func SendFreeBasByContract(key *keystore.Key,addr common.Address,amount int64){
	opts:=bind.NewKeyedTransactor(key.PrivateKey)
	value := big.NewInt(amount) // in wei (1 eth) := big.NewInt(amount) // in wei (1 eth)
	tx,err:=Bas_Ethereum.FreeBas().SendTokenByContract(opts,addr,value)
	if err!=nil{
		logger.Error("send BAS error : ",err)
		return
	}
	receipt, err := bind.WaitMined(context.Background(), Bas_Ethereum.GetConn(), tx)
	if err!=nil{
		logger.Error("mine send BAS error : ", err)
		return
	}else{
		logger.Info("send free bas to : ",addr.String()," on block : ",receipt.BlockNumber.String())
	}
}