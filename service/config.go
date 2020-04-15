package service

import (
	"github.com/BASChain/go-bas/Bas_Ethereum"
	"github.com/BASChain/go-bas/DataSync"
	"github.com/BASChain/go-bas/Market"
	"github.com/BASChain/go-bas/Miner"
)

//this should be called before any sync
func ChangeAccessPoint(points []string){
	Bas_Ethereum.AccessPoints = points
}

//this should be called before any sync
func ChangeContractAddresses(
	token_addr,
	ownership_addr,
	asset_addr,
	dns_addr,
	oann_addr,
	miner_addr,
	market_addr string){
	DataSync.TokenAddr     = token_addr
	DataSync.OwnershipAddr = ownership_addr
	DataSync.AssetAddr     = asset_addr
	DataSync.DNSAddr       = dns_addr
	DataSync.OANNAddr      = oann_addr
	Miner.MinerAddr = miner_addr
	Market.MarketAddr = market_addr
}