package scram

import (
	"fmt"
	"scram-go/pkg/consts"
	"scram-go/pkg/salt"
)

// Encrypt encrypts a raw password with scram-sha-256
func Encrypt(raw []byte) (string, error) {
	if len(raw) == 0 {
		fmt.Printf("Len of raw password is %s\n", len(raw))
		return "", nil
	}

	salts, err := salt.Generate(consts.SaltSize)
	if err != nil {
		fmt.Printf("Salt generation failed: %s\n", err)
		return "", err
	}

	return generate(raw, salts, consts.IterationCount, consts.DigestLen), err
}
