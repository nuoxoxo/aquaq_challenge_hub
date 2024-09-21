package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    _"strconv"
)

var Map [6]string

func main() {

    // got the map
    //for _, line := range Map { fmt.Println("Map/", line) }

    // get the steps
    URL := "https://challenges.aquaq.co.uk/challenge/3/input.txt"
    body := strings.TrimSpace(string(getbody(URL)))
    res := 0
    r, c := 0, 2
    N := len(body)
    for i := 0; i < N; i++ {
        char := string(body[i])
        if char == "U" && r - 1 >= 0 && string(Map[r - 1][c]) == "#" {r -= 1}
        if char == "D" && r + 1 <= 5 && string(Map[r + 1][c]) == "#" {r += 1}
        if char == "L" && c - 1 >= 0 && string(Map[r][c - 1]) == "#" {c -= 1}
        if char == "R" && c + 1 <= 5 && string(Map[r][c + 1]) == "#" {c += 1}
        res += r + c
    }
    fmt.Println("res/", res)
}

func init() {
    Map[0] = "  ##  "
    Map[1] = " #### "
    Map[2] = "######"
    Map[3] = "######"
    Map[4] = " #### "
    Map[5] = "  ##  "
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
