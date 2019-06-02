package main

import (
	"errors"
	"flag"

	"github.com/darkraiden/aws-kms-encrypter/pkg/random"
)

var usageMessage = "USAGE: ake -kms-id <yourKMSID> -context <encryptionContext>"

// Flag defines the type of flags passed by the user to the application
type Flag struct {
	kmsID           *string
	kmsContext      *string
	pwdLength       *int
	pwdSpecialChars *bool
	pwdDigits       *bool
}

func main() {
	// Get Flags
	var flags Flag
	flags.kmsID = flag.String("kms-id", "", "the ID of your KMS key")
	flags.kmsContext = flag.String("context", "", "the KMS encryption context")
	flags.pwdLength = flag.Int("length", 12, "the length of the random generated password")
	flags.pwdSpecialChars = flag.Bool("special-chars", true, "whether the password should contain special characters")
	flags.pwdDigits = flag.Bool("digits", true, "whether the password should contain numbers")
	flag.Parse()

	// Generate random password
	password := random.New(*flags.pwdLength, *flags.pwdSpecialChars, *flags.pwdDigits)

	// Encrypt with KMS
	// Return Payload and raw password
}

func checkFlags(flags Flag) error {
	if *flags.kmsID == "" || *flags.kmsContext == "" {
		return errors.New("Invalid flags.\nUsage: `aws-kms-encrypter -kms-id=\"ThisIsTheIDOfYourKMSKey\" -context=\"KMSEncryptionContext\" -length=12 -special-chars=false -digits=false")
	}

	return nil
}
