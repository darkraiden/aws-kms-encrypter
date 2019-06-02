package kms

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

// New creates a new Encrypter object to be used to interact with the package methods
func New(kmsID, kmsContext string) Encrypter {
	return Encrypter{
		ID:      kmsID,
		Context: kmsContext,
	}
}

// Encrypt uses the AWS API to encrypt a string using KMS and returns a slice of bytes
// (Encryption Text blob in base64), and an error
func (e *Encrypter) Encrypt(str string) ([]byte, error) {
	svc := kms.New(session.New())
	input := &kms.EncryptInput{
		KeyId:     &e.ID,
		Plaintext: []byte(str),
	}
	result, err := svc.Encrypt(input)
	if err != nil {
		return nil, err
	}
	return result.CiphertextBlob, nil
}
