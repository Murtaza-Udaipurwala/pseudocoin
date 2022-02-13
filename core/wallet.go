package core

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivKey ecdsa.PrivateKey
	PubKey  []byte
}

func newKeyPair() (ecdsa.PrivateKey, []byte, error) {
	curve := elliptic.P256()
	priv, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return ecdsa.PrivateKey{}, nil, err
	}

	pub := append(priv.PublicKey.X.Bytes(), priv.PublicKey.Y.Bytes()...)
	return *priv, pub, nil
}

func NewWallet() (Wallet, error) {
	priv, pub, err := newKeyPair()
	if err != nil {
		return Wallet{}, err
	}

	wallet := Wallet{
		PrivKey: priv,
		PubKey:  pub,
	}

	return wallet, nil
}

func HashPubKey(pub []byte) ([]byte, error) {
	pubSHA256 := sha256.Sum256(pub)

	RIPEMD160 := ripemd160.New()
	_, err := RIPEMD160.Write(pubSHA256[:])
	if err != nil {
		return nil, err
	}

	pubRIPEMD160 := RIPEMD160.Sum(nil)
	return pubRIPEMD160, nil
}
