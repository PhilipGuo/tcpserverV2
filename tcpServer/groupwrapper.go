package tcpserver

import (
	"sync"
)

// WaitGroupWrapper object
//
type WaitGroupWrapper struct {
	sync.WaitGroup
}

//Wrap a goroutine
//
func (w *WaitGroupWrapper) Wrap(cb func()) {
	w.Add(1)
	go func() {
		cb()
		w.Done()
	}()
}
