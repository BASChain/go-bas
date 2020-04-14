package Bas_Ethereum

type SyncHelper struct {
	Connection *Conn
	JustStart bool
	// this method is for sync history data
	SyncGap func(from, to uint64)
	// this method is doing the watch loop
	WatchLogic func(lastBlockNumber uint64)
	// this method will be called when remote peer closed connection
	ResetConn func()
	// do once when started
	StartupLogic func()
}

func NewSyncHelper (conn *Conn,
					SyncGapFunc func(from ,to uint64),
					WatchLogicFunc func(lastBlockNumber uint64),
					ResetConnFunc func(),
					StartupLogicFunc func()) *SyncHelper{
	return &SyncHelper{
		Connection:         conn,
		JustStart:          true,
		SyncGap:            SyncGapFunc,
		WatchLogic:         WatchLogicFunc,
		ResetConn:          ResetConnFunc,
		StartupLogic:       StartupLogicFunc,
	}
}

func (sync *SyncHelper) Watch(lastBlockNumber uint64){
	// this will block until remote peer disconnect
	sync.WatchLogic(lastBlockNumber)
	sync.ResetConn()
	currentBlock := sync.Connection.GetLastBlockNumber()
	logger.Info("connection dropped by remote peer, re-watch")
	sync.Watch(currentBlock)
}

func (sync *SyncHelper) Sync(){
	if sync.JustStart {
		sync.JustStart = false
		sync.StartupLogic()
	}
	currentBlock := sync.Connection.GetLastBlockNumber()
	go sync.Watch(currentBlock)
	sync.SyncGap(0,currentBlock)
}

func (sync *SyncHelper) ReplaySync(from,to uint64){
	go sync.SyncGap(from,to)
}