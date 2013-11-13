package cbor

import "math"
import "math/big"

func decodeBigint(tag byte, data []byte) *big.Int {
	num := &big.Int{}
	num = num.SetBytes(data)
	if tag == tagPosBigint {
		return num
	}
	return num.Neg(num)
}

// tagDecBignum byte = typeTag | 4
// Array: [exponent(Uint/Negint, base 10), significand(Uint/Negint/Bignum)]

// tagBinBignum byte = typeTag | 5
// Array: [exponent(Uint/Negint, base  2), significand(Uint/Negint/Bignum)]

func decodeBignum(tag byte, data []byte) *big.Rat {
	// TODO
	return nil
}
