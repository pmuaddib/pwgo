package main

import (
    "bufio"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "strings"
    "time"
)

func main() {
    var locations  = make(map[string]string)
    for _, v := range os.Args {
       loc := strings.Split(v, "=")
       if len(loc) != 2 {
           continue
       }
       locations[loc[0]] = loc[1]
    }
    var conns = make(map[string]net.Conn)
    for n, a := range locations {
      conn, err := net.Dial("tcp", a)
      if err != nil {
          log.Fatal(err)
      }
      conns[n] = conn
    }

    for {
       for n, c := range conns {
           go showTime(n, c)
       }
    }

}

func showTime(location string, conn io.Reader) {
    str, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Fprintf(os.Stdout, "%s = %s\r", location, str)
    time.Sleep(1 * time.Second)
}
