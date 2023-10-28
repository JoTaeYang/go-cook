package main

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"unsafe"

	"github.com/JoTaeYang/go-cook/module/packet"
	"github.com/JoTaeYang/go-cook/module/ringbuffer"
)

type User struct {
	socket net.Conn

	RecvBuffer *ringbuffer.RingBuffer
	SendBuffer *ringbuffer.RingBuffer

	GameBuffer *ringbuffer.RingBuffer
}

var Users [500]*User
var HeaderSize int32 = 5

func (user *User) CheckRecv(size int32) {
	var tmpSize int32 = size
	var header packet.Header
	var recvSize int32
	for tmpSize > 0 {
		header = packet.Header{}

		user.RecvBuffer.Peek((*[]byte)(unsafe.Pointer(&header)), HeaderSize)

		if header.ByCode[0] != 0x89 {
			log.Println("header code : %x", header.ByCode[0])
			break
		}

		packetSize := binary.BigEndian.Uint16(header.Size)

		buffer := make([]byte, 0, HeaderSize+int32(packetSize))
		recvSize, _ = user.RecvBuffer.Dequeue(&buffer, int32(packetSize))

		tmpSize -= recvSize
	}

}

func (user *User) Recv() {
	//user.socket.SetReadDeadline(time.Now().Add(time.Microsecond))
	for {
		n, err := user.socket.Read(user.RecvBuffer.GetRearPos())
		//연결 종료
		if err == io.EOF {
			log.Println("user connection out")
			break
		}

		if err != nil {
			if os.IsTimeout(err) == true {
				log.Println("user timeout")
			}
		}

		user.RecvBuffer.MoveRearPos(int32(n))

		user.CheckRecv(int32(n))
	}
}

func (user *User) Send() {

}

func main() {
	var wait sync.WaitGroup

	//listen socket create
	listen, err := net.Listen("tcp", "localhost:15000")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer listen.Close()

	//accept thread
	wait.Add(1)
	go func() {
		var user *User
		for {
			accept, err := listen.Accept()
			if err != nil {
				log.Println(err.Error())
			}

			//create user
			user = &User{
				socket: accept,

				RecvBuffer: ringbuffer.NewRingBuffer(1000),
				SendBuffer: ringbuffer.NewRingBuffer(1000),
				GameBuffer: ringbuffer.NewRingBuffer(1000),
			}

			for idx, v := range Users {
				if v == nil {
					user = Users[idx]
					break
				}
			}

			user.socket = accept
			user.RecvBuffer = ringbuffer.NewRingBuffer(1000)
			user.SendBuffer = ringbuffer.NewRingBuffer(1000)

			go user.Recv()

			log.Println("accept success")
		}
		wait.Done()
	}()

}
