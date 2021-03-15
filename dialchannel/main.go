package main

import (
    "fmt"
    "io"
    "log"
    "net"
    "os"
)

func main() {
    tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        log.Fatal(err)
    }
    ch := make(chan struct{})
    go func() {
        io.Copy(os.Stdout, conn)
        fmt.Println("Done")
        ch <- struct{}{}
    }()

    mustCopy(conn, os.Stdin)

    conn.CloseWrite()
    <-ch
}

func mustCopy(dst io.Writer, rsv io.Reader)  {
    if _, err := io.Copy(dst, rsv); err != nil {
        if err == io.EOF {
            return
        }
        log.Fatal(err)
    }
}
