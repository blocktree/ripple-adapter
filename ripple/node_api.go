/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package ripple

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/blocktree/openwallet/log"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

type ClientInterface interface {
	Call(path string, request []interface{}) (*gjson.Result, error)
}

// A Client is a Elastos RPC client. It performs RPCs over HTTP using JSON
// request and responses. A Client must be configured with a secret token
// to authenticate with other Cores on the network.
type Client struct {
	BaseURL     string
	AccessToken string
	Debug       bool
	client      *req.Req
	//Client *req.Req
}

type Response struct {
	Code    int         `json:"code,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
	Message string      `json:"message,omitempty"`
	Id      string      `json:"id,omitempty"`
}

func NewClient(url string /*token string,*/, debug bool) *Client {
	c := Client{
		BaseURL: url,
		//	AccessToken: token,
		Debug: debug,
	}

	api := req.New()
	//trans, _ := api.Client().Transport.(*http.Transport)
	//trans.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	c.client = api

	return &c
}

// Call calls a remote procedure on another node, specified by the path.
func (c *Client) Call(path string, request []interface{}) (*gjson.Result, error) {

	var (
		body = make(map[string]interface{}, 0)
	)

	if c.client == nil {
		return nil, errors.New("API url is not setup. ")
	}

	authHeader := req.Header{
		"Accept":        "application/json",
		"Authorization": "Basic " + c.AccessToken,
	}

	//json-rpc
	body["jsonrpc"] = "2.0"
	body["id"] = "curltext"
	body["method"] = path
	body["params"] = request

	if c.Debug {
		log.Std.Info("Start Request API...")
	}

	r, err := c.client.Post(c.BaseURL, req.BodyJSON(&body), authHeader)

	if c.Debug {
		log.Std.Info("Request API Completed")
	}

	if c.Debug {
		log.Std.Info("%+v", r)
	}

	if err != nil {
		return nil, err
	}

	resp := gjson.ParseBytes(r.Bytes())
	err = isError(&resp)
	if err != nil {
		return nil, err
	}

	result := resp.Get("result")

	return &result, nil
}

// See 2 (end of page 4) http://www.ietf.org/rfc/rfc2617.txt
// "To receive authorization, the client sends the userid and password,
// separated by a single colon (":") character, within a base64
// encoded string in the credentials."
// It is not meant to be urlencoded.
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))

	//return username + ":" + password
}

//isError 是否报错
func isError(result *gjson.Result) error {
	var (
		err error
	)

	/*
		//failed 返回错误
		{
			"result": null,
			"error": {
				"code": -8,
				"message": "Block height out of range"
			},
			"id": "foo"
		}
	*/

	if !result.Get("error").IsObject() {

		if !result.Get("result").Exists() {
			return errors.New("Response is empty! ")
		}

		return nil
	}

	errInfo := fmt.Sprintf("[%d]%s",
		result.Get("error.code").Int(),
		result.Get("error.message").String())
	err = errors.New(errInfo)

	return err
}

// 获取当前区块高度
func (c *Client) getBlockHeight() (uint64, error) {

	request := []interface{}{
		map[string]interface{}{
			"ledger_index": "validated",
			"accounts":     false,
			"full":         false,
			"transactions": false,
			"expand":       false,
			"owner_funds":  false,
		},
	}

	resp, err := c.Call("ledger", request)
	if err != nil {
		return 0, err
	}
	return resp.Get("ledger").Get("ledger_index").Uint(), nil
}

// 通过高度获取区块哈希
func (c *Client) getBlockHash(height uint64) (string, error) {
	request := []interface{}{
		map[string]interface{}{
			"ledger_index": height,
			"accounts":     false,
			"full":         false,
			"transactions": false,
			"expand":       false,
			"owner_funds":  false,
		},
	}
	resp, err := c.Call("ledger", request)

	if err != nil {
		return "", err
	}

	return resp.Get("ledger").Get("hash").String(), nil
}

func (c *Client) getSequence(address string) (uint32, error) {
	request := []interface{}{
		map[string]interface{}{
			"account":      address,
			"strict":       true,
			"ledger_index": "current",
			"queue":        true,
		},
	}

	r, err := c.Call("account_info", request)

	if err != nil {
		return 0, err
	}

	if r.Get("error").String() == "actNotFound" {
		return 0, nil
	}
	return uint32(r.Get("account_data").Get("Sequence").Int()), nil
}

// 获取地址余额
func (c *Client) getBalance(address string, ignoreReserve bool, reserveAmount int64) (*AddrBalance, error) {
	request := []interface{}{
		map[string]interface{}{
			"account":      address,
			"strict":       true,
			"ledger_index": "current",
			"queue":        true,
		},
	}

	r, err := c.Call("account_info", request)

	if err != nil {
		return nil, err
	}

	if r.Get("error").String() == "actNotFound" {
		return &AddrBalance{Address: address, Balance: big.NewInt(0), Actived: false}, nil
	}

	if ignoreReserve {
		return &AddrBalance{Address: address, Balance: big.NewInt(r.Get("account_data").Get("Balance").Int() - reserveAmount), Actived: true}, nil
	}

	return &AddrBalance{Address: address, Balance: big.NewInt(r.Get("account_data").Get("Balance").Int()), Actived: true}, nil
}

func (c *Client) isActived(address string) (bool, error) {
	request := []interface{}{
		map[string]interface{}{
			"account":      address,
			"strict":       true,
			"ledger_index": "current",
			"queue":        true,
		},
	}

	r, err := c.Call("account_info", request)

	if err != nil {
		return false, err
	}

	if r.Get("error").String() == "actNotFound" {
		return false, nil
	}
	return true, nil
}

// 获取区块信息
func (c *Client) getBlock(hash string) (*Block, error) {
	request := []interface{}{
		map[string]interface{}{
			"ledger_index": "validated",
			"accounts":     false,
			"full":         false,
			"transactions": true,
			"expand":       false,
			"owner_funds":  false,
		},
	}
	resp, err := c.Call("ledger", request)

	if err != nil {
		return nil, err
	}
	return NewBlock(resp), nil
}

func (c *Client) getBlockByHeight(height uint64) (*Block, error) {
	request := []interface{}{
		map[string]interface{}{
			"ledger_index": height,
			"accounts":     false,
			"full":         false,
			"transactions": true,
			"expand":       false,
			"owner_funds":  false,
		},
	}
	resp, err := c.Call("ledger", request)

	if err != nil {
		return nil, err
	}
	return NewBlock(resp), nil
}

func (c *Client) getTransaction(txid string) (*Transaction, error) {
	request := []interface{}{
		map[string]interface{}{
			"transaction": txid,
			"binary":      false,
		},
	}
	resp, err := c.Call("tx", request)
	if err != nil {
		return nil, err
	}
	return c.NewTransaction(resp), nil
}

func (c *Client) sendTransaction(rawTx string) (string, error) {
	request := []interface{}{
		map[string]interface{}{
			"tx_blob": rawTx,
		},
	}

	resp, err := c.Call("submit", request)
	fmt.Println(resp)
	if err != nil {
		return "", err
	}

	time.Sleep(time.Duration(1) * time.Second)

	if resp.Get("engine_result").String() != "tesSUCCESS" && resp.Get("engine_result").String() != "terQUEUED" {
		return "", errors.New("Submit transaction with error: " + resp.Get("engine_result_message").String())
	}

	return resp.Get("tx_json").Get("hash").String(), nil
}
