package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := FromFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid config:", err)
		os.Exit(2)
	}

	log.Printf("Snake config -> %s", cfg)
}
