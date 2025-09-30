package main

import (
	"log"

	"github.com/Molert511/TextScan/internal/server"
)

func main() {
	s := server.NewServer()
	if err := s.Run(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
