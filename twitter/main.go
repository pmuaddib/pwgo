package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

var client = http.Client{
    Transport: &http.Transport{

    },
    Timeout: 3 * time.Second,
}

func main() {
    tc := time.Tick(3 * time.Second)
    for {
        status, err := ping("twitter.com/home")
        if err != nil {
            log.Print(err)
        } else {
            fmt.Println(status)
        }
        <-tc
    }
}

func ping(domain string) (int, error) {
    url := "https://" + domain
    req, err := http.NewRequest("HEAD", url, nil)
    if err != nil {
        return 0, err
    }
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    resp.Body.Close()
    return resp.StatusCode, nil
}