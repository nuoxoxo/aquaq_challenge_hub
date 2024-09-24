package main

import (
    "fmt"
    "container/heap"
)

var dists map[string]int

func (g *Graph) Dijkstra(src, des string) int {

    pq := make(MinHeap, 0)
    heap.Init( & pq)

    // distances
    dists = make(map[string]int)
    for v := range g.Vertices {
        dists[v] = 2147583648
    }
    dists[src] = 0

    heap.Push(&pq, & HeapNode{ Vertex: g.Vertices[src] })

    // dijkstra proper
    for pq.Len() > 0 {
        node := heap.Pop( & pq).(*HeapNode)
        if node.Vertex.Key == des {
            break
        }
        for nei, cost := range node.Vertex.Costs {
            if dists[nei.Key] > cost + dists[ node.Vertex.Key ] {
                // search for min cost
                dists[nei.Key] = cost + dists[ node.Vertex.Key ]
                heap.Push( & pq, & HeapNode{ Vertex: nei })
            }
        }
    }

    return dists[des]
}

// min heap - priority Q

type HeapNode struct {
    Vertex  * Vertex
    Index   int
}

type MinHeap [] * HeapNode

func (pq MinHeap) Len() int { return len(pq) }
func (pq MinHeap) Less(L, R int) bool {
    return dists[ pq[L].Vertex.Key ] < dists[ pq[R].Vertex.Key ]
}

func (pq MinHeap) Swap(L, R int) {
    pq[L], pq[R] = pq[R], pq[L]
    pq[L].Index, pq[R].Index = L, R
}

func (pq * MinHeap) Push(x interface{}) {
    end := len(* pq)
    node := x.(* HeapNode)
    node.Index = end
    *pq = append(*pq, node)
}

func (pq * MinHeap) Pop() interface{} {
    prev := *pq
    N := len(prev)
    node := prev[N - 1]
    prev[N - 1] = nil // in case of leak
    node.Index = -1
    *pq = prev[0 : N - 1]

    return node
}

// graph - adj list

type Vertex struct {
    Key     string
    Costs   map[* Vertex]int // destination, cost
}

type Graph struct {
    Vertices map[string] * Vertex
}

func (G * Graph) Printer() {
    for _, v := range G.Vertices {
        fmt.Println("\nfrom/", YELL(v.Key))
        for nei, cost := range v.Costs {
            fmt.Println("  to/", CYAN(nei.Key), cost)
        }
    }
}

// color def. after main

