package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func Sum(key, msg []byte) []byte {
	h := hmac.New(sha256.New, key)
	_, err := h.Write(msg)
	if err != nil {
		fmt.Printf("hmac.Write err: %v\n", err)
	}
	return h.Sum(nil)
}
