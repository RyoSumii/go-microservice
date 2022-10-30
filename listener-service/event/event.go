package event

import "github.com/rabbitmq/amqp091-go"

func declareExchange(ch *amqp091.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name
		"topic",      //type
		true,         // durable
		false,        //auto-delete
		false,        //internal
		false,        //no-wait
		nil,          //arguments
	)
}

func declareRandomQueue(ch *amqp091.Channel) (amqp091.Queue, error) {
	return ch.QueueDeclare(
		"",    //name
		false, //durable
		false, //delete
		true,  //exclusive
		false, //no-wait
		nil,   // arguments
	)
}
