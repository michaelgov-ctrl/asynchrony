package main

import "github.com/michaelgov-ctrl/asynchrony/io"

func main() {
	/*
		With a listener and a client in the same process, the two
		must overlap, so there is a requirement of concurrency.
	*/

	listener, client := NewTCPListener(), NewClient()

	// This can not be run serially because it will
	// deadlock on Accept forever and the client never runs
	io.NewIo(io.Serial).Do(
		listener.Accept,
		client.EstablishConnection,
	)

	// Since concurrency is required for this workflow to function.
	// The code should express that these tasks MUST be progressed concurrently.
	// Under the hood io.Concurrent can be parallel execution or task switching
	io.NewIo(io.Concurrent).Do(
		listener.Accept,
		client.EstablishConnection,
	)

	// This example exhibits asynchrony because it doesn't matter whether the listener
	// accepts or the client reaches out first, however there is a dependency on concurrency.
}
