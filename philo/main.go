package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Philosopher struct {
	Chopstick sync.Mutex //chan bool
	Name      string
	Neighbor  *Philosopher
}

func New(name string, n *Philosopher) *Philosopher {
	//c := make(chan bool, 1)
	p := &Philosopher{
		Name:     name,
		Neighbor: n,
		//Chopstick: c,
	}

	//p.Chopstick <- true
	return p
}

func (p *Philosopher) Eat() {
	fmt.Println(p.Name, "is eating")
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (p *Philosopher) Think() {
	fmt.Println(p.Name, "is thinking")
	time.Sleep(time.Duration(rand.Int63n(1e9)))
}

func (p *Philosopher) GetChopsticks() {
	timeout := make(chan bool, 1)
	fmt.Println(p.Name, " want chopsticks.")
	go func() { time.Sleep(1 * time.Second); timeout <- true }()
	//<-p.Chopstick
	p.Chopstick.Lock()
	fmt.Println(p.Name, " got his chopstick.")
	select {
	/*case <-p.Neighbor.Chopstick:
	fmt.Println(p.Name, "got ", p.Neighbor.Name, "'s chopstick.")
	fmt.Println(p.Name, " has two chopsticks.")
	return*/
	case <-timeout:
		//p.Chopstick <- true
		p.Chopstick.Unlock()
		p.Think()
		p.GetChopsticks()
	default:
		p.Neighbor.Chopstick.Lock()
		fmt.Println(p.Name, "got ", p.Neighbor.Name, "'s chopstick.")
		fmt.Println(p.Name, " has two chopsticks.")
		return
	}
}

func (p *Philosopher) ReleaseChopsticks() {
	/*p.Chopstick <- true
	p.Neighbor.Chopstick <- true*/
	p.Chopstick.Unlock()
	p.Neighbor.Chopstick.Unlock()
}

func (p *Philosopher) dine(table chan *Philosopher) {
	p.Think()
	p.GetChopsticks()
	p.Eat()
	p.ReleaseChopsticks()
	table <- p
}

func main() {
	names := []string{
		"Melkor",
		"Dak",
		"Cognet",
		"JP",
		"ZLeTraitre",
	}
	philosophers := make([]*Philosopher, len(names))
	var phil *Philosopher
	for i, name := range names {
		phil = New(name, phil)
		philosophers[i] = phil
	}
	philosophers[0].Neighbor = phil
	fmt.Printf("There are %v philosophers sitting at a table.\n", len(philosophers))
	fmt.Println("They each have one chopstick, and must borrow from their neighbor to eat.")
	announce := make(chan *Philosopher)
	for _, phil := range philosophers {
		go phil.dine(announce)
	}
	for i := 0; i < len(names); i++ {
		phil := <-announce
		fmt.Printf("%v is done dining.\n", phil.Name)
	}
}
