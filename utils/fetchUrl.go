package utils

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func Fetch(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
        os.Exit(1)
    }
    b, err := ioutil.ReadAll(resp.Body)
    resp.Body.Close()
    if err != nil {
        fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
        os.Exit(1)
    }
    return fmt.Sprintf("%s", b)
}
