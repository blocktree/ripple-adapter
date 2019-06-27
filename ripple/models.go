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
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/blocktree/openwallet/crypto"
	"github.com/blocktree/openwallet/openwallet"
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
	TxType      string
	TxID        string
	Fee         uint64
	TimeStamp   uint64
	From        string
	To          string
	Amount      uint64
	BlockHeight uint64
	BlockHash   string
	Status      string
	MemoData    string
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
	obj.Status = gjson.Get(json.Raw, "status").String()
	memos := gjson.Get(json.Raw, "Memos").Array()
	if memos != nil && len(memos) >= 1 {
		memoData := memos[0].Get("Memo").Get(memoScan).String()
		memo, _ := hex.DecodeString(memoData)
		obj.MemoData = string(memo)
	}
	return obj
}

func NewBlock(json *gjson.Result) *Block {
	obj := &Block{}
	// 解 析
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
