package main

import (
    "fmt"
    "time"
)

func main() {
    const n = 6
    go spinner(100 * time.Millisecond)
    fmt.Printf("\r%d", fib(n))
}

func spinner(duration time.Duration) {
    for {
        for _, l := range `-\|/` {
            fmt.Printf("\r%c", l)
            time.Sleep(duration)
        }
    }
}

func fib(x int) int {
    if x < 2 {
        return x
    }
    return fib(x - 1) + fib(x - 2)
}