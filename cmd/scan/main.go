package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/radisvaliullin/challenge20231220/pkg/scan"
)

func main() {
	fmt.Println("scanner run...")
	// config flag
	host := flag.String("host", "0.0.0.0", "host IP address of scanned server")
	port := flag.Int("port", 3306, "port address of scanned server")
	flag.Parse()

	// scan config
	config := scan.Config{
		Host: *host,
		Port: *port,
	}
	// scan
	s := scan.New(config)
	if err := s.Run(); err != nil {
		log.Fatalf("main: scanner run error: %v", err)
	}
}
