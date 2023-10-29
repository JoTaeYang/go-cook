package api

import (
	"bytes"
	"context"
	"log"
	"net"
	"reflect"
	"sync"
	"syscall"

	"github.com/JoTaeYang/go-cook/common/cache"
	"github.com/JoTaeYang/go-cook/websocket-server/api/internal"
	"github.com/gobwas/ws/wsutil"
	"golang.org/x/sys/unix"
)

type Epoll struct {
	fd          int
	connections map[int]*internal.Session
	lock        *sync.RWMutex
}

var epoller *Epoll

func MkEpoll() error {
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return err
	}

	epoller = &Epoll{
		fd:          fd,
		lock:        &sync.RWMutex{},
		connections: make(map[int]*internal.Session),
	}

	return nil
}

func InitPubSub() {
	chatChannel := []string{
		"CHANNEL_01",
	}

	subManager := cache.PubRedisClient.PSubscribe(context.Background(), chatChannel...)

	cache.SetSubScribe(subManager)
}

func PubSubGoRoutine() {
	for {
		select {
		case rMsg, ok := <-cache.Channel():
			if !ok {
				log.Println("redis subscribe channel ok!")
				continue
			}

			log.Println("channel: ", rMsg.Channel, " message: ", rMsg.Payload)
		}
	}
}

func StartEpoll() {
	for {
		connections, err := epoller.Wait()
		if err != nil {
			log.Printf("Failed to epoll wait %v", err)
			continue
		}
		for _, conn := range connections {
			var data []byte
			if conn == nil {
				break
			}
			if data, _, err = wsutil.ReadClientData(conn); err != nil {
				if err := epoller.Remove(conn); err != nil {
					log.Printf("Failed to remove %v", err)
				}
				conn.Close()
			} else {
				// This is commented out since in demo usage, stdout is showing messages sent from > 1M connections at very high rate
				//log.Printf("msg: %s", string(msg))
			}
			log.Println(string(data))

			err := cache.Publish(context.Background(), "CHANNEL_01", bytes.NewBuffer(data).String())
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (e *Epoll) Add(conn net.Conn) error {
	// Extract file descriptor associated with the connection
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_ADD, fd, &unix.EpollEvent{Events: unix.POLLIN | unix.POLLHUP | unix.POLLOUT, Fd: int32(fd)})
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	e.connections[fd] = &internal.Session{
		Conn: &conn,
	}
	if len(e.connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.connections))
	}
	return nil
}

func (e *Epoll) Remove(conn net.Conn) error {
	fd := websocketFD(conn)
	err := unix.EpollCtl(e.fd, syscall.EPOLL_CTL_DEL, fd, nil)
	if err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.connections, fd)
	if len(e.connections)%100 == 0 {
		log.Printf("Total number of connections: %v", len(e.connections))
	}
	return nil
}

func (e *Epoll) Wait() ([]net.Conn, error) {
	events := make([]unix.EpollEvent, 100)
	n, err := unix.EpollWait(e.fd, events, 100)
	if err != nil {
		return nil, err
	}
	e.lock.RLock()
	defer e.lock.RUnlock()
	var connections []net.Conn
	for i := 0; i < n; i++ {
		conn := e.connections[int(events[i].Fd)]
		connections = append(connections, *conn.Conn)
	}
	return connections, nil
}

func websocketFD(conn net.Conn) int {
	tcpConn := reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn")
	fdVal := tcpConn.FieldByName("fd")
	pfdVal := reflect.Indirect(fdVal).FieldByName("pfd")

	return int(pfdVal.FieldByName("Sysfd").Int())
}
