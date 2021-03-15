package main

import (
    "fmt"
    "net/http"
    "sync"
)

var urls = []string {
    "https://pravda.com.ua",
    "https://censor.net",
    "https://youtube.com",
    "https://programming.guide",
}

//var ch = make(chan []string)
var wg sync.WaitGroup

func main() {
    http.HandleFunc("/", handle)
    http.ListenAndServe(":8000", nil)


}

func handle(resp http.ResponseWriter, req *http.Request) {
    for _, url := range urls {
        wg.Add(1)
        go func(u string) {
            defer wg.Done()
            r, err := http.Get(u)
            //ch <- []string{u}
            if err != nil {
                fmt.Fprintln(resp, err)
                return
            }
            fmt.Fprintln(resp, u + " " + r.Status)
        }(url)
    }
        wg.Wait()
}
