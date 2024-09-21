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

    URL := "https://challenges.aquaq.co.uk/challenge/22/input.txt"
    nums := []int{}
    roms := []string{}

    // romanize
    for _, num := range strings.Split(strings.TrimSpace(string(getbody(URL))), " "){
        n, _ := strconv.Atoi(num)
        nums = append(nums, n)
        roms = append(roms, romanize(n))
    }
    fmt.Println("nums/", nums, "\nlen/", len(nums))
    fmt.Println("roms/", roms, "\nlen/", len(roms))

    // caesar
    res := 0
    for _, rom := range roms {
        for _, char := range rom {
            res += int(char - 'A' + 1)
        }
    }
    fmt.Println("res/", res)
}

func romanize(n int) string {
    vn := [13]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
    vs := [13]string{"I","IV","V","IX","X","XL","L","XC","C","CD","D","CM","M"}
    i := 12
    res := ""
    for true {
        tmp := n / vn[i]
        n %= vn[i]
        for tmp > 0 {
            res += vs[i]
            tmp--
        }
        i--
        if n < 1 {
            break
        }
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

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }
