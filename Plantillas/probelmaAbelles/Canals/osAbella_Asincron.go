package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const (
	Procs           = 4
	THREADS_ABELLES = 10
	BUFFER_SIZE     = 10
	REPETICIONS     = 20
)

// var pot = 0
type Empty struct{}

func abella(id int, buffer chan int, done chan Empty, notPle chan Empty, notBuid chan Empty) {
	//produeixen mel mentre l'ós no esta menjant
	fmt.Printf("Hola, som l'abella %d \n", id)
	for i := 0; i < REPETICIONS; i++ {
		<-notPle //espera que el buffer no estigui ple
		time.Sleep(100 * time.Millisecond)
		data := rand.Intn(30)
		buffer <- data
		//pot+=1
		fmt.Printf("L'abella %d ha posat una porció %d de mel al pot \n", id, data)
		if len(buffer) == BUFFER_SIZE {
			//if pot==BUFFER_SIZE{
			fmt.Printf("Abella %d : El pot està ple. Avisaré a l'ós\n", id)
			notBuid <- Empty{}
		} else {
			notPle <- Empty{}
		}
	}
	done <- Empty{}

}

func os(done chan Empty, buffer chan int, notPle chan Empty, notBuid chan Empty) {
	for i := 0; i < REPETICIONS; i++ {
		<-notBuid //espera que le notifiquen les abelles
		fmt.Printf("Menjaré tota la mel, yyuppie\n")
		for len(buffer) > 0 {
			msg := <-buffer
			fmt.Printf("L'ós ha consumit la porció %d\n", msg)
		}
		//pot = 0
		fmt.Printf("M'en vaig a dormir\n")
		notPle <- Empty{}

		time.Sleep(100 * time.Millisecond)

	}
	done <- Empty{}

}

func main() {
	runtime.GOMAXPROCS(Procs)
	buffer := make(chan int, BUFFER_SIZE)
	done := make(chan Empty, 1)
	notPle := make(chan Empty, 1)
	notBuid := make(chan Empty, 1)
	//donar pas a les abelles
	notPle <- Empty{}
	for i := 0; i < THREADS_ABELLES; i++ {
		go abella(i, buffer, done, notPle, notBuid)
	}
	go os(done, buffer, notPle, notBuid)
	for i := 0; i < THREADS_ABELLES; i++ {
		<-done
	}
	<-done

	fmt.Print("Simulació acabada\n")
}
