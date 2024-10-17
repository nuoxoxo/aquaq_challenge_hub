package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "regexp"
    "strings"
    "sort"
    "sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var cipher string
var wordlist []string
var demo_number int = 0

func main() {

    //wg.Wait()

    fmt.Println(CYAN("wordlist/size"), len(wordlist))
    fmt.Println(CYAN("cipher/"), cipher, len(cipher))

    src := "WE ARE DISCOVERED FLEE AT ONCE"

    // demo
    cipher_from_GLASS := demo(src, "GLASS")
    fmt.Println(YELL("\ndemo/ends"), cipher_from_GLASS)

    cipher_from_LEVER := demo(src, "LEVER")
    fmt.Println(YELL("\ndemo/ends"), cipher_from_LEVER)

    // reversed_demo func test
    arr, res := reversed_demo(cipher_from_GLASS, "GLASS")
    fmt.Println(YELL("\nsoln/ends"), CYAN(res), arr)
    assert (res == src)

    arr, res = reversed_demo(cipher_from_LEVER, "LEVER")//, wordlist)
    assert (res == src)

    // soln
    end := 99996
    codeword := wordlist[ end ]
    arr, res = reversed_demo(cipher, codeword)
    fmt.Println(YELL("\nsoln/ends"), CYAN(res), arr)

    // real soln
    record_word := ""
    record_accu := -1.0
    record_indi := -1
    for idx, codeword := range wordlist {
        arr, res = reversed_real (cipher, codeword)
        // check pertinence of decoded text against the give wordlist
        N := float64(len(arr))
        count_correct := 0
        for _, word := range arr {
            l, r := 0, len(wordlist) - 1
            for l <= r {
                mid := (l + r) / 2
                tar := wordlist[mid]
                // Compare function
                //  > 0 we are smaller ie. we need to go lower
                //  == 0 matches
                //  < 0 we are bigger ie. go higher/right
                cmp := strings.Compare(tar, word)
                if cmp == 0 {
                    count_correct++
                    break
                } else if cmp < 0 {
                    l = mid + 1
                } else {
                    r = mid - 1
                }
            }
        }
        accuracy := float64(count_correct) / N
        if record_accu < accuracy {
            record_word = codeword
            record_accu = accuracy
            record_indi = idx
        }
        // printer/DBG
        //fmt.Printf("%d - %.2f - %s \n", idx, accuracy, codeword)
    }
    //fmt.Println("found/", record_word, "- rate/", record_accu * 100)
    fmt.Printf("found/ %.2f - codeword/ %s - @ index/ %d \n",
        record_accu * 100, YELL(record_word), record_indi)
}

func reversed_real(ciphertext, codeword string/*, wl []string*/) ([]string, string) {
    demo_number++
    mutex.Lock()
    mutex.Unlock()
    // ------ redo ------------------------------------------
    N := len(codeword)
    order, sorted := make([]int, N), []string{}
    for _, c := range codeword { sorted = append(sorted, string(c)) }
    sort.Strings(sorted)
    for pos, r := range codeword {
        for idx, l := range sorted {
            if l == string(r) {
                order[pos] = idx
                sorted[idx] = " " // emptying the sorted char list
                break
            }
        }
    }
    cols := len(ciphertext) / len(order)
    temp := ""
    reslist, res := []string{}, ""
    for c := 0; c < cols; c++ {
        for _, pos := range order {
            char := ciphertext[pos * cols + c] // algo
            if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
                temp += string (char)
            } else if len(temp) > 0 {
                reslist = append (reslist, strings.ToLower(temp))
                res += temp + string(char)
                if string(char) != " " { res += " " }
                temp = ""
            }
        }
    }
    if len(temp) > 0 { // last word grapped
        reslist = append (reslist, strings.ToLower(temp))
        res += temp + " "
        temp = "" //obsolete
    }
    return reslist, res[:len(res) - 1]
}

func reversed_demo(ciphertext, codeword string/*, wl []string*/) ([]string, string) {
    demo_number++
    mutex.Lock()
    fmt.Println("\n- Demo", demo_number, "-")
    fmt.Print("\n- Code : ", YELL(codeword), " -\n\n")
    mutex.Unlock()
    // ------ redo ------------------------------------------
    N := len(codeword)
    order, sorted := make([]int, N), []string{}
    for _, c := range codeword { sorted = append(sorted, string(c)) }
    sort.Strings(sorted)
    fmt.Println(CYAN("sorted/"), sorted)
    //for _, char := range sorted {fmt.Print(char, " ")}
    //fmt.Println()
    for pos, r := range codeword {
        for idx, l := range sorted {
            if l == string(r) {
                order[pos] = idx
                sorted[idx] = " " // emptying the sorted char list
                break
            }
        }
    }
    fmt.Println(CYAN("order/ "), order)
    fmt.Println(CYAN("sorted/emptied "), sorted)

    cols := len(ciphertext) / len(order)
    temp := ""
    reslist, res := []string{}, ""
    for c := 0; c < cols; c++ {
        for _, pos := range order {
            char := ciphertext[pos * cols + c] // algo
            if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
                temp += string (char)
            } else if len(temp) > 0 {
                reslist = append (reslist, strings.ToLower(temp))
                res += temp + string(char)
                if string(char) != " " { res += " " }
                temp = ""
            }
        }
    }
    if len(temp) > 0 { // last word grapped
        reslist = append (reslist, strings.ToLower(temp))
        res += temp + " "
        temp = "" //obsolete
    }
    return reslist, res[:len(res) - 1]
}

func index(arr[]int, val int)int{
    for i, v := range arr { if v == val { return i } }
    return -1
}

func demo(text, codeword string)string{
    demo_number++
    fmt.Println("\n- Demo", demo_number, "-\n")

    // from codeword, get a selection order eg. 1 2 0 3 4
    N := len(codeword)
    order, sorted := make([]int, N), []string{}
    for _, c := range codeword { sorted = append(sorted, string(c)) }
    sort.Strings(sorted)
    fmt.Println(CYAN("sorted/"), sorted)
    for pos, r := range codeword {
        for idx, l := range sorted {
            if l == string(r) {
                order[pos] = idx
                sorted[idx] = " " // emptying the sorted char list
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
    return res + "\n"
}

func init() {
    /*
    wg.Add(1)
    go func() {
        defer wg.Done()
    */
    // input - ciphertext
    URL := "https://challenges.aquaq.co.uk/challenge/35/input.txt"
    body := string(getbody(URL))
    // DBG
    //fmt.Println(YELL("-1/dbg"), string(body[len(body) - 1]))
    //fmt.Println(YELL("-2/dbg"), string(body[len(body) - 2]))
    cipher = body[ :len(body) - 1 ]
    // wordlist
    URL = "https://challenges.aquaq.co.uk/challenge/35"
    re := regexp.MustCompile(`(?s)ave a <a href="(.*?)">handy list`)
    sub := re.FindAllStringSubmatch(string(getbody(URL)), -1)[0][1]
    URL = "https://challenges.aquaq.co.uk" + sub
    words := string(getbody(URL))
    words = words[:len(words) - 1]
    wordlist = strings.Split(strings.TrimSpace(words), string(rune(13)) + "\n")

    //}()
    // 👆 --- for wg - go func
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
