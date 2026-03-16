package types

type QPayTransaction struct {
	TxID        string
	Buyer       string
	Seller      string
	Amount      float64
	Message     string
	CommitProof string
	QBlockFlag  bool
}

type ApprovalRequest struct {
	Tx QPayTransaction
}

type PartialSignature struct {
	NodeID string
	Sig    []byte
}
