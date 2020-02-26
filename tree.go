package main

import (
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"log"
)

// Tree represents the XML structure of a OriginStamp Proof
type Tree struct {
	// The XMLName value is ignored; it starts as `node` but then becomes either `left` or `right`
	// Unclear whether there's any utility capturing T(ype); it's either "key" or "mesh"
	// Left|Right must be pointers to avoid complaints about this being a recursive struct

	XMLName xml.Name `xml`
	V       string   `xml:"value,attr"`
	T       string   `xml:"type,attr"`
	Left    *Tree    `xml:"left"`
	Right   *Tree    `xml:"right"`
}

// Value recursively descends the comparing the node's value (V) with the hash of its concatenated leaves
func (t *Tree) Value() (string, error) {
	if t == nil {
		return "", fmt.Errorf("Can't compute hash of an empty tree")
	}
	if t.V == "" {
		return "", fmt.Errorf("V is a SHA-256 hash and must have a value")
	}

	log.Printf("Value: [%s]", t.V)

	if (t.Left == nil && t.Right != nil) || (t.Left != nil && t.Right == nil) {
		return "", fmt.Errorf("Tree is unbalanced: neither or both branches must be present")
	}
	// If no branches are present, the value is V
	if t.Left == nil && t.Right == nil {
		return t.V, nil
	}

	// Both branches are present
	// The value is V iff V = SHA(LR)
	var l, r string
	var err error
	if l, err = t.Left.Value(); err != nil {
		return "", err
	}
	log.Printf("Value: [%s]  L=%s", t.V, l)

	if r, err = t.Right.Value(); err != nil {
		return "", err
	}
	log.Printf("Value: [%s]  R=%s", t.V, r)

	// Compute the SHA-256 hash of LR
	b := sha256.Sum256([]byte(fmt.Sprintf("%s%s", l, r)))
	h := fmt.Sprintf("%x", b)

	log.Printf("Value: [%s] LR=%s", t.V, h)

	if h != t.V {
		return "", fmt.Errorf("Hash of branches does not match V")
	}

	return h, nil
}
