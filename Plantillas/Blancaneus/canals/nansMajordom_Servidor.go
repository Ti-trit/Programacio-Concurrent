/*
*
Solució del problema de nans i el majordom amb canals síncrons
*
*/
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

// globar channels
var permisSeure = make(chan int)
var permisMenjar = make(chan int) //canal senyalitzar que un nan vol menjar
var permisAixecarse = make(chan int)
var esperaMenjar = make(chan Empty)
var esperaCadira = make(chan Empty)
var esperaAixecarse = make(chan Empty)
var noms = [NANS]string{"Beethoven", "Tenma", "Johan", "Titrit", "Belqis", "Mario", "Andreas"}

func majordom() {
	var cadires = CADIRES
	var nomNanEsperant string
	numNansEsperant := 0
	var nansEsperants [NANS] int
	
	for {
		select {
		case id := <-permisSeure:
			if cadires > 0 {
				fmt.Printf("*******El majordom fa seure a %s \n", noms[id])
				cadires--
				//donar permis
				esperaCadira <- Empty{}
			} else {
				fmt.Printf("******* El majordom fa esperar a %s, totes les cadires estan ocupades\n", noms[id])
				nansEsperants[id] = 1
				numNansEsperant++
			}

		case id := <-permisMenjar:
			time.Sleep(time.Duration(rand.Intn(17)) * time.Millisecond)
			fmt.Printf("*******El majordom serveix a %s \n", noms[id])
			esperaMenjar <- Empty{}

		case id := <-permisAixecarse:
			fmt.Printf("******* El majordom dona permís per anar-se'n a %s\n", noms[id])
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			esperaAixecarse <- Empty{}

			//comprova si hi ha nans esperant per una cadira

			if numNansEsperant > 0 {
				for i := 0; i < NANS; i++ {
					if nansEsperants[i] == 1 {
						nomNanEsperant = noms[i]
						nansEsperants[i]=0 // alliberar		
						break
					}
				}
				fmt.Printf("******* El majordom fa seure a %s a la cadira de %s \n", nomNanEsperant, noms[id])
				
				numNansEsperant--
				esperaCadira <- Empty{}

			} else {
				//ningu esta esperant per una cadira
				cadires++
			}

		}
	}

}

func nan(id int, done chan Empty) {

	fmt.Printf("Hola el meu nom és %s\n", noms[id])

	for i := 0; i < MENJADES; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		//anar a la mina
		fmt.Printf("%s treballa a la mina\n", noms[id])
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("%s ha arribat de la mina i espera una cadira\n", noms[id])

		//demanar cadira
		permisSeure <- id
		<-esperaCadira
		fmt.Printf("%s ja seu i demana ser servit\n", noms[id])
		//demenar menjar
		permisMenjar <- id
		<-esperaMenjar
		fmt.Printf("%s ja menja!!!!!\n", noms[id])
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)

		//permis per aixecar-se
		fmt.Printf(" %s ha acabat de menjar i demana permís per aixecar-se\n", noms[id])
		permisAixecarse <- id
		<-esperaAixecarse

	}
	fmt.Printf(" %s se'n va a dormir\n", noms[id])
	time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
	done <- Empty{}
}

func main() {
	runtime.GOMAXPROCS(Procs)
	done := make(chan Empty, 1)

	//definir noms pels nans

	go majordom()
	for i := 0; i < NANS; i++ {
		go nan(i, done)
	}
	for i := 0; i < NANS; i++ {
		<-done
	}
	fmt.Print("Simulació acabada\n")
}
