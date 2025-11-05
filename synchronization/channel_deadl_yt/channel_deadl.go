package main

import (
	"fmt"
	"log"
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

	ch := make(chan *Message, 2)

	go getUserChats(id, ch)
	go getUserFriends(id, ch)

	for msg := range ch {
		log.Println(msg)
	}

	log.Println(time.Since(now))
}

func getUserFriends(id string, ch chan<- *Message) {
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
}

func getUserChats(id string, ch chan<- *Message) {
	time.Sleep(time.Second * 2)

	ch <- &Message{
		chats: []string{
			"john",
			"jane",
			"joe",
		},
	}
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("%s", name)
}
