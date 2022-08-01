package main

import (
	"fmt"
	"sync"
)

type Chopstick = sync.Mutex

type Philosopher struct {
	id    int
	meals int
	left  *Chopstick
	right *Chopstick
}

var wg sync.WaitGroup

const PhilosopherCount = 5
const MaxConcurrentMeals = 2
const MaxMeals = 3

func (p *Philosopher) eat(request chan *Philosopher, finish chan bool) {
	for meals := 1; meals <= MaxMeals; meals++ {
		request <- p

		fmt.Printf("starting to eat %d [%d/3]\n", p.id, meals)

		// pick up the chopsticks in any order
		wgLock := sync.WaitGroup{}
		wgLock.Add(2)
		go func() { p.left.Lock(); wgLock.Done() }()
		go func() { p.right.Lock(); wgLock.Done() }()
		wgLock.Wait()

		fmt.Printf("finished to eat %d [%d/3]\n", p.id, meals)
		p.right.Unlock()
		p.left.Unlock()
		finish <- true
	}

	wg.Done()
}

// host is waiting finish event before allowing a new request
func host(request chan *Philosopher, finish chan bool) {
	for range finish {
		<-request
		fmt.Println("Exit")
	}
}

func main() {
	// Set up the objects
	chopsticks := make([]*Chopstick, PhilosopherCount)

	for i := range chopsticks {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, PhilosopherCount)

	for i := range chopsticks {
		philosophers[i] = &Philosopher{i + 1, 0, chopsticks[i], chopsticks[(i+1)%PhilosopherCount]}
	}

	// request is a channel of size 2, allowing maximum 2 concurrent meals
	request := make(chan *Philosopher, MaxConcurrentMeals)
	finish := make(chan bool)

	// Start the process with the host function
	wg.Add(5)
	go host(request, finish)

	for _, p := range philosophers {
		go p.eat(request, finish)
	}
	wg.Wait()
}
