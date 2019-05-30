package random_test

import (
	"testing"

	"github.com/darkraiden/aws-kms-encrypter/pkg/random"
)

var numbers = "0123456789"
var specials = "~=+%^*/()[]{}/!@#$?|"
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func TestNewRandomLength(t *testing.T) {
	lengths := [...]int{
		12,
		34,
		24,
		16,
	}

	for _, l := range lengths {
		result := random.New(l, true, true)
		if len(result) != l {
			t.Errorf("Error validating new random string. Length Expected: %d, Length Got: %d", l, len(result))
		}
	}
}
