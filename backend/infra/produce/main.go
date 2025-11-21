package produce

import amqp "github.com/rabbitmq/amqp091-go"

type Produce struct {
	EmailService *EmailService
}

var produceInstance *Produce

func InitProduce(channel *amqp.Channel) *Produce {
	if produceInstance != nil {
		return produceInstance
	}

	emailService := InitEmailService(channel)
	if emailService == nil {
		panic("Failed to initialize Email service")
	}

	produceInstance = &Produce{
		EmailService: emailService,
	}

	return produceInstance
}

func GetProduce() *Produce {
	if produceInstance == nil {
		panic("Produce not initialized. Call InitProduce() first.")
	}
	return produceInstance
}
