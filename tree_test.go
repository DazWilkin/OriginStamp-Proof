package main

import (
	"encoding/xml"
	"testing"
)

func TestValue(t *testing.T) {
	for _, proof := range proofs {
		t.Run(proof.Hash, func(t *testing.T) {
			tree := Tree{}
			if err := xml.Unmarshal(proof.XML, &tree); err != nil {
				t.Errorf(err.Error())
			}

			want := tree.V

			got, err := tree.Value()
			if err != nil {
				t.Errorf(err.Error())
			}

			if got != want {
				t.Errorf("got:  %s\nwant: %s\n", got, want)
			}
		})
	}
}
