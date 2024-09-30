package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "strconv"
)

const debugging bool = false
var iter_side_pairs [][2]int
var coor_pairs [][][2]int
var D [4][2]int = [4][2]int{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}

func main() {
    //iter_side_pairs, coor_pairs := [][2]int{{350,6}}, [][][2]int{{{2,2}, {2,3}}}
    // 👆 TEST
    var gol func([2]int, [][2]int)int = func(pair[2]int, coors[][2]int)int{
        iter, N := pair[0], pair[1]
        if debugging {
            fmt.Println("gol/", iter, N, coors)
        }
        cache := [][2]int{}
        for _, coor := range coors { cache = append(cache, coor) }
        for it := 0; it < iter; it++ {
            to_cache := [][2]int{}
            nei_of_life := make(map[[2]int]int)
            for _, cell := range cache {
                r, c := cell[0], cell[1]
                for i := 0; i < 4; i++ {
                    rr, cc := r + D[i][0], c + D[i][1]
                    if rr < 0 || rr >= N || cc < 0 || cc >= N { continue }
                    nei_of_life[ [2]int{rr, cc} ]++
                }
            }
            for nei, n := range nei_of_life {
                if n % 2 != 0 { to_cache = append(to_cache, nei) }
            }
            cache = to_cache
        }
        fmt.Println(CYAN("res/iter"), len(cache), "from", iter, N, coors)
        if debugging {
            fmt.Println("cache/", cache)
        }
        return len(cache)
    }
    res := 0
    for i, slc := range iter_side_pairs { res += gol( slc, coor_pairs[i] ) }
    fmt.Println(CYAN("res/"), res)
}

func init() {

    // input
    URL := "https://challenges.aquaq.co.uk/challenge/19/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    // for _, line := range lines { fmt.Println("line/", line) }
    for idx, line := range lines {
        parts := strings.Split(line, " ")
        iter, _ := strconv.Atoi(parts[0])
        side, _ := strconv.Atoi(parts[1])
        temp := [][2]int{}
        for i := 2; i < len(parts); i += 2 {
            r, _ := strconv.Atoi( parts[i] )
            c, _ := strconv.Atoi( parts[i + 1] )
            temp = append(temp, [2]int{r, c})
        }
        coor_pairs = append(coor_pairs, temp)
        iter_side_pairs = append(iter_side_pairs, [2]int{iter, side})
        if debugging {
            fmt.Println(idx, "-", temp)
        }
    }
    if debugging {
        fmt.Println(len(iter_side_pairs), len(coor_pairs))
    }
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
