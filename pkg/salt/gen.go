package salt

import (
	"crypto/rand"
	"fmt"
	"io"
)

func Generate(size int) (salt []byte, err error) {
	salt = make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		fmt.Printf("failed to generate salt: %v", err)
		return nil, err
	}
	return salt, err
}
