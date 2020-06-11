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
	"fmt"
	"strconv"
	"time"

	"github.com/blocktree/openwallet/v2/crypto"
	"github.com/blocktree/openwallet/v2/openwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tidwall/gjson"
)

// type Vin struct {
// 	Coinbase string
// 	TxID     string
// 	Vout     uint64
// 	N        uint64
// 	Addr     string
// 	Value    string
// }

// type Vout struct {
// 	N            uint64
// 	Addr         string
// 	Value        string
// 	ScriptPubKey string
// 	Type         string
// }

type Block struct {
	Hash                  string // actually block signature in XRP chain
	PrevBlockHash         string // actually block signature in XRP chain
	TransactionMerkleRoot string
	Timestamp             uint64
	Height                uint64
	Transactions          []string
}

type Transaction struct {
	TxType         string
	TxID           string
	Fee            uint64
	TimeStamp      uint64
	From           string
	To             string
	Amount         uint64
	BlockHeight    uint64
	BlockHash      string
	Status         string
	DestinationTag string
}

func (c *Client) NewTransaction(json *gjson.Result, memoScan string) *Transaction {
	obj := &Transaction{}
	obj.TxType = gjson.Get(json.Raw, "TransactionType").String()
	obj.TxID = gjson.Get(json.Raw, "hash").String()
	fee, _ := strconv.ParseInt(gjson.Get(json.Raw, "Fee").String(), 10, 64)
	obj.Fee = uint64(fee)
	obj.TimeStamp = uint64(946612800) + gjson.Get(json.Raw, "date").Uint() //1999-12-31 12:00:00 + date
	obj.From = gjson.Get(json.Raw, "Account").String()
	obj.BlockHeight = gjson.Get(json.Raw, "inLedger").Uint()
	obj.BlockHash, _ = c.getBlockHash(obj.BlockHeight)
	amount, err := strconv.ParseInt(gjson.Get(json.Raw, "Amount").String(), 10, 64)
	if err == nil {
		obj.Amount = uint64(amount)
		obj.To = gjson.Get(json.Raw, "Destination").String()
	}
	obj.Status = gjson.Get(json.Raw, "meta").Get("TransactionResult").String()
	obj.DestinationTag = gjson.Get(json.Raw, "DestinationTag").String()
	return obj
}

func (c *WSClient) NewTransaction(json *gjson.Result, memoScan string) *Transaction {
	if gjson.Get(json.Raw, "TransactionType").String() != "Payment" {
		return &Transaction{}
	}
	for count := 0; count <= 10; count++ {
		if gjson.Get(json.Raw, "meta").Get("TransactionResult").String() == "" {
			time.Sleep(500 * time.Millisecond)
			request := map[string]interface{}{
				"transaction": gjson.Get(json.Raw, "hash").String(),
				"binary":      false,
			}
			resp, err := c.Call("tx", request)
			if err != nil {
				continue
			}

			if resp.Get("error").String() != "" {
				continue
			}

			if resp.Get("result").Get("meta").Get("TransactionResult").String() == "" {
				continue
			}

			tx := resp.Get("result")
			json = &tx
			break
		}
	}

	obj := &Transaction{}
	obj.TxType = gjson.Get(json.Raw, "TransactionType").String()
	obj.TxID = gjson.Get(json.Raw, "hash").String()
	fee, _ := strconv.ParseInt(gjson.Get(json.Raw, "Fee").String(), 10, 64)
	obj.Fee = uint64(fee)
	obj.TimeStamp = uint64(946612800) + gjson.Get(json.Raw, "date").Uint() //1999-12-31 12:00:00 + date
	obj.From = gjson.Get(json.Raw, "Account").String()
	obj.BlockHeight = gjson.Get(json.Raw, "inLedger").Uint()
	if obj.BlockHeight != 0 {
		obj.BlockHash, _ = c.getBlockHash(obj.BlockHeight)
	}
	amount, err := strconv.ParseInt(gjson.Get(json.Raw, "Amount").String(), 10, 64)
	if err == nil {
		obj.Amount = uint64(amount)
		obj.To = gjson.Get(json.Raw, "Destination").String()
	}

	//if gjson.Get(json.Raw, "meta").Get("TransactionResult").String() == "tesSUCCESS" {
	//	obj.Status = "success"
	//}
	obj.Status = gjson.Get(json.Raw, "meta").Get("TransactionResult").String()
	//if obj.Status != "tesSUCCESS" {
	//	fmt.Println("[XRP:wrong_status] tx-",json.String())
	//}

	obj.DestinationTag = gjson.Get(json.Raw, "DestinationTag").String()
	return obj
}

func NewTransaction(json *gjson.Result) *Transaction {
	obj := &Transaction{}
	obj.BlockHash = json.Get("ledger_hash").String()
	obj.BlockHeight = json.Get("ledger_index").Uint()
	obj.TxType = json.Get("tx_json").Get("TransactionType").String()
	obj.TxID = json.Get("tx_json").Get("hash").String()
	fee, _ := strconv.ParseInt(json.Get("tx_json").Get("Fee").String(), 10, 64)
	obj.Fee = uint64(fee)
	obj.From = json.Get("tx_json").Get("Account").String()
	amount, err := strconv.ParseInt(json.Get("tx_json").Get("Amount").String(), 10, 64)
	if err == nil {
		obj.Amount = uint64(amount)
		obj.To = json.Get("tx_json").Get("Destination").String()
	}
	//if json.Get( "metadata").Get("TransactionResult").String() == "tesSUCCESS" {
	//	obj.Status = "success"
	//}
	obj.Status = json.Get("metadata").Get("TransactionResult").String()
	return obj
}

func NewBlock(json *gjson.Result) *Block {
	obj := &Block{}
	// 解  析
	obj.Hash = gjson.Get(json.Raw, "ledger").Get("ledger_hash").String()
	obj.PrevBlockHash = gjson.Get(json.Raw, "ledger").Get("parent_hash").String()
	obj.TransactionMerkleRoot = gjson.Get(json.Raw, "ledger").Get("transaction_hash").String()
	obj.Timestamp = gjson.Get(json.Raw, "ledger").Get("close_time").Uint()
	obj.Height = gjson.Get(json.Raw, "ledger").Get("ledger_index").Uint()

	for _, tx := range gjson.Get(json.Raw, "ledger").Get("transactions").Array() {
		obj.Transactions = append(obj.Transactions, tx.String())
	}

	if obj.Hash == "" {
		time.Sleep(5 * time.Second)
	}
	return obj
}

//BlockHeader 区块链头
func (b *Block) BlockHeader() *openwallet.BlockHeader {

	obj := openwallet.BlockHeader{}
	//解析json
	obj.Hash = b.Hash
	//obj.Confirmations = b.Confirmations
	obj.Merkleroot = b.TransactionMerkleRoot
	obj.Previousblockhash = b.PrevBlockHash
	obj.Height = b.Height
	//obj.Version = uint64(b.Version)
	obj.Time = b.Timestamp
	obj.Symbol = Symbol

	return &obj
}

//UnscanRecords 扫描失败的区块及交易
type UnscanRecord struct {
	ID          string `storm:"id"` // primary key
	BlockHeight uint64
	TxID        string
	Reason      string
}

func NewUnscanRecord(height uint64, txID, reason string) *UnscanRecord {
	obj := UnscanRecord{}
	obj.BlockHeight = height
	obj.TxID = txID
	obj.Reason = reason
	obj.ID = common.Bytes2Hex(crypto.SHA256([]byte(fmt.Sprintf("%d_%s", height, txID))))
	return &obj
}
