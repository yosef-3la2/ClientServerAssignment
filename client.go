package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	clientName, _ := reader.ReadString('\n')
	clientName = strings.TrimSpace(clientName)

	fmt.Println("\nConnected to chat server as", clientName)
	fmt.Println("Type messages and press Enter. Type 'exit' to quit.\n")

	for {
		fmt.Print(clientName, ": ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		var reply []string
		err = client.Call("ChatServer.SendMessage", struct {
			ClientName string
			Message    string
		}{ClientName: clientName, Message: text}, &reply)

		if err != nil {
			fmt.Println("Error calling RPC:", err)
			break
		}

		fmt.Println("\n💬 Chat history:")
		for _, msg := range reply {
			fmt.Println(" -", msg)
		}
		fmt.Println()
	}
}
