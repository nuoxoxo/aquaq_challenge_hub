package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
)

var demo_number int = 1
const N int = 5 // grid must be 5*5 due to there being 26 alphabets

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/23/input.txt"
    line := strings.TrimSpace(string(getbody(URL)))
    fmt.Println(YELL("line/"), line)
    playfair_demo ("playfair", "tree")
    playfair_demo ("playfair", "flawless")
    reversed_playfair("playfair", "pabapgxyxy")
    reversed_playfair("power plant", line)
}

func reversed_playfair(keyword, todo string) string {

    grid := make_playfair_grid(keyword)
    bgm := []string{}
    fmt.Println(YELL("reversed_playfair/grid"), len(grid), len(grid[0]))
    fmt.Println(CYAN("todo/"), todo)
    for i := 0; i < len(todo) - 1; i += 2 {bgm = append(bgm, todo[i:i + 2])}
    fmt.Println(CYAN("bi-grams/"), bgm)
    for i, pair := range bgm {
        l, r := pair[0], pair[1]
        for _, line := range grid { // check rows
            L, R := -1, -1
            for col, char := range line {
                if byte(char) == l { L = col }
                if byte(char) == r { R = col }
            }
            if L != -1 && R != -1 {
                L, R = (L - 1 + len(line)) % len(line), (R - 1 + len(line)) % len(line)
                bgm[i] = string(line[L]) + string(line[R])
                fmt.Println(YELL("rows/"), bgm, YELL(bgm[i]))
                break
            }
        }
        for col := 0; col < len(grid[0]); col++ { // check cols
            U, D := -1, -1
            for row := 0; row < len(grid); row++ {
                if grid[row][col] == l { U = row }
                if grid[row][col] == r { D = row }
            }
            if U != -1 && D != -1 {
                U = (U - 1 + len(grid[0])) % len(grid[0]) 
                D = (D - 1 + len(grid[0])) % len(grid[0])
                bgm[i] = string(grid[U][col]) + string(grid[D][col])
                fmt.Println(YELL("cols/"), bgm, YELL(bgm[i]))
                break
            }
        }
        afound, bfound, A, B := false, false, [2]int{}, [2]int{}
        for row := 0; row < len(grid); row++ { // diagonals
            for col := 0; col < len(grid[0]); col++ {
                if grid[row][col] == l { A, afound = [2]int{row, col}, true }
                if grid[row][col] == r { B, bfound = [2]int{row, col}, true }
            }
            if A[0] == B[0] || A[1] == B[1] { continue }
            if afound && bfound {
                bgm[i] = string(grid[A[0]][B[1]]) + string(grid[B[0]][A[1]])
                fmt.Println(YELL("diag/"), bgm, YELL(bgm[i]))
                break
            }
        }
        fmt.Println(CYAN("curr/"), bgm)
    }
    fmt.Println("res/bi-grams", bgm)
    res, alt := "", ""
    for _, pair := range bgm {
        alt += pair
        l, r := pair[0], pair[1]
        if string(l) != "x" { res += string(l) }
        if string(r) != "x" { res += string(r) }
    }
    fmt.Println(YELL("alt/string"), alt)
    fmt.Println(YELL("res/string"), res)
    return res
}

func playfair_demo(keyword, todo string) {

    grid := make_playfair_grid( keyword )
    fmt.Println(CYAN("todo/original"), todo)
    for i := 1; i < len(todo); i++ {
        if todo[i - 1] == todo[i] { todo = todo[:i] + "x" + todo[i:] }
    }
    if len(todo) % 2 == 1 { todo += "x" }
    fmt.Println(CYAN("todo/prepared"), todo)
    bgm := []string{}
    for i := 0; i < len(todo) - 1; i += 2 { bgm = append(bgm, todo[i:i + 2]) }
    fmt.Println(CYAN("bi-grams/"), bgm)
    // do the todo
    for i, pair := range bgm {
        l, r := pair[0], pair[1]
        for _, line := range grid { // check rows
            L, R := -1, -1
            for col, char := range line {
                if byte(char) == l { L = col }
                if byte(char) == r { R = col }
            }
            if L != -1 && R != -1 {
                L, R = (L + 1) % len(line), (R + 1) % len(line)
                bgm[i] = string(line[L]) + string(line[R])
                fmt.Println(YELL("rows/"), bgm, YELL(bgm[i]))
                break
            }
        }
        for col := 0; col < len(grid[0]); col++ { // check cols
            U, D := -1, -1
            for row := 0; row < len(grid); row++ {
                if grid[row][col] == l { U = row }
                if grid[row][col] == r { D = row }
            }
            if U != -1 && D != -1 {
                U, D = (U + 1) % len(grid[0]), (D + 1) % len(grid[0])
                bgm[i] = string(grid[U][col]) + string(grid[D][col])
                fmt.Println(YELL("cols/"), bgm, YELL(bgm[i]))
                break
            }
        }
        afound, bfound, A, B := false, false, [2]int{}, [2]int{}
        for row := 0; row < len(grid); row++ { // diagonals
            for col := 0; col < len(grid[0]); col++ {
                if grid[row][col] == l { A, afound = [2]int{row, col}, true }
                if grid[row][col] == r { B, bfound = [2]int{row, col}, true }
            }
            if A[0] == B[0] || A[1] == B[1] { continue }
            if afound && bfound {
                bgm[i] = string(grid[A[0]][B[1]]) + string(grid[B[0]][A[1]])
                fmt.Println(YELL("diag/"), bgm, YELL(bgm[i]))
                break
            }
        }
        fmt.Println(CYAN("curr/"), bgm)
    }
}

func make_playfair_grid(keyword string) []string {

    for i, char := range keyword {
        if string(char) == " " { keyword = keyword[:i] + keyword[i + 1:] }
    }
    demo_number++
    fmt.Println("\n- Demo", demo_number, "-\n")
    used := make(map[byte]bool)
    flat := ""
    for i := range keyword {
        if used[keyword[i]] { continue }
        used[keyword[i]] = true
        flat += string(keyword[i])
    }
    for i := byte('a'); i <= byte('z'); i++ {
        if used[i] || string(i) == "j" { continue } // omit j
        used[i] = true
        flat += string(i)
    }
    fmt.Println(CYAN("flat/"), flat)
    fmt.Println(CYAN("grid/"), len(flat))
    grid := []string{}
    for i := 0; i < len(flat) - 4; i += 5 {grid = append(grid, flat[i:i + 5])}
    for i, g := range grid { fmt.Println(g, "-", i) }
    return grid
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
