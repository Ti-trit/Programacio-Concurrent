package main

import (
	"fmt"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var numTabacs int = 0
var numMistos int = 0

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//obrir canal de conexio
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//defincio de coes a usar

	//Cues per enviar tabac i mistos
	var nomCues = [3]string{"tabac", "mistos", "peticions"}
	for i := 0; i < len(nomCues); i++ {
		_, err = ch.QueueDeclare(
			nomCues[i], //name
			false,      //durable
			false,      //delete when unused
			false,      //exclusive
			false,      //no-wait
			nil,        //arguments
		)
		if err != nil {
			log.Panicf("%s: %s %s", "Failed to declare queue", err, nomCues[i])

		}
	}

	//cua de comsum
	messages, err := ch.Consume(
		"peticions", //name
		"",          //consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")

	//exchange per rebre avisos de la policia pel fumador Xivato

	err = ch.ExchangeDeclare(
		"avisPolicia", //name
		"fanout",      //exchange type
		false,         //durable
		true,          // auto-deleted
		false,         // internal
		false,         // no-wait
		nil,           // arguments
	)
	failOnError(err, "Failed to declare exchange")

	_, err = ch.QueueDeclare("Avisos_estanquer", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	//cua per la qual rebra l'avis de policia

	fmt.Print("Hola, som l'estanquer il·legal\n")
	go veLaPolicia(ch)
	estanquer(ch, messages)

}

func estanquer(ch *amqp.Channel, messages <-chan amqp.Delivery) {

	forver := make(chan bool)
	go func() {
		for d := range messages {
			if string(d.Body) == "tabac" {
				numTabacs++
				log.Printf("He posat el tabac %d damunt la taula", numTabacs)
				//publicar el missatge per la cua de fumados de tabac
				msg := fmt.Sprintf("%d", numTabacs)
				err := ch.Publish(
					"",      //exchange
					"tabac", //routing key,
					false,   //mandatory
					false,
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,
						ContentType:  "text/plain",
						Body:         []byte(msg),
					})
				failOnError(err, "Failed to publish a message")

			} else if string(d.Body) == "mistos" {
				numMistos++
				log.Printf("He posat el misto %d damunt la taula", numMistos)
				//publicar el missatge per la cua de fumados de misos
				msg := fmt.Sprintf("%d", numMistos)

				err := ch.Publish(
					"",       //exchange
					"mistos", //routing key,
					false,    //mandatory
					false,    //immediate

					amqp.Publishing{
						DeliveryMode: amqp.Persistent,
						ContentType:  "text/plain",
						Body:         []byte(msg),
					})
				if err != nil {
					log.Printf("Error al publicar el missatge: %s", err)
				}

			}

		}
	}()
	<-forver
}

func veLaPolicia(ch *amqp.Channel) {

	messages, err := ch.Consume("Avisos_estanquer", "", false, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	for d := range messages {
		fmt.Println("\n Uyuyuy la policia! Men vaig")
		time.Sleep(1 * time.Second)
		d.Ack(false)
		//esborrar les cues
		var nomCues = [6]string{"tabac", "mistos", "peticions", "messages", "Avisos_FumadorMistos", "Avisos_FumadorTabac"}
		for i := 0; i < len(nomCues); i++ {
			time.Sleep(1 * time.Second)
			ch.QueueDelete(nomCues[i], false, false, true)
		}

		ch.QueueDelete("Avisos_estanquer", false, false, false)

		fmt.Println(". . . Men duc la taula ! ! ! !")
		//acaba

		os.Exit(0)
	}

}
