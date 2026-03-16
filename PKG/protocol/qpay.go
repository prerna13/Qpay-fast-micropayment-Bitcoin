package protocol

import (
	"crypto/sha256"
	"qpay/pkg/types"
)

func AggregateSignatures(sigs []types.PartialSignature) []byte {

	h := sha256.New()

	for _, s := range sigs {
		h.Write(s.Sig)
	}

	return h.Sum(nil)
}

func GenerateCommitProof(tx *types.QPayTransaction, sigs []types.PartialSignature) {

	agg := AggregateSignatures(sigs)

	tx.CommitProof = string(agg)

	tx.QBlockFlag = true
}
