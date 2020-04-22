package main

import (
	"fmt"
)

type MessageHandler interface {
	handleMessage(msg string)
}

type PrintMessageHandler struct {}

func (PrintMessageHandler) handleMessage(msg string) {

	if msg == "" {
		panic("Message mustn't be empty")
	}

	fmt.Printf("New message: %s\n", msg)
}

type Connection struct {
	handler MessageHandler
}

func (c *Connection) processMessages(messages []string) {

	for _, msg := range messages {
		c.handler.handleMessage(msg)
	}
}

func (c *Connection) Close() {
	fmt.Println("Connection closed")
}

func main() {

	var m MessageHandler = PrintMessageHandler{}

	messages := []string{"Hello", "these", "are", "messages", ""}

	c := Connection{handler: m}
	defer c.Close()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic")
		}
	}()

	c.processMessages(messages)
}
