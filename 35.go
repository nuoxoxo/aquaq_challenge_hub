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
    fmt.Println(CYAN("wordlist/size"), len(wordlist))
    fmt.Println(CYAN("cipher/"), cipher, len(cipher))
    s := ""
    var p float64
    s = columnar_transposition_demo("WE ARE DISCOVERED FLEE AT ONCE", "GLASS")
    fmt.Println(YELL("end/"), s)
    s = columnar_transposition_demo("WE ARE DISCOVERED FLEE AT ONCE", "LEVER")
    fmt.Println(YELL("end/"), s)

    // 
    s, p = reversed(cipher, "GLASS", wordlist)
    fmt.Println(YELL("end/rev"), s, p)
    s, p = reversed(cipher, "LEVER", wordlist)
    fmt.Println(YELL("end/rev"), s, p)

    /*
    END := 99996

    for i := END; i < END + 1; i++ {
        //fmt.Println("in/", i, word, precision)
        t, p := reversed (cipher, wordlist[i], wordlist)
        if precision < p {
            precision = p
            fulltext = t
        }
        fmt.Println("i =", i , "word =", wordlist[i])
        fmt.Printf("\tprecision = %0.2f - p = %0.8f \n", precision, p)
        //fmt.Println("in/", i, word, precision)
    }
    fmt.Println(fulltext, precision, "/res")
    *//*
    fulltext, precision = reversed (" DV  NWECEE E ODEOAIEFACRSRLTE", "GLASS", wordlist)
    fmt.Println(fulltext, precision, "/res")
    */
}

func reversed(ciphertext, codeword string, wl []string) (string, float64) {

    demo_number++
    fmt.Println("\n- Demo", demo_number, "-\n")

    return "", 0.0
}

func columnar_transposition_demo(text, codeword string)string{
    demo_number++
    fmt.Println("\n- Demo", demo_number, "-\n")

    // from codeword, get a selection order eg. 1 2 0 3 4
    N := len(codeword)
    order, tosort := make([]int, N), []string{}
    for _, c := range codeword { tosort = append(tosort, string(c)) }
    sort.Strings(tosort)
    fmt.Println(CYAN("sorted/"), tosort)
    for pos, r := range codeword {
        for idx, l := range tosort {
            if l == string(r) {
                order[pos] = idx
                tosort[idx] = " "
                break
            }
        }
    }
    fmt.Println(CYAN("order/ "), order)

    // from original text, get the chopped grid
    chopped := []string{}
    for i := 0; i < len(text) - 1; i += N {
        end := i + N
        if end > len(text) { end = len(text) }
        chopped = append(chopped, text[i : end])
        fmt.Println(YELL(chopped[i / N]), "- idx/", i)
    }

    // from chopped grid & order get a new chopped grid + its string
    chopped2 := make([]string, len(order))
    for idx, col := range order {
        colword := ""
        for row := 0; row < len(chopped); row++ {
            colword += string(chopped[row][idx])
        }
        chopped2[col] = colword
        fmt.Println(YELL("chopped/new"), colword, "-", col, idx)
    }
    res := strings.Join( chopped2, "" )
    fmt.Println(CYAN("res/"), res, YELL("-"), codeword)

    // debugger/assert
    if codeword == "GLASS" {
        cmp := " DV  NWECEE E ODEOAIEFACRSRLTE"
        fmt.Println(CYAN("cmp/"), cmp)
        assert (res == cmp)
    }
    return res
}

func init() {
    // input - ciphertext
    URL := "https://challenges.aquaq.co.uk/challenge/35/input.txt"
    body := string(getbody(URL))
    // DBG
    //fmt.Println(YELL("-1/"), string(body[len(body) - 1]))
    //fmt.Println(YELL("-2/"), string(body[len(body) - 2]))
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
        fmt.Print(YELL("assert/false "))
        panic(expression)
    }
}

const Yell, Cyan, Rest string = "\033[33m", "\033[36m", "\033[0m"
func YELL(s string)string{ return Yell + s + Rest }
func CYAN(s string)string{ return Cyan + s + Rest }
