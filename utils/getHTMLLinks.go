package utils

import (
    "fmt"
    "golang.org/x/net/html"
    "log"
    "strings"
)

var htmlTree []string

func GetLinks(htmlData string) []string {
    doc, err := html.Parse(strings.NewReader(htmlData))
    if err != nil {
        log.Fatal(err)
    }

    return visit(nil, doc)
}

func visit(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode {
        htmlTree = append(htmlTree, n.Data)
        fmt.Println(htmlTree)
    }

    if n.Type == html.ElementNode && n.Data == "a" {
        for _, a := range n.Attr {
            if a.Key == "href" {
                links = append(links, a.Val)
            }
        }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
        links = visit(links, c)
    }

    return links
}
