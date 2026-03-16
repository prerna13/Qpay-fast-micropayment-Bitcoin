package protocol

import (
	"encoding/json"
	"fmt"

	"qpay-artifact/pkg/bitcoin"
	"qpay-artifact/pkg/types"
)

func AggregateProof(tx types.QPayTransaction, sigs []types.PartialSignature) {

	data, _ := json.Marshal(tx)

	proof := fmt.Sprintf("%x", data)

	fmt.Println("Commit proof generated:", proof)

	bitcoin.AnchorCommitProof(proof)
}
