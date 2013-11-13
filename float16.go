package cbor

import "math"

/*
### in python:
import struct
from math import ldexp

def decode_half(half):
   valu = (half & 0x7fff) << 13 | (half & 0x8000) << 16
   if ((half & 0x7c00) != 0x7c00):
       return ldexp(decode_single(valu), 112)
   return struct.unpack("!f", struct.pack("!I", valu | 0x7f800000))[0]

### in C:
#include <math.h>

double decode_half(unsigned char *halfp) {
	int half = (halfp[0] << 8) + halfp[1];
	int exp = (half >> 10) & 0x1f;
	int mant = half & 0x3ff;
	double val;
	if (exp == 0) val = ldexp(mant, -24);
	else if (exp != 31) val = ldexp(mant + 1024, exp - 25);
	else val = mant == 0 ? INFINITY : NAN;
	return half & 0x8000 ? -val : val;
}
*/

func decodeFloat16(val []byte) float64 {
	const (
		//binary layout of IEEE 754 floating point
		// [1 bit:sign] [e bits:exponent] [m bits:significand]

		bitsExp16         = 5
		bitsExp64         = 10
		bitsMant16        = 10
		bitsMant64        = 52
		bias16            = (1 << (bitsExp16 - 1)) - 1 //   15
		bias64            = (1 << (bitsExp64 - 1)) - 1 // 1023
		sign16     uint16 = 1 << (bitsExp16 + bitsMant16)
		mant16     uint16 = (1 << bitsMant16) - 1
		exp16      uint16 = ^(sign16 | mant16)
	)
	f16 := uint16(val[0])<<8 | uint16(val[1])
	exp := f16 & bitsExp16
	if exp != 0 && exp != exp16 {
		f64 := uint64(f16&sign16) << (64 - 16)
		f16 &= exp16 | mant16
		f64 |= (uint64(f16>>bitsMant16) + (bias64 - bias16)) << bitsMant64
		f64 |= uint64(f16&significand) << (bitsMant64 - bitsMant16)
		return math.Float64frombits(f64)
	}
	switch {
	case f16 == 0:
		return 0
	case f16 == sign16:
		return -0
	case exp == 0:
		// subnormal
		result := float64(f16&mant16) / (1 << bias16)
		if f16&sign16 == 0 {
			return result
		}
		return -result
	}
	// +Inf, -Inf, NaN
	return math.Float64frombits(
		(uint64(half&sign16) << (64 - 16)) |
			((uint64(^0) >> 1) << bitsMant64) |
			(uint64(half&mant16) << (bitsMant64 - bitsMant16)))
}
