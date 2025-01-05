package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

const (
	Procs    = 4
	NANS     = 7
	MENJADES = 2
	CADIRES  = 4
)

type Empty struct{}

func nan(nom string, perMenjar chan Empty, cadires chan Empty, Adormir chan Empty, prepararMenjar chan string, done chan Empty) {

	fmt.Printf("Hola, som el nan %s\n", nom)
	for i := 0; i < MENJADES; i++ {
		//va a la mina
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("El nan %s ha anat a la mina\n", nom)
		//tornar
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("El nan %s ha tornat de la mina\n", nom)
		//demanra cadira
		<-cadires
		fmt.Printf("El nan %s ha agafat una cadira\n", nom)
		//demana menjar
		prepararMenjar <- nom
		//espera que le sirveixen el plat
		<-perMenjar
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("<<---El nan %s ha menjat el %d seu plat\n", nom, i+1)
		//alliberar una cadia
		cadires <- Empty{}

	}
	//els nans se'n van a dormir
	fmt.Printf("----->>El nan %s ha anat a dormir\n", nom)
	Adormir <- Empty{}

	done <- Empty{}
	

}

func Blancaneus(perMenjar chan Empty, Adormir chan Empty, prepararMenjar chan string, done chan Empty) {
	var comptador int = 0
	for comptador < NANS*MENJADES {
		//espera que els nans demanen menjar
		select {
		case nom := <-prepararMenjar:
			fmt.Println("Blancaneus està preparant el menjar pel nan [", nom, "]")
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Simula temps preparant el menjar

			perMenjar <- Empty{}
			comptador++
		default:
			// No hi ha més demandes, Blancaneus pot anar a passejar
			fmt.Println("Blancaneus s'ha anat a passejar")
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}

	for i := 0; i < NANS; i++ {
		<-Adormir //espera que tots els nans se'n vagin a dormir
	}

	fmt.Printf("Els nans s'han anat a dormir\n")
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	fmt.Printf("Blancaneus ha anat a dormir també\n")

	done <- Empty{}

}

func main() {
	runtime.GOMAXPROCS(Procs)
	done := make(chan Empty, 1)
	perMenjar := make(chan Empty, CADIRES) //canal senyalitzar que un nan vol menjar
	cadires := make(chan Empty, CADIRES)
	Adormir := make(chan Empty, NANS)
	prepararMenjar := make(chan string, 1)

	//definir noms pels nans
	noms := [NANS]string{"Beethoven", "Tenma", "Johan", "Titrit", "Belqis", "Mario", "Andreas"}
	perMenjar <- Empty{}
	cadires <- Empty{}
	Adormir <- Empty{}
	go Blancaneus(perMenjar, Adormir, prepararMenjar, done)
	for i := 0; i < NANS; i++ {
		go nan(noms[i], perMenjar, cadires, Adormir, prepararMenjar, done)
	}
	for i := 0; i < NANS; i++ {
		<-done
	}
	<-done //blancaneus
	fmt.Print("Simulació acabada\n")
}
