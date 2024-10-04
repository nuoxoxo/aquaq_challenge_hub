package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "regexp"
    "time"
)

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/25/input.txt"
    line := strings.TrimSpace(string(getbody(URL)))
    timestamps := strings.Split(line, "\n")
    format := "15:04:05.000"

    diffs := [][]time.Duration{}
    temp := []time.Duration{}
    for i := 1; i < len(timestamps); i++ {
        prev, err := time.Parse(format, timestamps[i - 1])
        if err != nil {
            continue
        }
        curr, err := time.Parse(format, timestamps[i])
        if err != nil {
            diffs = append(diffs, temp)
            temp = []time.Duration{  }
            continue
        }
        diff := curr.Sub(prev)
        temp = append(temp, diff)
    }

    URL = "https://challenges.aquaq.co.uk/challenge/25"
    re := regexp.MustCompile(`(?s)<pre>(.*?)</pre></div>`)
    line = re.FindAllStringSubmatch( string(getbody(URL)), -1 )[0][1]
    lines := strings.Split(strings.TrimSpace(line), "<br>")
    morse := make(map[string]string)
    for _, line := range lines {
        pair := strings.Split(line, "| ")
        morse[pair[1]] = pair[0]
    }
    test_cipher := ".--- .- --   -.. --- -. ..- -"
    fmt.Println("test/", YELL(dcode( morse, test_cipher )))

    // soln
    res := ""
    for _, diff := range diffs {
        unit_size := diff[0]
        morsechar := ""
        for i := 1; i < len(diff); i++ { if unit_size > diff[i] { unit_size = diff[i] }}
        for i, duration := range diff {
            size := int((duration / unit_size))// / time.Nanosecond)
            if i % 2 == 0 { // btw on and off
                if size == 3 { morsechar += "-" }
                if size == 1 { morsechar += "." }
            } else if size == 3 || size == 7 { // btw off and the next on
                res += morse[morsechar]
                morsechar = ""
                if size == 7 { res += " " }
            }
        }
        if morsechar != "" { res += morse[morsechar] }
        res += "\n"
        fmt.Println(CYAN("res/add"), res)
    }
    fmt.Println(YELL("res/"), res)
}

func dcode (morse map[string]string, cipher string) string {
    res := ""
    slc := strings.Split(strings.Replace(cipher, "   ", "  ", -1), " ")
    for _, key := range slc {
        if key == "" {
            res += " "
        } else {
            res += morse[key]
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
