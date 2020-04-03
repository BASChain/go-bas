package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strconv"
	"time"
)

type  DNSRecord struct {
	Ipv4      [4]byte
	Ipv6      [16]byte
	Bca       []byte
	OpData    []byte
	AliasName string
}

type RootRecord struct{
	RootName      []byte
	OpenToPublic  bool
	IsCustomed    bool
	CustomedPrice *big.Int
}

type SubRecord struct {
	TotalName []byte
	RootHash  [32]byte
}



func QueryDNS(key Bas_Ethereum.Hash, blockNumber uint64)  (DNSRecord,error){
	return _QueryDNS(key,blockNumber,0)
}
func _QueryDNS(key Bas_Ethereum.Hash, blockNumber uint64, tryTimes int)  (DNSRecord,error){
	dns,err:=BasDNS().DNS(Bas_Ethereum.GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return dns,err
		}else{
			logger.Error("query dns error : " , err , "retry in "+ strconv.Itoa(Bas_Ethereum.RetryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(Bas_Ethereum.RetryRule[tryTimes])*time.Second)
			return _QueryDNS(key,blockNumber,tryTimes)
		}
	}else{
		return dns,nil
	}
}

func QueryRoot(key Bas_Ethereum.Hash,blockNumber uint64) (RootRecord,error) {
	return _QueryRoot(key,blockNumber,0)
}
func _QueryRoot(key Bas_Ethereum.Hash,blockNumber uint64, tryTimes int) (RootRecord,error) {
	root,err:=BasAsset().Root(Bas_Ethereum.GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return root,err
		}else{
			logger.Error("query root error : " , err , "retry in "+ strconv.Itoa(Bas_Ethereum.RetryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(Bas_Ethereum.RetryRule[tryTimes])*time.Second)
			return _QueryRoot(key,blockNumber,tryTimes)
		}
	}else{
		return root,nil
	}
}

func QuerySub(key Bas_Ethereum.Hash,blockNumber uint64) (SubRecord,error) {
	return _QuerySub(key,blockNumber,0)
}
func _QuerySub(key Bas_Ethereum.Hash,blockNumber uint64, tryTimes int) (SubRecord,error) {
	sub,err:=BasAsset().Sub(Bas_Ethereum.GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return sub,err
		}else{
			logger.Error("query sub error : " , err , "retry in "+ strconv.Itoa(Bas_Ethereum.RetryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(Bas_Ethereum.RetryRule[tryTimes])*time.Second)
			return _QuerySub(key,blockNumber,tryTimes)
		}
	}else{
		return sub,nil
	}
}

func QueryOwnership(key Bas_Ethereum.Hash,blockNumber uint64) (common.Address, *big.Int, error){
	return _QueryOwnership(key,blockNumber,0)
}
func _QueryOwnership(key Bas_Ethereum.Hash,blockNumber uint64, tryTimes int) (common.Address, *big.Int, error){
	newOwner,expire,err := BasOwnership().Query(Bas_Ethereum.GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return newOwner,expire,err
		}else{
			logger.Error("query ownership error : " , err , "retry in "+ strconv.Itoa(Bas_Ethereum.RetryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(Bas_Ethereum.RetryRule[tryTimes])*time.Second)
			return _QueryOwnership(key,blockNumber,tryTimes)
		}
	}else{
		return newOwner,expire,nil
	}
}

