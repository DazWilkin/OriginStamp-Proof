package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"golang.org/x/crypto/ripemd160"
)

func Public(privateKey string) (publicKey string) {
	var e ecdsa.PrivateKey
	e.D, _ = new(big.Int).SetString(privateKey, 16)
	e.PublicKey.Curve = secp256k1.S256()
	e.PublicKey.X, e.PublicKey.Y = e.PublicKey.Curve.ScalarBaseMult(e.D.Bytes())
	//return fmt.Sprintf("%x", elliptic.Marshal(secp256k1.S256(), e.X, e.Y))
	return hex.EncodeToString(elliptic.Marshal(secp256k1.S256(), e.X, e.Y))
}
func Address(publicKey string) (address string) {
	h := sha256.New()

	// Step #1
	s1, _ := hex.DecodeString(publicKey)

	// Step #2: SHA-256
	h.Reset()
	h.Write(s1)
	s2 := h.Sum(nil)
	log.Printf("2: %x", s2)

	// Step #3: RIPEMD-160
	r := ripemd160.New()
	r.Write(s2)
	s3 := r.Sum(nil)
	log.Printf("3: %x", s3)

	// Step #4: Prefix w/ "00"
	// s4, err := hex.DecodeString(fmt.Sprintf("00%x", s3))
	s4 := append([]byte{0}, s3...)
	log.Printf("4: %x", s4)

	// Step #5: SHA-256
	h.Reset()
	h.Write(s4)
	s5 := h.Sum(nil)
	log.Printf("5: %x", s5)

	// Step #6: SHA-256
	h.Reset()
	h.Write(s5)
	s6 := h.Sum(nil)
	log.Printf("6: %x", s6)

	// Step #7: 1st 4 bytes
	s7 := s6[:4]
	log.Printf("7: %x", s7)

	// Step #8: Append #7 to #4
	s8 := append(s4, s7...)
	log.Printf("8: %x", s8)

	// Step #9: Base58
	s9 := base58.Encode(s8)
	log.Printf("9: %s", s9)

	return s9
}
