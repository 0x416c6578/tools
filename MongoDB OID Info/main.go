package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"os"
	"time"
)

const usage = `MongoDB OID Information Tool
An OID consists of:
  A 4 byte timestamp
  A 3 byte machine identifier
  A 2 byte process ID, unique to a process
  A 3 byte counter, random offset
Usage: oid <mongo object oid>
`

func main() {
	if len(os.Args) != 2 {
		fmt.Print(usage)
		os.Exit(1)
	}

	oid := os.Args[1]

	if len(oid) != 24 {
		fmt.Println("Invalid MongoDB ID")
		os.Exit(1)
	}

	timestampBytes, err := hex.DecodeString(oid[:8])
	if err != nil {
		fmt.Println("Failed to parse timestamp")
		os.Exit(1)
	}
	machineIdentifier := oid[8:14]
	processId := oid[14:18]
	counterBytes, err := hex.DecodeString(oid[18:24])
	if err != nil {
		fmt.Println("Failed to parse counter")
		os.Exit(1)
	}
	counter := uint32(counterBytes[0])<<16 | uint32(counterBytes[1])<<8 | uint32(counterBytes[2])

	timestamp := time.Unix(int64(binary.BigEndian.Uint32(timestampBytes)), 0)

	fmt.Printf("Timestamp: %s\nMachine ID: %s\nProcess ID: %s\nCounter: %d\n",
		timestamp.String(), machineIdentifier, processId, counter)
}
