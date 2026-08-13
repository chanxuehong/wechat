package security

import (
	"crypto/subtle"
)

func SecureCompare(given, actual []byte) bool {
	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		if subtle.ConstantTimeCompare(given, actual) == 1 {
			return true
		}
		return false
	}
	// Securely compare actual to itself to keep constant time, but always return false
	if subtle.ConstantTimeCompare(actual, actual) == 1 {
		return false
	}
	return false
}

func SecureCompareString(given, actual string) bool {
	// The following code is incorrect:
	// return SecureCompare([]byte(given), []byte(actual))

	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		if subtle.ConstantTimeCompare([]byte(given), []byte(actual)) == 1 {
			return true
		}
		return false
	}
	// Securely compare actual to itself to keep constant time, but always return false
	if subtle.ConstantTimeCompare([]byte(actual), []byte(actual)) == 1 {
		return false
	}
	return false
}
