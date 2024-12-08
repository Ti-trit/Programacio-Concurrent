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

	//Cua per enviar mistos
	mistos, err := ch.QueueDeclare(
		"mistos", //name
		false,    //durable
		false,    //delete when unused
		false,    //exclusive
		false,    //no-wait
		nil,      //arguments
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

	_, err = ch.QueueDeclare("Avisos_FumadorMistos", false, false, false, false, nil)
	failOnError(err, "Failed to declare Avisos_FumadorMistos queue")

	// err = ch.ExchangeDeclare("avisPolicia", //name
	// 	"fanout", //exchange type
	// 	false,    //durable
	// 	false,    // auto-deleted
	// 	false,    // internal
	// 	false,    // no-wait
	// 	nil,      // arguments
	// )
	// failOnError(err, "Failed to declare an exchange")

	//establir la vinculacio en cada instancia del fumadorTabac
	// err = ch.QueueBind("Avisos_FumadorMistos", "", "avisPolicia", false, nil)
	// failOnError(err, "Failed to bind queue Avisos_FumadorMistos to avisPolicia") //cua de consum de tabac

	//cua de consum de tabac
	consumicions_mistos, err := ch.Consume(
		mistos.Name, //queue
		"",          //consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	failOnError(err, "Failed to register a consumer")
	fmt.Print("Sóc fumador. Tinc tabac però me falten mistos\n")

	go veLaPolicia_FM(ch)

	fumadorMistos(ch, consumicions_mistos, peticions)

}

func fumadorMistos(ch *amqp.Channel, consumicions_mistos <-chan amqp.Delivery, peticions amqp.Queue) {

	mes_peticio := make(chan bool)
	go func() {
		for {
			err := ch.Publish(
				"",             //exchange
				peticions.Name, //routing key,
				false,          //mandatory
				false,          //immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte("mistos"),
				})
			failOnError(err, "Failed to publish a message")
			<-mes_peticio //espera rebre una resposta per demanar més mistos

			time.Sleep(time.Second * 1) //espera 1 segons
			//demana més mistos
			fmt.Println(". . .\nMe dones un altre misto?")
		}
	}()
	//agafar els mistos
	for d := range consumicions_mistos {
		numMistos, err := strconv.Atoi(string(d.Body))
		failOnError(err, "Failed to convert d.Body to an integer")

		fmt.Printf("He agafat el misto %d. Gràcies!\n", numMistos)
		//permis per demanar mes mistos
		mes_peticio <- true
		d.Ack(false) //confirmar que ha rebut el missatge actual
		time.Sleep(time.Second * 1)

	}

}

func veLaPolicia_FM(ch *amqp.Channel) {
	messages, err := ch.Consume("Avisos_FumadorMistos", "", false, false, false, false, nil)
	failOnError(err, "Failed to register a consumer")

	for d := range messages {
		fmt.Println("\n Anem que ve la policia!")
		d.Ack(false)
		//time.Sleep(2 * time.Second)
		err = ch.Publish("avisPolicia", "", false, false, amqp.Publishing{Body: []byte("policia")})
		failOnError(err, "Failed to publish a messsage")

		//acaba
		os.Exit(0)

	}
}