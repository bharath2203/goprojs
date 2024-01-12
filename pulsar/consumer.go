package pulsar

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"time"
)

func ListenAndServe() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		ConnectionTimeout: 10 * time.Second,
		OperationTimeout:  10 * time.Second,
		KeepAliveInterval: 10 * time.Second,
	})

	if err != nil {
		log.Fatalf("error initilizing client")
	}

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "persistent://public/default/test-topic",
		SubscriptionName: "test-topic",
		Type:             pulsar.Shared,
	})

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	w := NewWorkerPool(30)

	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			fmt.Println("error " + err.Error())
			continue
		}
		w.HandleMessage(msg.Topic())
		err = consumer.Ack(msg)
		if err != nil {
			fmt.Println("error " + err.Error())
			continue
		}
	}
}
