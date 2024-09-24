package main

import (

    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    _"reflect"
    "strings"
)

func main() {
    URL := "https://challenges.aquaq.co.uk/challenge/1/input.txt"
    body := string( getbody(URL) )
    fmt.Println("body/", YELL(strings.TrimSpace(body)))

    // sim
    
    // body = "kdb4life" // TEST

    HexChars := "0123456789abcdef"
    cipher := ""
    for i := 0; i < len(body); i++ {
        char := string(body[i])
        if ! strings.Contains(HexChars, char) {
            cipher += "0"
        } else if char != " " {
            cipher += char
        }
    }

    need := 3 - len(cipher) % 3
    //fmt.Println("cipher/", cipher, "\nlen/", len(cipher), "\nneed/", need)

    // padding
    for i := 0; i < need; i++ {
        cipher += "0"
    }
    partlen := len(cipher) / 3

    //fmt.Println("partlen/", partlen)

    res := make([]uint8, partlen * 2)
    for i := 0; i < 2; i++ {
        res[0 + i] = cipher[partlen * 0 + i]
        res[2 + i] = cipher[partlen * 1 + i]
        res[4 + i] = cipher[partlen * 2 + i]
    }
    // 0 2 4 - i=0
    // 1 3 5 - i=1

    fmt.Println("res/", CYAN(string(res)))

    //fmt.Println("len/cipher", len(cipher))
    //fmt.Println("cmp/", "0d40fe", "TEST/")
}

/*
  Set the string's non-hexadecimal characters to 0.
  Pad the string length to the next multiple of 3.
  Split the result into 3 equal sections.
  The first two digits of each remaining section are the hex components.
*/

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

