package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	amqpURL      = "amqp://guest:guest@localhost:5672/"
	exchangeName = "demo.events"
	exchangeType = "topic"
	queueName    = "demo.video-worker"
	bindingKey   = "video.*"
	routingKey   = "video.created"
)

type Event struct {
	VideoID    uint      `json:"video_id"`
	Action     string    `json:"action"`
	OccurredAt time.Time `json:"occurred_at"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: go run . consume|publish")
		os.Exit(2)
	}

	var err error

	switch os.Args[1] {
	case "consume":
		err = consume()
	case "publish":
		err = publish()
	default:
		fmt.Fprintln(os.Stderr, "usage: go run . consume|publish")
		os.Exit(2)
	}

	if err != nil {
		log.Fatal(err)
	}
}

func openChannel() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		_ = conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}

func declareTopology(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,  // durable
		false, // auto-delete
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	return ch.QueueBind(
		q.Name,
		bindingKey,
		exchangeName,
		false, // no-wait
		nil,   // arguments
	)
}

func publish() error {
	conn, ch, err := openChannel()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	if err := declareTopology(ch); err != nil {
		return err
	}

	event := Event{
		VideoID:    42,
		Action:     "created",
		OccurredAt: time.Now(),
	}

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ch.PublishWithContext(
		ctx,
		exchangeName,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Body:         body,
		},
	)
}

func consume() error {
	conn, ch, err := openChannel()
	if err != nil {
		return err
	}
	defer conn.Close()
	defer ch.Close()

	if err := declareTopology(ch); err != nil {
		return err
	}

	deliveries, err := ch.Consume(
		queueName,
		"",    // consumer tag
		false, // auto-ack; false means acknowledge manually
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	log.Printf("listening on queue %q", queueName)

	for delivery := range deliveries {
		fmt.Printf(
			"received: exchange=%s routing_key=%s body=%s\n",
			delivery.Exchange,
			delivery.RoutingKey,
			delivery.Body,
		)

		if err := delivery.Ack(false); err != nil {
			return err
		}
	}

	return nil
}
