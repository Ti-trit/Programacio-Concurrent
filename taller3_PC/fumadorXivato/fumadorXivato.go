package main

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var nomCues = [3]string{"Avisos_estanquer", "Avisos_FumadorMistos", "Avisos_FumadorTabac"}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//obrir canal de conexio
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//declarar les cues d'avis

	for i := 0; i < len(nomCues); i++ {
		_, err = ch.QueueDeclare(
			nomCues[i], //name
			false,      //durable
			false,      //delete when unused
			false,      //exclusive
			false,      //no-wait
			nil,        //arguments
		)
		failOnError(err, "Failed to declare a queue")

	}

	//exchange declare

	err = ch.ExchangeDeclare("avisPolicia", //name
		"fanout", //exchange type
		false,    //durable
		true,     // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//vincular cada una de les cues a les que enviarem l'avis de la policia

	for i := 0; i < len(nomCues); i++ {
		err = ch.QueueBind(nomCues[i], "", "avisPolicia", false, nil)
		if err != nil {
			fmt.Fprintln(os.Stdout, []any{"Failed to bin queue %s to exchange avisPolicia", nomCues[i]}...)
		}
	}

	//publicara el missatge de "policia" per la cua Avisos_estanquer, i per les cues
	// de Avisos_FumadorMistos i Avisos_FumadorTabac
	err = ch.Publish("avisPolicia", "", false, false, amqp.Publishing{Body: []byte("policia")})
	failOnError(err, "Failed to publish a messsage: policia ")

	fmt.Println("No sóm fumador. ALERTA! Que ve la policia!\n. . .")

}
