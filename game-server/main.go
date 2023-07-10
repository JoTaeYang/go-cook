package main

import (
	"log"

	"github.com/JoTaeYang/go-cook/module/ringbuffer"
)

func main() {
	log.Println("Hello")

	ringbuffer.NewRingBuffer(10)
}
