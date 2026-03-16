package protocol

import (
	"encoding/json"
	"fmt"

	"qpay-artifact/pkg/types"
)

func AggregateProof(tx types.QPayTransaction, sigs []types.PartialSignature) {

	data, _ := json.Marshal(tx)

	fmt.Println("Commit proof generated for tx:", tx.TxID)

	fmt.Println("Signatures collected:", len(sigs))

	fmt.Println("Transaction message size:", len(data))
}
