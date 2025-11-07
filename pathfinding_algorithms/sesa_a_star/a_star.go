package main

import (
	"container/heap"
	"fmt"
	"math"
)

type AStarNode struct {
	ID        int
	Cost      int
	Distance  int
	Heuristic int
}

type AStarPriorityQueue []*AStarNode

func (pq AStarPriorityQueue) Len() int {
	return len(pq)
}

func (pq AStarPriorityQueue) Less(i, j int) bool {
	return pq[i].Heuristic+pq[i].Distance < pq[j].Heuristic+pq[j].Distance
}

func (pq AStarPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *AStarPriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*AStarNode))
}

func (pq *AStarPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func heuristic(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func AStar(graph map[int]map[int]int, start, goal int) map[int]int {
	distances := make(map[int]int)
	for node := range graph {
		distances[node] = 1<<31 - 1 // infinite
	}
	distances[start] = 0

	pq := &AStarPriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &AStarNode{ID: start, Distance: 0, Heuristic: heuristic(start, goal)})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*AStarNode)
		if current.ID == goal {
			break
		}

		for neighbor, weight := range graph[current.ID] {
			alt := distances[current.ID] + weight

			if alt < distances[neighbor] {
				distances[neighbor] = alt
				est := heuristic(neighbor, goal)
				heap.Push(pq, &AStarNode{ID: neighbor, Distance: alt, Heuristic: est})
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

	start, goal := 0, 4
	distances := AStar(graph, start, goal)
	fmt.Printf("Shortest path from node %d to node %d:\n", start, goal)

	for node, distance := range distances {
		fmt.Printf("Node %d: %d\n", node, distance)
	}
}
