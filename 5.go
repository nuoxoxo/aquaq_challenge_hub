package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
)

// object

type Die struct {
    faces map[string]int
}

func New_Die (initial_state map[string]int) *Die {
    return & Die { faces: initial_state }
}

func (d *Die) Faces() map[string]int {
    return d.faces
}

func (d *Die) Show (idx *int) {
    if idx != nil {
        fmt.Println(*idx, d.faces)
    } else {
        fmt.Println(d.faces)
    }
}

func (d *Die) Roll(op string) {
    if op == "L" { d.Roll_Left() }
    if op == "R" { d.Roll_Right() }
    if op == "U" { d.Roll_Up() }
    if op == "D" { d.Roll_Down() }
}

func (d *Die) Roll_Up(){
    up := d.faces["front"]
    front := d.faces["down"]
    d.faces["up"] = up
    d.faces["down"] = 7 - up
    d.faces["front"] = front
    d.faces["back"] = 7 - front
}


func (d *Die) Roll_Down(){
    up := d.faces["back"]
    front := d.faces["up"]
    d.faces["up"] = up
    d.faces["down"] = 7 - up
    d.faces["front"] = front
    d.faces["back"] = 7 - front
}

func (d *Die) Roll_Left(){
    left := d.faces["front"]
    front := d.faces["right"]
    d.faces["left"] = left
    d.faces["right"] = 7 - left
    d.faces["front"] = front
    d.faces["back"] = 7 - front
}

func (d *Die) Roll_Right(){
    left := d.faces["back"]
    front := d.faces["left"]
    d.faces["left"] = left
    d.faces["right"] = 7 - left
    d.faces["front"] = front
    d.faces["back"] = 7 - front
}

// solve

func main() {

    URL := "https://challenges.aquaq.co.uk/challenge/5/input.txt"
    ops := strings.TrimSpace(string(getbody(URL)))

    //ops = "LRDLU" // TEST

    da := New_Die(map[string]int{ "up": 3, "front": 1, "left": 2,
        "down": 7 - 3, "back": 7 - 1, "right": 7 - 2,
    })

    db := New_Die(map[string]int{ "up": 2, "front": 1, "left": 3,
        "down": 7 - 2, "back": 7 - 1, "right": 7 - 3,
    })

    da.Show(nil)
    db.Show(nil)
    res := 0
    matches := []int{}
    for i, op := range ops {
        da.Roll(string(op))
        db.Roll(string(op))
        //da.Show(&i)
        //db.Show(&i)
        if da.Faces()["front"] == db.Faces()["front"] {
            matches = append(matches, i)
            res += i
        }
    }

    // fmt.Println("matches/", matches) // DBG
    fmt.Println("res/" + Yell, res, Rest)
    da.Show(nil)
    db.Show(nil)
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

