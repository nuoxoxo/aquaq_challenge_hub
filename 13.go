package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {
    URL := "https://challenges.aquaq.co.uk/challenge/13/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    //for _, line := range lines {fmt.Println("line/", CYAN(line[:17]), "...") }

    lines = []string{ "ABCABCABCABCABC", "AAAAAAB", "ABCABCABCABCABCAAAAAAB" }

    var max_repetion_bruteforce func(s string) (int, string)
    max_repetion_bruteforce = func(s string) (int, string) {
        pattern := ""
        count := 1
        N := len(s)
        // DBG
        // fmt.Println( CYAN ("\n" + s))
        for i := range s {
            size := 1
            for size < N - i + 1 {
                curr := 0
                substr := s[i : i + size]
                // DBG
                // fmt.Println( YELL ("\n" + substr))
                for i + size * curr < N + 1 && i + size * (curr + 1) < N + 1 &&
                    s[i + size * curr : i + size * (curr + 1)] == substr {
                    // DBG
                    //fmt.Println(s[i + size * curr : i + size * (curr + 1)], substr, curr)
                    curr++
                }
                if count < curr {
                    count = curr
                    pattern = substr
                }
                size++
            }
        }
        return count, pattern
    }
    res := 0
    for _, line := range lines {
        size, pattern := max_repetion_bruteforce (line)
        fmt.Println(size, YELL (pattern))
        res += size
    }
    fmt.Println("res/" + Cyan, res, Rest)
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
