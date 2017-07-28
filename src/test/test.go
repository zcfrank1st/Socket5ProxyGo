package main

import (
    "github.com/golang/glog"
    "flag"
    "fmt"
)

func init() {
    flag.Parse()
}

func main()  {
    glog.Info("this is a demo, hello world")
    fmt.Println("hello world")
}