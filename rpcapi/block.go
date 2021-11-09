package rpcapi

type BlockRequest struct {
	BlockHash interface{} `json:"block_hash"`
}

type GetBlock struct {
	*RPCResponse
	Confirmations          int      `json:"confirmations"`
	DifficultyTarget       int      `json:"difficulty_target"`
	Height                 int      `json:"height"`
	MerkleRoot             int      `json:"merkle_root"`
	Nonce                  int      `json:"nonce"`
	PedersenMerkleRootHash int      `json:"pedersen_merkle_root_hash"`
	PreviousBlockHash      string   `json:"previous_block_hash"`
	Proof                  string   `json:"proof"`
	Size                   int      `json:"size"`
	Time                   int      `json:"time"`
	Transactions           []string `json:"transactions"`

	Hash string `json:"hash"`
}

type BlockTemplate struct {
	*RPCResponse
	PreviousBlockHash string   `json:"previous_block_hash"`
	BlockHeight       int      `json:"block_height"`
	Time              int      `json:"time"`
	DifficultyTarget  int      `json:"difficulty_target"`
	CoinbaseValue     int      `json:"coinbase_value"`
	Transactions      []string `json:"transactions"`
}

type BestBlockHashResponse struct {
	*RPCResponse
	Result string `json:"result"`
}

type BlockCountResponse struct {
	*RPCResponse
	Result int `json:"result"`
}

type BlockHashResponse struct {
	*RPCResponse
	Result string `json:"result"`
}
