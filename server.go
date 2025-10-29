package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// ChatServer type with message history
type ChatServer struct {
	messages []string
	mu       sync.Mutex
}

// Args type for sending a message
type Args struct {
	ClientName string
	Message    string
}

// SendMessage adds a message to the history and returns full history
func (c *ChatServer) SendMessage(args Args, reply *[]string) error {
	c.mu.Lock()
	formatted := fmt.Sprintf("%s: %s", args.ClientName, args.Message)
	c.messages = append(c.messages, formatted)

	// ✅ اطبع الرسالة على السيرفر
	fmt.Println("📩", formatted)

	*reply = c.messages
	c.mu.Unlock()
	return nil
}

func main() {
	chatServer := new(ChatServer)
	rpc.Register(chatServer)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("🚀 Server started on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
