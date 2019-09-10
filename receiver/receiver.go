package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
)

type MessageEvent struct {
	Message string
}

// Receive received a cloud event
func Receive(event cloudevents.Event) {
	m := &MessageEvent{}
	err := event.DataAs(m)
	if err != nil {
		log.Printf("Error decoding message: %v", err)
	}
	log.Printf("Received event %v", m)
}

func main() {
	log.Println("Starting CloudEvent Client...")

	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	log.Fatal(c.StartReceiver(context.Background(), Receive))
}
