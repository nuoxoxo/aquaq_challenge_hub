package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/29/input.txt"
    body := strings.TrimSpace(string(getbody(URL)))
    fmt.Println("body/", body)

    end, _ := strconv.Atoi(body)
    fmt.Println("end/", end)

    var isgood func(int) bool
    isgood = func(n int) bool {
        s := strconv.Itoa(n)
        i := 0
        for i < len(s) - 1 {
            if s[i] > s[i + 1] { return false }
            i++
        }
        return true
    }

    res := 0
    n := 0
    for n <= end {
        if isgood(n) { res++ }
        n++
    }
    fmt.Println("res/", res)
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
