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


func GetFreeBasAmount() (*big.Int,error){
	return Bas_Ethereum.FreeBas().Amount(nil)
}

func SendFreeBas(key *keystore.Key,addr common.Address){
	amount,err:=GetFreeBasAmount()
	if err!=nil{
		logger.Error("can't get Free Bas Amount", err)
		return
	}
	opts:=bind.NewKeyedTransactor(key.PrivateKey)
	tx,err:=Bas_Ethereum.BasToken().Transfer(opts,addr,amount)
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