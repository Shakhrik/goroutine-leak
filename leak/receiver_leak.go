package leak

import (
	"fmt"
	"time"
)

func ReceiverLeak() {
	processRecords([]string{"record1", "record2", "record3", "record4", "record5"}, 3)
	time.Sleep(2 * time.Second)
}
func processRecords(list []string, workers int) string {
	ch := make(chan string, len(list))
	for _, item := range list {
		ch <- item
	}
	for i := 0; i < workers; i++ {
		go worker(ch)
	}

	return "some value"
}
func worker(input <-chan string) {
	for item := range input {
		fmt.Println("item received", item)
	}
	fmt.Println("[worker]: shutting down")
}
