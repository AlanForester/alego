package rpcapi

// Transaction methods represents operations with transactions
type RawTransactionRequest struct {
	TransactionBytes string `json:"transaction_bytes"`
}

type TransactionResponse struct {
	Result string `json:"result"`
}

type ValidateTransactionResponse struct {
	Result bool `json:"result"`
}

type RawTransactionResponse struct {
	TxID                string      `json:"txid"`
	Size                int         `json:"size"`
	OldSerialNumbers    []string    `json:"old_serial_numbers"`
	NewCommitments      []string    `json:"new_commitments"`
	Memo                string      `json:"memo"`
	NetworkId           int         `json:"network_id"`
	Digest              string      `json:"digest"`
	TransactionProof    string      `json:"transaction_proof"`
	ProgramCommitment   string      `json:"program_commitment"`
	LocalDataRoot       string      `json:"local_data_root"`
	ValueBalance        int64       `json:"value_balance"`
	Signatures          []string    `json:"signatures"`
	EncryptedRecords    []string    `json:"encrypted_records"`
	TransactionMetadata interface{} `json:"transaction_metadata"`
}
