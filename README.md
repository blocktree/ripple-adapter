# ripple-adapter

本项目适配了openwallet.AssetsAdapter接口，给应用提供了底层的区块链协议支持。

## 如何测试

openwtester包下的测试用例已经集成了openwallet钱包体系，创建conf文件，新建XRP.ini文件，编辑如下内容：

```ini

# node api url
nodeAPI = "http://ip:port"

# rpc user
rpcUser = "user"
# rpc password
rpcPassword = "password"

# fixed Fee in sawi
fixedFee = 10000
# register fee in sawi
registerFee = 10000
# min transfer amount in sawi
minTransferAmount = 10000

```

# 账户激活
```
所有地址需要20个XRP进行激活，且激活之后该20XRP将被永久冻结。
```