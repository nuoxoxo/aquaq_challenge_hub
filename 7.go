package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
    "math"
    "sort"
)

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/7/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")[1:]
    //fmt.Println("lines/", lines) // DBG

    var ELO_DEFAULT float64 = 1200.0
    vs := make( map[string]map[string]bool ) // a dict of dicts
    elo := make( map[string]float64 )
    for _, line := range lines {
        // parse
        parts := strings.Split(line, ",")
        points := strings.Split(parts[2], "-")
        a, b := parts[0], parts[1]
        if _, ok := vs[a]; ! ok {
            vs[a] = make(map[string]bool)
            elo[a] = ELO_DEFAULT
        }
        if _, ok := vs[b]; ! ok {
            vs[b] = make(map[string]bool)
            elo[b] = ELO_DEFAULT
        }    
        ap, _ := strconv.Atoi(points[0])
        bp, _ := strconv.Atoi(points[1])
        player_a_wins := false
        if ap > bp {
            player_a_wins = true
        }
        vs[a][b] = player_a_wins
        vs[b][a] = !player_a_wins
        Rank_a := elo[a]
        Rank_b := elo[b]
        // solve
        Expected_a := expected_win_rate(Rank_a, Rank_b)
        elo[a] = update_score(Expected_a, Rank_a, player_a_wins)
        Expected_b := expected_win_rate(Rank_b, Rank_a)
        elo[b] = update_score(Expected_b, Rank_b, !player_a_wins)
    }
    A := []int{}
    for _, v := range elo {
        A = append(A, int(v))
    }
    sort.Ints(A)
    //fmt.Println("sort/", A) // DBG
    fmt.Println("diff/", A[len(A) - 1] - A[0])
}

// Elo:
//  expected win rate:  Ea = 1 / (1 + 10 ^ ((Rb - Ra) / 400))
//  updated ranking:    Ri' = Ri + 20 * (1 - Ei)
//      note/ the '1' should be 'has_won': 1 on win, 0 on lose

func expected_win_rate(Rank_a, Rank_b float64) float64 {
    return 1.0 / (1 + math.Pow(10, (Rank_b - Rank_a) / 400))
}

func update_score(Expected, Rank float64, has_won bool) float64 {
    curr := 0
    if has_won { curr = 1 }
    return Rank + 20 * (float64(curr) - Expected)
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
