package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "sort"
    "strconv"
)

func main() {
    URL := "https://challenges.aquaq.co.uk/challenge/26/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    fmt.Println("lines/top", CYAN(lines[0]))
    fmt.Println("lines/btm", CYAN(lines[len(lines) - 1]))
    // TEST
    //lines = []string{ "1423", "121", "10290" }
    res := 0
    for _, line := range lines {
        N := len(line)
        if N == 0 { continue }
        i := N - 2
        for i > -1 {
            if line[i] < line[i + 1] {
                break
            }
            i--
        }
        if i == -1 { continue } // descending entirely
        j, smallestlarger := -1, '9' + 1
        for right := N - 1; right > i; right-- {
            if smallestlarger > rune(line[right]) && line[right] > line[i] {
                smallestlarger = rune(line[right])
                j = right
            }
        }
        // grab the original int from substr
        prev, _ := strconv.Atoi(line[i:])
        // swap [i, j] of line
        fmt.Println(line, i, j)
        runes := []rune(line)
        runes[i], runes[j] = runes[j], runes[i]
        line = string(runes)
        fmt.Println(line, i, j)
        // sort the right part of i
        sub := line[i + 1:]
        runes = []rune(sub)
        sort.Slice(runes, func(l ,r int) bool {
            return runes[l] < runes[r]
        })
        sub = string(runes)
        line = line[:i + 1] + sub
        // grab the modified int
        curr, _ := strconv.Atoi(line[i:])
        res += curr - prev
        size := N - i
        fmt.Println(YELL(line), i, j, CYAN("- size/"), size)
        fmt.Println("diff/", curr, "-", prev, "=", curr - prev, "\n")
        assert(size <= 10)
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

func assert(expression bool) {
    if ! expression {
        fmt.Print(CYAN("assert/false "))
        panic(expression)
    }
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }
