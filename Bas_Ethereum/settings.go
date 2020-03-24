package Bas_Ethereum

import (
	"math/big"
	"regexp"
)

var AROOTGAS *big.Int
var BROOTGAS *big.Int
var SUBGAS *big.Int
var CUSTOMPRICEGAS *big.Int
var MAXYEAR *big.Int
var RARETYPELENGTH *big.Int

func Settings(blockNumber uint64)  {
	if gas,err:=BasOANN().AROOTGAS(GetOpts(blockNumber));err!=nil{
		logger.Fatal("can't get sys param AROOTGAS")
	}else{
		AROOTGAS = gas
	}
	if gas,err:=BasOANN().BROOTGAS(GetOpts(blockNumber));err!=nil{
		logger.Fatal("can't get sys param BROOTGAS")
	}else{
		BROOTGAS = gas
	}
	if gas,err:=BasOANN().SUBGAS(GetOpts(blockNumber));err!=nil{
		logger.Fatal("can't get sys param SUBGAS")
	}else{
		SUBGAS = gas
	}
	if gas,err:=BasOANN().CUSTOMEDPRICEGAS(GetOpts(blockNumber));err!=nil{
		logger.Fatal("can't get sys param CUSTOMPRICEGAS")
	}else{
		CUSTOMPRICEGAS = gas
	}
	if year,err:=BasOANN().MAXYEAR(GetOpts(blockNumber));err!=nil{
		logger.Fatal("can't get sys param MAXYEAR")
	}else{
		MAXYEAR = year
	}
	if length,err:=BasOANN().RARETYPELENGTH(GetOpts(blockNumber));err!=nil{
		logger.Fatal("can't get sys param MAXYEAR")
	}else{
		RARETYPELENGTH = length
	}
}

func IsRare(name string)  bool{
	matched, _ := regexp.MatchString(`^[0-9a-z]{0,`+ RARETYPELENGTH.String() + `}$`, "123456")
	return matched
}