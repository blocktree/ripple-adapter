package ripple_txsigner

import (
	"encoding/hex"
	"fmt"
	"github.com/blocktree/go-owcdrivers/rippleTransaction"
)

var Default = &TransactionSigner{}

type TransactionSigner struct {
}

// SignTransactionHash 交易哈希签名算法
// required
func (singer *TransactionSigner) SignTransactionHash(msg []byte, privateKey []byte, eccType uint32) ([]byte, error) {

	sigStr, err := rippleTransaction.SignRawTransaction(hex.EncodeToString(msg), privateKey)
	if err != nil {
		return nil, fmt.Errorf("ECC sign hash failed")
	}
	signature,_ := hex.DecodeString(sigStr)

	if signature[32] >= 0x80 {
		return nil, fmt.Errorf("Failed to serilize S in txsigner!")
	}

	return signature, nil
}
