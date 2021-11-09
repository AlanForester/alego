package rpcapi

type RPCResponse struct {
	JsonRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result;omitempty"`
	Error   interface{} `json:"error;omitempty"`
	ID      string      `json:"id"`
}
