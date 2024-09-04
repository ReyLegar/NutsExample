package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect("nats://nuts:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subject := "updates"
	msg := "Hello, NUTS!"
	err = nc.Publish(subject, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Sent message: %s", msg)
}
