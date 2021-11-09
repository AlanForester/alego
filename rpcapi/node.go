package rpcapi

type NodeResponse *RPCResponse

type GetPeerInfoResponse struct {
	*RPCResponse
	Peers []string `json:"peers"`
}

type NodeStatsResponse struct {
	*RPCResponse
	Connections struct {
		AllAccepted  uint64 `json:"connections.all_accepted"`
		AllInitiated uint64 `json:"connections.all_initiated"`
		AllRejected  uint64 `json:"connections.all_rejected"`

		ConnectionsConnectedPeers    uint `json:"connections.connected_peers"`
		ConnectionsConnectingPeers   uint `json:"connections.connecting_peers"`
		ConnectionsDisconnectedPeers uint `json:"connections.disconnected_peers"`
		HandshakesFailuresInit       uint `json:"handshakes.failures_init"`
		HandshakesFailuresResp       uint `json:"handshakes.failures_resp"`
		HandshakesSuccessesInit      uint `json:"handshakes.successes_init"`
		HandshakesSuccessesResp      uint `json:"handshakes.successes_resp"`
		HandshakesTimeoutsInit       uint `json:"handshakes.timeouts_init"`
		HandshakesTimeoutsResp       uint `json:"handshakes.timeouts_resp"`
		InboundAllSuccesses          uint `json:"inbound.all_successes"`
		InboundAllFailures           uint `json:"inbound.all_failures"`
		InboundBlocks                uint `json:"inbound.blocks"`
		InboundGetblocks             uint `json:"inbound.getblocks"`
		InboundGetmemorypool         uint `json:"inbound.getmemorypoolå"`
		InboundGetpeers              uint `json:"inbound.getpeers"`
		InboundGetsync               uint `json:"inbound.getsync"`
		InboundMemorypool            uint `json:"inbound.memorypool"`
		InboundPeers                 uint `json:"inbound.peers"`
		InboundPings                 uint `json:"inbound.pings"`
		InboundPongs                 uint `json:"inbound.pongs"`
		InboundSyncs                 uint `json:"inbound.syncs"`
		InboundSyncblocks            uint `json:"inbound.syncblocks"`
		InboundTransactions          uint `json:"inbound.transactions"`
		InboundUnknown               uint `json:"inbound.unknown"`
		MiscBlockHeight              uint `json:"misc.block_height"`
		MiscBlocksMined              uint `json:"misc.blocks_mined"`
		MiscDuplicateBlocks          uint `json:"misc.duplicate_blocks"`
		MiscDuplicateSyncB           uint `json:"misc.duplicate_sync_b"`
		OutboundAllSuccesse          uint `json:"outbound.all_successe"`
		OutboundAllFailures          uint `json:"outbound.all_failures"`
		QueuesInbound                uint `json:"queues.inbound"`
		QueuesOutbound               uint `json:"queues.outbound"`
	}
}

type NodeInfoResponse struct {
	IsBootNode    bool   `json:"is_bootnode"`
	IsMiner       bool   `json:"is_miner"`
	IsSyncing     bool   `json:"is_syncing"`
	Launched      string `json:"launched"`
	ListeningAddr string `json:"listening_addr"`
	Version       string `json:"version"`
}
