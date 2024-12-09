package main

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//obrir canal de conexio
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//declarar la cues d'avis

	_, err = ch.QueueDeclare(
		"Avisos_FumadorMistos", //name
		false,                  //durable
		false,                  //delete when unused
		false,                  //exclusive
		false,                  //no-wait
		nil,                    //arguments
	)
	failOnError(err, "Failed to declare a queue")

	_, err = ch.QueueDeclare(
		"Avisos_FumadorTabac", //name
		false,                 //durable
		false,                 //delete when unused
		false,                 //exclusive
		false,                 //no-wait
		nil,                   //arguments
	)

	failOnError(err, "Failed to declare a queue")

	_, err = ch.QueueDeclare("Avisos_estanquer", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")
	//exchange declare
	err = ch.ExchangeDeclare("avisPolicia", //name
		"fanout", //exchange type
		false,    //durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//vincular cada una de les cues a les que enviarem l'avis de la policia

	err = ch.QueueBind("Avisos_FumadorMistos", "", "avisPolicia", false, nil)
	failOnError(err, "Failed to bind queue Avisos_FumadorMistos to avisPolicia")
	//fumadorTabac
	err = ch.QueueBind("Avisos_FumadorTabac", "", "avisPolicia", false, nil)
	failOnError(err, "Failed to bind queue Avisos_FumadorTabac to avisPolicia")
	//estanquer
	err = ch.QueueBind("Avisos_estanquer", "", "avisPolicia", false, nil)
	failOnError(err, "Failed to bind queue Avisos_FumadorTabac to avisPolicia")

	//publicara el missatge de "policia" per la cua d'avisos per l'estanque, i per les cues
	// de Avisos_FumadorMistos i Avisos_FumadorTabac
	err = ch.Publish("avisPolicia", "", false, false, amqp.Publishing{Body: []byte("policia")})
	failOnError(err, "Failed to publish a messsage")

	fmt.Println("No sóm fumador. ALERTA! Que ve la policia!")
	fmt.Println(". . .")

}
