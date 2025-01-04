package main

import (
    "fmt"
    "runtime"
    "time"
)

const (
    Procs    = 4
    FILOSOFS = 5
    EatCount = 100
)

type Empty struct{}
type Fork chan Empty
// Array de canals 1 per bastonet
type Forks [FILOSOFS]Fork

func think() {
    time.Sleep(50 * time.Millisecond)
}

func right(i int) int {
    return (i + 1) % FILOSOFS
}

func pickForks(id int, forks Forks) {
  // tots els filòsofs agafen primer el bastonet de l'esquerra
    if id < right(id) {
        <-forks[id]
        <-forks[right(id)]
  // escepte el darrer que agafa primer el de la dreta
    } else {
        <-forks[right(id)]
        <-forks[id]
    }
}

func eat(id int) {
    fmt.Printf("%d start eat\n", id)
    time.Sleep(100 * time.Millisecond)
    fmt.Printf("%d end eat\n", id)
}

func releaseForks(id int, forks Forks) {
    forks[id] <- Empty{}
    forks[right(id)] <- Empty{}
}

func philosopher(id int, done chan Empty, forks Forks) {
    for i := 0; i < EatCount; i++ {
        think()
        pickForks(id, forks)
        eat(id)
        releaseForks(id, forks)
    }
    done <- Empty{}
}

func main() {
    runtime.GOMAXPROCS(Procs)
    done := make(chan Empty, 1)
    var forks Forks

    for i := range forks {
      // Canal asíncron buffer 1
        forks[i] = make(chan Empty, 1)
        forks[i] <- Empty{}
    }
    for i := 0; i < FILOSOFS; i++ {
        go philosopher(i, done, forks)
    }
    for i := 0; i < FILOSOFS; i++ {
        <-done
    }

    fmt.Printf("End\n")
}
