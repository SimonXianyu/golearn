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

    fmt.Println(" MyTestController ------")
    var c1 MyTestController
    c1.doSome()
    c1.doSome2()

    fmt.Println(" MyTest2Controller ------")
    var c2 MyTest2Controller
    c2.doSome()
    c2.doSome2()
    fmt.Println(" MyTest2Controller as MyTestController ------")
    c2.MyTestController.doSome()
}

func (*MyTestController) doSome() {
    fmt.Println("Test1 Do Some")
}

func (*MyTestController) doSome2() {
    fmt.Println("Test1 Do Some 2")
}

type MyTest2Controller struct {
    MyTestController
}

func (*MyTest2Controller) doSome() {
    fmt.Println("Test2 Do some")
}