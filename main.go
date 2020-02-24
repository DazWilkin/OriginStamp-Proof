package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"hash"
	"log"

	"github.com/btcsuite/btcutil/base58"

	"golang.org/x/crypto/ripemd160"
)

func main() {
	tree := Tree{}
	if err := xml.Unmarshal(proofs[0].XML, &tree); err != nil {
		log.Fatal(err)
	}

	// Pretty-print the resulting structure to confirm (visually) that it worked
	pretty, err := xml.MarshalIndent(tree, " ", "   ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(pretty))

	// Recursively descend the tree calculating SHA-256 hashes
	key, err := tree.Value()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Correct: %v", key)

	// Determine Bitcoin address from this private key
	// See https://gobittest.appspot.com/Address

	// Step #0
	privateKey := "69AF555EFCC31073FCB82977A5BEB6FD84F20FE01F663CB1585C84945193F950"
	// privateKey := "4eac8a92f8846ea801669b5834aa733e5345cc5e90875152ea6b9f38c724701e"

	// Step #1
	publicKey, err := PublicKey(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("1: %s", publicKey)

	var h hash.Hash

	// Step #2: SHA-256

	// h := sha256.New()
	// h.Write([]byte(publicKey))
	// a := h.Sum(nil)
	// log.Printf("2: %x\n", a)

	h = sha256.New()
	x, _ := hex.DecodeString(publicKey)
	h.Write(x)
	a := h.Sum(nil)
	log.Printf("2: %x\n", a)

	// Step #3: RIPEMD-160
	h = ripemd160.New()
	h.Write(a)
	b := h.Sum(nil)
	log.Printf("3: %x\n", b)

	// Step #4: Prefix w/ "00"
	c := fmt.Sprintf("00%x", b)
	log.Printf("4: %s", c)

	// Step #5: SHA-256
	h = sha256.New()
	x, _ = hex.DecodeString(c)
	h.Write(x)
	d := h.Sum(nil)
	log.Printf("5: %x\n", d)

	// Step #6: SHA-256
	h = sha256.New()
	h.Write(d)
	e := h.Sum(nil)
	log.Printf("6: %x\n", e)

	// Step #7: 1st 4 bytes
	f := e[:4]
	log.Printf("7: %x\n", f)

	// Step #8: Append #7 to #4
	g := fmt.Sprintf("%s%x", c, f)
	log.Printf("8: %s", g)

	// Step #9: Base58
	x, _ = hex.DecodeString(g)
	i := base58.Encode(x)
	log.Printf("9: %s", i)

}
func PublicKey(privateKey string) (publicKey string, err error) {
	switch privateKey {
	case "69AF555EFCC31073FCB82977A5BEB6FD84F20FE01F663CB1585C84945193F950":
		return "04D5CDB17FE99AB803F7D695CDAD7B4FC2D18D6C79FBBEFDF26115C142C5A3FBA961E9C00A7E448AD605B45021B98F2BC9509E930F51586B950B11C8B1A97A09FE", nil
	case "4eac8a92f8846ea801669b5834aa733e5345cc5e90875152ea6b9f38c724701e":
		return "0441809144428B4F78AAB75EB5572EE24162377A841A59D7024D3C627A8B433E74BFA7C7A79C89603B552159A9B62A529EBF5F96A6839CD43B6C24117E753CA883", nil
	default:
		return "", fmt.Errorf("invalid private key")
	}
}
