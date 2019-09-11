package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/google/uuid"
)

type MessageEvent struct {
	Message string
}

func main() {

	// Create HTTP transport with the target receiver
	host := os.Getenv("EVENT_RECEIVER_HOST")
	path := os.Getenv("EVENT_RECEIVER_PATH")

	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(u.String()),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to create transport: %v", err.Error()))
	}

	// Create cloud events client
	c, err := cloudevents.NewClient(t)
	if err != nil {
		panic(fmt.Sprintf("unable to create cloudevent client: %v", err.Error()))
	}

	// Create an event
	event := createAquaEvent()

	// Send the event using the client
	log.Printf("Sending CloudEvent %s", event.Data)
	_, rsp, err := c.Send(context.Background(), event)
	if err != nil {
		log.Printf("failed to send cloudevent: %v", err.Error())
	}

	if rsp != nil {
		log.Printf("Response %v", rsp)
	}
}

func createAquaEvent() cloudevents.Event {
	m := MessageEvent{Message: "hello"}

	event := cloudevents.NewEvent()
	event.SetID(uuid.New().String())
	event.SetType("aquaEvent")
	event.SetSource("http://localhost:80/")
	event.SetDataContentType("application/json")
	event.SetData(m)
	return event
}
