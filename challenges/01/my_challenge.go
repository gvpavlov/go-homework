package main

import "fmt"

type PubSub struct {
	readers []chan string
	writer  chan string
}

func (ps *PubSub) read() {
	for msg := range ps.writer {
		for _, reader := range ps.readers {
			go func(reader chan string, msg string) {
				reader <- msg
			}(reader, msg)
		}
	}
}

func NewPubSub() *PubSub {
	ps := &PubSub{}
	ps.writer = make(chan string)
	go ps.read()
	return ps
}

func (ps *PubSub) Subscribe() chan string {
	c := make(chan string)
	ps.readers = append(ps.readers, c)
	return c
}

func (ps *PubSub) Publish() chan string {
	return ps.writer
}

func main() {
	ps := NewPubSub()
	a := ps.Subscribe()
	b := ps.Subscribe()
	c := ps.Subscribe()
	go func() {
		ps.Publish() <- "wat"
		ps.Publish() <- ("wat" + <-c)
	}()
	fmt.Printf("A recieved %s, B recieved %s and we ignore C!\n", <-a, <-b)
	fmt.Printf("A recieved %s, B recieved %s and C received %s\n", <-a, <-b, <-c)
}
