package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

const testing bool = false

func main() {
    URL := "https://challenges.aquaq.co.uk/challenge/40/input.txt"
    line := string(getbody(URL))
    if testing {line = "0 1 2 4 6 8 9 8 6 4 2 3 5 6 5 4 5 7 8 6 4 2 1 0"}
    lines := strings.Split(strings.TrimSpace(line), " ")
    nums := []int{}
    for _, n := range lines {
        num, _ := strconv.Atoi(n)
        nums = append(nums, num)
    }
    // fmt.Println("nums/", nums, "- len/", len(nums)) // DBG
    N := len(nums)
    hipos, peaks, proms := []int{}, []int{}, []int{}
    highest := -1
    // get peaks' heights and positions
    for i, n := range nums {
        if i == 0 && n > nums[i + 1] ||
        i == N - 1 && n > nums[i - 1] || 
        0 < i && i < N && nums[i - 1] < n && n > nums[i + 1] {
            hipos, peaks = append(hipos, i), append(peaks, n)
            if highest < n { highest = n }
        }
    }
    //fmt.Println(hipos, "hipos/") // DBG
    //fmt.Println(peaks, "peaks/") // DBG
    // get peaks' prominences
    size := len(hipos)
    for i := 0; i < size; i++ {        
        idx, curr := hipos[i], peaks[i]
        if curr == highest {
            proms = append(proms, highest)
            continue
        }
        lmin, rmin := highest + 1, highest + 1
        assert( idx < N - 1 && idx > 0 )
        for j := idx - 1; j > -1 && nums[j] < curr; j-- {
            if lmin > nums[j] { lmin = nums[j] }
        }
        for j := idx + 1; j < N && nums[j] < curr; j++ {
            if rmin > nums[j] { rmin = nums[j] }
        }
        realmin := lmin
        if realmin < rmin { realmin = rmin }
        if realmin < highest { proms = append(proms, curr - realmin) }
    }
    //fmt.Println(proms, "proms/") // DBG
    res := 0
    for _, n := range proms {res += n}
    fmt.Println("res/", res)
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

