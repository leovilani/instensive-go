package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	// how messages I want to read before confirmation that was red.
	ch.Qos(100, 0, false)
	if err != nil {
		panic(err)
	}
	return ch, nil
}

// the RabbitMQ channel is diferent than the GO channel, in this case we will use GO channel to consume the RMQ channel with goroutines
func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"orders",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		out <- msg
	}
	return nil
}
