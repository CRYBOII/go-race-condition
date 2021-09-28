package main

import (
	"fmt"
	"sync"
)

var chDeposit = make(chan int)
var chBalance = make(chan int)

func main() {

	go bankSystem()

	w := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		w.Add(1)

		go deposit(100, w)

	}

	w.Wait()

	fmt.Println("Balance is : ", finalBalance())

}

func deposit(x int, w *sync.WaitGroup) {
	chDeposit <- x
	w.Done()

}

func finalBalance() int {

	return <-chBalance

}

func bankSystem() {

	balance := 0

	for {
		select {
		case x := <-chDeposit:
			balance += x

		case chBalance <- balance:

		}

	}

}
