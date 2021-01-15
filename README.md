# ripple-adapter

本项目适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建XRP.ini文件，编辑如下内容：

```ini
# ws - wsAPI  rpc - nodeAPI
apiChoose = "ws"
# node api url
nodeAPI = "http://"
# websocket api url
wsAPI = "ws://"
# fixed Fee in smallest unit
fixedFee = 12
# reserve amount in smallest unit
reserveAmount = 20000000
# ignore reserve amount
ignoreReserve = true
# register fee in sawi
registerFee = 10000
# last ledger sequence number
lastLedgerSequenceNumber = 20
# memo type
memoType = "withdraw"
# memo format
memoFormat = "text/plain"
# which feild of memo to scan
memoScan = "MemoData"
# Cache data file directory, default = "", current directory: ./data
dataDir = "/home/golang/data"
```

# 账户激活
```
所有地址需要20个XRP进行激活，且激活之后该20XRP将被永久冻结。
```