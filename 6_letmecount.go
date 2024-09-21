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

    URL := "https://challenges.aquaq.co.uk/challenge/6/input.txt"
    body := strings.TrimSpace(string(getbody(URL)))
    fmt.Println("body/", body)

    words := strings.Split(body, " ")
    last := words[len(words) - 1]
    num, _ := strconv.Atoi(last)
    fmt.Println("num/", num)

    res := 0
    for l := 0; l <= num; l++ {
        for m := 0; m <= num - l; m++ {
            r := num - l - m
            res += f(strconv.Itoa(l)) + f(strconv.Itoa(m)) + f(strconv.Itoa(r))
        }
    }
    fmt.Println("res/", res)
}

func f(s string) int {
    res := 0
    for _, char := range s {
        if char == '1' { res++ }
    }
    return res
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
