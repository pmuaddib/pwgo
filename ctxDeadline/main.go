package main

import (
    "context"
    "fmt"
    "time"
)

const shortDuration = 2 * time.Second

func main() {
    d := time.Now().Add(shortDuration)
    ctx, cancel := context.WithDeadline(context.Background(), d)

    defer cancel()

    select {
    case <-time.After(time.Second*1):
        fmt.Println("overslept")
    case <-ctx.Done():
        fmt.Println(ctx.Err())
    }
}
