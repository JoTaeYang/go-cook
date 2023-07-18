package main

import (
	"log"

	"github.com/JoTaeYang/go-cook/module/ringbuffer"
)

func main() {
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
	ring.Dequeue(&out, 5)

	log.Println(string(out))
}
