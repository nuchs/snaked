package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	cfg, err := FromFlags()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid config:", err)
		os.Exit(1)
	}

	f, err := initLogging(cfg)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to set up logging")
		os.Exit(2)
	}
	defer f.Close()

	if err := Run(cfg); err != nil {
		log.Fatalf("Pants: %s", err)
	}

	log.Println("Ok lady, I love you buhbye!")
}

func initLogging(cfg Config) (*os.File, error) {
	if !cfg.Debug {
		log.SetOutput((io.Discard))
		return nil, nil
	}

	f, err := os.OpenFile(cfg.LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Printf("======================================\n")
	log.Printf("Hello, it's a me, snaked\n")
	log.Println(cfg.String())

	return f, nil
}
