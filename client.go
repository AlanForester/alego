package alego

import (
	"github.com/AdamSLevy/jsonrpc2/v14"
	"github.com/AlexCollin/alego/rpcapi"
)

// Client is a low-level Telegram client
type Client struct {
	username  string
	password  string
	baseURL   string
	url       string
	rpcClient *jsonrpc2.Client
	logger    Logger
	timeout   int
}

// NewClient creates new Telegram API client
func NewClient(username string, pass string, baseURL string) *Client {
	var c = new(Client)
	c.baseURL = baseURL
	c.url = "http://" + baseURL
	c.logger = new(BasicLogger)
	c.username = username
	c.password = pass

	c.rpcClient = &jsonrpc2.Client{
		DebugRequest: false,
		User:         username,
		Password:     pass,
		Log:          BasicLogger{},
	}

	return c
}

// GetMe returns info about bot as a User object
func (c *Client) GetNodeInfo() (rpcapi.NodeInfoResponse, error) {
	me := rpcapi.NodeInfoResponse{}
	err := c.doRequest("getnodeinfo", nil, &me)
	return me, err
}

func (c *Client) GetNodeStats() (rpcapi.NodeStatsResponse, error) {
	me := rpcapi.NodeStatsResponse{}
	err := c.doRequest("getnodestats", nil, &me)
	return me, err
}

func (c *Client) GetNodePeers() (rpcapi.GetPeerInfoResponse, error) {
	me := rpcapi.GetPeerInfoResponse{}
	err := c.doRequest("getpeerinfo", nil, &me)
	return me, err
}

func (c *Client) GetBlock(hash string) (interface{}, error) {
	me := rpcapi.GetBlock{}
	err := c.doRequest("getblock", []string{hash}, &me)
	return me, err
}

func (c *Client) GetBestBlockHash() (interface{}, error) {
	me := rpcapi.BestBlockHashResponse{}
	err := c.doRequest("getbestblockhash", nil, &me)
	return me, err
}

func (c *Client) GetBlockCount() (interface{}, error) {
	me := rpcapi.BlockCountResponse{}
	err := c.doRequest("getblockcount", nil, &me)
	return me, err
}

func (c *Client) GetBlockHash(block_height string) (interface{}, error) {
	me := rpcapi.BlockHashResponse{}
	err := c.doRequest("getblockhash", []string{block_height}, &me)
	return me, err
}

func (c *Client) GetBlockTemplate() (interface{}, error) {
	me := rpcapi.BlockTemplate{}
	err := c.doRequest("getblocktemplate", nil, &me)
	return me, err
}

func (c *Client) ValidateRawTransaction(transaction_bytes string) (interface{}, error) {
	me := rpcapi.RawTransactionResponse{}
	err := c.doRequest("validaterawtransaction", []string{transaction_bytes}, &me)
	return me, err
}

func (c *Client) SendTransaction(transaction_bytes string) (interface{}, error) {
	me := rpcapi.TransactionResponse{}
	err := c.doRequest("sendtransaction", []string{transaction_bytes}, &me)
	return me, err
}

func (c *Client) GetTransactionInfo(transaction_id string) (interface{}, error) {
	me := rpcapi.RawTransactionResponse{}
	err := c.doRequest("gettransactioninfo", []string{transaction_id}, &me)
	return me, err
}

func (c *Client) GetRawTransaction(transaction_id string) (interface{}, error) {
	me := rpcapi.TransactionResponse{}
	err := c.doRequest("getrawtransaction", []string{transaction_id}, &me)
	return me, err
}

func (c *Client) DecodeRawTransaction(transaction_bytes string) (interface{}, error) {
	me := rpcapi.RawTransactionResponse{}
	err := c.doRequest("decoderawtransaction", []string{transaction_bytes}, &me)
	return me, err
}

func (c *Client) doRequest(method string, request interface{}, response interface{}) error {
	err := c.rpcClient.Request(nil, c.url, method, request, &response)
	if _, ok := err.(jsonrpc2.Error); ok {
		c.logger.Errorf("Error on RPC req: %v", err.Error())
	}
	return err
}
