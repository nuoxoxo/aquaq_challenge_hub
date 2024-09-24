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

    URL := "https://challenges.aquaq.co.uk/challenge/10/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")[1:]

    adj := & Graph {
        Vertices: make(map[string]*Vertex),
    }

    for _, line := range lines {
        trio := strings.Split(line, ",")
        src, des := trio[0], trio[1]
        cost, _ := strconv.Atoi(trio[2])
        // add node/vertex - both vertices "A" "B"
        if ok := adj.Vertices[src]; ok == nil {
            node := & Vertex {
                Key:    src,
                Costs:  make(map[*Vertex]int),
            }
            adj.Vertices[src] = node
        }
        if ok := adj.Vertices[des]; ok == nil {
            node := & Vertex {
                Key:    des,
                Costs:  make(map[*Vertex]int),
            }
            adj.Vertices[des] = node
        }
        // add edge/cost
        adj.Vertices[src].Costs[adj.Vertices[des]] = cost
    }
    adj.Printer()
    
}

// priority Q



// graph - adj list

type Vertex struct {
    Key     string
    Costs   map[*Vertex]int // destination, cost
}

type Graph struct {
    Vertices map[string]*Vertex
}

func (G * Graph) Printer() {
    for _, v := range G.Vertices {
        fmt.Println("\nfrom/", YELL(v.Key))
        for nei, cost := range v.Costs {
            fmt.Println("  to/", CYAN(nei.Key), cost)
        }
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

