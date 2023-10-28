package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/JoTaeYang/go-cook/module/ringbuffer"
)

type BufferT struct {
	data interface{}
	len  int
	cap  int
}

func (s BufferT) Len() int {
	return s.len
}

func (s BufferT) Cap() int {
	return s.cap
}

func (s BufferT) Data() interface{} {
	return s.data
}

var randStr = []byte("1234567890 abcdefghijklmnopqrstuvwxyz 1234567890 abcdefghijklmnopqrstuvwxyz 12345")

func main() {
	var leftCount = len(randStr)
	var frontLeftCount = 0
	ring := ringbuffer.NewRingBuffer(1000)

	var randEque, randDque int64
	var retEque int64
	var retDque int32
	//tmp := (*[]byte)(unsafe.Pointer(&randStr))
	var tmp []byte
	var tmpDeq []byte
	for {

		tmp = randStr[frontLeftCount:leftCount]

		if frontLeftCount == leftCount {
			frontLeftCount = 0
			tmp = randStr[frontLeftCount:leftCount]
		}

		tmpCnt := leftCount - frontLeftCount
		randEque = rand.Int63()%int64(tmpCnt) + 1
		retEque = int64(ring.Enqueue(&tmp, int32(randEque)))

		randDque = rand.Int63()%int64(tmpCnt) + 1
		tmpDeq = []byte{}
		retDque, _ = ring.Dequeue(&tmpDeq, int32(randDque))

		if retDque > 0 {
			fmt.Print(string(tmpDeq))
		}

		frontLeftCount += int(retEque)
		time.Sleep(time.Second)
	}
}
