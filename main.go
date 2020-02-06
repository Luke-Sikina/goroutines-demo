package main

import (
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	threads := parseArgs()

	var waitGroup sync.WaitGroup
	waitGroup.Add(threads)

	for i := 0; i < threads; i ++ {
		go doSomething(i, &waitGroup)
	}

	waitGroup.Wait()

	log.Printf("All routines complete. Exiting.\n")
	os.Exit(0)
}

func parseArgs() int {
	if len(os.Args) == 1 {
		log.Fatalf("Provide thread count")
	}
	threads, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Invalid thread count: %s", os.Args[1])
	}
	return threads
}

func doSomething(input int, signal *sync.WaitGroup) {
	defer signal.Done()
	time.Sleep(time.Duration(input % 5) * time.Second)
	log.Printf("Routine %d exiting", input)
}