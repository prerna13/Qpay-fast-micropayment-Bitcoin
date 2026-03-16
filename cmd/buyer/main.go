package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"qpay/pkg/types"
)

func main() {

	tx := types.QPayTransaction{
		TxID:    fmt.Sprintf("tx-%d", time.Now().UnixNano()),
		Buyer:   "Alice",
		Seller:  "Bob",
		Amount:  0.001,
		Message: "Q_pay micropayment",
	}

	data, _ := json.Marshal(tx)

	resp, err := http.Post(
		"http://localhost:9000/request_payment",
		"application/json",
		bytes.NewBuffer(data),
	)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("Buyer received response:", resp.Status)
}
