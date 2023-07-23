package main

import (
	"log"
	"net"
)

func main() {
	sock, err := net.Dial("tcp", ":30000")
	if err != nil {
		log.Fatal(err.Error())
	}

	_ = sock
}
