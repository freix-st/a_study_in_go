package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handlerHello)
	http.ListenAndServe("localhost:8080", nil)
}

func handlerHello(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Println("HELLO, WORLD!")
}
