package network

import (
	"bytes"
	"encoding/json"
	"net/http"
	"qpay/pkg/types"
)

func SendApprovalRequest(nodeURL string, tx types.QPayTransaction) (types.PartialSignature, error) {

	req := types.ApprovalRequest{
		Tx: tx,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return types.PartialSignature{}, err
	}

	resp, err := http.Post(nodeURL+"/approve", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return types.PartialSignature{}, err
	}

	defer resp.Body.Close()

	var sig types.PartialSignature
	err = json.NewDecoder(resp.Body).Decode(&sig)

	return sig, err
}
