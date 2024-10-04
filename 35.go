package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
    "sort"
)

var cipher string
var wordlist []string
var demo_number int = 0

func main() {
    fmt.Println(YELL("wordlist/size"), len(wordlist))
    fmt.Println(YELL("cipher/"), cipher, len(cipher))

    demo("WE ARE DISCOVERED FLEE AT ONCE", "GLASS")
    demo("WE ARE DISCOVERED FLEE AT ONCE", "LEVER")
}
func demo(text, codeword string) {
    demo_number++
    fmt.Println("\n- Demo", demo_number, "-\n")
    // the chopped grid
    N := len(codeword)
    chopped := []string{}
    for i := 0; i < len(text) - 1; i += N {
        end := i + N
        if end > len(text) { end = len(text) }
        chopped = append(chopped, text[i : end])
        fmt.Println(CYAN(chopped[i / N]), "- idx/", i)
    }
    // get a selection order like 1 2 0 3 4
    order, tosort := make([]int, N), []string{}
    for _, c := range codeword { tosort = append(tosort, string(c)) }
    sort.Strings(tosort)
    fmt.Println(YELL("sorted/"), tosort)
    for pos, r := range codeword {
        for idx, l := range tosort {
            if l == string(r) {
                order[pos] = idx
                tosort[idx] = " "
                break
            }
        }
    }
    fmt.Println(YELL("order/ "), order)

    res := ""
    chopped2 := make([]string, len(order))
    for idx, col := range order {
        colword := ""
        for row := 0; row < len(chopped); row++ {
            colword += string(chopped[row][idx])
        }
        chopped2[col] = colword
        fmt.Println(CYAN("chopped/new"), colword, "-", col, idx)
    }
    for _, line := range chopped2 {
        res += line
    }
    fmt.Println(YELL("res/"), res)
    fmt.Println(CYAN("cmp/"), " DV  NWECEE E ODEOAIEFACRSRLTE", CYAN("-"), codeword)
}

func init() {
    // input - ciphertext
    URL := "https://challenges.aquaq.co.uk/challenge/35/input.txt"
    body := string(getbody(URL))
    // DBG
    //fmt.Println(CYAN("-1/"), string(body[len(body) - 1]))
    //fmt.Println(CYAN("-2/"), string(body[len(body) - 2]))
    cipher = body[ :len(body) - 2 ]
    // wordlist
    URL = "https://challenges.aquaq.co.uk/challenge/35"
    re := regexp.MustCompile(`(?s)ave a <a href="(.*?)">handy list`)
    sub := re.FindAllStringSubmatch(string(getbody(URL)), -1)[0][1]
    URL = "https://challenges.aquaq.co.uk" + sub
    wordlist = strings.Split(strings.TrimSpace(string(getbody(URL))), "\n")
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
