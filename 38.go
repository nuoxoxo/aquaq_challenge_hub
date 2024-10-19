package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

var lines []string

func main() {
    //for i, line := range lines { fmt.Println("i/", i, "-", line) }
    res := 0
    for _, line := range lines {
        nums := []int{}
        for _, n := range strings.Split(line, " ") {
            num, _ := strconv.Atoi(n)
            nums = append(nums, num)
        }
        N := len(nums)
        var l, r int
        l = 0
        for l < N {
            r = 1
            for r < N {
                L := 0
                if l > r { L = l - r }
                R := l + 1
                if l + r + 1 > N { R = N - r }
                found := false
                for L < R {
                    sum := 0
                    for _, num := range nums[L : L + r + 1] { sum += num }
                    if sum % (r + 1) == 0 {
                        found = true
                        res++
                        break
                    }
                    L++
                }
                if !found { break }
                r++
            }
            l++
        }
        res += N
    }
    fmt.Println("res/", res)
}

func init() {
    URL := "https://challenges.aquaq.co.uk/challenge/38/input.txt"
    body := strings.TrimSpace(string(getbody(URL)))
    lines = strings.Split(body, "\n")
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
        fmt.Print(YELL("assert/false "))
        panic(expression)
    }
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }

