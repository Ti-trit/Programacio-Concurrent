package main

import (
    "fmt"
    "runtime"
    "time"
)

const (
    Procs        = 4
    Philosophers = 5
    EatCount     = 100

    Thinking = iota
    Hungry
    Eating
)

type Empty struct{}
// Missatges que envien els filòsofs
type Request struct {
    c      chan Empty
    id     int
    status int
}

func philosopher(id int, done chan Empty, provider chan Request) {
    myChan := make(chan Empty)
    think := func() {
        time.Sleep(50 * time.Millisecond)
    }
    eat := func() {
        fmt.Printf("%d start eat\n", id)
        time.Sleep(100 * time.Millisecond)
        fmt.Printf("%d end eat\n", id)
    }
    pick := func() {
        // Quan vol menjar envia missatge al canal del proveïdor
        provider <- Request{id: id, c: myChan, status: Hungry}
        <-myChan
    }
    release := func() {
        // Ha acabat de menjar envia missatge al proveïdor
        provider <- Request{id: id, c: myChan, status: Thinking}
    }
    for i := 0; i < EatCount; i++ {
        think()
        pick()
        eat()
        release()
    }
    done <- Empty{}
}

func provider(channel chan Request) {
    var philosophers [Philosophers]Request
    right := func(i int) int {
        return (i + 1) % Philosophers
    }
    left := func(i int) int {
        return (i + Philosophers - 1) % Philosophers
    }
    canEat := func(i int) {
        r := right(i)
        l := left(i)
        if philosophers[i].status == Hungry &&
            philosophers[l].status != Eating &&
            philosophers[r].status != Eating {
            philosophers[i].status = Eating
            // Envia permís per menjar
            philosophers[i].c <- Empty{}
        }
    }
    for i := range philosophers {
        philosophers[i].status = Thinking
    }
    for {
        // Quan arriba un missatge el còpia
        m := <-channel
        philosophers[m.id] = m
        switch m.status {
        case Hungry:
            // Mira si el filòsof pot menjar
            canEat(m.id)
        case Thinking:
            // Mira si poden menjar els veïns
            canEat(left(m.id))
            canEat(right(m.id))
        }
    }
}

func main() {
    runtime.GOMAXPROCS(Procs)
    done := make(chan Empty, 1)
    // Procés per tota la taula per rebre
    // els missatges dels filòsofs
    providerChan := make(chan Request)
    go provider(providerChan)
    for i := 0; i < Philosophers; i++ {
        go philosopher(i, done, providerChan)
    }
    for i := 0; i < Philosophers; i++ {
        <-done
    }
    fmt.Printf("End\n")
}
