package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

type Tile struct { // ie. HeapNode
    Gain    int
    Coor    [2]int
}

var tiles [][]int

func main() {
    fmt.Println(YELL("res/"), DP(tiles))
}

func DP (tiles [][]int) int {
    //tiles=[][]int{{3, 4, 5, 1, 3},{9, 3, 4, 0, 9},{4, 5, 4, 4, 7},{3, 7, 9, 8, 2},{2,1,1,1,1}}
    R, C := len(tiles), len(tiles[0]) - 5
    dp := make([][]int, R)
    for r := 0; r < R; r++ {
        dp[r] = make([]int, C + 1)
    }
    // 0-th row
    for c := 0; c <= C; c++ {
        for sub := c; sub < c + 5; sub++ {
            dp[0][c] += tiles[0][sub]
        }
    }
    for r := 1; r < R; r++ {
        for c := 0; c <= C; c++ {
            curr := dp[r - 1][c]
            if c > 0 { curr = max2(curr, dp[r - 1][c - 1]) }
            if c < C { curr = max2(curr, dp[r - 1][c + 1]) }
            todo := 0
            for sub := c; sub < c + 5; sub++ {
                todo += tiles[r][sub]
            }
            dp[r][c] = curr + todo
        }
    }
    res := 0
    for c := 0; c <= C; c++ {
        if res < dp[R - 1][c] { res = dp[R - 1][c] }
    }
    return res
}

func max2 (l, r int)int{
    if l > r {return l}
    return r
}

func max3 (l, m, r int)int{
    return max2(max2(l, m), r)
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

