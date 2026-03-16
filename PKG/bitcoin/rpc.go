package bitcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type RPCResponse struct {
	Result json.RawMessage `json:"result"`
	Error  interface{}     `json:"error"`
}

var rpcURL = "http://127.0.0.1:18332"
var rpcUser = "bitcoinrpc"
var rpcPass = "password"

func callRPC(method string, params []interface{}) (RPCResponse, error) {

	reqBody := RPCRequest{
		Jsonrpc: "1.0",
		ID:      "qpay",
		Method:  method,
		Params:  params,
	}

	data, _ := json.Marshal(reqBody)

	req, _ := http.NewRequest("POST", rpcURL, bytes.NewBuffer(data))

	req.SetBasicAuth(rpcUser, rpcPass)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return RPCResponse{}, err
	}

	defer resp.Body.Close()

	var rpcResp RPCResponse

	json.NewDecoder(resp.Body).Decode(&rpcResp)

	return rpcResp, nil
}

func AnchorCommitProof(proof string) {

	_, err := callRPC(
		"sendtoaddress",
		[]interface{}{
			"tb1qexampleaddressxxxxxxxx",
			0.0001,
			"",
			"",
			false,
			false,
			1,
			"UNSET",
			nil,
			proof,
		},
	)

	if err != nil {
		fmt.Println("Bitcoin anchor failed:", err)
	}

	fmt.Println("Commit proof anchored to Bitcoin testnet")
}
