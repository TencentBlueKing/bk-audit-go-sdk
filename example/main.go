package main

import "time"

func main() {
	// init client
	initClient()
	// add event
	client.AddEvent(&viewHost, &host, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
	// wait channel
	time.Sleep(3 * time.Second)
}
