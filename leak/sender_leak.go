package leak

import (
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"time"
)

func SenderLeak() {
	// Capture starting number of goroutines.
	startingGs := runtime.NumGoroutine()

	if err := process("gophers"); err != nil {
		log.Print(err)
	}

	// Hold the program from terminating for 1 second to see
	// if any goroutines created by process terminate.
	time.Sleep(time.Second)

	// Capture ending number of goroutines.
	endingGs := runtime.NumGoroutine()

	// Report the results.
	fmt.Println("========================================")
	fmt.Println("Number of goroutines before:", startingGs)
	fmt.Println("Number of goroutines after :", endingGs)
	fmt.Println("Number of goroutines leaked:", endingGs-startingGs)
}

// process is the work for the program. It finds a record
// then prints it. It fails if it takes more than 100ms.
func process(term string) error {

	// Create a context that will be canceled in 100ms.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	ch := make(chan struct{})

	// Launch a goroutine to find the record
	go func() {
		search(term)
		ch <- struct{}{}
	}()

	// Block waiting to either receive from the goroutine's
	// channel or for the context to be canceled.
	select {
	case <-ctx.Done():
		return errors.New("search canceled")
	case <-ch:
		fmt.Println("Received:")
		return nil
	}
}

// search simulates a function that finds a record based
// on a search term. It takes 200ms to perform this work.
func search(term string) {
	time.Sleep(200 * time.Millisecond)
}
