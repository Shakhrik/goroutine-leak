package main

import "goroutine-leak/leak"

func main() {
	// Sender leak
	leak.SenderLeak()

	// Receiver leak
	leak.ReceiverLeak()
}
