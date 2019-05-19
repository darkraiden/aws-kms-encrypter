package main

import (
	"errors"
	"flag"
)

var usageMessage = "USAGE: ake -kms-id <yourKMSID> -context <encryptionContext>"

// Flag defines the type of flags passed by the user to the application
type Flag struct {
	kmsID      *string
	kmsContext *string
}

func main() {
	// Get Flags
	var flags Flag
	flags.kmsID = flag.String("kms-id", "", "the ID of your KMS key")
	flags.kmsContext = flag.String("context", "", "the KMS encryption context")
	flag.Parse()

	// Generate random password
	// Encrypt with KMS
	// Return Payload and raw password
}

func checkFlags(flags Flag) error {
	if *flags.kmsID == "" || *flags.kmsContext == "" {
		return errors.New("Invalid flags.\nUsage: `aws-kms-encrypter -kms-id=\"ThisIsTheIDOfYourKMSKey\" -context=\"KMSEncryptionContext\"")
	}

	return nil
}
