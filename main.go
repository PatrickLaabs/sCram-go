package main

import (
	"fmt"
	"scram-go/cmd"
)

func main() {
	err := cmd.Command()
	if err != nil {
		fmt.Printf("Error running command %v\n", err)
	}
}
