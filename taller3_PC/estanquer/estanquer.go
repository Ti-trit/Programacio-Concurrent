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

//variables globals per comptar numero de tabacs i mistos que es van posant a la taula
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

	var nomCues = [4]string{"tabac", "mistos", "peticions", "Avisos_estanquer"}
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

	fmt.Println("Hola, som l'estanquer il·legal")
	go veLaPolicia(ch)
	estanquer(ch)

}

func estanquer(ch *amqp.Channel) {
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
	failOnError(err, "Failed to register a consumer: messages")

	forver := make(chan bool)
	go func() {
		for d := range messages {
			if string(d.Body) == "tabac" {
				numTabacs++
				fmt.Printf("He posat el tabac %d damunt la taula\n", numTabacs)
				//publicar el missatge per la cua de fumados de tabac
				//l'estanquer envia el numero de tabacs com a missatge a la cua
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
				fmt.Printf("He posat el misto %d damunt la taula\n", numMistos)
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
		fmt.Println("\nUyuyuy la policia! Men vaig")
		time.Sleep(1 * time.Second)
		//confirmar que s'ha rebut el missatge actual
		d.Ack(false)
		//esborrar les cues amb un marge de temps, per donar temps a que l'avis arribi a tots els clients
		var nomCues = [6]string{"tabac", "mistos", "peticions", "messages", "Avisos_FumadorMistos", "Avisos_FumadorTabac"}
		for i := 0; i < len(nomCues); i++ {
			time.Sleep(1 * time.Second)
			ch.QueueDelete(nomCues[i], false, false, true)
		}
		//no es borra en el bucle, perque té el parametre noWait different de la resta de cues		
		ch.QueueDelete("Avisos_estanquer", false, false, false)

		fmt.Println(". . . Men duc la taula ! ! ! !")
		//acaba

		os.Exit(0)
	}

}
