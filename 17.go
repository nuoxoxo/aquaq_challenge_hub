package main

import (
    "fmt"
    "time"
    "reflect"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

const format string = "2006-01-02"
const tostring string = "20060102"
const nogo, goal = -1, 1 // no goal, goal, no record

func main() {
    // parsing
    URL := "https://challenges.aquaq.co.uk/challenge/17/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")[1:]
    dates := []time.Time{}
    hasdate := make(map[time.Time]bool)
    nations := []string{}
    hasnation:= make(map[string]bool)
    records := make(map[string]map[time.Time]int)
    for _, line := range lines {
        sub := strings.Split(line, ",")
        d, l, r, lp, rp := sub[0], sub[1], sub[2], sub[3], sub[4]
        date, _ := time.Parse(format, d)
        if !hasdate[date] {
            hasdate[date] = true
            dates = append(dates, date)
        }
        if !hasnation[l] {
            hasnation[l] = true
            nations = append(nations, l)
        }
        if !hasnation[r] {
            hasnation[r] = true
            nations = append(nations, r)
        }
        lvalue, _ := strconv.Atoi(lp)
        rvalue, _ := strconv.Atoi(rp)
        if _, ok := records[l]; !ok { records[l] = make(map[time.Time]int) }
        if _, ok := records[r]; !ok { records[r] = make(map[time.Time]int) }
        if lvalue == 0 {records[l][date] = nogo} else {records[l][date] = goal}
        if rvalue == 0 {records[r][date] = nogo} else {records[r][date] = goal}
    }
    // soln - loop through nations then dates
    res := 0
    team := ""
    lend, rend := "", ""
    for _, nation := range nations {
        diff := 0
        begin := -1
        for i, date := range dates {
            record := records[nation][date] // has goal: > 0 - no goal: < 0
            if record < 0 && begin == -1 {
                begin = i
            } else if record > 0 && begin != -1 {
                diff = int(date.Sub( dates[begin] ).Hours() / 24)
                if res < diff {
                    lend, rend = dates[begin].Format(tostring), date.Format(tostring)
                    fmt.Println(CYAN("upd/"), nation, lend, rend, YELL("dif/"), diff)
                    team = nation
                    res = diff
                }
                begin = -1
            }
        }
    }
    fmt.Println(YELL("res/"), team, lend, rend, CYAN("diff/"), res)
}

func init() {

    // TEST - time module sanity test

    l := "1900-01-03"
    r := "1902-01-01"
    ll, _ := time.Parse(format, l)
    rr, _ := time.Parse(format, r)
    fmt.Println(ll)
    fmt.Println(rr, "type/rr -", reflect.TypeOf(ll))
    hours := rr.Sub(ll)
    div := hours.Hours() / 24
    diff := int(div)

    fmt.Println("hours/", hours, reflect.TypeOf(hours))
    fmt.Println("hours/", hours.Hours(), reflect.TypeOf(hours.Hours()))
    fmt.Println("div/", div, reflect.TypeOf(div))
    fmt.Println("diff/", diff)
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

