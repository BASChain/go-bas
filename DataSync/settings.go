package DataSync

import (
	Contract "github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"regexp"
	"sync"
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

func loopOverSettingChanged(opts *bind.FilterOpts,wg *sync.WaitGroup){
	defer wg.Done()
	it,err:=BasOANN().FilterSettingChanged(opts)
	if err==nil{
		changed:=false
		for it.Next() {
			changed = true
			break
		}
		if changed {
			Settings()
		}
	}else{
		logger.Error("loop over setting changed err :" , err)
	}
}

func watchSettingChanged(opts *bind.WatchOpts,subs *[]event.Subscription,wg *sync.WaitGroup){
	logs := make(chan *Contract.BasOANNSettingChanged)
	sub,err:=BasOANN().WatchSettingChanged(opts,logs)
	defer wg.Done()
	if err==nil{
		logger.Info("watching setting changed")
		*subs = append(*subs, sub)
		for {
			select {
			case e :=<-sub.Err():
				logger.Error("subscript setting changed runtime error", e)
				return
			case log:= <-logs:
				SyncGapWithNoTrust(log.Raw.BlockNumber)
				Settings()
				logger.Info("setting changed")
			}
		}
	}else{
		logger.Error("subscript setting changed failed " ,err)
	}
}


func IsRare(name string)  bool{
	if RARETYPELENGTH == nil {
		RARETYPELENGTH = big.NewInt(6)
	}
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

