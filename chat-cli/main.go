package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
)

type User struct {
	UserName string

	RoomNo   int32
	RoomName string

	sendBuffer *bytes.Buffer

	socket net.Conn
}

var myUser *User

var otherUser []*User

func main() {
	var userInput string
	sock, err := net.Dial("tcp", ":30000")
	if err != nil {
		log.Fatal(err.Error())
	}

	_ = sock

	//var wg sync.WaitGroup
	//wg.Add(1)

	myUser := &User{
		sendBuffer: new(bytes.Buffer),
		socket:     sock,
	}
	fmt.Print("닉네임 입력하기 : ")

	fmt.Scanln(&myUser.UserName)

	{

		if myUser.RoomNo == 0 {
			fmt.Println("1. 방 생성하기")
			fmt.Println("2. 방 입장하기")

			fmt.Scanln(&userInput)

			switch userInput {
			case "1":
				fmt.Print("생성할 방 이름을 입력하세요 : ")
				fmt.Scanln(&myUser.RoomName)

				// header := protocol.ChatHeader{
				// 	Type: int32(protocol.ANSE_ROOMCREATE),
				// 	Len:  int32(len(myUser.UserName)),
				// }

				// tmpBuf := (*[]byte)(unsafe.Pointer(&header))
				// tmpName := (*[]byte)(unsafe.Pointer(&myUser.RoomName))
				// myUser.sendBuffer.Write(*tmpBuf)
				// myUser.sendBuffer.Write(*tmpName)

				log.Println(myUser.sendBuffer.String())

				myUser.socket.Write([]byte("hello"))
				//myUser.socket.Write(myUser.sendBuffer.Bytes())
			case "2":
			}

		}
	}
	//wg.Wait()
}
