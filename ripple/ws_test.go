package ripple

//
//
//import (
//	"encoding/json"
//	"flag"
//"log"
//"net/url"
//"os"
//"os/signal"
//	"testing"
//	"time"
//
//"github.com/gorilla/websocket"
//)
//
//var addr = flag.String("addr", "47.244.179.69:20029", "http service address")
//
//func Test_123(t *testing.T) {
//	flag.Parse()
//	log.SetFlags(0)
//
//	interrupt := make(chan os.Signal, 1)
//	signal.Notify(interrupt, os.Interrupt)
//
//	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
//	log.Printf("connecting to %s", u.String())
//	//request := map[string]interface{}{
//	//	"id" : 14,
//	//	"command":"ledger",
//	//	"ledger_index": "validated",
//	//	"accounts":     false,
//	//	"full":         false,
//	//	"transactions": false,
//	//	"expand":       false,
//	//	"owner_funds":  false,
//	//}
//
//	request :=        map[string]interface{}{
//		"id":           14,
//		"command":      "ledger",
//		"ledger_index": 48554232,
//		"accounts":     false,
//		"full":         false,
//		"transactions": true,
//		"expand":       false,
//		"owner_funds":  false,
//	}
//	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
//	if err != nil {
//		log.Fatal("dial:", err)
//	}
//	defer c.Close()
//
//	done := make(chan struct{})
//
//	go func() {
//		defer close(done)
//
//			tmp,_ := json.Marshal(request)
//			c.WriteMessage(2, tmp)
//			_, message, err := c.ReadMessage()
//			if err != nil {
//				log.Println("read:", err)
//				return
//			}
//			log.Printf("recv: %s", message)
//
//	}()
//
//	ticker := time.NewTicker(time.Second)
//	defer ticker.Stop()
//
//	for {
//		select {
//		case <-done:
//			return
//		case t := <-ticker.C:
//			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
//			if err != nil {
//				log.Println("write:", err)
//				return
//			}
//		case <-interrupt:
//			log.Println("interrupt")
//
//			// Cleanly close the connection by sending a close message and then
//			// waiting (with timeout) for the server to close the connection.
//			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
//			if err != nil {
//				log.Println("write close:", err)
//				return
//			}
//			select {
//			case <-done:
//			case <-time.After(time.Second):
//			}
//			return
//		}
//	}
//}