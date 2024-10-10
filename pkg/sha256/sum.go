package sha256

import "crypto/sha256"

func Sum(key []byte) []byte {
	h := sha256.New()
	_, _ = h.Write(key)
	return h.Sum(nil)
}
