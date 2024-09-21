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

    // get a []int{}
    URL := "https://challenges.aquaq.co.uk/challenge/2/input.txt"
    body := string(getbody(URL))
    res := []int{}
    A := []int{}
    for _, num := range strings.Split(body[:len(body) - 1], " ") {
        n, err := strconv.Atoi(num)
        if err != nil {
            fmt.Println("err/", err)
        }
        A = append(A, n)
    }
    N := len(A)
    L := 0

    // algo
    for L < N {
        R := L + 1
        for R < N && A[L] != A[R] {R++}
        if R == N {
            res = append(res, A[L])
            L++
        } else {
            L = R
        }
        fmt.Println("res/", res)
    }
    fmt.Println("res/", res)
    sum := 0
    for _, n := range res { sum += n }
    fmt.Println("sum/", sum)
}

// DRAFT
//  1 4 3 2 4 7 2 6 3 6
//    4     4
//        2     2
//                6   6
//  1 4   2       6 ---> which is incorrect
//
//  1 4 3 2 4 7 2 6 3 6
//    4     4
//  1 4       7 2 6 3 6
//                6   6
//  1 4 7 2 6 ------> correct - which means no overlapping is concerned

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


