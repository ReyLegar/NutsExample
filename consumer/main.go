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
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		log.Printf("Received message: %s", string(m.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// Ожидание сообщений
	select {}
}
