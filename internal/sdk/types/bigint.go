// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import (
	"fmt"
	"math/big"
)

// MustNewBigIntFromString returns an instance of big.Int from a string
// The string is assumed to be base 10 and if it is not a valid big.Int
// then the function panics.
// Avoid using this function in production code.
func MustNewBigIntFromString(s string) *big.Int {
	i, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic(fmt.Errorf("failed to parse string as big.Int"))
	}

	return i
}
