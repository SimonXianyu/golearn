package main
/**
Simple program read text file and hard wrap line into 50 characters per line.

 */
import (
    "flag"
    "bufio"
    "os"
    "fmt"
    "strings"
)

func main() {
    fname := flag.String("f", "", "target file name");
    tname := flag.String("t", "", "target file name");
    flag.Parse()

    if len(*fname ) == 0 || len(*tname) == 0 {
        fmt.Errorf("Missing filename \n")
        return
    }

    in, err := os.Open(*fname)
    if nil != err {
        fmt.Errorf("Failed to open file %s \n", fname)
        return
    }
    defer in.Close()

    sc := bufio.NewScanner(in)
    var lines []string
    var empty bool = false
    var enc []rune
    for sc.Scan() {
        l := sc.Text()
        strings.TrimSpace(l)
        if len(l) ==0 {
            if (!empty) {
                lines = append(lines, "\r\n\r\n")
            }
            empty = true
        } else {
            enc = []rune(l)
            if len(enc) >50 {
                var i int =0
                var end int = i+50
                for i<len(enc) {
                    end = i+50
                    if end > len(enc) {
                        end = len(enc)
                    }
                    l1 := string(enc[i:end])
                    lines = append(lines,l1)
                    i+=50
                }

            } else {
                lines = append(lines, l)
            }
        }
    }

    of, err1 := os.OpenFile(*tname, os.O_CREATE|os.O_WRONLY, 0644)
    if err1 != nil {
        fmt.Errorf("Failed to open output")
    }

    o := bufio.NewWriter(of)
    for _,vl := range(lines) {
        o.WriteString(vl)
        o.WriteString("\n")
    }
    o.Flush()
    of.Close()
}
