package scram

import (
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
	"scram-go/pkg/base64"
	"scram-go/pkg/hmac"
	sha "scram-go/pkg/sha256"
	"scram-go/pkg/vars"
)

func generate(raw, salt []byte, iteration, keyLen int) string {
	digestKey := pbkdf2.Key(raw, salt, iteration, keyLen, sha256.New)
	clientKey := hmac.Sum(digestKey, vars.ClientRawKey)
	storedKey := sha.Sum(clientKey)
	serverKey := hmac.Sum(digestKey, vars.ServerRawKey)

	return fmt.Sprintf(
		"SCRAM-SHA-256$%d:%s$%s:%s",
		iteration,
		string(base64.Encode(salt)),
		string(base64.Encode(storedKey)),
		string(base64.Encode(serverKey)),
	)
}
