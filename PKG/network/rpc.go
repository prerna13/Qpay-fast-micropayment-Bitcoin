package network

import (
	"bytes"
	"encoding/json"
	"net/http"

	"qpay-artifact/pkg/types"
)

func RequestApproval(node string, tx types.QPayTransaction) (types.PartialSignature, error) {

	req := types.ApprovalRequest{Tx: tx}

	data, _ := json.Marshal(req)

	resp, err := http.Post(node+"/approve", "application/json", bytes.NewBuffer(data))

	if err != nil {
		return types.PartialSignature{}, err
	}

	defer resp.Body.Close()

	var sig types.PartialSignature

	json.NewDecoder(resp.Body).Decode(&sig)

	return sig, nil
}
