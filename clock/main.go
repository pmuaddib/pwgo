package main

import (
    "flag"
    "io"
    "log"
    "net"
    "time"
)

var tz = flag.String("tz", "", "Time Zone")
var port = flag.String("p", "8010", "Port number")

func main() {
    flag.Parse()
    listener, err := net.Listen("tcp", "localhost:" + *port)
    if err != nil {
        log.Fatal(err)
    }
    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }
        go handleConn(conn)
    }
}

func handleConn(c net.Conn) {
    loc, err := time.LoadLocation(*tz)
    if err != nil {
        log.Fatal(err)
    }

    for {
        _, err := io.WriteString(c, time.Now().In(loc).Format("Mon Jan 2 15:04:05PM\n"))
        if err != nil {
            log.Print("Client quit")
            return
        }
        time.Sleep(1 * time.Second)
    }
    defer c.Close()
}