package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <mongo object id>")
		os.Exit(1)
	}

	id := os.Args[1]

	if len(id) != 24 {
		fmt.Println("Invalid MongoDB ID")
		os.Exit(1)
	}

	bytes, err := hex.DecodeString(id[:8])

	if err != nil {
		fmt.Println("Failed to parse timestamp")
		os.Exit(1)
	}

	timestamp := time.Unix(int64(binary.BigEndian.Uint32(bytes)), 0)

	fmt.Println(timestamp)
}
