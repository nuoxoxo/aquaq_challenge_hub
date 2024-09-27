package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    _"reflect"
)

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/32/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    ///*
    lines = []string{
        "()",
        "([]{})",
        "(a[b[]]c){}",
        ")()",
        "([a)]",
        "]{}[",
        "((a)){]",
    }//*/

    for _, line := range lines { fmt.Println(CYAN("line/"), len(line), YELL(line)) }
    fmt.Println(YELL("size/"), len(lines))

    res := 0
    left, right := "([{", ")]}"
    dict := make(map[string]string)
    dict[")"], dict["]"], dict["}"] = "(", "[", "{"
    for _, line := range lines {
        deque := []string{}
        found := true
        for _, c := range line {
            char := string(c)
            if strings.Contains(left, char) {
                deque = append(deque, char)
            } else if strings.Contains(right, char) {
                if len(deque) == 0 {
                    found = false
                    break
                }
                last := deque[ len(deque) - 1 ]
                if last != dict[char] {
                    found = false
                    break
                } else {
                    deque = deque[ : len(deque) - 1]
                }
            }
        }
        if found && len(deque) == 0 { res++ }
    }
    fmt.Println(YELL("res/"), res)
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
