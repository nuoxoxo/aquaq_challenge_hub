package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/27/input.txt"
    lines := strings.Split(string(getbody(URL)), "\n")
    lines = lines[:len(lines) - 1]
    /*
    lines = []string{
        "                roulette            ",
        "                e      l            ",
        "                v      e            ",
        "                e      c            ",
        "                netulg t            ",
        "    invalidly        n i            ",
        "            a        i o            ",
        "            c        y n            ",
        "            h        r sharpness    ",
        "            t        r              ",
        "            i        u              ",
        "            n        c              ",
        "            grumpiness              ",
    }
    */
    R, C := len(lines), len(lines[0])
    //for _, line := range lines { fmt.Println(CYAN("line/"), line, len(line)) }

    res := 0
    for r := 0; r < R; r ++ { // all horizontal words
        left := -1
        for c := 0; c < C; c++ {
            char := lines[r][c]
            if char != ' ' { // char
                if left == -1 {
                    left = c
                }
            } else if char == ' ' && left != -1 {
                temp := calc( lines[r][left : c] )
                if temp > 0 {
                    res += temp
                }
                left = -1
            }
        }
    }
    for c := 0; c < C; c ++ { // verticle words
        top := -1
        for r := 0; r < R; r++ {
            char := lines[r][c]
            if char != ' ' && top == -1 {
                top = r
            } else if (char == ' ' || r == R - 1) && top != -1 {
                substr := ""
                for i := top; i < r; i++ {
                    substr += string(lines[i][c])
                }
                if r == R - 1 { substr += string(char) }
                temp := calc(substr)
                if temp > 0 {
                    res += temp
                }
                top = -1
            }
        }
    }
    fmt.Println(CYAN("res/"), res)
}

func calc(s string) int {
    if len(s) < 2 { return -42 }
    res := 0
    for _, c := range s { res += int(c - 'a') + 1 }
    return res * len(s)
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
