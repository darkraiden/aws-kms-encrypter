package random_test

import (
	"regexp"
	"testing"

	"github.com/darkraiden/aws-kms-encrypter/pkg/random"
)

type randomPasswordInput struct {
	length       int
	specialChars bool
	digits       bool
}

type testRandomPassword struct {
	input  randomPasswordInput
	output string
}

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

func TestNewRandomContent(t *testing.T) {
	var testArgs = []testRandomPassword{
		{
			input: randomPasswordInput{
				length:       12,
				specialChars: true,
				digits:       true,
			},
			output: "[A-Za-z0-9\\~\\=\\+\\%\\^\\*\\/\\(\\)\\[\\]\\{\\}\\/\\!\\@\\#\\$\\?\\|]{12}",
		},
		{
			input: randomPasswordInput{
				length:       34,
				specialChars: true,
				digits:       false,
			},
			output: "[A-Za-z\\~\\=\\+\\%\\^\\*\\/\\(\\)\\[\\]\\{\\}\\/\\!\\@\\#\\$\\?\\|]{34}",
		},
		{
			input: randomPasswordInput{
				length:       24,
				specialChars: false,
				digits:       true,
			},
			output: "[A-Za-z0-9]{24}",
		},
		{
			input: randomPasswordInput{
				length:       16,
				specialChars: false,
				digits:       false,
			},
			output: "[A-Za-z]{16}",
		},
	}

	for _, arg := range testArgs {
		result := random.New(arg.input.length, arg.input.specialChars, arg.input.digits)
		match, _ := regexp.MatchString(arg.output, result)
		if len(result) != arg.input.length {
			t.Errorf("Error validating new random string. Length Expected: %d, Length Got: %d", arg.input.length, len(result))
		} else if !match {
			t.Errorf("Error validating new random string. Pattern Expected: %s, String Got: %s", arg.output, result)
		}
	}
}
