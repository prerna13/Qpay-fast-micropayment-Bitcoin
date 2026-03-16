package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"qpay/pkg/network"
	"qpay/pkg/types"
)

var committeeNodes = []string{
	"http://localhost:8000",
	"http://localhost:8001",
	"http://localhost:8002",
}

func requestPaymentHandler(w http.ResponseWriter, r *http.Request) {

	var tx types.QPayTransaction

	json.NewDecoder(r.Body).Decode(&tx)

	var sigs []types.PartialSignature

	for _, node := range committeeNodes {

		sig, err := network.SendApprovalRequest(node, tx)

		if err == nil {
			sigs = append(sigs, sig)
		}
	}

	req := map[string]interface{}{
		"Tx":   tx,
		"Sigs": sigs,
	}

	data, _ := json.Marshal(req)

	resp, err := http.Post(
		"http://localhost:8500/aggregate",
		"application/json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	fmt.Println("Commit proof generated")

}

func main() {

	http.HandleFunc("/request_payment", requestPaymentHandler)

	fmt.Println("Seller running on port 9000")

	http.ListenAndServe(":9000", nil)
}
