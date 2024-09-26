package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
)

func main() {

    // wordlist
    URL := "https://challenges.aquaq.co.uk/challenge/15"
    re := regexp.MustCompile(`(?s)list <a href="(.*?)">here</a>`)
    txt := re.FindAllStringSubmatch(string(getbody(URL)), -1)[0][1]
    URL = "https://challenges.aquaq.co.uk" + txt
    wordlist := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    W := map[string]bool{}
    for _, word := range wordlist { 
        W[strings.TrimSpace(word)] = true
    }

    // input - word pairs
    URL = "https://challenges.aquaq.co.uk/challenge/15/input.txt"
    lines := strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
    pairs := [][2]string{}
    for _, line := range lines {
        pair := strings.Split(line, ",")
        pairs = append(pairs, [2]string{pair[0], pair[1]})
    }

    // TEST
    //pairs = [][2] string {[2]string{"fly", "try"}, [2]string{"try", "fly"}, [2]string{"word", "maze"}}

    // bi-directional BFS
    // assert len(set(wordlist)) == len(wordlist)
    var BI_BFS func() [][]string

    BI_BFS = func() [][]string {
        chains := [][]string{}
        for _, pair := range pairs {
            src, des := pair[0], pair[1]
            lq := [][]string{ []string{src} }
            rq := [][]string{ []string{des} }
            lchains := map[string][] string { src: []string{src} }
            rchains := map[string][] string { des: []string{des} }
            found := false

            for !found && len(lq) > 0 && len(rq) > 0 {
                // swapping - proceed the shorted end first
                if len(lq) > len(rq) {
                    lq, rq = rq, lq
                    lchains, rchains = rchains, lchains
                }

                // bfs popping
                N := len(lq)
                for i := 0; i < N; i++ {
                    lchain := lq[0]
                    lq = lq[1:]
                    curr := lchain[len(lchain) - 1]
                    // gen all possible words
                    possible := []string{}
                    for idx := 0; idx < len(curr); idx++ {
                        for c := 'a'; c <= 'z'; c++ {
                            if c == rune(curr[idx]) {continue}
                            nxt := curr[:idx] + string(c) + curr[idx + 1:]
                            if W[nxt] {
                                possible = append(possible, nxt)
                            }
                        }
                    }
                    for _, w := range possible {
                        if rchain, haskey := rchains[ w ]; haskey {
                            // reverse R chain
                            for l, r := 0, len(rchain) - 1; l < r; l, r = l + 1, r - 1 {
                                rchain[l], rchain[r] = rchain[r], rchain[l]
                            }
                            chains = append(chains, append( lchain, rchain ... ))
                            found = true
                            break
                        }
                        if _, haskey := lchains[ w ]; !haskey {
                            todo := append(append([]string{}, lchain ...), w)
                            lchains[w] = todo
                            lq = append(lq, todo)
                        }
                    }
                    if found { break }
                }
            }
        }
        return chains /// empty
    }
    chains := BI_BFS()
    res := 1
    for _, chain := range chains {
        fmt.Println("chain/", chain)
        res *= len(chain)
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
