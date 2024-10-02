package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
)

var N int

func main() {
    URL := "https://challenges.aquaq.co.uk/challenge/33/input.txt"
    N, _ = strconv.Atoi(strings.TrimSpace(string(getbody(URL))))
    fmt.Println("target/", N)

    // set up scores
    scores := []int{}
    for n := 1; n < 21; n++ {
        scores = append(scores, n)
    }
    for n := 21; n < 61; n++ {
        if 20<n && n<41 && n%2==0 || 20<n && n<61 && n%3==0 || n == 25 || n == 50{
            scores = append(scores, n)
        }
    }
    
    // DP
    INF := 2147483647
    dp := make([]int, N + 1)
    dp[0] = 0
    for n := 1; n <= N; n++ { dp[n] = INF }
    for n := 1; n <= N; n++ {
        for _, score := range scores {
            diff := n - score
            if diff > -1 && dp[n] > dp[ diff ] + 1 {
                dp[n] = dp[diff] + 1
            }
        }
    }
    res := 0
    for n := 1; n <= N; n++ { res += dp[n] }
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
