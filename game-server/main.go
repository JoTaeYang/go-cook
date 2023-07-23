package main

import (
	"log"
	"net"
	"sync"

	"github.com/JoTaeYang/go-cook/module/ringbuffer"
)

func main() {
	var wait sync.WaitGroup
	log.Println("Hello")

	ring := ringbuffer.NewRingBuffer(1000)

	str := "hello"
	tt := make([]byte, 0, 5)

	for _, v := range []byte(str) {
		tt = append(tt, v)
	}

	ring.Enqueue(&tt, int32(len(tt)))

	log.Println(string(ring.Buffer))

	out := make([]byte, 0, 5)
	ring.Peek(&out, 5)

	log.Println(string(out))

	//listen socket create
	listen, err := net.Listen("tcp", "localhost:30000")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer listen.Close()

	//accept thread
	wait.Add(1)
	go func() {
		for {
			accept, err := listen.Accept()
			if err != nil {
				log.Println(err.Error())
			}

			_ = accept

			log.Println("accept success")
		}
		wait.Done()
	}()
	wait.Wait()
}
