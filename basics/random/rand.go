package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"time"
)

func main() {
	// seed
	seed := time.Now().UnixNano()
	sourceStr := strconv.FormatInt(seed, 10)
	fmt.Println(sourceStr[:len(sourceStr)/2])

	firstHalf, err := strconv.ParseUint(sourceStr[:len(sourceStr)/2], 10, 64)
	if err != nil {
		panic(err)
	}

	secondHalf, err := strconv.ParseUint(sourceStr[len(sourceStr)/2:], 10, 64)
	if err != nil {
		panic(err)
	}

	source := rand.NewPCG(firstHalf, secondHalf)

	// generate
	generator := rand.New(source)
	fmt.Println(generator.IntN(100))
}
