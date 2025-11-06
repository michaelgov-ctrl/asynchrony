package io

import (
	"log"
	"sync"
)

type ExecutionImplementation int

const (
	Serial ExecutionImplementation = iota
	Concurrent
)

type Io struct {
	impl ExecutionImplementation
	wg   sync.WaitGroup
}

func NewIo(impl ExecutionImplementation) *Io {
	return &Io{
		impl: impl,
	}
}

func (io *Io) Do(funcs ...func()) {
	for _, f := range funcs {
		switch io.impl {
		case Serial:
			f()

		case Concurrent:
			io.wg.Go(f)

		default:
			log.Fatal("how the func did we get here")
		}
	}

	io.Wait()
}

func (io *Io) Wait() {
	io.wg.Wait()
}
