package cbor

// See http://tools.ietf.org/html/rfc7049#section-2

// high order 3 bits of first byte
const (
	maskType   byte = 0xff << 5
	typeUint   byte = 0 << 5 //  0 + additional;  500 == (0:25) 000:01010 x01f4
	typeNegint byte = 1 << 5 // -1 - additional; -500 == (1:25) 001:01010 x01f3
	typeBytes  byte = 2 << 5 // additional specifies length
	typeUtf8   byte = 3 << 5 // additional specifies length (in bytes)
	typeArray  byte = 4 << 5 // additional specifies length (number of items)
	typeMap    byte = 5 << 5 // additional specifies length (number of key-value pairs)
	typeTag    byte = 6 << 5
	typeOther  byte = 7 << 5

	// additional 5 bits:
	// < 24 => small unsigned int
	//   24 => + 1 byte
	//   25 => + 2 byte
	//   26 => + 4 byte
	//   27 => + 8 byte
	//  ... => reserved
	//   31 => indefinite length (only for typeBytes, typeUtf8, typeArray, typeMap)
	// indefinite length for typeBytes and typeUtf8: sequence of definite length chunks

	// values for typeTag:
	tagDatetime  byte = typeTag | 0 // Utf8: Standard date/time string; 2003-12-13T18:30:02[.25](+01:00|Z)
	tagEpochtime byte = typeTag | 1 // INT/FLOAT: Epoch-based date/time, delta seconds to 1970-01-01T00:00Z
	tagPosBigint byte = typeTag | 2 // Bytes: Positive bigint;  0 + ... (in network byte order)
	tagNegBigint byte = typeTag | 3 // Bytes: Negative bigint; -1 - ... (in network byte order)
	tagDecBignum byte = typeTag | 4 // Array: [exponent(Uint/Negint, base 10), significand(Uint/Negint/Bignum)]
	tagBinBignum byte = typeTag | 5 // Array: [exponent(Uint/Negint, base  2), significand(Uint/Negint/Bignum)]
	// ... | 6..20 => unassigned
	tagConvBase64Url byte = typeTag | 21 // Bytes/Any: convert to base64url
	tagConvBase64    byte = typeTag | 22 // Bytes/Any: convert to base64
	tagConvBase16    byte = typeTag | 23 // Bytes/Any: convert to base16 (upper case hex)
	tagKeepAsCbor    byte = typeTag | 24 // Bytes: CBOR not to be decoded
	// ... | 25..31 => unassigned
	tagIsUri       byte = typeTag | 32 // Utf8: contains uri [RFC3986]
	tagIsBase64Url byte = typeTag | 33 // Utf8: contains base64url encoded string [RFC4648]
	tagIsBase64    byte = typeTag | 34 // Utf8: contains base64 encoded string [RFC4648]
	tagIsRegexp    byte = typeTag | 35 // Utf8: contains regular expression in PCRE or JS syntax
	tagIsMime      byte = typeTag | 36 // Utf8: contains MIME message (with headers) [RFC2045]
	// ... | 37..55798 => unassigned
	tagSelfDescCbor1 byte = typeTag | 25 // header: CBOR follows; tag + 2 bytes(55799) => 0xd9 d9f7
	tagSelfDescCbor2 byte = 0xd9         // the full tag can be used to identify CBOR content
	tagSelfDescCbor3 byte = 0xf7         // following after the tag
	// ... | 55800.. => unassigned

	// subtypes of typeOther
	// ... |  0..19 => unassigned
	tokenFalse     byte = typeOther | 20
	tokenTrue      byte = typeOther | 21
	tokenNull      byte = typeOther | 22
	tokenUndefined byte = typeOther | 23
	tokenSimple    byte = typeOther | 24 // + 1 byte
	tokenFloat16   byte = typeOther | 25 // + 2 byte
	tokenFloat32   byte = typeOther | 26 // + 4 byte
	tokenFloat64   byte = typeOther | 27 // + 8 byte
	// ... |  28..30 => unassigned
	tokenBreak byte = typeOther | 31 // 0xff: token to end indefinite types
	// ... | 31+ (after tokenSimple) => unassigned
)

// base64url is like base64 without trailing "="
