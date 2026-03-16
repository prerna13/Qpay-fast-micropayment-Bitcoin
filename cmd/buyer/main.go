package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"qpay-artifact/pkg/types"
)

func main() {

	tx := types.QPayTransaction{
		TxID:   fmt.Sprintf("%d", time.Now().UnixNano()),
		Buyer:  "Alice",
		Seller: "Bob",
		Amount: 0.001,
	}

	data, _ := json.Marshal(tx)

	start := time.Now()

	http.Post("http://localhost:9000/request",
		"application/json",
		bytes.NewBuffer(data))

	latency := time.Since(start)

	fmt.Println("Buyer service latency:", latency)
}
