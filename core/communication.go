package core

import "image"

// Communication :
type Communication struct {
	CPU      chan CPU
	Screen   chan image.Image
	Sound    chan bool
	keypress chan int
	shutdown chan struct{}
}

// NewChan :
func newChan() Communication {
	return Communication{
		CPU:      make(chan CPU),
		Screen:   make(chan image.Image),
		Sound:    make(chan bool),
		keypress: make(chan int),
		shutdown: make(chan struct{}),
	}
}
