package main
// 237786x - 242142x
import (
    "fmt"
    "os"
    "io/ioutil"
    "net/http"
    "strings"
    "regexp"
)

var A map[string] []string // char:ascii graphic
var D map[string] [][2]int // char:left- and right- (not incl.) most # index
var wall [6]string
var input string

func main() {
    fmt.Println(input)
    jettison := 0
    for i, c := range input {
        key := string(c)
        for _, line:= range A[key] { fmt.Println("ln/", YELL(line) ) }
        gap := 42
        if i > 0 {
            prevkey := string( input[i - 1] )
            for j, _ := range wall {
                L, R := D[prevkey][j][1], D[key][j][0]
                if gap > R + len(A[prevkey][j]) - L {
                    gap = R + len(A[prevkey][j]) - L
                }
            }
        }
        for j, _ := range wall { wall[j] += A[key][j] }
        fmt.Println(YELL("gap/"), gap)
        if i == 0 { continue }
        if gap > 1 {
            jettison += (gap - 1) * 6
        } else if gap == 0 {
            jettison -= 6
        } else if gap != 1 { panic(YELL("never here/")) }
        fmt.Println("gap/", gap, "jet/", jettison)
    }
    // print out entire graphic
    //for _, line := range wall { fmt.Println(line, len(line)) }
    fmt.Println("all/", count_sp())
    fmt.Println("jet/", jettison)
    fmt.Println("res/", count_sp() - jettison)
}

func count_sp() int {
    res := 0
    for i := range wall {
        for j := range wall[i] {
            if string(wall[i][j]) == " " { res++ }
        }
    }
    return res
}

func init() {

    // input
    URL := "https://challenges.aquaq.co.uk/challenge/16/input.txt"
    input = strings.TrimSpace(string(getbody(URL)))

    // wall - playground
    for i := 0; i < 6; i++ { wall[i] = "" }

    // alphabet
    URL = "https://challenges.aquaq.co.uk/challenge/16"
    re := regexp.MustCompile(`(?s)found <a href="(.*?)">here</a>`)
    txt := re.FindAllStringSubmatch( string(getbody(URL)), -1 )[0][1]
    URL = "https://challenges.aquaq.co.uk" + txt
    lines := strings.Split(string(getbody(URL)), "\n")
    A = make(map[string][]string)
    D = make(map[string][][2]int)
    letter := 'A'
    temp := []string{}
    tempd := [][2]int{}
    for i, line := range lines {
        N := len(line)
        l, r := 0, N - 1
        for l < N && line[l] != '#' { l++ }
        for r > -1 && line[r] != '#' { r-- }
        if (i + 1) % 6 == 0 {
            key := string(letter)

            // add to A
            temp = append(temp, line)
            A[key] = temp
            temp = []string{}

            // add to D
            tempd = append (tempd, [2]int{l, r + 1})
            D[key] = tempd
            tempd = [][2]int{}
            letter++
        } else {
            temp = append(temp, line)
            tempd = append(tempd, [2]int{l, r + 1})
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
