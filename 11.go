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

    URL := "https://challenges.aquaq.co.uk/challenge/11/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")[1:]

    // TEST
    //lines = []string{"lx,ly,ux,uy", "0,0,3,3", "2,2,4,5", "6,3,8,7",}

    //fmt.Println("lines/", lines[:4], len(lines))
    coors := make([][4]int, len(lines))
    cells := make(map[[2]int]int)
    seen := make(map[[2]int]bool)

    // 1st pass
    var i, igrow, j, jgrow int
    for idx, line := range lines {
        nums := strings.Split(line, ",")
        temp := [4]int{}
        i = 0
        for i < 4 {
            n, _ := strconv.Atoi(nums[i])
            temp[i] = n
            i++
        }
        coors[idx] = temp
        // update this area in cells
        igrow, jgrow = 1, 1 // i grows, otherwise shrinks
        if temp[0] > temp[2] { igrow = -1 }
        if temp[1] > temp[3] { jgrow = -1 }
        i = temp[0]
        for i != temp[2] {
            j = temp[1]
            for j != temp[3] {
                cells[[2]int{i, j}] += 1
                seen[[2]int{i, j}] = true
                j += jgrow
            }
            i += igrow
        }
    }

    // 2nd pass
    for _, coor := range coors {
        alone := true
        if coor[0] > coor[2] { igrow = -1 }
        if coor[1] > coor[3] { jgrow = -1 }
        i = coor[0]
        for i != coor[2] {
            j = coor[1]
            found := false
            for j != coor[3] {
                if cells [[2]int{i, j}] > 1 {
                    alone = false
                    found = true
                    break
                }
                j += jgrow
            }
            if found { break }
            i += igrow
        }
        if !alone { continue } // keep otherwise discard entire area
        i = coor[0]
        for i != coor[2] {
            j = coor[1]
            for j != coor[3] {
                seen[[2]int{i, j}] = false
                j += jgrow
            }
            i += igrow
        }
    }
    res := 0
    for _, ok := range seen {
        if ok { res++ }
    }
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

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }
