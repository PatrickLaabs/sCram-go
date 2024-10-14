package sha256

import (
	"crypto/sha256"
	"fmt"
)

func Sum(key []byte) []byte {
	h := sha256.New()

	_, err := h.Write(key)
	if err != nil {
		fmt.Printf("Error writing sha256 key %v\n", err)
	}

	return h.Sum(nil)
}
