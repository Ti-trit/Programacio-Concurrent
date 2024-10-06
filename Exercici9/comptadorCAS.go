/**
*Implimentació d'un comptador amb exclusió mútua usant la primitiva
* compare&swap
 */
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

const (
	Procs      = 4 //fils d'execució del sistema operatiu
	MaxCount   = 10000000
	Goroutines = 4 //fils del programa
)

var (
	counter int
	mutex   int32
)

func lock() {
	var local int32
	for !atomic.CompareAndSwapInt32(&mutex, local, 1) {
	}
}

func unlock() {
	mutex = 0
}

func run(id, counts int, done chan bool) {
	for i := 0; i < counts; i++ {
		lock()
		counter++
		unlock()
	}
	fmt.Printf("End %d counter: %d\n", id, counter)
	done <- true
}

func main() {
	runtime.GOMAXPROCS(Procs)
	done := make(chan bool, 1)

	for i := 0; i < Goroutines; i++ {
		go run(i, MaxCount/Goroutines, done)
	}

	for i := 0; i < Goroutines; i++ {
		<-done
	}

	fmt.Printf("Counter value: %d Expected: %d\n", counter, MaxCount)
}
