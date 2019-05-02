package main

import (
	"flag"
)

var usageMessage = "USAGE: ake -kms-id <yourKMSID> -context <encryptionContext>"

// Flags defines the type of flags passed by the user to the application
type Flags struct {
	kmsID      *string
	kmsContext *string
}

func main() {
	// Get Flags
	var flags Flags
	flags.kmsID = flag.String("kms-id", "", "the ID of your KMS key")
	flags.kmsContext = flag.String("context", "", "the KMS encryption context")
	flag.Parse()

	// Generate random password
	// Encrypt with KMS
	// Return Payload and raw password
}

func checkFlags(flags Flags) error {

	return nil
}
