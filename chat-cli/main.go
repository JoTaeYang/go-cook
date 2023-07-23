package main

import (
	"log"
	"net"
	"sync"
	"time"
)

func main() {
	sock, err := net.Dial("tcp", ":30000")
	if err != nil {
		log.Fatal(err.Error())
	}

	_ = sock

	var wg sync.WaitGroup
	wg.Add(1)
	for {
		defer wg.Done()
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}
