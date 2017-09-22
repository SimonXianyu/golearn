package mytest

import (
    "github.com/astaxie/beego"
    "fmt"
)

type MyTestController struct {
    beego.Controller
}

func init() {
    fmt.Println("myTest init executed")
}