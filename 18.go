package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
    _"reflect"
    _"regexp"
    _"strconv"
)

const format string = "15:04:05"

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/18/input.txt"
    src, des := "", ""
    var res time.Duration
    for _, line := range strings.Split(strings.TrimSpace(string(getbody(URL))), "\n") {
        curr, _ := time.Parse(format, line)
        back, fore := curr, curr
        for true {
            if ispld( fore.Format(format) ) { // forward --->
                diff := fore.Sub(curr)
                res += diff
                break
            } else if ispld( back.Format(format) ) { // <--- backward
                diff := curr.Sub(back)
                res += diff
                break
            }
            back, fore = back.Add( - time.Second), fore.Add(time.Second)
        }
    }
    fmt.Println(YELL("res/"), res.Seconds(), src, des)
    
}

func ispld (s string) bool {
    for l, r := 0, len(s) - 1; l < r; l, r = l + 1, r - 1 {
        if s[l] != s[r] { return false }
    }
    return true
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
