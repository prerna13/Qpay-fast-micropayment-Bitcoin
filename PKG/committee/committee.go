package committee

import (
	"qpay/pkg/crypto"
	"qpay/pkg/types"
)

type CommitteeNode struct {
	ID  string
	Key crypto.KeyPair
}

func (c *CommitteeNode) Approve(tx types.QPayTransaction) types.PartialSignature {

	msg := []byte(tx.TxID)

	sig := crypto.Sign(msg, c.Key.Private)

	return types.PartialSignature{
		NodeID: c.ID,
		Sig:    sig,
	}
}
