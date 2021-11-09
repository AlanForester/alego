package rpcapi

// Transaction methods represents operations with transactions
type RawTransactionRequest struct {
	TransactionBytes string `json:"transaction_bytes"`
}

//type RawTransactionEndpoint func() alego.RPCQuery

//// IsInline says whether message is an inline message.
//func (c *RawTransactionEndpoint) RawDecode() alego.RPCObject {
//	return &alego.RPCQuery{
//		method:   "",
//		params:   nil,
//		response: nil,
//	}
//}
//
//// Check that the transaction is valid
//func (c *RawTransactionEndpoint) RawValidate() *ValidateTransactionResponse {
//	return &ValidateTransactionResponse{}
//}
//
//// Send raw transaction bytes to this node to be added into the mempool. If valid, the transaction will be stored and propagated to all peers.
//func (c *RawTransactionEndpoint) Send() *TransactionResponse {
//	return &TransactionResponse{}
//}
//
//type TransactionEndpoint struct{}
//
//// Transaction methods represents operations with transactions
//type TransactionRequest struct {
//	TransactionId string `json:"transaction_id"`
//}
//
//// Check that the transaction is valid
//func (c *TransactionEndpoint) GetInfo() *RawTransactionResponse {
//	return &RawTransactionResponse{}
//}
//

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
