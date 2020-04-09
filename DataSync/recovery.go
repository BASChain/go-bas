package DataSync

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/Contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/op/go-logging"
)

var logger, _ = logging.GetLogger("DataSync")

var (

	TokenAddr     = "0x809e5b1508fA95127Aa5E19879aAC8aFAc298915"
	OwnershipAddr = "0xaAd042bBcE3c4Bca5D217D0ef7306804ee71bF9E"
	AssetAddr     = "0x97D51d2C06E9c77D849CB8647DD5dD841b3aD0aC"
	DNSAddr       = "0x3744E474307b321a8E42b4fa0A12BA4b39815a78"
	OANNAddr      = "0xd00745e93e8E920967444aBB5bA31bD762C4BbCB"

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
















