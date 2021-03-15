package google

import (
    "context"
    "encoding/json"
    "fmt"
    "golang.org/x/net/html"
    "log"
    "net/http"
    "strings"
    "thegoprlang/ctx/userip"
)

type Results []Result

type Result struct {
    Title, URL string
}

const jsonReader = `
{
   "ResponseData":{
      "Results":[
        {
           "Title":"Title1",
           "URL":"https:example.com/title1"
        },
        {
           "Title":"Title2",
           "URL":"https:example.com/title2"
        },
        {
           "Title":"Title3",
           "URL":"https:example.com/title3"
        },
        {
           "Title":"Title4",
           "URL":"https:example.com/title4"
        }
      ]
  }
}
`

func Search(ctx context.Context, query string) (Results, error) {
    var data struct{
        ResponseData struct{
            Results[]struct{
                Title string
                URL string
            }
        }
    }
    req, err := http.NewRequest("GET", "https://www.pravda.com.ua", nil)
    if err != nil {
        return nil, err
    }
    q := req.URL.Query()
    q.Set("search_query", query)

    if userIp, ok := userip.FromContext(ctx); ok {
        q.Set("user_ip", userIp.String())
        log.Print(userIp.String())
    }

    req.URL.RawQuery = q.Encode()

    var results Results
    err = doRequest(ctx, req, func(resp *http.Response, err error) error {
        if err != nil {
            return err
        }
        defer resp.Body.Close()
        doc, err := html.Parse(resp.Body)
        if err != nil {
            log.Print(err)
        } else {
            visitNode := func(n *html.Node) {
                if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
                    fmt.Printf("Title str q: %[1]s, %[1]q\n", n.FirstChild.Data)
                }
            }
            forEachNode(doc, visitNode)
        }

        fmt.Printf("Server answer code %d\n", resp.StatusCode)
        if err := json.NewDecoder(strings.NewReader(jsonReader)).Decode(&data); err != nil {
            return err
        }

        for _, res := range data.ResponseData.Results {
            results = append(results, Result{Title: res.Title, URL: res.URL})
        }
        return nil
    })

    return results, err
}

func doRequest(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
    c := make(chan error, 1)
    req = req.WithContext(ctx)
    go func() {
        c<- f(http.DefaultClient.Do(req))
    }()
    select {
    case <-ctx.Done():
        <-c
        return ctx.Err()
    case err := <-c:
        return err
    }
}

func forEachNode(n *html.Node, pre func(n *html.Node)) {
    if pre != nil {
        pre(n)
    }

    for c:= n.FirstChild; c != nil; c = c.NextSibling {
        forEachNode(c, pre)
    }
}