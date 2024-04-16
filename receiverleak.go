package main

import (
	"fmt"
	"time"
)

func main() {
	processRecords([]string{"record1", "record2", "record3", "record4", "record5"}, 3)
	time.Sleep(2 * time.Second)
}

func processRecords(list []string, workers int) {
	ch := make(chan string, len(list))
	for _, item := range list {
		ch <- item
	}
	for i := 0; i < workers; i++ {
		go worker(ch)
	}
}
func worker(input <-chan string) {
	for item := range input {
		fmt.Println("item received", item)
	}
	fmt.Printf("[worker]: shutting down\n")
}
