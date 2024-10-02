package main
// 86009x - 86205
import (
    "fmt"
    "container/heap"
)

// min heap for levels

type Tile struct { // ie. HeapNode
    Gain    int
    Coor    [2]int
}

type MinHeap [] *Tile

func (q MinHeap) Len() int { return len(q) }
func (q MinHeap) Less(L, R int) bool { return q[L].Gain > q[R].Gain }
func (q MinHeap) Swap(L, R int) { q[L], q[R] = q[R], q[L] }
func (q *MinHeap) Push(x interface{}) { *q = append(*q, x.(*Tile)) }

func (q *MinHeap) Pop() interface{} {
    cp := *q
    size := len(cp)
    last := cp[size - 1]
    *q = cp[:size - 1]
    return last
}

func Dijkstra (tiles[][]int) int {

//    tiles=[][]int{{3, 4, 5, 1, 3},{9, 3, 4, 0, 9},{4, 5, 4, 4, 7},{3, 7, 9, 8, 2}}
    if len(tiles) == 0 || len(tiles[0]) < 3 { return 0 }
    H, W := len(tiles), len(tiles[0])
    q := make(MinHeap, 0)
    heap.Init(& q)
    // src - 0th row --- des - last row
    for i := 1; i < W - 1; i++ {
        gain := tiles[0][i - 1] + tiles[0][i] + tiles[0][i + 1]
        heap.Push(& q, & Tile { Gain: gain, Coor: [2]int{0, i} })
    }
    res := 0
    for q.Len() > 0 {
        level := heap.Pop(& q).(*Tile)
        //fmt.Println("lvl/", level)
        r, c := level.Coor[0], level.Coor[1]
        if r == H - 1 {
            res = level.Gain
            break
        }
        begin, end := c - 1, c + 1
        if begin == 0 && r != H - 2 { begin/*, end*/ = 1/*, 2*/ }
        if end == W - 1 && r != H - 2 { /*begin, */end = /*W - 3, */W - 2 }
        // assert begin < end
        //fmt.Println(begin, end)
        for i := begin; i < end + 1; i++ {
            gain := tiles[r + 1][i - 1] + tiles[r + 1][i] + tiles[r + 1][i + 1]
            heap.Push(& q, & Tile { Gain: level.Gain + gain, Coor: [2]int{r + 1, i} })
        }
    }
    for _, qq := range q { if qq.Gain > 86000 {fmt.Println(qq, H) }}
    return res
}



