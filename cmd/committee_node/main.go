package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"qpay-artifact/pkg/crypto"
	"qpay-artifact/pkg/types"
)

var priv, _ = crypto.GenerateKey()

func approve(w http.ResponseWriter, r *http.Request) {

	var req types.ApprovalRequest

	json.NewDecoder(r.Body).Decode(&req)

	sig := crypto.Sign(priv, []byte(req.Tx.TxID))

	resp := types.PartialSignature{
		NodeID: "committee",
		Sig:    sig,
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/approve", approve)

	fmt.Println("Committee node running:", port)

	http.ListenAndServe(":"+port, nil)
}
