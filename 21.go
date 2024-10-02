package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

var tiles [][]int

func main() {
    //fmt.Println(len(tiles), len(tiles[0]))
    res := Dijkstra(tiles)
    fmt.Println(YELL("res/"), res)
    fmt.Println(YELL("res/"), DP(tiles))
    //fmt.Println(YELL("res/"), DP2(tiles))
    //res = BFS (tiles)
    //fmt.Println(YELL("res/"), res)
}

func DP (tiles [][]int) [2]int {
    tiles=[][]int{{3, 4, 5, 1, 3},{9, 3, 4, 0, 9},{4, 5, 4, 4, 7},{3, 7, 9, 8, 2},{2,1,1,1,1}}
    R, C := len(tiles), len(tiles[0])
    dp := make([][]int, R)
    for r := 0; r < R; r++ {
        dp[r] = make([]int, C - 4)
    }
    // 0-th row
    for c := 0; c < C - 4; c++ {
        for sub := c; sub < c + 5; sub++ {
            dp[0][c] += tiles[0][sub]
        }
    }
    for r := 1; r < R; r++ {
        if r < R {
            dp[r][0] = tiles[r][0] + tiles[r][1]
            dp[r][C - 1] = tiles[r][C - 2] + tiles[r][C - 1]
            for c := 1; c < C - 1; c++ {
                dp[r][c] = tiles[r][c - 1] + tiles[r][c] + tiles[r][c + 1]
            }
        } else  {
            for c := 0; c < C; c++ { dp[r][c] = 0 }
        }
    }
    fmt.Println(dp[0])
    for r := 1; r < R + 1; r++ {
        left_end := max2(dp[r - 1][0], dp[r - 1][1])
        right_end := max2 (dp[r - 1][C - 2], dp[r - 1][C - 1])
        dp[r][0] += left_end
        dp[r][C - 1] += right_end
        for c := 1; c < C - 1; c++ {
            mid := max3 (dp[r - 1][c - 1], dp[r - 1][c], dp[r - 1][c + 1])
            dp[r][c] += mid
        }
        fmt.Println("in/", r, dp[r])
    }
    res := 0
    for c := 0; c < C; c++ {
        if res < dp[R][c] { res = dp[R][c] }
    }
    res2 := 0
    for c := 0; c < C; c++ {
        if res2 < dp[R - 1][c] { res2 = dp[R - 1][c] }
    }
    return [2]int{res, res2}
}

func max2 (l, r int)int{
    if l > r {return l}
    return r
}

func max3 (l, m, r int)int{
    return max2(max2(l, m), r)
}

func BFS (tiles [][]int) int {

    res := 0
    // tiles=[][]int{{3, 4, 5, 1, 3},{9, 3, 4, 0, 9},{4, 5, 4, 4, 7},{3, 7, 9, 8, 2}}
    if len(tiles) == 0 || len(tiles[0]) < 3 { return 0 }
    H, W := len(tiles), len(tiles[0])
    q := []Tile{}
    for i := 1; i < W - 1; i++ {
        gain := tiles[0][i - 1] + tiles[0][i] + tiles[0][i + 1]
        q = append(q, Tile{Gain: gain, Coor: [2]int{0, i}})
    }
    for len(q) > 0 {
        level := q[0]
        fmt.Println("lvl/", level)
        q = q[1:]
        r, c := level.Coor[0], level.Coor[1]
        begin, end := c - 1, c + 1
        if begin == 0 { begin/*, end*/ = 1/*, 2*/ }
        if end == W - 1 { /*begin, */end = /*W - 3, */W - 2 }
        // assert begin < end
        if r == H - 1 {
            if res < level.Gain { res = level.Gain }
            break
        }
        for i := begin; i < end + 1; i++ {
            if r == H - 1 { continue }
            gain := tiles[r + 1][i - 1] + tiles[r + 1][i] + tiles[r + 1][i + 1]
            q = append(q, Tile{ Gain: level.Gain + gain, Coor: [2]int{r + 1, i} })
            //if res < level.Gain + gain { res = level.Gain + gain }
            //heap.Push(& q, & Tile { Gain: level.Gain + gain, Coor: [2]int{r + 1, i} })
        }
    }
    return res
}

func init() {
    // input
    URL := "https://challenges.aquaq.co.uk/challenge/21/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    for _, line := range lines {
        nums := strings.Split(line, " ")
        temp := []int{}
        for _, num := range nums {
            n, _ := strconv.Atoi(num)
            temp = append(temp, n)
        }
        tiles = append(tiles, temp)
    }
}

func getbody(URL string) []uint8 {
    session_value := os.Getenv("COOK")
    if session_value == "" {
        panic("session/empty")
    }
    conn := &http.Client {}
    req, _ := http.NewRequest("GET", URL, nil)
    session := &http.Cookie {
        Name:   "session",
        Value:  session_value,
    }
    req.AddCookie( session )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    return body
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }

