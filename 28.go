package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
)

func main() {
    URL := "https://challenges.aquaq.co.uk/challenge/28/input.txt"
    body := string(getbody(URL))
    // TEST
    //body = " ABCD \nA\\  /A\nB /\\ B\nC/ \\ C\nD/ / D\n ABCD \n "
    input := "FISSION_MAILED"//"DAD"
    lines := strings.Split(body, "\n")
    for i := range lines {fmt.Println(YELL("line/"), lines[i])}
    H := [][]string{}
    for i, line := range lines[:len(lines) - 1] {
        var temp []string
        for _, c := range line {
            temp = append(temp, string(c))
        }
        H = append(H, temp)
        fmt.Printf("%s %02d %v \n", CYAN("grid/"), i, H[i])
    }
    // prepare directions
    D := map[string][2]int{ "U": {-1, 0}, "D": {1, 0}, "L": {0, -1}, "R": {0, 1} }
    // soln
    R, C := len(H), len(H[0])
    res := ""
    for _, ch := range input {
        char := string(ch)
        d, r, c := "R", 0, 0
        hall := H
        for r < R {
            if char == hall[r][0] {
                break
            }
            r++
        }
        c++ // 1st step
        for true {
            if r == 0 || r == R - 1 || c == C - 1 || c == 0 {
                //assert hall[r][c] is alnum or underscore
                res += hall[r][c]
                break
            }
            // DBG
            // fmt.Println("running/", r, c, R, C, d)
            if hall[r][c] == "/" {
                hall[r][c] = "\\"
                if d == "R" {
                    d = "U"
                } else if d == "U" {
                    d = "R"
                } else if d == "L" {
                    d = "D"
                } else if d == "D" {
                    d = "L"
                }
            } else if hall[r][c] == "\\" {
                hall[r][c] = "/"
                if d == "R" {
                    d = "D"
                } else if d == "D" {
                    d = "R"
                } else if d == "L" {
                    d = "U"
                } else if d == "U" {
                    d = "L"
                }
            }
            r, c = r + D[d][0], c + D[d][1]
        }
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
