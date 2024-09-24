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

    for _, row := range G { fmt.Println(CYAN ("grid/"), len(row), row) }
    for _, sq := range numseq {fmt.Println("nseq/" + Yell, sq[:12], Rest, len(sq))}
    for k, v := range Coors { fmt.Println(k, v) }

    // TEST
    //numseq = [][]int{[]int{10,5,21,45,53,70,66,4},[]int{10,5,21,45,53,70,66,4},[]int{10,5,21,45,53,70,66,4},[]int{10,5,21,45,53,70,66,4}}

    // bruteforce
    res := 0
    for idx, sq := range numseq {
        temp := len(sq)
        // fmt.Println(sq, len(sq))
        for left := 0; left < len(sq) - N + 1; left++ {
            //var g [N][N][2]int
            //g = G
            g := G
            for right := left; right < len(sq); right++ {
                //fmt.Println(left, right)
                // fmt.Printf("G/ %p - g/ %p \n", &G, &g) // proof of copying
                key := sq[right]
                if _, has := Coors[key]; !has { continue }
                r, c := Coors[key][0], Coors[key][1]
                g[r][c][1] = 1
                dist := right - left + 1
                //fmt.Println(CYAN("win/"), wins (g))
                // if left == 0 && right == 6 { fmt.Println(g, key, Coors[key], r, c) }
                if dist >= N && wins (g) && temp > dist {
                    temp = dist
                    fmt.Println(idx, "- winning sub/", sq[left : right + 1], "l - r/", left, right)
                    for _, gg := range g {fmt.Println(gg)}
                    break
                }
            }
        }
        res += temp
    }
    fmt.Println("res/", res)
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

    // DBG
    // fmt.Println(CYAN ("lines/"), lines, len(lines))

    Coors = make(map[int][2]int)
    r := 0
    for _, line := range lines {
        subs := strings.Split(line, " ")
        // DBG
        fmt.Println(YELL ("subs/"), subs, len(subs))
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
    // DBG
    // for _, row := range G { fmt.Println(CYAN ("row/"), len(row), row) }

    // Step - number sequences

    URL = "https://challenges.aquaq.co.uk/challenge/14/input.txt"
    lines = strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    // DBG
    // for _, line := range lines {fmt.Println("line/", CYAN (line[:21]), "...")}
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
