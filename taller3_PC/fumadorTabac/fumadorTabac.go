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

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//obrir canal de conexio
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//defincio de cues a usar

	//Cua per enviar tabac
	tabac, err := ch.QueueDeclare(
		"tabac", //name
		false,   //durable
		false,   //delete when unused
		false,   //exclusive
		false,   //no-wait
		nil,     //arguments
	)
	failOnError(err, "Failed to declare a queue")

	//cua per solicitar mistos
	peticions, err := ch.QueueDeclare(
		"peticions", //name
		false,       //durable
		false,       //delete when unused
		false,       //exclusive
		false,       //no-wait
		nil,         //arguments
	)
	failOnError(err, "Failed to declare a queue")

	//cua per rebre avisos de policia
	_, err = ch.QueueDeclare("Avisos_FumadorTabac", false, false, false, false, nil)
	failOnError(err, "Failed to declare a queue")

	fmt.Print("Sóc fumador. Tinc tabac però me falten mistos\n")

	//cua de consum de tabac
	consumicions_tabac, err := ch.Consume(
		tabac.Name, //queue
		"",         //consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")
	go veLaPolicia_FT(ch)

	fumadorTabac(ch, consumicions_tabac, peticions)

}

func fumadorTabac(ch *amqp.Channel, consumicions_tabac <-chan amqp.Delivery, peticions amqp.Queue) {

	for {
		err := ch.Publish(
			"",             //exchange
			peticions.Name, //routing key,
			false,          //mandatory
			false,          //immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte("tabac"),
			})

		failOnError(err, "Failed to publish a message")

		//agafar els mistos
		for d := range consumicions_tabac {
			fmt.Printf("He agafat el tabac %d. Gràcies!\n", d.Body)
		}

		time.Sleep(time.Second * 1) //espera 2 segons
		//demana més mistos
		fmt.Println("Me dones més tabac?")

	}

}

func veLaPolicia_FT(ch *amqp.Channel) {
	messages, err := ch.Consume("Avisos_FumadorTabac", "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	for range messages {
		fmt.Println("\n Anem que ve la policia!")
		time.Sleep(2 * time.Second)

		//acaba
		os.Exit(0)

	}
}
