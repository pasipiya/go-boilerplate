package rabbitmq

import (
	"fmt"
	"log"

	"go-boilerplate/config"
	"go-boilerplate/pkg/logger"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
	Config     config.RabbitMQConfig
}

func NewRabbitMQ(cfg config.RabbitMQConfig) (*RabbitMQ, error) {
	uri := fmt.Sprintf("%s://%s:%s@%s:%d/",
		cfg.Protocol,
		cfg.Username,
		cfg.Password,
		cfg.BrokerHost,
		cfg.BrokerPort,
	)

	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to create RabbitMQ channel: %w", err)
	}

	// QoS settings
	if err := ch.Qos(cfg.QoSLevel, 0, false); err != nil {
		return nil, fmt.Errorf("failed to set QoS: %w", err)
	}

	log.Println("🐇 RabbitMQ connected")

	return &RabbitMQ{
		Connection: conn,
		Channel:    ch,
		Config:     cfg,
	}, nil
}

func (r *RabbitMQ) Publish(queue string, body []byte) error {
	_, err := r.Channel.QueueDeclare(
		queue,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	return r.Channel.Publish(
		"",
		queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *RabbitMQ) Close() {
	if err := r.Channel.Close(); err != nil {
		logger.Warn("failed to close channel: " + err.Error())
	}

	if err := r.Connection.Close(); err != nil {
		logger.Warn("failed to close connection: " + err.Error())
	}

}
