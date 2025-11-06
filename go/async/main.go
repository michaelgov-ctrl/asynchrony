package main

import (
	"fmt"

	"github.com/michaelgov-ctrl/asynchrony/io"
)

func main() {
	/*
		Writing FileA and FileB is asynchronous, their relative order doesn’t matter
		for correctness.
		They can be executed serially, or at the same time, in any order and all
		outcomes are still correct. Asynchrony does not require concurrency.
	*/

	fmt.Println("started serial execution")
	io.NewIo(io.Serial).Do(
		appendFileA,
		appendFileB,
	)

	fmt.Println("started concurrent execution")
	io.NewIo(io.Concurrent).Do(
		appendFileA,
		appendFileB,
		appendFileA,
		appendFileB,
	)
}
