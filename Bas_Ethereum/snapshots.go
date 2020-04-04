package Bas_Ethereum

import "errors"

type Snapshot struct {
	BlockNumber uint64
	TxIndex uint
	Data interface{}
}

type Gallery []Snapshot

func (g *Gallery) GetClosest(blockNumber uint64, index uint) (error, interface{}) {
	l := len(*g)
	if l==0{
		return errors.New("empty gallery"),nil
	}
	for i:=l-1;i>=0;i--{
		if (*g)[i].BlockNumber < blockNumber ||
			((*g)[i].BlockNumber == blockNumber && (*g)[i].TxIndex <= index){
			return nil,(*g)[i].Data
		}
	}
	return errors.New("not found"),nil
}

func (g *Gallery) Insert (blockNumber uint64, index uint,Data interface{}) {
	l := len(*g)
	s:=Snapshot{
		BlockNumber: blockNumber,
		Data:        Data,
		TxIndex : index,
	}
	if l==0 {
		*g = append(*g, s)
		return
	}
	found := false

	for i:=l-1;i>=0;i--{
		if (*g)[i].BlockNumber < blockNumber ||
			((*g)[i].BlockNumber == blockNumber && (*g)[i].TxIndex <= index){
			t := &Gallery{}
			*t = append(*t, (*g)[:i+1]...)
			*t = append(*t, s)
			*g = append(*t,(*g)[i+1:]...)
			found = true
			break
		}
	}
	if !found{
		*g = append([]Snapshot{s},*g...)
	}
}
