package ripple

import (
	"fmt"
	"testing"
)

//var origin = "http://"
var wsurl = ":"

//
//func Test_ws(t *testing.T){
//
//	ws, err := websocket.Dial(wsurl, "", origin)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	defer func() {
//		err = ws.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Printf("WebSocket close\n")
//	}()
//	request := []interface{}{
//		map[string]interface{}{
//			"id" : 14,
//			"command":"ledger",
//			"ledger_index": "validated",
//			"accounts":     false,
//			"full":         false,
//			"transactions": false,
//			"expand":       false,
//			"owner_funds":  false,
//		},
//	}
//
//
//	//message := []byte("{\"id\": 14,\"command\": \"ledger\",\"ledger_index\": \"validated\",\"full\": false,\"accounts\": false,\"transactions\": false,\"expand\": false,\"owner_funds\": false}")
//message,_ := json.Marshal(request[0])
//	fmt.Println(string(message))
//	n, err := ws.Write(message)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Send(%d): %s\n", n, message)
//
//	var msg = make([]byte, 1024)
//	n, err = ws.Read(msg)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Receive(%d): %s\n", n, msg)
//}


func Test_ws_getBlockHeight(t *testing.T){

	height, err := tw.WSClient.getBlockHeight()
	if err != nil {
		t.Error(err)
	}

	fmt.Println("height : ", height)
}

func Test_ws_gwtBlockHash(t *testing.T) {
	//c := NewWSClient(wsurl, 0, true)
	height := uint64(48551264)
	hash, err := tw.WSClient.getBlockHash(height)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("hash : ", hash)
}

func Test_ws_getSequence(t *testing.T){
	c := tw.WSClient
	addr := "rMzax7NdBeVe5dqwo87VQepccSh9AWyP1m"

	sequence, err := c.getSequence(addr)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("sequence : ", sequence)

	addr = "rUTEn2jLLv4ESmrUqQmhZfEfDN3LorhgvZ"

	sequence, err = c.getSequence(addr)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("sequence : ", sequence)
}

func Test_ws_getBalance(t *testing.T){
	c := tw.WSClient
	addr := "rMzax7NdBeVe5dqwo87VQepccSh9AWyP1m"

	balance, err := c.getBalance(addr, true, 20000000)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println("balance : ", balance)
	}


	addr = "rUTEn2jLLv4ESmrUqQmhZfEfDN3LorhgvZ"
	balance, err = c.getBalance(addr, true, 20000000)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println("balance : ", balance)
	}

	addr = "rUTEn2jLLv4ESmrUqQmhZfEfDN3LorhgvZ"
	balance, err = c.getBalance(addr, true, 20000000)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println("balance : ", balance)
	}

	addr = "rUTEn2jLLv4ESmrUqQmhZfEfDN3LorhgvZ"
	balance, err = c.getBalance(addr, true, 20000000)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println("balance : ", balance)
	}

	addr = "rUTEn2jLLv4ESmrUqQmhZfEfDN3LorhgvZ"
	balance, err = c.getBalance(addr, true, 20000000)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println("balance : ", balance)
	}
}


func Test_ws_isActived(t *testing.T){
	c := tw.WSClient
	addr := "rMzax7NdBeVe5dqwo87VQepccSh9AWyP1m"

	isActived, err := c.isActived(addr)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println("isActived : ", isActived)
	}


	addr = "rUTEn2jLLv4ESmrUqQmhZfEfDN3LorhgvZ"
	isActived, err = c.isActived(addr)

	if err != nil {
		t.Error(err)
	}else{
		fmt.Println("isActived : ", isActived)
	}
}

func Test_ws_getBlockByHeight(t *testing.T) {
	c := tw.WSClient
	r, err := c.getBlockByHeight(48554232)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
}



func Test_ws_getTransaction(t *testing.T) {

	c := tw.WSClient
	txid := "7B5CE804B39DAD4F1EAF0BC147923B68E12E91D2C9A8F3F0E370848B6E7675E9"
	r, err := c.getTransaction(txid, "MemoData")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r, err = c.getTransaction(txid, "MemoData")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r, err = c.getTransaction(txid, "MemoData")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	r, err = c.getTransaction(txid, "MemoData")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}
	for i := 0 ;i<100 ; i ++ {
		r, err = c.getTransaction(txid, "MemoData")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(r)
		}
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



