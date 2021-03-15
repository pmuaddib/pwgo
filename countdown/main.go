package main

import (
    "fmt"
    "os"
    "time"
)

const limit = 5

func main() {
    abort := make(chan struct{})
    go func() {
        os.Stdin.Read(make([]byte, 1))
        abort<- struct{}{}
    }()
    fmt.Println("Waiting 5 sec, abort if need to stop lunch")
    select {
    case <-abort:
        fmt.Println("Abort")
        return
    case <-time.After(5 * time.Second):
        fmt.Println("LUNCH!!!")
    }
}
