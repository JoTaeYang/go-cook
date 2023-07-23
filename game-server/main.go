package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/JoTaeYang/go-cook/module/ringbuffer"
)

//
// func (tn *TelnetLib) GetData() {
//     tn.Conn.SetReadDeadline(time.Second)
//     recvData := make([]byte, 1024)
//     n, err := tn.Conn.Read(recvData)
//     if n > 0 {
//        // do something with recvData[:n]
//     }
//     if e, ok := err.(interface{ Timeout() bool }); ok && e.Timeout() {
//         // handle timeout
//     } else if err != nil {
//        // handle error
//     }
// }

type User struct {
	Name []byte

	socket net.Conn

	RecvBuffer *ringbuffer.RingBuffer
	SendBuffer *ringbuffer.RingBuffer
}

var ChatUsers [500]*User

func (user *User) Recv() {

	user.socket.SetReadDeadline(time.Now().Add(time.Microsecond))
	n, err := user.socket.Read(user.RecvBuffer.GetRearPos())

	//연결 종료
	if err == io.EOF {
		log.Println("user connection out")
	}

	if err != nil {
		if os.IsTimeout(err) == true {
			log.Println("user timeout")
		}
	} else {
		log.Println(n)
	}
}

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
			}

			for idx, v := range ChatUsers {
				if v == nil {
					ChatUsers[idx] = user
					break
				}
			}

			log.Println("accept success")
		}
		wait.Done()
	}()

	wait.Add(1)
	go func() {
		for {
			for _, v := range ChatUsers {
				if v != nil {
					v.Recv()
				}
			}

			time.Sleep(1 * time.Second)
		}
		wait.Done()
	}()
	wait.Wait()
}
