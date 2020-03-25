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

func Settings()  {
	if gas,err:=BasOANN().AROOTGAS(GetOpts(0));err!=nil{
		logger.Fatal("can't get sys param AROOTGAS")
	}else{
		AROOTGAS = gas
	}
	if gas,err:=BasOANN().BROOTGAS(GetOpts(0));err!=nil{
		logger.Fatal("can't get sys param BROOTGAS")
	}else{
		BROOTGAS = gas
	}
	if gas,err:=BasOANN().SUBGAS(GetOpts(0));err!=nil{
		logger.Fatal("can't get sys param SUBGAS")
	}else{
		SUBGAS = gas
	}
	if gas,err:=BasOANN().CUSTOMEDPRICEGAS(GetOpts(0));err!=nil{
		logger.Fatal("can't get sys param CUSTOMPRICEGAS")
	}else{
		CUSTOMPRICEGAS = gas
	}
	if year,err:=BasOANN().MAXYEAR(GetOpts(0));err!=nil{
		logger.Fatal("can't get sys param MAXYEAR")
	}else{
		MAXYEAR = year
	}
	if length,err:=BasOANN().RARETYPELENGTH(GetOpts(0));err!=nil{
		logger.Fatal("can't get sys param MAXYEAR")
	}else{
		RARETYPELENGTH = length
	}
}

func IsRare(name string)  bool{
	matched, _ := regexp.MatchString(`^[0-9a-z]{0,`+ RARETYPELENGTH.String() + `}$`, name)
	return matched
}

func CheckSettings()   {
	if AROOTGAS == nil || BROOTGAS==nil ||
		SUBGAS == nil || CUSTOMPRICEGAS == nil ||
		MAXYEAR == nil || RARETYPELENGTH == nil{
			Settings()
	}
}


func GetARootGas() string {
	if AROOTGAS == nil{
		return ""
	}
	return AROOTGAS.String()
}


func GetBRootGas() string{
	if BROOTGAS == nil{
		return ""
	}
	return BROOTGAS.String()
}

func GetSubGas() string {
	if SUBGAS == nil{
		return ""
	}
	return SUBGAS.String()
}

func GetCustomPriceGas() string  {
	if CUSTOMPRICEGAS == nil{
		return ""
	}
	return CUSTOMPRICEGAS.String()
}

func GetMaxYear() int64  {
	if MAXYEAR == nil{
		return 0
	}
	return MAXYEAR.Int64()
}

func GetRareTypeLength() int64 {
	if RARETYPELENGTH == nil{
		return 0
	}

	return RARETYPELENGTH.Int64()
}

