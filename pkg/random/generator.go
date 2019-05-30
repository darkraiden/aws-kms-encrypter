package random

import (
	"math/rand"
)

var numbers = "0123456789"
var specials = "~=+%^*/()[]{}/!@#$?|"
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// New returns a new random string with a number of characters defined by
// the function parameter `length` and, eventually, some special chars
// and digits, depending on whether `specialChars` and/or `digits` are `true`
func New(length int, specialChars, digits bool) string {
	all := alphabet
	var buf = make([]byte, length)
	if specialChars && digits {
		all += specials + numbers
		buf[0] = specials[rand.Intn(len(specials))]
		buf[1] = numbers[rand.Intn(len(numbers))]
		for i := 2; i < length; i++ {
			buf[i] = all[rand.Intn(len(all))]
		}
	} else if specialChars && !digits {
		all += specials
		buf[0] = specials[rand.Intn(len(specials))]
		for i := 1; i < length; i++ {
			buf[i] = all[rand.Intn(len(all))]
		}
	} else if !specialChars && digits {
		all += numbers
		buf[0] = numbers[rand.Intn(len(numbers))]
		for i := 1; i < length; i++ {
			buf[i] = all[rand.Intn(len(all))]
		}
	} else {
		for i := 0; i < length; i++ {
			buf[i] = all[rand.Intn(len(all))]
		}
	}

	rand.Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return string(buf)
}
