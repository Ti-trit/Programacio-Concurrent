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

var nomCues = [3]string{"tabac", "peticions", "Avisos_FumadorTabac"}

func main() {
	//establir conexió amb rabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	//obrir canal de conexio
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//defincio de cues a usar

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

	fmt.Print("Sóc fumador. Tinc mistos però me falta tabac\n")

	go veLaPolicia_FT(ch)

	fumadorTabac(ch)

}

func fumadorTabac(ch *amqp.Channel) {
	//cua per consumir els tabacs que envia l'estanquer per la cua tabac
	consumicions_tabac, err := ch.Consume(
		"tabac", //queue
		"",      //consumer
		false,   // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	failOnError(err, "Failed to register consumer: consumicions_tabac")
	//canal per sincronitzar l'eviament i el processament de peticions
	mes_peticio := make(chan bool)
	go func() {

		for {
			err := ch.Publish(
				"",          //exchange
				"peticions", //routing key,
				false,       //mandatory
				false,       //immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte("tabac"),
				})

			failOnError(err, "Failed to publish a message to get tabac")
			<-mes_peticio //esperem rebre una resposta per poder demanar més tabac

			time.Sleep(time.Second * 1) //espera 1 segons
			//demana més tabac
			fmt.Println(". . .\nMe dones més tabac?")
		}
	}()
	//processar els missatges que envia l'estanquer
	for d := range consumicions_tabac {
		//convertim el numero de tabac, enviat com a string, a un enter
		numTabacs, err := strconv.Atoi(string(d.Body))
		failOnError(err, "Failed to convert d.Body to an integer")

		fmt.Printf("He agafat el tabac %d. Gràcies!\n", numTabacs)
		//permis per poder demanar mes mistos
		mes_peticio <- true
		d.Ack(false) //confirmar que s'ha rebut el missatge actual
		time.Sleep(time.Second * 1)
	}

}

func veLaPolicia_FT(ch *amqp.Channel) {
	//Definem la cua que consumirá els avisos que s'enviaren als fumadors de tabac
	messages, err := ch.Consume("Avisos_FumadorTabac", "", false, false, false, false, nil)
	failOnError(err, "Failed to register a consumer:Avisos_FumadorTabac ")
	//Com que l'exchange només envia un missatge a una cua, i no a tots els clients
	//que consumeixn d'ella. Per aixo tornam a ficas el missatge d'avis per la resta de fumadors de tabac

	for d := range messages {
		fmt.Println("\nAnem que ve la policia!")

		d.Ack(false) //afirmar que s'ha rebut l'avis correctament
		//publicar de nou el missatge d'avis a les cues vinculades a l'exchange avisPolicia
		err = ch.Publish("avisPolicia", "", false, false, amqp.Publishing{Body: []byte("policia")})
		failOnError(err, "Failed to publish a messsage")

		//acaba
		os.Exit(0)

	}

}
