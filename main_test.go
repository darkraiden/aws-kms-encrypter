package main

import "testing"

type testFlag struct {
	kmsID      string
	kmsContext string
}

var flags = []testFlag{
	{
		kmsID:      "safhasfk",
		kmsContext: "foo=bar",
	},
}

func testCheckFlags(t *testing.T) {
	for _, f := range flags {
		flags := Flags{
			kmsID:      &f.kmsID,
			kmsContext: &f.kmsContext,
		}
		err := checkFlags(flags)
		if err != nil {
			t.Errorf("Invalid flags! Expected: KMS ID: %s, KMS Context: %s; Got: %v", *flags.kmsID, *flags.kmsContext, err)
		}
	}
}
