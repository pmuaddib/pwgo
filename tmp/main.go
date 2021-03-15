package main

import (
    "os"
    "text/template"
)

type T struct {
    Name string
    Age int
}

func main() {
    me := T{Name: "Misha", Age: 34}
    t := template.New("person")
    tmpl := `Hello my name is {{.Name}}{{print "\n"}}I am {{.Age}} years old.`
    t, err := t.Parse(tmpl)
    if err != nil {
        panic(err)
    }

    if err := t.Execute(os.Stdout, me); err != nil {
        panic(err)
    }
}
