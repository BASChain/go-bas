package Transactions

import (
"context"
"github.com/BASChain/go-bas/Bas_Ethereum"
"github.com/ethereum/go-ethereum/accounts/abi/bind"
"github.com/ethereum/go-ethereum/accounts/keystore"
)

type  DNSRecord struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}

func QueryDNS(key Bas_Ethereum.Hash)  (DNSRecord,error){
	dns,client:=BasDNS()
	defer client.Close()
	record,err:=dns.DNS(Bas_Ethereum.GetOpts(0), key)
	if err!=nil{
		return DNSRecord{},err
	}else{
		return record,nil
	}
}

func ChangeDNSInfo(key *keystore.Key,nameHash Bas_Ethereum.Hash,rcd DNSRecord) error {
	opts:=bind.NewKeyedTransactor(key.PrivateKey)
	oann,client:=BasOANN()
	defer client.Close()
	tx,err:=oann.SetRecord(
		opts,
		nameHash,
		rcd.Ipv4,
		rcd.Ipv6,
		rcd.Bca,
		rcd.OpData,
		rcd.AliasName,
	)
	if err!=nil{
		logger.Error("change DNS info, pack transaction error : ",err)
		return err
	}
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err!=nil{
		logger.Error("change DNS info, send transaction error : ", err)
		return err
	}else{
		_,e:=QueryDNS(nameHash)
		if e!=nil{
			logger.Error("query after change error : ", e)
		} else{
			logger.Info("success, gas used : ",receipt.GasUsed)
		}
	}
	return nil
}