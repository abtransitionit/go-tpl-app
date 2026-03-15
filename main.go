package main

import (
	"log"

	// "github.com/abtransitionit/go-app-test/lsd"
	"github.com/abtransitionit/go-app-test/logx"
)

func main() {
	if err := logx.Get(); err != nil {
		log.Fatalf("runtime error: %v", err)
	}
}
