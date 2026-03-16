package crypto

import (
	"crypto/rand"
	"crypto/sha256"

	"github.com/btcsuite/btcd/btcec/v2"
)

func GenerateKey() (*btcec.PrivateKey, *btcec.PublicKey) {
	priv, _ := btcec.NewPrivateKey()
	return priv, priv.PubKey()
}

func Sign(priv *btcec.PrivateKey, msg []byte) []byte {

	hash := sha256.Sum256(msg)

	sig, _ := priv.Sign(hash[:])

	return sig.Serialize()
}
