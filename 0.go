package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "os"
    "strings"
    "strconv"
)

func main() {

    // failed - missing cookie session value
    /*
    resp, _ := http.Get("https://challenges.aquaq.co.uk/challenge/0/input.txt")
    defer resp.Body.Close()
    */

    // switch to using http.NewRequest
    URL := "https://challenges.aquaq.co.uk/challenge/0/input.txt"
    session_value := os.Getenv("COOK")
    if session_value == "" {
        panic("session/empty")
    }
    conn := & http.Client{}
    req, _ := http.NewRequest("GET", URL, nil)
    session := & http.Cookie {
        Name:   "session",
        Value:  session_value,
    }
    req.AddCookie( session )
    resp, _ := conn.Do(req)
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)
    data := strings.Split(string(body), "\n")

    // numpad setup
    np := make(map[int]string)
    np[2] = "abc"
    np[3] = "def"
    np[4] = "ghi"
    np[5] = "jkl"
    np[6] = "mno"
    np[7] = "pqrs"
    np[8] = "tuv"
    np[9] = "wxyz"
    np[0] = " "

    // solve
    res := ""
    for _, line := range data {
        if line == "" {
            continue
        }
        pair := strings.Split(line, " ")
        l, r := pair[0], pair[1]
        num, _ := strconv.Atoi(l)
        press, _ := strconv.Atoi(r)
        res += string(np[num][press - 1])
    }
    fmt.Println("res/", CYAN(res))
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }

