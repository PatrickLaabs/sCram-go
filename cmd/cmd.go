package cmd

import (
	"fmt"
	"os"
	"scram-go/pkg/scram"
)

func Command() (err error) {
	arg := os.Args

	if len(arg) < 2 {
		fmt.Printf("Missing password to hash.\nUsage: sCram-go <password>\n")
		os.Exit(1)
	}

	password := []byte(arg[1])
	hash, err := scram.Encrypt(password)
	if err != nil {
		fmt.Printf("error while encrypting password: %s\n", err)
	}

	fmt.Printf("Hashed Password:\n%s\n", hash)
	return err
}
