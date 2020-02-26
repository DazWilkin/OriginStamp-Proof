package main

import (
	"encoding/xml"
	"log"
	"strings"
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
	// The root (if correct) is the private key
	privateKey, err := tree.Value()
	if err != nil {
		log.Fatal(err)
	}

	// Step #0 of 9
	// See: https://gobittest.appspot.com/Address
	log.Printf("Correct: %v", privateKey)

	// Step #15
	publicKey := Public(privateKey)
	log.Printf("1: %s", strings.ToUpper(publicKey))

	// Step #9
	address := Address(publicKey)
	log.Printf("9: %s", address)

}
