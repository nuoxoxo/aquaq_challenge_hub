package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
    _"reflect"
)

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/12/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")

    ops := [][]int{}
    for _, line := range lines {
        pair := strings.Split(line, " ")
        d, _ := strconv.Atoi(pair[0])
        n, _ := strconv.Atoi(pair[1])
        ops = append(ops, []int{d, n})
    }

    var sim func([][]int) int
    sim = func(ops [][]int) int {
        res := 1
        lvl := 0
        dir := 1
        for lvl < len(ops) {
            op := ops[lvl]
            d, n := op[0], op[1]
            if d == 0 {
                dir *= -1
            }
            lvl += dir * n
            res++
        }
        return res
    }

    fmt.Println("res/", sim(ops))

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
