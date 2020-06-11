package ripple

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

const (

	testNodeAPI = "https://"
)

func Test_getBlockHeight(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	r, err := c.getBlockHeight()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("height:", r)
	}

}

func Test_getBlockByHeight(t *testing.T) {

	c := NewClient(testNodeAPI, true)
	r, err := c.getBlockByHeight(55172763)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}
func Test_getBlockByHash(t *testing.T) {
	hash := "3Uvb87ukKKwVeU6BFsZ21hy9sSbSd3Rd5QZTWbNop1d3TaY9ZzceJAT54vuY8XXQmw6nDx8ZViPV3cVznAHTtiVE"

	c := NewClient(testNodeAPI, true)

	r, err := c.Call("blocks/signature/"+hash, nil)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}

func Test_getBlockHash(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	height := uint64(48142783)

	r, err := c.getBlockHash(height)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getBalance(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	address := "rUTEn2jLLv4ESmrUqQmhZfEfDN3LorhgvZ"

	r, err := c.getBalance(address, true, 20000000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	address = "rP9YxN6yjw5HJj5LeK55gtVr8RznEPLwRc"

	r, err = c.getBalance(address, true, 20000000)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

}

func Test_getTransaction(t *testing.T) {

	c := NewClient(testNodeAPI, true)
	//txid := "852C8C41FC914FCBCF2A598FCA82738475ABAB35F5C79014E9C0236E53016521"
	//txid := "27495DD615B7DA81B2198C849520DD4C2D4722B8F91E4382D779489E7D1CD8B6"
	txid := "72BCC9D59F29C9F086C494756F87CD4BB7260C418854816CFAD85D964B12EBCA"
	r, err := c.getTransaction(txid, "MemoData")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	// txid = "1b0147a6b5660215e9a37ca34fe9a6298988e45f7eefbd8c4b98993f4e762c3e" //"9KBoALfTjvZLJ6CAuJCGyzRA1aWduiNFMvbqTchfBVpF"

	// r, err = c.getTransaction(txid)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(r)
	// }

	// txid = "3ca0888b232df90d910a921d2f4004bb61a80bbbe27caee7107de282576e38a0" //"9KBoALfTjvZLJ6CAuJCGyzRA1aWduiNFMvbqTchfBVpF"

	// r, err = c.getTransaction(txid)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(r)
	// }
}

func Test_convert(t *testing.T) {

	amount := uint64(5000000001)

	amountStr := fmt.Sprintf("%d", amount)

	fmt.Println(amountStr)

	d, _ := decimal.NewFromString(amountStr)

	w, _ := decimal.NewFromString("100000000")

	d = d.Div(w)

	fmt.Println(d.String())

	d = d.Mul(w)

	fmt.Println(d.String())

	r, _ := strconv.ParseInt(d.String(), 10, 64)

	fmt.Println(r)

	fmt.Println(time.Now().UnixNano())
}

func Test_getTransactionByAddresses(t *testing.T) {
	addrs := "ARAA8AnUYa4kWwWkiZTTyztG5C6S9MFTx11"

	c := NewClient(testNodeAPI, true)
	result, err := c.getMultiAddrTransactions("MemoData", 0, -1, addrs)

	if err != nil {
		t.Error("get transactions failed!")
	} else {
		for _, tx := range result {
			fmt.Println(tx.TxID)
		}
	}
}

func Test_tmp(t *testing.T) {

	c := NewClient(testNodeAPI, true)

	block, err := c.getBlockByHeight(48059631)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(block)
}
