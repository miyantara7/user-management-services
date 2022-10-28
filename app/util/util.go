package util

import (
	crand "crypto/rand"
	"math/big"
)

func GenerateID() string {
	number0, err := crand.Int(crand.Reader, big.NewInt(999999999999))
	if err != nil {
		panic(err)
	}

	return PadRight(number0.String(), "0", 12)
}

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}
