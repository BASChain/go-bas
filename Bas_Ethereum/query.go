package Bas_Ethereum

import (
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"strconv"
	"time"
)

type Hash [32]byte

func (hash Hash)String() string{
	return hex.EncodeToString(hash[:])
}

func GetHash(key string) Hash{
	hash := solsha3.SoliditySHA3(solsha3.String(key))
	var ret [32]byte
	for i:=0;i<32;i++ {
		ret[i] = hash[i]
	}
	return ret
}

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

func GetOpts(blockNumber uint64) *bind.CallOpts {
	var opts = new(bind.CallOpts)
	if blockNumber == 0 {
		opts = nil
	}else{
		opts.BlockNumber = new(big.Int).SetUint64(blockNumber)
	}
	return opts
}

var retryRule = map[int]int{
	1:5,
	2:10,
	3:15,
}

func QueryDNS(key Hash, blockNumber uint64)  (DNSRecord,error){
	return _QueryDNS(key,blockNumber,0)
}
func _QueryDNS(key Hash, blockNumber uint64, tryTimes int)  (DNSRecord,error){
	dns,err:=BasDNS().DNS(GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return dns,err
		}else{
			logger.Error("query dns error : " , err , "retry in "+ strconv.Itoa(retryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(retryRule[tryTimes])*time.Second)
			return _QueryDNS(key,blockNumber,tryTimes)
		}
	}else{
		return dns,nil
	}
}

func QueryRoot(key Hash,blockNumber uint64) (RootRecord,error) {
	return _QueryRoot(key,blockNumber,0)
}
func _QueryRoot(key Hash,blockNumber uint64, tryTimes int) (RootRecord,error) {
	root,err:=BasAsset().Root(GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return root,err
		}else{
			logger.Error("query root error : " , err , "retry in "+ strconv.Itoa(retryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(retryRule[tryTimes])*time.Second)
			return _QueryRoot(key,blockNumber,tryTimes)
		}
	}else{
		return root,nil
	}
}

func QuerySub(key Hash,blockNumber uint64) (SubRecord,error) {
	return _QuerySub(key,blockNumber,0)
}
func _QuerySub(key Hash,blockNumber uint64, tryTimes int) (SubRecord,error) {
	sub,err:=BasAsset().Sub(GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return sub,err
		}else{
			logger.Error("query sub error : " , err , "retry in "+ strconv.Itoa(retryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(retryRule[tryTimes])*time.Second)
			return _QuerySub(key,blockNumber,tryTimes)
		}
	}else{
		return sub,nil
	}
}

func QueryOwnership(key Hash,blockNumber uint64) (common.Address, *big.Int, error){
	return _QueryOwnership(key,blockNumber,0)
}
func _QueryOwnership(key Hash,blockNumber uint64, tryTimes int) (common.Address, *big.Int, error){
	newOwner,expire,err := BasOwnership().Query(GetOpts(blockNumber),key)
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return newOwner,expire,err
		}else{
			logger.Error("query ownership error : " , err , "retry in "+ strconv.Itoa(retryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(retryRule[tryTimes])*time.Second)
			return _QueryOwnership(key,blockNumber,tryTimes)
		}
	}else{
		return newOwner,expire,nil
	}
}

var blockTimeMapping = make(map[uint64]uint64)

func GetTimestamp(blockNumber uint64) (uint64,error){
	return _GetTimestamp(blockNumber,0)
}
func _GetTimestamp(blockNumber uint64, tryTimes int) (uint64,error){
	if blockTimeMapping[blockNumber] != 0 {
		return blockTimeMapping[blockNumber],nil
	}
	header,err:=GetConn().HeaderByNumber(context.Background(),big.NewInt(int64(blockNumber)))
	if err!=nil{
		tryTimes +=1
		if tryTimes>3{
			return header.Time,err
		}else{
			logger.Error("query blockheader error : " , err , "retry in "+ strconv.Itoa(retryRule[tryTimes])+" seconds")
			time.Sleep(time.Duration(retryRule[tryTimes])*time.Second)
			return _GetTimestamp(blockNumber,tryTimes)
		}
	}else{
		blockTimeMapping[blockNumber] = header.Time
		return header.Time,nil
	}
}