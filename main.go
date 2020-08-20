package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)

	go populate(channel1)
	go fanOutIn(channel1, channel2)

	for value := range channel2 {
		fmt.Println(value)
	}

	fmt.Println("Going to the run")
}

func populate(channel chan int) {
	for i := 0; i < 100; i++ {
		channel <- 1
	}
	close(channel)
}

func fanOutIn(channel1, channel2 chan int) {
	var wg sync.WaitGroup
	for channel1Message := range channel1 {
		wg.Add(1)
		go func(channel2Message int) {
			channel2 <- timeConsumingWork(channel2Message)
			wg.Done()
		}(channel1Message)
	}
	wg.Wait()
	close(channel2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

