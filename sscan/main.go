package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Type a command:\n")
    for scanner.Scan() {
        input := strings.Fields(scanner.Text())
        if len(input) == 0 {
            continue
        }
        switch input[0] {
        case "hi":
            fmt.Fprint(os.Stdout, "Hello!\n")
        case "b":
            fmt.Fprint(os.Stdout, "Bye!\n")
            return
        default:
            fmt.Fprint(os.Stdout, "Try again...\n")
        }
    }
}
