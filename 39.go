package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
    "strconv"
)

var nums, test []int

func main() {
    //fmt.Println(nums[:42], YELL("\n/nums - size/"), len(nums))
    //fmt.Println(test, CYAN("\n/test - size/"), len(test))
    res := soln(test)
    fmt.Println(YELL("test/res"), res, "\n")
    res = soln(nums)
    fmt.Println(YELL("nums/res"), res)
}

func soln(A []int) int {
    who_begins := 0 // switched btw 0 1
    who_plays := 0 // switched btw 0 1
    pts_pair := [2]int{0, 0}
    A_won := 0
    darts_final := 0
    darts_played := 0 // count to 3 and switch player
    for _, n := range A {
        pts_pair[ who_plays ] += n
        if pts_pair[ who_plays ] == 501 {
            if who_plays == 0 { A_won++ }
            pts_pair = [2]int{0, 0}
            who_begins = 1 - who_begins
            who_plays = who_begins
            darts_played = 0
            darts_final += n
        } else if darts_played == 2 {
            who_plays = 1 - who_plays
            darts_played = 0
        } else {
            darts_played++
        }
    }
    res := A_won * darts_final
    fmt.Println(CYAN("won/"), A_won)
    fmt.Println(CYAN("fin/"), darts_final)
    return res
}

func init() {
    URL := "https://challenges.aquaq.co.uk/challenge/39/input.txt"
    nums = getnums(string(getbody(URL)))
    re := regexp.MustCompile(`(?s)<pre>(.*?)</pre>`)
    URL = "https://challenges.aquaq.co.uk/challenge/39"
    test = getnums(re.FindAllStringSubmatch(string(getbody(URL)), -1)[0][1])
}

func getnums(s string) []int {
    res := []int{}
    for _, s := range strings.Split(strings.TrimSpace(s), " ") {
        n, _ := strconv.Atoi(s)
        res = append(res, n)
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

func assert(expression bool) {
    if ! expression {
        fmt.Print(YELL("assert/false "))
        panic(expression)
    }
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }

