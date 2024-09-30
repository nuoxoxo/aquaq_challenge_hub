package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

var hands []string

func main() {
    res := 0 // games won
    score := 0
    aces := 0
    for _, hand := range hands {
        face, err := strconv.Atoi(hand)
        if err == nil {
            score += face
        } else {
            if hand == "J" || hand == "K" || hand == "Q" {
                score += 10
            } else if hand == "A" {
                score += 1
                aces ++
            }
        }
        if score < 21 {
            t := 1
            for t < aces + 1 {
                if score + 10 * t == 21 {
                    score, aces, res = 0, 0, res + 1
                    break
                }
                t++
            }
        } else {
            if score == 21 { res++ }
            score, aces = 0, 0
        }
    }
    fmt.Println(YELL("res/"), res)
}

func init() {
    // input
    URL := "https://challenges.aquaq.co.uk/challenge/20/input.txt"
    hands = strings.Split(strings.TrimSpace(string(getbody(URL))), " ")
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
