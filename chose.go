package main

/**
  尝试用动态规划思路写的从 数字序列中找出最少大于目标值的组合。
  用法：
  ./chose -file=文件名 -target=目标数字
  file 文件名缺省 values.txt 一个浮点数一行
  target 目标值 缺省 500.0
 */
import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "sort"
    "flag"
)
func readPrices(filename string) ([]float64, error) {
    //f, err := os.Open("values.txt")
    f, err := os.Open(filename)
    if err!= nil {
        return nil, err
    }
    defer f.Close()

    var values []float64
    var curLine string
    scanner := bufio.NewScanner(f);
    for scanner.Scan() {
        curLine = scanner.Text()
        v,perr := strconv.ParseFloat(curLine, 64)
        if perr != nil {
            continue
        }
        values = append(values, v)
    }
    sort.Sort(sort.Reverse(sort.Float64Slice(values)))
    return values, nil
}
type Combine struct {
    sum float64
    elements []float64
    indexes []int
}
func minSumGe(target float64, offset int, values []float64) *Combine {
    //fmt.Printf("  >> process offset : %d \n", offset)
    if offset >= len(values) {
        return nil
    }
    var result Combine
    if offset == len(values) -1 {
        result.sum = values[offset]
        result.elements = values[offset:]
        fmt.Printf(">>> %d %s",offset, result)
        return &result
    }

    var sub *Combine
    if values[offset] >= target {
        sub = minSumGe(target, offset +1 , values)
        if nil == sub || sub.sum > values[offset ]{
            result.sum = values[offset]
            result.elements = values[offset:offset+1]
            //fmt.Printf(">>> %d %s\n",offset, result)
            return &result
        }
        return sub
    } else {
        left := target - values[offset]
        if left > sumArr[offset+1] {
            return nil
        }
        sub = minSumGe(target - values[offset], offset+1, values)
        if sub == nil {
            return nil
        }
        result.sum = values[offset] + sub.sum
        result.elements = append(sub.elements, values[offset])
    }

    return &result
}
var sumArr []float64
func main() {
    target :=flag.Float64("target", 500.0, "Target value")
    filename := flag.String("file", "values.txt", "Value source file")
    flag.Parse()
    fmt.Printf("File : %s , target: %f \n", *filename, *target)
    values,err := readPrices(*filename)
    if err != nil {
        fmt.Errorf("Failed to read values\n")
    }
    sumArr = make([]float64, len(values))
    var sum float64 = 0
    // build sum array
    for idx:=len(values)-1;idx>=0; idx -- {
        fmt.Printf("%5d : %f\n",idx, values[idx])
        sum += values[idx]
        sumArr[idx]=sum
    }
    result := minSumGe(*target, 0, values)
    if result == nil {
        fmt.Println("No result found")
        return
    }
    fmt.Printf("Chosen sum %s ", result.sum)
    for idx, v := range result.elements {
        fmt.Printf("%d: %f ",idx,v)
    }
    fmt.Println()
}
