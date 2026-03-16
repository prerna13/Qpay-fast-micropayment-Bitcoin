package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"qpay-artifact/pkg/network"
	"qpay-artifact/pkg/protocol"
	"qpay-artifact/pkg/types"
)

var committee = []string{
	"http://localhost:8000",
	"http://localhost:8001",
	"http://localhost:8002",
}

func requestHandler(w http.ResponseWriter, r *http.Request) {

	var tx types.QPayTransaction

	json.NewDecoder(r.Body).Decode(&tx)

	var sigs []types.PartialSignature

	for _, node := range committee {

		sig, err := network.RequestApproval(node, tx)

		if err == nil {
			sigs = append(sigs, sig)
		}
	}

	protocol.AggregateProof(tx, sigs)

	w.Write([]byte("service delivered"))
}

func main() {

	http.HandleFunc("/request", requestHandler)

	fmt.Println("Seller running 9000")

	http.ListenAndServe(":9000", nil)
}
