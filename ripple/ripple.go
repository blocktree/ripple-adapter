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
	"errors"
	"fmt"
	"path/filepath"

	"github.com/astaxie/beego/config"
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openwallet"
)

//初始化配置流程
func (wm *WalletManager) InitConfigFlow() error {

	wm.Config.InitConfig()
	file := filepath.Join(wm.Config.configFilePath, wm.Config.configFileName)
	fmt.Printf("You can run 'vim %s' to edit wallet's Config.\n", file)
	return nil
}

//查看配置信息
func (wm *WalletManager) ShowConfig() error {
	return wm.Config.PrintConfig()
}

//InstallNode 安装节点
func (wm *WalletManager) InstallNodeFlow() error {
	return errors.New("Install node is unsupport now. ")
}

//InitNodeConfig 初始化节点配置文件
func (wm *WalletManager) InitNodeConfigFlow() error {
	return errors.New("Install node is unsupport now. ")
}

//StartNodeFlow 开启节点
func (wm *WalletManager) StartNodeFlow() error {

	//return wm.startNode()
	return nil
}

//StopNodeFlow 关闭节点
func (wm *WalletManager) StopNodeFlow() error {

	//return wm.stopNode()
	return nil
}

//RestartNodeFlow 重启节点
func (wm *WalletManager) RestartNodeFlow() error {
	return errors.New("Install node is unsupport now. ")
}

//ShowNodeInfo 显示节点信息
func (wm *WalletManager) ShowNodeInfo() error {
	return errors.New("Install node is unsupport now. ")
}

//SetConfigFlow 初始化配置流程
func (wm *WalletManager) SetConfigFlow(subCmd string) error {
	file := wm.Config.configFilePath + wm.Config.configFileName
	fmt.Printf("You can run 'vim %s' to edit %s Config.\n", file, subCmd)
	return nil
}

//ShowConfigInfo 查看配置信息
func (wm *WalletManager) ShowConfigInfo(subCmd string) error {
	wm.Config.PrintConfig()
	return nil
}

//CurveType 曲线类型
func (wm *WalletManager) CurveType() uint32 {
	return wm.Config.CurveType
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "Ripple"
}

//Symbol 币种标识
func (wm *WalletManager) Symbol() string {
	return wm.Config.Symbol
}

//小数位精度
func (wm *WalletManager) Decimal() int32 {
	return 6
}

//AddressDecode 地址解析器
func (wm *WalletManager) GetAddressDecode() openwallet.AddressDecoder {
	return wm.Decoder
}

//TransactionDecoder 交易单解析器
func (wm *WalletManager) GetTransactionDecoder() openwallet.TransactionDecoder {
	return wm.TxDecoder
}

//GetBlockScanner 获取区块链
func (wm *WalletManager) GetBlockScanner() openwallet.BlockScanner {

	return wm.Blockscanner
}

//LoadAssetsConfig 加载外部配置
func (wm *WalletManager) LoadAssetsConfig(c config.Configer) error {

	wm.Config.NodeAPI = c.String("nodeAPI")
	wm.Config.WSAPI = c.String("wsAPI")
	wm.Config.APIChoose = c.String("apiChoose")
	if wm.Config.APIChoose == "rpc" {
		wm.Client = NewClient(wm.Config.NodeAPI, false)
	}else if wm.Config.APIChoose == "ws" {
		wm.WSClient = NewWSClient(wm, wm.Config.WSAPI, 0, false)
	}

	wm.Config.FixedFee, _ = c.Int64("fixedFee")
	wm.Config.ReserveAmount, _ = c.Int64("reserveAmount")
	wm.Config.IgnoreReserve, _ = c.Bool("ignoreReserve")
	wm.Config.LastLedgerSequenceNumber, _ = c.Int64("lastLedgerSequenceNumber")
	wm.Config.DataDir = c.String("dataDir")

	wm.Config.MemoType = c.String("memoType")
	wm.Config.MemoFormat = c.String("memoFormat")
	wm.Config.MemoScan = c.String("memoScan")
	//数据文件夹
	wm.Config.makeDataDir()

	return nil
}

//InitAssetsConfig 初始化默认配置
func (wm *WalletManager) InitAssetsConfig() (config.Configer, error) {
	return config.NewConfigData("ini", []byte(wm.Config.DefaultConfig))
}

//GetAssetsLogger 获取资产账户日志工具
func (wm *WalletManager) GetAssetsLogger() *log.OWLogger {
	//return nil
	return wm.Log
}

// GetSmartContractDecoder 获取智能合约解析器
func (wm *WalletManager) GetSmartContractDecoder() openwallet.SmartContractDecoder {
	return wm.ContractDecoder
}
