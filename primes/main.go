package main

import (
	"fmt"
	"os"
	"primes"
	"strconv"
)

//import primes
const EXPECTED_ARGS int = 3

func worker(numberToCheck <-chan int, foundPrimes chan<- int, finishedSignal chan<- bool) {
	for num := range numberToCheck {
		if true == primes.IsItAPrime(num) {
			foundPrimes <- num
		}
	}
	finishedSignal <- true
}

func printer(foundPrimes <-chan int) {
	for prime := range foundPrimes {
		fmt.Printf("Found a Prime: %d", prime)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < EXPECTED_ARGS {
		fmt.Printf("Not enough arguments \n")
		return
	}

	rangeStart, _ := strconv.Atoi(args[0])
	rangeStop, _ := strconv.Atoi(args[1])
	numberOfWorkers, _ := strconv.Atoi(args[2])

	numberToCheck := make(chan int, numberOfWorkers)
	foundPrimes := make(chan int, numberOfWorkers*4)
	finishedSignal := make(chan bool, numberOfWorkers)

	for i := 0; i < numberOfWorkers; i++ {
		go worker(numberToCheck, foundPrimes, finishedSignal)
	}

	go printer(foundPrimes)

	for num := rangeStart; num < rangeStop; num++ {
		numberToCheck <- num
	}
	//Signal to end the work for
	close(numberToCheck)

	for i := 0; i < numberOfWorkers; i++ {
		<-finishedSignal
	}
	//Signal to close the printer
	close(foundPrimes)

	//TODO: Wait for the printer to be closed
}
