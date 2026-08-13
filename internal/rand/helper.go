package rand

import "encoding/hex"

// New returns 16 random bytes.
//
// The result is not printable, you can use encoding/hex or encoding/base64 to print it.
//
// Deprecated: Please use the Read function directly.
func New() (result [16]byte) {
	Read(result[:])
	return
}

// NewHex returns 32 hex-encoded random bytes.
//
// Deprecated: Please use the Read function directly.
func NewHex() (result []byte) {
	var buf [16]byte
	Read(buf[:])
	result = make([]byte, hex.EncodedLen(len(buf)))
	hex.Encode(result, buf[:])
	return
}
