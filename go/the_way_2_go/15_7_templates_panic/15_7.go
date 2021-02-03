// little fun with panicking templates
package main

import (
    "fmt"
    "text/template"
)

func main() {
    tmplOK := template.New("ok")
    // no need to panic here, `Must' is called here for its side effects only
    template.Must(tmplOK.Parse("/* some comment */ some static text {{ .Name}}"))
    defer func() {
        if rec := recover(); rec != nil {
            fmt.Println("`template.Must' recovered in `main()'")
        }
    }()
    // if we are here, that means previous `Must' succeded
    fmt.Println("tmplOK parsed OK.")
    fmt.Println("Next one ought to fail.")
    tmplErr := template.New("error")
    // intentional template error: missing closing `}'
    template.Must(tmplErr.Parse("some static text {{ .Name}"))
}
