package main

import (
	"context"
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

	// Host:   "10.98.202.209:4000",
	// Path:   "/hello",
	host := os.Getenv("EVENT_RECEIVER_HOST")
	path := os.Getenv("EVENT_RECEIVER_PATH")

	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	m := MessageEvent{Message: "hello"}

	log.Printf("Sending CloudEvent %v", m)

	event := cloudevents.NewEvent()
	event.SetID(uuid.New().String())
	event.SetType("lizEvent")
	event.SetSource("http://localhost:80/")
	event.SetData(m)
	event.SetDataContentType("application/json")

	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(u.String()),
	)
	if err != nil {
		panic("failed to create transport, " + err.Error())
	}

	c, err := cloudevents.NewClient(t)
	if err != nil {
		panic("unable to create cloudevent client: " + err.Error())
	}

	_, rsp, err := c.Send(context.Background(), event)
	if err != nil {
		log.Printf("failed to send cloudevent: " + err.Error())
	}

	if rsp != nil {
		log.Printf("Response %v", rsp)
	}
}
