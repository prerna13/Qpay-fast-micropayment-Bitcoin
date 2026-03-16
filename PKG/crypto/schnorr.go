package crypto

import (
	"crypto/rand"
	"crypto/sha256"
)

type KeyPair struct {
	Private []byte
	Public  []byte
}

func GenerateKey() KeyPair {
	priv := make([]byte, 32)
	rand.Read(priv)

	hash := sha256.Sum256(priv)

	return KeyPair{
		Private: priv,
		Public:  hash[:],
	}
}

func Sign(message []byte, priv []byte) []byte {
	h := sha256.New()
	h.Write(priv)
	h.Write(message)

	return h.Sum(nil)
}

func Verify(message []byte, sig []byte, pub []byte) bool {

	h := sha256.New()
	h.Write(pub)
	h.Write(message)

	expected := h.Sum(nil)

	return string(expected) == string(sig)
}
