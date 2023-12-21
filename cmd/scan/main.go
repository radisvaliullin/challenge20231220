package main

import (
	"fmt"
	"log"

	"github.com/radisvaliullin/challenge20231220/pkg/scan"
)

func main() {
	fmt.Println("main")

	s := scan.New()
	if err := s.Run(); err != nil {
		log.Fatalf("main: run error: %v", err)
	}
}
