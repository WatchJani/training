package main

import (
	"fmt"
	"time"
)

type Communication struct {
	trigger chan struct{}
}

func NewCommunication() *Communication {
	return &Communication{
		trigger: make(chan struct{}),
	}
}

func RunTraining1() {
	c := NewCommunication()

	go c.Send()

	c.Listen()

	fmt.Println("done")
}

func (c *Communication) Send() {
	time.Sleep(2 * time.Second)
	c.trigger <- struct{}{}
}

func (c *Communication) Listen() {
	<-c.trigger
}
