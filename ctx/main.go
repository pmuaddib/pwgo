package main

import (
    "context"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "thegoprlang/ctx/google"
    "thegoprlang/ctx/userip"
    "time"
)

func main()  {
    log.Println("Starting server")
    http.HandleFunc("/search", handleSearch)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSearch(w http.ResponseWriter, req *http.Request) {
    var (
        ctx context.Context
        cancel context.CancelFunc
    )
    timeout, err := time.ParseDuration(req.FormValue("timeout"))
    if err == nil {
        ctx, cancel = context.WithTimeout(context.Background(), timeout)
    } else {
        ctx, cancel = context.WithCancel(context.Background())
    }
    defer cancel()
    query := req.FormValue("q")
    if query == "" {
        http.Error(w, "Query empty", http.StatusBadRequest)
        return
    }
    userIP, err := userip.FromRequest(req)
    if err != nil {
        http.Error(w, "no IP", http.StatusBadRequest)
        return
    }
    ctx = userip.NewContext(ctx, userIP)
    start := time.Now()
    result, err := google.Search(ctx, query)
    elapsed := time.Since(start)
    if err != nil {
        http.Error(w, "Bad search", http.StatusBadRequest)
        fmt.Println(err)
        return
    }
    if err = resultsTemplate.Execute(w, struct {
            Results google.Results
            Timeout, Elapsed time.Duration
        }{
            Results: result,
            Timeout: timeout,
            Elapsed: elapsed,
        }); err != nil {
            log.Print(err)
            return
    }
}
var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))
