package main

// 3009x - 2960x

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "regexp"
    "strconv"
)

const N int = 5
var G [N][N][2]int
var Coors map[int][2]int
var numseq [][]int

func main() {
    res := 0
    for _, sq := range numseq {
        g := G // this already deepcopies
        for i := 0; i < len(sq); i++ {
            key := sq[i]
            if _, has := Coors[key]; !has {
                continue
            }
            r, c := Coors[key][0], Coors[key][1]
            g[r][c][1] = 1
            if wins (g) {
                res += i + 1
                break
            }
        }
    }
    fmt.Println("res/" + Yell, res, Rest)
}

func wins (g [N][N][2]int) bool {

    // rows and cols
    for r := 0; r < N; r++ {
        rows, cols := 0, 0 // win counts
        for c := 0; c < N; c++ {
            if g[r][c][1] == 1 { rows++ }
            if g[c][r][1] == 1 { cols++ }
        }
        if rows == N || cols == N { return true }
    }
    // diagonal
    da, db := 0, 0
    for i := 0; i < N; i++ {
        if g[i][i][1] == 1 { da++ }
        if g[N - i - 1][i][1] == 1 { db++ }
    }
    if da == N || db == N { return true }
    return false
}

func init() {

    // Step - board

    URL := "https://challenges.aquaq.co.uk/challenge/14"
    body := getbody(URL)
    re := regexp.MustCompile(`(?s)<pre>(.*?)</pre>`)
    lines := strings.Split(re.FindAllStringSubmatch(string(body), -1)[0][1], "<br>")[1:]


    Coors = make(map[int][2]int)
    r := 0
    for _, line := range lines {
        subs := strings.Split(line, " ")
        c := 0
        for _, sub := range subs {
            if sub == "" || sub == " " {
                continue
            }
            n, _ := strconv.Atoi(sub)
            G[r][c] = [2]int{n, 0}
            Coors[n] = [2]int{r, c}
            c++
        }
        r++
    }

    URL = "https://challenges.aquaq.co.uk/challenge/14/input.txt"
    lines = strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    for _, line := range lines {
        sub := strings.Split(line, " ")
        temp := []int{}
        for _, num := range sub {
            n, _ := strconv.Atoi(num)
            temp = append(temp, n)
        }
        numseq = append(numseq, temp)
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
