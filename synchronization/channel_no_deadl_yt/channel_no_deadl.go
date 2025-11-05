package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	chats   []string
	friends []string
}

func main() {
	now := time.Now()
	id := getUserByName("john")
	fmt.Println(id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message, 2)

	wg.Add(2)

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	wg.Wait()
	close(ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println(time.Since(now))
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)

	ch <- &Message{
		friends: []string{
			"john",
			"jane",
			"joe",
			"james",
			"tiago",
		},
	}

	/*
		You can close the channel here, and it may or may not work:
		- if it works, if you add more go routines, then you transfer the closing to the last one
		- if it doesn't work (incomplete data), it's because this go routine finished faster than the other one
	*/

	wg.Done()
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)

	ch <- &Message{
		chats: []string{
			"john",
			"jane",
			"joe",
		},
	}

	wg.Done()
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s", name)
}
