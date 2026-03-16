package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"qpay/pkg/committee"
	"qpay/pkg/types"
)

var node = committee.CommitteeNode{
	ID: "node",
}

func approveHandler(w http.ResponseWriter, r *http.Request) {

	var req types.ApprovalRequest

	json.NewDecoder(r.Body).Decode(&req)

	sig := node.Approve(req.Tx)

	json.NewEncoder(w).Encode(sig)
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.HandleFunc("/approve", approveHandler)

	fmt.Println("Committee node running on port", port)

	http.ListenAndServe(":"+port, nil)
}
