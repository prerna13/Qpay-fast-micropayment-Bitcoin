package types

type QPayTransaction struct {
	TxID   string
	Buyer  string
	Seller string
	Amount float64
}

type PartialSignature struct {
	NodeID string
	Sig    []byte
}

type ApprovalRequest struct {
	Tx QPayTransaction
}
