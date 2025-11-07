package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	ID       int
	Cost     int
	Distance int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph map[int]map[int]int, start int) map[int]int {
	distances := make(map[int]int)
	for node := range graph {
		distances[node] = 1<<31 - 1 // infinite
	}
	distances[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Node{ID: start, Distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		for neighbor, weight := range graph[current.ID] {
			alt := distances[current.ID] + weight

			if alt < distances[neighbor] {
				distances[neighbor] = alt
				heap.Push(pq, &Node{ID: neighbor, Distance: alt})
			}
		}
	}

	return distances
}

func main() {
	graph := map[int]map[int]int{
		0: {1: 4, 7: 8},
		1: {0: 4, 2: 8, 7: 11},
		2: {1: 8, 3: 7, 8: 2, 5: 4},
		3: {2: 7, 4: 9, 5: 14},
		4: {3: 9, 5: 10},
		5: {2: 4, 3: 14, 4: 10, 6: 2},
		6: {5: 2, 7: 1, 8: 6},
		7: {0: 8, 1: 11, 6: 1, 8: 7},
		8: {2: 2, 6: 6, 7: 7},
	}

	distances := Dijkstra(graph, 0)
	fmt.Println("Shortest paths from node 0:")

	for node, distance := range distances {
		fmt.Printf("Node %d: %d\n", node, distance)
	}
}
