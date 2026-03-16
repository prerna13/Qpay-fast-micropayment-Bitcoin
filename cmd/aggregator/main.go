package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"qpay/pkg/protocol"
	"qpay/pkg/types"
)

type AggregationRequest struct {
	Tx   types.QPayTransaction
	Sigs []types.PartialSignature
}

func aggregateHandler(w http.ResponseWriter, r *http.Request) {

	var req AggregationRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	protocol.GenerateCommitProof(&req.Tx, req.Sigs)

	json.NewEncoder(w).Encode(req.Tx)
}

func main() {

	http.HandleFunc("/aggregate", aggregateHandler)

	fmt.Println("Aggregator running on port 8500")

	http.ListenAndServe(":8500", nil)
}
