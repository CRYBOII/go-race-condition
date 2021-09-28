package main

import (
	"fmt"
	"sync"
)

var chDeposit = make(chan int, 1000)
var chBalance = make(chan int)

func main() {

	go bankSystem()
	fmt.Println("Balance is : ", finalBalance())

	w := &sync.WaitGroup{}

	for i := 0; i < 100000; i++ {

		w.Add(1)

		go deposit(100, w)

	}
	fmt.Println("Balance is : ", finalBalance())

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
		if len(chDeposit) == 0 {
			select {
			case x := <-chDeposit:
				balance += x

			case chBalance <- balance:

			}
		} else {
			x := <-chDeposit
			balance += x

		}

	}

}
