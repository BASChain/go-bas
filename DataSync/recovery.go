package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("DataSync")

var (

	TokenAddr     = "0x105B1413461394148023FEB5bE3b4307448872d5"
	OwnershipAddr = "0x35D5FE9dfbED34e0d404A8073D8Ee9618E8dbC16"
	AssetAddr     = "0x36631a815bbecfb8947e814196DbF1768397d75b"
	DNSAddr       = "0xEc784426d352fF80E6c4192a10B009dc45e92DBD"
	OANNAddr      = "0x6a76585B037988281Aa2c80E6E42d689bA940Cef"

	conn = Bas_Ethereum.NewConn()

	token *Contract.BasToken
	ownership *Contract.BasOwnership
	asset *Contract.BasAsset
	dns *Contract.BasDNS
	oann  *Contract.BasOANN
)





func ResetConnAndService(){
	conn.Reset()
	token = nil
	ownership = nil
	asset = nil
	dns = nil
	oann = nil
}


func BasToken() *Contract.BasToken{
	if token==nil{
		if t,err:=Contract.NewBasToken(common.HexToAddress(TokenAddr), conn.GetClient());err==nil{
			token = t
		}else{
			logger.Fatal("can't recover BasToken",err)
		}
	}
	return token
}

func BasOwnership() *Contract.BasOwnership  {
	if ownership==nil{
		if o,err:=Contract.NewBasOwnership(common.HexToAddress(OwnershipAddr), conn.GetClient());err==nil{
			ownership = o
		}else{
			logger.Fatal("can't recover ownership",err)
		}
	}
	return ownership
}

func BasAsset() *Contract.BasAsset{
	if asset==nil{
		if a,err:=Contract.NewBasAsset(common.HexToAddress(AssetAddr), conn.GetClient());err==nil{
			asset = a
		}else{
			logger.Fatal("can't recover Asset",err)
		}
	}
	return asset
}

func BasDNS() *Contract.BasDNS{
	if dns==nil{
		if d,err:=Contract.NewBasDNS(common.HexToAddress(DNSAddr), conn.GetClient());err==nil{
			dns = d
		}else{
			logger.Fatal("can't recover dns",err)
		}
	}
	return dns
}


func BasOANN() *Contract.BasOANN{
	if oann==nil{
		if o,err:=Contract.NewBasOANN(common.HexToAddress(OANNAddr), conn.GetClient());err==nil{
			oann = o
		}else{
			logger.Fatal("can't recover OANN",err)
		}
	}
	return oann
}
















