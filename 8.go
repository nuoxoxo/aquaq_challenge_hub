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

    URL := "https://challenges.aquaq.co.uk/challenge/8/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")[1:]
    fmt.Println("body/", lines[:2])

    buyin := [][5]int{}
    has_ym := make(map[string]bool)
    ym_keys := []string{}
    check_days_in_a_month := make(map[string]int)

    // Parse - get a list for each line - month-days sanity checker
    for _, line := range lines {
        item, datestring := parse_item(line)
        buyin = append(buyin, item)
        ym := datestring[:7]
        check_days_in_a_month[datestring[:7]]++
        if !has_ym[ym] {
            has_ym[ym] = true
            ym_keys = append(ym_keys, ym)
        }
    }
    for _, key := range ym_keys {
        fmt.Println(CYAN("checker/"), key, check_days_in_a_month[key])
        // sanity checker result: days are recorded continuously w/o missiong day
    }

    // Solve - as it turns out the dates(y-m-d) are redundant
    C := 0
    M := [5]int{0, 0, 0, 0, 0}
    for _, buy := range buyin {
        milk, cereal := buy[3], buy[4]
        C += cereal // cereal restocked before breakfast
        // breakfast
        for i := range M {
            if M[i] >= 100 && C >= 100 {
                M[i] -= 100
                C -= 100
                break
            }
        }
        for i := range M {
            if i == 4 {
                M[i] = 0
            } else {
                M[i] = M[i + 1]
            }
        }
        M [4] += milk // milk restocked after breakfast
    }
    res := C
    for i := range M { res += M[i] }
    fmt.Println(YELL("res/"), res)
}

func parse_item(line string) ([5]int, string) {
    trio := strings.Split(line, ",")
    date := strings.Split(trio[0], "-")
    y, _ := strconv.Atoi(date[0])
    m, _ := strconv.Atoi(date[1])
    d, _ := strconv.Atoi(date[2])
    milk, _ := strconv.Atoi(trio[1])
    cereal, _ := strconv.Atoi(trio[2])
    item := [5]int{y, m, d, milk, cereal}
    return item, trio[0] // return: item, "2015-02-12"
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
