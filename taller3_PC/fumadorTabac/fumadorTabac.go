package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
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
	failOnError(err, "Failed to declare a tabac queue")

	//cua per solicitar mistos
	peticions, err := ch.QueueDeclare(
		"peticions", //name
		false,       //durable
		false,       //delete when unused
		false,       //exclusive
		false,       //no-wait
		nil,         //arguments
	)
	failOnError(err, "Failed to declare a peticions queue")

	//cua per rebre avisos de policia
	_, err = ch.QueueDeclare("Avisos_FumadorTabac", false, false, false, false, nil)
	failOnError(err, "Failed to declare queue Avisos_FumadorTabac ")

	err = ch.ExchangeDeclare("avisPolicia", //name
		"fanout", //exchange type
		false,    //durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//establir la vinculacio en cada instancia del fumadorTabac
	err = ch.QueueBind("Avisos_FumadorTabac", "", "avisPolicia", false, nil)
	failOnError(err, "Failed to bind queue Avisos_FumadorTabac to avisPolicia") //cua de consum de tabac

	consumicions_tabac, err := ch.Consume(
		tabac.Name, //queue
		"",         //consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register consumer tabac")

	fmt.Print("Sóc fumador. Tinc mistos però me falta tabac\n")

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

		failOnError(err, "Failed to publish a message to get tabac")

		//agafar els mistos
		for d := range consumicions_tabac {
			numTabacs, err := strconv.Atoi(string(d.Body))
			if err != nil {
				log.Printf("Error convertint el missatge a número: %v", err)
				continue
			}
			fmt.Printf("He agafat el tabac %d. Gràcies!\n", numTabacs)
			break
		}

		time.Sleep(time.Second * 1) //espera 1 segons
		fmt.Println(". . .")
		//demana més mistos
		fmt.Println("Me dones més tabac?")

	}

}

func veLaPolicia_FT(ch *amqp.Channel) {
	messages, err := ch.Consume("Avisos_FumadorTabac", "", true, false, false, false, nil)
	failOnError(err, "Failed to register a consumer:Avisos_FumadorTabac ")
	//tornar a ficas el missatge d'avis per la resta de fumadors de tabac

	for range messages {
		fmt.Println("\n Anem que ve la policia!")
		//	time.Sleep(1 * time.Second)
		err = ch.Publish("avisPolicia", "", false, false, amqp.Publishing{Body: []byte("policia")})
		failOnError(err, "Failed to publish a messsage")
		os.Exit(0)

	}
	//avisar als altres possibles fumados de tabac

}
