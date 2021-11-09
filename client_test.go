package alego

import (
	"github.com/AdamSLevy/jsonrpc2/v14"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNodeInfo(t *testing.T) {
	var client = client()
	data, err := client.GetNodeInfo()
	if err != nil {
		client.logger.Errorf("Error on TestGetNodeInfo: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on TestGetNodeInfo: %v", err.Error())
	}

	assert.False(t, data.IsMiner)
}

func TestGetNodeStats(t *testing.T) {
	var client = client()
	_, err := client.GetNodeStats()
	if err != nil {
		client.logger.Errorf("Error on TestGetNodeStats: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on TestGetNodeStats: %v", err.Error())
	}

	assert.Empty(t, err)
}

func TestGetNodePeers(t *testing.T) {
	var client = client()
	_, err := client.GetNodeStats()
	if err != nil {
		client.logger.Errorf("Error on TestGetNodePeers: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on TestGetNodePeers: %v", err.Error())
	}

	assert.Empty(t, err)
}

func TestGetBlock(t *testing.T) {
	var client = client()
	_, err := client.GetBlock("de24fc1ddb8f55a46277aee00e5e13e283337a006491ca9c01c22a5d764ea8f8")
	if err != nil {
		client.logger.Errorf("Error on TestGetBlock: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on TestGetBlock: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestDecodeRawTransaction(t *testing.T) {
	var client = client()
	_, err := client.DecodeRawTransaction("de24fc1ddb8f55a46277aee00e5e13e283337a006491ca9c01c22a5d764ea8f8")
	if err != nil {
		client.logger.Errorf("Error on DecodeRawTransaction: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on DecodeRawTransaction: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestGetRawTransaction(t *testing.T) {
	var client = client()
	_, err := client.GetRawTransaction("de24fc1ddb8f55a46277aee00e5e13e283337a006491ca9c01c22a5d764ea8f8")
	if err != nil {
		client.logger.Errorf("Error on GetRawTransaction: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on GetRawTransaction: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestGetTransactionInfo(t *testing.T) {
	var client = client()
	_, err := client.GetTransactionInfo("de24fc1ddb8f55a46277aee00e5e13e283337a006491ca9c01c22a5d764ea8f8")
	if err != nil {
		client.logger.Errorf("Error on GetTransactionInfo: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on GetTransactionInfo: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestSendTransaction(t *testing.T) {
	var client = client()
	_, err := client.SendTransaction("de24fc1ddb8f55a46277aee00e5e13e283337a006491ca9c01c22a5d764ea8f8")
	if err != nil {
		client.logger.Errorf("Error on SendTransaction: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on SendTransaction: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestValidateRawTransaction(t *testing.T) {
	var client = client()
	_, err := client.ValidateRawTransaction("de24fc1ddb8f55a46277aee00e5e13e283337a006491ca9c01c22a5d764ea8f8")
	if err != nil {
		client.logger.Errorf("Error on ValidateRawTransaction: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on ValidateRawTransaction: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestGetBlockTemplate(t *testing.T) {
	var client = client()
	_, err := client.GetBlockTemplate()
	if err != nil {
		client.logger.Errorf("Error on GetBlockTemplate: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on GetBlockTemplate: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestGetBlockHash(t *testing.T) {
	var client = client()
	_, err := client.GetBlockHash("10000")
	if err != nil {
		client.logger.Errorf("Error on GetBlockHash: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on GetBlockHash: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestGetBlockCount(t *testing.T) {
	var client = client()
	_, err := client.GetBlockCount()
	if err != nil {
		client.logger.Errorf("Error on GetBlockCount: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on GetBlockCount: %v", err.Error())
	}

	assert.Error(t, err)
}

func TestGetBestBlockHash(t *testing.T) {
	var client = client()
	_, err := client.GetBestBlockHash()
	if err != nil {
		client.logger.Errorf("Error on GetBestBlockHash: %v", err.Error())
	}
	if _, ok := err.(jsonrpc2.Error); ok {
		client.logger.Errorf("Error on GetBestBlockHash: %v", err.Error())
	}

	assert.Error(t, err)
}
func client() *Client {
	return NewClient("alex", "842655", "52.8.172.237:443")
}
