package main

import (
	"context"
	// "encoding/json"
	// "fmt"
	"log"
	// "os"
	"sync"
	"time"

	"message-tagger/settings"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var wg = sync.WaitGroup{}

func publishMessage(ctx context.Context, body string, ch *amqp.Channel, mark string, settings settings.Settings) {
	err := ch.PublishWithContext(ctx,
		settings.Rabbit.Exhange,
		settings.Rabbit.RoutingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent message for: %v\n", mark)
	wg.Done()
}

func main() {

	settings := &settings.Settings{}
	settings.GetSettings()
	amqpAddress := settings.GetRabbitmqUrl()

	conn, err := amqp.Dial(amqpAddress)
	log.Println(amqpAddress)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	chConsume, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//Consumer part
	queue, err := ch.QueueDeclarePassive(
		settings.Rabbit.InputQueue, // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	messages, err := chConsume.Consume(
		queue.Name, // queue
		"TestConsumerName",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range messages {
		  log.Printf("Received a message: %s", d.Body)
		  wg.Add(1)
		  go publishMessage(ctx, string(d.Body), ch, string(d.Body), *settings)
		}
	  }()
	  
	  log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	  <-forever

	wg.Wait()
}
