package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "net/url"
    "os"
    "strings"
    "thegoprlang/links"
    "time"
)

func main() {
    worklist := make(chan []string)
    go func() {
        worklist <- os.Args[1:]
        fmt.Println(worklist)
    }()

    mainDomains := make(map[string]bool)

    for _, v := range os.Args[1:] {
        u, err := url.Parse(v)
        if err != nil {
            continue // ignore
        }
        mainDomains[u.Host] = true
    }
    for list := range worklist {
        for _, link := range list {
            fmt.Println(link)
        }
        fmt.Println(list)
    }

    seen := make(map[string]bool)
    for list := range worklist {
      for _, link := range list {
           u, err := url.Parse(link)
           if err != nil {
             log.Println(err)
           }
           if !mainDomains[u.Host] {
             continue
           }
          if !seen[link] {
              seen[link] = true
              go func(link string) {
                 worklist <- crawl(link)
              }(link)
          }
      }
    }

    breadthFirst(crawl, os.Args[1:])
}

func handleConn(c net.Conn) {
    input := bufio.NewScanner(c)
    for input.Scan() {
        go echo(c, input.Text(), 1*time.Second)
    }
    c.Close()
}


func breadthFirst(f func(item string) []string, worklist []string) {
    seen := make(map[string]bool)
    mainDomains := make(map[string]bool)
    for _, v := range worklist {
        u, err := url.Parse(v)
        if err != nil {
            continue // ignore
        }
        mainDomains[u.Host] = true
    }
    for len(worklist) > 0 {
        items := worklist
        worklist = nil
        for _, item := range items {
            u, err := url.Parse(item)
            if err != nil {
                log.Println(err)
            }
            if !mainDomains[u.Host] {
                continue
            }
            if !seen[item] {
                seen[item] = true
                worklist = append(worklist, f(item)...)
            }
        }
    }
}

func crawl(url string) []string {
    fmt.Println(url)
    result, err := links.Extract(url)
    if err != nil {
        log.Print(err)
    }

    return result
}
