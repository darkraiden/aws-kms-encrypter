package main

import (
	"errors"
	"testing"
)

type flagAsStrings struct {
	kmsID      string
	kmsContext string
}

type testFlagOpts struct {
	input  flagAsStrings
	output error
}

var length = 12
var specialChars = true
var digits = true

func TestCheckFlagsNoErrors(t *testing.T) {
	var testFlags = []testFlagOpts{
		{
			input: flagAsStrings{
				kmsID:      "thisIsAnID",
				kmsContext: "thisIsAContext",
			},
			output: nil,
		},
		{
			input: flagAsStrings{
				kmsID:      "123thisIsAnID456",
				kmsContext: "thisIsAContext999",
			},
			output: nil,
		},
		{
			input: flagAsStrings{
				kmsID:      "$sfj(99237",
				kmsContext: "1234567890",
			},
			output: nil,
		},
	}
	for _, tf := range testFlags {
		err := checkFlags(Flag{&tf.input.kmsID, &tf.input.kmsContext, &length, &specialChars, &digits})
		if err != nil {
			t.Errorf("Error validating the input flags. Expected: %v, Got: %v", tf.output, err)
		}
	}
}

func TestCheckFlagsWithErrors(t *testing.T) {
	var testFlags = []testFlagOpts{
		{
			input: flagAsStrings{
				kmsID:      "",
				kmsContext: "thisIsAContext",
			},
			output: errors.New("Invalid flags.\nUsage: `aws-kms-encrypter -kms-id=\"ThisIsTheIDOfYourKMSKey\" -context=\"KMSEncryptionContext\""),
		},
		{
			input: flagAsStrings{
				kmsID:      "123thisIsAnID456",
				kmsContext: "",
			},
			output: errors.New("Invalid flags.\nUsage: `aws-kms-encrypter -kms-id=\"ThisIsTheIDOfYourKMSKey\" -context=\"KMSEncryptionContext\""),
		},
	}
	for _, tf := range testFlags {
		err := checkFlags(Flag{&tf.input.kmsID, &tf.input.kmsContext, &length, &specialChars, &digits})
		if err == nil {
			t.Errorf("Error validating the input flags. Expected: %v, Got: %v", tf.output, err)
		}
	}
}
