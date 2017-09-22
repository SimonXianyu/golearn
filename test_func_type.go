package main

import (
    "fmt"
    _ "simonxianyu/golearn/mytest"
)


type Sfun func(a *int)

func Test(a *int) {
    fmt.Printf("%03d\n",*a)
}

func Test1(a *int , b *int) {
    fmt.Printf("%04d : %05d\n", *a, *b)
}

func main() {
    num := 4234234
    dd := Sfun(Test)

    dd(&num)

}